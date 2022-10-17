//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// LiveTokenClient contains the methods for the LiveToken group.
// Don't use this type directly, use NewLiveTokenClient() instead.
type LiveTokenClient struct {
	host string
	pl   runtime.Pipeline
}

// NewLiveTokenClient creates a new instance of LiveTokenClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLiveTokenClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*LiveTokenClient, error) {
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
	client := &LiveTokenClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// Get - Gets an access token for live metrics stream data.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-10-14
// resourceURI - The identifier of the resource.
// options - LiveTokenClientGetOptions contains the optional parameters for the LiveTokenClient.Get method.
func (client *LiveTokenClient) Get(ctx context.Context, resourceURI string, options *LiveTokenClientGetOptions) (LiveTokenClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceURI, options)
	if err != nil {
		return LiveTokenClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LiveTokenClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LiveTokenClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LiveTokenClient) getCreateRequest(ctx context.Context, resourceURI string, options *LiveTokenClientGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/generatelivetoken"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-14")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LiveTokenClient) getHandleResponse(resp *http.Response) (LiveTokenClientGetResponse, error) {
	result := LiveTokenClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LiveTokenResponse); err != nil {
		return LiveTokenClientGetResponse{}, err
	}
	return result, nil
}