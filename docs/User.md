# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**HasMfa** | Pointer to **bool** |  | [optional] 
**LastEulaAccepted** | Pointer to **int32** |  | [optional] 
**VerifiedPhone** | Pointer to **bool** |  | [optional] 
**DefaultTeamId** | Pointer to **string** |  | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *User) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *User) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *User) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *User) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *User) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPhone

`func (o *User) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *User) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *User) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *User) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetHasMfa

`func (o *User) GetHasMfa() bool`

GetHasMfa returns the HasMfa field if non-nil, zero value otherwise.

### GetHasMfaOk

`func (o *User) GetHasMfaOk() (*bool, bool)`

GetHasMfaOk returns a tuple with the HasMfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMfa

`func (o *User) SetHasMfa(v bool)`

SetHasMfa sets HasMfa field to given value.

### HasHasMfa

`func (o *User) HasHasMfa() bool`

HasHasMfa returns a boolean if a field has been set.

### GetLastEulaAccepted

`func (o *User) GetLastEulaAccepted() int32`

GetLastEulaAccepted returns the LastEulaAccepted field if non-nil, zero value otherwise.

### GetLastEulaAcceptedOk

`func (o *User) GetLastEulaAcceptedOk() (*int32, bool)`

GetLastEulaAcceptedOk returns a tuple with the LastEulaAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEulaAccepted

`func (o *User) SetLastEulaAccepted(v int32)`

SetLastEulaAccepted sets LastEulaAccepted field to given value.

### HasLastEulaAccepted

`func (o *User) HasLastEulaAccepted() bool`

HasLastEulaAccepted returns a boolean if a field has been set.

### GetVerifiedPhone

`func (o *User) GetVerifiedPhone() bool`

GetVerifiedPhone returns the VerifiedPhone field if non-nil, zero value otherwise.

### GetVerifiedPhoneOk

`func (o *User) GetVerifiedPhoneOk() (*bool, bool)`

GetVerifiedPhoneOk returns a tuple with the VerifiedPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedPhone

`func (o *User) SetVerifiedPhone(v bool)`

SetVerifiedPhone sets VerifiedPhone field to given value.

### HasVerifiedPhone

`func (o *User) HasVerifiedPhone() bool`

HasVerifiedPhone returns a boolean if a field has been set.

### GetDefaultTeamId

`func (o *User) GetDefaultTeamId() string`

GetDefaultTeamId returns the DefaultTeamId field if non-nil, zero value otherwise.

### GetDefaultTeamIdOk

`func (o *User) GetDefaultTeamIdOk() (*string, bool)`

GetDefaultTeamIdOk returns a tuple with the DefaultTeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTeamId

`func (o *User) SetDefaultTeamId(v string)`

SetDefaultTeamId sets DefaultTeamId field to given value.

### HasDefaultTeamId

`func (o *User) HasDefaultTeamId() bool`

HasDefaultTeamId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


