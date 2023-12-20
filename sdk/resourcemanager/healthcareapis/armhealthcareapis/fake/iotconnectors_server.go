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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis/v2"
	"net/http"
	"net/url"
	"regexp"
)

// IotConnectorsServer is a fake server for instances of the armhealthcareapis.IotConnectorsClient type.
type IotConnectorsServer struct {
	// BeginCreateOrUpdate is the fake for method IotConnectorsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, iotConnectorName string, iotConnector armhealthcareapis.IotConnector, options *armhealthcareapis.IotConnectorsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armhealthcareapis.IotConnectorsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method IotConnectorsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, iotConnectorName string, workspaceName string, options *armhealthcareapis.IotConnectorsClientBeginDeleteOptions) (resp azfake.PollerResponder[armhealthcareapis.IotConnectorsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method IotConnectorsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, iotConnectorName string, options *armhealthcareapis.IotConnectorsClientGetOptions) (resp azfake.Responder[armhealthcareapis.IotConnectorsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByWorkspacePager is the fake for method IotConnectorsClient.NewListByWorkspacePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByWorkspacePager func(resourceGroupName string, workspaceName string, options *armhealthcareapis.IotConnectorsClientListByWorkspaceOptions) (resp azfake.PagerResponder[armhealthcareapis.IotConnectorsClientListByWorkspaceResponse])

	// BeginUpdate is the fake for method IotConnectorsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, iotConnectorName string, workspaceName string, iotConnectorPatchResource armhealthcareapis.IotConnectorPatchResource, options *armhealthcareapis.IotConnectorsClientBeginUpdateOptions) (resp azfake.PollerResponder[armhealthcareapis.IotConnectorsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewIotConnectorsServerTransport creates a new instance of IotConnectorsServerTransport with the provided implementation.
// The returned IotConnectorsServerTransport instance is connected to an instance of armhealthcareapis.IotConnectorsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewIotConnectorsServerTransport(srv *IotConnectorsServer) *IotConnectorsServerTransport {
	return &IotConnectorsServerTransport{
		srv:                     srv,
		beginCreateOrUpdate:     newTracker[azfake.PollerResponder[armhealthcareapis.IotConnectorsClientCreateOrUpdateResponse]](),
		beginDelete:             newTracker[azfake.PollerResponder[armhealthcareapis.IotConnectorsClientDeleteResponse]](),
		newListByWorkspacePager: newTracker[azfake.PagerResponder[armhealthcareapis.IotConnectorsClientListByWorkspaceResponse]](),
		beginUpdate:             newTracker[azfake.PollerResponder[armhealthcareapis.IotConnectorsClientUpdateResponse]](),
	}
}

// IotConnectorsServerTransport connects instances of armhealthcareapis.IotConnectorsClient to instances of IotConnectorsServer.
// Don't use this type directly, use NewIotConnectorsServerTransport instead.
type IotConnectorsServerTransport struct {
	srv                     *IotConnectorsServer
	beginCreateOrUpdate     *tracker[azfake.PollerResponder[armhealthcareapis.IotConnectorsClientCreateOrUpdateResponse]]
	beginDelete             *tracker[azfake.PollerResponder[armhealthcareapis.IotConnectorsClientDeleteResponse]]
	newListByWorkspacePager *tracker[azfake.PagerResponder[armhealthcareapis.IotConnectorsClientListByWorkspaceResponse]]
	beginUpdate             *tracker[azfake.PollerResponder[armhealthcareapis.IotConnectorsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for IotConnectorsServerTransport.
func (i *IotConnectorsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "IotConnectorsClient.BeginCreateOrUpdate":
		resp, err = i.dispatchBeginCreateOrUpdate(req)
	case "IotConnectorsClient.BeginDelete":
		resp, err = i.dispatchBeginDelete(req)
	case "IotConnectorsClient.Get":
		resp, err = i.dispatchGet(req)
	case "IotConnectorsClient.NewListByWorkspacePager":
		resp, err = i.dispatchNewListByWorkspacePager(req)
	case "IotConnectorsClient.BeginUpdate":
		resp, err = i.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *IotConnectorsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if i.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := i.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HealthcareApis/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/iotconnectors/(?P<iotConnectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armhealthcareapis.IotConnector](req)
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
		iotConnectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("iotConnectorName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, iotConnectorNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		i.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		i.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		i.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (i *IotConnectorsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if i.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := i.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HealthcareApis/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/iotconnectors/(?P<iotConnectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		iotConnectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("iotConnectorName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginDelete(req.Context(), resourceGroupNameParam, iotConnectorNameParam, workspaceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		i.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		i.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		i.beginDelete.remove(req)
	}

	return resp, nil
}

func (i *IotConnectorsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HealthcareApis/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/iotconnectors/(?P<iotConnectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	iotConnectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("iotConnectorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, iotConnectorNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).IotConnector, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IotConnectorsServerTransport) dispatchNewListByWorkspacePager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListByWorkspacePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByWorkspacePager not implemented")}
	}
	newListByWorkspacePager := i.newListByWorkspacePager.get(req)
	if newListByWorkspacePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HealthcareApis/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/iotconnectors`
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
		resp := i.srv.NewListByWorkspacePager(resourceGroupNameParam, workspaceNameParam, nil)
		newListByWorkspacePager = &resp
		i.newListByWorkspacePager.add(req, newListByWorkspacePager)
		server.PagerResponderInjectNextLinks(newListByWorkspacePager, req, func(page *armhealthcareapis.IotConnectorsClientListByWorkspaceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByWorkspacePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListByWorkspacePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByWorkspacePager) {
		i.newListByWorkspacePager.remove(req)
	}
	return resp, nil
}

func (i *IotConnectorsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if i.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := i.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HealthcareApis/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/iotconnectors/(?P<iotConnectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armhealthcareapis.IotConnectorPatchResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		iotConnectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("iotConnectorName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginUpdate(req.Context(), resourceGroupNameParam, iotConnectorNameParam, workspaceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		i.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		i.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		i.beginUpdate.remove(req)
	}

	return resp, nil
}
