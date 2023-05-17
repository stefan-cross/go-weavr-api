# CorporateCreateRequestCompany

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CompanyType**](CompanyType.md) |  | 
**BusinessAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**Name** | **interface{}** | The registered name of the company. | 
**RegistrationNumber** | Pointer to **interface{}** | The registration number of the company. | [optional] 
**RegistrationCountry** | **interface{}** | The country of company registration, in ISO 3166 alpha-2 format. | 

## Methods

### NewCorporateCreateRequestCompany

`func NewCorporateCreateRequestCompany(type_ CompanyType, name interface{}, registrationCountry interface{}, ) *CorporateCreateRequestCompany`

NewCorporateCreateRequestCompany instantiates a new CorporateCreateRequestCompany object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporateCreateRequestCompanyWithDefaults

`func NewCorporateCreateRequestCompanyWithDefaults() *CorporateCreateRequestCompany`

NewCorporateCreateRequestCompanyWithDefaults instantiates a new CorporateCreateRequestCompany object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CorporateCreateRequestCompany) GetType() CompanyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CorporateCreateRequestCompany) GetTypeOk() (*CompanyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CorporateCreateRequestCompany) SetType(v CompanyType)`

SetType sets Type field to given value.


### GetBusinessAddress

`func (o *CorporateCreateRequestCompany) GetBusinessAddress() Address`

GetBusinessAddress returns the BusinessAddress field if non-nil, zero value otherwise.

### GetBusinessAddressOk

`func (o *CorporateCreateRequestCompany) GetBusinessAddressOk() (*Address, bool)`

GetBusinessAddressOk returns a tuple with the BusinessAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessAddress

`func (o *CorporateCreateRequestCompany) SetBusinessAddress(v Address)`

SetBusinessAddress sets BusinessAddress field to given value.

### HasBusinessAddress

`func (o *CorporateCreateRequestCompany) HasBusinessAddress() bool`

HasBusinessAddress returns a boolean if a field has been set.

### GetName

`func (o *CorporateCreateRequestCompany) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CorporateCreateRequestCompany) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CorporateCreateRequestCompany) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *CorporateCreateRequestCompany) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CorporateCreateRequestCompany) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRegistrationNumber

`func (o *CorporateCreateRequestCompany) GetRegistrationNumber() interface{}`

GetRegistrationNumber returns the RegistrationNumber field if non-nil, zero value otherwise.

### GetRegistrationNumberOk

`func (o *CorporateCreateRequestCompany) GetRegistrationNumberOk() (*interface{}, bool)`

GetRegistrationNumberOk returns a tuple with the RegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationNumber

`func (o *CorporateCreateRequestCompany) SetRegistrationNumber(v interface{})`

SetRegistrationNumber sets RegistrationNumber field to given value.

### HasRegistrationNumber

`func (o *CorporateCreateRequestCompany) HasRegistrationNumber() bool`

HasRegistrationNumber returns a boolean if a field has been set.

### SetRegistrationNumberNil

`func (o *CorporateCreateRequestCompany) SetRegistrationNumberNil(b bool)`

 SetRegistrationNumberNil sets the value for RegistrationNumber to be an explicit nil

### UnsetRegistrationNumber
`func (o *CorporateCreateRequestCompany) UnsetRegistrationNumber()`

UnsetRegistrationNumber ensures that no value is present for RegistrationNumber, not even an explicit nil
### GetRegistrationCountry

`func (o *CorporateCreateRequestCompany) GetRegistrationCountry() interface{}`

GetRegistrationCountry returns the RegistrationCountry field if non-nil, zero value otherwise.

### GetRegistrationCountryOk

`func (o *CorporateCreateRequestCompany) GetRegistrationCountryOk() (*interface{}, bool)`

GetRegistrationCountryOk returns a tuple with the RegistrationCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationCountry

`func (o *CorporateCreateRequestCompany) SetRegistrationCountry(v interface{})`

SetRegistrationCountry sets RegistrationCountry field to given value.


### SetRegistrationCountryNil

`func (o *CorporateCreateRequestCompany) SetRegistrationCountryNil(b bool)`

 SetRegistrationCountryNil sets the value for RegistrationCountry to be an explicit nil

### UnsetRegistrationCountry
`func (o *CorporateCreateRequestCompany) UnsetRegistrationCountry()`

UnsetRegistrationCountry ensures that no value is present for RegistrationCountry, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


