# SpendLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | [**CurrencyAmount**](CurrencyAmount.md) | The spend limit, as amount and currency, (for the given interval). | 
**Interval** | [**SpendLimitInterval**](SpendLimitInterval.md) |  | 

## Methods

### NewSpendLimit

`func NewSpendLimit(value CurrencyAmount, interval SpendLimitInterval, ) *SpendLimit`

NewSpendLimit instantiates a new SpendLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpendLimitWithDefaults

`func NewSpendLimitWithDefaults() *SpendLimit`

NewSpendLimitWithDefaults instantiates a new SpendLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *SpendLimit) GetValue() CurrencyAmount`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SpendLimit) GetValueOk() (*CurrencyAmount, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SpendLimit) SetValue(v CurrencyAmount)`

SetValue sets Value field to given value.


### GetInterval

`func (o *SpendLimit) GetInterval() SpendLimitInterval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *SpendLimit) GetIntervalOk() (*SpendLimitInterval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *SpendLimit) SetInterval(v SpendLimitInterval)`

SetInterval sets Interval field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


