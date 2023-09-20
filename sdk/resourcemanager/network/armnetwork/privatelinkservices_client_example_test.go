//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/PrivateLinkServiceDelete.json
func ExamplePrivateLinkServicesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateLinkServicesClient().BeginDelete(ctx, "rg1", "testPls", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/PrivateLinkServiceGet.json
func ExamplePrivateLinkServicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkServicesClient().Get(ctx, "rg1", "testPls", &armnetwork.PrivateLinkServicesClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkService = armnetwork.PrivateLinkService{
	// 	Name: to.Ptr("testPls"),
	// 	Type: to.Ptr("Microsoft.Network/privateLinkServices"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armnetwork.PrivateLinkServiceProperties{
	// 		Alias: to.Ptr("ContosoService.{guid}.azure.privatelinkservice"),
	// 		AutoApproval: &armnetwork.PrivateLinkServicePropertiesAutoApproval{
	// 			Subscriptions: []*string{
	// 				to.Ptr("subscription1"),
	// 				to.Ptr("subscription2")},
	// 			},
	// 			Fqdns: []*string{
	// 				to.Ptr("fqdn1"),
	// 				to.Ptr("fqdn2"),
	// 				to.Ptr("fqdn3")},
	// 				IPConfigurations: []*armnetwork.PrivateLinkServiceIPConfiguration{
	// 					{
	// 						Name: to.Ptr("fe-lb"),
	// 						Properties: &armnetwork.PrivateLinkServiceIPConfigurationProperties{
	// 							PrivateIPAddress: to.Ptr("10.0.1.4"),
	// 							PrivateIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
	// 							PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodStatic),
	// 							Subnet: &armnetwork.Subnet{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb"),
	// 							},
	// 						},
	// 				}},
	// 				LoadBalancerFrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
	// 				}},
	// 				NetworkInterfaces: []*armnetwork.Interface{
	// 					{
	// 						ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/provders/Microsoft.Network/networkInterfaces/testPls.nic.abcd1234"),
	// 				}},
	// 				PrivateEndpointConnections: []*armnetwork.PrivateEndpointConnection{
	// 					{
	// 						Name: to.Ptr("privateEndpointConnection"),
	// 						Properties: &armnetwork.PrivateEndpointConnectionProperties{
	// 							PrivateEndpoint: &armnetwork.PrivateEndpoint{
	// 								ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/testPe"),
	// 							},
	// 							PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
	// 								Description: to.Ptr("approved it for some reason."),
	// 								Status: to.Ptr("Approved"),
	// 							},
	// 						},
	// 				}},
	// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				Visibility: &armnetwork.PrivateLinkServicePropertiesVisibility{
	// 					Subscriptions: []*string{
	// 						to.Ptr("subscription1"),
	// 						to.Ptr("subscription2"),
	// 						to.Ptr("subscription3")},
	// 					},
	// 				},
	// 			}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/PrivateLinkServiceCreate.json
func ExamplePrivateLinkServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateLinkServicesClient().BeginCreateOrUpdate(ctx, "rg1", "testPls", armnetwork.PrivateLinkService{
		Location: to.Ptr("eastus"),
		Properties: &armnetwork.PrivateLinkServiceProperties{
			AutoApproval: &armnetwork.PrivateLinkServicePropertiesAutoApproval{
				Subscriptions: []*string{
					to.Ptr("subscription1"),
					to.Ptr("subscription2")},
			},
			Fqdns: []*string{
				to.Ptr("fqdn1"),
				to.Ptr("fqdn2"),
				to.Ptr("fqdn3")},
			IPConfigurations: []*armnetwork.PrivateLinkServiceIPConfiguration{
				{
					Name: to.Ptr("fe-lb"),
					Properties: &armnetwork.PrivateLinkServiceIPConfigurationProperties{
						PrivateIPAddress:          to.Ptr("10.0.1.4"),
						PrivateIPAddressVersion:   to.Ptr(armnetwork.IPVersionIPv4),
						PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodStatic),
						Subnet: &armnetwork.Subnet{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb"),
						},
					},
				}},
			LoadBalancerFrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
				{
					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
				}},
			Visibility: &armnetwork.PrivateLinkServicePropertiesVisibility{
				Subscriptions: []*string{
					to.Ptr("subscription1"),
					to.Ptr("subscription2"),
					to.Ptr("subscription3")},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkService = armnetwork.PrivateLinkService{
	// 	Name: to.Ptr("testPls"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armnetwork.PrivateLinkServiceProperties{
	// 		Alias: to.Ptr("ContosoService.{guid}.azure.privatelinkservice"),
	// 		AutoApproval: &armnetwork.PrivateLinkServicePropertiesAutoApproval{
	// 			Subscriptions: []*string{
	// 				to.Ptr("subscription1"),
	// 				to.Ptr("subscription2")},
	// 			},
	// 			Fqdns: []*string{
	// 				to.Ptr("fqdn1"),
	// 				to.Ptr("fqdn2"),
	// 				to.Ptr("fqdn3")},
	// 				IPConfigurations: []*armnetwork.PrivateLinkServiceIPConfiguration{
	// 					{
	// 						Name: to.Ptr("fe-lb"),
	// 						Properties: &armnetwork.PrivateLinkServiceIPConfigurationProperties{
	// 							PrivateIPAddress: to.Ptr("10.0.1.4"),
	// 							PrivateIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
	// 							PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodStatic),
	// 							Subnet: &armnetwork.Subnet{
	// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb"),
	// 							},
	// 						},
	// 				}},
	// 				LoadBalancerFrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
	// 					{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
	// 				}},
	// 				NetworkInterfaces: []*armnetwork.Interface{
	// 					{
	// 						ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/provders/Microsoft.Network/networkInterfaces/testPls.nic.abcd1234"),
	// 				}},
	// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				Visibility: &armnetwork.PrivateLinkServicePropertiesVisibility{
	// 					Subscriptions: []*string{
	// 						to.Ptr("subscription1"),
	// 						to.Ptr("subscription2"),
	// 						to.Ptr("subscription3")},
	// 					},
	// 				},
	// 			}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/PrivateLinkServiceList.json
func ExamplePrivateLinkServicesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkServicesClient().NewListPager("rg1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.PrivateLinkServiceListResult = armnetwork.PrivateLinkServiceListResult{
		// 	Value: []*armnetwork.PrivateLinkService{
		// 		{
		// 			Name: to.Ptr("testPls1"),
		// 			Type: to.Ptr("Microsoft.Network/privateLinkServices"),
		// 			ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls1"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armnetwork.PrivateLinkServiceProperties{
		// 				Alias: to.Ptr("ContosoService.{guid}.azure.privatelinkservice"),
		// 				AutoApproval: &armnetwork.PrivateLinkServicePropertiesAutoApproval{
		// 					Subscriptions: []*string{
		// 						to.Ptr("subscription1")},
		// 					},
		// 					Fqdns: []*string{
		// 						to.Ptr("fqdn1"),
		// 						to.Ptr("fqdn2")},
		// 						IPConfigurations: []*armnetwork.PrivateLinkServiceIPConfiguration{
		// 							{
		// 								Name: to.Ptr("fe-lb1"),
		// 								Properties: &armnetwork.PrivateLinkServiceIPConfigurationProperties{
		// 									PrivateIPAddress: to.Ptr("10.0.1.4"),
		// 									PrivateIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		// 									PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodStatic),
		// 									Subnet: &armnetwork.Subnet{
		// 										ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb1"),
		// 									},
		// 								},
		// 						}},
		// 						LoadBalancerFrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
		// 							{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb1"),
		// 						}},
		// 						NetworkInterfaces: []*armnetwork.Interface{
		// 							{
		// 								ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/provders/Microsoft.Network/networkInterfaces/testPls1.nic.abcd1234"),
		// 						}},
		// 						PrivateEndpointConnections: []*armnetwork.PrivateEndpointConnection{
		// 							{
		// 								Name: to.Ptr("pec1"),
		// 								Properties: &armnetwork.PrivateEndpointConnectionProperties{
		// 									PrivateEndpoint: &armnetwork.PrivateEndpoint{
		// 										ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/testPe1"),
		// 									},
		// 									PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
		// 										Description: to.Ptr("approved it for some reason."),
		// 										Status: to.Ptr("Approved"),
		// 									},
		// 								},
		// 						}},
		// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 						Visibility: &armnetwork.PrivateLinkServicePropertiesVisibility{
		// 							Subscriptions: []*string{
		// 								to.Ptr("subscription1")},
		// 							},
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("testPls2"),
		// 						Type: to.Ptr("Microsoft.Network/privateLinkServices"),
		// 						ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls2"),
		// 						Location: to.Ptr("eastus"),
		// 						Properties: &armnetwork.PrivateLinkServiceProperties{
		// 							Alias: to.Ptr("ContosoService.{guid}.azure.privatelinkservice"),
		// 							AutoApproval: &armnetwork.PrivateLinkServicePropertiesAutoApproval{
		// 								Subscriptions: []*string{
		// 									to.Ptr("subscription1"),
		// 									to.Ptr("subscription2")},
		// 								},
		// 								Fqdns: []*string{
		// 									to.Ptr("fqdn1"),
		// 									to.Ptr("fqdn2"),
		// 									to.Ptr("fqdn3")},
		// 									IPConfigurations: []*armnetwork.PrivateLinkServiceIPConfiguration{
		// 										{
		// 											Name: to.Ptr("fe-lb2"),
		// 											Properties: &armnetwork.PrivateLinkServiceIPConfigurationProperties{
		// 												PrivateIPAddress: to.Ptr("10.0.1.5"),
		// 												PrivateIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		// 												PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodStatic),
		// 												Subnet: &armnetwork.Subnet{
		// 													ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb2"),
		// 												},
		// 											},
		// 									}},
		// 									LoadBalancerFrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
		// 										{
		// 											ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb2"),
		// 									}},
		// 									NetworkInterfaces: []*armnetwork.Interface{
		// 										{
		// 											ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/provders/Microsoft.Network/networkInterfaces/testPls2.nic.efgh5678"),
		// 									}},
		// 									PrivateEndpointConnections: []*armnetwork.PrivateEndpointConnection{
		// 										{
		// 											Name: to.Ptr("pec2"),
		// 											Properties: &armnetwork.PrivateEndpointConnectionProperties{
		// 												PrivateEndpoint: &armnetwork.PrivateEndpoint{
		// 													ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/testPe2"),
		// 												},
		// 												PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
		// 													Description: to.Ptr("approved it for some reason."),
		// 													Status: to.Ptr("Approved"),
		// 												},
		// 											},
		// 									}},
		// 									ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 									Visibility: &armnetwork.PrivateLinkServicePropertiesVisibility{
		// 										Subscriptions: []*string{
		// 											to.Ptr("subscription1"),
		// 											to.Ptr("subscription2"),
		// 											to.Ptr("subscription3")},
		// 										},
		// 									},
		// 							}},
		// 						}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/PrivateLinkServiceListAll.json
func ExamplePrivateLinkServicesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkServicesClient().NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.PrivateLinkServiceListResult = armnetwork.PrivateLinkServiceListResult{
		// 	Value: []*armnetwork.PrivateLinkService{
		// 		{
		// 			Name: to.Ptr("testPls1"),
		// 			Type: to.Ptr("Microsoft.Network/privateLinkServices"),
		// 			ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls1"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armnetwork.PrivateLinkServiceProperties{
		// 				Alias: to.Ptr("ContosoService.{guid}.azure.privatelinkservice"),
		// 				AutoApproval: &armnetwork.PrivateLinkServicePropertiesAutoApproval{
		// 					Subscriptions: []*string{
		// 						to.Ptr("subscription1"),
		// 						to.Ptr("subscription2")},
		// 					},
		// 					Fqdns: []*string{
		// 						to.Ptr("fqdn1"),
		// 						to.Ptr("fqdn2"),
		// 						to.Ptr("fqdn3")},
		// 						IPConfigurations: []*armnetwork.PrivateLinkServiceIPConfiguration{
		// 							{
		// 								Name: to.Ptr("fe-lb1"),
		// 								Properties: &armnetwork.PrivateLinkServiceIPConfigurationProperties{
		// 									PrivateIPAddress: to.Ptr("10.0.1.4"),
		// 									PrivateIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		// 									PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodStatic),
		// 									Subnet: &armnetwork.Subnet{
		// 										ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb1"),
		// 									},
		// 								},
		// 						}},
		// 						LoadBalancerFrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
		// 							{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb1"),
		// 						}},
		// 						NetworkInterfaces: []*armnetwork.Interface{
		// 							{
		// 								ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/provders/Microsoft.Network/networkInterfaces/testPls1.nic.abcd1234"),
		// 						}},
		// 						PrivateEndpointConnections: []*armnetwork.PrivateEndpointConnection{
		// 							{
		// 								Name: to.Ptr("pec1"),
		// 								Properties: &armnetwork.PrivateEndpointConnectionProperties{
		// 									PrivateEndpoint: &armnetwork.PrivateEndpoint{
		// 										ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/testPe1"),
		// 									},
		// 									PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
		// 										Description: to.Ptr("approved it for some reason."),
		// 										Status: to.Ptr("Approved"),
		// 									},
		// 								},
		// 						}},
		// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 						Visibility: &armnetwork.PrivateLinkServicePropertiesVisibility{
		// 							Subscriptions: []*string{
		// 								to.Ptr("subscription1"),
		// 								to.Ptr("subscription2"),
		// 								to.Ptr("subscription3")},
		// 							},
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("testPls2"),
		// 						Type: to.Ptr("Microsoft.Network/privateLinkServices"),
		// 						ID: to.Ptr("/subscriptions/subId/resourceGroups/rg2/providers/Microsoft.Network/privateLinkServices/testPls2"),
		// 						Location: to.Ptr("eastus"),
		// 						Properties: &armnetwork.PrivateLinkServiceProperties{
		// 							Alias: to.Ptr("ContosoService.{guid}.azure.privatelinkservice"),
		// 							AutoApproval: &armnetwork.PrivateLinkServicePropertiesAutoApproval{
		// 								Subscriptions: []*string{
		// 									to.Ptr("subscription1"),
		// 									to.Ptr("subscription2")},
		// 								},
		// 								Fqdns: []*string{
		// 									to.Ptr("fqdn1"),
		// 									to.Ptr("fqdn2")},
		// 									IPConfigurations: []*armnetwork.PrivateLinkServiceIPConfiguration{
		// 										{
		// 											Name: to.Ptr("fe-lb2"),
		// 											Properties: &armnetwork.PrivateLinkServiceIPConfigurationProperties{
		// 												PrivateIPAddress: to.Ptr("10.0.1.5"),
		// 												PrivateIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		// 												PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodStatic),
		// 												Subnet: &armnetwork.Subnet{
		// 													ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb2"),
		// 												},
		// 											},
		// 									}},
		// 									LoadBalancerFrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
		// 										{
		// 											ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb2"),
		// 									}},
		// 									NetworkInterfaces: []*armnetwork.Interface{
		// 										{
		// 											ID: to.Ptr("/subscriptions/subId/resourceGroups/rg2/provders/Microsoft.Network/networkInterfaces/testPls2.nic.efgh5678"),
		// 									}},
		// 									PrivateEndpointConnections: []*armnetwork.PrivateEndpointConnection{
		// 										{
		// 											Name: to.Ptr("pec1"),
		// 											Properties: &armnetwork.PrivateEndpointConnectionProperties{
		// 												PrivateEndpoint: &armnetwork.PrivateEndpoint{
		// 													ID: to.Ptr("/subscriptions/subId/resourceGroups/rg2/providers/Microsoft.Network/privateEndpoints/testPe2"),
		// 												},
		// 												PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
		// 													Description: to.Ptr("approved it for some reason."),
		// 													Status: to.Ptr("Approved"),
		// 												},
		// 											},
		// 									}},
		// 									ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 									Visibility: &armnetwork.PrivateLinkServicePropertiesVisibility{
		// 										Subscriptions: []*string{
		// 											to.Ptr("subscription1"),
		// 											to.Ptr("subscription2")},
		// 										},
		// 									},
		// 							}},
		// 						}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/PrivateLinkServiceGetPrivateEndpointConnection.json
func ExamplePrivateLinkServicesClient_GetPrivateEndpointConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkServicesClient().GetPrivateEndpointConnection(ctx, "rg1", "testPls", "testPlePeConnection", &armnetwork.PrivateLinkServicesClientGetPrivateEndpointConnectionOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armnetwork.PrivateEndpointConnection{
	// 	Name: to.Ptr("testPlePeConnection"),
	// 	Properties: &armnetwork.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armnetwork.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/testPe"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("approved it for some reason."),
	// 			Status: to.Ptr("Approved"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/PrivateLinkServiceUpdatePrivateEndpointConnection.json
func ExamplePrivateLinkServicesClient_UpdatePrivateEndpointConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkServicesClient().UpdatePrivateEndpointConnection(ctx, "rg1", "testPls", "testPlePeConnection", armnetwork.PrivateEndpointConnection{
		Name: to.Ptr("testPlePeConnection"),
		Properties: &armnetwork.PrivateEndpointConnectionProperties{
			PrivateEndpoint: &armnetwork.PrivateEndpoint{
				ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/testPe"),
			},
			PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
				Description: to.Ptr("approved it for some reason."),
				Status:      to.Ptr("Approved"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armnetwork.PrivateEndpointConnection{
	// 	Name: to.Ptr("testPlePeConnection"),
	// 	Properties: &armnetwork.PrivateEndpointConnectionProperties{
	// 		PrivateEndpoint: &armnetwork.PrivateEndpoint{
	// 			ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/testPe"),
	// 		},
	// 		PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
	// 			Description: to.Ptr("approved it for some reason."),
	// 			Status: to.Ptr("Approved"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/PrivateLinkServiceDeletePrivateEndpointConnection.json
func ExamplePrivateLinkServicesClient_BeginDeletePrivateEndpointConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateLinkServicesClient().BeginDeletePrivateEndpointConnection(ctx, "rg1", "testPls", "testPlePeConnection", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/PrivateLinkServiceListPrivateEndpointConnection.json
func ExamplePrivateLinkServicesClient_NewListPrivateEndpointConnectionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkServicesClient().NewListPrivateEndpointConnectionsPager("rg1", "testPls", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.PrivateEndpointConnectionListResult = armnetwork.PrivateEndpointConnectionListResult{
		// 	Value: []*armnetwork.PrivateEndpointConnection{
		// 		{
		// 			Name: to.Ptr("testPlePeConnection1"),
		// 			Properties: &armnetwork.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armnetwork.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/testPe1"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("approved it for some reason."),
		// 					Status: to.Ptr("Approved"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testPlePeConnection2"),
		// 			Properties: &armnetwork.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armnetwork.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/testPe2"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armnetwork.PrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("rejected by some reason."),
		// 					Status: to.Ptr("Rejected"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/CheckPrivateLinkServiceVisibility.json
func ExamplePrivateLinkServicesClient_BeginCheckPrivateLinkServiceVisibility() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateLinkServicesClient().BeginCheckPrivateLinkServiceVisibility(ctx, "westus", armnetwork.CheckPrivateLinkServiceVisibilityRequest{
		PrivateLinkServiceAlias: to.Ptr("mypls.00000000-0000-0000-0000-000000000000.azure.privatelinkservice"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkServiceVisibility = armnetwork.PrivateLinkServiceVisibility{
	// 	Visible: to.Ptr(true),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/CheckPrivateLinkServiceVisibilityByResourceGroup.json
func ExamplePrivateLinkServicesClient_BeginCheckPrivateLinkServiceVisibilityByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateLinkServicesClient().BeginCheckPrivateLinkServiceVisibilityByResourceGroup(ctx, "westus", "rg1", armnetwork.CheckPrivateLinkServiceVisibilityRequest{
		PrivateLinkServiceAlias: to.Ptr("mypls.00000000-0000-0000-0000-000000000000.azure.privatelinkservice"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkServiceVisibility = armnetwork.PrivateLinkServiceVisibility{
	// 	Visible: to.Ptr(true),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/AutoApprovedPrivateLinkServicesGet.json
func ExamplePrivateLinkServicesClient_NewListAutoApprovedPrivateLinkServicesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkServicesClient().NewListAutoApprovedPrivateLinkServicesPager("regionName", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AutoApprovedPrivateLinkServicesResult = armnetwork.AutoApprovedPrivateLinkServicesResult{
		// 	Value: []*armnetwork.AutoApprovedPrivateLinkService{
		// 		{
		// 			PrivateLinkService: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls1"),
		// 		},
		// 		{
		// 			PrivateLinkService: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls2"),
		// 		},
		// 		{
		// 			PrivateLinkService: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls3"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/AutoApprovedPrivateLinkServicesResourceGroupGet.json
func ExamplePrivateLinkServicesClient_NewListAutoApprovedPrivateLinkServicesByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkServicesClient().NewListAutoApprovedPrivateLinkServicesByResourceGroupPager("regionName", "rg1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AutoApprovedPrivateLinkServicesResult = armnetwork.AutoApprovedPrivateLinkServicesResult{
		// 	Value: []*armnetwork.AutoApprovedPrivateLinkService{
		// 		{
		// 			PrivateLinkService: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls1"),
		// 		},
		// 		{
		// 			PrivateLinkService: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls2"),
		// 		},
		// 		{
		// 			PrivateLinkService: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls3"),
		// 	}},
		// }
	}
}
