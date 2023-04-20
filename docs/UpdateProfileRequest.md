# UpdateProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**User**](User.md) |  | [optional] 
**ExistingPassword** | Pointer to **string** |  | [optional] 
**Passcode** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateProfileRequest

`func NewUpdateProfileRequest() *UpdateProfileRequest`

NewUpdateProfileRequest instantiates a new UpdateProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProfileRequestWithDefaults

`func NewUpdateProfileRequestWithDefaults() *UpdateProfileRequest`

NewUpdateProfileRequestWithDefaults instantiates a new UpdateProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UpdateProfileRequest) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UpdateProfileRequest) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UpdateProfileRequest) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *UpdateProfileRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetExistingPassword

`func (o *UpdateProfileRequest) GetExistingPassword() string`

GetExistingPassword returns the ExistingPassword field if non-nil, zero value otherwise.

### GetExistingPasswordOk

`func (o *UpdateProfileRequest) GetExistingPasswordOk() (*string, bool)`

GetExistingPasswordOk returns a tuple with the ExistingPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingPassword

`func (o *UpdateProfileRequest) SetExistingPassword(v string)`

SetExistingPassword sets ExistingPassword field to given value.

### HasExistingPassword

`func (o *UpdateProfileRequest) HasExistingPassword() bool`

HasExistingPassword returns a boolean if a field has been set.

### GetPasscode

`func (o *UpdateProfileRequest) GetPasscode() string`

GetPasscode returns the Passcode field if non-nil, zero value otherwise.

### GetPasscodeOk

`func (o *UpdateProfileRequest) GetPasscodeOk() (*string, bool)`

GetPasscodeOk returns a tuple with the Passcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscode

`func (o *UpdateProfileRequest) SetPasscode(v string)`

SetPasscode sets Passcode field to given value.

### HasPasscode

`func (o *UpdateProfileRequest) HasPasscode() bool`

HasPasscode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


