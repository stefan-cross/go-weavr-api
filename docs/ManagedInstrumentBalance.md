# ManagedInstrumentBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableBalance** | Pointer to **interface{}** | The funds available for transactions on the instrument. Funds that are blocked or pending are not included in the available balance. | [optional] 
**ActualBalance** | Pointer to **interface{}** | The funds that are actually on the instrument. Funds that are blocked or pending, due to for example, a purchase authorisation or a pending deposit, are included in the actual balance. | [optional] 

## Methods

### NewManagedInstrumentBalance

`func NewManagedInstrumentBalance() *ManagedInstrumentBalance`

NewManagedInstrumentBalance instantiates a new ManagedInstrumentBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedInstrumentBalanceWithDefaults

`func NewManagedInstrumentBalanceWithDefaults() *ManagedInstrumentBalance`

NewManagedInstrumentBalanceWithDefaults instantiates a new ManagedInstrumentBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableBalance

`func (o *ManagedInstrumentBalance) GetAvailableBalance() interface{}`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *ManagedInstrumentBalance) GetAvailableBalanceOk() (*interface{}, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *ManagedInstrumentBalance) SetAvailableBalance(v interface{})`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *ManagedInstrumentBalance) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.

### SetAvailableBalanceNil

`func (o *ManagedInstrumentBalance) SetAvailableBalanceNil(b bool)`

 SetAvailableBalanceNil sets the value for AvailableBalance to be an explicit nil

### UnsetAvailableBalance
`func (o *ManagedInstrumentBalance) UnsetAvailableBalance()`

UnsetAvailableBalance ensures that no value is present for AvailableBalance, not even an explicit nil
### GetActualBalance

`func (o *ManagedInstrumentBalance) GetActualBalance() interface{}`

GetActualBalance returns the ActualBalance field if non-nil, zero value otherwise.

### GetActualBalanceOk

`func (o *ManagedInstrumentBalance) GetActualBalanceOk() (*interface{}, bool)`

GetActualBalanceOk returns a tuple with the ActualBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualBalance

`func (o *ManagedInstrumentBalance) SetActualBalance(v interface{})`

SetActualBalance sets ActualBalance field to given value.

### HasActualBalance

`func (o *ManagedInstrumentBalance) HasActualBalance() bool`

HasActualBalance returns a boolean if a field has been set.

### SetActualBalanceNil

`func (o *ManagedInstrumentBalance) SetActualBalanceNil(b bool)`

 SetActualBalanceNil sets the value for ActualBalance to be an explicit nil

### UnsetActualBalance
`func (o *ManagedInstrumentBalance) UnsetActualBalance()`

UnsetActualBalance ensures that no value is present for ActualBalance, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


