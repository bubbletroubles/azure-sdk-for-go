//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

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

// LibrariesClient contains the methods for the Libraries group.
// Don't use this type directly, use NewLibrariesClient() instead.
type LibrariesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLibrariesClient creates a new instance of LibrariesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLibrariesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LibrariesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LibrariesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListByWorkspacePager - List libraries in a workspace.
//
// Generated from API version 2021-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - LibrariesClientListByWorkspaceOptions contains the optional parameters for the LibrariesClient.NewListByWorkspacePager
//     method.
func (client *LibrariesClient) NewListByWorkspacePager(resourceGroupName string, workspaceName string, options *LibrariesClientListByWorkspaceOptions) *runtime.Pager[LibrariesClientListByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[LibrariesClientListByWorkspaceResponse]{
		More: func(page LibrariesClientListByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LibrariesClientListByWorkspaceResponse) (LibrariesClientListByWorkspaceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "LibrariesClient.NewListByWorkspacePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
			}, nil)
			if err != nil {
				return LibrariesClientListByWorkspaceResponse{}, err
			}
			return client.listByWorkspaceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *LibrariesClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *LibrariesClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/libraries"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *LibrariesClient) listByWorkspaceHandleResponse(resp *http.Response) (LibrariesClientListByWorkspaceResponse, error) {
	result := LibrariesClientListByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LibraryListResponse); err != nil {
		return LibrariesClientListByWorkspaceResponse{}, err
	}
	return result, nil
}
