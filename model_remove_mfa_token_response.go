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

// checks if the RemoveMFATokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveMFATokenResponse{}

// RemoveMFATokenResponse struct for RemoveMFATokenResponse
type RemoveMFATokenResponse struct {
	User *User `json:"user,omitempty"`
}

// NewRemoveMFATokenResponse instantiates a new RemoveMFATokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveMFATokenResponse() *RemoveMFATokenResponse {
	this := RemoveMFATokenResponse{}
	return &this
}

// NewRemoveMFATokenResponseWithDefaults instantiates a new RemoveMFATokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveMFATokenResponseWithDefaults() *RemoveMFATokenResponse {
	this := RemoveMFATokenResponse{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *RemoveMFATokenResponse) GetUser() User {
	if o == nil || IsNil(o.User) {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveMFATokenResponse) GetUserOk() (*User, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *RemoveMFATokenResponse) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *RemoveMFATokenResponse) SetUser(v User) {
	o.User = &v
}

func (o RemoveMFATokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveMFATokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableRemoveMFATokenResponse struct {
	value *RemoveMFATokenResponse
	isSet bool
}

func (v NullableRemoveMFATokenResponse) Get() *RemoveMFATokenResponse {
	return v.value
}

func (v *NullableRemoveMFATokenResponse) Set(val *RemoveMFATokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveMFATokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveMFATokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveMFATokenResponse(val *RemoveMFATokenResponse) *NullableRemoveMFATokenResponse {
	return &NullableRemoveMFATokenResponse{value: val, isSet: true}
}

func (v NullableRemoveMFATokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveMFATokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


