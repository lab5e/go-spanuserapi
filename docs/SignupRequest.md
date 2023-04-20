# SignupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Name is a required field. | [optional] 
**Address** | Pointer to **string** | Address is a free-form field and may be left blank. | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**AcceptedEulaVersion** | Pointer to **int32** | parentheses are ignored.  This version must be set to the current EULA version | [optional] 

## Methods

### NewSignupRequest

`func NewSignupRequest() *SignupRequest`

NewSignupRequest instantiates a new SignupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignupRequestWithDefaults

`func NewSignupRequestWithDefaults() *SignupRequest`

NewSignupRequestWithDefaults instantiates a new SignupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SignupRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SignupRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SignupRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SignupRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *SignupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SignupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SignupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SignupRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *SignupRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SignupRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SignupRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SignupRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPhone

`func (o *SignupRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SignupRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SignupRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SignupRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAcceptedEulaVersion

`func (o *SignupRequest) GetAcceptedEulaVersion() int32`

GetAcceptedEulaVersion returns the AcceptedEulaVersion field if non-nil, zero value otherwise.

### GetAcceptedEulaVersionOk

`func (o *SignupRequest) GetAcceptedEulaVersionOk() (*int32, bool)`

GetAcceptedEulaVersionOk returns a tuple with the AcceptedEulaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedEulaVersion

`func (o *SignupRequest) SetAcceptedEulaVersion(v int32)`

SetAcceptedEulaVersion sets AcceptedEulaVersion field to given value.

### HasAcceptedEulaVersion

`func (o *SignupRequest) HasAcceptedEulaVersion() bool`

HasAcceptedEulaVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


