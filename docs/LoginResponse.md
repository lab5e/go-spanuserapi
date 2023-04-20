# LoginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**LoginResponseResult**](LoginResponseResult.md) |  | [optional] [default to LOGINRESPONSERESULT_UNSPECIFIED]
**Credentials** | Pointer to **string** |  | [optional] 

## Methods

### NewLoginResponse

`func NewLoginResponse() *LoginResponse`

NewLoginResponse instantiates a new LoginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginResponseWithDefaults

`func NewLoginResponseWithDefaults() *LoginResponse`

NewLoginResponseWithDefaults instantiates a new LoginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *LoginResponse) GetResult() LoginResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *LoginResponse) GetResultOk() (*LoginResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *LoginResponse) SetResult(v LoginResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *LoginResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetCredentials

`func (o *LoginResponse) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *LoginResponse) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *LoginResponse) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *LoginResponse) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


