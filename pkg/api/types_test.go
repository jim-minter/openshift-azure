package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unicode"
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
				Count:        1,
				VMSize:       "properties.agentPoolProfiles.0.vmSize",
				VnetSubnetID: "properties.agentPoolProfiles.0.vnetSubnetID",
				Role:         "properties.agentPoolProfiles.0.role",
			},
			{
				Name:         "properties.agentPoolProfiles.0.name",
				Count:        2,
				VMSize:       "properties.agentPoolProfiles.0.vmSize",
				OSType:       "properties.agentPoolProfiles.0.osType",
				VnetSubnetID: "properties.agentPoolProfiles.0.vnetSubnetID",
				Role:         "properties.agentPoolProfiles.0.role",
			},
			{
				Name:         "properties.agentPoolProfiles.0.name",
				Count:        1,
				VMSize:       "properties.agentPoolProfiles.0.vmSize",
				OSType:       "properties.agentPoolProfiles.0.osType",
				VnetSubnetID: "properties.agentPoolProfiles.0.vnetSubnetID",
				Role:         "master",
			},
		},
		ServicePrincipalProfile: &ServicePrincipalProfile{
			ClientID: "properties.servicePrincipalProfile.clientId",
			Secret:   "properties.servicePrincipalProfile.secret",
		},
	},
}

var marshalled = []byte(`{
	"id": "id",
	"location": "location",
	"name": "name",
	"plan": {
		"name": "plan.name",
		"product": "plan.product",
		"promotionCode": "plan.promotionCode",
		"publisher": "plan.publisher"
	},
	"tags": {
		"tags.k1": "v1",
		"tags.k2": "v2"
	},
	"type": "type",
	"properties": {
		"provisioningState": "properties.provisioningState",
		"openShiftVersion": "properties.openShiftVersion",
		"publicHostname": "properties.publicHostname",
		"fqdn": "properties.fqdn",
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
				"vnetSubnetID": "properties.agentPoolProfiles.0.vnetSubnetID",
				"role": "properties.agentPoolProfiles.0.role"
			},
			{
				"name": "properties.agentPoolProfiles.0.name",
				"count": 2,
				"vmSize": "properties.agentPoolProfiles.0.vmSize",
				"vnetSubnetID": "properties.agentPoolProfiles.0.vnetSubnetID",
				"osType": "properties.agentPoolProfiles.0.osType",
				"role": "properties.agentPoolProfiles.0.role"
			},
			{
				"name": "properties.agentPoolProfiles.0.name",
				"count": 1,
				"vmSize": "properties.agentPoolProfiles.0.vmSize",
				"vnetSubnetID": "properties.agentPoolProfiles.0.vnetSubnetID",
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
	}
}`)

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
	for _, err := range testJSONTags(reflect.TypeOf(OpenShiftManagedCluster{})) {
		t.Error(err)
	}
}

// toTag converts a field name like ImageSKU to imageSku
func toTag(s string) string {
	if s == "ServiceCatalogAPIClient" {
		// should be serviceCatalogApiClient, but isn't.  TODO: maybe fix this
		// one day.
		return "serviceCatalogAPIClient"
	}

	for _, acronym := range []string{"API", "FQDN", "HTTP", "ID", "SKU", "SSH", "VM"} {
		lower := string(acronym[0]) + strings.Map(unicode.ToLower, acronym[1:])
		s = strings.Replace(s, acronym, lower, -1)
	}

	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

func testJSONTags(t reflect.Type) (errs []error) {
	switch t.Kind() {
	case reflect.Ptr:
		errs = append(errs, testJSONTags(t.Elem())...)

	case reflect.Struct:
		if !strings.HasPrefix(t.PkgPath(), "github.com/openshift/openshift-azure/") ||
			strings.HasPrefix(t.PkgPath(), "github.com/openshift/openshift-azure/vendor/") {
			return
		}

		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			tag := f.Tag.Get("json")
			parts := strings.Split(tag, ",")
			if len(parts) != 2 || parts[1] != "omitempty" {
				errs = append(errs, fmt.Errorf("invalid tag %s", tag))
			}
			if toTag(f.Name) != parts[0] {
				errs = append(errs, fmt.Errorf("%s: tag name mismatch: wanted %s, got %s", f.Name, toTag(f.Name), parts[0]))
			}

			errs = append(errs, testJSONTags(f.Type)...)
		}
	}

	return
}
