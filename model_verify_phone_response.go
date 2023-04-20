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

// checks if the VerifyPhoneResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyPhoneResponse{}

// VerifyPhoneResponse struct for VerifyPhoneResponse
type VerifyPhoneResponse struct {
	User *User `json:"user,omitempty"`
}

// NewVerifyPhoneResponse instantiates a new VerifyPhoneResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyPhoneResponse() *VerifyPhoneResponse {
	this := VerifyPhoneResponse{}
	return &this
}

// NewVerifyPhoneResponseWithDefaults instantiates a new VerifyPhoneResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyPhoneResponseWithDefaults() *VerifyPhoneResponse {
	this := VerifyPhoneResponse{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *VerifyPhoneResponse) GetUser() User {
	if o == nil || IsNil(o.User) {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyPhoneResponse) GetUserOk() (*User, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *VerifyPhoneResponse) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *VerifyPhoneResponse) SetUser(v User) {
	o.User = &v
}

func (o VerifyPhoneResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyPhoneResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableVerifyPhoneResponse struct {
	value *VerifyPhoneResponse
	isSet bool
}

func (v NullableVerifyPhoneResponse) Get() *VerifyPhoneResponse {
	return v.value
}

func (v *NullableVerifyPhoneResponse) Set(val *VerifyPhoneResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyPhoneResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyPhoneResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyPhoneResponse(val *VerifyPhoneResponse) *NullableVerifyPhoneResponse {
	return &NullableVerifyPhoneResponse{value: val, isSet: true}
}

func (v NullableVerifyPhoneResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyPhoneResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


