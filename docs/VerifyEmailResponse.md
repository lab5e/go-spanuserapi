# VerifyEmailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**VerifyEmailResponseResult**](VerifyEmailResponseResult.md) |  | [optional] [default to VERIFYEMAILRESPONSERESULT_UNSPECIFIED]
**Credentials** | Pointer to **string** |  | [optional] 

## Methods

### NewVerifyEmailResponse

`func NewVerifyEmailResponse() *VerifyEmailResponse`

NewVerifyEmailResponse instantiates a new VerifyEmailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyEmailResponseWithDefaults

`func NewVerifyEmailResponseWithDefaults() *VerifyEmailResponse`

NewVerifyEmailResponseWithDefaults instantiates a new VerifyEmailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *VerifyEmailResponse) GetResult() VerifyEmailResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *VerifyEmailResponse) GetResultOk() (*VerifyEmailResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *VerifyEmailResponse) SetResult(v VerifyEmailResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *VerifyEmailResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetCredentials

`func (o *VerifyEmailResponse) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *VerifyEmailResponse) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *VerifyEmailResponse) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *VerifyEmailResponse) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


