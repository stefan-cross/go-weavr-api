# AvailableToSpend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**CurrencyAmount**](CurrencyAmount.md) | The amount and currency that is available to spend, (for the given interval). | [optional] 
**Interval** | Pointer to [**SpendLimitInterval**](SpendLimitInterval.md) |  | [optional] 

## Methods

### NewAvailableToSpend

`func NewAvailableToSpend() *AvailableToSpend`

NewAvailableToSpend instantiates a new AvailableToSpend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableToSpendWithDefaults

`func NewAvailableToSpendWithDefaults() *AvailableToSpend`

NewAvailableToSpendWithDefaults instantiates a new AvailableToSpend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *AvailableToSpend) GetValue() CurrencyAmount`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AvailableToSpend) GetValueOk() (*CurrencyAmount, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AvailableToSpend) SetValue(v CurrencyAmount)`

SetValue sets Value field to given value.

### HasValue

`func (o *AvailableToSpend) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetInterval

`func (o *AvailableToSpend) GetInterval() SpendLimitInterval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *AvailableToSpend) GetIntervalOk() (*SpendLimitInterval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *AvailableToSpend) SetInterval(v SpendLimitInterval)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *AvailableToSpend) HasInterval() bool`

HasInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


