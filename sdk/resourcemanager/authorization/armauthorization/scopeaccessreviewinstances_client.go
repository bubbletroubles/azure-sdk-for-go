//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

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

// ScopeAccessReviewInstancesClient contains the methods for the ScopeAccessReviewInstances group.
// Don't use this type directly, use NewScopeAccessReviewInstancesClient() instead.
type ScopeAccessReviewInstancesClient struct {
	internal *arm.Client
}

// NewScopeAccessReviewInstancesClient creates a new instance of ScopeAccessReviewInstancesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewScopeAccessReviewInstancesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ScopeAccessReviewInstancesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ScopeAccessReviewInstancesClient{
		internal: cl,
	}
	return client, nil
}

// Create - Update access review instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - scope - The scope of the resource.
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - id - The id of the access review instance.
//   - properties - Access review instance properties.
//   - options - ScopeAccessReviewInstancesClientCreateOptions contains the optional parameters for the ScopeAccessReviewInstancesClient.Create
//     method.
func (client *ScopeAccessReviewInstancesClient) Create(ctx context.Context, scope string, scheduleDefinitionID string, id string, properties AccessReviewInstanceProperties, options *ScopeAccessReviewInstancesClientCreateOptions) (ScopeAccessReviewInstancesClientCreateResponse, error) {
	var err error
	const operationName = "ScopeAccessReviewInstancesClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, scope, scheduleDefinitionID, id, properties, options)
	if err != nil {
		return ScopeAccessReviewInstancesClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScopeAccessReviewInstancesClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ScopeAccessReviewInstancesClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *ScopeAccessReviewInstancesClient) createCreateRequest(ctx context.Context, scope string, scheduleDefinitionID string, id string, properties AccessReviewInstanceProperties, options *ScopeAccessReviewInstancesClientCreateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *ScopeAccessReviewInstancesClient) createHandleResponse(resp *http.Response) (ScopeAccessReviewInstancesClientCreateResponse, error) {
	result := ScopeAccessReviewInstancesClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewInstance); err != nil {
		return ScopeAccessReviewInstancesClientCreateResponse{}, err
	}
	return result, nil
}

// GetByID - Get access review instances
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-12-01-preview
//   - scope - The scope of the resource.
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - id - The id of the access review instance.
//   - options - ScopeAccessReviewInstancesClientGetByIDOptions contains the optional parameters for the ScopeAccessReviewInstancesClient.GetByID
//     method.
func (client *ScopeAccessReviewInstancesClient) GetByID(ctx context.Context, scope string, scheduleDefinitionID string, id string, options *ScopeAccessReviewInstancesClientGetByIDOptions) (ScopeAccessReviewInstancesClientGetByIDResponse, error) {
	var err error
	const operationName = "ScopeAccessReviewInstancesClient.GetByID"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getByIDCreateRequest(ctx, scope, scheduleDefinitionID, id, options)
	if err != nil {
		return ScopeAccessReviewInstancesClientGetByIDResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScopeAccessReviewInstancesClientGetByIDResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ScopeAccessReviewInstancesClientGetByIDResponse{}, err
	}
	resp, err := client.getByIDHandleResponse(httpResp)
	return resp, err
}

// getByIDCreateRequest creates the GetByID request.
func (client *ScopeAccessReviewInstancesClient) getByIDCreateRequest(ctx context.Context, scope string, scheduleDefinitionID string, id string, options *ScopeAccessReviewInstancesClientGetByIDOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByIDHandleResponse handles the GetByID response.
func (client *ScopeAccessReviewInstancesClient) getByIDHandleResponse(resp *http.Response) (ScopeAccessReviewInstancesClientGetByIDResponse, error) {
	result := ScopeAccessReviewInstancesClientGetByIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewInstance); err != nil {
		return ScopeAccessReviewInstancesClientGetByIDResponse{}, err
	}
	return result, nil
}

// NewListPager - Get access review instances
//
// Generated from API version 2021-12-01-preview
//   - scope - The scope of the resource.
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - options - ScopeAccessReviewInstancesClientListOptions contains the optional parameters for the ScopeAccessReviewInstancesClient.NewListPager
//     method.
func (client *ScopeAccessReviewInstancesClient) NewListPager(scope string, scheduleDefinitionID string, options *ScopeAccessReviewInstancesClientListOptions) *runtime.Pager[ScopeAccessReviewInstancesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ScopeAccessReviewInstancesClientListResponse]{
		More: func(page ScopeAccessReviewInstancesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ScopeAccessReviewInstancesClientListResponse) (ScopeAccessReviewInstancesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ScopeAccessReviewInstancesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, scope, scheduleDefinitionID, options)
			}, nil)
			if err != nil {
				return ScopeAccessReviewInstancesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ScopeAccessReviewInstancesClient) listCreateRequest(ctx context.Context, scope string, scheduleDefinitionID string, options *ScopeAccessReviewInstancesClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", url.PathEscape(scope))
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ScopeAccessReviewInstancesClient) listHandleResponse(resp *http.Response) (ScopeAccessReviewInstancesClientListResponse, error) {
	result := ScopeAccessReviewInstancesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewInstanceListResult); err != nil {
		return ScopeAccessReviewInstancesClientListResponse{}, err
	}
	return result, nil
}
