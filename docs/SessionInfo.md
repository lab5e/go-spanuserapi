# SessionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **string** |  | [optional] 
**End** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Browser** | Pointer to **string** |  | [optional] 
**Device** | Pointer to [**SessionInfoType**](SessionInfoType.md) |  | [optional] [default to SESSIONINFOTYPE_UNSPECIFIED]
**CurrentSession** | Pointer to **bool** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 

## Methods

### NewSessionInfo

`func NewSessionInfo() *SessionInfo`

NewSessionInfo instantiates a new SessionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionInfoWithDefaults

`func NewSessionInfoWithDefaults() *SessionInfo`

NewSessionInfoWithDefaults instantiates a new SessionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *SessionInfo) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SessionInfo) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SessionInfo) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *SessionInfo) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *SessionInfo) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SessionInfo) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SessionInfo) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *SessionInfo) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetIp

`func (o *SessionInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *SessionInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *SessionInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *SessionInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLocation

`func (o *SessionInfo) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SessionInfo) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SessionInfo) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SessionInfo) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPlatform

`func (o *SessionInfo) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *SessionInfo) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *SessionInfo) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *SessionInfo) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetBrowser

`func (o *SessionInfo) GetBrowser() string`

GetBrowser returns the Browser field if non-nil, zero value otherwise.

### GetBrowserOk

`func (o *SessionInfo) GetBrowserOk() (*string, bool)`

GetBrowserOk returns a tuple with the Browser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowser

`func (o *SessionInfo) SetBrowser(v string)`

SetBrowser sets Browser field to given value.

### HasBrowser

`func (o *SessionInfo) HasBrowser() bool`

HasBrowser returns a boolean if a field has been set.

### GetDevice

`func (o *SessionInfo) GetDevice() SessionInfoType`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *SessionInfo) GetDeviceOk() (*SessionInfoType, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *SessionInfo) SetDevice(v SessionInfoType)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *SessionInfo) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetCurrentSession

`func (o *SessionInfo) GetCurrentSession() bool`

GetCurrentSession returns the CurrentSession field if non-nil, zero value otherwise.

### GetCurrentSessionOk

`func (o *SessionInfo) GetCurrentSessionOk() (*bool, bool)`

GetCurrentSessionOk returns a tuple with the CurrentSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSession

`func (o *SessionInfo) SetCurrentSession(v bool)`

SetCurrentSession sets CurrentSession field to given value.

### HasCurrentSession

`func (o *SessionInfo) HasCurrentSession() bool`

HasCurrentSession returns a boolean if a field has been set.

### GetActive

`func (o *SessionInfo) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SessionInfo) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SessionInfo) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SessionInfo) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


