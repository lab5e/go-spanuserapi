# TeamMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberId** | Pointer to **string** |  | [optional] 
**Role** | Pointer to [**OrgRole**](OrgRole.md) |  | [optional] [default to ORGROLE_UNSPECIFIED]

## Methods

### NewTeamMemberRequest

`func NewTeamMemberRequest() *TeamMemberRequest`

NewTeamMemberRequest instantiates a new TeamMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamMemberRequestWithDefaults

`func NewTeamMemberRequestWithDefaults() *TeamMemberRequest`

NewTeamMemberRequestWithDefaults instantiates a new TeamMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberId

`func (o *TeamMemberRequest) GetMemberId() string`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *TeamMemberRequest) GetMemberIdOk() (*string, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *TeamMemberRequest) SetMemberId(v string)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *TeamMemberRequest) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetRole

`func (o *TeamMemberRequest) GetRole() OrgRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *TeamMemberRequest) GetRoleOk() (*OrgRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *TeamMemberRequest) SetRole(v OrgRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *TeamMemberRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


