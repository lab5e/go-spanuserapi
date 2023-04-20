# UpdatePasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExistingPassword** | Pointer to **string** |  | [optional] 
**NewPassword** | Pointer to **string** |  | [optional] 
**Passcode** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdatePasswordRequest

`func NewUpdatePasswordRequest() *UpdatePasswordRequest`

NewUpdatePasswordRequest instantiates a new UpdatePasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePasswordRequestWithDefaults

`func NewUpdatePasswordRequestWithDefaults() *UpdatePasswordRequest`

NewUpdatePasswordRequestWithDefaults instantiates a new UpdatePasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExistingPassword

`func (o *UpdatePasswordRequest) GetExistingPassword() string`

GetExistingPassword returns the ExistingPassword field if non-nil, zero value otherwise.

### GetExistingPasswordOk

`func (o *UpdatePasswordRequest) GetExistingPasswordOk() (*string, bool)`

GetExistingPasswordOk returns a tuple with the ExistingPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingPassword

`func (o *UpdatePasswordRequest) SetExistingPassword(v string)`

SetExistingPassword sets ExistingPassword field to given value.

### HasExistingPassword

`func (o *UpdatePasswordRequest) HasExistingPassword() bool`

HasExistingPassword returns a boolean if a field has been set.

### GetNewPassword

`func (o *UpdatePasswordRequest) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *UpdatePasswordRequest) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *UpdatePasswordRequest) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *UpdatePasswordRequest) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.

### GetPasscode

`func (o *UpdatePasswordRequest) GetPasscode() string`

GetPasscode returns the Passcode field if non-nil, zero value otherwise.

### GetPasscodeOk

`func (o *UpdatePasswordRequest) GetPasscodeOk() (*string, bool)`

GetPasscodeOk returns a tuple with the Passcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscode

`func (o *UpdatePasswordRequest) SetPasscode(v string)`

SetPasscode sets Passcode field to given value.

### HasPasscode

`func (o *UpdatePasswordRequest) HasPasscode() bool`

HasPasscode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


