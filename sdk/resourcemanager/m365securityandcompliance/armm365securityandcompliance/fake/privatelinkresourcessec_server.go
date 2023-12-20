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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/m365securityandcompliance/armm365securityandcompliance"
	"net/http"
	"net/url"
	"regexp"
)

// PrivateLinkResourcesSecServer is a fake server for instances of the armm365securityandcompliance.PrivateLinkResourcesSecClient type.
type PrivateLinkResourcesSecServer struct {
	// Get is the fake for method PrivateLinkResourcesSecClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, resourceName string, groupName string, options *armm365securityandcompliance.PrivateLinkResourcesSecClientGetOptions) (resp azfake.Responder[armm365securityandcompliance.PrivateLinkResourcesSecClientGetResponse], errResp azfake.ErrorResponder)

	// ListByService is the fake for method PrivateLinkResourcesSecClient.ListByService
	// HTTP status codes to indicate success: http.StatusOK
	ListByService func(ctx context.Context, resourceGroupName string, resourceName string, options *armm365securityandcompliance.PrivateLinkResourcesSecClientListByServiceOptions) (resp azfake.Responder[armm365securityandcompliance.PrivateLinkResourcesSecClientListByServiceResponse], errResp azfake.ErrorResponder)
}

// NewPrivateLinkResourcesSecServerTransport creates a new instance of PrivateLinkResourcesSecServerTransport with the provided implementation.
// The returned PrivateLinkResourcesSecServerTransport instance is connected to an instance of armm365securityandcompliance.PrivateLinkResourcesSecClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateLinkResourcesSecServerTransport(srv *PrivateLinkResourcesSecServer) *PrivateLinkResourcesSecServerTransport {
	return &PrivateLinkResourcesSecServerTransport{srv: srv}
}

// PrivateLinkResourcesSecServerTransport connects instances of armm365securityandcompliance.PrivateLinkResourcesSecClient to instances of PrivateLinkResourcesSecServer.
// Don't use this type directly, use NewPrivateLinkResourcesSecServerTransport instead.
type PrivateLinkResourcesSecServerTransport struct {
	srv *PrivateLinkResourcesSecServer
}

// Do implements the policy.Transporter interface for PrivateLinkResourcesSecServerTransport.
func (p *PrivateLinkResourcesSecServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PrivateLinkResourcesSecClient.Get":
		resp, err = p.dispatchGet(req)
	case "PrivateLinkResourcesSecClient.ListByService":
		resp, err = p.dispatchListByService(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PrivateLinkResourcesSecServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateLinkResources/(?P<groupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	groupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, resourceNameParam, groupNameParam, nil)
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

func (p *PrivateLinkResourcesSecServerTransport) dispatchListByService(req *http.Request) (*http.Response, error) {
	if p.srv.ListByService == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListByService not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.M365SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateLinkResources`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.ListByService(req.Context(), resourceGroupNameParam, resourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateLinkResourceListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
