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

// checks if the RecoverPasswordRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverPasswordRequest{}

// RecoverPasswordRequest struct for RecoverPasswordRequest
type RecoverPasswordRequest struct {
	Email *string `json:"email,omitempty"`
	RecoverToken *string `json:"recoverToken,omitempty"`
	Passcode *string `json:"passcode,omitempty"`
	NewPassword *string `json:"newPassword,omitempty"`
}

// NewRecoverPasswordRequest instantiates a new RecoverPasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverPasswordRequest() *RecoverPasswordRequest {
	this := RecoverPasswordRequest{}
	return &this
}

// NewRecoverPasswordRequestWithDefaults instantiates a new RecoverPasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverPasswordRequestWithDefaults() *RecoverPasswordRequest {
	this := RecoverPasswordRequest{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *RecoverPasswordRequest) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverPasswordRequest) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *RecoverPasswordRequest) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *RecoverPasswordRequest) SetEmail(v string) {
	o.Email = &v
}

// GetRecoverToken returns the RecoverToken field value if set, zero value otherwise.
func (o *RecoverPasswordRequest) GetRecoverToken() string {
	if o == nil || IsNil(o.RecoverToken) {
		var ret string
		return ret
	}
	return *o.RecoverToken
}

// GetRecoverTokenOk returns a tuple with the RecoverToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverPasswordRequest) GetRecoverTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RecoverToken) {
		return nil, false
	}
	return o.RecoverToken, true
}

// HasRecoverToken returns a boolean if a field has been set.
func (o *RecoverPasswordRequest) HasRecoverToken() bool {
	if o != nil && !IsNil(o.RecoverToken) {
		return true
	}

	return false
}

// SetRecoverToken gets a reference to the given string and assigns it to the RecoverToken field.
func (o *RecoverPasswordRequest) SetRecoverToken(v string) {
	o.RecoverToken = &v
}

// GetPasscode returns the Passcode field value if set, zero value otherwise.
func (o *RecoverPasswordRequest) GetPasscode() string {
	if o == nil || IsNil(o.Passcode) {
		var ret string
		return ret
	}
	return *o.Passcode
}

// GetPasscodeOk returns a tuple with the Passcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverPasswordRequest) GetPasscodeOk() (*string, bool) {
	if o == nil || IsNil(o.Passcode) {
		return nil, false
	}
	return o.Passcode, true
}

// HasPasscode returns a boolean if a field has been set.
func (o *RecoverPasswordRequest) HasPasscode() bool {
	if o != nil && !IsNil(o.Passcode) {
		return true
	}

	return false
}

// SetPasscode gets a reference to the given string and assigns it to the Passcode field.
func (o *RecoverPasswordRequest) SetPasscode(v string) {
	o.Passcode = &v
}

// GetNewPassword returns the NewPassword field value if set, zero value otherwise.
func (o *RecoverPasswordRequest) GetNewPassword() string {
	if o == nil || IsNil(o.NewPassword) {
		var ret string
		return ret
	}
	return *o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverPasswordRequest) GetNewPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.NewPassword) {
		return nil, false
	}
	return o.NewPassword, true
}

// HasNewPassword returns a boolean if a field has been set.
func (o *RecoverPasswordRequest) HasNewPassword() bool {
	if o != nil && !IsNil(o.NewPassword) {
		return true
	}

	return false
}

// SetNewPassword gets a reference to the given string and assigns it to the NewPassword field.
func (o *RecoverPasswordRequest) SetNewPassword(v string) {
	o.NewPassword = &v
}

func (o RecoverPasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverPasswordRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.RecoverToken) {
		toSerialize["recoverToken"] = o.RecoverToken
	}
	if !IsNil(o.Passcode) {
		toSerialize["passcode"] = o.Passcode
	}
	if !IsNil(o.NewPassword) {
		toSerialize["newPassword"] = o.NewPassword
	}
	return toSerialize, nil
}

type NullableRecoverPasswordRequest struct {
	value *RecoverPasswordRequest
	isSet bool
}

func (v NullableRecoverPasswordRequest) Get() *RecoverPasswordRequest {
	return v.value
}

func (v *NullableRecoverPasswordRequest) Set(val *RecoverPasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverPasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverPasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverPasswordRequest(val *RecoverPasswordRequest) *NullableRecoverPasswordRequest {
	return &NullableRecoverPasswordRequest{value: val, isSet: true}
}

func (v NullableRecoverPasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverPasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


