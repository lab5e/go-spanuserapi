# OrgTeamMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberId** | Pointer to **string** |  | [optional] 
**TeamRole** | Pointer to [**OrgRole**](OrgRole.md) |  | [optional] [default to ORGROLE_UNSPECIFIED]
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 

## Methods

### NewOrgTeamMember

`func NewOrgTeamMember() *OrgTeamMember`

NewOrgTeamMember instantiates a new OrgTeamMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgTeamMemberWithDefaults

`func NewOrgTeamMemberWithDefaults() *OrgTeamMember`

NewOrgTeamMemberWithDefaults instantiates a new OrgTeamMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberId

`func (o *OrgTeamMember) GetMemberId() string`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *OrgTeamMember) GetMemberIdOk() (*string, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *OrgTeamMember) SetMemberId(v string)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *OrgTeamMember) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetTeamRole

`func (o *OrgTeamMember) GetTeamRole() OrgRole`

GetTeamRole returns the TeamRole field if non-nil, zero value otherwise.

### GetTeamRoleOk

`func (o *OrgTeamMember) GetTeamRoleOk() (*OrgRole, bool)`

GetTeamRoleOk returns a tuple with the TeamRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamRole

`func (o *OrgTeamMember) SetTeamRole(v OrgRole)`

SetTeamRole sets TeamRole field to given value.

### HasTeamRole

`func (o *OrgTeamMember) HasTeamRole() bool`

HasTeamRole returns a boolean if a field has been set.

### GetName

`func (o *OrgTeamMember) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgTeamMember) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgTeamMember) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrgTeamMember) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *OrgTeamMember) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrgTeamMember) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrgTeamMember) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrgTeamMember) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


