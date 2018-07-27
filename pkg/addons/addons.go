package addons

//go:generate go get github.com/go-bindata/go-bindata/go-bindata
//go:generate go-bindata -nometadata -pkg $GOPACKAGE -prefix data data/...
//go:generate gofmt -s -l -w bindata.go

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"sort"
	"time"

	acsapi "github.com/Azure/acs-engine/pkg/api"
	"github.com/ghodss/yaml"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
	apiregistrationv1 "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1"

	"github.com/openshift/openshift-azure/pkg/config"
	"github.com/openshift/openshift-azure/pkg/jsonpath"
	"github.com/openshift/openshift-azure/pkg/util"
)

func unmarshal(b []byte) (unstructured.Unstructured, error) {
	// can't use straight yaml.Unmarshal() because it universally mangles yaml
	// integers into float64s, whereas the Kubernetes client library uses int64s
	// wherever it can.  Such a difference can cause us to update objects when
	// we don't actually need to.
	json, err := yaml.YAMLToJSON(b)
	if err != nil {
		return unstructured.Unstructured{}, err
	}

	var o unstructured.Unstructured
	_, _, err = unstructured.UnstructuredJSONScheme.Decode(json, nil, &o)
	if err != nil {
		return unstructured.Unstructured{}, err
	}

	return o, nil
}

// readDB reads previously exported objects into a map via go-bindata as well as
// populating configuration items via Translate().
func readDB(cs *acsapi.ContainerService, c *config.Config) (map[string]unstructured.Unstructured, error) {
	db := map[string]unstructured.Unstructured{}

	for _, asset := range AssetNames() {
		b, err := Asset(asset)
		if err != nil {
			return nil, err
		}

		o, err := unmarshal(b)
		if err != nil {
			return nil, err
		}

		ts := Translations[KeyFunc(o.GroupVersionKind().GroupKind(), o.GetNamespace(), o.GetName())]
		for _, tr := range ts {
			b, err := util.Template(tr.Template, nil, cs, c, nil)
			if err != nil {
				return nil, err
			}

			err = Translate(o.Object, tr.Path, tr.NestedPath, tr.NestedFlags, string(b))
			if err != nil {
				return nil, err
			}
		}

		db[KeyFunc(o.GroupVersionKind().GroupKind(), o.GetNamespace(), o.GetName())] = o
	}

	return db, nil
}

// syncWorkloadsConfig iterates over all workload controllers (deployments,
// daemonsets, statefulsets), walks their volumes, and updates their pod
// templates with annotations that include the hashes of the content for
// each configmap or secret.
func syncWorkloadsConfig(db map[string]unstructured.Unstructured) error {
	// map config resources to their hashed content
	configToHash := make(map[string][]byte)
	for _, o := range db {
		if o.GroupVersionKind().Kind != "Secret" &&
			o.GroupVersionKind().Kind != "ConfigMap" {
			continue
		}

		h := sha256.New()
		for _, v := range jsonpath.MustCompile("$.data").Get(o.Object) {
			// NOTE: this relies on the fact that %#v on a map sorts by key
			fmt.Fprintf(h, "%#v", v)
		}
		for _, v := range jsonpath.MustCompile("$.stringData").Get(o.Object) {
			fmt.Fprintf(h, "%#v", v)
		}
		configToHash[KeyFunc(o.GroupVersionKind().GroupKind(), o.GetNamespace(), o.GetName())] = h.Sum(nil)
	}

	secretGk := schema.GroupKind{Kind: "Secret"}
	configMapGk := schema.GroupKind{Kind: "ConfigMap"}

	// iterate over all workload controllers and add annotations with the hashes
	// of every config map or secret appropriately to force redeployments on config
	// updates.
	for _, o := range db {
		if o.GroupVersionKind().Kind != "DaemonSet" &&
			o.GroupVersionKind().Kind != "Deployment" &&
			o.GroupVersionKind().Kind != "StatefulSet" {
			continue
		}

		volumes := jsonpath.MustCompile("$.spec.template.spec.volumes.*").MustGetObject(o.Object)

		if secretData, found := volumes["secret"]; found {
			secretName := jsonpath.MustCompile("$.secretName").MustGetString(secretData)
			key := fmt.Sprintf("checksum/secret-%s", secretName)
			secretKey := KeyFunc(secretGk, o.GetNamespace(), secretName)
			if hash, found := configToHash[secretKey]; found {
				setPodTemplateAnnotation(key, base64.StdEncoding.EncodeToString(hash), o)
			}
		}

		if configMapData, found := volumes["configMap"]; found {
			configMapName := jsonpath.MustCompile("$.name").MustGetString(configMapData)
			key := fmt.Sprintf("checksum/configmap-%s", configMapName)
			configMapKey := KeyFunc(configMapGk, o.GetNamespace(), configMapName)
			if hash, found := configToHash[configMapKey]; found {
				setPodTemplateAnnotation(key, base64.StdEncoding.EncodeToString(hash), o)
			}
		}
	}

	return nil
}

// setPodTemplateAnnotation sets the provided key-value pair as an annotation
// inside the provided object's pod template.
func setPodTemplateAnnotation(key, value string, o unstructured.Unstructured) {
	annotations, _, _ := unstructured.NestedStringMap(o.Object, "spec", "template", "metadata", "annotations")
	if annotations == nil {
		annotations = make(map[string]string)
	}
	annotations[key] = value
	unstructured.SetNestedStringMap(o.Object, annotations, "spec", "template", "metadata", "annotations")
}

// resource filters
var (
	nsFilter = func(o unstructured.Unstructured) bool {
		return o.GroupVersionKind().Kind == "Namespace"
	}
	saFilter = func(o unstructured.Unstructured) bool {
		return o.GroupVersionKind().Kind == "ServiceAccount"
	}
	cfgFilter = func(o unstructured.Unstructured) bool {
		return o.GroupVersionKind().Kind == "Secret" || o.GroupVersionKind().Kind == "ConfigMap"
	}
	nonScFilter = func(o unstructured.Unstructured) bool {
		return o.GroupVersionKind().Group != "servicecatalog.k8s.io"
	}
	scFilter = func(o unstructured.Unstructured) bool {
		return o.GroupVersionKind().Group == "servicecatalog.k8s.io"
	}
)

// writeDB uses the discovery and dynamic clients to synchronise an API server's
// objects with db.
// TODO: need to implement deleting objects which we don't want any more.
func writeDB(client *client, db map[string]unstructured.Unstructured) error {
	// impose an order to improve debuggability.
	var keys []string
	for k := range db {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// namespaces must exist before namespaced objects.
	if err := client.createResources(nsFilter, db, keys); err != nil {
		return err
	}
	// create serviceaccounts
	if err := client.createResources(saFilter, db, keys); err != nil {
		return err
	}
	// create all secrets and configmaps
	if err := client.createResources(cfgFilter, db, keys); err != nil {
		return err
	}
	// create all non-service catalog resources
	if err := client.createResources(nonScFilter, db, keys); err != nil {
		return err
	}

	// wait for the service catalog api extension to arrive. TODO: we should do
	// this dynamically, and should not PollInfinite.
	err := wait.PollInfinite(time.Second, func() (bool, error) {
		svc, err := client.ac.ApiregistrationV1().APIServices().Get("v1beta1.servicecatalog.k8s.io", metav1.GetOptions{})
		switch {
		case kerrors.IsNotFound(err):
			return false, nil
		case err != nil:
			return false, err
		}
		for _, cond := range svc.Status.Conditions {
			if cond.Type == apiregistrationv1.Available &&
				cond.Status == apiregistrationv1.ConditionTrue {
				return true, nil
			}
		}
		return false, nil
	})
	if err != nil {
		return err
	}

	// refresh dynamic client
	if err := client.updateDynamicClient(); err != nil {
		return err
	}

	// now write the servicecatalog configurables.
	return client.createResources(scFilter, db, keys)
}

func Main(cs *acsapi.ContainerService, c *config.Config, dryRun bool) error {
	client, err := newClient(dryRun)
	if err != nil {
		return err
	}

	db, err := readDB(cs, c)
	if err != nil {
		return err
	}

	if err := syncWorkloadsConfig(db); err != nil {
		return err
	}

	if dryRun {
		// impose an order to improve debuggability.
		var keys []string
		for k := range db {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			b, err := yaml.Marshal(db[k].Object)
			if err != nil {
				return err
			}

			fmt.Println(string(b))
		}

		return nil
	}

	return writeDB(client, db)
}
