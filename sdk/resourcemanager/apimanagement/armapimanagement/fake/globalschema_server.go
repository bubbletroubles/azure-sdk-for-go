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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// GlobalSchemaServer is a fake server for instances of the armapimanagement.GlobalSchemaClient type.
type GlobalSchemaServer struct {
	// BeginCreateOrUpdate is the fake for method GlobalSchemaClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, serviceName string, schemaID string, parameters armapimanagement.GlobalSchemaContract, options *armapimanagement.GlobalSchemaClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armapimanagement.GlobalSchemaClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method GlobalSchemaClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, serviceName string, schemaID string, ifMatch string, options *armapimanagement.GlobalSchemaClientDeleteOptions) (resp azfake.Responder[armapimanagement.GlobalSchemaClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method GlobalSchemaClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serviceName string, schemaID string, options *armapimanagement.GlobalSchemaClientGetOptions) (resp azfake.Responder[armapimanagement.GlobalSchemaClientGetResponse], errResp azfake.ErrorResponder)

	// GetEntityTag is the fake for method GlobalSchemaClient.GetEntityTag
	// HTTP status codes to indicate success: http.StatusOK
	GetEntityTag func(ctx context.Context, resourceGroupName string, serviceName string, schemaID string, options *armapimanagement.GlobalSchemaClientGetEntityTagOptions) (resp azfake.Responder[armapimanagement.GlobalSchemaClientGetEntityTagResponse], errResp azfake.ErrorResponder)

	// NewListByServicePager is the fake for method GlobalSchemaClient.NewListByServicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServicePager func(resourceGroupName string, serviceName string, options *armapimanagement.GlobalSchemaClientListByServiceOptions) (resp azfake.PagerResponder[armapimanagement.GlobalSchemaClientListByServiceResponse])
}

// NewGlobalSchemaServerTransport creates a new instance of GlobalSchemaServerTransport with the provided implementation.
// The returned GlobalSchemaServerTransport instance is connected to an instance of armapimanagement.GlobalSchemaClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGlobalSchemaServerTransport(srv *GlobalSchemaServer) *GlobalSchemaServerTransport {
	return &GlobalSchemaServerTransport{
		srv:                   srv,
		beginCreateOrUpdate:   newTracker[azfake.PollerResponder[armapimanagement.GlobalSchemaClientCreateOrUpdateResponse]](),
		newListByServicePager: newTracker[azfake.PagerResponder[armapimanagement.GlobalSchemaClientListByServiceResponse]](),
	}
}

// GlobalSchemaServerTransport connects instances of armapimanagement.GlobalSchemaClient to instances of GlobalSchemaServer.
// Don't use this type directly, use NewGlobalSchemaServerTransport instead.
type GlobalSchemaServerTransport struct {
	srv                   *GlobalSchemaServer
	beginCreateOrUpdate   *tracker[azfake.PollerResponder[armapimanagement.GlobalSchemaClientCreateOrUpdateResponse]]
	newListByServicePager *tracker[azfake.PagerResponder[armapimanagement.GlobalSchemaClientListByServiceResponse]]
}

// Do implements the policy.Transporter interface for GlobalSchemaServerTransport.
func (g *GlobalSchemaServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "GlobalSchemaClient.BeginCreateOrUpdate":
		resp, err = g.dispatchBeginCreateOrUpdate(req)
	case "GlobalSchemaClient.Delete":
		resp, err = g.dispatchDelete(req)
	case "GlobalSchemaClient.Get":
		resp, err = g.dispatchGet(req)
	case "GlobalSchemaClient.GetEntityTag":
		resp, err = g.dispatchGetEntityTag(req)
	case "GlobalSchemaClient.NewListByServicePager":
		resp, err = g.dispatchNewListByServicePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GlobalSchemaServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if g.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := g.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schemas/(?P<schemaId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armapimanagement.GlobalSchemaContract](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		schemaIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("schemaId")])
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		var options *armapimanagement.GlobalSchemaClientBeginCreateOrUpdateOptions
		if ifMatchParam != nil {
			options = &armapimanagement.GlobalSchemaClientBeginCreateOrUpdateOptions{
				IfMatch: ifMatchParam,
			}
		}
		respr, errRespr := g.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, serviceNameParam, schemaIDParam, body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		g.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		g.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		g.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (g *GlobalSchemaServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if g.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schemas/(?P<schemaId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	schemaIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("schemaId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Delete(req.Context(), resourceGroupNameParam, serviceNameParam, schemaIDParam, getHeaderValue(req.Header, "If-Match"), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GlobalSchemaServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if g.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schemas/(?P<schemaId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	schemaIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("schemaId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Get(req.Context(), resourceGroupNameParam, serviceNameParam, schemaIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GlobalSchemaContract, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (g *GlobalSchemaServerTransport) dispatchGetEntityTag(req *http.Request) (*http.Response, error) {
	if g.srv.GetEntityTag == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetEntityTag not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schemas/(?P<schemaId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	schemaIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("schemaId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.GetEntityTag(req.Context(), resourceGroupNameParam, serviceNameParam, schemaIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (g *GlobalSchemaServerTransport) dispatchNewListByServicePager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListByServicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServicePager not implemented")}
	}
	newListByServicePager := g.newListByServicePager.get(req)
	if newListByServicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schemas`
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
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
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
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armapimanagement.GlobalSchemaClientListByServiceOptions
		if filterParam != nil || topParam != nil || skipParam != nil {
			options = &armapimanagement.GlobalSchemaClientListByServiceOptions{
				Filter: filterParam,
				Top:    topParam,
				Skip:   skipParam,
			}
		}
		resp := g.srv.NewListByServicePager(resourceGroupNameParam, serviceNameParam, options)
		newListByServicePager = &resp
		g.newListByServicePager.add(req, newListByServicePager)
		server.PagerResponderInjectNextLinks(newListByServicePager, req, func(page *armapimanagement.GlobalSchemaClientListByServiceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		g.newListByServicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServicePager) {
		g.newListByServicePager.remove(req)
	}
	return resp, nil
}
