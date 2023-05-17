# CurrencyAmount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **interface{}** | The currency expressed in ISO-4217 code. Example: GBP, EUR, USD. | 
**Amount** | **interface{}** | The monetary amount, scaled to the lowest denomination of the currency.  Example, an amount of 1000 for a EUR currency is actually 1000 Euro cents or EUR 10.00.  | 

## Methods

### NewCurrencyAmount

`func NewCurrencyAmount(currency interface{}, amount interface{}, ) *CurrencyAmount`

NewCurrencyAmount instantiates a new CurrencyAmount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyAmountWithDefaults

`func NewCurrencyAmountWithDefaults() *CurrencyAmount`

NewCurrencyAmountWithDefaults instantiates a new CurrencyAmount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *CurrencyAmount) GetCurrency() interface{}`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CurrencyAmount) GetCurrencyOk() (*interface{}, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CurrencyAmount) SetCurrency(v interface{})`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *CurrencyAmount) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *CurrencyAmount) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetAmount

`func (o *CurrencyAmount) GetAmount() interface{}`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CurrencyAmount) GetAmountOk() (*interface{}, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CurrencyAmount) SetAmount(v interface{})`

SetAmount sets Amount field to given value.


### SetAmountNil

`func (o *CurrencyAmount) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *CurrencyAmount) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


