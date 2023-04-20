# AddAPITokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Write** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddAPITokenRequest

`func NewAddAPITokenRequest() *AddAPITokenRequest`

NewAddAPITokenRequest instantiates a new AddAPITokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAPITokenRequestWithDefaults

`func NewAddAPITokenRequestWithDefaults() *AddAPITokenRequest`

NewAddAPITokenRequestWithDefaults instantiates a new AddAPITokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *AddAPITokenRequest) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AddAPITokenRequest) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AddAPITokenRequest) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *AddAPITokenRequest) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetName

`func (o *AddAPITokenRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddAPITokenRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddAPITokenRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddAPITokenRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWrite

`func (o *AddAPITokenRequest) GetWrite() bool`

GetWrite returns the Write field if non-nil, zero value otherwise.

### GetWriteOk

`func (o *AddAPITokenRequest) GetWriteOk() (*bool, bool)`

GetWriteOk returns a tuple with the Write field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrite

`func (o *AddAPITokenRequest) SetWrite(v bool)`

SetWrite sets Write field to given value.

### HasWrite

`func (o *AddAPITokenRequest) HasWrite() bool`

HasWrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


