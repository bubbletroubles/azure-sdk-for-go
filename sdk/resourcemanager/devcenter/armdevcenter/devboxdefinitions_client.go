//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevcenter

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// DevBoxDefinitionsClient contains the methods for the DevBoxDefinitions group.
// Don't use this type directly, use NewDevBoxDefinitionsClient() instead.
type DevBoxDefinitionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDevBoxDefinitionsClient creates a new instance of DevBoxDefinitionsClient with the specified values.
// subscriptionID - Unique identifier of the Azure subscription. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDevBoxDefinitionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DevBoxDefinitionsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &DevBoxDefinitionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a Dev Box definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - Name of the resource group within the Azure subscription.
// devCenterName - The name of the devcenter.
// devBoxDefinitionName - The name of the Dev Box definition.
// body - Represents a Dev Box definition.
// options - DevBoxDefinitionsClientBeginCreateOrUpdateOptions contains the optional parameters for the DevBoxDefinitionsClient.BeginCreateOrUpdate
// method.
func (client *DevBoxDefinitionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, devCenterName string, devBoxDefinitionName string, body DevBoxDefinition, options *DevBoxDefinitionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[DevBoxDefinitionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, devCenterName, devBoxDefinitionName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[DevBoxDefinitionsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DevBoxDefinitionsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a Dev Box definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
func (client *DevBoxDefinitionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, devCenterName string, devBoxDefinitionName string, body DevBoxDefinition, options *DevBoxDefinitionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, devCenterName, devBoxDefinitionName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DevBoxDefinitionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, devBoxDefinitionName string, body DevBoxDefinition, options *DevBoxDefinitionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/devboxdefinitions/{devBoxDefinitionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	if devBoxDefinitionName == "" {
		return nil, errors.New("parameter devBoxDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devBoxDefinitionName}", url.PathEscape(devBoxDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Deletes a Dev Box definition
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - Name of the resource group within the Azure subscription.
// devCenterName - The name of the devcenter.
// devBoxDefinitionName - The name of the Dev Box definition.
// options - DevBoxDefinitionsClientBeginDeleteOptions contains the optional parameters for the DevBoxDefinitionsClient.BeginDelete
// method.
func (client *DevBoxDefinitionsClient) BeginDelete(ctx context.Context, resourceGroupName string, devCenterName string, devBoxDefinitionName string, options *DevBoxDefinitionsClientBeginDeleteOptions) (*runtime.Poller[DevBoxDefinitionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, devCenterName, devBoxDefinitionName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[DevBoxDefinitionsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DevBoxDefinitionsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a Dev Box definition
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
func (client *DevBoxDefinitionsClient) deleteOperation(ctx context.Context, resourceGroupName string, devCenterName string, devBoxDefinitionName string, options *DevBoxDefinitionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, devCenterName, devBoxDefinitionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DevBoxDefinitionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, devBoxDefinitionName string, options *DevBoxDefinitionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/devboxdefinitions/{devBoxDefinitionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	if devBoxDefinitionName == "" {
		return nil, errors.New("parameter devBoxDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devBoxDefinitionName}", url.PathEscape(devBoxDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a Dev Box definition
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - Name of the resource group within the Azure subscription.
// devCenterName - The name of the devcenter.
// devBoxDefinitionName - The name of the Dev Box definition.
// options - DevBoxDefinitionsClientGetOptions contains the optional parameters for the DevBoxDefinitionsClient.Get method.
func (client *DevBoxDefinitionsClient) Get(ctx context.Context, resourceGroupName string, devCenterName string, devBoxDefinitionName string, options *DevBoxDefinitionsClientGetOptions) (DevBoxDefinitionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, devCenterName, devBoxDefinitionName, options)
	if err != nil {
		return DevBoxDefinitionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DevBoxDefinitionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DevBoxDefinitionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DevBoxDefinitionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, devBoxDefinitionName string, options *DevBoxDefinitionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/devboxdefinitions/{devBoxDefinitionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	if devBoxDefinitionName == "" {
		return nil, errors.New("parameter devBoxDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devBoxDefinitionName}", url.PathEscape(devBoxDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DevBoxDefinitionsClient) getHandleResponse(resp *http.Response) (DevBoxDefinitionsClientGetResponse, error) {
	result := DevBoxDefinitionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DevBoxDefinition); err != nil {
		return DevBoxDefinitionsClientGetResponse{}, err
	}
	return result, nil
}

// GetByProject - Gets a Dev Box definition configured for a project
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - Name of the resource group within the Azure subscription.
// projectName - The name of the project.
// devBoxDefinitionName - The name of the Dev Box definition.
// options - DevBoxDefinitionsClientGetByProjectOptions contains the optional parameters for the DevBoxDefinitionsClient.GetByProject
// method.
func (client *DevBoxDefinitionsClient) GetByProject(ctx context.Context, resourceGroupName string, projectName string, devBoxDefinitionName string, options *DevBoxDefinitionsClientGetByProjectOptions) (DevBoxDefinitionsClientGetByProjectResponse, error) {
	req, err := client.getByProjectCreateRequest(ctx, resourceGroupName, projectName, devBoxDefinitionName, options)
	if err != nil {
		return DevBoxDefinitionsClientGetByProjectResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DevBoxDefinitionsClientGetByProjectResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DevBoxDefinitionsClientGetByProjectResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByProjectHandleResponse(resp)
}

// getByProjectCreateRequest creates the GetByProject request.
func (client *DevBoxDefinitionsClient) getByProjectCreateRequest(ctx context.Context, resourceGroupName string, projectName string, devBoxDefinitionName string, options *DevBoxDefinitionsClientGetByProjectOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}/devboxdefinitions/{devBoxDefinitionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	if devBoxDefinitionName == "" {
		return nil, errors.New("parameter devBoxDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devBoxDefinitionName}", url.PathEscape(devBoxDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByProjectHandleResponse handles the GetByProject response.
func (client *DevBoxDefinitionsClient) getByProjectHandleResponse(resp *http.Response) (DevBoxDefinitionsClientGetByProjectResponse, error) {
	result := DevBoxDefinitionsClientGetByProjectResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DevBoxDefinition); err != nil {
		return DevBoxDefinitionsClientGetByProjectResponse{}, err
	}
	return result, nil
}

// NewListByDevCenterPager - List Dev Box definitions for a devcenter.
// Generated from API version 2022-09-01-preview
// resourceGroupName - Name of the resource group within the Azure subscription.
// devCenterName - The name of the devcenter.
// options - DevBoxDefinitionsClientListByDevCenterOptions contains the optional parameters for the DevBoxDefinitionsClient.ListByDevCenter
// method.
func (client *DevBoxDefinitionsClient) NewListByDevCenterPager(resourceGroupName string, devCenterName string, options *DevBoxDefinitionsClientListByDevCenterOptions) *runtime.Pager[DevBoxDefinitionsClientListByDevCenterResponse] {
	return runtime.NewPager(runtime.PagingHandler[DevBoxDefinitionsClientListByDevCenterResponse]{
		More: func(page DevBoxDefinitionsClientListByDevCenterResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DevBoxDefinitionsClientListByDevCenterResponse) (DevBoxDefinitionsClientListByDevCenterResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByDevCenterCreateRequest(ctx, resourceGroupName, devCenterName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DevBoxDefinitionsClientListByDevCenterResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DevBoxDefinitionsClientListByDevCenterResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DevBoxDefinitionsClientListByDevCenterResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDevCenterHandleResponse(resp)
		},
	})
}

// listByDevCenterCreateRequest creates the ListByDevCenter request.
func (client *DevBoxDefinitionsClient) listByDevCenterCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, options *DevBoxDefinitionsClientListByDevCenterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/devboxdefinitions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDevCenterHandleResponse handles the ListByDevCenter response.
func (client *DevBoxDefinitionsClient) listByDevCenterHandleResponse(resp *http.Response) (DevBoxDefinitionsClientListByDevCenterResponse, error) {
	result := DevBoxDefinitionsClientListByDevCenterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DevBoxDefinitionListResult); err != nil {
		return DevBoxDefinitionsClientListByDevCenterResponse{}, err
	}
	return result, nil
}

// NewListByProjectPager - List Dev Box definitions configured for a project.
// Generated from API version 2022-09-01-preview
// resourceGroupName - Name of the resource group within the Azure subscription.
// projectName - The name of the project.
// options - DevBoxDefinitionsClientListByProjectOptions contains the optional parameters for the DevBoxDefinitionsClient.ListByProject
// method.
func (client *DevBoxDefinitionsClient) NewListByProjectPager(resourceGroupName string, projectName string, options *DevBoxDefinitionsClientListByProjectOptions) *runtime.Pager[DevBoxDefinitionsClientListByProjectResponse] {
	return runtime.NewPager(runtime.PagingHandler[DevBoxDefinitionsClientListByProjectResponse]{
		More: func(page DevBoxDefinitionsClientListByProjectResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DevBoxDefinitionsClientListByProjectResponse) (DevBoxDefinitionsClientListByProjectResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByProjectCreateRequest(ctx, resourceGroupName, projectName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DevBoxDefinitionsClientListByProjectResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DevBoxDefinitionsClientListByProjectResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DevBoxDefinitionsClientListByProjectResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByProjectHandleResponse(resp)
		},
	})
}

// listByProjectCreateRequest creates the ListByProject request.
func (client *DevBoxDefinitionsClient) listByProjectCreateRequest(ctx context.Context, resourceGroupName string, projectName string, options *DevBoxDefinitionsClientListByProjectOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/projects/{projectName}/devboxdefinitions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if projectName == "" {
		return nil, errors.New("parameter projectName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{projectName}", url.PathEscape(projectName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByProjectHandleResponse handles the ListByProject response.
func (client *DevBoxDefinitionsClient) listByProjectHandleResponse(resp *http.Response) (DevBoxDefinitionsClientListByProjectResponse, error) {
	result := DevBoxDefinitionsClientListByProjectResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DevBoxDefinitionListResult); err != nil {
		return DevBoxDefinitionsClientListByProjectResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Partially updates a Dev Box definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
// resourceGroupName - Name of the resource group within the Azure subscription.
// devCenterName - The name of the devcenter.
// devBoxDefinitionName - The name of the Dev Box definition.
// body - Represents a Dev Box definition.
// options - DevBoxDefinitionsClientBeginUpdateOptions contains the optional parameters for the DevBoxDefinitionsClient.BeginUpdate
// method.
func (client *DevBoxDefinitionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, devCenterName string, devBoxDefinitionName string, body DevBoxDefinitionUpdate, options *DevBoxDefinitionsClientBeginUpdateOptions) (*runtime.Poller[DevBoxDefinitionsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, devCenterName, devBoxDefinitionName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[DevBoxDefinitionsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DevBoxDefinitionsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Partially updates a Dev Box definition.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-preview
func (client *DevBoxDefinitionsClient) update(ctx context.Context, resourceGroupName string, devCenterName string, devBoxDefinitionName string, body DevBoxDefinitionUpdate, options *DevBoxDefinitionsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, devCenterName, devBoxDefinitionName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *DevBoxDefinitionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, devBoxDefinitionName string, body DevBoxDefinitionUpdate, options *DevBoxDefinitionsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/devboxdefinitions/{devBoxDefinitionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	if devBoxDefinitionName == "" {
		return nil, errors.New("parameter devBoxDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devBoxDefinitionName}", url.PathEscape(devBoxDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}