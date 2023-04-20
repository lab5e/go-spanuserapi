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

// checks if the OrgMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgMember{}

// OrgMember Member in an organisation.
type OrgMember struct {
	MemberId *string `json:"memberId,omitempty"`
	OrgId *string `json:"orgId,omitempty"`
	UserId *string `json:"userId,omitempty"`
	OrgRole *OrgRole `json:"orgRole,omitempty"`
	Created *string `json:"created,omitempty"`
	Name *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
	Status *MemberStatus `json:"status,omitempty"`
}

// NewOrgMember instantiates a new OrgMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgMember() *OrgMember {
	this := OrgMember{}
	var orgRole OrgRole = ORGROLE_UNSPECIFIED
	this.OrgRole = &orgRole
	var status MemberStatus = MEMBERSTATUS_UNSPECIFIED
	this.Status = &status
	return &this
}

// NewOrgMemberWithDefaults instantiates a new OrgMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgMemberWithDefaults() *OrgMember {
	this := OrgMember{}
	var orgRole OrgRole = ORGROLE_UNSPECIFIED
	this.OrgRole = &orgRole
	var status MemberStatus = MEMBERSTATUS_UNSPECIFIED
	this.Status = &status
	return &this
}

// GetMemberId returns the MemberId field value if set, zero value otherwise.
func (o *OrgMember) GetMemberId() string {
	if o == nil || IsNil(o.MemberId) {
		var ret string
		return ret
	}
	return *o.MemberId
}

// GetMemberIdOk returns a tuple with the MemberId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgMember) GetMemberIdOk() (*string, bool) {
	if o == nil || IsNil(o.MemberId) {
		return nil, false
	}
	return o.MemberId, true
}

// HasMemberId returns a boolean if a field has been set.
func (o *OrgMember) HasMemberId() bool {
	if o != nil && !IsNil(o.MemberId) {
		return true
	}

	return false
}

// SetMemberId gets a reference to the given string and assigns it to the MemberId field.
func (o *OrgMember) SetMemberId(v string) {
	o.MemberId = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *OrgMember) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgMember) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *OrgMember) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *OrgMember) SetOrgId(v string) {
	o.OrgId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *OrgMember) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgMember) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *OrgMember) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *OrgMember) SetUserId(v string) {
	o.UserId = &v
}

// GetOrgRole returns the OrgRole field value if set, zero value otherwise.
func (o *OrgMember) GetOrgRole() OrgRole {
	if o == nil || IsNil(o.OrgRole) {
		var ret OrgRole
		return ret
	}
	return *o.OrgRole
}

// GetOrgRoleOk returns a tuple with the OrgRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgMember) GetOrgRoleOk() (*OrgRole, bool) {
	if o == nil || IsNil(o.OrgRole) {
		return nil, false
	}
	return o.OrgRole, true
}

// HasOrgRole returns a boolean if a field has been set.
func (o *OrgMember) HasOrgRole() bool {
	if o != nil && !IsNil(o.OrgRole) {
		return true
	}

	return false
}

// SetOrgRole gets a reference to the given OrgRole and assigns it to the OrgRole field.
func (o *OrgMember) SetOrgRole(v OrgRole) {
	o.OrgRole = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *OrgMember) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgMember) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *OrgMember) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *OrgMember) SetCreated(v string) {
	o.Created = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrgMember) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgMember) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrgMember) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrgMember) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *OrgMember) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgMember) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *OrgMember) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *OrgMember) SetEmail(v string) {
	o.Email = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrgMember) GetStatus() MemberStatus {
	if o == nil || IsNil(o.Status) {
		var ret MemberStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgMember) GetStatusOk() (*MemberStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrgMember) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given MemberStatus and assigns it to the Status field.
func (o *OrgMember) SetStatus(v MemberStatus) {
	o.Status = &v
}

func (o OrgMember) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MemberId) {
		toSerialize["memberId"] = o.MemberId
	}
	if !IsNil(o.OrgId) {
		toSerialize["orgId"] = o.OrgId
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.OrgRole) {
		toSerialize["orgRole"] = o.OrgRole
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableOrgMember struct {
	value *OrgMember
	isSet bool
}

func (v NullableOrgMember) Get() *OrgMember {
	return v.value
}

func (v *NullableOrgMember) Set(val *OrgMember) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgMember) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgMember(val *OrgMember) *NullableOrgMember {
	return &NullableOrgMember{value: val, isSet: true}
}

func (v NullableOrgMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

