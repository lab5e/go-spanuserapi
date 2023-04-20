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

// checks if the VerifyNewMFATokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyNewMFATokenRequest{}

// VerifyNewMFATokenRequest struct for VerifyNewMFATokenRequest
type VerifyNewMFATokenRequest struct {
	Passcode *string `json:"passcode,omitempty"`
}

// NewVerifyNewMFATokenRequest instantiates a new VerifyNewMFATokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyNewMFATokenRequest() *VerifyNewMFATokenRequest {
	this := VerifyNewMFATokenRequest{}
	return &this
}

// NewVerifyNewMFATokenRequestWithDefaults instantiates a new VerifyNewMFATokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyNewMFATokenRequestWithDefaults() *VerifyNewMFATokenRequest {
	this := VerifyNewMFATokenRequest{}
	return &this
}

// GetPasscode returns the Passcode field value if set, zero value otherwise.
func (o *VerifyNewMFATokenRequest) GetPasscode() string {
	if o == nil || IsNil(o.Passcode) {
		var ret string
		return ret
	}
	return *o.Passcode
}

// GetPasscodeOk returns a tuple with the Passcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyNewMFATokenRequest) GetPasscodeOk() (*string, bool) {
	if o == nil || IsNil(o.Passcode) {
		return nil, false
	}
	return o.Passcode, true
}

// HasPasscode returns a boolean if a field has been set.
func (o *VerifyNewMFATokenRequest) HasPasscode() bool {
	if o != nil && !IsNil(o.Passcode) {
		return true
	}

	return false
}

// SetPasscode gets a reference to the given string and assigns it to the Passcode field.
func (o *VerifyNewMFATokenRequest) SetPasscode(v string) {
	o.Passcode = &v
}

func (o VerifyNewMFATokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyNewMFATokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Passcode) {
		toSerialize["passcode"] = o.Passcode
	}
	return toSerialize, nil
}

type NullableVerifyNewMFATokenRequest struct {
	value *VerifyNewMFATokenRequest
	isSet bool
}

func (v NullableVerifyNewMFATokenRequest) Get() *VerifyNewMFATokenRequest {
	return v.value
}

func (v *NullableVerifyNewMFATokenRequest) Set(val *VerifyNewMFATokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyNewMFATokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyNewMFATokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyNewMFATokenRequest(val *VerifyNewMFATokenRequest) *NullableVerifyNewMFATokenRequest {
	return &NullableVerifyNewMFATokenRequest{value: val, isSet: true}
}

func (v NullableVerifyNewMFATokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyNewMFATokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


