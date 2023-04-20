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

// checks if the CreateTestUserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTestUserResponse{}

// CreateTestUserResponse struct for CreateTestUserResponse
type CreateTestUserResponse struct {
	Id *string `json:"id,omitempty"`
}

// NewCreateTestUserResponse instantiates a new CreateTestUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTestUserResponse() *CreateTestUserResponse {
	this := CreateTestUserResponse{}
	return &this
}

// NewCreateTestUserResponseWithDefaults instantiates a new CreateTestUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTestUserResponseWithDefaults() *CreateTestUserResponse {
	this := CreateTestUserResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateTestUserResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTestUserResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateTestUserResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateTestUserResponse) SetId(v string) {
	o.Id = &v
}

func (o CreateTestUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTestUserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableCreateTestUserResponse struct {
	value *CreateTestUserResponse
	isSet bool
}

func (v NullableCreateTestUserResponse) Get() *CreateTestUserResponse {
	return v.value
}

func (v *NullableCreateTestUserResponse) Set(val *CreateTestUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTestUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTestUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTestUserResponse(val *CreateTestUserResponse) *NullableCreateTestUserResponse {
	return &NullableCreateTestUserResponse{value: val, isSet: true}
}

func (v NullableCreateTestUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTestUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


