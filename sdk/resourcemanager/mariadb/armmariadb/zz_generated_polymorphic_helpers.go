//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmariadb

import "encoding/json"

func unmarshalServerPropertiesForCreateClassification(rawMsg json.RawMessage) (ServerPropertiesForCreateClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ServerPropertiesForCreateClassification
	switch m["createMode"] {
	case string(CreateModeDefault):
		b = &ServerPropertiesForDefaultCreate{}
	case string(CreateModeGeoRestore):
		b = &ServerPropertiesForGeoRestore{}
	case string(CreateModePointInTimeRestore):
		b = &ServerPropertiesForRestore{}
	case string(CreateModeReplica):
		b = &ServerPropertiesForReplica{}
	default:
		b = &ServerPropertiesForCreate{}
	}
	return b, json.Unmarshal(rawMsg, b)
}