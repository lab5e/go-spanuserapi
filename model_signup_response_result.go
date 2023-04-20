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

// SignupResponseResult the model 'SignupResponseResult'
type SignupResponseResult string

// List of SignupResponse.Result
const (
	SIGNUPRESPONSERESULT_UNSPECIFIED SignupResponseResult = "RESULT_UNSPECIFIED"
	SIGNUPRESPONSERESULT_SUCCESS SignupResponseResult = "RESULT_SUCCESS"
)

// All allowed values of SignupResponseResult enum
var AllowedSignupResponseResultEnumValues = []SignupResponseResult{
	"RESULT_UNSPECIFIED",
	"RESULT_SUCCESS",
}

func (v *SignupResponseResult) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SignupResponseResult(value)
	for _, existing := range AllowedSignupResponseResultEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SignupResponseResult", value)
}

// NewSignupResponseResultFromValue returns a pointer to a valid SignupResponseResult
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSignupResponseResultFromValue(v string) (*SignupResponseResult, error) {
	ev := SignupResponseResult(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SignupResponseResult: valid values are %v", v, AllowedSignupResponseResultEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SignupResponseResult) IsValid() bool {
	for _, existing := range AllowedSignupResponseResultEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SignupResponse.Result value
func (v SignupResponseResult) Ptr() *SignupResponseResult {
	return &v
}

type NullableSignupResponseResult struct {
	value *SignupResponseResult
	isSet bool
}

func (v NullableSignupResponseResult) Get() *SignupResponseResult {
	return v.value
}

func (v *NullableSignupResponseResult) Set(val *SignupResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSignupResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSignupResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignupResponseResult(val *SignupResponseResult) *NullableSignupResponseResult {
	return &NullableSignupResponseResult{value: val, isSet: true}
}

func (v NullableSignupResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignupResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

