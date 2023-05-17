# ConsumerRootUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**IdentityId**](IdentityId.md) |  | 
**Name** | **interface{}** | First name of the root user. | 
**Surname** | **interface{}** | Last name of the root user. | 
**Email** | **interface{}** | E-mail Address of the user | 
**Mobile** | [**Mobile**](Mobile.md) |  | 
**Occupation** | Pointer to [**Occupation**](Occupation.md) |  | [optional] 
**Active** | **interface{}** | The state of the root user. If false, then the user will not be able to log in. | 
**EmailVerified** | **interface{}** | Indicates if the root user&#39;s email has been verified. | 
**MobileNumberVerified** | **interface{}** | Indicates if the root user&#39;s mobile number has been verified. | 
**DateOfBirth** | Pointer to [**Date**](Date.md) | Date of birth of the consumer root user. | [optional] 
**Address** | Pointer to [**AddressWithNoRequiredFields**](AddressWithNoRequiredFields.md) | Address of the consumer root user. | [optional] 
**Nationality** | Pointer to **interface{}** | Nationality of the user - using ISO 3166 alpha-2. | [optional] 
**PlaceOfBirth** | Pointer to **interface{}** | The place of birth of the consumer root user. | [optional] 

## Methods

### NewConsumerRootUser

`func NewConsumerRootUser(id IdentityId, name interface{}, surname interface{}, email interface{}, mobile Mobile, active interface{}, emailVerified interface{}, mobileNumberVerified interface{}, ) *ConsumerRootUser`

NewConsumerRootUser instantiates a new ConsumerRootUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerRootUserWithDefaults

`func NewConsumerRootUserWithDefaults() *ConsumerRootUser`

NewConsumerRootUserWithDefaults instantiates a new ConsumerRootUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsumerRootUser) GetId() IdentityId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsumerRootUser) GetIdOk() (*IdentityId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsumerRootUser) SetId(v IdentityId)`

SetId sets Id field to given value.


### GetName

`func (o *ConsumerRootUser) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsumerRootUser) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsumerRootUser) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *ConsumerRootUser) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConsumerRootUser) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSurname

`func (o *ConsumerRootUser) GetSurname() interface{}`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *ConsumerRootUser) GetSurnameOk() (*interface{}, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *ConsumerRootUser) SetSurname(v interface{})`

SetSurname sets Surname field to given value.


### SetSurnameNil

`func (o *ConsumerRootUser) SetSurnameNil(b bool)`

 SetSurnameNil sets the value for Surname to be an explicit nil

### UnsetSurname
`func (o *ConsumerRootUser) UnsetSurname()`

UnsetSurname ensures that no value is present for Surname, not even an explicit nil
### GetEmail

`func (o *ConsumerRootUser) GetEmail() interface{}`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ConsumerRootUser) GetEmailOk() (*interface{}, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ConsumerRootUser) SetEmail(v interface{})`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *ConsumerRootUser) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ConsumerRootUser) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetMobile

`func (o *ConsumerRootUser) GetMobile() Mobile`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *ConsumerRootUser) GetMobileOk() (*Mobile, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *ConsumerRootUser) SetMobile(v Mobile)`

SetMobile sets Mobile field to given value.


### GetOccupation

`func (o *ConsumerRootUser) GetOccupation() Occupation`

GetOccupation returns the Occupation field if non-nil, zero value otherwise.

### GetOccupationOk

`func (o *ConsumerRootUser) GetOccupationOk() (*Occupation, bool)`

GetOccupationOk returns a tuple with the Occupation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupation

`func (o *ConsumerRootUser) SetOccupation(v Occupation)`

SetOccupation sets Occupation field to given value.

### HasOccupation

`func (o *ConsumerRootUser) HasOccupation() bool`

HasOccupation returns a boolean if a field has been set.

### GetActive

`func (o *ConsumerRootUser) GetActive() interface{}`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ConsumerRootUser) GetActiveOk() (*interface{}, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ConsumerRootUser) SetActive(v interface{})`

SetActive sets Active field to given value.


### SetActiveNil

`func (o *ConsumerRootUser) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *ConsumerRootUser) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetEmailVerified

`func (o *ConsumerRootUser) GetEmailVerified() interface{}`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *ConsumerRootUser) GetEmailVerifiedOk() (*interface{}, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *ConsumerRootUser) SetEmailVerified(v interface{})`

SetEmailVerified sets EmailVerified field to given value.


### SetEmailVerifiedNil

`func (o *ConsumerRootUser) SetEmailVerifiedNil(b bool)`

 SetEmailVerifiedNil sets the value for EmailVerified to be an explicit nil

### UnsetEmailVerified
`func (o *ConsumerRootUser) UnsetEmailVerified()`

UnsetEmailVerified ensures that no value is present for EmailVerified, not even an explicit nil
### GetMobileNumberVerified

`func (o *ConsumerRootUser) GetMobileNumberVerified() interface{}`

GetMobileNumberVerified returns the MobileNumberVerified field if non-nil, zero value otherwise.

### GetMobileNumberVerifiedOk

`func (o *ConsumerRootUser) GetMobileNumberVerifiedOk() (*interface{}, bool)`

GetMobileNumberVerifiedOk returns a tuple with the MobileNumberVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumberVerified

`func (o *ConsumerRootUser) SetMobileNumberVerified(v interface{})`

SetMobileNumberVerified sets MobileNumberVerified field to given value.


### SetMobileNumberVerifiedNil

`func (o *ConsumerRootUser) SetMobileNumberVerifiedNil(b bool)`

 SetMobileNumberVerifiedNil sets the value for MobileNumberVerified to be an explicit nil

### UnsetMobileNumberVerified
`func (o *ConsumerRootUser) UnsetMobileNumberVerified()`

UnsetMobileNumberVerified ensures that no value is present for MobileNumberVerified, not even an explicit nil
### GetDateOfBirth

`func (o *ConsumerRootUser) GetDateOfBirth() Date`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *ConsumerRootUser) GetDateOfBirthOk() (*Date, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *ConsumerRootUser) SetDateOfBirth(v Date)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *ConsumerRootUser) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetAddress

`func (o *ConsumerRootUser) GetAddress() AddressWithNoRequiredFields`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ConsumerRootUser) GetAddressOk() (*AddressWithNoRequiredFields, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ConsumerRootUser) SetAddress(v AddressWithNoRequiredFields)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ConsumerRootUser) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetNationality

`func (o *ConsumerRootUser) GetNationality() interface{}`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *ConsumerRootUser) GetNationalityOk() (*interface{}, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *ConsumerRootUser) SetNationality(v interface{})`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *ConsumerRootUser) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### SetNationalityNil

`func (o *ConsumerRootUser) SetNationalityNil(b bool)`

 SetNationalityNil sets the value for Nationality to be an explicit nil

### UnsetNationality
`func (o *ConsumerRootUser) UnsetNationality()`

UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
### GetPlaceOfBirth

`func (o *ConsumerRootUser) GetPlaceOfBirth() interface{}`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *ConsumerRootUser) GetPlaceOfBirthOk() (*interface{}, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *ConsumerRootUser) SetPlaceOfBirth(v interface{})`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *ConsumerRootUser) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### SetPlaceOfBirthNil

`func (o *ConsumerRootUser) SetPlaceOfBirthNil(b bool)`

 SetPlaceOfBirthNil sets the value for PlaceOfBirth to be an explicit nil

### UnsetPlaceOfBirth
`func (o *ConsumerRootUser) UnsetPlaceOfBirth()`

UnsetPlaceOfBirth ensures that no value is present for PlaceOfBirth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


