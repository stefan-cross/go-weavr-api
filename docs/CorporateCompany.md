# CorporateCompany

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | The registered name of the company. | 
**Type** | [**CompanyType**](CompanyType.md) |  | 
**RegistrationNumber** | Pointer to **interface{}** | The company registration number. | [optional] 
**RegisteredAddress** | Pointer to [**LegalAddress**](LegalAddress.md) |  | [optional] 
**BusinessAddress** | Pointer to [**Address**](Address.md) | The address where the business is based. | [optional] 
**CountryOfRegistration** | **interface{}** | The country of company registration in ISO 3166 alpha-2. | 
**IncorporatedOn** | Pointer to [**Date**](Date.md) | The company&#39;s date of incorporation | [optional] 

## Methods

### NewCorporateCompany

`func NewCorporateCompany(name interface{}, type_ CompanyType, countryOfRegistration interface{}, ) *CorporateCompany`

NewCorporateCompany instantiates a new CorporateCompany object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporateCompanyWithDefaults

`func NewCorporateCompanyWithDefaults() *CorporateCompany`

NewCorporateCompanyWithDefaults instantiates a new CorporateCompany object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CorporateCompany) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CorporateCompany) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CorporateCompany) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *CorporateCompany) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CorporateCompany) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *CorporateCompany) GetType() CompanyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CorporateCompany) GetTypeOk() (*CompanyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CorporateCompany) SetType(v CompanyType)`

SetType sets Type field to given value.


### GetRegistrationNumber

`func (o *CorporateCompany) GetRegistrationNumber() interface{}`

GetRegistrationNumber returns the RegistrationNumber field if non-nil, zero value otherwise.

### GetRegistrationNumberOk

`func (o *CorporateCompany) GetRegistrationNumberOk() (*interface{}, bool)`

GetRegistrationNumberOk returns a tuple with the RegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationNumber

`func (o *CorporateCompany) SetRegistrationNumber(v interface{})`

SetRegistrationNumber sets RegistrationNumber field to given value.

### HasRegistrationNumber

`func (o *CorporateCompany) HasRegistrationNumber() bool`

HasRegistrationNumber returns a boolean if a field has been set.

### SetRegistrationNumberNil

`func (o *CorporateCompany) SetRegistrationNumberNil(b bool)`

 SetRegistrationNumberNil sets the value for RegistrationNumber to be an explicit nil

### UnsetRegistrationNumber
`func (o *CorporateCompany) UnsetRegistrationNumber()`

UnsetRegistrationNumber ensures that no value is present for RegistrationNumber, not even an explicit nil
### GetRegisteredAddress

`func (o *CorporateCompany) GetRegisteredAddress() LegalAddress`

GetRegisteredAddress returns the RegisteredAddress field if non-nil, zero value otherwise.

### GetRegisteredAddressOk

`func (o *CorporateCompany) GetRegisteredAddressOk() (*LegalAddress, bool)`

GetRegisteredAddressOk returns a tuple with the RegisteredAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredAddress

`func (o *CorporateCompany) SetRegisteredAddress(v LegalAddress)`

SetRegisteredAddress sets RegisteredAddress field to given value.

### HasRegisteredAddress

`func (o *CorporateCompany) HasRegisteredAddress() bool`

HasRegisteredAddress returns a boolean if a field has been set.

### GetBusinessAddress

`func (o *CorporateCompany) GetBusinessAddress() Address`

GetBusinessAddress returns the BusinessAddress field if non-nil, zero value otherwise.

### GetBusinessAddressOk

`func (o *CorporateCompany) GetBusinessAddressOk() (*Address, bool)`

GetBusinessAddressOk returns a tuple with the BusinessAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessAddress

`func (o *CorporateCompany) SetBusinessAddress(v Address)`

SetBusinessAddress sets BusinessAddress field to given value.

### HasBusinessAddress

`func (o *CorporateCompany) HasBusinessAddress() bool`

HasBusinessAddress returns a boolean if a field has been set.

### GetCountryOfRegistration

`func (o *CorporateCompany) GetCountryOfRegistration() interface{}`

GetCountryOfRegistration returns the CountryOfRegistration field if non-nil, zero value otherwise.

### GetCountryOfRegistrationOk

`func (o *CorporateCompany) GetCountryOfRegistrationOk() (*interface{}, bool)`

GetCountryOfRegistrationOk returns a tuple with the CountryOfRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfRegistration

`func (o *CorporateCompany) SetCountryOfRegistration(v interface{})`

SetCountryOfRegistration sets CountryOfRegistration field to given value.


### SetCountryOfRegistrationNil

`func (o *CorporateCompany) SetCountryOfRegistrationNil(b bool)`

 SetCountryOfRegistrationNil sets the value for CountryOfRegistration to be an explicit nil

### UnsetCountryOfRegistration
`func (o *CorporateCompany) UnsetCountryOfRegistration()`

UnsetCountryOfRegistration ensures that no value is present for CountryOfRegistration, not even an explicit nil
### GetIncorporatedOn

`func (o *CorporateCompany) GetIncorporatedOn() Date`

GetIncorporatedOn returns the IncorporatedOn field if non-nil, zero value otherwise.

### GetIncorporatedOnOk

`func (o *CorporateCompany) GetIncorporatedOnOk() (*Date, bool)`

GetIncorporatedOnOk returns a tuple with the IncorporatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncorporatedOn

`func (o *CorporateCompany) SetIncorporatedOn(v Date)`

SetIncorporatedOn sets IncorporatedOn field to given value.

### HasIncorporatedOn

`func (o *CorporateCompany) HasIncorporatedOn() bool`

HasIncorporatedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


