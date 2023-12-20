//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataboxedge

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

// StorageAccountCredentialsClient contains the methods for the StorageAccountCredentials group.
// Don't use this type directly, use NewStorageAccountCredentialsClient() instead.
type StorageAccountCredentialsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewStorageAccountCredentialsClient creates a new instance of StorageAccountCredentialsClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewStorageAccountCredentialsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*StorageAccountCredentialsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &StorageAccountCredentialsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates the storage account credential.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01
//   - deviceName - The device name.
//   - name - The storage account credential name.
//   - resourceGroupName - The resource group name.
//   - storageAccountCredential - The storage account credential.
//   - options - StorageAccountCredentialsClientBeginCreateOrUpdateOptions contains the optional parameters for the StorageAccountCredentialsClient.BeginCreateOrUpdate
//     method.
func (client *StorageAccountCredentialsClient) BeginCreateOrUpdate(ctx context.Context, deviceName string, name string, resourceGroupName string, storageAccountCredential StorageAccountCredential, options *StorageAccountCredentialsClientBeginCreateOrUpdateOptions) (*runtime.Poller[StorageAccountCredentialsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, deviceName, name, resourceGroupName, storageAccountCredential, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[StorageAccountCredentialsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[StorageAccountCredentialsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates the storage account credential.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01
func (client *StorageAccountCredentialsClient) createOrUpdate(ctx context.Context, deviceName string, name string, resourceGroupName string, storageAccountCredential StorageAccountCredential, options *StorageAccountCredentialsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "StorageAccountCredentialsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, deviceName, name, resourceGroupName, storageAccountCredential, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *StorageAccountCredentialsClient) createOrUpdateCreateRequest(ctx context.Context, deviceName string, name string, resourceGroupName string, storageAccountCredential StorageAccountCredential, options *StorageAccountCredentialsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccountCredentials/{name}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, storageAccountCredential); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the storage account credential.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01
//   - deviceName - The device name.
//   - name - The storage account credential name.
//   - resourceGroupName - The resource group name.
//   - options - StorageAccountCredentialsClientBeginDeleteOptions contains the optional parameters for the StorageAccountCredentialsClient.BeginDelete
//     method.
func (client *StorageAccountCredentialsClient) BeginDelete(ctx context.Context, deviceName string, name string, resourceGroupName string, options *StorageAccountCredentialsClientBeginDeleteOptions) (*runtime.Poller[StorageAccountCredentialsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, deviceName, name, resourceGroupName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[StorageAccountCredentialsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[StorageAccountCredentialsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the storage account credential.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01
func (client *StorageAccountCredentialsClient) deleteOperation(ctx context.Context, deviceName string, name string, resourceGroupName string, options *StorageAccountCredentialsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "StorageAccountCredentialsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, deviceName, name, resourceGroupName, options)
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
func (client *StorageAccountCredentialsClient) deleteCreateRequest(ctx context.Context, deviceName string, name string, resourceGroupName string, options *StorageAccountCredentialsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccountCredentials/{name}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the properties of the specified storage account credential.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01
//   - deviceName - The device name.
//   - name - The storage account credential name.
//   - resourceGroupName - The resource group name.
//   - options - StorageAccountCredentialsClientGetOptions contains the optional parameters for the StorageAccountCredentialsClient.Get
//     method.
func (client *StorageAccountCredentialsClient) Get(ctx context.Context, deviceName string, name string, resourceGroupName string, options *StorageAccountCredentialsClientGetOptions) (StorageAccountCredentialsClientGetResponse, error) {
	var err error
	const operationName = "StorageAccountCredentialsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, deviceName, name, resourceGroupName, options)
	if err != nil {
		return StorageAccountCredentialsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StorageAccountCredentialsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StorageAccountCredentialsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *StorageAccountCredentialsClient) getCreateRequest(ctx context.Context, deviceName string, name string, resourceGroupName string, options *StorageAccountCredentialsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccountCredentials/{name}"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
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
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *StorageAccountCredentialsClient) getHandleResponse(resp *http.Response) (StorageAccountCredentialsClientGetResponse, error) {
	result := StorageAccountCredentialsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageAccountCredential); err != nil {
		return StorageAccountCredentialsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDataBoxEdgeDevicePager - Gets all the storage account credentials in a Data Box Edge/Data Box Gateway device.
//
// Generated from API version 2022-03-01
//   - deviceName - The device name.
//   - resourceGroupName - The resource group name.
//   - options - StorageAccountCredentialsClientListByDataBoxEdgeDeviceOptions contains the optional parameters for the StorageAccountCredentialsClient.NewListByDataBoxEdgeDevicePager
//     method.
func (client *StorageAccountCredentialsClient) NewListByDataBoxEdgeDevicePager(deviceName string, resourceGroupName string, options *StorageAccountCredentialsClientListByDataBoxEdgeDeviceOptions) *runtime.Pager[StorageAccountCredentialsClientListByDataBoxEdgeDeviceResponse] {
	return runtime.NewPager(runtime.PagingHandler[StorageAccountCredentialsClientListByDataBoxEdgeDeviceResponse]{
		More: func(page StorageAccountCredentialsClientListByDataBoxEdgeDeviceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *StorageAccountCredentialsClientListByDataBoxEdgeDeviceResponse) (StorageAccountCredentialsClientListByDataBoxEdgeDeviceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "StorageAccountCredentialsClient.NewListByDataBoxEdgeDevicePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByDataBoxEdgeDeviceCreateRequest(ctx, deviceName, resourceGroupName, options)
			}, nil)
			if err != nil {
				return StorageAccountCredentialsClientListByDataBoxEdgeDeviceResponse{}, err
			}
			return client.listByDataBoxEdgeDeviceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByDataBoxEdgeDeviceCreateRequest creates the ListByDataBoxEdgeDevice request.
func (client *StorageAccountCredentialsClient) listByDataBoxEdgeDeviceCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, options *StorageAccountCredentialsClientListByDataBoxEdgeDeviceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccountCredentials"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
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
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDataBoxEdgeDeviceHandleResponse handles the ListByDataBoxEdgeDevice response.
func (client *StorageAccountCredentialsClient) listByDataBoxEdgeDeviceHandleResponse(resp *http.Response) (StorageAccountCredentialsClientListByDataBoxEdgeDeviceResponse, error) {
	result := StorageAccountCredentialsClientListByDataBoxEdgeDeviceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageAccountCredentialList); err != nil {
		return StorageAccountCredentialsClientListByDataBoxEdgeDeviceResponse{}, err
	}
	return result, nil
}
