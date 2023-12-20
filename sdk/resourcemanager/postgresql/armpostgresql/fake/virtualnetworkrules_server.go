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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
	"net/http"
	"net/url"
	"regexp"
)

// VirtualNetworkRulesServer is a fake server for instances of the armpostgresql.VirtualNetworkRulesClient type.
type VirtualNetworkRulesServer struct {
	// BeginCreateOrUpdate is the fake for method VirtualNetworkRulesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, parameters armpostgresql.VirtualNetworkRule, options *armpostgresql.VirtualNetworkRulesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armpostgresql.VirtualNetworkRulesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VirtualNetworkRulesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, options *armpostgresql.VirtualNetworkRulesClientBeginDeleteOptions) (resp azfake.PollerResponder[armpostgresql.VirtualNetworkRulesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VirtualNetworkRulesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serverName string, virtualNetworkRuleName string, options *armpostgresql.VirtualNetworkRulesClientGetOptions) (resp azfake.Responder[armpostgresql.VirtualNetworkRulesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByServerPager is the fake for method VirtualNetworkRulesClient.NewListByServerPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServerPager func(resourceGroupName string, serverName string, options *armpostgresql.VirtualNetworkRulesClientListByServerOptions) (resp azfake.PagerResponder[armpostgresql.VirtualNetworkRulesClientListByServerResponse])
}

// NewVirtualNetworkRulesServerTransport creates a new instance of VirtualNetworkRulesServerTransport with the provided implementation.
// The returned VirtualNetworkRulesServerTransport instance is connected to an instance of armpostgresql.VirtualNetworkRulesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualNetworkRulesServerTransport(srv *VirtualNetworkRulesServer) *VirtualNetworkRulesServerTransport {
	return &VirtualNetworkRulesServerTransport{
		srv:                  srv,
		beginCreateOrUpdate:  newTracker[azfake.PollerResponder[armpostgresql.VirtualNetworkRulesClientCreateOrUpdateResponse]](),
		beginDelete:          newTracker[azfake.PollerResponder[armpostgresql.VirtualNetworkRulesClientDeleteResponse]](),
		newListByServerPager: newTracker[azfake.PagerResponder[armpostgresql.VirtualNetworkRulesClientListByServerResponse]](),
	}
}

// VirtualNetworkRulesServerTransport connects instances of armpostgresql.VirtualNetworkRulesClient to instances of VirtualNetworkRulesServer.
// Don't use this type directly, use NewVirtualNetworkRulesServerTransport instead.
type VirtualNetworkRulesServerTransport struct {
	srv                  *VirtualNetworkRulesServer
	beginCreateOrUpdate  *tracker[azfake.PollerResponder[armpostgresql.VirtualNetworkRulesClientCreateOrUpdateResponse]]
	beginDelete          *tracker[azfake.PollerResponder[armpostgresql.VirtualNetworkRulesClientDeleteResponse]]
	newListByServerPager *tracker[azfake.PagerResponder[armpostgresql.VirtualNetworkRulesClientListByServerResponse]]
}

// Do implements the policy.Transporter interface for VirtualNetworkRulesServerTransport.
func (v *VirtualNetworkRulesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualNetworkRulesClient.BeginCreateOrUpdate":
		resp, err = v.dispatchBeginCreateOrUpdate(req)
	case "VirtualNetworkRulesClient.BeginDelete":
		resp, err = v.dispatchBeginDelete(req)
	case "VirtualNetworkRulesClient.Get":
		resp, err = v.dispatchGet(req)
	case "VirtualNetworkRulesClient.NewListByServerPager":
		resp, err = v.dispatchNewListByServerPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualNetworkRulesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := v.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualNetworkRules/(?P<virtualNetworkRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armpostgresql.VirtualNetworkRule](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkRuleName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, serverNameParam, virtualNetworkRuleNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		v.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		v.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		v.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (v *VirtualNetworkRulesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := v.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualNetworkRules/(?P<virtualNetworkRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkRuleName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceGroupNameParam, serverNameParam, virtualNetworkRuleNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		v.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		v.beginDelete.remove(req)
	}

	return resp, nil
}

func (v *VirtualNetworkRulesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualNetworkRules/(?P<virtualNetworkRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	virtualNetworkRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameParam, serverNameParam, virtualNetworkRuleNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualNetworkRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualNetworkRulesServerTransport) dispatchNewListByServerPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListByServerPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServerPager not implemented")}
	}
	newListByServerPager := v.newListByServerPager.get(req)
	if newListByServerPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualNetworkRules`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListByServerPager(resourceGroupNameParam, serverNameParam, nil)
		newListByServerPager = &resp
		v.newListByServerPager.add(req, newListByServerPager)
		server.PagerResponderInjectNextLinks(newListByServerPager, req, func(page *armpostgresql.VirtualNetworkRulesClientListByServerResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServerPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListByServerPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServerPager) {
		v.newListByServerPager.remove(req)
	}
	return resp, nil
}
