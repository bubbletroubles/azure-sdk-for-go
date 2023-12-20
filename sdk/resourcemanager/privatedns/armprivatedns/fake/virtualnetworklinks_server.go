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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// VirtualNetworkLinksServer is a fake server for instances of the armprivatedns.VirtualNetworkLinksClient type.
type VirtualNetworkLinksServer struct {
	// BeginCreateOrUpdate is the fake for method VirtualNetworkLinksClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, parameters armprivatedns.VirtualNetworkLink, options *armprivatedns.VirtualNetworkLinksClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armprivatedns.VirtualNetworkLinksClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VirtualNetworkLinksClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, options *armprivatedns.VirtualNetworkLinksClientBeginDeleteOptions) (resp azfake.PollerResponder[armprivatedns.VirtualNetworkLinksClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VirtualNetworkLinksClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, options *armprivatedns.VirtualNetworkLinksClientGetOptions) (resp azfake.Responder[armprivatedns.VirtualNetworkLinksClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method VirtualNetworkLinksClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, privateZoneName string, options *armprivatedns.VirtualNetworkLinksClientListOptions) (resp azfake.PagerResponder[armprivatedns.VirtualNetworkLinksClientListResponse])

	// BeginUpdate is the fake for method VirtualNetworkLinksClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, parameters armprivatedns.VirtualNetworkLink, options *armprivatedns.VirtualNetworkLinksClientBeginUpdateOptions) (resp azfake.PollerResponder[armprivatedns.VirtualNetworkLinksClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewVirtualNetworkLinksServerTransport creates a new instance of VirtualNetworkLinksServerTransport with the provided implementation.
// The returned VirtualNetworkLinksServerTransport instance is connected to an instance of armprivatedns.VirtualNetworkLinksClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualNetworkLinksServerTransport(srv *VirtualNetworkLinksServer) *VirtualNetworkLinksServerTransport {
	return &VirtualNetworkLinksServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armprivatedns.VirtualNetworkLinksClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armprivatedns.VirtualNetworkLinksClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armprivatedns.VirtualNetworkLinksClientListResponse]](),
		beginUpdate:         newTracker[azfake.PollerResponder[armprivatedns.VirtualNetworkLinksClientUpdateResponse]](),
	}
}

// VirtualNetworkLinksServerTransport connects instances of armprivatedns.VirtualNetworkLinksClient to instances of VirtualNetworkLinksServer.
// Don't use this type directly, use NewVirtualNetworkLinksServerTransport instead.
type VirtualNetworkLinksServerTransport struct {
	srv                 *VirtualNetworkLinksServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armprivatedns.VirtualNetworkLinksClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armprivatedns.VirtualNetworkLinksClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armprivatedns.VirtualNetworkLinksClientListResponse]]
	beginUpdate         *tracker[azfake.PollerResponder[armprivatedns.VirtualNetworkLinksClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for VirtualNetworkLinksServerTransport.
func (v *VirtualNetworkLinksServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualNetworkLinksClient.BeginCreateOrUpdate":
		resp, err = v.dispatchBeginCreateOrUpdate(req)
	case "VirtualNetworkLinksClient.BeginDelete":
		resp, err = v.dispatchBeginDelete(req)
	case "VirtualNetworkLinksClient.Get":
		resp, err = v.dispatchGet(req)
	case "VirtualNetworkLinksClient.NewListPager":
		resp, err = v.dispatchNewListPager(req)
	case "VirtualNetworkLinksClient.BeginUpdate":
		resp, err = v.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualNetworkLinksServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := v.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/privateDnsZones/(?P<privateZoneName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualNetworkLinks/(?P<virtualNetworkLinkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armprivatedns.VirtualNetworkLink](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateZoneNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateZoneName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkLinkNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkLinkName")])
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
		var options *armprivatedns.VirtualNetworkLinksClientBeginCreateOrUpdateOptions
		if ifMatchParam != nil || ifNoneMatchParam != nil {
			options = &armprivatedns.VirtualNetworkLinksClientBeginCreateOrUpdateOptions{
				IfMatch:     ifMatchParam,
				IfNoneMatch: ifNoneMatchParam,
			}
		}
		respr, errRespr := v.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, privateZoneNameParam, virtualNetworkLinkNameParam, body, options)
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

func (v *VirtualNetworkLinksServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := v.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/privateDnsZones/(?P<privateZoneName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualNetworkLinks/(?P<virtualNetworkLinkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateZoneNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateZoneName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkLinkNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkLinkName")])
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		var options *armprivatedns.VirtualNetworkLinksClientBeginDeleteOptions
		if ifMatchParam != nil {
			options = &armprivatedns.VirtualNetworkLinksClientBeginDeleteOptions{
				IfMatch: ifMatchParam,
			}
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceGroupNameParam, privateZoneNameParam, virtualNetworkLinkNameParam, options)
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

func (v *VirtualNetworkLinksServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/privateDnsZones/(?P<privateZoneName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualNetworkLinks/(?P<virtualNetworkLinkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	privateZoneNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateZoneName")])
	if err != nil {
		return nil, err
	}
	virtualNetworkLinkNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkLinkName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameParam, privateZoneNameParam, virtualNetworkLinkNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualNetworkLink, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualNetworkLinksServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := v.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/privateDnsZones/(?P<privateZoneName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualNetworkLinks`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateZoneNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateZoneName")])
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armprivatedns.VirtualNetworkLinksClientListOptions
		if topParam != nil {
			options = &armprivatedns.VirtualNetworkLinksClientListOptions{
				Top: topParam,
			}
		}
		resp := v.srv.NewListPager(resourceGroupNameParam, privateZoneNameParam, options)
		newListPager = &resp
		v.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armprivatedns.VirtualNetworkLinksClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		v.newListPager.remove(req)
	}
	return resp, nil
}

func (v *VirtualNetworkLinksServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := v.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/privateDnsZones/(?P<privateZoneName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualNetworkLinks/(?P<virtualNetworkLinkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armprivatedns.VirtualNetworkLink](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateZoneNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateZoneName")])
		if err != nil {
			return nil, err
		}
		virtualNetworkLinkNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualNetworkLinkName")])
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		var options *armprivatedns.VirtualNetworkLinksClientBeginUpdateOptions
		if ifMatchParam != nil {
			options = &armprivatedns.VirtualNetworkLinksClientBeginUpdateOptions{
				IfMatch: ifMatchParam,
			}
		}
		respr, errRespr := v.srv.BeginUpdate(req.Context(), resourceGroupNameParam, privateZoneNameParam, virtualNetworkLinkNameParam, body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		v.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		v.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		v.beginUpdate.remove(req)
	}

	return resp, nil
}
