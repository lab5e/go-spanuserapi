/*
The Span User API

API for sessions, teams and API tokens

API version: [version] [name]
Contact: contact@lab5e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanuserapi

import (
	"encoding/json"
	"fmt"
)

// SessionInfoType the model 'SessionInfoType'
type SessionInfoType string

// List of SessionInfo.Type
const (
	SESSIONINFOTYPE_UNSPECIFIED SessionInfoType = "TYPE_UNSPECIFIED"
	SESSIONINFOTYPE_DESKTOP SessionInfoType = "TYPE_DESKTOP"
	SESSIONINFOTYPE_MOBILE SessionInfoType = "TYPE_MOBILE"
	SESSIONINFOTYPE_TABLET SessionInfoType = "TYPE_TABLET"
	SESSIONINFOTYPE_BOT SessionInfoType = "TYPE_BOT"
)

// All allowed values of SessionInfoType enum
var AllowedSessionInfoTypeEnumValues = []SessionInfoType{
	"TYPE_UNSPECIFIED",
	"TYPE_DESKTOP",
	"TYPE_MOBILE",
	"TYPE_TABLET",
	"TYPE_BOT",
}

func (v *SessionInfoType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SessionInfoType(value)
	for _, existing := range AllowedSessionInfoTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SessionInfoType", value)
}

// NewSessionInfoTypeFromValue returns a pointer to a valid SessionInfoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSessionInfoTypeFromValue(v string) (*SessionInfoType, error) {
	ev := SessionInfoType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SessionInfoType: valid values are %v", v, AllowedSessionInfoTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SessionInfoType) IsValid() bool {
	for _, existing := range AllowedSessionInfoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SessionInfo.Type value
func (v SessionInfoType) Ptr() *SessionInfoType {
	return &v
}

type NullableSessionInfoType struct {
	value *SessionInfoType
	isSet bool
}

func (v NullableSessionInfoType) Get() *SessionInfoType {
	return v.value
}

func (v *NullableSessionInfoType) Set(val *SessionInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionInfoType(val *SessionInfoType) *NullableSessionInfoType {
	return &NullableSessionInfoType{value: val, isSet: true}
}

func (v NullableSessionInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

