# UpdateAPITokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | Pointer to **string** |  | [optional] 
**Write** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NewTeamId** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateAPITokenRequest

`func NewUpdateAPITokenRequest() *UpdateAPITokenRequest`

NewUpdateAPITokenRequest instantiates a new UpdateAPITokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAPITokenRequestWithDefaults

`func NewUpdateAPITokenRequestWithDefaults() *UpdateAPITokenRequest`

NewUpdateAPITokenRequestWithDefaults instantiates a new UpdateAPITokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *UpdateAPITokenRequest) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *UpdateAPITokenRequest) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *UpdateAPITokenRequest) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *UpdateAPITokenRequest) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetWrite

`func (o *UpdateAPITokenRequest) GetWrite() bool`

GetWrite returns the Write field if non-nil, zero value otherwise.

### GetWriteOk

`func (o *UpdateAPITokenRequest) GetWriteOk() (*bool, bool)`

GetWriteOk returns a tuple with the Write field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrite

`func (o *UpdateAPITokenRequest) SetWrite(v bool)`

SetWrite sets Write field to given value.

### HasWrite

`func (o *UpdateAPITokenRequest) HasWrite() bool`

HasWrite returns a boolean if a field has been set.

### GetName

`func (o *UpdateAPITokenRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAPITokenRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAPITokenRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAPITokenRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNewTeamId

`func (o *UpdateAPITokenRequest) GetNewTeamId() string`

GetNewTeamId returns the NewTeamId field if non-nil, zero value otherwise.

### GetNewTeamIdOk

`func (o *UpdateAPITokenRequest) GetNewTeamIdOk() (*string, bool)`

GetNewTeamIdOk returns a tuple with the NewTeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTeamId

`func (o *UpdateAPITokenRequest) SetNewTeamId(v string)`

SetNewTeamId sets NewTeamId field to given value.

### HasNewTeamId

`func (o *UpdateAPITokenRequest) HasNewTeamId() bool`

HasNewTeamId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


