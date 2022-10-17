//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevops_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devops/armdevops"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devops/resource-manager/Microsoft.DevOps/preview/2019-07-01-preview/examples/CreateAzurePipeline-Sample-AspNet-WindowsWebApp.json
func ExamplePipelinesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevops.NewPipelinesClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myAspNetWebAppPipeline-rg",
		"myAspNetWebAppPipeline",
		armdevops.Pipeline{
			Location: to.Ptr("South India"),
			Tags:     map[string]*string{},
			Properties: &armdevops.PipelineProperties{
				BootstrapConfiguration: &armdevops.BootstrapConfiguration{
					Template: &armdevops.PipelineTemplate{
						ID: to.Ptr("ms.vss-continuous-delivery-pipeline-templates.aspnet-windowswebapp"),
						Parameters: map[string]*string{
							"appInsightLocation": to.Ptr("South India"),
							"appServicePlan":     to.Ptr("S1 Standard"),
							"azureAuth":          to.Ptr("{\"scheme\":\"ServicePrincipal\",\"parameters\":{\"tenantid\":\"{subscriptionTenantId}\",\"objectid\":\"{appObjectId}\",\"serviceprincipalid\":\"{appId}\",\"serviceprincipalkey\":\"{appSecret}\"}}"),
							"location":           to.Ptr("South India"),
							"resourceGroup":      to.Ptr("myAspNetWebAppPipeline-rg"),
							"subscriptionId":     to.Ptr("{subscriptionId}"),
							"webAppName":         to.Ptr("myAspNetWebApp"),
						},
					},
				},
				Organization: &armdevops.OrganizationReference{
					Name: to.Ptr("myAspNetWebAppPipeline-org"),
				},
				Project: &armdevops.ProjectReference{
					Name: to.Ptr("myAspNetWebAppPipeline-project"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devops/resource-manager/Microsoft.DevOps/preview/2019-07-01-preview/examples/GetAzurePipeline.json
func ExamplePipelinesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevops.NewPipelinesClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"myAspNetWebAppPipeline-rg",
		"myAspNetWebAppPipeline",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devops/resource-manager/Microsoft.DevOps/preview/2019-07-01-preview/examples/UpdateAzurePipeline.json
func ExamplePipelinesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevops.NewPipelinesClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"myAspNetWebAppPipeline-rg",
		"myAspNetWebAppPipeline",
		armdevops.PipelineUpdateParameters{
			Tags: map[string]*string{
				"tagKey": to.Ptr("tagvalue"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devops/resource-manager/Microsoft.DevOps/preview/2019-07-01-preview/examples/DeleteAzurePipeline.json
func ExamplePipelinesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevops.NewPipelinesClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"myAspNetWebAppPipeline-rg",
		"myAspNetWebAppPipeline",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devops/resource-manager/Microsoft.DevOps/preview/2019-07-01-preview/examples/ListAzurePipelinesByResourceGroup.json
func ExamplePipelinesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevops.NewPipelinesClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupPager("myAspNetWebAppPipeline-rg",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devops/resource-manager/Microsoft.DevOps/preview/2019-07-01-preview/examples/ListAzurePipelinesBySubscription.json
func ExamplePipelinesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevops.NewPipelinesClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListBySubscriptionPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}