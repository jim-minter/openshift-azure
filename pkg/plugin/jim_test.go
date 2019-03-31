package plugin

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2018-02-01/storage"
	azstorage "github.com/Azure/azure-sdk-for-go/storage"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"

	"github.com/openshift/openshift-azure/pkg/api"
	"github.com/openshift/openshift-azure/pkg/cluster"
	"github.com/openshift/openshift-azure/pkg/util/mocks/mock_azureclient"
	"github.com/openshift/openshift-azure/pkg/util/mocks/mock_azureclient/mock_storage"
	"github.com/openshift/openshift-azure/pkg/util/mocks/mock_kubeclient"
	"github.com/openshift/openshift-azure/pkg/util/mocks/mock_scaler"
	"github.com/openshift/openshift-azure/pkg/util/mocks/mock_wait"
	"github.com/openshift/openshift-azure/test/util/populate"
	"github.com/openshift/openshift-azure/test/util/tls"
)

func create(t *testing.T, cs *api.OpenShiftManagedCluster) (blob []byte) {
	ctx := context.Background()

	gmc := gomock.NewController(t)
	defer gmc.Finish()

	log := logrus.NewEntry(logrus.StandardLogger())
	httpClient := mock_wait.NewMockSimpleHTTPClient(gmc)
	kubeclient := mock_kubeclient.NewMockKubeclient(gmc)
	scalerFactory := mock_scaler.NewMockFactory(gmc)
	accountsClient := mock_azureclient.NewMockAccountsClient(gmc)
	vmc := mock_azureclient.NewMockVirtualMachineScaleSetVMsClient(gmc)
	ssc := mock_azureclient.NewMockVirtualMachineScaleSetsClient(gmc)
	kvc := mock_azureclient.NewMockKeyVaultClient(gmc)
	storageClient := mock_storage.NewMockClient(gmc)
	blobService := mock_storage.NewMockBlobStorageClient(gmc)
	configContainer := mock_storage.NewMockContainer(gmc)
	etcdContainer := mock_storage.NewMockContainer(gmc)
	updateContainer := mock_storage.NewMockContainer(gmc)
	masterStartupBlob := mock_storage.NewMockBlob(gmc)
	workerStartupBlob := mock_storage.NewMockBlob(gmc)
	updateBlob := mock_storage.NewMockBlob(gmc)
	syncBlob := mock_storage.NewMockBlob(gmc)

	c := blobService.EXPECT().GetContainerReference("update").Return(updateContainer)
	// CreateOrUpdateConfigStorageAccount
	c = accountsClient.EXPECT().Create(ctx, "", "", gomock.Any()).Return(nil).After(c)
	c = storageClient.EXPECT().GetBlobService().Return(blobService).After(c)
	c = blobService.EXPECT().GetContainerReference("config").Return(configContainer).After(c)
	c = configContainer.EXPECT().CreateIfNotExists(nil).Return(true, nil).After(c)
	c = blobService.EXPECT().GetContainerReference("etcd").Return(etcdContainer).After(c)
	c = etcdContainer.EXPECT().CreateIfNotExists(nil).Return(true, nil).After(c)
	c = blobService.EXPECT().GetContainerReference("update").Return(updateContainer).After(c)
	c = updateContainer.EXPECT().CreateIfNotExists(nil).Return(true, nil).After(c)
	// WriteStartupBlobs
	c = storageClient.EXPECT().GetBlobService().Return(blobService).After(c)
	c = blobService.EXPECT().GetContainerReference("config").Return(configContainer).After(c)
	c = configContainer.EXPECT().GetBlobReference("master-startup").Return(masterStartupBlob).After(c)
	c = masterStartupBlob.EXPECT().GetSASURI(gomock.Any()).Return("", nil).After(c)
	c = configContainer.EXPECT().GetBlobReference("worker-startup").Return(workerStartupBlob).After(c)
	c = workerStartupBlob.EXPECT().GetSASURI(gomock.Any()).Return("", nil).After(c)
	c = storageClient.EXPECT().GetBlobService().Return(blobService).After(c)
	c = blobService.EXPECT().GetContainerReference("config").Return(configContainer).After(c)
	c = configContainer.EXPECT().GetBlobReference("master-startup").Return(masterStartupBlob).After(c)
	c = masterStartupBlob.EXPECT().CreateBlockBlobFromReader(gomock.Any(), nil).After(c)
	c = storageClient.EXPECT().GetBlobService().Return(blobService).After(c)
	c = blobService.EXPECT().GetContainerReference("config").Return(configContainer).After(c)
	c = configContainer.EXPECT().GetBlobReference("worker-startup").Return(workerStartupBlob).After(c)
	c = workerStartupBlob.EXPECT().CreateBlockBlobFromReader(gomock.Any(), nil).After(c)
	// EnrichCertificatesFromVault
	c = kvc.EXPECT().GetSecret(ctx, "", ".", "").Return(keyvault.SecretBundle{Value: &tls.DummyBundle}, nil).After(c)
	c = kvc.EXPECT().GetSecret(ctx, "", ".", "").Return(keyvault.SecretBundle{Value: &tls.DummyBundle}, nil).After(c)
	// EnrichStorageAccountKeys
	c = accountsClient.EXPECT().ListKeys(ctx, "", "").Return(storage.AccountListKeysResult{Keys: &[]storage.AccountKey{{Value: to.StringPtr("")}}}, nil).After(c)
	c = accountsClient.EXPECT().ListKeys(ctx, "", "").Return(storage.AccountListKeysResult{Keys: &[]storage.AccountKey{{Value: to.StringPtr("")}}}, nil).After(c)
	// InitializeUpdateBlob
	c = updateContainer.EXPECT().GetBlobReference("update").Return(updateBlob).After(c)
	c = updateBlob.EXPECT().CreateBlockBlobFromReader(gomock.Any(), nil).DoAndReturn(func(r io.Reader, options *azstorage.PutBlobOptions) (err error) {
		blob, err = ioutil.ReadAll(r)
		return
	}).After(c)
	// WaitForHealthzStatusOk
	c = httpClient.EXPECT().Do(gomock.Any()).Return(&http.Response{StatusCode: http.StatusOK}, nil).After(c)
	// CreateOrUpdateSyncPod
	c = storageClient.EXPECT().GetBlobService().Return(blobService).After(c)
	c = blobService.EXPECT().GetContainerReference("config").Return(configContainer).After(c)
	c = configContainer.EXPECT().GetBlobReference("sync").Return(syncBlob).After(c)
	c = syncBlob.EXPECT().CreateBlockBlobFromReader(gomock.Any(), nil).After(c)
	c = kubeclient.EXPECT().EnsureSyncPod(ctx, "", gomock.Any()).Return(nil).After(c)
	// WaitForNodesInAgentPoolProfile
	c = vmc.EXPECT().List(ctx, "", "ss-master", "", "", "").Return([]compute.VirtualMachineScaleSetVM{
		{
			VirtualMachineScaleSetVMProperties: &compute.VirtualMachineScaleSetVMProperties{
				OsProfile: &compute.OSProfile{
					ComputerName: to.StringPtr("master-000000"),
				},
			},
		},
	}, nil).After(c)
	c = kubeclient.EXPECT().WaitForReadyMaster(ctx, "master-000000").Return(nil).After(c)
	// WaitForNodesInAgentPoolProfile
	c = vmc.EXPECT().List(ctx, "", "ss-infra-0", "", "", "").Return([]compute.VirtualMachineScaleSetVM{
		{
			VirtualMachineScaleSetVMProperties: &compute.VirtualMachineScaleSetVMProperties{
				OsProfile: &compute.OSProfile{
					ComputerName: to.StringPtr("ss-infra-0-000000"),
				},
			},
		},
	}, nil).After(c)
	c = kubeclient.EXPECT().WaitForReadyWorker(ctx, "ss-infra-0-000000").Return(nil).After(c)
	// WaitForNodesInAgentPoolProfile
	c = vmc.EXPECT().List(ctx, "", "ss-compute-0", "", "", "").Return([]compute.VirtualMachineScaleSetVM{
		{
			VirtualMachineScaleSetVMProperties: &compute.VirtualMachineScaleSetVMProperties{
				OsProfile: &compute.OSProfile{
					ComputerName: to.StringPtr("ss-compute-0-000000"),
				},
			},
		},
	}, nil).After(c)
	c = kubeclient.EXPECT().WaitForReadyWorker(ctx, "ss-compute-0-000000").Return(nil).After(c)
	// WaitForReadySyncPod
	c = kubeclient.EXPECT().WaitForReadySyncPod(ctx).Return(nil).After(c)
	// HealthCheck
	c = httpClient.EXPECT().Do(gomock.Any()).Return(&http.Response{StatusCode: http.StatusOK}, nil).After(c)

	p := &plugin{
		log: log,
		upgraderFactory: func(ctx context.Context, log *logrus.Entry, cs *api.OpenShiftManagedCluster, initializeStorageClients, disableKeepAlives bool, testConfig api.TestConfig) (cluster.Upgrader, error) {
			return cluster.NewFakeUpgrader(ctx, log, cs, accountsClient, vmc, ssc, kvc, storageClient, blobService, httpClient, kubeclient, scalerFactory, api.TestConfig{})
		},
		now: func() time.Time { return time.Unix(0, 0) },
	}
	err := p.CreateOrUpdate(ctx, cs, false, func(context.Context, map[string]interface{}) error { return nil })
	if err != nil {
		t.Error(err)
	}

	return
}

func update(t *testing.T, cs *api.OpenShiftManagedCluster, blob []byte) {
	ctx := context.Background()

	gmc := gomock.NewController(t)
	defer gmc.Finish()

	log := logrus.NewEntry(logrus.StandardLogger())
	httpClient := mock_wait.NewMockSimpleHTTPClient(gmc)
	kubeclient := mock_kubeclient.NewMockKubeclient(gmc)
	scalerFactory := mock_scaler.NewMockFactory(gmc)
	targetScaler := mock_scaler.NewMockScaler(gmc)
	sourceScaler := mock_scaler.NewMockScaler(gmc)
	accountsClient := mock_azureclient.NewMockAccountsClient(gmc)
	vmc := mock_azureclient.NewMockVirtualMachineScaleSetVMsClient(gmc)
	ssc := mock_azureclient.NewMockVirtualMachineScaleSetsClient(gmc)
	kvc := mock_azureclient.NewMockKeyVaultClient(gmc)
	storageClient := mock_storage.NewMockClient(gmc)
	blobService := mock_storage.NewMockBlobStorageClient(gmc)
	configContainer := mock_storage.NewMockContainer(gmc)
	etcdContainer := mock_storage.NewMockContainer(gmc)
	updateContainer := mock_storage.NewMockContainer(gmc)
	masterStartupBlob := mock_storage.NewMockBlob(gmc)
	workerStartupBlob := mock_storage.NewMockBlob(gmc)
	updateBlob := mock_storage.NewMockBlob(gmc)
	syncBlob := mock_storage.NewMockBlob(gmc)

	c := blobService.EXPECT().GetContainerReference("update").Return(updateContainer)
	// CreateOrUpdateConfigStorageAccount
	c = accountsClient.EXPECT().Create(ctx, "", "", gomock.Any()).Return(nil).After(c)
	c = storageClient.EXPECT().GetBlobService().Return(blobService).After(c)
	c = blobService.EXPECT().GetContainerReference("config").Return(configContainer).After(c)
	c = configContainer.EXPECT().CreateIfNotExists(nil).Return(true, nil).After(c)
	c = blobService.EXPECT().GetContainerReference("etcd").Return(etcdContainer).After(c)
	c = etcdContainer.EXPECT().CreateIfNotExists(nil).Return(true, nil).After(c)
	c = blobService.EXPECT().GetContainerReference("update").Return(updateContainer).After(c)
	c = updateContainer.EXPECT().CreateIfNotExists(nil).Return(true, nil).After(c)
	// WriteStartupBlobs
	c = storageClient.EXPECT().GetBlobService().Return(blobService).After(c)
	c = blobService.EXPECT().GetContainerReference("config").Return(configContainer).After(c)
	c = configContainer.EXPECT().GetBlobReference("master-startup").Return(masterStartupBlob).After(c)
	c = masterStartupBlob.EXPECT().GetSASURI(gomock.Any()).Return("", nil).After(c)
	c = configContainer.EXPECT().GetBlobReference("worker-startup").Return(workerStartupBlob).After(c)
	c = workerStartupBlob.EXPECT().GetSASURI(gomock.Any()).Return("", nil).After(c)
	c = storageClient.EXPECT().GetBlobService().Return(blobService).After(c)
	c = blobService.EXPECT().GetContainerReference("config").Return(configContainer).After(c)
	c = configContainer.EXPECT().GetBlobReference("master-startup").Return(masterStartupBlob).After(c)
	c = masterStartupBlob.EXPECT().CreateBlockBlobFromReader(gomock.Any(), nil).After(c)
	c = storageClient.EXPECT().GetBlobService().Return(blobService).After(c)
	c = blobService.EXPECT().GetContainerReference("config").Return(configContainer).After(c)
	c = configContainer.EXPECT().GetBlobReference("worker-startup").Return(workerStartupBlob).After(c)
	c = workerStartupBlob.EXPECT().CreateBlockBlobFromReader(gomock.Any(), nil).After(c)
	// EnrichCertificatesFromVault
	c = kvc.EXPECT().GetSecret(ctx, "", ".", "").Return(keyvault.SecretBundle{Value: &tls.DummyBundle}, nil).After(c)
	c = kvc.EXPECT().GetSecret(ctx, "", ".", "").Return(keyvault.SecretBundle{Value: &tls.DummyBundle}, nil).After(c)
	// EnrichStorageAccountKeys
	c = accountsClient.EXPECT().ListKeys(ctx, "", "").Return(storage.AccountListKeysResult{Keys: &[]storage.AccountKey{{Value: to.StringPtr("")}}}, nil).After(c)
	c = accountsClient.EXPECT().ListKeys(ctx, "", "").Return(storage.AccountListKeysResult{Keys: &[]storage.AccountKey{{Value: to.StringPtr("")}}}, nil).After(c)
	// UpdateMasterAgentPool
	c = updateContainer.EXPECT().GetBlobReference("update").Return(updateBlob).After(c)
	c = updateBlob.EXPECT().Get(nil).Return(ioutil.NopCloser(bytes.NewReader(blob)), nil).After(c)
	if false {
		c = kubeclient.EXPECT().DeleteMaster("master-000000").Return(nil).After(c)
		c = vmc.EXPECT().Deallocate(ctx, "", "ss-master", "0").Return(nil).After(c)
		c = ssc.EXPECT().UpdateInstances(ctx, "", "ss-master", compute.VirtualMachineScaleSetVMInstanceRequiredIDs{InstanceIds: &[]string{"0"}}).Return(nil).After(c)
		c = vmc.EXPECT().Reimage(ctx, "", "ss-master", "0", nil).Return(nil).After(c)
		c = vmc.EXPECT().Start(ctx, "", "ss-master", "0").Return(nil).After(c)
		c = kubeclient.EXPECT().WaitForReadyMaster(ctx, "master-000000").Return(nil).After(c)
		c = updateContainer.EXPECT().GetBlobReference("update").Return(updateBlob).After(c)
		c = updateBlob.EXPECT().CreateBlockBlobFromReader(gomock.Any(), nil).DoAndReturn(func(r io.Reader, options *azstorage.PutBlobOptions) (err error) {
			blob, err = ioutil.ReadAll(r)
			return
		}).After(c)
	}
	// UpdateWorkerAgentPool
	c = updateContainer.EXPECT().GetBlobReference("update").Return(updateBlob).After(c)
	c = updateBlob.EXPECT().Get(nil).Return(ioutil.NopCloser(bytes.NewReader(blob)), nil).After(c)
	infraVmss := compute.VirtualMachineScaleSet{
		Sku: &compute.Sku{
			Capacity: to.Int64Ptr(1),
		},
		Name: to.StringPtr("ss-infra-0"),
	}
	c = ssc.EXPECT().List(ctx, "").Return([]compute.VirtualMachineScaleSet{infraVmss}, nil).After(c)
	if false {
		c = ssc.EXPECT().CreateOrUpdate(ctx, "", "ss-infra-1", gomock.Any()).After(c)
		c = updateContainer.EXPECT().GetBlobReference("update").Return(updateBlob).After(c)
		c = updateBlob.EXPECT().CreateBlockBlobFromReader(gomock.Any(), nil).DoAndReturn(func(r io.Reader, options *azstorage.PutBlobOptions) (err error) {
			blob, err = ioutil.ReadAll(r)
			return
		}).After(c)
	}
	c = scalerFactory.EXPECT().New(log, ssc, vmc, kubeclient, "", gomock.Any()).Return(targetScaler).After(c)
	if false {
		c = scalerFactory.EXPECT().New(log, ssc, vmc, kubeclient, "", gomock.Any()).Return(sourceScaler).After(c)
		c = targetScaler.EXPECT().Scale(ctx, int64(1)).Return(nil).After(c)
		c = sourceScaler.EXPECT().Scale(ctx, int64(0)).DoAndReturn(func(ctx context.Context, count int64) *api.PluginError {
			infraVmss.Sku.Capacity = to.Int64Ptr(0)
			return nil
		}).After(c)
		c = updateContainer.EXPECT().GetBlobReference("update").Return(updateBlob).After(c)
		c = updateBlob.EXPECT().CreateBlockBlobFromReader(gomock.Any(), nil).DoAndReturn(func(r io.Reader, options *azstorage.PutBlobOptions) (err error) {
			blob, err = ioutil.ReadAll(r)
			return
		}).After(c)
		c = ssc.EXPECT().Delete(ctx, "", "ss-infra-0").Return(nil).After(c)
	}
	c = targetScaler.EXPECT().Scale(ctx, int64(1)).Return(nil).After(c)
	// UpdateWorkerAgentPool
	c = updateContainer.EXPECT().GetBlobReference("update").Return(updateBlob).After(c)
	c = updateBlob.EXPECT().Get(nil).Return(ioutil.NopCloser(bytes.NewReader(blob)), nil).After(c)
	computeVmss := compute.VirtualMachineScaleSet{
		Sku: &compute.Sku{
			Capacity: to.Int64Ptr(1),
		},
		Name: to.StringPtr("ss-compute-0"),
	}
	c = ssc.EXPECT().List(ctx, "").Return([]compute.VirtualMachineScaleSet{computeVmss}, nil).After(c)
	if false {
		c = ssc.EXPECT().CreateOrUpdate(ctx, "", "ss-compute-1", gomock.Any()).After(c)
		c = updateContainer.EXPECT().GetBlobReference("update").Return(updateBlob).After(c)
		c = updateBlob.EXPECT().CreateBlockBlobFromReader(gomock.Any(), nil).DoAndReturn(func(r io.Reader, options *azstorage.PutBlobOptions) (err error) {
			blob, err = ioutil.ReadAll(r)
			return
		}).After(c)
	}
	c = scalerFactory.EXPECT().New(log, ssc, vmc, kubeclient, "", gomock.Any()).Return(targetScaler).After(c)
	if false {
		c = scalerFactory.EXPECT().New(log, ssc, vmc, kubeclient, "", gomock.Any()).Return(sourceScaler).After(c)
		c = targetScaler.EXPECT().Scale(ctx, int64(1)).Return(nil).After(c)
		c = sourceScaler.EXPECT().Scale(ctx, int64(0)).DoAndReturn(func(ctx context.Context, count int64) *api.PluginError {
			computeVmss.Sku.Capacity = to.Int64Ptr(0)
			return nil
		}).After(c)
		c = updateContainer.EXPECT().GetBlobReference("update").Return(updateBlob).After(c)
		c = updateBlob.EXPECT().CreateBlockBlobFromReader(gomock.Any(), nil).DoAndReturn(func(r io.Reader, options *azstorage.PutBlobOptions) (err error) {
			blob, err = ioutil.ReadAll(r)
			return
		}).After(c)
		c = ssc.EXPECT().Delete(ctx, "", "ss-compute-0").Return(nil).After(c)
	}
	c = targetScaler.EXPECT().Scale(ctx, int64(1)).Return(nil).After(c)
	// CreateOrUpdateSyncPod
	c = storageClient.EXPECT().GetBlobService().Return(blobService).After(c)
	c = blobService.EXPECT().GetContainerReference("config").Return(configContainer).After(c)
	c = configContainer.EXPECT().GetBlobReference("sync").Return(syncBlob).After(c)
	c = syncBlob.EXPECT().CreateBlockBlobFromReader(gomock.Any(), nil).After(c)
	c = kubeclient.EXPECT().EnsureSyncPod(ctx, "", gomock.Any()).Return(nil).After(c)
	// WaitForReadySyncPod
	c = kubeclient.EXPECT().WaitForReadySyncPod(ctx).Return(nil).After(c)
	// HealthCheck
	c = httpClient.EXPECT().Do(gomock.Any()).Return(&http.Response{StatusCode: http.StatusOK}, nil).After(c)

	p := &plugin{
		log: log,
		upgraderFactory: func(ctx context.Context, log *logrus.Entry, cs *api.OpenShiftManagedCluster, initializeStorageClients, disableKeepAlives bool, testConfig api.TestConfig) (cluster.Upgrader, error) {
			return cluster.NewFakeUpgrader(ctx, log, cs, accountsClient, vmc, ssc, kvc, storageClient, blobService, httpClient, kubeclient, scalerFactory, api.TestConfig{})
		},
		now: func() time.Time { return time.Unix(1, 0) },
	}
	err := p.CreateOrUpdate(ctx, cs, true, func(context.Context, map[string]interface{}) error { return nil })
	if err != nil {
		t.Error(err)
	}
}

func TestJim(t *testing.T) {
	cs := &api.OpenShiftManagedCluster{
		Properties: api.Properties{
			RouterProfiles: []api.RouterProfile{
				{},
			},
			AgentPoolProfiles: []api.AgentPoolProfile{
				{
					Name:  "master",
					Count: 1,
					Role:  api.AgentPoolProfileRoleMaster,
				},
				{
					Name:  "infra",
					Count: 1,
					Role:  api.AgentPoolProfileRoleInfra,
				},
				{
					Name:   "compute",
					Count:  1,
					Role:   api.AgentPoolProfileRoleCompute,
					VMSize: api.StandardD2sV3,
				},
			},
			AuthProfile: api.AuthProfile{
				IdentityProviders: []api.IdentityProvider{
					{
						Provider: &api.AADIdentityProvider{},
					},
				},
			},
		},
		Config: api.Config{
			PluginVersion: "v4.0",
			ImageVersion:  "311.0.0",
			Images: api.ImageConfig{
				AlertManager:             ":",
				ConfigReloader:           ":",
				Grafana:                  ":",
				KubeRbacProxy:            ":",
				KubeStateMetrics:         ":",
				NodeExporter:             ":",
				OAuthProxy:               ":",
				Prometheus:               ":",
				PrometheusConfigReloader: ":",
				PrometheusOperator:       ":",
			},
		},
	}
	populate.DummyCertsAndKeys(cs)

	updateBlob := create(t, cs)
	update(t, cs, updateBlob)
}
