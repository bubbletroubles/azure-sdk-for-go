//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armquantum

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

// OfferingsClient contains the methods for the Offerings group.
// Don't use this type directly, use NewOfferingsClient() instead.
type OfferingsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewOfferingsClient creates a new instance of OfferingsClient with the specified values.
// subscriptionID - The Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOfferingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OfferingsClient, error) {
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
	client := &OfferingsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - Returns the list of all provider offerings available for the given location.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-10-preview
// locationName - Location.
// options - OfferingsClientListOptions contains the optional parameters for the OfferingsClient.List method.
func (client *OfferingsClient) NewListPager(locationName string, options *OfferingsClientListOptions) *runtime.Pager[OfferingsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[OfferingsClientListResponse]{
		More: func(page OfferingsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OfferingsClientListResponse) (OfferingsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, locationName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OfferingsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return OfferingsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OfferingsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *OfferingsClient) listCreateRequest(ctx context.Context, locationName string, options *OfferingsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Quantum/locations/{locationName}/offerings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-10-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OfferingsClient) listHandleResponse(resp *http.Response) (OfferingsClientListResponse, error) {
	result := OfferingsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OfferingsListResult); err != nil {
		return OfferingsClientListResponse{}, err
	}
	return result, nil
}