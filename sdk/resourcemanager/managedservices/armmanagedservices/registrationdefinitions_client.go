//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagedservices

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

// RegistrationDefinitionsClient contains the methods for the RegistrationDefinitions group.
// Don't use this type directly, use NewRegistrationDefinitionsClient() instead.
type RegistrationDefinitionsClient struct {
	internal *arm.Client
}

// NewRegistrationDefinitionsClient creates a new instance of RegistrationDefinitionsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRegistrationDefinitionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*RegistrationDefinitionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RegistrationDefinitionsClient{
		internal: cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a registration definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-01-preview
//   - registrationDefinitionID - The GUID of the registration definition.
//   - scope - The scope of the resource.
//   - requestBody - The parameters required to create a new registration definition.
//   - options - RegistrationDefinitionsClientBeginCreateOrUpdateOptions contains the optional parameters for the RegistrationDefinitionsClient.BeginCreateOrUpdate
//     method.
func (client *RegistrationDefinitionsClient) BeginCreateOrUpdate(ctx context.Context, registrationDefinitionID string, scope string, requestBody RegistrationDefinition, options *RegistrationDefinitionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[RegistrationDefinitionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, registrationDefinitionID, scope, requestBody, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RegistrationDefinitionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[RegistrationDefinitionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates a registration definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-01-preview
func (client *RegistrationDefinitionsClient) createOrUpdate(ctx context.Context, registrationDefinitionID string, scope string, requestBody RegistrationDefinition, options *RegistrationDefinitionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "RegistrationDefinitionsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, registrationDefinitionID, scope, requestBody, options)
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
func (client *RegistrationDefinitionsClient) createOrUpdateCreateRequest(ctx context.Context, registrationDefinitionID string, scope string, requestBody RegistrationDefinition, options *RegistrationDefinitionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions/{registrationDefinitionId}"
	if registrationDefinitionID == "" {
		return nil, errors.New("parameter registrationDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registrationDefinitionId}", url.PathEscape(registrationDefinitionID))
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, requestBody); err != nil {
		return nil, err
	}
	return req, nil
}

// Delete - Deletes the registration definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-01-preview
//   - registrationDefinitionID - The GUID of the registration definition.
//   - scope - The scope of the resource.
//   - options - RegistrationDefinitionsClientDeleteOptions contains the optional parameters for the RegistrationDefinitionsClient.Delete
//     method.
func (client *RegistrationDefinitionsClient) Delete(ctx context.Context, registrationDefinitionID string, scope string, options *RegistrationDefinitionsClientDeleteOptions) (RegistrationDefinitionsClientDeleteResponse, error) {
	var err error
	const operationName = "RegistrationDefinitionsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, registrationDefinitionID, scope, options)
	if err != nil {
		return RegistrationDefinitionsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RegistrationDefinitionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return RegistrationDefinitionsClientDeleteResponse{}, err
	}
	return RegistrationDefinitionsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RegistrationDefinitionsClient) deleteCreateRequest(ctx context.Context, registrationDefinitionID string, scope string, options *RegistrationDefinitionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions/{registrationDefinitionId}"
	if registrationDefinitionID == "" {
		return nil, errors.New("parameter registrationDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registrationDefinitionId}", url.PathEscape(registrationDefinitionID))
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the registration definition details.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-01-preview
//   - scope - The scope of the resource.
//   - registrationDefinitionID - The GUID of the registration definition.
//   - options - RegistrationDefinitionsClientGetOptions contains the optional parameters for the RegistrationDefinitionsClient.Get
//     method.
func (client *RegistrationDefinitionsClient) Get(ctx context.Context, scope string, registrationDefinitionID string, options *RegistrationDefinitionsClientGetOptions) (RegistrationDefinitionsClientGetResponse, error) {
	var err error
	const operationName = "RegistrationDefinitionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, scope, registrationDefinitionID, options)
	if err != nil {
		return RegistrationDefinitionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RegistrationDefinitionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RegistrationDefinitionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RegistrationDefinitionsClient) getCreateRequest(ctx context.Context, scope string, registrationDefinitionID string, options *RegistrationDefinitionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions/{registrationDefinitionId}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if registrationDefinitionID == "" {
		return nil, errors.New("parameter registrationDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registrationDefinitionId}", url.PathEscape(registrationDefinitionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RegistrationDefinitionsClient) getHandleResponse(resp *http.Response) (RegistrationDefinitionsClientGetResponse, error) {
	result := RegistrationDefinitionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegistrationDefinition); err != nil {
		return RegistrationDefinitionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of the registration definitions.
//
// Generated from API version 2022-01-01-preview
//   - scope - The scope of the resource.
//   - options - RegistrationDefinitionsClientListOptions contains the optional parameters for the RegistrationDefinitionsClient.NewListPager
//     method.
func (client *RegistrationDefinitionsClient) NewListPager(scope string, options *RegistrationDefinitionsClientListOptions) *runtime.Pager[RegistrationDefinitionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RegistrationDefinitionsClientListResponse]{
		More: func(page RegistrationDefinitionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RegistrationDefinitionsClientListResponse) (RegistrationDefinitionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RegistrationDefinitionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, scope, options)
			}, nil)
			if err != nil {
				return RegistrationDefinitionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *RegistrationDefinitionsClient) listCreateRequest(ctx context.Context, scope string, options *RegistrationDefinitionsClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.ManagedServices/registrationDefinitions"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RegistrationDefinitionsClient) listHandleResponse(resp *http.Response) (RegistrationDefinitionsClientListResponse, error) {
	result := RegistrationDefinitionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegistrationDefinitionList); err != nil {
		return RegistrationDefinitionsClientListResponse{}, err
	}
	return result, nil
}
