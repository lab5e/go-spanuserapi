# UpdateOrgMemberRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to [**OrgRole**](OrgRole.md) |  | [optional] [default to ORGROLE_UNSPECIFIED]

## Methods

### NewUpdateOrgMemberRoleRequest

`func NewUpdateOrgMemberRoleRequest() *UpdateOrgMemberRoleRequest`

NewUpdateOrgMemberRoleRequest instantiates a new UpdateOrgMemberRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrgMemberRoleRequestWithDefaults

`func NewUpdateOrgMemberRoleRequestWithDefaults() *UpdateOrgMemberRoleRequest`

NewUpdateOrgMemberRoleRequestWithDefaults instantiates a new UpdateOrgMemberRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *UpdateOrgMemberRoleRequest) GetRole() OrgRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UpdateOrgMemberRoleRequest) GetRoleOk() (*OrgRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UpdateOrgMemberRoleRequest) SetRole(v OrgRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *UpdateOrgMemberRoleRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


