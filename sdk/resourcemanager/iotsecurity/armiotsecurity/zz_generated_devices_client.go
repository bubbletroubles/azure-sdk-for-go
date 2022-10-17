//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiotsecurity

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

// DevicesClient contains the methods for the Devices group.
// Don't use this type directly, use NewDevicesClient() instead.
type DevicesClient struct {
	host                string
	subscriptionID      string
	iotDefenderLocation string
	pl                  runtime.Pipeline
}

// NewDevicesClient creates a new instance of DevicesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// iotDefenderLocation - Defender for IoT location
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDevicesClient(subscriptionID string, iotDefenderLocation string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DevicesClient, error) {
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
	client := &DevicesClient{
		subscriptionID:      subscriptionID,
		iotDefenderLocation: iotDefenderLocation,
		host:                ep,
		pl:                  pl,
	}
	return client, nil
}

// Get - Get device
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-02-01-preview
// deviceGroupName - Device group name
// deviceID - Device Id
// options - DevicesClientGetOptions contains the optional parameters for the DevicesClient.Get method.
func (client *DevicesClient) Get(ctx context.Context, deviceGroupName string, deviceID string, options *DevicesClientGetOptions) (DevicesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, deviceGroupName, deviceID, options)
	if err != nil {
		return DevicesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DevicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DevicesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DevicesClient) getCreateRequest(ctx context.Context, deviceGroupName string, deviceID string, options *DevicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTSecurity/locations/{iotDefenderLocation}/deviceGroups/{deviceGroupName}/devices/{deviceId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if client.iotDefenderLocation == "" {
		return nil, errors.New("parameter client.iotDefenderLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iotDefenderLocation}", url.PathEscape(client.iotDefenderLocation))
	if deviceGroupName == "" {
		return nil, errors.New("parameter deviceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceGroupName}", url.PathEscape(deviceGroupName))
	if deviceID == "" {
		return nil, errors.New("parameter deviceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceId}", url.PathEscape(deviceID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DevicesClient) getHandleResponse(resp *http.Response) (DevicesClientGetResponse, error) {
	result := DevicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeviceModel); err != nil {
		return DevicesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List devices
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-02-01-preview
// deviceGroupName - Device group name
// options - DevicesClientListOptions contains the optional parameters for the DevicesClient.List method.
func (client *DevicesClient) NewListPager(deviceGroupName string, options *DevicesClientListOptions) *runtime.Pager[DevicesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DevicesClientListResponse]{
		More: func(page DevicesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DevicesClientListResponse) (DevicesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, deviceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DevicesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DevicesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DevicesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DevicesClient) listCreateRequest(ctx context.Context, deviceGroupName string, options *DevicesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTSecurity/locations/{iotDefenderLocation}/deviceGroups/{deviceGroupName}/devices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if client.iotDefenderLocation == "" {
		return nil, errors.New("parameter client.iotDefenderLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iotDefenderLocation}", url.PathEscape(client.iotDefenderLocation))
	if deviceGroupName == "" {
		return nil, errors.New("parameter deviceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceGroupName}", url.PathEscape(deviceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DevicesClient) listHandleResponse(resp *http.Response) (DevicesClientListResponse, error) {
	result := DevicesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeviceList); err != nil {
		return DevicesClientListResponse{}, err
	}
	return result, nil
}