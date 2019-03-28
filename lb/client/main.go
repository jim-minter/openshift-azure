package main

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-06-01/network"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/go-autorest/autorest/to"
)

var (
	resourceGroup = flag.String("resourceGroup", "", "resource group")

	template = []byte(`{
	"$schema": "https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
	"contentVersion": "1.0.0.0",
	"parameters": {
		"domainNameLabel": {
			"type": "string"
		},
		"keyData": {
			"type": "string"
		},
		"script": {
			"type": "string"
		}
	},
	"variables": {},
	"resources": [
		{
			"type": "Microsoft.Network/virtualNetworks",
			"name": "vnet",
			"apiVersion": "2018-07-01",
			"location": "eastus",
			"properties": {
				"addressSpace": {
					"addressPrefixes": [
						"10.0.0.0/8"
					]
				},
				"subnets": [
					{
						"name": "default",
						"properties": {
							"addressPrefix": "10.0.0.0/24"
						}
					}
				]
			}
		},
		{
			"type": "Microsoft.Network/publicIPAddresses",
			"sku": {
				"name": "Standard"
			},
			"name": "ip-apiserver",
			"apiVersion": "2018-07-01",
			"location": "eastus",
			"properties": {
				"dnsSettings": {
					"domainNameLabel": "[parameters('domainNameLabel')]"
				},
				"idleTimeoutInMinutes": 15,
				"publicIPAllocationMethod": "Static"
			}
		},
		{
			"type": "Microsoft.Network/loadBalancers",
			"sku": {
				"name": "Standard"
			},
			"name": "lb-apiserver",
			"apiVersion": "2018-07-01",
			"location": "eastus",
			"properties": {
				"backendAddressPools": [
					{
						"name": "backend"
					}
				],
				"frontendIPConfigurations": [
					{
						"name": "frontend",
						"properties": {
							"privateIPAllocationMethod": "Dynamic",
							"publicIPAddress": {
								"id": "[resourceId('Microsoft.Network/publicIPAddresses', 'ip-apiserver')]"
							}
						}
					}
				],
				"loadBalancingRules": [
					{
						"name": "port-443",
						"properties": {
							"backendAddressPool": {
								"id": "[concat(resourceId('Microsoft.Network/loadBalancers', 'lb-apiserver'), '/backendAddressPools/backend')]"
							},
							"backendPort": 443,
							"enableFloatingIP": false,
							"frontendIPConfiguration": {
								"id": "[concat(resourceId('Microsoft.Network/loadBalancers', 'lb-apiserver'), '/frontendIPConfigurations/frontend')]"
							},
							"frontendPort": 443,
							"idleTimeoutInMinutes": 15,
							"loadDistribution": "Default",
							"probe": {
								"id": "[concat(resourceId('Microsoft.Network/loadBalancers', 'lb-apiserver'), '/probes/port-443')]"
							},
							"protocol": "Tcp"
						}
					}
				],
				"probes": [
					{
						"name": "port-443",
						"properties": {
							"intervalInSeconds": 5,
							"numberOfProbes": 2,
							"port": 443,
							"protocol": "Https",
							"requestPath": "/healthz"
						}
					}
				]
			},
			"dependsOn": [
				"[resourceId('Microsoft.Network/publicIPAddresses', 'ip-apiserver')]"
			]
		},
		{
			"type": "Microsoft.Network/networkSecurityGroups",
			"name": "nsg-master",
			"apiVersion": "2018-07-01",
			"location": "eastus",
			"properties": {
				"securityRules": [
					{
						"name": "allow_ssh",
						"properties": {
							"access": "Allow",
							"description": "Allow SSH traffic",
							"destinationAddressPrefix": "*",
							"destinationPortRange": "22-22",
							"direction": "Inbound",
							"priority": 101,
							"protocol": "Tcp",
							"sourceAddressPrefixes": [
								"0.0.0.0/0"
							],
							"sourcePortRange": "*"
						}
					},
					{
						"name": "allow_https",
						"properties": {
							"access": "Allow",
							"description": "Allow HTTPS traffic",
							"destinationAddressPrefix": "*",
							"destinationPortRange": "443-443",
							"direction": "Inbound",
							"priority": 102,
							"protocol": "Tcp",
							"sourceAddressPrefixes": [
								"0.0.0.0/0"
							],
							"sourcePortRange": "*"
						}
					}
				]
			}
		},
		{
			"type": "Microsoft.Compute/virtualMachineScaleSets",
			"sku": {
				"capacity": 3,
				"name": "Standard_D2s_v3",
				"tier": "Standard"
			},
			"name": "ss-master",
			"apiVersion": "2018-10-01",
			"location": "eastus",
			"properties": {
                "upgradePolicy": {
                    "mode": "Manual"
                },
				"virtualMachineProfile": {
					"extensionProfile": {
						"extensions": [
							{
								"name": "cse",
								"properties": {
									"autoUpgradeMinorVersion": true,
									"protectedSettings": {
										"script": "[parameters('script')]"
									},
									"publisher": "Microsoft.Azure.Extensions",
									"type": "CustomScript",
									"typeHandlerVersion": "2.0"
								}
							}
						]
					},
					"networkProfile": {
						"networkInterfaceConfigurations": [
							{
								"name": "nic",
								"properties": {
									"enableIPForwarding": true,
									"ipConfigurations": [
										{
											"name": "ipconfig",
											"properties": {
												"loadBalancerBackendAddressPools": [
													{
														"id": "[concat(resourceId('Microsoft.Network/loadBalancers', 'lb-apiserver'), '/backendAddressPools/backend')]"
													}
												],
												"primary": true,
												"publicIPAddressConfiguration": {
													"name": "ip",
													"properties": {
														"idleTimeoutInMinutes": 15
													}
												},
												"subnet": {
													"id": "[concat(resourceId('Microsoft.Network/virtualNetworks', 'vnet'), '/subnets/default')]"
												}
											}
										}
									],
									"networkSecurityGroup": {
										"id": "[resourceId('Microsoft.Network/networkSecurityGroups', 'nsg-master')]"
									},
									"primary": true
								}
							}
						]
					},
					"osProfile": {
						"adminUsername": "cloud-user",
						"computerNamePrefix": "master-",
						"linuxConfiguration": {
							"disablePasswordAuthentication": true,
							"ssh": {
								"publicKeys": [
									{
										"keyData": "[parameters('keyData')]",
										"path": "/home/cloud-user/.ssh/authorized_keys"
									}
								]
							}
						}
					},
					"storageProfile": {
						"imageReference": {
							"offer": "RHEL",
							"publisher": "RedHat",
							"sku": "7-RAW",
							"version": "latest"
						},
						"osDisk": {
							"caching": "ReadWrite",
							"createOption": "FromImage",
							"managedDisk": {
								"storageAccountType": "Premium_LRS"
							}
						}
					}
				}
			},
			"dependsOn": [
				"[resourceId('Microsoft.Network/loadBalancers', 'lb-apiserver')]",
				"[resourceId('Microsoft.Network/networkSecurityGroups', 'nsg-master')]",
				"[resourceId('Microsoft.Network/virtualNetworks', 'vnet')]"
			]
		}
	],
	"outputs": {}
}`)

	script = []byte(`#!/bin/bash -ex

exec 2>&1

export HOME=/root
cd

yum -y update -x WALinuxAgent
yum -y install git golang

firewall-cmd --zone=public --add-port=443/tcp

git clone -b lb https://github.com/jim-minter/openshift-azure.git go/src/github.com/openshift/openshift-azure
go run ./go/src/github.com/openshift/openshift-azure/lb/server/main.go &>/var/log/server &
`)
)

type client struct {
	groups            resources.GroupsClient
	deployments       resources.DeploymentsClient
	publicIPAddresses network.PublicIPAddressesClient
}

func newClient() (*client, error) {
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		return nil, err
	}

	groups := resources.NewGroupsClient(os.Getenv("AZURE_SUBSCRIPTION_ID"))
	groups.Authorizer = authorizer
	deployments := resources.NewDeploymentsClient(os.Getenv("AZURE_SUBSCRIPTION_ID"))
	deployments.Authorizer = authorizer
	publicIPAddresses := network.NewPublicIPAddressesClient(os.Getenv("AZURE_SUBSCRIPTION_ID"))
	publicIPAddresses.Authorizer = authorizer

	return &client{
		groups:            groups,
		deployments:       deployments,
		publicIPAddresses: publicIPAddresses,
	}, nil
}

func (c *client) deploy(ctx context.Context) error {
	_, err := c.groups.CreateOrUpdate(ctx, *resourceGroup, resources.Group{
		Location: to.StringPtr("eastus"),
	})
	if err != nil {
		return err
	}

	sshPublicKey, err := ioutil.ReadFile(os.Getenv("HOME") + "/.ssh/id_rsa.pub")
	if err != nil {
		return err
	}

	future, err := c.deployments.CreateOrUpdate(ctx, *resourceGroup, "azuredeploy", resources.Deployment{
		Properties: &resources.DeploymentProperties{
			Template: json.RawMessage(template),
			Parameters: map[string]interface{}{
				"domainNameLabel": map[string]interface{}{"value": *resourceGroup},
				"keyData":         map[string]interface{}{"value": string(sshPublicKey)},
				"script":          map[string]interface{}{"value": base64.StdEncoding.EncodeToString(script)},
			},
			Mode: resources.Incremental,
		},
	})
	if err != nil {
		return err
	}

	return future.WaitForCompletionRef(ctx, c.deployments.Client)
}

func (c *client) getURLs(ctx context.Context) ([]string, error) {
	urls := []string{
		"https://" + *resourceGroup + ".eastus.cloudapp.azure.com/healthz",
	}
	ips, err := c.publicIPAddresses.ListVirtualMachineScaleSetPublicIPAddressesComplete(ctx, *resourceGroup, "ss-master")
	if err != nil {
		return nil, err
	}
	for ips.NotDone() {
		urls = append(urls, "https://"+*ips.Value().IPAddress+"/healthz")
		err = ips.Next()
		if err != nil {
			return nil, err
		}
	}

	return urls, nil
}

func monitor(url string) error {
	cli := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			DisableKeepAlives: true,
		},
		Timeout: time.Second,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	go func() {
		t := time.NewTicker(100 * time.Millisecond)

		for {
			start := time.Now()

			resp, err := cli.Do(req)

			end := time.Now()

			if err == nil && resp.StatusCode != http.StatusOK {
				err = fmt.Errorf("invalid status %d", resp.StatusCode)
			}

			log.Printf("%-60s %4dms %v", url, end.Sub(start)/time.Millisecond, err)

			<-t.C
		}
	}()

	return nil
}

func run(ctx context.Context) error {
	cli, err := newClient()
	if err != nil {
		return err
	}

	err = cli.deploy(ctx)
	if err != nil {
		return err
	}

	urls, err := cli.getURLs(ctx)
	if err != nil {
		return err
	}

	for _, url := range urls {
		err = monitor(url)
		if err != nil {
			return err
		}
	}

	select {}
}

func main() {
	flag.Parse()

	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
