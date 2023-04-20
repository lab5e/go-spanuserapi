# ListAPITokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tokens** | Pointer to [**[]APIToken**](APIToken.md) |  | [optional] 

## Methods

### NewListAPITokenResponse

`func NewListAPITokenResponse() *ListAPITokenResponse`

NewListAPITokenResponse instantiates a new ListAPITokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAPITokenResponseWithDefaults

`func NewListAPITokenResponseWithDefaults() *ListAPITokenResponse`

NewListAPITokenResponseWithDefaults instantiates a new ListAPITokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokens

`func (o *ListAPITokenResponse) GetTokens() []APIToken`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *ListAPITokenResponse) GetTokensOk() (*[]APIToken, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *ListAPITokenResponse) SetTokens(v []APIToken)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *ListAPITokenResponse) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


