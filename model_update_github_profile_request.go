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

// checks if the UpdateGithubProfileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateGithubProfileRequest{}

// UpdateGithubProfileRequest Request from client when profile is updated. Token is forwarded to the client via redirects on GH auth.
type UpdateGithubProfileRequest struct {
	Name *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
	Address *string `json:"address,omitempty"`
	Phone *string `json:"phone,omitempty"`
	EulaAccepted *bool `json:"eulaAccepted,omitempty"`
	EulaVersion *int32 `json:"eulaVersion,omitempty"`
}

// NewUpdateGithubProfileRequest instantiates a new UpdateGithubProfileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGithubProfileRequest() *UpdateGithubProfileRequest {
	this := UpdateGithubProfileRequest{}
	return &this
}

// NewUpdateGithubProfileRequestWithDefaults instantiates a new UpdateGithubProfileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGithubProfileRequestWithDefaults() *UpdateGithubProfileRequest {
	this := UpdateGithubProfileRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateGithubProfileRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGithubProfileRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateGithubProfileRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateGithubProfileRequest) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UpdateGithubProfileRequest) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGithubProfileRequest) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateGithubProfileRequest) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UpdateGithubProfileRequest) SetEmail(v string) {
	o.Email = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *UpdateGithubProfileRequest) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGithubProfileRequest) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *UpdateGithubProfileRequest) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *UpdateGithubProfileRequest) SetAddress(v string) {
	o.Address = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *UpdateGithubProfileRequest) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGithubProfileRequest) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *UpdateGithubProfileRequest) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *UpdateGithubProfileRequest) SetPhone(v string) {
	o.Phone = &v
}

// GetEulaAccepted returns the EulaAccepted field value if set, zero value otherwise.
func (o *UpdateGithubProfileRequest) GetEulaAccepted() bool {
	if o == nil || IsNil(o.EulaAccepted) {
		var ret bool
		return ret
	}
	return *o.EulaAccepted
}

// GetEulaAcceptedOk returns a tuple with the EulaAccepted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGithubProfileRequest) GetEulaAcceptedOk() (*bool, bool) {
	if o == nil || IsNil(o.EulaAccepted) {
		return nil, false
	}
	return o.EulaAccepted, true
}

// HasEulaAccepted returns a boolean if a field has been set.
func (o *UpdateGithubProfileRequest) HasEulaAccepted() bool {
	if o != nil && !IsNil(o.EulaAccepted) {
		return true
	}

	return false
}

// SetEulaAccepted gets a reference to the given bool and assigns it to the EulaAccepted field.
func (o *UpdateGithubProfileRequest) SetEulaAccepted(v bool) {
	o.EulaAccepted = &v
}

// GetEulaVersion returns the EulaVersion field value if set, zero value otherwise.
func (o *UpdateGithubProfileRequest) GetEulaVersion() int32 {
	if o == nil || IsNil(o.EulaVersion) {
		var ret int32
		return ret
	}
	return *o.EulaVersion
}

// GetEulaVersionOk returns a tuple with the EulaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGithubProfileRequest) GetEulaVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.EulaVersion) {
		return nil, false
	}
	return o.EulaVersion, true
}

// HasEulaVersion returns a boolean if a field has been set.
func (o *UpdateGithubProfileRequest) HasEulaVersion() bool {
	if o != nil && !IsNil(o.EulaVersion) {
		return true
	}

	return false
}

// SetEulaVersion gets a reference to the given int32 and assigns it to the EulaVersion field.
func (o *UpdateGithubProfileRequest) SetEulaVersion(v int32) {
	o.EulaVersion = &v
}

func (o UpdateGithubProfileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateGithubProfileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.EulaAccepted) {
		toSerialize["eulaAccepted"] = o.EulaAccepted
	}
	if !IsNil(o.EulaVersion) {
		toSerialize["eulaVersion"] = o.EulaVersion
	}
	return toSerialize, nil
}

type NullableUpdateGithubProfileRequest struct {
	value *UpdateGithubProfileRequest
	isSet bool
}

func (v NullableUpdateGithubProfileRequest) Get() *UpdateGithubProfileRequest {
	return v.value
}

func (v *NullableUpdateGithubProfileRequest) Set(val *UpdateGithubProfileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGithubProfileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGithubProfileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGithubProfileRequest(val *UpdateGithubProfileRequest) *NullableUpdateGithubProfileRequest {
	return &NullableUpdateGithubProfileRequest{value: val, isSet: true}
}

func (v NullableUpdateGithubProfileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGithubProfileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


