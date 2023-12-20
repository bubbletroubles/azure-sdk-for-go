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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice"
	"net/http"
	"net/url"
	"regexp"
)

// Server is a fake server for instances of the armhybridcontainerservice.Client type.
type Server struct {
	// BeginDeleteKubernetesVersions is the fake for method Client.BeginDeleteKubernetesVersions
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDeleteKubernetesVersions func(ctx context.Context, customLocationResourceURI string, options *armhybridcontainerservice.ClientBeginDeleteKubernetesVersionsOptions) (resp azfake.PollerResponder[armhybridcontainerservice.ClientDeleteKubernetesVersionsResponse], errResp azfake.ErrorResponder)

	// BeginDeleteVMSKUs is the fake for method Client.BeginDeleteVMSKUs
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDeleteVMSKUs func(ctx context.Context, customLocationResourceURI string, options *armhybridcontainerservice.ClientBeginDeleteVMSKUsOptions) (resp azfake.PollerResponder[armhybridcontainerservice.ClientDeleteVMSKUsResponse], errResp azfake.ErrorResponder)

	// GetKubernetesVersions is the fake for method Client.GetKubernetesVersions
	// HTTP status codes to indicate success: http.StatusOK
	GetKubernetesVersions func(ctx context.Context, customLocationResourceURI string, options *armhybridcontainerservice.ClientGetKubernetesVersionsOptions) (resp azfake.Responder[armhybridcontainerservice.ClientGetKubernetesVersionsResponse], errResp azfake.ErrorResponder)

	// GetVMSKUs is the fake for method Client.GetVMSKUs
	// HTTP status codes to indicate success: http.StatusOK
	GetVMSKUs func(ctx context.Context, customLocationResourceURI string, options *armhybridcontainerservice.ClientGetVMSKUsOptions) (resp azfake.Responder[armhybridcontainerservice.ClientGetVMSKUsResponse], errResp azfake.ErrorResponder)

	// BeginPutKubernetesVersions is the fake for method Client.BeginPutKubernetesVersions
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginPutKubernetesVersions func(ctx context.Context, customLocationResourceURI string, kubernetesVersions armhybridcontainerservice.KubernetesVersionProfile, options *armhybridcontainerservice.ClientBeginPutKubernetesVersionsOptions) (resp azfake.PollerResponder[armhybridcontainerservice.ClientPutKubernetesVersionsResponse], errResp azfake.ErrorResponder)

	// BeginPutVMSKUs is the fake for method Client.BeginPutVMSKUs
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginPutVMSKUs func(ctx context.Context, customLocationResourceURI string, skus armhybridcontainerservice.VMSKUProfile, options *armhybridcontainerservice.ClientBeginPutVMSKUsOptions) (resp azfake.PollerResponder[armhybridcontainerservice.ClientPutVMSKUsResponse], errResp azfake.ErrorResponder)
}

// NewServerTransport creates a new instance of ServerTransport with the provided implementation.
// The returned ServerTransport instance is connected to an instance of armhybridcontainerservice.Client via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerTransport(srv *Server) *ServerTransport {
	return &ServerTransport{
		srv:                           srv,
		beginDeleteKubernetesVersions: newTracker[azfake.PollerResponder[armhybridcontainerservice.ClientDeleteKubernetesVersionsResponse]](),
		beginDeleteVMSKUs:             newTracker[azfake.PollerResponder[armhybridcontainerservice.ClientDeleteVMSKUsResponse]](),
		beginPutKubernetesVersions:    newTracker[azfake.PollerResponder[armhybridcontainerservice.ClientPutKubernetesVersionsResponse]](),
		beginPutVMSKUs:                newTracker[azfake.PollerResponder[armhybridcontainerservice.ClientPutVMSKUsResponse]](),
	}
}

// ServerTransport connects instances of armhybridcontainerservice.Client to instances of Server.
// Don't use this type directly, use NewServerTransport instead.
type ServerTransport struct {
	srv                           *Server
	beginDeleteKubernetesVersions *tracker[azfake.PollerResponder[armhybridcontainerservice.ClientDeleteKubernetesVersionsResponse]]
	beginDeleteVMSKUs             *tracker[azfake.PollerResponder[armhybridcontainerservice.ClientDeleteVMSKUsResponse]]
	beginPutKubernetesVersions    *tracker[azfake.PollerResponder[armhybridcontainerservice.ClientPutKubernetesVersionsResponse]]
	beginPutVMSKUs                *tracker[azfake.PollerResponder[armhybridcontainerservice.ClientPutVMSKUsResponse]]
}

// Do implements the policy.Transporter interface for ServerTransport.
func (s *ServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "Client.BeginDeleteKubernetesVersions":
		resp, err = s.dispatchBeginDeleteKubernetesVersions(req)
	case "Client.BeginDeleteVMSKUs":
		resp, err = s.dispatchBeginDeleteVMSKUs(req)
	case "Client.GetKubernetesVersions":
		resp, err = s.dispatchGetKubernetesVersions(req)
	case "Client.GetVMSKUs":
		resp, err = s.dispatchGetVMSKUs(req)
	case "Client.BeginPutKubernetesVersions":
		resp, err = s.dispatchBeginPutKubernetesVersions(req)
	case "Client.BeginPutVMSKUs":
		resp, err = s.dispatchBeginPutVMSKUs(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ServerTransport) dispatchBeginDeleteKubernetesVersions(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDeleteKubernetesVersions == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDeleteKubernetesVersions not implemented")}
	}
	beginDeleteKubernetesVersions := s.beginDeleteKubernetesVersions.get(req)
	if beginDeleteKubernetesVersions == nil {
		const regexStr = `/(?P<customLocationResourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridContainerService/kubernetesVersions/default`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		customLocationResourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("customLocationResourceUri")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginDeleteKubernetesVersions(req.Context(), customLocationResourceURIParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDeleteKubernetesVersions = &respr
		s.beginDeleteKubernetesVersions.add(req, beginDeleteKubernetesVersions)
	}

	resp, err := server.PollerResponderNext(beginDeleteKubernetesVersions, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		s.beginDeleteKubernetesVersions.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDeleteKubernetesVersions) {
		s.beginDeleteKubernetesVersions.remove(req)
	}

	return resp, nil
}

func (s *ServerTransport) dispatchBeginDeleteVMSKUs(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDeleteVMSKUs == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDeleteVMSKUs not implemented")}
	}
	beginDeleteVMSKUs := s.beginDeleteVMSKUs.get(req)
	if beginDeleteVMSKUs == nil {
		const regexStr = `/(?P<customLocationResourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridContainerService/skus/default`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		customLocationResourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("customLocationResourceUri")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginDeleteVMSKUs(req.Context(), customLocationResourceURIParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDeleteVMSKUs = &respr
		s.beginDeleteVMSKUs.add(req, beginDeleteVMSKUs)
	}

	resp, err := server.PollerResponderNext(beginDeleteVMSKUs, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		s.beginDeleteVMSKUs.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDeleteVMSKUs) {
		s.beginDeleteVMSKUs.remove(req)
	}

	return resp, nil
}

func (s *ServerTransport) dispatchGetKubernetesVersions(req *http.Request) (*http.Response, error) {
	if s.srv.GetKubernetesVersions == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetKubernetesVersions not implemented")}
	}
	const regexStr = `/(?P<customLocationResourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridContainerService/kubernetesVersions/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	customLocationResourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("customLocationResourceUri")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.GetKubernetesVersions(req.Context(), customLocationResourceURIParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).KubernetesVersionProfile, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerTransport) dispatchGetVMSKUs(req *http.Request) (*http.Response, error) {
	if s.srv.GetVMSKUs == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetVMSKUs not implemented")}
	}
	const regexStr = `/(?P<customLocationResourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridContainerService/skus/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	customLocationResourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("customLocationResourceUri")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.GetVMSKUs(req.Context(), customLocationResourceURIParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VMSKUProfile, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerTransport) dispatchBeginPutKubernetesVersions(req *http.Request) (*http.Response, error) {
	if s.srv.BeginPutKubernetesVersions == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPutKubernetesVersions not implemented")}
	}
	beginPutKubernetesVersions := s.beginPutKubernetesVersions.get(req)
	if beginPutKubernetesVersions == nil {
		const regexStr = `/(?P<customLocationResourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridContainerService/kubernetesVersions/default`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armhybridcontainerservice.KubernetesVersionProfile](req)
		if err != nil {
			return nil, err
		}
		customLocationResourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("customLocationResourceUri")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginPutKubernetesVersions(req.Context(), customLocationResourceURIParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPutKubernetesVersions = &respr
		s.beginPutKubernetesVersions.add(req, beginPutKubernetesVersions)
	}

	resp, err := server.PollerResponderNext(beginPutKubernetesVersions, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		s.beginPutKubernetesVersions.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPutKubernetesVersions) {
		s.beginPutKubernetesVersions.remove(req)
	}

	return resp, nil
}

func (s *ServerTransport) dispatchBeginPutVMSKUs(req *http.Request) (*http.Response, error) {
	if s.srv.BeginPutVMSKUs == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPutVMSKUs not implemented")}
	}
	beginPutVMSKUs := s.beginPutVMSKUs.get(req)
	if beginPutVMSKUs == nil {
		const regexStr = `/(?P<customLocationResourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridContainerService/skus/default`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armhybridcontainerservice.VMSKUProfile](req)
		if err != nil {
			return nil, err
		}
		customLocationResourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("customLocationResourceUri")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginPutVMSKUs(req.Context(), customLocationResourceURIParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPutVMSKUs = &respr
		s.beginPutVMSKUs.add(req, beginPutVMSKUs)
	}

	resp, err := server.PollerResponderNext(beginPutVMSKUs, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		s.beginPutVMSKUs.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPutVMSKUs) {
		s.beginPutVMSKUs.remove(req)
	}

	return resp, nil
}
