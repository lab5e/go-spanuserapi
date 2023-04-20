# ListSessionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sessions** | Pointer to [**[]SessionInfo**](SessionInfo.md) |  | [optional] 

## Methods

### NewListSessionsResponse

`func NewListSessionsResponse() *ListSessionsResponse`

NewListSessionsResponse instantiates a new ListSessionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSessionsResponseWithDefaults

`func NewListSessionsResponseWithDefaults() *ListSessionsResponse`

NewListSessionsResponseWithDefaults instantiates a new ListSessionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessions

`func (o *ListSessionsResponse) GetSessions() []SessionInfo`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *ListSessionsResponse) GetSessionsOk() (*[]SessionInfo, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *ListSessionsResponse) SetSessions(v []SessionInfo)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *ListSessionsResponse) HasSessions() bool`

HasSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


