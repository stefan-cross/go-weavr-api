# Mobile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | **interface{}** | The country code of the Corporate&#39;s root user mobile number (e.g. +44). | 
**Number** | **interface{}** | The mobile number of the Corporate&#39;s root user - excluding country code. | 

## Methods

### NewMobile

`func NewMobile(countryCode interface{}, number interface{}, ) *Mobile`

NewMobile instantiates a new Mobile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileWithDefaults

`func NewMobileWithDefaults() *Mobile`

NewMobileWithDefaults instantiates a new Mobile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *Mobile) GetCountryCode() interface{}`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Mobile) GetCountryCodeOk() (*interface{}, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Mobile) SetCountryCode(v interface{})`

SetCountryCode sets CountryCode field to given value.


### SetCountryCodeNil

`func (o *Mobile) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *Mobile) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetNumber

`func (o *Mobile) GetNumber() interface{}`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Mobile) GetNumberOk() (*interface{}, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Mobile) SetNumber(v interface{})`

SetNumber sets Number field to given value.


### SetNumberNil

`func (o *Mobile) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *Mobile) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


