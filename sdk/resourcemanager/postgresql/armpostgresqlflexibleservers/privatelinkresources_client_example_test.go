//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/PrivateLinkResourcesList.json
func ExamplePrivateLinkResourcesClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkResourcesClient().NewListByServerPager("Default", "test-svr", nil)
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
		// page.PrivateLinkResourceListResult = armpostgresqlflexibleservers.PrivateLinkResourceListResult{
		// 	Value: []*armpostgresqlflexibleservers.PrivateLinkResource{
		// 		{
		// 			Name: to.Ptr("plr"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/privateLinkResources"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/Default/providers/Microsoft.DBforPostgreSQL/flexibleServers/test-svr/privateLinkResources/plr"),
		// 			Properties: &armpostgresqlflexibleservers.PrivateLinkResourceProperties{
		// 				GroupID: to.Ptr("postgresqlServer"),
		// 				RequiredMembers: []*string{
		// 					to.Ptr("postgresqlServer")},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelink.meru.com")},
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/PrivateLinkResourcesGet.json
func ExamplePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().Get(ctx, "Default", "test-svr", "plr", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResource = armpostgresqlflexibleservers.PrivateLinkResource{
	// 	Name: to.Ptr("plr"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/privateLinkResources"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/Default/providers/Microsoft.DBforPostgreSQL/flexibleServers/test-svr/privateLinkResources/plr"),
	// 	Properties: &armpostgresqlflexibleservers.PrivateLinkResourceProperties{
	// 		GroupID: to.Ptr("postgresqlServer"),
	// 		RequiredMembers: []*string{
	// 			to.Ptr("postgresqlServer")},
	// 			RequiredZoneNames: []*string{
	// 				to.Ptr("privatelink.meru.com")},
	// 			},
	// 		}
}
