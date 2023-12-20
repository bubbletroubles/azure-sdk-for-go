//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
	"net/http"
	"net/url"
	"regexp"
)

// NodesServer is a fake server for instances of the armdataboxedge.NodesClient type.
type NodesServer struct {
	// NewListByDataBoxEdgeDevicePager is the fake for method NodesClient.NewListByDataBoxEdgeDevicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByDataBoxEdgeDevicePager func(deviceName string, resourceGroupName string, options *armdataboxedge.NodesClientListByDataBoxEdgeDeviceOptions) (resp azfake.PagerResponder[armdataboxedge.NodesClientListByDataBoxEdgeDeviceResponse])
}

// NewNodesServerTransport creates a new instance of NodesServerTransport with the provided implementation.
// The returned NodesServerTransport instance is connected to an instance of armdataboxedge.NodesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNodesServerTransport(srv *NodesServer) *NodesServerTransport {
	return &NodesServerTransport{
		srv:                             srv,
		newListByDataBoxEdgeDevicePager: newTracker[azfake.PagerResponder[armdataboxedge.NodesClientListByDataBoxEdgeDeviceResponse]](),
	}
}

// NodesServerTransport connects instances of armdataboxedge.NodesClient to instances of NodesServer.
// Don't use this type directly, use NewNodesServerTransport instead.
type NodesServerTransport struct {
	srv                             *NodesServer
	newListByDataBoxEdgeDevicePager *tracker[azfake.PagerResponder[armdataboxedge.NodesClientListByDataBoxEdgeDeviceResponse]]
}

// Do implements the policy.Transporter interface for NodesServerTransport.
func (n *NodesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "NodesClient.NewListByDataBoxEdgeDevicePager":
		resp, err = n.dispatchNewListByDataBoxEdgeDevicePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (n *NodesServerTransport) dispatchNewListByDataBoxEdgeDevicePager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListByDataBoxEdgeDevicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByDataBoxEdgeDevicePager not implemented")}
	}
	newListByDataBoxEdgeDevicePager := n.newListByDataBoxEdgeDevicePager.get(req)
	if newListByDataBoxEdgeDevicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/nodes`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := n.srv.NewListByDataBoxEdgeDevicePager(deviceNameParam, resourceGroupNameParam, nil)
		newListByDataBoxEdgeDevicePager = &resp
		n.newListByDataBoxEdgeDevicePager.add(req, newListByDataBoxEdgeDevicePager)
		server.PagerResponderInjectNextLinks(newListByDataBoxEdgeDevicePager, req, func(page *armdataboxedge.NodesClientListByDataBoxEdgeDeviceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByDataBoxEdgeDevicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListByDataBoxEdgeDevicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByDataBoxEdgeDevicePager) {
		n.newListByDataBoxEdgeDevicePager.remove(req)
	}
	return resp, nil
}
