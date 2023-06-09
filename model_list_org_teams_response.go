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

// checks if the ListOrgTeamsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOrgTeamsResponse{}

// ListOrgTeamsResponse struct for ListOrgTeamsResponse
type ListOrgTeamsResponse struct {
	Teams []OrgTeam `json:"teams,omitempty"`
}

// NewListOrgTeamsResponse instantiates a new ListOrgTeamsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOrgTeamsResponse() *ListOrgTeamsResponse {
	this := ListOrgTeamsResponse{}
	return &this
}

// NewListOrgTeamsResponseWithDefaults instantiates a new ListOrgTeamsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOrgTeamsResponseWithDefaults() *ListOrgTeamsResponse {
	this := ListOrgTeamsResponse{}
	return &this
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *ListOrgTeamsResponse) GetTeams() []OrgTeam {
	if o == nil || IsNil(o.Teams) {
		var ret []OrgTeam
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrgTeamsResponse) GetTeamsOk() ([]OrgTeam, bool) {
	if o == nil || IsNil(o.Teams) {
		return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *ListOrgTeamsResponse) HasTeams() bool {
	if o != nil && !IsNil(o.Teams) {
		return true
	}

	return false
}

// SetTeams gets a reference to the given []OrgTeam and assigns it to the Teams field.
func (o *ListOrgTeamsResponse) SetTeams(v []OrgTeam) {
	o.Teams = v
}

func (o ListOrgTeamsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOrgTeamsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Teams) {
		toSerialize["teams"] = o.Teams
	}
	return toSerialize, nil
}

type NullableListOrgTeamsResponse struct {
	value *ListOrgTeamsResponse
	isSet bool
}

func (v NullableListOrgTeamsResponse) Get() *ListOrgTeamsResponse {
	return v.value
}

func (v *NullableListOrgTeamsResponse) Set(val *ListOrgTeamsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListOrgTeamsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListOrgTeamsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOrgTeamsResponse(val *ListOrgTeamsResponse) *NullableListOrgTeamsResponse {
	return &NullableListOrgTeamsResponse{value: val, isSet: true}
}

func (v NullableListOrgTeamsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOrgTeamsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


