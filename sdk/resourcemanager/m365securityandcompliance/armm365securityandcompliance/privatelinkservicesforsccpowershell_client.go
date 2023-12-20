//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armm365securityandcompliance

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

// PrivateLinkServicesForSCCPowershellClient contains the methods for the PrivateLinkServicesForSCCPowershell group.
// Don't use this type directly, use NewPrivateLinkServicesForSCCPowershellClient() instead.
type PrivateLinkServicesForSCCPowershellClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPrivateLinkServicesForSCCPowershellClient creates a new instance of PrivateLinkServicesForSCCPowershellClient with the specified values.
//   - subscriptionID - The subscription identifier.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateLinkServicesForSCCPowershellClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateLinkServicesForSCCPowershellClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateLinkServicesForSCCPowershellClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update the metadata of a privateLinkServicesForSCCPowershell instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-25-preview
//   - resourceGroupName - The name of the resource group that contains the service instance.
//   - resourceName - The name of the service instance.
//   - privateLinkServicesForSCCPowershellDescription - The service instance metadata.
//   - options - PrivateLinkServicesForSCCPowershellClientBeginCreateOrUpdateOptions contains the optional parameters for the
//     PrivateLinkServicesForSCCPowershellClient.BeginCreateOrUpdate method.
func (client *PrivateLinkServicesForSCCPowershellClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, privateLinkServicesForSCCPowershellDescription PrivateLinkServicesForSCCPowershellDescription, options *PrivateLinkServicesForSCCPowershellClientBeginCreateOrUpdateOptions) (*runtime.Poller[PrivateLinkServicesForSCCPowershellClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceName, privateLinkServicesForSCCPowershellDescription, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PrivateLinkServicesForSCCPowershellClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[PrivateLinkServicesForSCCPowershellClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or update the metadata of a privateLinkServicesForSCCPowershell instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-25-preview
func (client *PrivateLinkServicesForSCCPowershellClient) createOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, privateLinkServicesForSCCPowershellDescription PrivateLinkServicesForSCCPowershellDescription, options *PrivateLinkServicesForSCCPowershellClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "PrivateLinkServicesForSCCPowershellClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, privateLinkServicesForSCCPowershellDescription, options)
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
func (client *PrivateLinkServicesForSCCPowershellClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, privateLinkServicesForSCCPowershellDescription PrivateLinkServicesForSCCPowershellDescription, options *PrivateLinkServicesForSCCPowershellClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForSCCPowershell/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, privateLinkServicesForSCCPowershellDescription); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a service instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-25-preview
//   - resourceGroupName - The name of the resource group that contains the service instance.
//   - resourceName - The name of the service instance.
//   - options - PrivateLinkServicesForSCCPowershellClientBeginDeleteOptions contains the optional parameters for the PrivateLinkServicesForSCCPowershellClient.BeginDelete
//     method.
func (client *PrivateLinkServicesForSCCPowershellClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateLinkServicesForSCCPowershellClientBeginDeleteOptions) (*runtime.Poller[PrivateLinkServicesForSCCPowershellClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PrivateLinkServicesForSCCPowershellClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[PrivateLinkServicesForSCCPowershellClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a service instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-25-preview
func (client *PrivateLinkServicesForSCCPowershellClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateLinkServicesForSCCPowershellClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "PrivateLinkServicesForSCCPowershellClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PrivateLinkServicesForSCCPowershellClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateLinkServicesForSCCPowershellClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForSCCPowershell/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the metadata of a privateLinkServicesForSCCPowershell resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-25-preview
//   - resourceGroupName - The name of the resource group that contains the service instance.
//   - resourceName - The name of the service instance.
//   - options - PrivateLinkServicesForSCCPowershellClientGetOptions contains the optional parameters for the PrivateLinkServicesForSCCPowershellClient.Get
//     method.
func (client *PrivateLinkServicesForSCCPowershellClient) Get(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateLinkServicesForSCCPowershellClientGetOptions) (PrivateLinkServicesForSCCPowershellClientGetResponse, error) {
	var err error
	const operationName = "PrivateLinkServicesForSCCPowershellClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return PrivateLinkServicesForSCCPowershellClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateLinkServicesForSCCPowershellClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateLinkServicesForSCCPowershellClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PrivateLinkServicesForSCCPowershellClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *PrivateLinkServicesForSCCPowershellClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForSCCPowershell/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateLinkServicesForSCCPowershellClient) getHandleResponse(resp *http.Response) (PrivateLinkServicesForSCCPowershellClientGetResponse, error) {
	result := PrivateLinkServicesForSCCPowershellClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkServicesForSCCPowershellDescription); err != nil {
		return PrivateLinkServicesForSCCPowershellClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get all the privateLinkServicesForSCCPowershell instances in a subscription.
//
// Generated from API version 2021-03-25-preview
//   - options - PrivateLinkServicesForSCCPowershellClientListOptions contains the optional parameters for the PrivateLinkServicesForSCCPowershellClient.NewListPager
//     method.
func (client *PrivateLinkServicesForSCCPowershellClient) NewListPager(options *PrivateLinkServicesForSCCPowershellClientListOptions) *runtime.Pager[PrivateLinkServicesForSCCPowershellClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateLinkServicesForSCCPowershellClientListResponse]{
		More: func(page PrivateLinkServicesForSCCPowershellClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateLinkServicesForSCCPowershellClientListResponse) (PrivateLinkServicesForSCCPowershellClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PrivateLinkServicesForSCCPowershellClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return PrivateLinkServicesForSCCPowershellClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *PrivateLinkServicesForSCCPowershellClient) listCreateRequest(ctx context.Context, options *PrivateLinkServicesForSCCPowershellClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForSCCPowershell"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PrivateLinkServicesForSCCPowershellClient) listHandleResponse(resp *http.Response) (PrivateLinkServicesForSCCPowershellClientListResponse, error) {
	result := PrivateLinkServicesForSCCPowershellClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkServicesForSCCPowershellDescriptionListResult); err != nil {
		return PrivateLinkServicesForSCCPowershellClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Get all the service instances in a resource group.
//
// Generated from API version 2021-03-25-preview
//   - resourceGroupName - The name of the resource group that contains the service instance.
//   - options - PrivateLinkServicesForSCCPowershellClientListByResourceGroupOptions contains the optional parameters for the
//     PrivateLinkServicesForSCCPowershellClient.NewListByResourceGroupPager method.
func (client *PrivateLinkServicesForSCCPowershellClient) NewListByResourceGroupPager(resourceGroupName string, options *PrivateLinkServicesForSCCPowershellClientListByResourceGroupOptions) *runtime.Pager[PrivateLinkServicesForSCCPowershellClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateLinkServicesForSCCPowershellClientListByResourceGroupResponse]{
		More: func(page PrivateLinkServicesForSCCPowershellClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateLinkServicesForSCCPowershellClientListByResourceGroupResponse) (PrivateLinkServicesForSCCPowershellClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PrivateLinkServicesForSCCPowershellClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return PrivateLinkServicesForSCCPowershellClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *PrivateLinkServicesForSCCPowershellClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *PrivateLinkServicesForSCCPowershellClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForSCCPowershell"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *PrivateLinkServicesForSCCPowershellClient) listByResourceGroupHandleResponse(resp *http.Response) (PrivateLinkServicesForSCCPowershellClientListByResourceGroupResponse, error) {
	result := PrivateLinkServicesForSCCPowershellClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkServicesForSCCPowershellDescriptionListResult); err != nil {
		return PrivateLinkServicesForSCCPowershellClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update the metadata of a privateLinkServicesForSCCPowershell instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-25-preview
//   - resourceGroupName - The name of the resource group that contains the service instance.
//   - resourceName - The name of the service instance.
//   - servicePatchDescription - The service instance metadata and security metadata.
//   - options - PrivateLinkServicesForSCCPowershellClientBeginUpdateOptions contains the optional parameters for the PrivateLinkServicesForSCCPowershellClient.BeginUpdate
//     method.
func (client *PrivateLinkServicesForSCCPowershellClient) BeginUpdate(ctx context.Context, resourceGroupName string, resourceName string, servicePatchDescription ServicesPatchDescription, options *PrivateLinkServicesForSCCPowershellClientBeginUpdateOptions) (*runtime.Poller[PrivateLinkServicesForSCCPowershellClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, resourceName, servicePatchDescription, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PrivateLinkServicesForSCCPowershellClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[PrivateLinkServicesForSCCPowershellClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update the metadata of a privateLinkServicesForSCCPowershell instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-03-25-preview
func (client *PrivateLinkServicesForSCCPowershellClient) update(ctx context.Context, resourceGroupName string, resourceName string, servicePatchDescription ServicesPatchDescription, options *PrivateLinkServicesForSCCPowershellClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "PrivateLinkServicesForSCCPowershellClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, servicePatchDescription, options)
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

// updateCreateRequest creates the Update request.
func (client *PrivateLinkServicesForSCCPowershellClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, servicePatchDescription ServicesPatchDescription, options *PrivateLinkServicesForSCCPowershellClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.M365SecurityAndCompliance/privateLinkServicesForSCCPowershell/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, servicePatchDescription); err != nil {
		return nil, err
	}
	return req, nil
}
