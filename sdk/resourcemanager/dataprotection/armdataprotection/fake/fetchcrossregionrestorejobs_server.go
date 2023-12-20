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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v3"
	"net/http"
	"net/url"
	"regexp"
)

// FetchCrossRegionRestoreJobsServer is a fake server for instances of the armdataprotection.FetchCrossRegionRestoreJobsClient type.
type FetchCrossRegionRestoreJobsServer struct {
	// NewListPager is the fake for method FetchCrossRegionRestoreJobsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, location string, parameters armdataprotection.CrossRegionRestoreJobsRequest, options *armdataprotection.FetchCrossRegionRestoreJobsClientListOptions) (resp azfake.PagerResponder[armdataprotection.FetchCrossRegionRestoreJobsClientListResponse])
}

// NewFetchCrossRegionRestoreJobsServerTransport creates a new instance of FetchCrossRegionRestoreJobsServerTransport with the provided implementation.
// The returned FetchCrossRegionRestoreJobsServerTransport instance is connected to an instance of armdataprotection.FetchCrossRegionRestoreJobsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFetchCrossRegionRestoreJobsServerTransport(srv *FetchCrossRegionRestoreJobsServer) *FetchCrossRegionRestoreJobsServerTransport {
	return &FetchCrossRegionRestoreJobsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armdataprotection.FetchCrossRegionRestoreJobsClientListResponse]](),
	}
}

// FetchCrossRegionRestoreJobsServerTransport connects instances of armdataprotection.FetchCrossRegionRestoreJobsClient to instances of FetchCrossRegionRestoreJobsServer.
// Don't use this type directly, use NewFetchCrossRegionRestoreJobsServerTransport instead.
type FetchCrossRegionRestoreJobsServerTransport struct {
	srv          *FetchCrossRegionRestoreJobsServer
	newListPager *tracker[azfake.PagerResponder[armdataprotection.FetchCrossRegionRestoreJobsClientListResponse]]
}

// Do implements the policy.Transporter interface for FetchCrossRegionRestoreJobsServerTransport.
func (f *FetchCrossRegionRestoreJobsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "FetchCrossRegionRestoreJobsClient.NewListPager":
		resp, err = f.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (f *FetchCrossRegionRestoreJobsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := f.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataProtection/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fetchCrossRegionRestoreJobs`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		body, err := server.UnmarshalRequestAsJSON[armdataprotection.CrossRegionRestoreJobsRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armdataprotection.FetchCrossRegionRestoreJobsClientListOptions
		if filterParam != nil {
			options = &armdataprotection.FetchCrossRegionRestoreJobsClientListOptions{
				Filter: filterParam,
			}
		}
		resp := f.srv.NewListPager(resourceGroupNameParam, locationParam, body, options)
		newListPager = &resp
		f.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armdataprotection.FetchCrossRegionRestoreJobsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		f.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		f.newListPager.remove(req)
	}
	return resp, nil
}
