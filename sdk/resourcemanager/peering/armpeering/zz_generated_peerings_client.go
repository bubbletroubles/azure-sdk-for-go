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

// PeeringsClient contains the methods for the Peerings group.
// Don't use this type directly, use NewPeeringsClient() instead.
type PeeringsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPeeringsClient creates a new instance of PeeringsClient with the specified values.
// subscriptionID - The Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPeeringsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PeeringsClient, error) {
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
	client := &PeeringsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates a new peering or updates an existing peering with the specified name under the given subscription
// and resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// peeringName - The name of the peering.
// peering - The properties needed to create or update a peering.
// options - PeeringsClientCreateOrUpdateOptions contains the optional parameters for the PeeringsClient.CreateOrUpdate method.
func (client *PeeringsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, peeringName string, peering Peering, options *PeeringsClientCreateOrUpdateOptions) (PeeringsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, peeringName, peering, options)
	if err != nil {
		return PeeringsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PeeringsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return PeeringsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PeeringsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, peeringName string, peering Peering, options *PeeringsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, peering)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PeeringsClient) createOrUpdateHandleResponse(resp *http.Response) (PeeringsClientCreateOrUpdateResponse, error) {
	result := PeeringsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Peering); err != nil {
		return PeeringsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an existing peering with the specified name under the given subscription and resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// peeringName - The name of the peering.
// options - PeeringsClientDeleteOptions contains the optional parameters for the PeeringsClient.Delete method.
func (client *PeeringsClient) Delete(ctx context.Context, resourceGroupName string, peeringName string, options *PeeringsClientDeleteOptions) (PeeringsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, peeringName, options)
	if err != nil {
		return PeeringsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PeeringsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return PeeringsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return PeeringsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PeeringsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, peeringName string, options *PeeringsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an existing peering with the specified name under the given subscription and resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// peeringName - The name of the peering.
// options - PeeringsClientGetOptions contains the optional parameters for the PeeringsClient.Get method.
func (client *PeeringsClient) Get(ctx context.Context, resourceGroupName string, peeringName string, options *PeeringsClientGetOptions) (PeeringsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, peeringName, options)
	if err != nil {
		return PeeringsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PeeringsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PeeringsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PeeringsClient) getCreateRequest(ctx context.Context, resourceGroupName string, peeringName string, options *PeeringsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}"
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
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PeeringsClient) getHandleResponse(resp *http.Response) (PeeringsClientGetResponse, error) {
	result := PeeringsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Peering); err != nil {
		return PeeringsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all of the peerings under the given subscription and resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// options - PeeringsClientListByResourceGroupOptions contains the optional parameters for the PeeringsClient.ListByResourceGroup
// method.
func (client *PeeringsClient) NewListByResourceGroupPager(resourceGroupName string, options *PeeringsClientListByResourceGroupOptions) *runtime.Pager[PeeringsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[PeeringsClientListByResourceGroupResponse]{
		More: func(page PeeringsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PeeringsClientListByResourceGroupResponse) (PeeringsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PeeringsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PeeringsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PeeringsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *PeeringsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *PeeringsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *PeeringsClient) listByResourceGroupHandleResponse(resp *http.Response) (PeeringsClientListByResourceGroupResponse, error) {
	result := PeeringsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return PeeringsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists all of the peerings under the given subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// options - PeeringsClientListBySubscriptionOptions contains the optional parameters for the PeeringsClient.ListBySubscription
// method.
func (client *PeeringsClient) NewListBySubscriptionPager(options *PeeringsClientListBySubscriptionOptions) *runtime.Pager[PeeringsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[PeeringsClientListBySubscriptionResponse]{
		More: func(page PeeringsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PeeringsClientListBySubscriptionResponse) (PeeringsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PeeringsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PeeringsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PeeringsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *PeeringsClient) listBySubscriptionCreateRequest(ctx context.Context, options *PeeringsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peerings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *PeeringsClient) listBySubscriptionHandleResponse(resp *http.Response) (PeeringsClientListBySubscriptionResponse, error) {
	result := PeeringsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResult); err != nil {
		return PeeringsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Updates tags for a peering with the specified name under the given subscription and resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01
// resourceGroupName - The name of the resource group.
// peeringName - The name of the peering.
// tags - The resource tags.
// options - PeeringsClientUpdateOptions contains the optional parameters for the PeeringsClient.Update method.
func (client *PeeringsClient) Update(ctx context.Context, resourceGroupName string, peeringName string, tags ResourceTags, options *PeeringsClientUpdateOptions) (PeeringsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, peeringName, tags, options)
	if err != nil {
		return PeeringsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PeeringsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PeeringsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *PeeringsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, peeringName string, tags ResourceTags, options *PeeringsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, tags)
}

// updateHandleResponse handles the Update response.
func (client *PeeringsClient) updateHandleResponse(resp *http.Response) (PeeringsClientUpdateResponse, error) {
	result := PeeringsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Peering); err != nil {
		return PeeringsClientUpdateResponse{}, err
	}
	return result, nil
}