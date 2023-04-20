# LoginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Passcode** | Pointer to **string** |  | [optional] 

## Methods

### NewLoginRequest

`func NewLoginRequest() *LoginRequest`

NewLoginRequest instantiates a new LoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginRequestWithDefaults

`func NewLoginRequestWithDefaults() *LoginRequest`

NewLoginRequestWithDefaults instantiates a new LoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *LoginRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LoginRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LoginRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *LoginRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPassword

`func (o *LoginRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LoginRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LoginRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *LoginRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasscode

`func (o *LoginRequest) GetPasscode() string`

GetPasscode returns the Passcode field if non-nil, zero value otherwise.

### GetPasscodeOk

`func (o *LoginRequest) GetPasscodeOk() (*string, bool)`

GetPasscodeOk returns a tuple with the Passcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscode

`func (o *LoginRequest) SetPasscode(v string)`

SetPasscode sets Passcode field to given value.

### HasPasscode

`func (o *LoginRequest) HasPasscode() bool`

HasPasscode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


