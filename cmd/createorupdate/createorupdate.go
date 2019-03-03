package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/ghodss/yaml"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/wait"

	v20180930preview "github.com/openshift/openshift-azure/pkg/api/2018-09-30-preview/api"
	admin "github.com/openshift/openshift-azure/pkg/api/admin/api"
	fakerp "github.com/openshift/openshift-azure/pkg/fakerp/client"
	"github.com/openshift/openshift-azure/pkg/fakerp/shared"
	"github.com/openshift/openshift-azure/pkg/util/aadapp"
	"github.com/openshift/openshift-azure/pkg/util/azureclient"
	v20180930previewclient "github.com/openshift/openshift-azure/pkg/util/azureclient/openshiftmanagedcluster/2018-09-30-preview"
	adminclient "github.com/openshift/openshift-azure/pkg/util/azureclient/openshiftmanagedcluster/admin"
)

var (
	method  = flag.String("request", http.MethodPut, "Specify request to send to the OpenShift resource provider. Supported methods are PUT and DELETE.")
	useProd = flag.Bool("use-prod", false, "If true, send the request to the production OpenShift resource provider.")

	adminManifest   = flag.String("admin-manifest", "", "If set, use the admin API to send this request.")
	restoreFromBlob = flag.String("restore-from-blob", "", "If set, request a restore of the cluster from the provided blob name.")
)

func validate() error {
	m := strings.ToUpper(*method)
	switch m {
	case http.MethodPut, http.MethodDelete:
	default:
		return fmt.Errorf("invalid request: %s, Supported methods are PUT and DELETE", strings.ToUpper(*method))
	}
	if *adminManifest != "" && *useProd {
		return errors.New("sending requests to the Admin API is not supported yet in the production RP")
	}
	if *restoreFromBlob != "" && *useProd {
		return errors.New("restoring clusters is not supported yet in the production RP")
	}
	if *restoreFromBlob != "" && m == http.MethodDelete {
		return errors.New("cannot restore a cluster while requesting a DELETE?")
	}
	return nil
}

func delete(ctx context.Context, log *logrus.Entry, rpc v20180930previewclient.OpenShiftManagedClustersClient, resourceGroup string, noWait bool) error {
	log.Info("deleting cluster")
	future, err := rpc.Delete(ctx, resourceGroup, resourceGroup)
	if err != nil {
		return err
	}
	if noWait {
		log.Info("will not wait for cluster deletion")
	} else {
		log.Info("waiting for cluster deletion")
		if err := future.WaitForCompletionRef(ctx, rpc.Client); err != nil {
			return err
		}
		log.Info("deleted cluster")
	}
	return nil
}

func createOrUpdatev20180930preview(ctx context.Context, log *logrus.Entry, rpc v20180930previewclient.OpenShiftManagedClustersClient, resourceGroup string, oc *v20180930preview.OpenShiftManagedCluster, manifestFile string) error {
	log.Info("creating/updating cluster")
	resp, err := rpc.CreateOrUpdateAndWait(ctx, resourceGroup, resourceGroup, *oc)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response: %s", resp.Status)
	}
	log.Info("created/updated cluster")
	return fakerp.WriteClusterConfigToManifest(&resp, manifestFile)
}

func createOrUpdateAdmin(ctx context.Context, log *logrus.Entry, rpc adminclient.OpenShiftManagedClustersClient, resourceGroup string, oc *admin.OpenShiftManagedCluster, manifestFile string) error {
	log.Info("creating/updating cluster")
	if oc.Properties != nil {
		oc.Properties.ProvisioningState = nil
	}
	resp, err := rpc.CreateOrUpdateAndWait(ctx, resourceGroup, resourceGroup, *oc)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response: %s", resp.Status)
	}
	data, err := yaml.Marshal(resp)
	if err != nil {
		return err
	}
	log.Info("created/updated cluster")
	return ioutil.WriteFile(manifestFile, data, 0600)
}

func execute(
	ctx context.Context,
	log *logrus.Entry,
	ac adminclient.OpenShiftManagedClustersClient,
	rpc v20180930previewclient.OpenShiftManagedClustersClient,
	conf *fakerp.Config,
	adminManifest string,
) error {
	dataDir, err := shared.FindDirectory(shared.DataDirectory)
	if err != nil {
		return fmt.Errorf("failed to find %s: %v", shared.DataDirectory, err)
	}

	if adminManifest != "" {
		oc, err := fakerp.GenerateManifestAdmin(adminManifest)
		if err != nil {
			return fmt.Errorf("failed reading admin manifest: %v", err)
		}
		defaultAdminManifest := filepath.Join(dataDir, "manifest-admin.yaml")
		return createOrUpdateAdmin(ctx, log, ac, conf.ResourceGroup, oc, defaultAdminManifest)
	}

	defaultManifestFile := filepath.Join(dataDir, "manifest.yaml")
	// TODO: Configuring this is probably not needed
	manifest := conf.Manifest
	// If no MANIFEST has been provided and this is a cluster
	// creation, default to the test manifest.
	if !shared.IsUpdate() && manifest == "" {
		manifest = "test/manifests/normal/create.yaml"
	}
	// If this is a cluster upgrade, reuse the existing manifest.
	if manifest == "" {
		manifest = defaultManifestFile
	}

	oc, err := fakerp.GenerateManifest(manifest)
	if err != nil {
		return fmt.Errorf("failed reading manifest: %v", err)
	}

	return createOrUpdatev20180930preview(ctx, log, rpc, conf.ResourceGroup, oc, defaultManifestFile)
}

func updateAadApplication(ctx context.Context, log *logrus.Entry, conf *fakerp.Config) error {
	if len(conf.AADClientID) > 0 && conf.AADClientID != conf.ClientID {
		log.Info("updating the aad application")
		authorizer, err := azureclient.NewAuthorizer(conf.ClientID, conf.ClientSecret, conf.TenantID, azure.PublicCloud.GraphEndpoint)
		if err != nil {
			return fmt.Errorf("cannot get authorizer: %v", err)
		}
		aadClient := azureclient.NewRBACApplicationsClient(ctx, conf.TenantID, authorizer)
		objID, err := aadapp.GetApplicationObjectIDFromAppID(ctx, aadClient, conf.AADClientID)
		if err != nil {
			return err
		}

		callbackURL := fmt.Sprintf("https://%s.%s.cloudapp.azure.com/oauth2callback/Azure%%20AD", conf.ResourceGroup, conf.Region)
		conf.AADClientSecret, err = aadapp.UpdateAADApp(ctx, aadClient, objID, callbackURL)
		if err != nil {
			return fmt.Errorf("cannot update aad app secret: %v", err)
		}
	} else {
		log.Info("using dummy aad")
		conf.AADClientID = conf.ClientID
		conf.AADClientSecret = conf.ClientSecret
	}
	// set env variable so enrich() still works
	err := os.Setenv("AZURE_AAD_CLIENT_ID", conf.AADClientID)
	if err != nil {
		return fmt.Errorf("failed setting AZURE_AAD_CLIENT_ID: %v", err)
	}
	err = os.Setenv("AZURE_AAD_CLIENT_SECRET", conf.AADClientSecret)
	if err != nil {
		return fmt.Errorf("failed setting AZURE_AAD_CLIENT_SECRET: %v", err)
	}
	return nil
}

func isConnectionRefused(err error) bool {
	if autoRestErr, ok := err.(autorest.DetailedError); ok {
		if urlErr, ok := autoRestErr.Original.(*url.Error); ok {
			if netErr, ok := urlErr.Err.(*net.OpError); ok {
				if sysErr, ok := netErr.Err.(*os.SyscallError); ok {
					if sysErr.Err == syscall.ECONNREFUSED {
						return true
					}
				}
			}
		}
	}
	return false
}

func main() {
	flag.Parse()
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	log := logrus.NewEntry(logrus.StandardLogger())

	if err := validate(); err != nil {
		log.Fatal(err)
	}

	isDelete := strings.ToUpper(*method) == http.MethodDelete
	conf, err := fakerp.NewConfig(log, !isDelete)
	if err != nil {
		log.Fatal(err)
	}

	if !isDelete {
		log.Infof("creating resource group %s", conf.ResourceGroup)
		if isCreate, err := fakerp.CreateResourceGroup(conf); err != nil {
			log.Fatal(err)
		} else if !isCreate {
			log.Infof("reusing existing resource group %s", conf.ResourceGroup)
		}
	}

	// simulate the RP
	rpURL := v20180930previewclient.DefaultBaseURI
	if !*useProd {
		rpURL = fmt.Sprintf("http://%s", shared.LocalHttpAddr)
	}

	// setup the osa clients
	adminClient := adminclient.NewOpenShiftManagedClustersClientWithBaseURI(rpURL+shared.AdminContext, conf.SubscriptionID)
	v20180930previewClient := v20180930previewclient.NewOpenShiftManagedClustersClientWithBaseURI(rpURL, conf.SubscriptionID)
	authorizer, err := azureclient.NewAuthorizer(conf.ClientID, conf.ClientSecret, conf.TenantID, "")
	if err != nil {
		log.Fatal(err)
	}
	adminClient.Authorizer = authorizer
	v20180930previewClient.Authorizer = authorizer

	ctx := context.Background()
	if isDelete {
		if err := wait.PollImmediate(time.Second, 1*time.Hour, func() (bool, error) {
			if err := delete(ctx, log, v20180930previewClient, conf.ResourceGroup, conf.NoWait); err != nil {
				if isConnectionRefused(err) {
					return false, nil
				}
				return false, err
			}
			return true, nil
		}); err != nil {
			log.Fatal(err)
		}
		return
	}

	if *restoreFromBlob != "" {
		err = wait.PollImmediate(time.Second, 1*time.Hour, func() (bool, error) {
			resp, err := adminClient.RestoreAndWait(ctx, conf.ResourceGroup, conf.ResourceGroup, *restoreFromBlob)
			if isConnectionRefused(err) {
				return false, nil
			}
			if err != nil {
				return false, err
			}
			if resp.StatusCode != http.StatusOK {
				return false, fmt.Errorf("expected 200 OK, got %v", resp.Status)
			}
			return true, nil
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	if !shared.IsUpdate() {
		if err := updateAadApplication(ctx, log, conf); err != nil {
			log.Fatal(err)
		}
	}

	err = wait.PollImmediate(time.Second, 1*time.Hour, func() (bool, error) {
		if err := execute(ctx, log, adminClient, v20180930previewClient, conf, *adminManifest); err != nil {
			if isConnectionRefused(err) {
				return false, nil
			}
			return false, err
		}
		return true, nil
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nCluster available at https://%s.%s.cloudapp.azure.com/\n", conf.ResourceGroup, conf.Region)
}
