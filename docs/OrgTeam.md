# OrgTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeamId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**Members** | Pointer to [**[]OrgTeamMember**](OrgTeamMember.md) | Members of team. | [optional] 

## Methods

### NewOrgTeam

`func NewOrgTeam() *OrgTeam`

NewOrgTeam instantiates a new OrgTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgTeamWithDefaults

`func NewOrgTeamWithDefaults() *OrgTeam`

NewOrgTeamWithDefaults instantiates a new OrgTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeamId

`func (o *OrgTeam) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *OrgTeam) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *OrgTeam) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *OrgTeam) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetName

`func (o *OrgTeam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgTeam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgTeam) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrgTeam) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreated

`func (o *OrgTeam) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OrgTeam) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OrgTeam) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *OrgTeam) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetMembers

`func (o *OrgTeam) GetMembers() []OrgTeamMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *OrgTeam) GetMembersOk() (*[]OrgTeamMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *OrgTeam) SetMembers(v []OrgTeamMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *OrgTeam) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


