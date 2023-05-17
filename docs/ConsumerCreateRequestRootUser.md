# ConsumerCreateRequestRootUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | The first name of the root user. | 
**Surname** | **interface{}** | The last name of the root user. | 
**Email** | **interface{}** | E-mail Address of the user | 
**Mobile** | [**Mobile**](Mobile.md) |  | 
**DateOfBirth** | [**Date**](Date.md) | Date of birth of the consumer root user. | 
**Occupation** | Pointer to [**Occupation**](Occupation.md) |  | [optional] 
**Address** | Pointer to [**AddressWithNoRequiredFields**](AddressWithNoRequiredFields.md) | Address of the consumer root user. | [optional] 
**PlaceOfBirth** | Pointer to **interface{}** | Place of birth of the consumer root user. | [optional] 
**Nationality** | Pointer to **interface{}** | Nationality of the user - using ISO 3166 alpha-2. | [optional] 

## Methods

### NewConsumerCreateRequestRootUser

`func NewConsumerCreateRequestRootUser(name interface{}, surname interface{}, email interface{}, mobile Mobile, dateOfBirth Date, ) *ConsumerCreateRequestRootUser`

NewConsumerCreateRequestRootUser instantiates a new ConsumerCreateRequestRootUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerCreateRequestRootUserWithDefaults

`func NewConsumerCreateRequestRootUserWithDefaults() *ConsumerCreateRequestRootUser`

NewConsumerCreateRequestRootUserWithDefaults instantiates a new ConsumerCreateRequestRootUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConsumerCreateRequestRootUser) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsumerCreateRequestRootUser) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsumerCreateRequestRootUser) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *ConsumerCreateRequestRootUser) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConsumerCreateRequestRootUser) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSurname

`func (o *ConsumerCreateRequestRootUser) GetSurname() interface{}`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *ConsumerCreateRequestRootUser) GetSurnameOk() (*interface{}, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *ConsumerCreateRequestRootUser) SetSurname(v interface{})`

SetSurname sets Surname field to given value.


### SetSurnameNil

`func (o *ConsumerCreateRequestRootUser) SetSurnameNil(b bool)`

 SetSurnameNil sets the value for Surname to be an explicit nil

### UnsetSurname
`func (o *ConsumerCreateRequestRootUser) UnsetSurname()`

UnsetSurname ensures that no value is present for Surname, not even an explicit nil
### GetEmail

`func (o *ConsumerCreateRequestRootUser) GetEmail() interface{}`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ConsumerCreateRequestRootUser) GetEmailOk() (*interface{}, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ConsumerCreateRequestRootUser) SetEmail(v interface{})`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *ConsumerCreateRequestRootUser) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ConsumerCreateRequestRootUser) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetMobile

`func (o *ConsumerCreateRequestRootUser) GetMobile() Mobile`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *ConsumerCreateRequestRootUser) GetMobileOk() (*Mobile, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *ConsumerCreateRequestRootUser) SetMobile(v Mobile)`

SetMobile sets Mobile field to given value.


### GetDateOfBirth

`func (o *ConsumerCreateRequestRootUser) GetDateOfBirth() Date`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *ConsumerCreateRequestRootUser) GetDateOfBirthOk() (*Date, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *ConsumerCreateRequestRootUser) SetDateOfBirth(v Date)`

SetDateOfBirth sets DateOfBirth field to given value.


### GetOccupation

`func (o *ConsumerCreateRequestRootUser) GetOccupation() Occupation`

GetOccupation returns the Occupation field if non-nil, zero value otherwise.

### GetOccupationOk

`func (o *ConsumerCreateRequestRootUser) GetOccupationOk() (*Occupation, bool)`

GetOccupationOk returns a tuple with the Occupation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupation

`func (o *ConsumerCreateRequestRootUser) SetOccupation(v Occupation)`

SetOccupation sets Occupation field to given value.

### HasOccupation

`func (o *ConsumerCreateRequestRootUser) HasOccupation() bool`

HasOccupation returns a boolean if a field has been set.

### GetAddress

`func (o *ConsumerCreateRequestRootUser) GetAddress() AddressWithNoRequiredFields`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ConsumerCreateRequestRootUser) GetAddressOk() (*AddressWithNoRequiredFields, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ConsumerCreateRequestRootUser) SetAddress(v AddressWithNoRequiredFields)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ConsumerCreateRequestRootUser) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPlaceOfBirth

`func (o *ConsumerCreateRequestRootUser) GetPlaceOfBirth() interface{}`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *ConsumerCreateRequestRootUser) GetPlaceOfBirthOk() (*interface{}, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *ConsumerCreateRequestRootUser) SetPlaceOfBirth(v interface{})`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *ConsumerCreateRequestRootUser) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### SetPlaceOfBirthNil

`func (o *ConsumerCreateRequestRootUser) SetPlaceOfBirthNil(b bool)`

 SetPlaceOfBirthNil sets the value for PlaceOfBirth to be an explicit nil

### UnsetPlaceOfBirth
`func (o *ConsumerCreateRequestRootUser) UnsetPlaceOfBirth()`

UnsetPlaceOfBirth ensures that no value is present for PlaceOfBirth, not even an explicit nil
### GetNationality

`func (o *ConsumerCreateRequestRootUser) GetNationality() interface{}`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *ConsumerCreateRequestRootUser) GetNationalityOk() (*interface{}, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *ConsumerCreateRequestRootUser) SetNationality(v interface{})`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *ConsumerCreateRequestRootUser) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### SetNationalityNil

`func (o *ConsumerCreateRequestRootUser) SetNationalityNil(b bool)`

 SetNationalityNil sets the value for Nationality to be an explicit nil

### UnsetNationality
`func (o *ConsumerCreateRequestRootUser) UnsetNationality()`

UnsetNationality ensures that no value is present for Nationality, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


