//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package armconnectedvmware

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

// DatastoresClient contains the methods for the Datastores group.
// Don't use this type directly, use NewDatastoresClient() instead.
type DatastoresClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDatastoresClient creates a new instance of DatastoresClient with the specified values.
//   - subscriptionID - The Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDatastoresClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DatastoresClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DatastoresClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Create Or Update datastore.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
//   - resourceGroupName - The Resource Group Name.
//   - datastoreName - Name of the datastore.
//   - body - Request payload.
//   - options - DatastoresClientBeginCreateOptions contains the optional parameters for the DatastoresClient.BeginCreate method.
func (client *DatastoresClient) BeginCreate(ctx context.Context, resourceGroupName string, datastoreName string, body Datastore, options *DatastoresClientBeginCreateOptions) (*runtime.Poller[DatastoresClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, datastoreName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DatastoresClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DatastoresClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Create Or Update datastore.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
func (client *DatastoresClient) create(ctx context.Context, resourceGroupName string, datastoreName string, body Datastore, options *DatastoresClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "DatastoresClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, datastoreName, body, options)
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

// createCreateRequest creates the Create request.
func (client *DatastoresClient) createCreateRequest(ctx context.Context, resourceGroupName string, datastoreName string, body Datastore, options *DatastoresClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/datastores/{datastoreName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if datastoreName == "" {
		return nil, errors.New("parameter datastoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{datastoreName}", url.PathEscape(datastoreName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Implements datastore DELETE method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
//   - resourceGroupName - The Resource Group Name.
//   - datastoreName - Name of the datastore.
//   - options - DatastoresClientBeginDeleteOptions contains the optional parameters for the DatastoresClient.BeginDelete method.
func (client *DatastoresClient) BeginDelete(ctx context.Context, resourceGroupName string, datastoreName string, options *DatastoresClientBeginDeleteOptions) (*runtime.Poller[DatastoresClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, datastoreName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DatastoresClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DatastoresClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Implements datastore DELETE method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
func (client *DatastoresClient) deleteOperation(ctx context.Context, resourceGroupName string, datastoreName string, options *DatastoresClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "DatastoresClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, datastoreName, options)
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
func (client *DatastoresClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, datastoreName string, options *DatastoresClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/datastores/{datastoreName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if datastoreName == "" {
		return nil, errors.New("parameter datastoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{datastoreName}", url.PathEscape(datastoreName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01")
	if options != nil && options.Force != nil {
		reqQP.Set("force", strconv.FormatBool(*options.Force))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Implements datastore GET method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
//   - resourceGroupName - The Resource Group Name.
//   - datastoreName - Name of the datastore.
//   - options - DatastoresClientGetOptions contains the optional parameters for the DatastoresClient.Get method.
func (client *DatastoresClient) Get(ctx context.Context, resourceGroupName string, datastoreName string, options *DatastoresClientGetOptions) (DatastoresClientGetResponse, error) {
	var err error
	const operationName = "DatastoresClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, datastoreName, options)
	if err != nil {
		return DatastoresClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatastoresClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DatastoresClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DatastoresClient) getCreateRequest(ctx context.Context, resourceGroupName string, datastoreName string, options *DatastoresClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/datastores/{datastoreName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if datastoreName == "" {
		return nil, errors.New("parameter datastoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{datastoreName}", url.PathEscape(datastoreName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DatastoresClient) getHandleResponse(resp *http.Response) (DatastoresClientGetResponse, error) {
	result := DatastoresClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Datastore); err != nil {
		return DatastoresClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List of datastores in a subscription.
//
// Generated from API version 2023-10-01
//   - options - DatastoresClientListOptions contains the optional parameters for the DatastoresClient.NewListPager method.
func (client *DatastoresClient) NewListPager(options *DatastoresClientListOptions) *runtime.Pager[DatastoresClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DatastoresClientListResponse]{
		More: func(page DatastoresClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DatastoresClientListResponse) (DatastoresClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DatastoresClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return DatastoresClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *DatastoresClient) listCreateRequest(ctx context.Context, options *DatastoresClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/datastores"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DatastoresClient) listHandleResponse(resp *http.Response) (DatastoresClientListResponse, error) {
	result := DatastoresClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatastoresList); err != nil {
		return DatastoresClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List of datastores in a resource group.
//
// Generated from API version 2023-10-01
//   - resourceGroupName - The Resource Group Name.
//   - options - DatastoresClientListByResourceGroupOptions contains the optional parameters for the DatastoresClient.NewListByResourceGroupPager
//     method.
func (client *DatastoresClient) NewListByResourceGroupPager(resourceGroupName string, options *DatastoresClientListByResourceGroupOptions) *runtime.Pager[DatastoresClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[DatastoresClientListByResourceGroupResponse]{
		More: func(page DatastoresClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DatastoresClientListByResourceGroupResponse) (DatastoresClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DatastoresClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return DatastoresClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DatastoresClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DatastoresClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/datastores"
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
	reqQP.Set("api-version", "2023-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DatastoresClient) listByResourceGroupHandleResponse(resp *http.Response) (DatastoresClientListByResourceGroupResponse, error) {
	result := DatastoresClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatastoresList); err != nil {
		return DatastoresClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// Update - API to update certain properties of the datastore resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01
//   - resourceGroupName - The Resource Group Name.
//   - datastoreName - Name of the datastore.
//   - body - Resource properties to update.
//   - options - DatastoresClientUpdateOptions contains the optional parameters for the DatastoresClient.Update method.
func (client *DatastoresClient) Update(ctx context.Context, resourceGroupName string, datastoreName string, body ResourcePatch, options *DatastoresClientUpdateOptions) (DatastoresClientUpdateResponse, error) {
	var err error
	const operationName = "DatastoresClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, datastoreName, body, options)
	if err != nil {
		return DatastoresClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatastoresClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DatastoresClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *DatastoresClient) updateCreateRequest(ctx context.Context, resourceGroupName string, datastoreName string, body ResourcePatch, options *DatastoresClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ConnectedVMwarevSphere/datastores/{datastoreName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if datastoreName == "" {
		return nil, errors.New("parameter datastoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{datastoreName}", url.PathEscape(datastoreName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *DatastoresClient) updateHandleResponse(resp *http.Response) (DatastoresClientUpdateResponse, error) {
	result := DatastoresClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Datastore); err != nil {
		return DatastoresClientUpdateResponse{}, err
	}
	return result, nil
}
