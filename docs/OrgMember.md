# OrgMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberId** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**OrgRole** | Pointer to [**OrgRole**](OrgRole.md) |  | [optional] [default to ORGROLE_UNSPECIFIED]
**Created** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**MemberStatus**](MemberStatus.md) |  | [optional] [default to MEMBERSTATUS_UNSPECIFIED]

## Methods

### NewOrgMember

`func NewOrgMember() *OrgMember`

NewOrgMember instantiates a new OrgMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgMemberWithDefaults

`func NewOrgMemberWithDefaults() *OrgMember`

NewOrgMemberWithDefaults instantiates a new OrgMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberId

`func (o *OrgMember) GetMemberId() string`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *OrgMember) GetMemberIdOk() (*string, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *OrgMember) SetMemberId(v string)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *OrgMember) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetOrgId

`func (o *OrgMember) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *OrgMember) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *OrgMember) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *OrgMember) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetUserId

`func (o *OrgMember) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OrgMember) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OrgMember) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *OrgMember) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetOrgRole

`func (o *OrgMember) GetOrgRole() OrgRole`

GetOrgRole returns the OrgRole field if non-nil, zero value otherwise.

### GetOrgRoleOk

`func (o *OrgMember) GetOrgRoleOk() (*OrgRole, bool)`

GetOrgRoleOk returns a tuple with the OrgRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgRole

`func (o *OrgMember) SetOrgRole(v OrgRole)`

SetOrgRole sets OrgRole field to given value.

### HasOrgRole

`func (o *OrgMember) HasOrgRole() bool`

HasOrgRole returns a boolean if a field has been set.

### GetCreated

`func (o *OrgMember) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OrgMember) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OrgMember) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *OrgMember) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetName

`func (o *OrgMember) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgMember) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgMember) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrgMember) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *OrgMember) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrgMember) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrgMember) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrgMember) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetStatus

`func (o *OrgMember) GetStatus() MemberStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrgMember) GetStatusOk() (*MemberStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrgMember) SetStatus(v MemberStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrgMember) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


