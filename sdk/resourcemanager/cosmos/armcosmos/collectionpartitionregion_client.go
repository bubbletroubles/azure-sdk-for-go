//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

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

// CollectionPartitionRegionClient contains the methods for the CollectionPartitionRegion group.
// Don't use this type directly, use NewCollectionPartitionRegionClient() instead.
type CollectionPartitionRegionClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCollectionPartitionRegionClient creates a new instance of CollectionPartitionRegionClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCollectionPartitionRegionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CollectionPartitionRegionClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CollectionPartitionRegionClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListMetricsPager - Retrieves the metrics determined by the given filter for the given collection and region, split by
// partition.
//
// Generated from API version 2023-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - region - Cosmos DB region, with spaces between words and each word capitalized.
//   - databaseRid - Cosmos DB database rid.
//   - collectionRid - Cosmos DB collection rid.
//   - filter - An OData filter expression that describes a subset of metrics to return. The parameters that can be filtered are
//     name.value (name of the metric, can have an or of multiple names), startTime, endTime,
//     and timeGrain. The supported operator is eq.
//   - options - CollectionPartitionRegionClientListMetricsOptions contains the optional parameters for the CollectionPartitionRegionClient.NewListMetricsPager
//     method.
func (client *CollectionPartitionRegionClient) NewListMetricsPager(resourceGroupName string, accountName string, region string, databaseRid string, collectionRid string, filter string, options *CollectionPartitionRegionClientListMetricsOptions) *runtime.Pager[CollectionPartitionRegionClientListMetricsResponse] {
	return runtime.NewPager(runtime.PagingHandler[CollectionPartitionRegionClientListMetricsResponse]{
		More: func(page CollectionPartitionRegionClientListMetricsResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *CollectionPartitionRegionClientListMetricsResponse) (CollectionPartitionRegionClientListMetricsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CollectionPartitionRegionClient.NewListMetricsPager")
			req, err := client.listMetricsCreateRequest(ctx, resourceGroupName, accountName, region, databaseRid, collectionRid, filter, options)
			if err != nil {
				return CollectionPartitionRegionClientListMetricsResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return CollectionPartitionRegionClientListMetricsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CollectionPartitionRegionClientListMetricsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listMetricsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listMetricsCreateRequest creates the ListMetrics request.
func (client *CollectionPartitionRegionClient) listMetricsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, region string, databaseRid string, collectionRid string, filter string, options *CollectionPartitionRegionClientListMetricsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/region/{region}/databases/{databaseRid}/collections/{collectionRid}/partitions/metrics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if region == "" {
		return nil, errors.New("parameter region cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{region}", url.PathEscape(region))
	if databaseRid == "" {
		return nil, errors.New("parameter databaseRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseRid}", url.PathEscape(databaseRid))
	if collectionRid == "" {
		return nil, errors.New("parameter collectionRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{collectionRid}", url.PathEscape(collectionRid))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15-preview")
	reqQP.Set("$filter", filter)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listMetricsHandleResponse handles the ListMetrics response.
func (client *CollectionPartitionRegionClient) listMetricsHandleResponse(resp *http.Response) (CollectionPartitionRegionClientListMetricsResponse, error) {
	result := CollectionPartitionRegionClientListMetricsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PartitionMetricListResult); err != nil {
		return CollectionPartitionRegionClientListMetricsResponse{}, err
	}
	return result, nil
}
