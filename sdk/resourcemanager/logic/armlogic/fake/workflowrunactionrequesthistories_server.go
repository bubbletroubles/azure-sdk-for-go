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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
	"net/http"
	"net/url"
	"regexp"
)

// WorkflowRunActionRequestHistoriesServer is a fake server for instances of the armlogic.WorkflowRunActionRequestHistoriesClient type.
type WorkflowRunActionRequestHistoriesServer struct {
	// Get is the fake for method WorkflowRunActionRequestHistoriesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workflowName string, runName string, actionName string, requestHistoryName string, options *armlogic.WorkflowRunActionRequestHistoriesClientGetOptions) (resp azfake.Responder[armlogic.WorkflowRunActionRequestHistoriesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method WorkflowRunActionRequestHistoriesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workflowName string, runName string, actionName string, options *armlogic.WorkflowRunActionRequestHistoriesClientListOptions) (resp azfake.PagerResponder[armlogic.WorkflowRunActionRequestHistoriesClientListResponse])
}

// NewWorkflowRunActionRequestHistoriesServerTransport creates a new instance of WorkflowRunActionRequestHistoriesServerTransport with the provided implementation.
// The returned WorkflowRunActionRequestHistoriesServerTransport instance is connected to an instance of armlogic.WorkflowRunActionRequestHistoriesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWorkflowRunActionRequestHistoriesServerTransport(srv *WorkflowRunActionRequestHistoriesServer) *WorkflowRunActionRequestHistoriesServerTransport {
	return &WorkflowRunActionRequestHistoriesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armlogic.WorkflowRunActionRequestHistoriesClientListResponse]](),
	}
}

// WorkflowRunActionRequestHistoriesServerTransport connects instances of armlogic.WorkflowRunActionRequestHistoriesClient to instances of WorkflowRunActionRequestHistoriesServer.
// Don't use this type directly, use NewWorkflowRunActionRequestHistoriesServerTransport instead.
type WorkflowRunActionRequestHistoriesServerTransport struct {
	srv          *WorkflowRunActionRequestHistoriesServer
	newListPager *tracker[azfake.PagerResponder[armlogic.WorkflowRunActionRequestHistoriesClientListResponse]]
}

// Do implements the policy.Transporter interface for WorkflowRunActionRequestHistoriesServerTransport.
func (w *WorkflowRunActionRequestHistoriesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "WorkflowRunActionRequestHistoriesClient.Get":
		resp, err = w.dispatchGet(req)
	case "WorkflowRunActionRequestHistoriesClient.NewListPager":
		resp, err = w.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WorkflowRunActionRequestHistoriesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runs/(?P<runName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/actions/(?P<actionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/requestHistories/(?P<requestHistoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workflowNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workflowName")])
	if err != nil {
		return nil, err
	}
	runNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("runName")])
	if err != nil {
		return nil, err
	}
	actionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("actionName")])
	if err != nil {
		return nil, err
	}
	requestHistoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("requestHistoryName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Get(req.Context(), resourceGroupNameParam, workflowNameParam, runNameParam, actionNameParam, requestHistoryNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RequestHistory, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkflowRunActionRequestHistoriesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := w.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/workflows/(?P<workflowName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runs/(?P<runName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/actions/(?P<actionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/requestHistories`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workflowNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workflowName")])
		if err != nil {
			return nil, err
		}
		runNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("runName")])
		if err != nil {
			return nil, err
		}
		actionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("actionName")])
		if err != nil {
			return nil, err
		}
		resp := w.srv.NewListPager(resourceGroupNameParam, workflowNameParam, runNameParam, actionNameParam, nil)
		newListPager = &resp
		w.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armlogic.WorkflowRunActionRequestHistoriesClientListResponse, createLink func() string) {
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
