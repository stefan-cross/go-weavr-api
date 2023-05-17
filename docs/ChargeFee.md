# ChargeFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | [**TransactionId**](TransactionId.md) | The unique identifier of the transaction. | 
**ProfileId** | **interface{}** | The profile Id which a specific identity, instrument or transaction type is linked to. | 
**FeeType** | **interface{}** | The fee type as defined in the Multi Portal, and as sent in the request. | 
**Source** | [**InstrumentId**](InstrumentId.md) |  | 
**AvailableBalanceAdjustment** | Pointer to [**CurrencyAmount**](CurrencyAmount.md) | The object representing a monetary amount in a particular currency. | [optional] 
**State** | Pointer to [**TransactionState**](TransactionState.md) | The transaction entry state. | [optional] 
**CreationTimestamp** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewChargeFee

`func NewChargeFee(transactionId TransactionId, profileId interface{}, feeType interface{}, source InstrumentId, ) *ChargeFee`

NewChargeFee instantiates a new ChargeFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeFeeWithDefaults

`func NewChargeFeeWithDefaults() *ChargeFee`

NewChargeFeeWithDefaults instantiates a new ChargeFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *ChargeFee) GetTransactionId() TransactionId`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ChargeFee) GetTransactionIdOk() (*TransactionId, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ChargeFee) SetTransactionId(v TransactionId)`

SetTransactionId sets TransactionId field to given value.


### GetProfileId

`func (o *ChargeFee) GetProfileId() interface{}`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *ChargeFee) GetProfileIdOk() (*interface{}, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *ChargeFee) SetProfileId(v interface{})`

SetProfileId sets ProfileId field to given value.


### SetProfileIdNil

`func (o *ChargeFee) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *ChargeFee) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
### GetFeeType

`func (o *ChargeFee) GetFeeType() interface{}`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *ChargeFee) GetFeeTypeOk() (*interface{}, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *ChargeFee) SetFeeType(v interface{})`

SetFeeType sets FeeType field to given value.


### SetFeeTypeNil

`func (o *ChargeFee) SetFeeTypeNil(b bool)`

 SetFeeTypeNil sets the value for FeeType to be an explicit nil

### UnsetFeeType
`func (o *ChargeFee) UnsetFeeType()`

UnsetFeeType ensures that no value is present for FeeType, not even an explicit nil
### GetSource

`func (o *ChargeFee) GetSource() InstrumentId`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ChargeFee) GetSourceOk() (*InstrumentId, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ChargeFee) SetSource(v InstrumentId)`

SetSource sets Source field to given value.


### GetAvailableBalanceAdjustment

`func (o *ChargeFee) GetAvailableBalanceAdjustment() CurrencyAmount`

GetAvailableBalanceAdjustment returns the AvailableBalanceAdjustment field if non-nil, zero value otherwise.

### GetAvailableBalanceAdjustmentOk

`func (o *ChargeFee) GetAvailableBalanceAdjustmentOk() (*CurrencyAmount, bool)`

GetAvailableBalanceAdjustmentOk returns a tuple with the AvailableBalanceAdjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalanceAdjustment

`func (o *ChargeFee) SetAvailableBalanceAdjustment(v CurrencyAmount)`

SetAvailableBalanceAdjustment sets AvailableBalanceAdjustment field to given value.

### HasAvailableBalanceAdjustment

`func (o *ChargeFee) HasAvailableBalanceAdjustment() bool`

HasAvailableBalanceAdjustment returns a boolean if a field has been set.

### GetState

`func (o *ChargeFee) GetState() TransactionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ChargeFee) GetStateOk() (*TransactionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ChargeFee) SetState(v TransactionState)`

SetState sets State field to given value.

### HasState

`func (o *ChargeFee) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCreationTimestamp

`func (o *ChargeFee) GetCreationTimestamp() interface{}`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *ChargeFee) GetCreationTimestampOk() (*interface{}, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *ChargeFee) SetCreationTimestamp(v interface{})`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *ChargeFee) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.

### SetCreationTimestampNil

`func (o *ChargeFee) SetCreationTimestampNil(b bool)`

 SetCreationTimestampNil sets the value for CreationTimestamp to be an explicit nil

### UnsetCreationTimestamp
`func (o *ChargeFee) UnsetCreationTimestamp()`

UnsetCreationTimestamp ensures that no value is present for CreationTimestamp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


