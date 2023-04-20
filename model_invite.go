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

// checks if the Invite type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Invite{}

// Invite struct for Invite
type Invite struct {
	InviteId *string `json:"inviteId,omitempty"`
	InviterName *string `json:"inviterName,omitempty"`
	InviterEmail *string `json:"inviterEmail,omitempty"`
	OrgName *string `json:"orgName,omitempty"`
	Created *string `json:"created,omitempty"`
}

// NewInvite instantiates a new Invite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvite() *Invite {
	this := Invite{}
	return &this
}

// NewInviteWithDefaults instantiates a new Invite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteWithDefaults() *Invite {
	this := Invite{}
	return &this
}

// GetInviteId returns the InviteId field value if set, zero value otherwise.
func (o *Invite) GetInviteId() string {
	if o == nil || IsNil(o.InviteId) {
		var ret string
		return ret
	}
	return *o.InviteId
}

// GetInviteIdOk returns a tuple with the InviteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invite) GetInviteIdOk() (*string, bool) {
	if o == nil || IsNil(o.InviteId) {
		return nil, false
	}
	return o.InviteId, true
}

// HasInviteId returns a boolean if a field has been set.
func (o *Invite) HasInviteId() bool {
	if o != nil && !IsNil(o.InviteId) {
		return true
	}

	return false
}

// SetInviteId gets a reference to the given string and assigns it to the InviteId field.
func (o *Invite) SetInviteId(v string) {
	o.InviteId = &v
}

// GetInviterName returns the InviterName field value if set, zero value otherwise.
func (o *Invite) GetInviterName() string {
	if o == nil || IsNil(o.InviterName) {
		var ret string
		return ret
	}
	return *o.InviterName
}

// GetInviterNameOk returns a tuple with the InviterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invite) GetInviterNameOk() (*string, bool) {
	if o == nil || IsNil(o.InviterName) {
		return nil, false
	}
	return o.InviterName, true
}

// HasInviterName returns a boolean if a field has been set.
func (o *Invite) HasInviterName() bool {
	if o != nil && !IsNil(o.InviterName) {
		return true
	}

	return false
}

// SetInviterName gets a reference to the given string and assigns it to the InviterName field.
func (o *Invite) SetInviterName(v string) {
	o.InviterName = &v
}

// GetInviterEmail returns the InviterEmail field value if set, zero value otherwise.
func (o *Invite) GetInviterEmail() string {
	if o == nil || IsNil(o.InviterEmail) {
		var ret string
		return ret
	}
	return *o.InviterEmail
}

// GetInviterEmailOk returns a tuple with the InviterEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invite) GetInviterEmailOk() (*string, bool) {
	if o == nil || IsNil(o.InviterEmail) {
		return nil, false
	}
	return o.InviterEmail, true
}

// HasInviterEmail returns a boolean if a field has been set.
func (o *Invite) HasInviterEmail() bool {
	if o != nil && !IsNil(o.InviterEmail) {
		return true
	}

	return false
}

// SetInviterEmail gets a reference to the given string and assigns it to the InviterEmail field.
func (o *Invite) SetInviterEmail(v string) {
	o.InviterEmail = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *Invite) GetOrgName() string {
	if o == nil || IsNil(o.OrgName) {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invite) GetOrgNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrgName) {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *Invite) HasOrgName() bool {
	if o != nil && !IsNil(o.OrgName) {
		return true
	}

	return false
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *Invite) SetOrgName(v string) {
	o.OrgName = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Invite) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invite) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Invite) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *Invite) SetCreated(v string) {
	o.Created = &v
}

func (o Invite) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Invite) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InviteId) {
		toSerialize["inviteId"] = o.InviteId
	}
	if !IsNil(o.InviterName) {
		toSerialize["inviterName"] = o.InviterName
	}
	if !IsNil(o.InviterEmail) {
		toSerialize["inviterEmail"] = o.InviterEmail
	}
	if !IsNil(o.OrgName) {
		toSerialize["orgName"] = o.OrgName
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	return toSerialize, nil
}

type NullableInvite struct {
	value *Invite
	isSet bool
}

func (v NullableInvite) Get() *Invite {
	return v.value
}

func (v *NullableInvite) Set(val *Invite) {
	v.value = val
	v.isSet = true
}

func (v NullableInvite) IsSet() bool {
	return v.isSet
}

func (v *NullableInvite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvite(val *Invite) *NullableInvite {
	return &NullableInvite{value: val, isSet: true}
}

func (v NullableInvite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


