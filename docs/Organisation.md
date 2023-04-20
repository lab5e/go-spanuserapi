# Organisation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganisationId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**Personal** | Pointer to **bool** |  | [optional] 
**Stats** | Pointer to [**OrgStats**](OrgStats.md) |  | [optional] 

## Methods

### NewOrganisation

`func NewOrganisation() *Organisation`

NewOrganisation instantiates a new Organisation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganisationWithDefaults

`func NewOrganisationWithDefaults() *Organisation`

NewOrganisationWithDefaults instantiates a new Organisation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganisationId

`func (o *Organisation) GetOrganisationId() string`

GetOrganisationId returns the OrganisationId field if non-nil, zero value otherwise.

### GetOrganisationIdOk

`func (o *Organisation) GetOrganisationIdOk() (*string, bool)`

GetOrganisationIdOk returns a tuple with the OrganisationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisationId

`func (o *Organisation) SetOrganisationId(v string)`

SetOrganisationId sets OrganisationId field to given value.

### HasOrganisationId

`func (o *Organisation) HasOrganisationId() bool`

HasOrganisationId returns a boolean if a field has been set.

### GetName

`func (o *Organisation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Organisation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Organisation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Organisation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreated

`func (o *Organisation) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Organisation) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Organisation) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Organisation) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetPersonal

`func (o *Organisation) GetPersonal() bool`

GetPersonal returns the Personal field if non-nil, zero value otherwise.

### GetPersonalOk

`func (o *Organisation) GetPersonalOk() (*bool, bool)`

GetPersonalOk returns a tuple with the Personal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonal

`func (o *Organisation) SetPersonal(v bool)`

SetPersonal sets Personal field to given value.

### HasPersonal

`func (o *Organisation) HasPersonal() bool`

HasPersonal returns a boolean if a field has been set.

### GetStats

`func (o *Organisation) GetStats() OrgStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Organisation) GetStatsOk() (*OrgStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Organisation) SetStats(v OrgStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *Organisation) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


