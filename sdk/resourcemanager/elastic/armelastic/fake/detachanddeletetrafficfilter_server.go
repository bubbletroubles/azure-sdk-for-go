//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic"
	"net/http"
	"net/url"
	"regexp"
)

// DetachAndDeleteTrafficFilterServer is a fake server for instances of the armelastic.DetachAndDeleteTrafficFilterClient type.
type DetachAndDeleteTrafficFilterServer struct {
	// Delete is the fake for method DetachAndDeleteTrafficFilterClient.Delete
	// HTTP status codes to indicate success: http.StatusOK
	Delete func(ctx context.Context, resourceGroupName string, monitorName string, options *armelastic.DetachAndDeleteTrafficFilterClientDeleteOptions) (resp azfake.Responder[armelastic.DetachAndDeleteTrafficFilterClientDeleteResponse], errResp azfake.ErrorResponder)
}

// NewDetachAndDeleteTrafficFilterServerTransport creates a new instance of DetachAndDeleteTrafficFilterServerTransport with the provided implementation.
// The returned DetachAndDeleteTrafficFilterServerTransport instance is connected to an instance of armelastic.DetachAndDeleteTrafficFilterClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDetachAndDeleteTrafficFilterServerTransport(srv *DetachAndDeleteTrafficFilterServer) *DetachAndDeleteTrafficFilterServerTransport {
	return &DetachAndDeleteTrafficFilterServerTransport{srv: srv}
}

// DetachAndDeleteTrafficFilterServerTransport connects instances of armelastic.DetachAndDeleteTrafficFilterClient to instances of DetachAndDeleteTrafficFilterServer.
// Don't use this type directly, use NewDetachAndDeleteTrafficFilterServerTransport instead.
type DetachAndDeleteTrafficFilterServerTransport struct {
	srv *DetachAndDeleteTrafficFilterServer
}

// Do implements the policy.Transporter interface for DetachAndDeleteTrafficFilterServerTransport.
func (d *DetachAndDeleteTrafficFilterServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DetachAndDeleteTrafficFilterClient.Delete":
		resp, err = d.dispatchDelete(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DetachAndDeleteTrafficFilterServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if d.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Elastic/monitors/(?P<monitorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/detachAndDeleteTrafficFilter`
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
	monitorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("monitorName")])
	if err != nil {
		return nil, err
	}
	rulesetIDUnescaped, err := url.QueryUnescape(qp.Get("rulesetId"))
	if err != nil {
		return nil, err
	}
	rulesetIDParam := getOptional(rulesetIDUnescaped)
	var options *armelastic.DetachAndDeleteTrafficFilterClientDeleteOptions
	if rulesetIDParam != nil {
		options = &armelastic.DetachAndDeleteTrafficFilterClientDeleteOptions{
			RulesetID: rulesetIDParam,
		}
	}
	respr, errRespr := d.srv.Delete(req.Context(), resourceGroupNameParam, monitorNameParam, options)
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
	return resp, nil
}
