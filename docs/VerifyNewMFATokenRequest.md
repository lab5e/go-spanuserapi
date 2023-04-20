# VerifyNewMFATokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Passcode** | Pointer to **string** |  | [optional] 

## Methods

### NewVerifyNewMFATokenRequest

`func NewVerifyNewMFATokenRequest() *VerifyNewMFATokenRequest`

NewVerifyNewMFATokenRequest instantiates a new VerifyNewMFATokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyNewMFATokenRequestWithDefaults

`func NewVerifyNewMFATokenRequestWithDefaults() *VerifyNewMFATokenRequest`

NewVerifyNewMFATokenRequestWithDefaults instantiates a new VerifyNewMFATokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPasscode

`func (o *VerifyNewMFATokenRequest) GetPasscode() string`

GetPasscode returns the Passcode field if non-nil, zero value otherwise.

### GetPasscodeOk

`func (o *VerifyNewMFATokenRequest) GetPasscodeOk() (*string, bool)`

GetPasscodeOk returns a tuple with the Passcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscode

`func (o *VerifyNewMFATokenRequest) SetPasscode(v string)`

SetPasscode sets Passcode field to given value.

### HasPasscode

`func (o *VerifyNewMFATokenRequest) HasPasscode() bool`

HasPasscode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


