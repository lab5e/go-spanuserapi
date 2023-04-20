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

// checks if the APIToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIToken{}

// APIToken API Tokens are used for
type APIToken struct {
	TokenId *string `json:"tokenId,omitempty"`
	TeamId *string `json:"teamId,omitempty"`
	Resource *string `json:"resource,omitempty"`
	Write *bool `json:"write,omitempty"`
	Token *string `json:"token,omitempty"`
	Name *string `json:"name,omitempty"`
	LastUse *string `json:"lastUse,omitempty"`
	Uses *string `json:"uses,omitempty"`
}

// NewAPIToken instantiates a new APIToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIToken() *APIToken {
	this := APIToken{}
	return &this
}

// NewAPITokenWithDefaults instantiates a new APIToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPITokenWithDefaults() *APIToken {
	this := APIToken{}
	return &this
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *APIToken) GetTokenId() string {
	if o == nil || IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIToken) GetTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *APIToken) HasTokenId() bool {
	if o != nil && !IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *APIToken) SetTokenId(v string) {
	o.TokenId = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *APIToken) GetTeamId() string {
	if o == nil || IsNil(o.TeamId) {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIToken) GetTeamIdOk() (*string, bool) {
	if o == nil || IsNil(o.TeamId) {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *APIToken) HasTeamId() bool {
	if o != nil && !IsNil(o.TeamId) {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *APIToken) SetTeamId(v string) {
	o.TeamId = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *APIToken) GetResource() string {
	if o == nil || IsNil(o.Resource) {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIToken) GetResourceOk() (*string, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *APIToken) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *APIToken) SetResource(v string) {
	o.Resource = &v
}

// GetWrite returns the Write field value if set, zero value otherwise.
func (o *APIToken) GetWrite() bool {
	if o == nil || IsNil(o.Write) {
		var ret bool
		return ret
	}
	return *o.Write
}

// GetWriteOk returns a tuple with the Write field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIToken) GetWriteOk() (*bool, bool) {
	if o == nil || IsNil(o.Write) {
		return nil, false
	}
	return o.Write, true
}

// HasWrite returns a boolean if a field has been set.
func (o *APIToken) HasWrite() bool {
	if o != nil && !IsNil(o.Write) {
		return true
	}

	return false
}

// SetWrite gets a reference to the given bool and assigns it to the Write field.
func (o *APIToken) SetWrite(v bool) {
	o.Write = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *APIToken) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIToken) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *APIToken) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *APIToken) SetToken(v string) {
	o.Token = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *APIToken) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIToken) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *APIToken) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *APIToken) SetName(v string) {
	o.Name = &v
}

// GetLastUse returns the LastUse field value if set, zero value otherwise.
func (o *APIToken) GetLastUse() string {
	if o == nil || IsNil(o.LastUse) {
		var ret string
		return ret
	}
	return *o.LastUse
}

// GetLastUseOk returns a tuple with the LastUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIToken) GetLastUseOk() (*string, bool) {
	if o == nil || IsNil(o.LastUse) {
		return nil, false
	}
	return o.LastUse, true
}

// HasLastUse returns a boolean if a field has been set.
func (o *APIToken) HasLastUse() bool {
	if o != nil && !IsNil(o.LastUse) {
		return true
	}

	return false
}

// SetLastUse gets a reference to the given string and assigns it to the LastUse field.
func (o *APIToken) SetLastUse(v string) {
	o.LastUse = &v
}

// GetUses returns the Uses field value if set, zero value otherwise.
func (o *APIToken) GetUses() string {
	if o == nil || IsNil(o.Uses) {
		var ret string
		return ret
	}
	return *o.Uses
}

// GetUsesOk returns a tuple with the Uses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIToken) GetUsesOk() (*string, bool) {
	if o == nil || IsNil(o.Uses) {
		return nil, false
	}
	return o.Uses, true
}

// HasUses returns a boolean if a field has been set.
func (o *APIToken) HasUses() bool {
	if o != nil && !IsNil(o.Uses) {
		return true
	}

	return false
}

// SetUses gets a reference to the given string and assigns it to the Uses field.
func (o *APIToken) SetUses(v string) {
	o.Uses = &v
}

func (o APIToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TokenId) {
		toSerialize["tokenId"] = o.TokenId
	}
	if !IsNil(o.TeamId) {
		toSerialize["teamId"] = o.TeamId
	}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	if !IsNil(o.Write) {
		toSerialize["write"] = o.Write
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.LastUse) {
		toSerialize["lastUse"] = o.LastUse
	}
	if !IsNil(o.Uses) {
		toSerialize["uses"] = o.Uses
	}
	return toSerialize, nil
}

type NullableAPIToken struct {
	value *APIToken
	isSet bool
}

func (v NullableAPIToken) Get() *APIToken {
	return v.value
}

func (v *NullableAPIToken) Set(val *APIToken) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIToken) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIToken(val *APIToken) *NullableAPIToken {
	return &NullableAPIToken{value: val, isSet: true}
}

func (v NullableAPIToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


