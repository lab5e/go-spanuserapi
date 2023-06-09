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

// checks if the EndUserLicenseAgreementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndUserLicenseAgreementResponse{}

// EndUserLicenseAgreementResponse struct for EndUserLicenseAgreementResponse
type EndUserLicenseAgreementResponse struct {
	Text *string `json:"text,omitempty"`
	Version *int32 `json:"version,omitempty"`
}

// NewEndUserLicenseAgreementResponse instantiates a new EndUserLicenseAgreementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndUserLicenseAgreementResponse() *EndUserLicenseAgreementResponse {
	this := EndUserLicenseAgreementResponse{}
	return &this
}

// NewEndUserLicenseAgreementResponseWithDefaults instantiates a new EndUserLicenseAgreementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndUserLicenseAgreementResponseWithDefaults() *EndUserLicenseAgreementResponse {
	this := EndUserLicenseAgreementResponse{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *EndUserLicenseAgreementResponse) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndUserLicenseAgreementResponse) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *EndUserLicenseAgreementResponse) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *EndUserLicenseAgreementResponse) SetText(v string) {
	o.Text = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *EndUserLicenseAgreementResponse) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndUserLicenseAgreementResponse) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *EndUserLicenseAgreementResponse) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *EndUserLicenseAgreementResponse) SetVersion(v int32) {
	o.Version = &v
}

func (o EndUserLicenseAgreementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndUserLicenseAgreementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableEndUserLicenseAgreementResponse struct {
	value *EndUserLicenseAgreementResponse
	isSet bool
}

func (v NullableEndUserLicenseAgreementResponse) Get() *EndUserLicenseAgreementResponse {
	return v.value
}

func (v *NullableEndUserLicenseAgreementResponse) Set(val *EndUserLicenseAgreementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEndUserLicenseAgreementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEndUserLicenseAgreementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndUserLicenseAgreementResponse(val *EndUserLicenseAgreementResponse) *NullableEndUserLicenseAgreementResponse {
	return &NullableEndUserLicenseAgreementResponse{value: val, isSet: true}
}

func (v NullableEndUserLicenseAgreementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndUserLicenseAgreementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


