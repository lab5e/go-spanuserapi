# RecoverPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**RecoverToken** | Pointer to **string** |  | [optional] 
**Passcode** | Pointer to **string** |  | [optional] 
**NewPassword** | Pointer to **string** |  | [optional] 

## Methods

### NewRecoverPasswordRequest

`func NewRecoverPasswordRequest() *RecoverPasswordRequest`

NewRecoverPasswordRequest instantiates a new RecoverPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverPasswordRequestWithDefaults

`func NewRecoverPasswordRequestWithDefaults() *RecoverPasswordRequest`

NewRecoverPasswordRequestWithDefaults instantiates a new RecoverPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *RecoverPasswordRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RecoverPasswordRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RecoverPasswordRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *RecoverPasswordRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRecoverToken

`func (o *RecoverPasswordRequest) GetRecoverToken() string`

GetRecoverToken returns the RecoverToken field if non-nil, zero value otherwise.

### GetRecoverTokenOk

`func (o *RecoverPasswordRequest) GetRecoverTokenOk() (*string, bool)`

GetRecoverTokenOk returns a tuple with the RecoverToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToken

`func (o *RecoverPasswordRequest) SetRecoverToken(v string)`

SetRecoverToken sets RecoverToken field to given value.

### HasRecoverToken

`func (o *RecoverPasswordRequest) HasRecoverToken() bool`

HasRecoverToken returns a boolean if a field has been set.

### GetPasscode

`func (o *RecoverPasswordRequest) GetPasscode() string`

GetPasscode returns the Passcode field if non-nil, zero value otherwise.

### GetPasscodeOk

`func (o *RecoverPasswordRequest) GetPasscodeOk() (*string, bool)`

GetPasscodeOk returns a tuple with the Passcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscode

`func (o *RecoverPasswordRequest) SetPasscode(v string)`

SetPasscode sets Passcode field to given value.

### HasPasscode

`func (o *RecoverPasswordRequest) HasPasscode() bool`

HasPasscode returns a boolean if a field has been set.

### GetNewPassword

`func (o *RecoverPasswordRequest) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *RecoverPasswordRequest) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *RecoverPasswordRequest) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *RecoverPasswordRequest) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


