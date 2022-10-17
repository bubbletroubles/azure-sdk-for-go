//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpeering

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ReceivedRoutesClient contains the methods for the ReceivedRoutes group.
// Don't use this type directly, use NewReceivedRoutesClient() instead.
type ReceivedRoutesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewReceivedRoutesClient creates a new instance of ReceivedRoutesClient with the specified values.
// subscriptionID - The Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewReceivedRoutesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReceivedRoutesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ReceivedRoutesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListByPeeringPager - Lists the prefixes received over the specified peering under the given subscription and resource
// group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// peeringName - The name of the peering.
// options - ReceivedRoutesClientListByPeeringOptions contains the optional parameters for the ReceivedRoutesClient.ListByPeering
// method.
func (client *ReceivedRoutesClient) NewListByPeeringPager(resourceGroupName string, peeringName string, options *ReceivedRoutesClientListByPeeringOptions) *runtime.Pager[ReceivedRoutesClientListByPeeringResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReceivedRoutesClientListByPeeringResponse]{
		More: func(page ReceivedRoutesClientListByPeeringResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReceivedRoutesClientListByPeeringResponse) (ReceivedRoutesClientListByPeeringResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByPeeringCreateRequest(ctx, resourceGroupName, peeringName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReceivedRoutesClientListByPeeringResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ReceivedRoutesClientListByPeeringResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReceivedRoutesClientListByPeeringResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByPeeringHandleResponse(resp)
		},
	})
}

// listByPeeringCreateRequest creates the ListByPeering request.
func (client *ReceivedRoutesClient) listByPeeringCreateRequest(ctx context.Context, resourceGroupName string, peeringName string, options *ReceivedRoutesClientListByPeeringOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/receivedRoutes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Prefix != nil {
		reqQP.Set("prefix", *options.Prefix)
	}
	if options != nil && options.AsPath != nil {
		reqQP.Set("asPath", *options.AsPath)
	}
	if options != nil && options.OriginAsValidationState != nil {
		reqQP.Set("originAsValidationState", *options.OriginAsValidationState)
	}
	if options != nil && options.RpkiValidationState != nil {
		reqQP.Set("rpkiValidationState", *options.RpkiValidationState)
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByPeeringHandleResponse handles the ListByPeering response.
func (client *ReceivedRoutesClient) listByPeeringHandleResponse(resp *http.Response) (ReceivedRoutesClientListByPeeringResponse, error) {
	result := ReceivedRoutesClientListByPeeringResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReceivedRouteListResult); err != nil {
		return ReceivedRoutesClientListByPeeringResponse{}, err
	}
	return result, nil
}