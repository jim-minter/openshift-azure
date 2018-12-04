package azure

import (
	"os"

	"github.com/openshift/openshift-azure/pkg/util/azureclient"
)

// Client is the main controller for azure client objects
type Client struct {
	Accounts                         azureclient.AccountsClient
	Applications                     azureclient.ApplicationsClient
	VirtualMachineScaleSets          azureclient.VirtualMachineScaleSetsClient
	VirtualMachineScaleSetExtensions azureclient.VirtualMachineScaleSetExtensionsClient
	VirtualMachineScaleSetVMs        azureclient.VirtualMachineScaleSetVMsClient
	Resources                        azureclient.ResourcesClient
}

// NewClientFromEnvironment creates a new azure client from environment variables
func NewClientFromEnvironment() (*Client, error) {
	authorizer, err := azureclient.NewAuthorizerFromEnvironment()
	if err != nil {
		return nil, err
	}

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")

	return &Client{
		Accounts:                         azureclient.NewAccountsClient(subscriptionID, authorizer, nil),
		Applications:                     azureclient.NewApplicationsClient(subscriptionID, authorizer, nil),
		VirtualMachineScaleSets:          azureclient.NewVirtualMachineScaleSetsClient(subscriptionID, authorizer, nil),
		VirtualMachineScaleSetExtensions: azureclient.NewVirtualMachineScaleSetExtensionsClient(subscriptionID, authorizer, nil),
		VirtualMachineScaleSetVMs:        azureclient.NewVirtualMachineScaleSetVMsClient(subscriptionID, authorizer, nil),
		Resources:                        azureclient.NewResourcesClient(subscriptionID, authorizer, nil),
	}, nil
}
