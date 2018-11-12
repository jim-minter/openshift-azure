package api

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/openshift/openshift-azure/pkg/util/structtags"
	"github.com/openshift/openshift-azure/test/util/populate"
)

var marshalled = []byte(`{
	"plan": {
		"name": "Plan.Name",
		"product": "Plan.Product",
		"promotionCode": "Plan.PromotionCode",
		"publisher": "Plan.Publisher"
	},
	"properties": {
		"provisioningState": "Properties.ProvisioningState",
		"openShiftVersion": "Properties.OpenShiftVersion",
		"publicHostname": "Properties.PublicHostname",
		"fqdn": "Properties.FQDN",
		"networkProfile": {
			"vnetCidr": "Properties.NetworkProfile.VnetCIDR",
			"peerVnetId": "Properties.NetworkProfile.PeerVnetID"
		},
		"routerProfiles": [
			{
				"name": "Properties.RouterProfiles[0].Name",
				"publicSubdomain": "Properties.RouterProfiles[0].PublicSubdomain",
				"fqdn": "Properties.RouterProfiles[0].FQDN"
			}
		],
		"agentPoolProfiles": [
			{
				"name": "Properties.AgentPoolProfiles[0].Name",
				"count": 1,
				"vmSize": "Properties.AgentPoolProfiles[0].VMSize",
				"subnetCidr": "Properties.AgentPoolProfiles[0].SubnetCIDR",
				"osType": "Properties.AgentPoolProfiles[0].OSType",
				"role": "Properties.AgentPoolProfiles[0].Role"
			}
		],
		"authProfile": {
			"identityProviders": [
				{
					"name": "Properties.AuthProfile.IdentityProviders[0].Name",
					"provider": {
						"kind": "AADIdentityProvider",
						"clientId": "Properties.AuthProfile.IdentityProviders[0].Provider.ClientID",
						"secret": "Properties.AuthProfile.IdentityProviders[0].Provider.Secret",
						"tenantId": "Properties.AuthProfile.IdentityProviders[0].Provider.TenantID"
					}
				}
			]
		},
		"servicePrincipalProfile": {
			"clientId": "Properties.ServicePrincipalProfile.ClientID",
			"secret": "Properties.ServicePrincipalProfile.Secret"
		},
		"azProfile": {
			"tenantId": "Properties.AzProfile.TenantID",
			"subscriptionId": "Properties.AzProfile.SubscriptionID",
			"resourceGroup": "Properties.AzProfile.ResourceGroup"
		}
	},
	"id": "ID",
	"name": "Name",
	"type": "Type",
	"location": "Location",
	"tags": {
		"Tags.key": "Tags.val"
	},
	"config": {
		"imageOffer": "Config.ImageOffer",
		"imagePublisher": "Config.ImagePublisher",
		"imageSku": "Config.ImageSKU",
		"imageVersion": "Config.ImageVersion",
		"sshKey": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==",
		"configStorageAccount": "Config.ConfigStorageAccount",
		"registryStorageAccount": "Config.RegistryStorageAccount",
		"loggingWorkspace": "Config.LoggingWorkspace",
		"loggingLocation": "Config.LoggingLocation",
		"certificates": {
			"etcdCa": {
				"key": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==",
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
			},
			"ca": {
				"key": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==",
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
			},
			"frontProxyCa": {
				"key": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==",
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
			},
			"serviceSigningCa": {
				"key": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==",
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
			},
			"serviceCatalogCa": {
				"key": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==",
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
			},
			"etcdServer": {
				"key": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==",
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
			},
			"etcdPeer": {
				"key": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==",
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
			},
			"etcdClient": {
				"key": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==",
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
			},
			"masterServer": {
				"key": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==",
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
			},
			"openshiftConsole": {
				"key": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==",
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
			},
			"admin": {
				"key": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==",
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
			},
			"aggregatorFrontProxy": {
				"key": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==",
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
			},
			"masterKubeletClient": {
				"key": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==",
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
			},
			"masterProxyClient": {
				"key": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==",
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
			},
			"openShiftMaster": {
				"key": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==",
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
			},
			"nodeBootstrap": {
				"key": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==",
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
			},
			"registry": {
				"key": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==",
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
			},
			"router": {
				"key": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==",
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
			},
			"serviceCatalogServer": {
				"key": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==",
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
			},
			"serviceCatalogAPIClient": {
				"key": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==",
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
			},
			"azureClusterReader": {
				"key": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==",
				"cert": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K"
			}
		},
		"images": {
			"format": "Config.Images.Format",
			"masterEtcd": "Config.Images.MasterEtcd",
			"controlPlane": "Config.Images.ControlPlane",
			"node": "Config.Images.Node",
			"serviceCatalog": "Config.Images.ServiceCatalog",
			"sync": "Config.Images.Sync",
			"templateServiceBroker": "Config.Images.TemplateServiceBroker",
			"prometheusNodeExporter": "Config.Images.PrometheusNodeExporter",
			"registry": "Config.Images.Registry",
			"router": "Config.Images.Router",
			"registryConsole": "Config.Images.RegistryConsole",
			"ansibleServiceBroker": "Config.Images.AnsibleServiceBroker",
			"webConsole": "Config.Images.WebConsole",
			"oAuthProxy": "Config.Images.OAuthProxy",
			"prometheus": "Config.Images.Prometheus",
			"prometheusAlertBuffer": "Config.Images.PrometheusAlertBuffer",
			"prometheusAlertManager": "Config.Images.PrometheusAlertManager",
			"logBridge": "Config.Images.LogBridge",
			"etcdOperator": "Config.Images.EtcdOperator",
			"kubeStateMetrics": "Config.Images.KubeStateMetrics",
			"addonsResizer": "Config.Images.AddonsResizer"
		},
		"adminKubeconfig": "eyJwcmVmZXJlbmNlcyI6e30sImNsdXN0ZXJzIjpudWxsLCJ1c2VycyI6bnVsbCwiY29udGV4dHMiOm51bGwsImN1cnJlbnQtY29udGV4dCI6IiJ9",
		"masterKubeconfig": "eyJwcmVmZXJlbmNlcyI6e30sImNsdXN0ZXJzIjpudWxsLCJ1c2VycyI6bnVsbCwiY29udGV4dHMiOm51bGwsImN1cnJlbnQtY29udGV4dCI6IiJ9",
		"nodeBootstrapKubeconfig": "eyJwcmVmZXJlbmNlcyI6e30sImNsdXN0ZXJzIjpudWxsLCJ1c2VycyI6bnVsbCwiY29udGV4dHMiOm51bGwsImN1cnJlbnQtY29udGV4dCI6IiJ9",
		"azureClusterReaderKubeconfig": "eyJwcmVmZXJlbmNlcyI6e30sImNsdXN0ZXJzIjpudWxsLCJ1c2VycyI6bnVsbCwiY29udGV4dHMiOm51bGwsImN1cnJlbnQtY29udGV4dCI6IiJ9",
		"serviceAccountKey": "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==",
		"sessionSecretAuth": "Q29uZmlnLlNlc3Npb25TZWNyZXRBdXRo",
		"sessionSecretEnc": "Q29uZmlnLlNlc3Npb25TZWNyZXRFbmM=",
		"runningUnderTest": true,
		"htPasswd": "Q29uZmlnLkh0UGFzc3dk",
		"customerAdminPasswd": "Config.CustomerAdminPasswd",
		"customerReaderPasswd": "Config.CustomerReaderPasswd",
		"endUserPasswd": "Config.EndUserPasswd",
		"registryHttpSecret": "Q29uZmlnLlJlZ2lzdHJ5SFRUUFNlY3JldA==",
		"prometheusProxySessionSecret": "Q29uZmlnLlByb21ldGhldXNQcm94eVNlc3Npb25TZWNyZXQ=",
		"alertManagerProxySessionSecret": "Q29uZmlnLkFsZXJ0TWFuYWdlclByb3h5U2Vzc2lvblNlY3JldA==",
		"alertsProxySessionSecret": "Q29uZmlnLkFsZXJ0c1Byb3h5U2Vzc2lvblNlY3JldA==",
		"registryConsoleOAuthSecret": "Config.RegistryConsoleOAuthSecret",
		"routerStatsPassword": "Config.RouterStatsPassword",
		"serviceCatalogClusterId": "01010101-0101-0101-0101-010101010101"
	}
}`)

func TestMarshal(t *testing.T) {
	mogrify := func(v reflect.Value) {
		switch v.Interface().(type) {
		case []IdentityProvider:
			// set the Provider to AADIdentityProvider
			v.Set(reflect.ValueOf([]IdentityProvider{{Provider: &AADIdentityProvider{Kind: "AADIdentityProvider"}}}))
		}
	}

	populatedOc := OpenShiftManagedCluster{}
	populate.Walk(&populatedOc, mogrify)

	b, err := json.MarshalIndent(populatedOc, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(b, marshalled) {
		t.Errorf("json.MarshalIndent returned unexpected result\n%s\n", string(b))
	}
}

func TestUnmarshal(t *testing.T) {
	mogrify := func(v reflect.Value) {
		switch v.Interface().(type) {
		case []IdentityProvider:
			// set the Provider to AADIdentityProvider
			v.Set(reflect.ValueOf([]IdentityProvider{{Provider: &AADIdentityProvider{Kind: "AADIdentityProvider"}}}))
		}
	}

	t.Skip("skipping this test until we are able to unmarshall x509.Certificate")
	populatedOc := OpenShiftManagedCluster{}
	populate.Walk(&populatedOc, mogrify)

	var unmarshalledOc OpenShiftManagedCluster
	err := json.Unmarshal(marshalled, &unmarshalledOc)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(populatedOc, unmarshalledOc) {
		t.Errorf("json.Unmarshal returned unexpected result\n%#v\n", unmarshalledOc)
	}
}

// TestJSONTags ensures that all the `json:"..."` struct field tags under
// OpenShiftManagedCluster correspond with their field names
func TestJSONTags(t *testing.T) {
	o := OpenShiftManagedCluster{}
	for _, err := range structtags.CheckJsonTags(o) {
		t.Errorf("mismatch in struct tags for %T: %s", o, err.Error())
	}
}
