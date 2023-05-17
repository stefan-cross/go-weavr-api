# LegalAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | **interface{}** |  | 
**AddressLine2** | Pointer to **interface{}** |  | [optional] 
**City** | **interface{}** |  | 
**PostCode** | Pointer to **interface{}** |  | [optional] 
**State** | Pointer to **interface{}** |  | [optional] 
**Country** | **interface{}** | Country of the identity in ISO 3166 alpha-2 format. | 

## Methods

### NewLegalAddress

`func NewLegalAddress(addressLine1 interface{}, city interface{}, country interface{}, ) *LegalAddress`

NewLegalAddress instantiates a new LegalAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalAddressWithDefaults

`func NewLegalAddressWithDefaults() *LegalAddress`

NewLegalAddressWithDefaults instantiates a new LegalAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *LegalAddress) GetAddressLine1() interface{}`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *LegalAddress) GetAddressLine1Ok() (*interface{}, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *LegalAddress) SetAddressLine1(v interface{})`

SetAddressLine1 sets AddressLine1 field to given value.


### SetAddressLine1Nil

`func (o *LegalAddress) SetAddressLine1Nil(b bool)`

 SetAddressLine1Nil sets the value for AddressLine1 to be an explicit nil

### UnsetAddressLine1
`func (o *LegalAddress) UnsetAddressLine1()`

UnsetAddressLine1 ensures that no value is present for AddressLine1, not even an explicit nil
### GetAddressLine2

`func (o *LegalAddress) GetAddressLine2() interface{}`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *LegalAddress) GetAddressLine2Ok() (*interface{}, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *LegalAddress) SetAddressLine2(v interface{})`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *LegalAddress) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *LegalAddress) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *LegalAddress) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetCity

`func (o *LegalAddress) GetCity() interface{}`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *LegalAddress) GetCityOk() (*interface{}, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *LegalAddress) SetCity(v interface{})`

SetCity sets City field to given value.


### SetCityNil

`func (o *LegalAddress) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *LegalAddress) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetPostCode

`func (o *LegalAddress) GetPostCode() interface{}`

GetPostCode returns the PostCode field if non-nil, zero value otherwise.

### GetPostCodeOk

`func (o *LegalAddress) GetPostCodeOk() (*interface{}, bool)`

GetPostCodeOk returns a tuple with the PostCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCode

`func (o *LegalAddress) SetPostCode(v interface{})`

SetPostCode sets PostCode field to given value.

### HasPostCode

`func (o *LegalAddress) HasPostCode() bool`

HasPostCode returns a boolean if a field has been set.

### SetPostCodeNil

`func (o *LegalAddress) SetPostCodeNil(b bool)`

 SetPostCodeNil sets the value for PostCode to be an explicit nil

### UnsetPostCode
`func (o *LegalAddress) UnsetPostCode()`

UnsetPostCode ensures that no value is present for PostCode, not even an explicit nil
### GetState

`func (o *LegalAddress) GetState() interface{}`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LegalAddress) GetStateOk() (*interface{}, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LegalAddress) SetState(v interface{})`

SetState sets State field to given value.

### HasState

`func (o *LegalAddress) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *LegalAddress) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *LegalAddress) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCountry

`func (o *LegalAddress) GetCountry() interface{}`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *LegalAddress) GetCountryOk() (*interface{}, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *LegalAddress) SetCountry(v interface{})`

SetCountry sets Country field to given value.


### SetCountryNil

`func (o *LegalAddress) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *LegalAddress) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


