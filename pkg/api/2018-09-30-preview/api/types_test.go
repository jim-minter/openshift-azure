package api_test

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ghodss/yaml"

	. "github.com/openshift/openshift-azure/pkg/api/2018-09-30-preview/api"
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
		FQDN:              "properties.fqdn",
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
		NetworkProfile: &NetworkProfile{
			VnetCIDR:   "properties.networkProfile.vnetCidr",
			PeerVnetID: "properties.networkProfile.peerVnetId",
		},
		MasterPoolProfile: &MasterPoolProfile{
			Count:      1,
			VMSize:     "properties.agentPoolProfiles.0.vmSize",
			SubnetCIDR: "properties.agentPoolProfiles.0.subnetCidr",
		},
		AgentPoolProfiles: []AgentPoolProfile{
			{
				Name:       "properties.agentPoolProfiles.0.name",
				Count:      1,
				VMSize:     "properties.agentPoolProfiles.0.vmSize",
				SubnetCIDR: "properties.agentPoolProfiles.0.subnetCidr",
				OSType:     "properties.agentPoolProfiles.0.osType",
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
		"masterPoolProfile": {
			"count": 1,
			"vmSize": "properties.agentPoolProfiles.0.vmSize",
			"subnetCidr": "properties.agentPoolProfiles.0.subnetCidr"
		},
		"agentPoolProfiles": [
			{
				"name": "properties.agentPoolProfiles.0.name",
				"count": 1,
				"vmSize": "properties.agentPoolProfiles.0.vmSize",
				"subnetCidr": "properties.agentPoolProfiles.0.subnetCidr",
				"osType": "properties.agentPoolProfiles.0.osType",
				"role": "properties.agentPoolProfiles.0.role"
			},
			{
				"name": "properties.agentPoolProfiles.0.name",
				"count": 2,
				"vmSize": "properties.agentPoolProfiles.0.vmSize",
				"subnetCidr": "properties.agentPoolProfiles.0.subnetCidr",
				"osType": "properties.agentPoolProfiles.0.osType",
				"role": "properties.agentPoolProfiles.0.role"
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

var marshalledYaml = []byte(`id: ID
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
  fqdn: Properties.FQDN
  masterPoolProfile:
    count: 1
    subnetCidr: Properties.MasterPoolProfile.SubnetCIDR
    vmSize: Properties.MasterPoolProfile.VMSize
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

func TestStructTypes(t *testing.T) {
	populateFields := func(t reflect.Type) map[string]reflect.StructField {
		m := map[string]reflect.StructField{}

		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			f.Index = nil
			f.Offset = 0

			m[f.Name] = f
		}

		return m
	}

	appFields := populateFields(reflect.TypeOf(AgentPoolProfile{}))
	mppFields := populateFields(reflect.TypeOf(MasterPoolProfile{}))

	// every field (except Name, OSType, Role) in AgentPoolProfile should be
	// identical in MasterPoolProfile
	for name := range appFields {
		switch name {
		case "Name", "OSType", "Role":
			continue
		}

		if !reflect.DeepEqual(appFields[name], mppFields[name]) {
			t.Errorf("mismatch in field %s:\n%#v\n%#v", name, appFields[name], mppFields[name])
		}
	}

	// every field in MasterPoolProfile should be identical in
	// AgentPoolProfile
	for name := range mppFields {
		if !reflect.DeepEqual(appFields[name], mppFields[name]) {
			t.Errorf("mismatch in field %s:\n%#v\n%#v", name, appFields[name], mppFields[name])
		}
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
