//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/738ab25fe76324897f273645906d4a0415068a6c/specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/WorkflowRunActionRepetitionsRequestHistories_List.json
func ExampleWorkflowRunActionRepetitionsRequestHistoriesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkflowRunActionRepetitionsRequestHistoriesClient().NewListPager("test-resource-group", "test-name", "test-workflow", "08586776228332053161046300351", "HTTP_Webhook", "000001", nil)
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
		// page.RequestHistoryListResult = armappservice.RequestHistoryListResult{
		// 	Value: []*armappservice.RequestHistory{
		// 		{
		// 			Name: to.Ptr("08586611142732800686"),
		// 			Type: to.Ptr("/workflows/runs/actions/requestHistories"),
		// 			ID: to.Ptr("/workflows/test-workflow/runs/08586611142736787787412824395CU21/actions/HTTP_Webhook/requestHistories/08586611142732800686"),
		// 			Properties: &armappservice.RequestHistoryProperties{
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-25T18:36:52.186Z"); return t}()),
		// 				Response: &armappservice.Response{
		// 					BodyLink: &armappservice.ContentLink{
		// 						ContentHash: &armappservice.ContentHash{
		// 							Algorithm: to.Ptr("md5"),
		// 							Value: to.Ptr("2LOOAR8Eh2pd7AvRHXUhRg=="),
		// 						},
		// 						ContentSize: to.Ptr[int64](137),
		// 						ContentVersion: to.Ptr("2LOOAR8Eh2pd7AvRHXUhRg=="),
		// 						URI: to.Ptr("https://tempuri.org"),
		// 					},
		// 					Headers: map[string]any{
		// 						"Cache-Control": "private",
		// 						"Date": "Thu, 25 Oct 2018 18:36:51 GMT",
		// 						"Location": "http://www.bing.com/",
		// 						"Server": "Microsoft-IIS/10.0",
		// 						"X-AspNet-Version": "4.0.30319",
		// 						"X-Powered-By": "ASP.NET",
		// 					},
		// 					StatusCode: to.Ptr[int32](302),
		// 				},
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-25T18:36:51.920Z"); return t}()),
		// 				Request: &armappservice.Request{
		// 					Method: to.Ptr("GET"),
		// 					Headers: map[string]any{
		// 						"Accept-Language": "en-US",
		// 						"User-Agent": "azure-logic-apps/1.0,(workflow 80244732be3648f59d2084fd979cdd56; version 08586611142904036539)",
		// 						"x-ms-action-tracking-id": "ad27f634-6523-492f-924e-9a75e28619c8",
		// 						"x-ms-client-request-id": "ad484925-4148-4dd0-9488-07aed418b256",
		// 						"x-ms-client-tracking-id": "08586611142736787787412824395CU21",
		// 						"x-ms-correlation-id": "ad484925-4148-4dd0-9488-07aed418b256",
		// 						"x-ms-execution-location": "brazilsouth",
		// 						"x-ms-tracking-id": "ad484925-4148-4dd0-9488-07aed418b256",
		// 						"x-ms-workflow-id": "80244732be3648f59d2084fd979cdd56",
		// 						"x-ms-workflow-name": "test-workflow",
		// 						"x-ms-workflow-operation-name": "HTTP_Webhook",
		// 						"x-ms-workflow-resourcegroup-name": "test-resource-group",
		// 						"x-ms-workflow-run-id": "08586611142736787787412824395CU21",
		// 						"x-ms-workflow-run-tracking-id": "b4cd2e77-f949-4d8c-8753-791407aebde8",
		// 						"x-ms-workflow-subscription-capacity": "Large",
		// 						"x-ms-workflow-subscription-id": "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345",
		// 						"x-ms-workflow-system-id": "/locations/brazilsouth/scaleunits/prod-17/workflows/80244732be3648f59d2084fd979cdd56",
		// 						"x-ms-workflow-version": "08586611142904036539",
		// 					},
		// 					URI: to.Ptr("http://tempuri.org"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/738ab25fe76324897f273645906d4a0415068a6c/specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/WorkflowRunActionRepetitionsRequestHistories_Get.json
func ExampleWorkflowRunActionRepetitionsRequestHistoriesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkflowRunActionRepetitionsRequestHistoriesClient().Get(ctx, "test-resource-group", "test-name", "test-workflow", "08586776228332053161046300351", "HTTP_Webhook", "000001", "08586611142732800686", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RequestHistory = armappservice.RequestHistory{
	// 	Name: to.Ptr("08586611142732800686"),
	// 	Type: to.Ptr("/workflows/runs/actions/requestHistories"),
	// 	ID: to.Ptr("/workflows/test-workflow/runs/08586611142736787787412824395CU21/actions/HTTP_Webhook/requestHistories/08586611142732800686"),
	// 	Properties: &armappservice.RequestHistoryProperties{
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-25T18:36:52.186Z"); return t}()),
	// 		Response: &armappservice.Response{
	// 			BodyLink: &armappservice.ContentLink{
	// 				ContentHash: &armappservice.ContentHash{
	// 					Algorithm: to.Ptr("md5"),
	// 					Value: to.Ptr("2LOOAR8Eh2pd7AvRHXUhRg=="),
	// 				},
	// 				ContentSize: to.Ptr[int64](137),
	// 				ContentVersion: to.Ptr("2LOOAR8Eh2pd7AvRHXUhRg=="),
	// 				URI: to.Ptr("https://tempuri.org"),
	// 			},
	// 			Headers: map[string]any{
	// 				"Cache-Control": "private",
	// 				"Date": "Thu, 25 Oct 2018 18:36:51 GMT",
	// 				"Location": "http://www.bing.com/",
	// 				"Server": "Microsoft-IIS/10.0",
	// 				"X-AspNet-Version": "4.0.30319",
	// 				"X-Powered-By": "ASP.NET",
	// 			},
	// 			StatusCode: to.Ptr[int32](302),
	// 		},
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-10-25T18:36:51.920Z"); return t}()),
	// 		Request: &armappservice.Request{
	// 			Method: to.Ptr("GET"),
	// 			Headers: map[string]any{
	// 				"Accept-Language": "en-US",
	// 				"User-Agent": "azure-logic-apps/1.0,(workflow 80244732be3648f59d2084fd979cdd56; version 08586611142904036539)",
	// 				"x-ms-action-tracking-id": "ad27f634-6523-492f-924e-9a75e28619c8",
	// 				"x-ms-client-request-id": "ad484925-4148-4dd0-9488-07aed418b256",
	// 				"x-ms-client-tracking-id": "08586611142736787787412824395CU21",
	// 				"x-ms-correlation-id": "ad484925-4148-4dd0-9488-07aed418b256",
	// 				"x-ms-execution-location": "brazilsouth",
	// 				"x-ms-tracking-id": "ad484925-4148-4dd0-9488-07aed418b256",
	// 				"x-ms-workflow-id": "80244732be3648f59d2084fd979cdd56",
	// 				"x-ms-workflow-name": "test-workflow",
	// 				"x-ms-workflow-operation-name": "HTTP_Webhook",
	// 				"x-ms-workflow-resourcegroup-name": "test-resource-group",
	// 				"x-ms-workflow-run-id": "08586611142736787787412824395CU21",
	// 				"x-ms-workflow-run-tracking-id": "b4cd2e77-f949-4d8c-8753-791407aebde8",
	// 				"x-ms-workflow-subscription-capacity": "Large",
	// 				"x-ms-workflow-subscription-id": "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345",
	// 				"x-ms-workflow-system-id": "/locations/brazilsouth/scaleunits/prod-17/workflows/80244732be3648f59d2084fd979cdd56",
	// 				"x-ms-workflow-version": "08586611142904036539",
	// 			},
	// 			URI: to.Ptr("http://tempuri.org"),
	// 		},
	// 	},
	// }
}
