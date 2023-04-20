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

// checks if the LoginRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginRequest{}

// LoginRequest struct for LoginRequest
type LoginRequest struct {
	Email *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
	Passcode *string `json:"passcode,omitempty"`
}

// NewLoginRequest instantiates a new LoginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginRequest() *LoginRequest {
	this := LoginRequest{}
	return &this
}

// NewLoginRequestWithDefaults instantiates a new LoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginRequestWithDefaults() *LoginRequest {
	this := LoginRequest{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *LoginRequest) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginRequest) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *LoginRequest) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *LoginRequest) SetEmail(v string) {
	o.Email = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *LoginRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *LoginRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *LoginRequest) SetPassword(v string) {
	o.Password = &v
}

// GetPasscode returns the Passcode field value if set, zero value otherwise.
func (o *LoginRequest) GetPasscode() string {
	if o == nil || IsNil(o.Passcode) {
		var ret string
		return ret
	}
	return *o.Passcode
}

// GetPasscodeOk returns a tuple with the Passcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginRequest) GetPasscodeOk() (*string, bool) {
	if o == nil || IsNil(o.Passcode) {
		return nil, false
	}
	return o.Passcode, true
}

// HasPasscode returns a boolean if a field has been set.
func (o *LoginRequest) HasPasscode() bool {
	if o != nil && !IsNil(o.Passcode) {
		return true
	}

	return false
}

// SetPasscode gets a reference to the given string and assigns it to the Passcode field.
func (o *LoginRequest) SetPasscode(v string) {
	o.Passcode = &v
}

func (o LoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Passcode) {
		toSerialize["passcode"] = o.Passcode
	}
	return toSerialize, nil
}

type NullableLoginRequest struct {
	value *LoginRequest
	isSet bool
}

func (v NullableLoginRequest) Get() *LoginRequest {
	return v.value
}

func (v *NullableLoginRequest) Set(val *LoginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginRequest(val *LoginRequest) *NullableLoginRequest {
	return &NullableLoginRequest{value: val, isSet: true}
}

func (v NullableLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


