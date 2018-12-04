package azureclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/Azure/go-autorest/autorest"
)

// DeploymentsClient is a minimal interface for azure DeploymentsClient
type DeploymentsClient interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, deploymentName string, parameters resources.Deployment) (result resources.DeploymentsCreateOrUpdateFuture, err error)
	Client
	DeploymentClient() resources.DeploymentsClient
}

type deploymentsClient struct {
	resources.DeploymentsClient
}

var _ DeploymentsClient = &deploymentsClient{}

// NewDeploymentsClient creates a new DeploymentsClient
func NewDeploymentsClient(subscriptionID string, authorizer autorest.Authorizer, languages []string) DeploymentsClient {
	client := resources.NewDeploymentsClient(subscriptionID)
	client.Authorizer = authorizer
	client.RequestInspector = addAcceptLanguages(languages)

	return &deploymentsClient{
		DeploymentsClient: client,
	}
}

func (c *deploymentsClient) Client() autorest.Client {
	return c.DeploymentsClient.Client
}

func (c *deploymentsClient) DeploymentClient() resources.DeploymentsClient {
	return c.DeploymentsClient
}

// ResourcesClient is a minimal interface for azure Resources Client
type ResourcesClient interface {
	ResourceClient() resources.Client
}

type resourcesClient struct {
	ResourcesClient resources.Client
}

var _ ResourcesClient = &resourcesClient{}

// NewResourcesClient creates a new ResourcesClient
func NewResourcesClient(subscriptionID string, authorizer autorest.Authorizer, languages []string) ResourcesClient {
	client := resources.NewClient(subscriptionID)
	client.Authorizer = authorizer
	client.RequestInspector = addAcceptLanguages(languages)

	return &resourcesClient{
		ResourcesClient: client,
	}
}

func (c *resourcesClient) Client() autorest.Client {
	return c.ResourcesClient.Client
}

func (c *resourcesClient) ResourceClient() resources.Client {
	return c.ResourcesClient
}
