package realrp

import (
	"context"
	"errors"
	"fmt"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-06-01/compute"
	"github.com/Azure/go-autorest/autorest/to"

	"github.com/openshift/openshift-azure/test/clients/azure"
)

const (
	fakepubkey = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC7laRyN4B3YZmVrDEZLZoIuUA72pQ0DpGuZBZWykCofIfCPrFZAJgFvonKGgKJl6FGKIunkZL9Us/mV4ZPkZhBlE7uX83AAf5i9Q8FmKpotzmaxN10/1mcnEE7pFvLoSkwqrQSkrrgSm8zaJ3g91giXSbtqvSIj/vk2f05stYmLfhAwNo3Oh27ugCakCoVeuCrZkvHMaJgcYrIGCuFo6q0Pfk9rsZyriIqEa9AtiUOtViInVYdby7y71wcbl0AbbCZsTSqnSoVxm2tRkOsXV6+8X4SnwcmZbao3H+zfO1GBhQOLxJ4NQbzAa8IJh810rYARNLptgmsd4cYXVOSosTX azureuser"
)

var _ = Describe("Resource provider e2e tests [Real]", func() {
	var (
		ctx = context.Background()
		cli *azure.Client
	)

	BeforeEach(func() {
		var err error
		cli, err = azure.NewClientFromEnvironment()
		Expect(err).ToNot(HaveOccurred())
		if os.Getenv("AZURE_REGION") == "" || os.Getenv("RESOURCEGROUP") == "" {
			Expect(errors.New("AZURE_REGION and/or RESOURCEGROUP not set")).ToNot(HaveOccurred())
		}
	})

	It("should keep the end user from deleting any azure resources", func() {
		resourcegroup, err := cli.OSAResourceGroup(ctx, os.Getenv("RESOURCEGROUP"), os.Getenv("RESOURCEGROUP"), os.Getenv("AZURE_REGION"))
		Expect(err).NotTo(HaveOccurred())
		By(fmt.Sprintf("OSA resource group is %s", resourcegroup))

		pages, err := cli.Resources.ListByResourceGroup(ctx, resourcegroup, "", "", nil)
		Expect(err).ToNot(HaveOccurred())
		// attempt to delete all resources in the resourcegroup
		for pages.NotDone() {
			for _, v := range pages.Values() {
				By(fmt.Sprintf("Attempting to delete %s/%s", *v.Type, *v.Name))
				_, err := cli.Resources.DeleteByID(ctx, *v.ID)
				Expect(err).To(HaveOccurred())
			}
			err = pages.Next()
			Expect(err).To(HaveOccurred())
		}
	})

	It("should keep the end user from reading the config blob", func() {
		resourcegroup, err := cli.OSAResourceGroup(ctx, os.Getenv("RESOURCEGROUP"), os.Getenv("RESOURCEGROUP"), os.Getenv("AZURE_REGION"))
		Expect(err).NotTo(HaveOccurred())
		By(fmt.Sprintf("OSA resource group is %s", resourcegroup))

		accts, err := cli.Accounts.ListByResourceGroup(ctx, resourcegroup)
		Expect(err).NotTo(HaveOccurred())

		for _, acct := range *accts.Value {
			By(fmt.Sprintf("trying to read account %s", *acct.Name))
			_, err := cli.Accounts.ListKeys(ctx, resourcegroup, *acct.Name)
			Expect(err).To(HaveOccurred())
		}
	})

	It("should not be possible for customer to mutate an osa scale set", func() {
		resourcegroup, err := cli.OSAResourceGroup(ctx, os.Getenv("RESOURCEGROUP"), os.Getenv("RESOURCEGROUP"), os.Getenv("AZURE_REGION"))
		Expect(err).NotTo(HaveOccurred())
		By(fmt.Sprintf("OSA resource group is %s", resourcegroup))

		scaleSets, err := cli.ListScaleSets(ctx, resourcegroup)
		Expect(err).NotTo(HaveOccurred())
		Expect(scaleSets).Should(HaveLen(3))

		for _, scaleSet := range scaleSets {
			vms, err := cli.ListScaleSetVMs(ctx, resourcegroup, *scaleSet.Name)
			Expect(err).NotTo(HaveOccurred())

			By("trying to update the scale set capacity")
			_, err = cli.VirtualMachineScaleSets.Update(ctx, resourcegroup, *scaleSet.Name, compute.VirtualMachineScaleSetUpdate{
				Sku: &compute.Sku{
					Capacity: to.Int64Ptr(int64(len(vms) + 1)),
				},
			})
			Expect(err).To(HaveOccurred())

			By("trying to update the scale set type")
			_, err = cli.VirtualMachineScaleSets.Update(ctx, resourcegroup, *scaleSet.Name, compute.VirtualMachineScaleSetUpdate{
				Sku: &compute.Sku{
					Name: to.StringPtr("Standard_DS1_v2"),
				},
			})
			Expect(err).To(HaveOccurred())

			By("trying to update the scale set SSH key")
			_, err = cli.VirtualMachineScaleSets.Update(ctx, resourcegroup, *scaleSet.Name, compute.VirtualMachineScaleSetUpdate{
				VirtualMachineScaleSetUpdateProperties: &compute.VirtualMachineScaleSetUpdateProperties{
					VirtualMachineProfile: &compute.VirtualMachineScaleSetUpdateVMProfile{
						OsProfile: &compute.VirtualMachineScaleSetUpdateOSProfile{
							LinuxConfiguration: &compute.LinuxConfiguration{
								SSH: &compute.SSHConfiguration{
									PublicKeys: &[]compute.SSHPublicKey{
										{
											Path:    to.StringPtr("/home/cloud-user/.ssh/authorized_keys"),
											KeyData: to.StringPtr(fakepubkey),
										},
									},
								},
							},
						},
					},
				},
			})

			Expect(err).To(HaveOccurred())

			By("trying to create scale set script extension")
			_, err = cli.VirtualMachineScaleSetExtensions.CreateOrUpdate(ctx, resourcegroup, *scaleSet.Name, "test", compute.VirtualMachineScaleSetExtension{
				VirtualMachineScaleSetExtensionProperties: &compute.VirtualMachineScaleSetExtensionProperties{
					Type:     to.StringPtr("CustomScript"),
					Settings: `{"fileUris":["https://raw.githubusercontent.com/Azure-Samples/compute-automation-configurations/master/automate_nginx.sh"],"commandToExecute":"./automate_nginx.sh"}`,
				},
			})
			Expect(err).To(HaveOccurred())

			for _, vm := range vms {
				By("trying to restart scale set instance vm")
				_, err = cli.VirtualMachineScaleSetVMs.Restart(ctx, resourcegroup, *scaleSet.Name, *vm.InstanceID)
				Expect(err).To(HaveOccurred())
			}
		}
	})
})
