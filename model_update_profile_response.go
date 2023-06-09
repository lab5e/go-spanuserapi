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

// checks if the UpdateProfileResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProfileResponse{}

// UpdateProfileResponse struct for UpdateProfileResponse
type UpdateProfileResponse struct {
	User *User `json:"user,omitempty"`
}

// NewUpdateProfileResponse instantiates a new UpdateProfileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProfileResponse() *UpdateProfileResponse {
	this := UpdateProfileResponse{}
	return &this
}

// NewUpdateProfileResponseWithDefaults instantiates a new UpdateProfileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProfileResponseWithDefaults() *UpdateProfileResponse {
	this := UpdateProfileResponse{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UpdateProfileResponse) GetUser() User {
	if o == nil || IsNil(o.User) {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProfileResponse) GetUserOk() (*User, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UpdateProfileResponse) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *UpdateProfileResponse) SetUser(v User) {
	o.User = &v
}

func (o UpdateProfileResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProfileResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableUpdateProfileResponse struct {
	value *UpdateProfileResponse
	isSet bool
}

func (v NullableUpdateProfileResponse) Get() *UpdateProfileResponse {
	return v.value
}

func (v *NullableUpdateProfileResponse) Set(val *UpdateProfileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProfileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProfileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProfileResponse(val *UpdateProfileResponse) *NullableUpdateProfileResponse {
	return &NullableUpdateProfileResponse{value: val, isSet: true}
}

func (v NullableUpdateProfileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProfileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


