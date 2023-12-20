//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration/v2"
	"net/http"
	"net/url"
	"regexp"
)

// PrivateLinkResourcesServer is a fake server for instances of the armappconfiguration.PrivateLinkResourcesClient type.
type PrivateLinkResourcesServer struct {
	// Get is the fake for method PrivateLinkResourcesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, configStoreName string, groupName string, options *armappconfiguration.PrivateLinkResourcesClientGetOptions) (resp azfake.Responder[armappconfiguration.PrivateLinkResourcesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByConfigurationStorePager is the fake for method PrivateLinkResourcesClient.NewListByConfigurationStorePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByConfigurationStorePager func(resourceGroupName string, configStoreName string, options *armappconfiguration.PrivateLinkResourcesClientListByConfigurationStoreOptions) (resp azfake.PagerResponder[armappconfiguration.PrivateLinkResourcesClientListByConfigurationStoreResponse])
}

// NewPrivateLinkResourcesServerTransport creates a new instance of PrivateLinkResourcesServerTransport with the provided implementation.
// The returned PrivateLinkResourcesServerTransport instance is connected to an instance of armappconfiguration.PrivateLinkResourcesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateLinkResourcesServerTransport(srv *PrivateLinkResourcesServer) *PrivateLinkResourcesServerTransport {
	return &PrivateLinkResourcesServerTransport{
		srv:                              srv,
		newListByConfigurationStorePager: newTracker[azfake.PagerResponder[armappconfiguration.PrivateLinkResourcesClientListByConfigurationStoreResponse]](),
	}
}

// PrivateLinkResourcesServerTransport connects instances of armappconfiguration.PrivateLinkResourcesClient to instances of PrivateLinkResourcesServer.
// Don't use this type directly, use NewPrivateLinkResourcesServerTransport instead.
type PrivateLinkResourcesServerTransport struct {
	srv                              *PrivateLinkResourcesServer
	newListByConfigurationStorePager *tracker[azfake.PagerResponder[armappconfiguration.PrivateLinkResourcesClientListByConfigurationStoreResponse]]
}

// Do implements the policy.Transporter interface for PrivateLinkResourcesServerTransport.
func (p *PrivateLinkResourcesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PrivateLinkResourcesClient.Get":
		resp, err = p.dispatchGet(req)
	case "PrivateLinkResourcesClient.NewListByConfigurationStorePager":
		resp, err = p.dispatchNewListByConfigurationStorePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PrivateLinkResourcesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppConfiguration/configurationStores/(?P<configStoreName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateLinkResources/(?P<groupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	configStoreNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configStoreName")])
	if err != nil {
		return nil, err
	}
	groupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, configStoreNameParam, groupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateLinkResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateLinkResourcesServerTransport) dispatchNewListByConfigurationStorePager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByConfigurationStorePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByConfigurationStorePager not implemented")}
	}
	newListByConfigurationStorePager := p.newListByConfigurationStorePager.get(req)
	if newListByConfigurationStorePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppConfiguration/configurationStores/(?P<configStoreName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateLinkResources`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		configStoreNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configStoreName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListByConfigurationStorePager(resourceGroupNameParam, configStoreNameParam, nil)
		newListByConfigurationStorePager = &resp
		p.newListByConfigurationStorePager.add(req, newListByConfigurationStorePager)
		server.PagerResponderInjectNextLinks(newListByConfigurationStorePager, req, func(page *armappconfiguration.PrivateLinkResourcesClientListByConfigurationStoreResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByConfigurationStorePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByConfigurationStorePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByConfigurationStorePager) {
		p.newListByConfigurationStorePager.remove(req)
	}
	return resp, nil
}
