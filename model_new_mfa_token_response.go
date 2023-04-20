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

// checks if the NewMFATokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewMFATokenResponse{}

// NewMFATokenResponse struct for NewMFATokenResponse
type NewMFATokenResponse struct {
	Secret *string `json:"secret,omitempty"`
	QrCode *string `json:"qrCode,omitempty"`
}

// NewNewMFATokenResponse instantiates a new NewMFATokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewMFATokenResponse() *NewMFATokenResponse {
	this := NewMFATokenResponse{}
	return &this
}

// NewNewMFATokenResponseWithDefaults instantiates a new NewMFATokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewMFATokenResponseWithDefaults() *NewMFATokenResponse {
	this := NewMFATokenResponse{}
	return &this
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *NewMFATokenResponse) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewMFATokenResponse) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *NewMFATokenResponse) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *NewMFATokenResponse) SetSecret(v string) {
	o.Secret = &v
}

// GetQrCode returns the QrCode field value if set, zero value otherwise.
func (o *NewMFATokenResponse) GetQrCode() string {
	if o == nil || IsNil(o.QrCode) {
		var ret string
		return ret
	}
	return *o.QrCode
}

// GetQrCodeOk returns a tuple with the QrCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewMFATokenResponse) GetQrCodeOk() (*string, bool) {
	if o == nil || IsNil(o.QrCode) {
		return nil, false
	}
	return o.QrCode, true
}

// HasQrCode returns a boolean if a field has been set.
func (o *NewMFATokenResponse) HasQrCode() bool {
	if o != nil && !IsNil(o.QrCode) {
		return true
	}

	return false
}

// SetQrCode gets a reference to the given string and assigns it to the QrCode field.
func (o *NewMFATokenResponse) SetQrCode(v string) {
	o.QrCode = &v
}

func (o NewMFATokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewMFATokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !IsNil(o.QrCode) {
		toSerialize["qrCode"] = o.QrCode
	}
	return toSerialize, nil
}

type NullableNewMFATokenResponse struct {
	value *NewMFATokenResponse
	isSet bool
}

func (v NullableNewMFATokenResponse) Get() *NewMFATokenResponse {
	return v.value
}

func (v *NullableNewMFATokenResponse) Set(val *NewMFATokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNewMFATokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNewMFATokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewMFATokenResponse(val *NewMFATokenResponse) *NullableNewMFATokenResponse {
	return &NullableNewMFATokenResponse{value: val, isSet: true}
}

func (v NullableNewMFATokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewMFATokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

