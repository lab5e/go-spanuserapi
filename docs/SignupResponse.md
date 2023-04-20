# SignupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**SignupResponseResult**](SignupResponseResult.md) |  | [optional] [default to SIGNUPRESPONSERESULT_UNSPECIFIED]

## Methods

### NewSignupResponse

`func NewSignupResponse() *SignupResponse`

NewSignupResponse instantiates a new SignupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignupResponseWithDefaults

`func NewSignupResponseWithDefaults() *SignupResponse`

NewSignupResponseWithDefaults instantiates a new SignupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *SignupResponse) GetResult() SignupResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SignupResponse) GetResultOk() (*SignupResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SignupResponse) SetResult(v SignupResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *SignupResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


