//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevtestlabs

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// UsersClient contains the methods for the Users group.
// Don't use this type directly, use NewUsersClient() instead.
type UsersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewUsersClient creates a new instance of UsersClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewUsersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*UsersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &UsersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or replace an existing user profile. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - name - The name of the user profile.
//   - userParam - Profile of a lab user.
//   - options - UsersClientBeginCreateOrUpdateOptions contains the optional parameters for the UsersClient.BeginCreateOrUpdate
//     method.
func (client *UsersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, name string, userParam User, options *UsersClientBeginCreateOrUpdateOptions) (*runtime.Poller[UsersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, labName, name, userParam, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[UsersClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[UsersClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or replace an existing user profile. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
func (client *UsersClient) createOrUpdate(ctx context.Context, resourceGroupName string, labName string, name string, userParam User, options *UsersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "UsersClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, labName, name, userParam, options)
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
func (client *UsersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, userParam User, options *UsersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, userParam); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete user profile. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - name - The name of the user profile.
//   - options - UsersClientBeginDeleteOptions contains the optional parameters for the UsersClient.BeginDelete method.
func (client *UsersClient) BeginDelete(ctx context.Context, resourceGroupName string, labName string, name string, options *UsersClientBeginDeleteOptions) (*runtime.Poller[UsersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, labName, name, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[UsersClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[UsersClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete user profile. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
func (client *UsersClient) deleteOperation(ctx context.Context, resourceGroupName string, labName string, name string, options *UsersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "UsersClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, labName, name, options)
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
func (client *UsersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, options *UsersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get user profile.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - name - The name of the user profile.
//   - options - UsersClientGetOptions contains the optional parameters for the UsersClient.Get method.
func (client *UsersClient) Get(ctx context.Context, resourceGroupName string, labName string, name string, options *UsersClientGetOptions) (UsersClientGetResponse, error) {
	var err error
	const operationName = "UsersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, labName, name, options)
	if err != nil {
		return UsersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UsersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return UsersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *UsersClient) getCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, options *UsersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *UsersClient) getHandleResponse(resp *http.Response) (UsersClientGetResponse, error) {
	result := UsersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.User); err != nil {
		return UsersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List user profiles in a given lab.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - options - UsersClientListOptions contains the optional parameters for the UsersClient.NewListPager method.
func (client *UsersClient) NewListPager(resourceGroupName string, labName string, options *UsersClientListOptions) *runtime.Pager[UsersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[UsersClientListResponse]{
		More: func(page UsersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *UsersClientListResponse) (UsersClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "UsersClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, labName, options)
			}, nil)
			if err != nil {
				return UsersClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *UsersClient) listCreateRequest(ctx context.Context, resourceGroupName string, labName string, options *UsersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *UsersClient) listHandleResponse(resp *http.Response) (UsersClientListResponse, error) {
	result := UsersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UserList); err != nil {
		return UsersClientListResponse{}, err
	}
	return result, nil
}

// Update - Allows modifying tags of user profiles. All other properties will be ignored.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-09-15
//   - resourceGroupName - The name of the resource group.
//   - labName - The name of the lab.
//   - name - The name of the user profile.
//   - userParam - Profile of a lab user.
//   - options - UsersClientUpdateOptions contains the optional parameters for the UsersClient.Update method.
func (client *UsersClient) Update(ctx context.Context, resourceGroupName string, labName string, name string, userParam UserFragment, options *UsersClientUpdateOptions) (UsersClientUpdateResponse, error) {
	var err error
	const operationName = "UsersClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, labName, name, userParam, options)
	if err != nil {
		return UsersClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UsersClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return UsersClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *UsersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, userParam UserFragment, options *UsersClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/users/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, userParam); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *UsersClient) updateHandleResponse(resp *http.Response) (UsersClientUpdateResponse, error) {
	result := UsersClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.User); err != nil {
		return UsersClientUpdateResponse{}, err
	}
	return result, nil
}
