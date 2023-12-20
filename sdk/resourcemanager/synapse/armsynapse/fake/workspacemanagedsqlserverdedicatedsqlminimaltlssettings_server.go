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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
	"net/http"
	"net/url"
	"regexp"
)

// WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsServer is a fake server for instances of the armsynapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient type.
type WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsServer struct {
	// Get is the fake for method WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, dedicatedSQLminimalTLSSettingsName string, options *armsynapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientGetOptions) (resp azfake.Responder[armsynapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workspaceName string, options *armsynapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientListOptions) (resp azfake.PagerResponder[armsynapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientListResponse])

	// BeginUpdate is the fake for method WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, dedicatedSQLminimalTLSSettingsName armsynapse.DedicatedSQLMinimalTLSSettingsName, parameters armsynapse.DedicatedSQLminimalTLSSettings, options *armsynapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientBeginUpdateOptions) (resp azfake.PollerResponder[armsynapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewWorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsServerTransport creates a new instance of WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsServerTransport with the provided implementation.
// The returned WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsServerTransport instance is connected to an instance of armsynapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsServerTransport(srv *WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsServer) *WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsServerTransport {
	return &WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armsynapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientListResponse]](),
		beginUpdate:  newTracker[azfake.PollerResponder[armsynapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientUpdateResponse]](),
	}
}

// WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsServerTransport connects instances of armsynapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient to instances of WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsServer.
// Don't use this type directly, use NewWorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsServerTransport instead.
type WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsServerTransport struct {
	srv          *WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsServer
	newListPager *tracker[azfake.PagerResponder[armsynapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientListResponse]]
	beginUpdate  *tracker[azfake.PollerResponder[armsynapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsServerTransport.
func (w *WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient.Get":
		resp, err = w.dispatchGet(req)
	case "WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient.NewListPager":
		resp, err = w.dispatchNewListPager(req)
	case "WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient.BeginUpdate":
		resp, err = w.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Synapse/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dedicatedSQLminimalTlsSettings/(?P<dedicatedSQLminimalTlsSettingsName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	dedicatedSQLminimalTLSSettingsNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dedicatedSQLminimalTlsSettingsName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, dedicatedSQLminimalTLSSettingsNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DedicatedSQLminimalTLSSettings, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := w.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Synapse/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dedicatedSQLminimalTlsSettings`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		resp := w.srv.NewListPager(resourceGroupNameParam, workspaceNameParam, nil)
		newListPager = &resp
		w.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsynapse.WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		w.newListPager.remove(req)
	}
	return resp, nil
}

func (w *WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if w.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := w.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Synapse/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dedicatedSQLminimalTlsSettings/(?P<dedicatedSQLminimalTlsSettingsName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsynapse.DedicatedSQLminimalTLSSettings](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		dedicatedSQLminimalTLSSettingsNameParam, err := parseWithCast(matches[regex.SubexpIndex("dedicatedSQLminimalTlsSettingsName")], func(v string) (armsynapse.DedicatedSQLMinimalTLSSettingsName, error) {
			p, unescapeErr := url.PathUnescape(v)
			if unescapeErr != nil {
				return "", unescapeErr
			}
			return armsynapse.DedicatedSQLMinimalTLSSettingsName(p), nil
		})
		if err != nil {
			return nil, err
		}
		respr, errRespr := w.srv.BeginUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, dedicatedSQLminimalTLSSettingsNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		w.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		w.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		w.beginUpdate.remove(req)
	}

	return resp, nil
}
