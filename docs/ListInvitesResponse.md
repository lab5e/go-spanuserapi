# ListInvitesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invites** | Pointer to [**[]Invite**](Invite.md) |  | [optional] 

## Methods

### NewListInvitesResponse

`func NewListInvitesResponse() *ListInvitesResponse`

NewListInvitesResponse instantiates a new ListInvitesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInvitesResponseWithDefaults

`func NewListInvitesResponseWithDefaults() *ListInvitesResponse`

NewListInvitesResponseWithDefaults instantiates a new ListInvitesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvites

`func (o *ListInvitesResponse) GetInvites() []Invite`

GetInvites returns the Invites field if non-nil, zero value otherwise.

### GetInvitesOk

`func (o *ListInvitesResponse) GetInvitesOk() (*[]Invite, bool)`

GetInvitesOk returns a tuple with the Invites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvites

`func (o *ListInvitesResponse) SetInvites(v []Invite)`

SetInvites sets Invites field to given value.

### HasInvites

`func (o *ListInvitesResponse) HasInvites() bool`

HasInvites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


