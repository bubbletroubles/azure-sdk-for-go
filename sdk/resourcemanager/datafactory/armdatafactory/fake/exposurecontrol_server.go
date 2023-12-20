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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v4"
	"net/http"
	"net/url"
	"regexp"
)

// ExposureControlServer is a fake server for instances of the armdatafactory.ExposureControlClient type.
type ExposureControlServer struct {
	// GetFeatureValue is the fake for method ExposureControlClient.GetFeatureValue
	// HTTP status codes to indicate success: http.StatusOK
	GetFeatureValue func(ctx context.Context, locationID string, exposureControlRequest armdatafactory.ExposureControlRequest, options *armdatafactory.ExposureControlClientGetFeatureValueOptions) (resp azfake.Responder[armdatafactory.ExposureControlClientGetFeatureValueResponse], errResp azfake.ErrorResponder)

	// GetFeatureValueByFactory is the fake for method ExposureControlClient.GetFeatureValueByFactory
	// HTTP status codes to indicate success: http.StatusOK
	GetFeatureValueByFactory func(ctx context.Context, resourceGroupName string, factoryName string, exposureControlRequest armdatafactory.ExposureControlRequest, options *armdatafactory.ExposureControlClientGetFeatureValueByFactoryOptions) (resp azfake.Responder[armdatafactory.ExposureControlClientGetFeatureValueByFactoryResponse], errResp azfake.ErrorResponder)

	// QueryFeatureValuesByFactory is the fake for method ExposureControlClient.QueryFeatureValuesByFactory
	// HTTP status codes to indicate success: http.StatusOK
	QueryFeatureValuesByFactory func(ctx context.Context, resourceGroupName string, factoryName string, exposureControlBatchRequest armdatafactory.ExposureControlBatchRequest, options *armdatafactory.ExposureControlClientQueryFeatureValuesByFactoryOptions) (resp azfake.Responder[armdatafactory.ExposureControlClientQueryFeatureValuesByFactoryResponse], errResp azfake.ErrorResponder)
}

// NewExposureControlServerTransport creates a new instance of ExposureControlServerTransport with the provided implementation.
// The returned ExposureControlServerTransport instance is connected to an instance of armdatafactory.ExposureControlClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewExposureControlServerTransport(srv *ExposureControlServer) *ExposureControlServerTransport {
	return &ExposureControlServerTransport{srv: srv}
}

// ExposureControlServerTransport connects instances of armdatafactory.ExposureControlClient to instances of ExposureControlServer.
// Don't use this type directly, use NewExposureControlServerTransport instead.
type ExposureControlServerTransport struct {
	srv *ExposureControlServer
}

// Do implements the policy.Transporter interface for ExposureControlServerTransport.
func (e *ExposureControlServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ExposureControlClient.GetFeatureValue":
		resp, err = e.dispatchGetFeatureValue(req)
	case "ExposureControlClient.GetFeatureValueByFactory":
		resp, err = e.dispatchGetFeatureValueByFactory(req)
	case "ExposureControlClient.QueryFeatureValuesByFactory":
		resp, err = e.dispatchQueryFeatureValuesByFactory(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *ExposureControlServerTransport) dispatchGetFeatureValue(req *http.Request) (*http.Response, error) {
	if e.srv.GetFeatureValue == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetFeatureValue not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/locations/(?P<locationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getFeatureValue`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdatafactory.ExposureControlRequest](req)
	if err != nil {
		return nil, err
	}
	locationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.GetFeatureValue(req.Context(), locationIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExposureControlResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExposureControlServerTransport) dispatchGetFeatureValueByFactory(req *http.Request) (*http.Response, error) {
	if e.srv.GetFeatureValueByFactory == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetFeatureValueByFactory not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getFeatureValue`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdatafactory.ExposureControlRequest](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.GetFeatureValueByFactory(req.Context(), resourceGroupNameParam, factoryNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExposureControlResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExposureControlServerTransport) dispatchQueryFeatureValuesByFactory(req *http.Request) (*http.Response, error) {
	if e.srv.QueryFeatureValuesByFactory == nil {
		return nil, &nonRetriableError{errors.New("fake for method QueryFeatureValuesByFactory not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/queryFeaturesValue`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdatafactory.ExposureControlBatchRequest](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.QueryFeatureValuesByFactory(req.Context(), resourceGroupNameParam, factoryNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExposureControlBatchResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
