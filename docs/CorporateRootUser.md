# CorporateRootUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**IdentityId**](IdentityId.md) |  | 
**Name** | **interface{}** | First name of the root user. | 
**Surname** | **interface{}** | Last name of the root user. | 
**Email** | **interface{}** | E-mail Address of the user | 
**Mobile** | [**Mobile**](Mobile.md) |  | 
**CompanyPosition** | [**CompanyPosition**](CompanyPosition.md) |  | 
**Active** | **interface{}** | The state of the root user. If false, then the user will not be able to log in. | 
**EmailVerified** | **interface{}** | Indicates if the root user&#39;s email has been verified. | 
**MobileNumberVerified** | **interface{}** | Indicates if the root user&#39;s mobile number has been verified. | 
**DateOfBirth** | Pointer to [**Date**](Date.md) | Date of birth of the authorised user. | [optional] 

## Methods

### NewCorporateRootUser

`func NewCorporateRootUser(id IdentityId, name interface{}, surname interface{}, email interface{}, mobile Mobile, companyPosition CompanyPosition, active interface{}, emailVerified interface{}, mobileNumberVerified interface{}, ) *CorporateRootUser`

NewCorporateRootUser instantiates a new CorporateRootUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporateRootUserWithDefaults

`func NewCorporateRootUserWithDefaults() *CorporateRootUser`

NewCorporateRootUserWithDefaults instantiates a new CorporateRootUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CorporateRootUser) GetId() IdentityId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CorporateRootUser) GetIdOk() (*IdentityId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CorporateRootUser) SetId(v IdentityId)`

SetId sets Id field to given value.


### GetName

`func (o *CorporateRootUser) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CorporateRootUser) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CorporateRootUser) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *CorporateRootUser) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CorporateRootUser) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSurname

`func (o *CorporateRootUser) GetSurname() interface{}`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *CorporateRootUser) GetSurnameOk() (*interface{}, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *CorporateRootUser) SetSurname(v interface{})`

SetSurname sets Surname field to given value.


### SetSurnameNil

`func (o *CorporateRootUser) SetSurnameNil(b bool)`

 SetSurnameNil sets the value for Surname to be an explicit nil

### UnsetSurname
`func (o *CorporateRootUser) UnsetSurname()`

UnsetSurname ensures that no value is present for Surname, not even an explicit nil
### GetEmail

`func (o *CorporateRootUser) GetEmail() interface{}`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CorporateRootUser) GetEmailOk() (*interface{}, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CorporateRootUser) SetEmail(v interface{})`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *CorporateRootUser) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CorporateRootUser) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetMobile

`func (o *CorporateRootUser) GetMobile() Mobile`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *CorporateRootUser) GetMobileOk() (*Mobile, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *CorporateRootUser) SetMobile(v Mobile)`

SetMobile sets Mobile field to given value.


### GetCompanyPosition

`func (o *CorporateRootUser) GetCompanyPosition() CompanyPosition`

GetCompanyPosition returns the CompanyPosition field if non-nil, zero value otherwise.

### GetCompanyPositionOk

`func (o *CorporateRootUser) GetCompanyPositionOk() (*CompanyPosition, bool)`

GetCompanyPositionOk returns a tuple with the CompanyPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyPosition

`func (o *CorporateRootUser) SetCompanyPosition(v CompanyPosition)`

SetCompanyPosition sets CompanyPosition field to given value.


### GetActive

`func (o *CorporateRootUser) GetActive() interface{}`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CorporateRootUser) GetActiveOk() (*interface{}, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CorporateRootUser) SetActive(v interface{})`

SetActive sets Active field to given value.


### SetActiveNil

`func (o *CorporateRootUser) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *CorporateRootUser) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetEmailVerified

`func (o *CorporateRootUser) GetEmailVerified() interface{}`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *CorporateRootUser) GetEmailVerifiedOk() (*interface{}, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *CorporateRootUser) SetEmailVerified(v interface{})`

SetEmailVerified sets EmailVerified field to given value.


### SetEmailVerifiedNil

`func (o *CorporateRootUser) SetEmailVerifiedNil(b bool)`

 SetEmailVerifiedNil sets the value for EmailVerified to be an explicit nil

### UnsetEmailVerified
`func (o *CorporateRootUser) UnsetEmailVerified()`

UnsetEmailVerified ensures that no value is present for EmailVerified, not even an explicit nil
### GetMobileNumberVerified

`func (o *CorporateRootUser) GetMobileNumberVerified() interface{}`

GetMobileNumberVerified returns the MobileNumberVerified field if non-nil, zero value otherwise.

### GetMobileNumberVerifiedOk

`func (o *CorporateRootUser) GetMobileNumberVerifiedOk() (*interface{}, bool)`

GetMobileNumberVerifiedOk returns a tuple with the MobileNumberVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNumberVerified

`func (o *CorporateRootUser) SetMobileNumberVerified(v interface{})`

SetMobileNumberVerified sets MobileNumberVerified field to given value.


### SetMobileNumberVerifiedNil

`func (o *CorporateRootUser) SetMobileNumberVerifiedNil(b bool)`

 SetMobileNumberVerifiedNil sets the value for MobileNumberVerified to be an explicit nil

### UnsetMobileNumberVerified
`func (o *CorporateRootUser) UnsetMobileNumberVerified()`

UnsetMobileNumberVerified ensures that no value is present for MobileNumberVerified, not even an explicit nil
### GetDateOfBirth

`func (o *CorporateRootUser) GetDateOfBirth() Date`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *CorporateRootUser) GetDateOfBirthOk() (*Date, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *CorporateRootUser) SetDateOfBirth(v Date)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *CorporateRootUser) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


