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

// checks if the OrgTeam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgTeam{}

// OrgTeam Team in organisation. Teams are subgroups of member in the organisations, a team could be \"dev team\", \"test test\" or \"prod team\" depending on how you organise your work. Teams (and consequently all the resources owned by the team) can be moved from one organisation to another
type OrgTeam struct {
	TeamId *string `json:"teamId,omitempty"`
	Name *string `json:"name,omitempty"`
	Created *string `json:"created,omitempty"`
	// Members of team.
	Members []OrgTeamMember `json:"members,omitempty"`
}

// NewOrgTeam instantiates a new OrgTeam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgTeam() *OrgTeam {
	this := OrgTeam{}
	return &this
}

// NewOrgTeamWithDefaults instantiates a new OrgTeam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgTeamWithDefaults() *OrgTeam {
	this := OrgTeam{}
	return &this
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *OrgTeam) GetTeamId() string {
	if o == nil || IsNil(o.TeamId) {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgTeam) GetTeamIdOk() (*string, bool) {
	if o == nil || IsNil(o.TeamId) {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *OrgTeam) HasTeamId() bool {
	if o != nil && !IsNil(o.TeamId) {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *OrgTeam) SetTeamId(v string) {
	o.TeamId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrgTeam) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgTeam) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrgTeam) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrgTeam) SetName(v string) {
	o.Name = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *OrgTeam) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgTeam) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *OrgTeam) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *OrgTeam) SetCreated(v string) {
	o.Created = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *OrgTeam) GetMembers() []OrgTeamMember {
	if o == nil || IsNil(o.Members) {
		var ret []OrgTeamMember
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgTeam) GetMembersOk() ([]OrgTeamMember, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *OrgTeam) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []OrgTeamMember and assigns it to the Members field.
func (o *OrgTeam) SetMembers(v []OrgTeamMember) {
	o.Members = v
}

func (o OrgTeam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgTeam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TeamId) {
		toSerialize["teamId"] = o.TeamId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	return toSerialize, nil
}

type NullableOrgTeam struct {
	value *OrgTeam
	isSet bool
}

func (v NullableOrgTeam) Get() *OrgTeam {
	return v.value
}

func (v *NullableOrgTeam) Set(val *OrgTeam) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgTeam) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgTeam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgTeam(val *OrgTeam) *NullableOrgTeam {
	return &NullableOrgTeam{value: val, isSet: true}
}

func (v NullableOrgTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgTeam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


