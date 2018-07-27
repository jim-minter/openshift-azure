package addons

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/ghodss/yaml"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestClean(t *testing.T) {
	tests := []struct {
		name    string
		index   int
		objects []string
		input   unstructured.Unstructured
		want    unstructured.Unstructured
	}{
		{
			name:    "Status clean",
			index:   1,
			objects: []string{"Pod"},
		},
		{
			name:    "DaemonSet.apps/Deployments/DeploymentConfigs clean",
			index:   2,
			objects: []string{"DaemonSet.apps", "Deployment.apps", "DeploymentConfig.apps.openshift.io"},
		},
		{
			name:    "Namespace clean",
			index:   3,
			objects: []string{"Namespace"},
		},
		{
			name:    "ImageStream.image.openshift.io clean",
			index:   4,
			objects: []string{"ImageStream.image.openshift.io"},
		},
		{
			name:    "Secret clean",
			index:   5,
			objects: []string{"Secret"},
		},
		{
			name:    "Secret convert",
			index:   6,
			objects: []string{"Secret"},
		},
		{
			name:    "Service clean",
			index:   7,
			objects: []string{"Service"},
		},
	}

	for _, tt := range tests {
		dir, err := os.Getwd()
		if err != nil {
			t.Errorf("Test failed: %v", err)
		}
		// input parsing
		input, err := ioutil.ReadFile(fmt.Sprintf("%s/../../tests/testdata/pkg/addons/clean/%02d-input.yaml", dir, tt.index))
		if err != nil {
			t.Errorf("Test failed: %v", err)
		}
		jsonInput, err := yaml.YAMLToJSON(input)
		if err != nil {
			t.Errorf("Test failed: %v", err)
		}
		_, _, err = unstructured.UnstructuredJSONScheme.Decode(jsonInput, nil, &tt.input)
		if err != nil {
			t.Errorf("Test failed: %v", err)
		}

		// want parsing
		want, err := ioutil.ReadFile(fmt.Sprintf("%s/../../tests/testdata/pkg/addons/clean/%02d-want.yaml", dir, tt.index))
		if err != nil {
			t.Errorf("Test failed: %v", err)
		}
		jsonWant, err := yaml.YAMLToJSON(want)
		if err != nil {
			t.Errorf("Test failed: %v", err)
		}
		_, _, err = unstructured.UnstructuredJSONScheme.Decode(jsonWant, nil, &tt.want)
		if err != nil {
			t.Errorf("Test failed: %v", err)
		}

		for _, obj := range tt.objects {
			tt.input.Object["kind"] = obj
			tt.want.Object["kind"] = obj
			Clean(tt.input)
			if !reflect.DeepEqual(tt.input, tt.want) {
				t.Errorf("fail: %#v \n%#v\n%#v", tt.name, tt.input, tt.want)
			}
		}
	}
}
