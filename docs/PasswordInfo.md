# PasswordInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityId** | [**IdentityId**](IdentityId.md) | The identity to which the user&#39;s password information belongs to. | 
**ExpiryDate** | Pointer to **interface{}** | The millisecond timestamp indicating when the password will expire. If 0, then this password will not expire. | [optional] 

## Methods

### NewPasswordInfo

`func NewPasswordInfo(identityId IdentityId, ) *PasswordInfo`

NewPasswordInfo instantiates a new PasswordInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordInfoWithDefaults

`func NewPasswordInfoWithDefaults() *PasswordInfo`

NewPasswordInfoWithDefaults instantiates a new PasswordInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityId

`func (o *PasswordInfo) GetIdentityId() IdentityId`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *PasswordInfo) GetIdentityIdOk() (*IdentityId, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *PasswordInfo) SetIdentityId(v IdentityId)`

SetIdentityId sets IdentityId field to given value.


### GetExpiryDate

`func (o *PasswordInfo) GetExpiryDate() interface{}`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *PasswordInfo) GetExpiryDateOk() (*interface{}, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *PasswordInfo) SetExpiryDate(v interface{})`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *PasswordInfo) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *PasswordInfo) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *PasswordInfo) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


