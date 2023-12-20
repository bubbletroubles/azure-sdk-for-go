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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
	"net/http"
	"net/url"
	"regexp"
)

// DenyAssignmentsServer is a fake server for instances of the armauthorization.DenyAssignmentsClient type.
type DenyAssignmentsServer struct {
	// Get is the fake for method DenyAssignmentsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, scope string, denyAssignmentID string, options *armauthorization.DenyAssignmentsClientGetOptions) (resp azfake.Responder[armauthorization.DenyAssignmentsClientGetResponse], errResp azfake.ErrorResponder)

	// GetByID is the fake for method DenyAssignmentsClient.GetByID
	// HTTP status codes to indicate success: http.StatusOK
	GetByID func(ctx context.Context, denyAssignmentID string, options *armauthorization.DenyAssignmentsClientGetByIDOptions) (resp azfake.Responder[armauthorization.DenyAssignmentsClientGetByIDResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method DenyAssignmentsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armauthorization.DenyAssignmentsClientListOptions) (resp azfake.PagerResponder[armauthorization.DenyAssignmentsClientListResponse])

	// NewListForResourcePager is the fake for method DenyAssignmentsClient.NewListForResourcePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListForResourcePager func(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, options *armauthorization.DenyAssignmentsClientListForResourceOptions) (resp azfake.PagerResponder[armauthorization.DenyAssignmentsClientListForResourceResponse])

	// NewListForResourceGroupPager is the fake for method DenyAssignmentsClient.NewListForResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListForResourceGroupPager func(resourceGroupName string, options *armauthorization.DenyAssignmentsClientListForResourceGroupOptions) (resp azfake.PagerResponder[armauthorization.DenyAssignmentsClientListForResourceGroupResponse])

	// NewListForScopePager is the fake for method DenyAssignmentsClient.NewListForScopePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListForScopePager func(scope string, options *armauthorization.DenyAssignmentsClientListForScopeOptions) (resp azfake.PagerResponder[armauthorization.DenyAssignmentsClientListForScopeResponse])
}

// NewDenyAssignmentsServerTransport creates a new instance of DenyAssignmentsServerTransport with the provided implementation.
// The returned DenyAssignmentsServerTransport instance is connected to an instance of armauthorization.DenyAssignmentsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDenyAssignmentsServerTransport(srv *DenyAssignmentsServer) *DenyAssignmentsServerTransport {
	return &DenyAssignmentsServerTransport{
		srv:                          srv,
		newListPager:                 newTracker[azfake.PagerResponder[armauthorization.DenyAssignmentsClientListResponse]](),
		newListForResourcePager:      newTracker[azfake.PagerResponder[armauthorization.DenyAssignmentsClientListForResourceResponse]](),
		newListForResourceGroupPager: newTracker[azfake.PagerResponder[armauthorization.DenyAssignmentsClientListForResourceGroupResponse]](),
		newListForScopePager:         newTracker[azfake.PagerResponder[armauthorization.DenyAssignmentsClientListForScopeResponse]](),
	}
}

// DenyAssignmentsServerTransport connects instances of armauthorization.DenyAssignmentsClient to instances of DenyAssignmentsServer.
// Don't use this type directly, use NewDenyAssignmentsServerTransport instead.
type DenyAssignmentsServerTransport struct {
	srv                          *DenyAssignmentsServer
	newListPager                 *tracker[azfake.PagerResponder[armauthorization.DenyAssignmentsClientListResponse]]
	newListForResourcePager      *tracker[azfake.PagerResponder[armauthorization.DenyAssignmentsClientListForResourceResponse]]
	newListForResourceGroupPager *tracker[azfake.PagerResponder[armauthorization.DenyAssignmentsClientListForResourceGroupResponse]]
	newListForScopePager         *tracker[azfake.PagerResponder[armauthorization.DenyAssignmentsClientListForScopeResponse]]
}

// Do implements the policy.Transporter interface for DenyAssignmentsServerTransport.
func (d *DenyAssignmentsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DenyAssignmentsClient.Get":
		resp, err = d.dispatchGet(req)
	case "DenyAssignmentsClient.GetByID":
		resp, err = d.dispatchGetByID(req)
	case "DenyAssignmentsClient.NewListPager":
		resp, err = d.dispatchNewListPager(req)
	case "DenyAssignmentsClient.NewListForResourcePager":
		resp, err = d.dispatchNewListForResourcePager(req)
	case "DenyAssignmentsClient.NewListForResourceGroupPager":
		resp, err = d.dispatchNewListForResourceGroupPager(req)
	case "DenyAssignmentsClient.NewListForScopePager":
		resp, err = d.dispatchNewListForScopePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DenyAssignmentsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/denyAssignments/(?P<denyAssignmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	denyAssignmentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("denyAssignmentId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), scopeParam, denyAssignmentIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DenyAssignment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DenyAssignmentsServerTransport) dispatchGetByID(req *http.Request) (*http.Response, error) {
	if d.srv.GetByID == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByID not implemented")}
	}
	const regexStr = `/(?P<denyAssignmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	denyAssignmentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("denyAssignmentId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.GetByID(req.Context(), denyAssignmentIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DenyAssignment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DenyAssignmentsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := d.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/denyAssignments`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armauthorization.DenyAssignmentsClientListOptions
		if filterParam != nil {
			options = &armauthorization.DenyAssignmentsClientListOptions{
				Filter: filterParam,
			}
		}
		resp := d.srv.NewListPager(options)
		newListPager = &resp
		d.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armauthorization.DenyAssignmentsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		d.newListPager.remove(req)
	}
	return resp, nil
}

func (d *DenyAssignmentsServerTransport) dispatchNewListForResourcePager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListForResourcePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListForResourcePager not implemented")}
	}
	newListForResourcePager := d.newListForResourcePager.get(req)
	if newListForResourcePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/(?P<resourceProviderNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<parentResourcePath>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<resourceType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/denyAssignments`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceProviderNamespaceParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceProviderNamespace")])
		if err != nil {
			return nil, err
		}
		parentResourcePathParam, err := url.PathUnescape(matches[regex.SubexpIndex("parentResourcePath")])
		if err != nil {
			return nil, err
		}
		resourceTypeParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceType")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armauthorization.DenyAssignmentsClientListForResourceOptions
		if filterParam != nil {
			options = &armauthorization.DenyAssignmentsClientListForResourceOptions{
				Filter: filterParam,
			}
		}
		resp := d.srv.NewListForResourcePager(resourceGroupNameParam, resourceProviderNamespaceParam, parentResourcePathParam, resourceTypeParam, resourceNameParam, options)
		newListForResourcePager = &resp
		d.newListForResourcePager.add(req, newListForResourcePager)
		server.PagerResponderInjectNextLinks(newListForResourcePager, req, func(page *armauthorization.DenyAssignmentsClientListForResourceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListForResourcePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListForResourcePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListForResourcePager) {
		d.newListForResourcePager.remove(req)
	}
	return resp, nil
}

func (d *DenyAssignmentsServerTransport) dispatchNewListForResourceGroupPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListForResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListForResourceGroupPager not implemented")}
	}
	newListForResourceGroupPager := d.newListForResourceGroupPager.get(req)
	if newListForResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/denyAssignments`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armauthorization.DenyAssignmentsClientListForResourceGroupOptions
		if filterParam != nil {
			options = &armauthorization.DenyAssignmentsClientListForResourceGroupOptions{
				Filter: filterParam,
			}
		}
		resp := d.srv.NewListForResourceGroupPager(resourceGroupNameParam, options)
		newListForResourceGroupPager = &resp
		d.newListForResourceGroupPager.add(req, newListForResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListForResourceGroupPager, req, func(page *armauthorization.DenyAssignmentsClientListForResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListForResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListForResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListForResourceGroupPager) {
		d.newListForResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (d *DenyAssignmentsServerTransport) dispatchNewListForScopePager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListForScopePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListForScopePager not implemented")}
	}
	newListForScopePager := d.newListForScopePager.get(req)
	if newListForScopePager == nil {
		const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/denyAssignments`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armauthorization.DenyAssignmentsClientListForScopeOptions
		if filterParam != nil {
			options = &armauthorization.DenyAssignmentsClientListForScopeOptions{
				Filter: filterParam,
			}
		}
		resp := d.srv.NewListForScopePager(scopeParam, options)
		newListForScopePager = &resp
		d.newListForScopePager.add(req, newListForScopePager)
		server.PagerResponderInjectNextLinks(newListForScopePager, req, func(page *armauthorization.DenyAssignmentsClientListForScopeResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListForScopePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListForScopePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListForScopePager) {
		d.newListForScopePager.remove(req)
	}
	return resp, nil
}
