//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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
	"strings"
)

// DataWarehouseUserActivitiesClient contains the methods for the DataWarehouseUserActivities group.
// Don't use this type directly, use NewDataWarehouseUserActivitiesClient() instead.
type DataWarehouseUserActivitiesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDataWarehouseUserActivitiesClient creates a new instance of DataWarehouseUserActivitiesClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDataWarehouseUserActivitiesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DataWarehouseUserActivitiesClient, error) {
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
	client := &DataWarehouseUserActivitiesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets the user activities of a data warehouse which includes running and suspended queries
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// databaseName - The name of the database.
// dataWarehouseUserActivityName - The activity name of the data warehouse.
// options - DataWarehouseUserActivitiesClientGetOptions contains the optional parameters for the DataWarehouseUserActivitiesClient.Get
// method.
func (client *DataWarehouseUserActivitiesClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, dataWarehouseUserActivityName DataWarehouseUserActivityName, options *DataWarehouseUserActivitiesClientGetOptions) (DataWarehouseUserActivitiesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, databaseName, dataWarehouseUserActivityName, options)
	if err != nil {
		return DataWarehouseUserActivitiesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataWarehouseUserActivitiesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataWarehouseUserActivitiesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DataWarehouseUserActivitiesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, dataWarehouseUserActivityName DataWarehouseUserActivityName, options *DataWarehouseUserActivitiesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/dataWarehouseUserActivities/{dataWarehouseUserActivityName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if dataWarehouseUserActivityName == "" {
		return nil, errors.New("parameter dataWarehouseUserActivityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataWarehouseUserActivityName}", url.PathEscape(string(dataWarehouseUserActivityName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataWarehouseUserActivitiesClient) getHandleResponse(resp *http.Response) (DataWarehouseUserActivitiesClientGetResponse, error) {
	result := DataWarehouseUserActivitiesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataWarehouseUserActivities); err != nil {
		return DataWarehouseUserActivitiesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDatabasePager - List the user activities of a data warehouse which includes running and suspended queries
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-11-01-preview
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// databaseName - The name of the database.
// options - DataWarehouseUserActivitiesClientListByDatabaseOptions contains the optional parameters for the DataWarehouseUserActivitiesClient.ListByDatabase
// method.
func (client *DataWarehouseUserActivitiesClient) NewListByDatabasePager(resourceGroupName string, serverName string, databaseName string, options *DataWarehouseUserActivitiesClientListByDatabaseOptions) *runtime.Pager[DataWarehouseUserActivitiesClientListByDatabaseResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataWarehouseUserActivitiesClientListByDatabaseResponse]{
		More: func(page DataWarehouseUserActivitiesClientListByDatabaseResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataWarehouseUserActivitiesClientListByDatabaseResponse) (DataWarehouseUserActivitiesClientListByDatabaseResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByDatabaseCreateRequest(ctx, resourceGroupName, serverName, databaseName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DataWarehouseUserActivitiesClientListByDatabaseResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DataWarehouseUserActivitiesClientListByDatabaseResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DataWarehouseUserActivitiesClientListByDatabaseResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDatabaseHandleResponse(resp)
		},
	})
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *DataWarehouseUserActivitiesClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, options *DataWarehouseUserActivitiesClientListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/dataWarehouseUserActivities"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *DataWarehouseUserActivitiesClient) listByDatabaseHandleResponse(resp *http.Response) (DataWarehouseUserActivitiesClientListByDatabaseResponse, error) {
	result := DataWarehouseUserActivitiesClientListByDatabaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataWarehouseUserActivitiesListResult); err != nil {
		return DataWarehouseUserActivitiesClientListByDatabaseResponse{}, err
	}
	return result, nil
}