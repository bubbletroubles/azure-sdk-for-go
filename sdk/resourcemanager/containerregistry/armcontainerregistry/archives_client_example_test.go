//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-11-01-preview/examples/ArchiveList.json
func ExampleArchivesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewArchivesClient().NewListPager("myResourceGroup", "myRegistry", "myPackageType", nil)
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
		// page.ArchiveListResult = armcontainerregistry.ArchiveListResult{
		// 	Value: []*armcontainerregistry.Archive{
		// 		{
		// 			Name: to.Ptr("string"),
		// 			Type: to.Ptr("string"),
		// 			ID: to.Ptr("string"),
		// 			SystemData: &armcontainerregistry.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-24T00:22:47.311Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armcontainerregistry.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-24T00:22:47.311Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armcontainerregistry.LastModifiedByTypeUser),
		// 			},
		// 			Properties: &armcontainerregistry.ArchiveProperties{
		// 				PackageSource: &armcontainerregistry.ArchivePackageSourceProperties{
		// 					Type: to.Ptr(armcontainerregistry.PackageSourceType("string")),
		// 					URL: to.Ptr("string"),
		// 				},
		// 				ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
		// 				PublishedVersion: to.Ptr("string"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-11-01-preview/examples/ArchiveGet.json
func ExampleArchivesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewArchivesClient().Get(ctx, "myResourceGroup", "myRegistry", "myPackageType", "myArchiveName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Archive = armcontainerregistry.Archive{
	// 	Name: to.Ptr("myArchiveName"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/packages/archives"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/packages/myPackageType/archives/myArchiveName"),
	// 	SystemData: &armcontainerregistry.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-24T00:13:46.836Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armcontainerregistry.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-24T00:13:46.836Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armcontainerregistry.LastModifiedByTypeUser),
	// 	},
	// 	Properties: &armcontainerregistry.ArchiveProperties{
	// 		PackageSource: &armcontainerregistry.ArchivePackageSourceProperties{
	// 			Type: to.Ptr(armcontainerregistry.PackageSourceType("string")),
	// 			URL: to.Ptr("string"),
	// 		},
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		PublishedVersion: to.Ptr("string"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-11-01-preview/examples/ArchiveCreate.json
func ExampleArchivesClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewArchivesClient().BeginCreate(ctx, "myResourceGroup", "myRegistry", "rpm", "myArchiveName", armcontainerregistry.Archive{
		Properties: &armcontainerregistry.ArchiveProperties{
			PackageSource: &armcontainerregistry.ArchivePackageSourceProperties{
				Type: to.Ptr(armcontainerregistry.PackageSourceTypeRemote),
				URL:  to.Ptr("string"),
			},
			PublishedVersion:         to.Ptr("string"),
			RepositoryEndpointPrefix: to.Ptr("string"),
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
	// res.Archive = armcontainerregistry.Archive{
	// 	Name: to.Ptr("myArchiveName"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/archives"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/packages/rpm/archives/myArchiveName"),
	// 	Properties: &armcontainerregistry.ArchiveProperties{
	// 		PackageSource: &armcontainerregistry.ArchivePackageSourceProperties{
	// 			Type: to.Ptr(armcontainerregistry.PackageSourceTypeRemote),
	// 			URL: to.Ptr("string"),
	// 		},
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		PublishedVersion: to.Ptr("string"),
	// 		RepositoryEndpoint: to.Ptr("string"),
	// 		RepositoryEndpointPrefix: to.Ptr("string"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-11-01-preview/examples/ArchiveDelete.json
func ExampleArchivesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewArchivesClient().BeginDelete(ctx, "myResourceGroup", "myRegistry", "myPackageType", "myArchiveName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2023-11-01-preview/examples/ArchiveUpdate.json
func ExampleArchivesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewArchivesClient().Update(ctx, "myResourceGroup", "myRegistry", "myPackageType", "myArchiveName", armcontainerregistry.ArchiveUpdateParameters{
		Properties: &armcontainerregistry.ArchiveUpdateProperties{
			PublishedVersion: to.Ptr("string"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Archive = armcontainerregistry.Archive{
	// 	Name: to.Ptr("myArchiveName"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/packages/archives"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/packages/myPackageType/archives/myArchiveName"),
	// 	SystemData: &armcontainerregistry.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T23:41:38.720Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armcontainerregistry.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-06-09T23:41:38.720Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armcontainerregistry.LastModifiedByTypeUser),
	// 	},
	// 	Properties: &armcontainerregistry.ArchiveProperties{
	// 		PackageSource: &armcontainerregistry.ArchivePackageSourceProperties{
	// 			Type: to.Ptr(armcontainerregistry.PackageSourceTypeRemote),
	// 			URL: to.Ptr("string"),
	// 		},
	// 		ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		PublishedVersion: to.Ptr("string"),
	// 		RepositoryEndpoint: to.Ptr("string"),
	// 		RepositoryEndpointPrefix: to.Ptr("string"),
	// 	},
	// }
}
