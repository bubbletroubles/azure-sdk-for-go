//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridnetwork

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

// NetworkServiceDesignVersionsClient contains the methods for the NetworkServiceDesignVersions group.
// Don't use this type directly, use NewNetworkServiceDesignVersionsClient() instead.
type NetworkServiceDesignVersionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNetworkServiceDesignVersionsClient creates a new instance of NetworkServiceDesignVersionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNetworkServiceDesignVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NetworkServiceDesignVersionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NetworkServiceDesignVersionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a network service design version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - networkServiceDesignGroupName - The name of the network service design group.
//   - networkServiceDesignVersionName - The name of the network service design version. The name should conform to the SemVer
//     2.0.0 specification: https://semver.org/spec/v2.0.0.html.
//   - parameters - Parameters supplied to the create or update network service design version operation.
//   - options - NetworkServiceDesignVersionsClientBeginCreateOrUpdateOptions contains the optional parameters for the NetworkServiceDesignVersionsClient.BeginCreateOrUpdate
//     method.
func (client *NetworkServiceDesignVersionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, networkServiceDesignVersionName string, parameters NetworkServiceDesignVersion, options *NetworkServiceDesignVersionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[NetworkServiceDesignVersionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, publisherName, networkServiceDesignGroupName, networkServiceDesignVersionName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkServiceDesignVersionsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkServiceDesignVersionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates a network service design version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *NetworkServiceDesignVersionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, networkServiceDesignVersionName string, parameters NetworkServiceDesignVersion, options *NetworkServiceDesignVersionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkServiceDesignVersionsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, publisherName, networkServiceDesignGroupName, networkServiceDesignVersionName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *NetworkServiceDesignVersionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, networkServiceDesignVersionName string, parameters NetworkServiceDesignVersion, options *NetworkServiceDesignVersionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/networkServiceDesignGroups/{networkServiceDesignGroupName}/networkServiceDesignVersions/{networkServiceDesignVersionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if networkServiceDesignGroupName == "" {
		return nil, errors.New("parameter networkServiceDesignGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkServiceDesignGroupName}", url.PathEscape(networkServiceDesignGroupName))
	if networkServiceDesignVersionName == "" {
		return nil, errors.New("parameter networkServiceDesignVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkServiceDesignVersionName}", url.PathEscape(networkServiceDesignVersionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the specified network service design version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - networkServiceDesignGroupName - The name of the network service design group.
//   - networkServiceDesignVersionName - The name of the network service design version. The name should conform to the SemVer
//     2.0.0 specification: https://semver.org/spec/v2.0.0.html.
//   - options - NetworkServiceDesignVersionsClientBeginDeleteOptions contains the optional parameters for the NetworkServiceDesignVersionsClient.BeginDelete
//     method.
func (client *NetworkServiceDesignVersionsClient) BeginDelete(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, networkServiceDesignVersionName string, options *NetworkServiceDesignVersionsClientBeginDeleteOptions) (*runtime.Poller[NetworkServiceDesignVersionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, publisherName, networkServiceDesignGroupName, networkServiceDesignVersionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkServiceDesignVersionsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkServiceDesignVersionsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the specified network service design version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *NetworkServiceDesignVersionsClient) deleteOperation(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, networkServiceDesignVersionName string, options *NetworkServiceDesignVersionsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkServiceDesignVersionsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, publisherName, networkServiceDesignGroupName, networkServiceDesignVersionName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *NetworkServiceDesignVersionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, networkServiceDesignVersionName string, options *NetworkServiceDesignVersionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/networkServiceDesignGroups/{networkServiceDesignGroupName}/networkServiceDesignVersions/{networkServiceDesignVersionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if networkServiceDesignGroupName == "" {
		return nil, errors.New("parameter networkServiceDesignGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkServiceDesignGroupName}", url.PathEscape(networkServiceDesignGroupName))
	if networkServiceDesignVersionName == "" {
		return nil, errors.New("parameter networkServiceDesignVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkServiceDesignVersionName}", url.PathEscape(networkServiceDesignVersionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets information about a network service design version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - networkServiceDesignGroupName - The name of the network service design group.
//   - networkServiceDesignVersionName - The name of the network service design version. The name should conform to the SemVer
//     2.0.0 specification: https://semver.org/spec/v2.0.0.html.
//   - options - NetworkServiceDesignVersionsClientGetOptions contains the optional parameters for the NetworkServiceDesignVersionsClient.Get
//     method.
func (client *NetworkServiceDesignVersionsClient) Get(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, networkServiceDesignVersionName string, options *NetworkServiceDesignVersionsClientGetOptions) (NetworkServiceDesignVersionsClientGetResponse, error) {
	var err error
	const operationName = "NetworkServiceDesignVersionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, publisherName, networkServiceDesignGroupName, networkServiceDesignVersionName, options)
	if err != nil {
		return NetworkServiceDesignVersionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NetworkServiceDesignVersionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NetworkServiceDesignVersionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *NetworkServiceDesignVersionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, networkServiceDesignVersionName string, options *NetworkServiceDesignVersionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/networkServiceDesignGroups/{networkServiceDesignGroupName}/networkServiceDesignVersions/{networkServiceDesignVersionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if networkServiceDesignGroupName == "" {
		return nil, errors.New("parameter networkServiceDesignGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkServiceDesignGroupName}", url.PathEscape(networkServiceDesignGroupName))
	if networkServiceDesignVersionName == "" {
		return nil, errors.New("parameter networkServiceDesignVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkServiceDesignVersionName}", url.PathEscape(networkServiceDesignVersionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NetworkServiceDesignVersionsClient) getHandleResponse(resp *http.Response) (NetworkServiceDesignVersionsClientGetResponse, error) {
	result := NetworkServiceDesignVersionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkServiceDesignVersion); err != nil {
		return NetworkServiceDesignVersionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByNetworkServiceDesignGroupPager - Gets information about a list of network service design versions under a network
// service design group.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - networkServiceDesignGroupName - The name of the network service design group.
//   - options - NetworkServiceDesignVersionsClientListByNetworkServiceDesignGroupOptions contains the optional parameters for
//     the NetworkServiceDesignVersionsClient.NewListByNetworkServiceDesignGroupPager method.
func (client *NetworkServiceDesignVersionsClient) NewListByNetworkServiceDesignGroupPager(resourceGroupName string, publisherName string, networkServiceDesignGroupName string, options *NetworkServiceDesignVersionsClientListByNetworkServiceDesignGroupOptions) *runtime.Pager[NetworkServiceDesignVersionsClientListByNetworkServiceDesignGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkServiceDesignVersionsClientListByNetworkServiceDesignGroupResponse]{
		More: func(page NetworkServiceDesignVersionsClientListByNetworkServiceDesignGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkServiceDesignVersionsClientListByNetworkServiceDesignGroupResponse) (NetworkServiceDesignVersionsClientListByNetworkServiceDesignGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "NetworkServiceDesignVersionsClient.NewListByNetworkServiceDesignGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByNetworkServiceDesignGroupCreateRequest(ctx, resourceGroupName, publisherName, networkServiceDesignGroupName, options)
			}, nil)
			if err != nil {
				return NetworkServiceDesignVersionsClientListByNetworkServiceDesignGroupResponse{}, err
			}
			return client.listByNetworkServiceDesignGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByNetworkServiceDesignGroupCreateRequest creates the ListByNetworkServiceDesignGroup request.
func (client *NetworkServiceDesignVersionsClient) listByNetworkServiceDesignGroupCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, options *NetworkServiceDesignVersionsClientListByNetworkServiceDesignGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/networkServiceDesignGroups/{networkServiceDesignGroupName}/networkServiceDesignVersions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if networkServiceDesignGroupName == "" {
		return nil, errors.New("parameter networkServiceDesignGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkServiceDesignGroupName}", url.PathEscape(networkServiceDesignGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByNetworkServiceDesignGroupHandleResponse handles the ListByNetworkServiceDesignGroup response.
func (client *NetworkServiceDesignVersionsClient) listByNetworkServiceDesignGroupHandleResponse(resp *http.Response) (NetworkServiceDesignVersionsClientListByNetworkServiceDesignGroupResponse, error) {
	result := NetworkServiceDesignVersionsClientListByNetworkServiceDesignGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkServiceDesignVersionListResult); err != nil {
		return NetworkServiceDesignVersionsClientListByNetworkServiceDesignGroupResponse{}, err
	}
	return result, nil
}

// Update - Updates a network service design version resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - networkServiceDesignGroupName - The name of the network service design group.
//   - networkServiceDesignVersionName - The name of the network service design version. The name should conform to the SemVer
//     2.0.0 specification: https://semver.org/spec/v2.0.0.html.
//   - parameters - Parameters supplied to the create or update network service design version operation.
//   - options - NetworkServiceDesignVersionsClientUpdateOptions contains the optional parameters for the NetworkServiceDesignVersionsClient.Update
//     method.
func (client *NetworkServiceDesignVersionsClient) Update(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, networkServiceDesignVersionName string, parameters TagsObject, options *NetworkServiceDesignVersionsClientUpdateOptions) (NetworkServiceDesignVersionsClientUpdateResponse, error) {
	var err error
	const operationName = "NetworkServiceDesignVersionsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, publisherName, networkServiceDesignGroupName, networkServiceDesignVersionName, parameters, options)
	if err != nil {
		return NetworkServiceDesignVersionsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NetworkServiceDesignVersionsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NetworkServiceDesignVersionsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *NetworkServiceDesignVersionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, networkServiceDesignVersionName string, parameters TagsObject, options *NetworkServiceDesignVersionsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/networkServiceDesignGroups/{networkServiceDesignGroupName}/networkServiceDesignVersions/{networkServiceDesignVersionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if networkServiceDesignGroupName == "" {
		return nil, errors.New("parameter networkServiceDesignGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkServiceDesignGroupName}", url.PathEscape(networkServiceDesignGroupName))
	if networkServiceDesignVersionName == "" {
		return nil, errors.New("parameter networkServiceDesignVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkServiceDesignVersionName}", url.PathEscape(networkServiceDesignVersionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *NetworkServiceDesignVersionsClient) updateHandleResponse(resp *http.Response) (NetworkServiceDesignVersionsClientUpdateResponse, error) {
	result := NetworkServiceDesignVersionsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkServiceDesignVersion); err != nil {
		return NetworkServiceDesignVersionsClientUpdateResponse{}, err
	}
	return result, nil
}

// BeginUpdateState - Update network service design version state.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - publisherName - The name of the publisher.
//   - networkServiceDesignGroupName - The name of the network service design group.
//   - networkServiceDesignVersionName - The name of the network service design version. The name should conform to the SemVer
//     2.0.0 specification: https://semver.org/spec/v2.0.0.html.
//   - parameters - Parameters supplied to update the state of network service design version.
//   - options - NetworkServiceDesignVersionsClientBeginUpdateStateOptions contains the optional parameters for the NetworkServiceDesignVersionsClient.BeginUpdateState
//     method.
func (client *NetworkServiceDesignVersionsClient) BeginUpdateState(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, networkServiceDesignVersionName string, parameters NetworkServiceDesignVersionUpdateState, options *NetworkServiceDesignVersionsClientBeginUpdateStateOptions) (*runtime.Poller[NetworkServiceDesignVersionsClientUpdateStateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateState(ctx, resourceGroupName, publisherName, networkServiceDesignGroupName, networkServiceDesignVersionName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkServiceDesignVersionsClientUpdateStateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[NetworkServiceDesignVersionsClientUpdateStateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// UpdateState - Update network service design version state.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
func (client *NetworkServiceDesignVersionsClient) updateState(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, networkServiceDesignVersionName string, parameters NetworkServiceDesignVersionUpdateState, options *NetworkServiceDesignVersionsClientBeginUpdateStateOptions) (*http.Response, error) {
	var err error
	const operationName = "NetworkServiceDesignVersionsClient.BeginUpdateState"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateStateCreateRequest(ctx, resourceGroupName, publisherName, networkServiceDesignGroupName, networkServiceDesignVersionName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateStateCreateRequest creates the UpdateState request.
func (client *NetworkServiceDesignVersionsClient) updateStateCreateRequest(ctx context.Context, resourceGroupName string, publisherName string, networkServiceDesignGroupName string, networkServiceDesignVersionName string, parameters NetworkServiceDesignVersionUpdateState, options *NetworkServiceDesignVersionsClientBeginUpdateStateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridNetwork/publishers/{publisherName}/networkServiceDesignGroups/{networkServiceDesignGroupName}/networkServiceDesignVersions/{networkServiceDesignVersionName}/updateState"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if networkServiceDesignGroupName == "" {
		return nil, errors.New("parameter networkServiceDesignGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkServiceDesignGroupName}", url.PathEscape(networkServiceDesignGroupName))
	if networkServiceDesignVersionName == "" {
		return nil, errors.New("parameter networkServiceDesignVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkServiceDesignVersionName}", url.PathEscape(networkServiceDesignVersionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
