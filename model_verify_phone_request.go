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
)

// checks if the VerifyPhoneRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyPhoneRequest{}

// VerifyPhoneRequest struct for VerifyPhoneRequest
type VerifyPhoneRequest struct {
	Code *string `json:"code,omitempty"`
}

// NewVerifyPhoneRequest instantiates a new VerifyPhoneRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyPhoneRequest() *VerifyPhoneRequest {
	this := VerifyPhoneRequest{}
	return &this
}

// NewVerifyPhoneRequestWithDefaults instantiates a new VerifyPhoneRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyPhoneRequestWithDefaults() *VerifyPhoneRequest {
	this := VerifyPhoneRequest{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *VerifyPhoneRequest) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPhoneRequest) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *VerifyPhoneRequest) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *VerifyPhoneRequest) SetCode(v string) {
	o.Code = &v
}

func (o VerifyPhoneRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyPhoneRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	return toSerialize, nil
}

type NullableVerifyPhoneRequest struct {
	value *VerifyPhoneRequest
	isSet bool
}

func (v NullableVerifyPhoneRequest) Get() *VerifyPhoneRequest {
	return v.value
}

func (v *NullableVerifyPhoneRequest) Set(val *VerifyPhoneRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyPhoneRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyPhoneRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyPhoneRequest(val *VerifyPhoneRequest) *NullableVerifyPhoneRequest {
	return &NullableVerifyPhoneRequest{value: val, isSet: true}
}

func (v NullableVerifyPhoneRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyPhoneRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


