# DeliveryAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** |  | 
**Surname** | **interface{}** |  | 
**AddressLine1** | **interface{}** |  | 
**AddressLine2** | Pointer to **interface{}** |  | [optional] 
**City** | **interface{}** |  | 
**PostCode** | **interface{}** |  | 
**State** | Pointer to **interface{}** |  | [optional] 
**Country** | **interface{}** | Country of the identity in ISO 3166 alpha-2 format. | 

## Methods

### NewDeliveryAddress

`func NewDeliveryAddress(name interface{}, surname interface{}, addressLine1 interface{}, city interface{}, postCode interface{}, country interface{}, ) *DeliveryAddress`

NewDeliveryAddress instantiates a new DeliveryAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryAddressWithDefaults

`func NewDeliveryAddressWithDefaults() *DeliveryAddress`

NewDeliveryAddressWithDefaults instantiates a new DeliveryAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeliveryAddress) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeliveryAddress) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeliveryAddress) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *DeliveryAddress) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DeliveryAddress) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSurname

`func (o *DeliveryAddress) GetSurname() interface{}`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *DeliveryAddress) GetSurnameOk() (*interface{}, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *DeliveryAddress) SetSurname(v interface{})`

SetSurname sets Surname field to given value.


### SetSurnameNil

`func (o *DeliveryAddress) SetSurnameNil(b bool)`

 SetSurnameNil sets the value for Surname to be an explicit nil

### UnsetSurname
`func (o *DeliveryAddress) UnsetSurname()`

UnsetSurname ensures that no value is present for Surname, not even an explicit nil
### GetAddressLine1

`func (o *DeliveryAddress) GetAddressLine1() interface{}`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *DeliveryAddress) GetAddressLine1Ok() (*interface{}, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *DeliveryAddress) SetAddressLine1(v interface{})`

SetAddressLine1 sets AddressLine1 field to given value.


### SetAddressLine1Nil

`func (o *DeliveryAddress) SetAddressLine1Nil(b bool)`

 SetAddressLine1Nil sets the value for AddressLine1 to be an explicit nil

### UnsetAddressLine1
`func (o *DeliveryAddress) UnsetAddressLine1()`

UnsetAddressLine1 ensures that no value is present for AddressLine1, not even an explicit nil
### GetAddressLine2

`func (o *DeliveryAddress) GetAddressLine2() interface{}`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *DeliveryAddress) GetAddressLine2Ok() (*interface{}, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *DeliveryAddress) SetAddressLine2(v interface{})`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *DeliveryAddress) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### SetAddressLine2Nil

`func (o *DeliveryAddress) SetAddressLine2Nil(b bool)`

 SetAddressLine2Nil sets the value for AddressLine2 to be an explicit nil

### UnsetAddressLine2
`func (o *DeliveryAddress) UnsetAddressLine2()`

UnsetAddressLine2 ensures that no value is present for AddressLine2, not even an explicit nil
### GetCity

`func (o *DeliveryAddress) GetCity() interface{}`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *DeliveryAddress) GetCityOk() (*interface{}, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *DeliveryAddress) SetCity(v interface{})`

SetCity sets City field to given value.


### SetCityNil

`func (o *DeliveryAddress) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *DeliveryAddress) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetPostCode

`func (o *DeliveryAddress) GetPostCode() interface{}`

GetPostCode returns the PostCode field if non-nil, zero value otherwise.

### GetPostCodeOk

`func (o *DeliveryAddress) GetPostCodeOk() (*interface{}, bool)`

GetPostCodeOk returns a tuple with the PostCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCode

`func (o *DeliveryAddress) SetPostCode(v interface{})`

SetPostCode sets PostCode field to given value.


### SetPostCodeNil

`func (o *DeliveryAddress) SetPostCodeNil(b bool)`

 SetPostCodeNil sets the value for PostCode to be an explicit nil

### UnsetPostCode
`func (o *DeliveryAddress) UnsetPostCode()`

UnsetPostCode ensures that no value is present for PostCode, not even an explicit nil
### GetState

`func (o *DeliveryAddress) GetState() interface{}`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DeliveryAddress) GetStateOk() (*interface{}, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DeliveryAddress) SetState(v interface{})`

SetState sets State field to given value.

### HasState

`func (o *DeliveryAddress) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *DeliveryAddress) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *DeliveryAddress) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCountry

`func (o *DeliveryAddress) GetCountry() interface{}`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *DeliveryAddress) GetCountryOk() (*interface{}, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *DeliveryAddress) SetCountry(v interface{})`

SetCountry sets Country field to given value.


### SetCountryNil

`func (o *DeliveryAddress) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *DeliveryAddress) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


