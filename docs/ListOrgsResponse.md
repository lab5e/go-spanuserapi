# ListOrgsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orgs** | Pointer to [**[]Organisation**](Organisation.md) |  | [optional] 

## Methods

### NewListOrgsResponse

`func NewListOrgsResponse() *ListOrgsResponse`

NewListOrgsResponse instantiates a new ListOrgsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrgsResponseWithDefaults

`func NewListOrgsResponseWithDefaults() *ListOrgsResponse`

NewListOrgsResponseWithDefaults instantiates a new ListOrgsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgs

`func (o *ListOrgsResponse) GetOrgs() []Organisation`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *ListOrgsResponse) GetOrgsOk() (*[]Organisation, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *ListOrgsResponse) SetOrgs(v []Organisation)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *ListOrgsResponse) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


