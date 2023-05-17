# CorporateCreateRequestRootUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | The first name of the root user. | 
**Surname** | **interface{}** | The last name of the root user. | 
**Email** | **interface{}** | E-mail Address of the user | 
**Mobile** | [**Mobile**](Mobile.md) |  | 
**CompanyPosition** | [**CompanyPosition**](CompanyPosition.md) |  | 
**DateOfBirth** | Pointer to [**Date**](Date.md) | Date of birth of the root user. | [optional] 

## Methods

### NewCorporateCreateRequestRootUser

`func NewCorporateCreateRequestRootUser(name interface{}, surname interface{}, email interface{}, mobile Mobile, companyPosition CompanyPosition, ) *CorporateCreateRequestRootUser`

NewCorporateCreateRequestRootUser instantiates a new CorporateCreateRequestRootUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporateCreateRequestRootUserWithDefaults

`func NewCorporateCreateRequestRootUserWithDefaults() *CorporateCreateRequestRootUser`

NewCorporateCreateRequestRootUserWithDefaults instantiates a new CorporateCreateRequestRootUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CorporateCreateRequestRootUser) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CorporateCreateRequestRootUser) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CorporateCreateRequestRootUser) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *CorporateCreateRequestRootUser) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CorporateCreateRequestRootUser) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSurname

`func (o *CorporateCreateRequestRootUser) GetSurname() interface{}`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *CorporateCreateRequestRootUser) GetSurnameOk() (*interface{}, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *CorporateCreateRequestRootUser) SetSurname(v interface{})`

SetSurname sets Surname field to given value.


### SetSurnameNil

`func (o *CorporateCreateRequestRootUser) SetSurnameNil(b bool)`

 SetSurnameNil sets the value for Surname to be an explicit nil

### UnsetSurname
`func (o *CorporateCreateRequestRootUser) UnsetSurname()`

UnsetSurname ensures that no value is present for Surname, not even an explicit nil
### GetEmail

`func (o *CorporateCreateRequestRootUser) GetEmail() interface{}`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CorporateCreateRequestRootUser) GetEmailOk() (*interface{}, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CorporateCreateRequestRootUser) SetEmail(v interface{})`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *CorporateCreateRequestRootUser) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CorporateCreateRequestRootUser) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetMobile

`func (o *CorporateCreateRequestRootUser) GetMobile() Mobile`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *CorporateCreateRequestRootUser) GetMobileOk() (*Mobile, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *CorporateCreateRequestRootUser) SetMobile(v Mobile)`

SetMobile sets Mobile field to given value.


### GetCompanyPosition

`func (o *CorporateCreateRequestRootUser) GetCompanyPosition() CompanyPosition`

GetCompanyPosition returns the CompanyPosition field if non-nil, zero value otherwise.

### GetCompanyPositionOk

`func (o *CorporateCreateRequestRootUser) GetCompanyPositionOk() (*CompanyPosition, bool)`

GetCompanyPositionOk returns a tuple with the CompanyPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyPosition

`func (o *CorporateCreateRequestRootUser) SetCompanyPosition(v CompanyPosition)`

SetCompanyPosition sets CompanyPosition field to given value.


### GetDateOfBirth

`func (o *CorporateCreateRequestRootUser) GetDateOfBirth() Date`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *CorporateCreateRequestRootUser) GetDateOfBirthOk() (*Date, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *CorporateCreateRequestRootUser) SetDateOfBirth(v Date)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *CorporateCreateRequestRootUser) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


