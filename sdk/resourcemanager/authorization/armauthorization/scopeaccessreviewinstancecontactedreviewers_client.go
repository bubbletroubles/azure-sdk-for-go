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

// ScopeAccessReviewInstanceContactedReviewersClient contains the methods for the ScopeAccessReviewInstanceContactedReviewers group.
// Don't use this type directly, use NewScopeAccessReviewInstanceContactedReviewersClient() instead.
type ScopeAccessReviewInstanceContactedReviewersClient struct {
	internal *arm.Client
}

// NewScopeAccessReviewInstanceContactedReviewersClient creates a new instance of ScopeAccessReviewInstanceContactedReviewersClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewScopeAccessReviewInstanceContactedReviewersClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ScopeAccessReviewInstanceContactedReviewersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ScopeAccessReviewInstanceContactedReviewersClient{
		internal: cl,
	}
	return client, nil
}

// NewListPager - Get access review instance contacted reviewers
//
// Generated from API version 2021-12-01-preview
//   - scope - The scope of the resource.
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - id - The id of the access review instance.
//   - options - ScopeAccessReviewInstanceContactedReviewersClientListOptions contains the optional parameters for the ScopeAccessReviewInstanceContactedReviewersClient.NewListPager
//     method.
func (client *ScopeAccessReviewInstanceContactedReviewersClient) NewListPager(scope string, scheduleDefinitionID string, id string, options *ScopeAccessReviewInstanceContactedReviewersClientListOptions) *runtime.Pager[ScopeAccessReviewInstanceContactedReviewersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ScopeAccessReviewInstanceContactedReviewersClientListResponse]{
		More: func(page ScopeAccessReviewInstanceContactedReviewersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ScopeAccessReviewInstanceContactedReviewersClientListResponse) (ScopeAccessReviewInstanceContactedReviewersClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ScopeAccessReviewInstanceContactedReviewersClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, scope, scheduleDefinitionID, id, options)
			}, nil)
			if err != nil {
				return ScopeAccessReviewInstanceContactedReviewersClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ScopeAccessReviewInstanceContactedReviewersClient) listCreateRequest(ctx context.Context, scope string, scheduleDefinitionID string, id string, options *ScopeAccessReviewInstanceContactedReviewersClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}/contactedReviewers"
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

// listHandleResponse handles the List response.
func (client *ScopeAccessReviewInstanceContactedReviewersClient) listHandleResponse(resp *http.Response) (ScopeAccessReviewInstanceContactedReviewersClientListResponse, error) {
	result := ScopeAccessReviewInstanceContactedReviewersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewContactedReviewerListResult); err != nil {
		return ScopeAccessReviewInstanceContactedReviewersClientListResponse{}, err
	}
	return result, nil
}
