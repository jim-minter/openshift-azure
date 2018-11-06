package api_test

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ghodss/yaml"

	. "github.com/openshift/openshift-azure/pkg/api"
	"github.com/openshift/openshift-azure/pkg/util/structtags"
	"github.com/openshift/openshift-azure/test/util/populate"
)

var unmarshalled = &OpenShiftManagedCluster{
	ID:       "id",
	Location: "location",
	Name:     "name",
	Plan: &ResourcePurchasePlan{
		Name:          "plan.name",
		Product:       "plan.product",
		PromotionCode: "plan.promotionCode",
		Publisher:     "plan.publisher",
	},
	Tags: map[string]string{
		"tags.k1": "v1",
		"tags.k2": "v2",
	},
	Type: "type",
	Properties: &Properties{
		ProvisioningState: "properties.provisioningState",
		OpenShiftVersion:  "properties.openShiftVersion",
		PublicHostname:    "properties.publicHostname",
		NetworkProfile: &NetworkProfile{
			VnetCIDR:   "properties.networkProfile.vnetCidr",
			PeerVnetID: "properties.networkProfile.peerVnetId",
		},
		RouterProfiles: []RouterProfile{
			{
				Name:            "properties.routerProfiles.0.name",
				PublicSubdomain: "properties.routerProfiles.0.publicSubdomain",
				FQDN:            "properties.routerProfiles.0.fqdn",
			},
			{
				Name:            "properties.routerProfiles.1.name",
				PublicSubdomain: "properties.routerProfiles.1.publicSubdomain",
				FQDN:            "properties.routerProfiles.1.fqdn",
			},
		},
		FQDN: "properties.fqdn",
		AuthProfile: &AuthProfile{
			IdentityProviders: []IdentityProvider{
				{
					Name: "properties.authProfile.identityProviders.0.name",
					Provider: &AADIdentityProvider{
						Kind:     "AADIdentityProvider",
						ClientID: "properties.authProfile.identityProviders.0.provider.clientId",
						Secret:   "properties.authProfile.identityProviders.0.provider.secret",
						TenantID: "properties.authProfile.identityProviders.0.provider.tenantId",
					},
				},
			},
		},
		AgentPoolProfiles: []AgentPoolProfile{
			{
				Count:      1,
				VMSize:     "properties.agentPoolProfiles.0.vmSize",
				SubnetCIDR: "properties.agentPoolProfiles.0.subnetCidr",
				Role:       "properties.agentPoolProfiles.0.role",
			},
			{
				Name:       "properties.agentPoolProfiles.0.name",
				Count:      2,
				VMSize:     "properties.agentPoolProfiles.0.vmSize",
				SubnetCIDR: "properties.agentPoolProfiles.0.subnetCidr",
				OSType:     "properties.agentPoolProfiles.0.osType",
				Role:       "properties.agentPoolProfiles.0.role",
			},
			{
				Name:       "properties.agentPoolProfiles.0.name",
				Count:      1,
				VMSize:     "properties.agentPoolProfiles.0.vmSize",
				SubnetCIDR: "properties.agentPoolProfiles.0.subnetCidr",
				OSType:     "properties.agentPoolProfiles.0.osType",
				Role:       "master",
			},
		},
		ServicePrincipalProfile: &ServicePrincipalProfile{
			ClientID: "properties.servicePrincipalProfile.clientId",
			Secret:   "properties.servicePrincipalProfile.secret",
		},
	},
}

var marshalled = []byte(`{
	"plan": {
		"name": "plan.name",
		"product": "plan.product",
		"promotionCode": "plan.promotionCode",
		"publisher": "plan.publisher"
	},
	"properties": {
		"provisioningState": "properties.provisioningState",
		"openShiftVersion": "properties.openShiftVersion",
		"publicHostname": "properties.publicHostname",
		"fqdn": "properties.fqdn",
		"networkProfile": {
			"vnetCidr": "properties.networkProfile.vnetCidr",
			"peerVnetId": "properties.networkProfile.peerVnetId"
		},
		"routerProfiles": [
			{
				"name": "properties.routerProfiles.0.name",
				"publicSubdomain": "properties.routerProfiles.0.publicSubdomain",
				"fqdn": "properties.routerProfiles.0.fqdn"
			},
			{
				"name": "properties.routerProfiles.1.name",
				"publicSubdomain": "properties.routerProfiles.1.publicSubdomain",
				"fqdn": "properties.routerProfiles.1.fqdn"
			}
		],
		"agentPoolProfiles": [
			{
				"count": 1,
				"vmSize": "properties.agentPoolProfiles.0.vmSize",
				"subnetCidr": "properties.agentPoolProfiles.0.subnetCidr",
				"role": "properties.agentPoolProfiles.0.role"
			},
			{
				"name": "properties.agentPoolProfiles.0.name",
				"count": 2,
				"vmSize": "properties.agentPoolProfiles.0.vmSize",
				"subnetCidr": "properties.agentPoolProfiles.0.subnetCidr",
				"osType": "properties.agentPoolProfiles.0.osType",
				"role": "properties.agentPoolProfiles.0.role"
			},
			{
				"name": "properties.agentPoolProfiles.0.name",
				"count": 1,
				"vmSize": "properties.agentPoolProfiles.0.vmSize",
				"subnetCidr": "properties.agentPoolProfiles.0.subnetCidr",
				"osType": "properties.agentPoolProfiles.0.osType",
				"role": "master"
			}
		],
		"authProfile": {
			"identityProviders": [
				{
					"name": "properties.authProfile.identityProviders.0.name",
					"provider": {
						"kind": "AADIdentityProvider",
						"clientId": "properties.authProfile.identityProviders.0.provider.clientId",
						"secret": "properties.authProfile.identityProviders.0.provider.secret",
						"tenantId": "properties.authProfile.identityProviders.0.provider.tenantId"
					}
				}
			]
		},
		"servicePrincipalProfile": {
			"clientId": "properties.servicePrincipalProfile.clientId",
			"secret": "properties.servicePrincipalProfile.secret"
		}
	},
	"id": "id",
	"name": "name",
	"type": "type",
	"location": "location",
	"tags": {
		"tags.k1": "v1",
		"tags.k2": "v2"
	}
}`)

var marshalledYaml = []byte(`config:
  adminKubeconfig: eyJwcmVmZXJlbmNlcyI6e30sImNsdXN0ZXJzIjpudWxsLCJ1c2VycyI6bnVsbCwiY29udGV4dHMiOm51bGwsImN1cnJlbnQtY29udGV4dCI6IiJ9
  alertManagerProxySessionSecret: Q29uZmlnLkFsZXJ0TWFuYWdlclByb3h5U2Vzc2lvblNlY3JldA==
  alertsProxySessionSecret: Q29uZmlnLkFsZXJ0c1Byb3h5U2Vzc2lvblNlY3JldA==
  azureClusterReaderKubeconfig: eyJwcmVmZXJlbmNlcyI6e30sImNsdXN0ZXJzIjpudWxsLCJ1c2VycyI6bnVsbCwiY29udGV4dHMiOm51bGwsImN1cnJlbnQtY29udGV4dCI6IiJ9
  certificates:
    admin:
      cert: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
      key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
    aggregatorFrontProxy:
      cert: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
      key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
    azureClusterReader:
      cert: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
      key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
    ca:
      cert: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
      key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
    etcdCa:
      cert: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
      key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
    etcdClient:
      cert: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
      key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
    etcdPeer:
      cert: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
      key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
    etcdServer:
      cert: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
      key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
    frontProxyCa:
      cert: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
      key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
    masterKubeletClient:
      cert: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
      key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
    masterProxyClient:
      cert: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
      key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
    masterServer:
      cert: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
      key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
    nodeBootstrap:
      cert: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
      key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
    openShiftMaster:
      cert: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
      key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
    openshiftConsole:
      cert: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
      key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
    registry:
      cert: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
      key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
    router:
      cert: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
      key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
    serviceCatalogAPIClient:
      cert: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
      key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
    serviceCatalogCa:
      cert: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
      key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
    serviceCatalogServer:
      cert: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
      key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
    serviceSigningCa:
      cert: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
      key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
  configStorageAccount: Config.ConfigStorageAccount
  customerAdminPasswd: Config.CustomerAdminPasswd
  customerReaderPasswd: Config.CustomerReaderPasswd
  endUserPasswd: Config.EndUserPasswd
  htPasswd: Q29uZmlnLkh0UGFzc3dk
  imageOffer: Config.ImageOffer
  imagePublisher: Config.ImagePublisher
  imageSku: Config.ImageSKU
  imageVersion: Config.ImageVersion
  images:
    addonsResizer: Config.Images.AddonsResizer
    ansibleServiceBroker: Config.Images.AnsibleServiceBroker
    controlPlane: Config.Images.ControlPlane
    etcdOperator: Config.Images.EtcdOperator
    format: Config.Images.Format
    kubeStateMetrics: Config.Images.KubeStateMetrics
    logBridge: Config.Images.LogBridge
    masterEtcd: Config.Images.MasterEtcd
    node: Config.Images.Node
    oAuthProxy: Config.Images.OAuthProxy
    prometheus: Config.Images.Prometheus
    prometheusAlertBuffer: Config.Images.PrometheusAlertBuffer
    prometheusAlertManager: Config.Images.PrometheusAlertManager
    prometheusNodeExporter: Config.Images.PrometheusNodeExporter
    registry: Config.Images.Registry
    registryConsole: Config.Images.RegistryConsole
    router: Config.Images.Router
    serviceCatalog: Config.Images.ServiceCatalog
    sync: Config.Images.Sync
    templateServiceBroker: Config.Images.TemplateServiceBroker
    webConsole: Config.Images.WebConsole
  loggingLocation: Config.LoggingLocation
  loggingWorkspace: Config.LoggingWorkspace
  masterKubeconfig: eyJwcmVmZXJlbmNlcyI6e30sImNsdXN0ZXJzIjpudWxsLCJ1c2VycyI6bnVsbCwiY29udGV4dHMiOm51bGwsImN1cnJlbnQtY29udGV4dCI6IiJ9
  nodeBootstrapKubeconfig: eyJwcmVmZXJlbmNlcyI6e30sImNsdXN0ZXJzIjpudWxsLCJ1c2VycyI6bnVsbCwiY29udGV4dHMiOm51bGwsImN1cnJlbnQtY29udGV4dCI6IiJ9
  prometheusProxySessionSecret: Q29uZmlnLlByb21ldGhldXNQcm94eVNlc3Npb25TZWNyZXQ=
  registryConsoleOAuthSecret: Config.RegistryConsoleOAuthSecret
  registryHttpSecret: Q29uZmlnLlJlZ2lzdHJ5SFRUUFNlY3JldA==
  registryStorageAccount: Config.RegistryStorageAccount
  routerStatsPassword: Config.RouterStatsPassword
  runningUnderTest: true
  serviceAccountKey: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
  serviceCatalogClusterId: 01010101-0101-0101-0101-010101010101
  sessionSecretAuth: Q29uZmlnLlNlc3Npb25TZWNyZXRBdXRo
  sessionSecretEnc: Q29uZmlnLlNlc3Npb25TZWNyZXRFbmM=
  sshKey: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNQ1VDQVFBQ0F3Q2lYd0lEQVFBQkFnSkQ4UUlDQU1VQ0FnRFRBZ0lBa1FJQ0FLMENBZ0MzCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
id: ID
location: Location
name: Name
plan:
  name: Plan.Name
  product: Plan.Product
  promotionCode: Plan.PromotionCode
  publisher: Plan.Publisher
properties:
  agentPoolProfiles:
  - count: 1
    name: Properties.AgentPoolProfiles[0].Name
    osType: Properties.AgentPoolProfiles[0].OSType
    role: Properties.AgentPoolProfiles[0].Role
    subnetCidr: Properties.AgentPoolProfiles[0].SubnetCIDR
    vmSize: Properties.AgentPoolProfiles[0].VMSize
  authProfile:
    identityProviders:
    - name: Properties.AuthProfile.IdentityProviders[0].Name
      provider:
        clientId: Properties.AuthProfile.IdentityProviders[0].Provider.ClientID
        kind: AADIdentityProvider
        secret: Properties.AuthProfile.IdentityProviders[0].Provider.Secret
        tenantId: Properties.AuthProfile.IdentityProviders[0].Provider.TenantID
  azProfile:
    resourceGroup: Properties.AzProfile.ResourceGroup
    subscriptionId: Properties.AzProfile.SubscriptionID
    tenantId: Properties.AzProfile.TenantID
  fqdn: Properties.FQDN
  networkProfile:
    peerVnetId: Properties.NetworkProfile.PeerVnetID
    vnetCidr: Properties.NetworkProfile.VnetCIDR
  openShiftVersion: Properties.OpenShiftVersion
  provisioningState: Properties.ProvisioningState
  publicHostname: Properties.PublicHostname
  routerProfiles:
  - fqdn: Properties.RouterProfiles[0].FQDN
    name: Properties.RouterProfiles[0].Name
    publicSubdomain: Properties.RouterProfiles[0].PublicSubdomain
  servicePrincipalProfile:
    clientId: Properties.ServicePrincipalProfile.ClientID
    secret: Properties.ServicePrincipalProfile.Secret
tags:
  Tags.key: Tags.val
type: Type
`)

func TestMarshal(t *testing.T) {
	b, err := json.MarshalIndent(unmarshalled, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(b, marshalled) {
		t.Errorf("json.MarshalIndent returned unexpected result\n%s\n", string(b))
	}
}

func TestUnmarshal(t *testing.T) {
	var oc *OpenShiftManagedCluster
	err := json.Unmarshal(marshalled, &oc)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(oc, unmarshalled) {
		t.Errorf("json.Unmarshal returned unexpected result\n%#v\n", oc)
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

func TestMarshallYaml(t *testing.T) {
	populatedOc := OpenShiftManagedCluster{}
	populate.Walk(&populatedOc)

	b, err := yaml.Marshal(populatedOc)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(b, marshalledYaml) {
		t.Errorf("yaml.Marshall returned unexpected result\n%s\n%s\n", string(b), marshalledYaml)
	}
}

func TestUnmarshallYaml(t *testing.T) {
	t.Skip("skipping this test until we are able to unmarshall x509.Certificate")
	populatedOc := OpenShiftManagedCluster{}
	populate.Walk(&populatedOc)

	var unmarshalledOc OpenShiftManagedCluster
	err := yaml.Unmarshal(marshalledYaml, &unmarshalledOc)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(populatedOc, unmarshalledOc) {
		t.Errorf("yaml.Unmarshal returned unexpected result\n%#v\n", unmarshalledOc)
	}
}
