//+build stress
// go test -timeout 0 -tags stress -v ./cmd/createorupdate

package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-02-01/resources"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/ghodss/yaml"
	"github.com/sirupsen/logrus"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"

	acsapi "github.com/openshift/openshift-azure/pkg/api"
	"github.com/openshift/openshift-azure/pkg/healthcheck"
	"github.com/openshift/openshift-azure/pkg/plugin"
	"github.com/openshift/openshift-azure/pkg/upgrade"
)

func TestCreate(t *testing.T) {
	err := os.Setenv("RUNNING_UNDER_TEST", "true")
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, acsapi.ContextKeyClientID, os.Getenv("AZURE_CLIENT_ID"))
	ctx = context.WithValue(ctx, acsapi.ContextKeyClientSecret, os.Getenv("AZURE_CLIENT_SECRET"))
	ctx = context.WithValue(ctx, acsapi.ContextKeyTenantID, os.Getenv("AZURE_TENANT_ID"))

	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{FullTimestamp: true}
	logger.SetLevel(logrus.DebugLevel)
	entry := logrus.NewEntry(logger)

	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		t.Fatal(err)
	}

	groups := resources.NewGroupsClient(os.Getenv("AZURE_SUBSCRIPTION_ID"))
	groups.Authorizer = authorizer

	p := plugin.NewPlugin(entry, "")

	location := "eastus"

	var cs *acsapi.OpenShiftManagedCluster
	var azuredeploy []byte

out:
	for {
		name, err := randomString(10)
		if err != nil {
			t.Fatal(err)
		}
		name = "jminter-" + name

		entry.Infof("create %s", name)
		_, err = groups.CreateOrUpdate(ctx, name, resources.Group{Name: &name, Location: &location})
		if err != nil {
			t.Fatal(err)
		}

		cs = &acsapi.OpenShiftManagedCluster{
			Location: location,
			Name:     name,
			Properties: &acsapi.Properties{
				OpenShiftVersion: "v3.10",
				FQDN:             fmt.Sprintf("%s.%s.cloudapp.azure.com", name, location),
				RouterProfiles: []acsapi.RouterProfile{
					{
						Name: "default",
						FQDN: fmt.Sprintf("%s-router.%s.cloudapp.azure.com", name, location),
					},
				},
				AgentPoolProfiles: []acsapi.AgentPoolProfile{
					{
						Name:   "master",
						Count:  3,
						VMSize: "Standard_D2s_v3",
						OSType: "Linux",
						Role:   "master",
					},
					{
						Name:   "infra",
						Count:  2,
						VMSize: "Standard_D2s_v3",
						OSType: "Linux",
						Role:   "infra",
					},
					{
						Name:   "compute",
						Count:  2,
						VMSize: "Standard_D2s_v3",
						OSType: "Linux",
						Role:   "compute",
					},
				},
				AuthProfile: &acsapi.AuthProfile{
					IdentityProviders: []acsapi.IdentityProvider{
						{
							Name: "Azure AD",
							Provider: &acsapi.AADIdentityProvider{
								Kind:     "AADIdentityProvider",
								ClientID: os.Getenv("AZURE_CLIENT_ID"),
								Secret:   os.Getenv("AZURE_CLIENT_SECRET"),
								TenantID: os.Getenv("AZURE_TENANT_ID"),
							},
						},
					},
				},
				ServicePrincipalProfile: &acsapi.ServicePrincipalProfile{
					ClientID: os.Getenv("AZURE_CLIENT_ID"),
					Secret:   os.Getenv("AZURE_CLIENT_SECRET"),
				},
				AzProfile: &acsapi.AzProfile{
					TenantID:       os.Getenv("AZURE_TENANT_ID"),
					SubscriptionID: os.Getenv("AZURE_SUBSCRIPTION_ID"),
					ResourceGroup:  name,
				},
			},
		}

		err = p.GenerateConfig(ctx, cs)
		if err != nil {
			t.Fatal(err)
		}
		azuredeploy, err = p.GenerateARM(ctx, cs, false)
		if err != nil {
			t.Fatal(err)
		}
		err = upgrade.Deploy(ctx, cs, p, azuredeploy)
		if err != nil {
			t.Fatal(err)
		}
		err = p.HealthCheck(ctx, cs)
		if err != nil {
			t.Fatal(err)
		}

		kc, err := healthcheck.NewClientSet(ctx, cs.Config.AdminKubeconfig)
		if err != nil {
			t.Fatal(err)
		}
		entry.Info("wait 20 minutes for PVCs")
		err = wait.Poll(2*time.Second, 20*time.Minute, func() (bool, error) {
			list, err := kc.CoreV1().PersistentVolumeClaims("").List(metav1.ListOptions{})
			if err != nil {
				return false, err
			}

			if len(list.Items) != 2 {
				return false, nil
			}

			for _, pvc := range list.Items {
				if pvc.Status.Phase != v1.ClaimBound {
					return false, nil
				}
			}

			return true, nil
		})
		switch {
		case err == wait.ErrWaitTimeout:
			entry.Info("timed out")
			break out
		case err == nil:
		default:
			t.Fatal(err)
		}

		entry.Info("deleting group")
		_, err = groups.Delete(ctx, name)
		if err != nil {
			t.Fatal(err)
		}
	}

	err = os.MkdirAll("../../_data/_out", 0777)
	if err != nil {
		t.Fatal(err)
	}
	b, err := yaml.Marshal(cs)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile("../../_data/containerservice.yaml", b, 0666)
	if err != nil {
		t.Fatal(err)
	}
	oc := acsapi.ConvertToV20180930preview(cs)
	b, err = yaml.Marshal(oc)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile("../../_data/manifest.yaml", b, 0666)
	if err != nil {
		t.Fatal(err)
	}
	err = writeHelpers(cs.Config, azuredeploy)
	if err != nil {
		t.Fatal(err)
	}
}

func randomString(length int) (string, error) {
	const letterBytes = "abcdefghijklmnopqrstuvwxyz0123456789"

	b := make([]byte, length)
	for i := range b {
		o, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterBytes))))
		if err != nil {
			return "", err
		}
		b[i] = letterBytes[o.Int64()]
	}

	return string(b), nil
}
