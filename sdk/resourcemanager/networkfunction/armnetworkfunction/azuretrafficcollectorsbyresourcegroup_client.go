//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetworkfunction

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AzureTrafficCollectorsByResourceGroupClient contains the methods for the AzureTrafficCollectorsByResourceGroup group.
// Don't use this type directly, use NewAzureTrafficCollectorsByResourceGroupClient() instead.
type AzureTrafficCollectorsByResourceGroupClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAzureTrafficCollectorsByResourceGroupClient creates a new instance of AzureTrafficCollectorsByResourceGroupClient with the specified values.
//   - subscriptionID - Azure Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAzureTrafficCollectorsByResourceGroupClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AzureTrafficCollectorsByResourceGroupClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AzureTrafficCollectorsByResourceGroupClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Return list of Azure Traffic Collectors in a Resource Group
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The name of the resource group.
//   - options - AzureTrafficCollectorsByResourceGroupClientListOptions contains the optional parameters for the AzureTrafficCollectorsByResourceGroupClient.NewListPager
//     method.
func (client *AzureTrafficCollectorsByResourceGroupClient) NewListPager(resourceGroupName string, options *AzureTrafficCollectorsByResourceGroupClientListOptions) *runtime.Pager[AzureTrafficCollectorsByResourceGroupClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AzureTrafficCollectorsByResourceGroupClientListResponse]{
		More: func(page AzureTrafficCollectorsByResourceGroupClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AzureTrafficCollectorsByResourceGroupClientListResponse) (AzureTrafficCollectorsByResourceGroupClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AzureTrafficCollectorsByResourceGroupClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return AzureTrafficCollectorsByResourceGroupClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *AzureTrafficCollectorsByResourceGroupClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *AzureTrafficCollectorsByResourceGroupClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkFunction/azureTrafficCollectors"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AzureTrafficCollectorsByResourceGroupClient) listHandleResponse(resp *http.Response) (AzureTrafficCollectorsByResourceGroupClientListResponse, error) {
	result := AzureTrafficCollectorsByResourceGroupClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureTrafficCollectorListResult); err != nil {
		return AzureTrafficCollectorsByResourceGroupClientListResponse{}, err
	}
	return result, nil
}
