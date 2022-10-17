//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsupport

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type CommunicationDetailsProperties.
func (c CommunicationDetailsProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "body", c.Body)
	populate(objectMap, "communicationDirection", c.CommunicationDirection)
	populate(objectMap, "communicationType", c.CommunicationType)
	populateTimeRFC3339(objectMap, "createdDate", c.CreatedDate)
	populate(objectMap, "sender", c.Sender)
	populate(objectMap, "subject", c.Subject)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CommunicationDetailsProperties.
func (c *CommunicationDetailsProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "body":
			err = unpopulate(val, "Body", &c.Body)
			delete(rawMsg, key)
		case "communicationDirection":
			err = unpopulate(val, "CommunicationDirection", &c.CommunicationDirection)
			delete(rawMsg, key)
		case "communicationType":
			err = unpopulate(val, "CommunicationType", &c.CommunicationType)
			delete(rawMsg, key)
		case "createdDate":
			err = unpopulateTimeRFC3339(val, "CreatedDate", &c.CreatedDate)
			delete(rawMsg, key)
		case "sender":
			err = unpopulate(val, "Sender", &c.Sender)
			delete(rawMsg, key)
		case "subject":
			err = unpopulate(val, "Subject", &c.Subject)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ContactProfile.
func (c ContactProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalEmailAddresses", c.AdditionalEmailAddresses)
	populate(objectMap, "country", c.Country)
	populate(objectMap, "firstName", c.FirstName)
	populate(objectMap, "lastName", c.LastName)
	populate(objectMap, "phoneNumber", c.PhoneNumber)
	populate(objectMap, "preferredContactMethod", c.PreferredContactMethod)
	populate(objectMap, "preferredSupportLanguage", c.PreferredSupportLanguage)
	populate(objectMap, "preferredTimeZone", c.PreferredTimeZone)
	populate(objectMap, "primaryEmailAddress", c.PrimaryEmailAddress)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type QuotaTicketDetails.
func (q QuotaTicketDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "quotaChangeRequestSubType", q.QuotaChangeRequestSubType)
	populate(objectMap, "quotaChangeRequestVersion", q.QuotaChangeRequestVersion)
	populate(objectMap, "quotaChangeRequests", q.QuotaChangeRequests)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServiceLevelAgreement.
func (s ServiceLevelAgreement) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "expirationTime", s.ExpirationTime)
	populate(objectMap, "slaMinutes", s.SLAMinutes)
	populateTimeRFC3339(objectMap, "startTime", s.StartTime)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ServiceLevelAgreement.
func (s *ServiceLevelAgreement) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "expirationTime":
			err = unpopulateTimeRFC3339(val, "ExpirationTime", &s.ExpirationTime)
			delete(rawMsg, key)
		case "slaMinutes":
			err = unpopulate(val, "SLAMinutes", &s.SLAMinutes)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, "StartTime", &s.StartTime)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TicketDetailsProperties.
func (t TicketDetailsProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "contactDetails", t.ContactDetails)
	populateTimeRFC3339(objectMap, "createdDate", t.CreatedDate)
	populate(objectMap, "description", t.Description)
	populate(objectMap, "enrollmentId", t.EnrollmentID)
	populateTimeRFC3339(objectMap, "modifiedDate", t.ModifiedDate)
	populate(objectMap, "problemClassificationDisplayName", t.ProblemClassificationDisplayName)
	populate(objectMap, "problemClassificationId", t.ProblemClassificationID)
	populateTimeRFC3339(objectMap, "problemStartTime", t.ProblemStartTime)
	populate(objectMap, "quotaTicketDetails", t.QuotaTicketDetails)
	populate(objectMap, "require24X7Response", t.Require24X7Response)
	populate(objectMap, "serviceDisplayName", t.ServiceDisplayName)
	populate(objectMap, "serviceId", t.ServiceID)
	populate(objectMap, "serviceLevelAgreement", t.ServiceLevelAgreement)
	populate(objectMap, "severity", t.Severity)
	populate(objectMap, "status", t.Status)
	populate(objectMap, "supportEngineer", t.SupportEngineer)
	populate(objectMap, "supportPlanType", t.SupportPlanType)
	populate(objectMap, "supportTicketId", t.SupportTicketID)
	populate(objectMap, "technicalTicketDetails", t.TechnicalTicketDetails)
	populate(objectMap, "title", t.Title)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TicketDetailsProperties.
func (t *TicketDetailsProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", t, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "contactDetails":
			err = unpopulate(val, "ContactDetails", &t.ContactDetails)
			delete(rawMsg, key)
		case "createdDate":
			err = unpopulateTimeRFC3339(val, "CreatedDate", &t.CreatedDate)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &t.Description)
			delete(rawMsg, key)
		case "enrollmentId":
			err = unpopulate(val, "EnrollmentID", &t.EnrollmentID)
			delete(rawMsg, key)
		case "modifiedDate":
			err = unpopulateTimeRFC3339(val, "ModifiedDate", &t.ModifiedDate)
			delete(rawMsg, key)
		case "problemClassificationDisplayName":
			err = unpopulate(val, "ProblemClassificationDisplayName", &t.ProblemClassificationDisplayName)
			delete(rawMsg, key)
		case "problemClassificationId":
			err = unpopulate(val, "ProblemClassificationID", &t.ProblemClassificationID)
			delete(rawMsg, key)
		case "problemStartTime":
			err = unpopulateTimeRFC3339(val, "ProblemStartTime", &t.ProblemStartTime)
			delete(rawMsg, key)
		case "quotaTicketDetails":
			err = unpopulate(val, "QuotaTicketDetails", &t.QuotaTicketDetails)
			delete(rawMsg, key)
		case "require24X7Response":
			err = unpopulate(val, "Require24X7Response", &t.Require24X7Response)
			delete(rawMsg, key)
		case "serviceDisplayName":
			err = unpopulate(val, "ServiceDisplayName", &t.ServiceDisplayName)
			delete(rawMsg, key)
		case "serviceId":
			err = unpopulate(val, "ServiceID", &t.ServiceID)
			delete(rawMsg, key)
		case "serviceLevelAgreement":
			err = unpopulate(val, "ServiceLevelAgreement", &t.ServiceLevelAgreement)
			delete(rawMsg, key)
		case "severity":
			err = unpopulate(val, "Severity", &t.Severity)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &t.Status)
			delete(rawMsg, key)
		case "supportEngineer":
			err = unpopulate(val, "SupportEngineer", &t.SupportEngineer)
			delete(rawMsg, key)
		case "supportPlanType":
			err = unpopulate(val, "SupportPlanType", &t.SupportPlanType)
			delete(rawMsg, key)
		case "supportTicketId":
			err = unpopulate(val, "SupportTicketID", &t.SupportTicketID)
			delete(rawMsg, key)
		case "technicalTicketDetails":
			err = unpopulate(val, "TechnicalTicketDetails", &t.TechnicalTicketDetails)
			delete(rawMsg, key)
		case "title":
			err = unpopulate(val, "Title", &t.Title)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", t, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UpdateContactProfile.
func (u UpdateContactProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalEmailAddresses", u.AdditionalEmailAddresses)
	populate(objectMap, "country", u.Country)
	populate(objectMap, "firstName", u.FirstName)
	populate(objectMap, "lastName", u.LastName)
	populate(objectMap, "phoneNumber", u.PhoneNumber)
	populate(objectMap, "preferredContactMethod", u.PreferredContactMethod)
	populate(objectMap, "preferredSupportLanguage", u.PreferredSupportLanguage)
	populate(objectMap, "preferredTimeZone", u.PreferredTimeZone)
	populate(objectMap, "primaryEmailAddress", u.PrimaryEmailAddress)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type UpdateSupportTicket.
func (u UpdateSupportTicket) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "contactDetails", u.ContactDetails)
	populate(objectMap, "severity", u.Severity)
	populate(objectMap, "status", u.Status)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}