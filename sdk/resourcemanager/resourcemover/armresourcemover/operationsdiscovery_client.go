//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcemover

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// OperationsDiscoveryClient contains the methods for the OperationsDiscovery group.
// Don't use this type directly, use NewOperationsDiscoveryClient() instead.
type OperationsDiscoveryClient struct {
	internal *arm.Client
}

// NewOperationsDiscoveryClient creates a new instance of OperationsDiscoveryClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOperationsDiscoveryClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*OperationsDiscoveryClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OperationsDiscoveryClient{
		internal: cl,
	}
	return client, nil
}

// Get -
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01
//   - options - OperationsDiscoveryClientGetOptions contains the optional parameters for the OperationsDiscoveryClient.Get method.
func (client *OperationsDiscoveryClient) Get(ctx context.Context, options *OperationsDiscoveryClientGetOptions) (OperationsDiscoveryClientGetResponse, error) {
	var err error
	const operationName = "OperationsDiscoveryClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return OperationsDiscoveryClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsDiscoveryClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OperationsDiscoveryClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *OperationsDiscoveryClient) getCreateRequest(ctx context.Context, options *OperationsDiscoveryClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Migrate/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OperationsDiscoveryClient) getHandleResponse(resp *http.Response) (OperationsDiscoveryClientGetResponse, error) {
	result := OperationsDiscoveryClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationsDiscoveryCollection); err != nil {
		return OperationsDiscoveryClientGetResponse{}, err
	}
	return result, nil
}
