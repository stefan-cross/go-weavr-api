# StatementEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | [**TransactionId**](TransactionId.md) | The unique identifier of the transaction. | 
**EntryState** | [**StatementEntryState**](StatementEntryState.md) |  | 
**OriginalAmount** | Pointer to [**CurrencyAmount**](CurrencyAmount.md) | The transaction amount as originally requested. The original amount currency may be different from the currency of the instrument.  In case of purchases, this will indicate the original currency and amount that the merchant requested.  | [optional] 
**ForexRate** | Pointer to [**ScaledAmount**](ScaledAmount.md) | If the &#x60;originalAmount&#x60; is in a different currency from the instrument&#39;s currency, the forex rate used will be provided. | [optional] 
**TransactionAmount** | [**CurrencyAmount**](CurrencyAmount.md) | The transaction amount in the instrument&#39;s currency. | 
**AvailableBalanceAdjustment** | Pointer to [**CurrencyAmount**](CurrencyAmount.md) | The amount of funds credited or debited on the available balance of the instrument. | [optional] 
**ActualBalanceAdjustment** | Pointer to [**CurrencyAmount**](CurrencyAmount.md) | The amount of funds credited or debited on the actual balance of the instrument. | [optional] 
**BalanceAfter** | Pointer to [**CurrencyAmount**](CurrencyAmount.md) | The balance of the instrument after the transaction was executed. | [optional] 
**AvailableBalanceAfter** | Pointer to [**CurrencyAmount**](CurrencyAmount.md) | The available balance of the instrument after the transaction was executed. | [optional] 
**ActualBalanceAfter** | Pointer to [**CurrencyAmount**](CurrencyAmount.md) | The actual balance of the instrument after the transaction was executed. | [optional] 
**TransactionFee** | Pointer to [**CurrencyAmount**](CurrencyAmount.md) | The fee amount that was captured for you (Multi account owner). You can set fees to be taken during transactions in the Multi Portal. | [optional] 
**CardholderFee** | Pointer to [**CurrencyAmount**](CurrencyAmount.md) | The fee amount that was captured for you (Multi account owner). You can set fees to be taken during transactions in the Multi Portal. This field is deprecated - use &#x60;transactionFee&#x60; instead. | [optional] 
**ProcessedTimestamp** | **interface{}** | The timestamp when the transaction was processed by the system, expressed in Epoch timestamp using millisecond precision. | 
**SourceAmount** | Pointer to [**CurrencyAmount**](CurrencyAmount.md) | The source amount of the transaction, if the transaction involves forex. | [optional] 
**AdditionalFields** | Pointer to  | A Map of additional fields. Possible values include fee information. The possible entries are &#x60;merchantName&#x60;, &#x60;merchantCategoryCode&#x60;, &#x60;merchantTerminalCountry&#x60;, &#x60;sourceInstrumentType&#x60;, &#x60;sourceInstrumentId&#x60;, &#x60;destinationInstrumentType&#x60;, &#x60;destinationInstrumentId&#x60;, &#x60;forexPaddingCurrency&#x60;, &#x60;forexPaddingAmount&#x60;, &#x60;note&#x60;, &#x60;sourceInstrumentFriendlyName&#x60;, &#x60;destinationInstrumentFriendlyName&#x60;, &#x60;sourceIdentityType&#x60;, &#x60;sourceIdentityId&#x60;, &#x60;sourceIdentityName&#x60;, &#x60;destinationIdentityType&#x60;, &#x60;destinationIdentityId&#x60;, &#x60;destinationIdentityName&#x60;, &#x60;exchangeRate&#x60;, &#x60;authorisationState&#x60;, &#x60;authorisationRelatedId&#x60;, &#x60;settlementRelatedId&#x60;, &#x60;chargeFeeType&#x60;, &#x60;relatedTransactionId&#x60;, &#x60;relatedTransactionIdType&#x60;, &#x60;beneficiaryName&#x60;, &#x60;beneficiaryAccount&#x60;, &#x60;beneficiaryBankCode&#x60;, &#x60;merchantId&#x60;, &#x60;merchantTransactionType&#x60;, &#x60;systemTransactionType&#x60;, &#x60;authorisationCode&#x60;, &#x60;relatedCardId&#x60;, &#x60;sender&#x60;, &#x60;forexFeeCurrency&#x60;, &#x60;forexFeeAmount&#x60;, &#x60;senderIban&#x60;, &#x60;senderReference&#x60;, &#x60;mandateId&#x60;, &#x60;merchantReference&#x60;, &#x60;description&#x60;, &#x60;declineReason&#x60; and &#x60;spendRuleFailedReason&#x60;. | [optional] 

## Methods

### NewStatementEntry

`func NewStatementEntry(transactionId TransactionId, entryState StatementEntryState, transactionAmount CurrencyAmount, processedTimestamp interface{}, ) *StatementEntry`

NewStatementEntry instantiates a new StatementEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementEntryWithDefaults

`func NewStatementEntryWithDefaults() *StatementEntry`

NewStatementEntryWithDefaults instantiates a new StatementEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *StatementEntry) GetTransactionId() TransactionId`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *StatementEntry) GetTransactionIdOk() (*TransactionId, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *StatementEntry) SetTransactionId(v TransactionId)`

SetTransactionId sets TransactionId field to given value.


### GetEntryState

`func (o *StatementEntry) GetEntryState() StatementEntryState`

GetEntryState returns the EntryState field if non-nil, zero value otherwise.

### GetEntryStateOk

`func (o *StatementEntry) GetEntryStateOk() (*StatementEntryState, bool)`

GetEntryStateOk returns a tuple with the EntryState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryState

`func (o *StatementEntry) SetEntryState(v StatementEntryState)`

SetEntryState sets EntryState field to given value.


### GetOriginalAmount

`func (o *StatementEntry) GetOriginalAmount() CurrencyAmount`

GetOriginalAmount returns the OriginalAmount field if non-nil, zero value otherwise.

### GetOriginalAmountOk

`func (o *StatementEntry) GetOriginalAmountOk() (*CurrencyAmount, bool)`

GetOriginalAmountOk returns a tuple with the OriginalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAmount

`func (o *StatementEntry) SetOriginalAmount(v CurrencyAmount)`

SetOriginalAmount sets OriginalAmount field to given value.

### HasOriginalAmount

`func (o *StatementEntry) HasOriginalAmount() bool`

HasOriginalAmount returns a boolean if a field has been set.

### GetForexRate

`func (o *StatementEntry) GetForexRate() ScaledAmount`

GetForexRate returns the ForexRate field if non-nil, zero value otherwise.

### GetForexRateOk

`func (o *StatementEntry) GetForexRateOk() (*ScaledAmount, bool)`

GetForexRateOk returns a tuple with the ForexRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexRate

`func (o *StatementEntry) SetForexRate(v ScaledAmount)`

SetForexRate sets ForexRate field to given value.

### HasForexRate

`func (o *StatementEntry) HasForexRate() bool`

HasForexRate returns a boolean if a field has been set.

### GetTransactionAmount

`func (o *StatementEntry) GetTransactionAmount() CurrencyAmount`

GetTransactionAmount returns the TransactionAmount field if non-nil, zero value otherwise.

### GetTransactionAmountOk

`func (o *StatementEntry) GetTransactionAmountOk() (*CurrencyAmount, bool)`

GetTransactionAmountOk returns a tuple with the TransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAmount

`func (o *StatementEntry) SetTransactionAmount(v CurrencyAmount)`

SetTransactionAmount sets TransactionAmount field to given value.


### GetAvailableBalanceAdjustment

`func (o *StatementEntry) GetAvailableBalanceAdjustment() CurrencyAmount`

GetAvailableBalanceAdjustment returns the AvailableBalanceAdjustment field if non-nil, zero value otherwise.

### GetAvailableBalanceAdjustmentOk

`func (o *StatementEntry) GetAvailableBalanceAdjustmentOk() (*CurrencyAmount, bool)`

GetAvailableBalanceAdjustmentOk returns a tuple with the AvailableBalanceAdjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalanceAdjustment

`func (o *StatementEntry) SetAvailableBalanceAdjustment(v CurrencyAmount)`

SetAvailableBalanceAdjustment sets AvailableBalanceAdjustment field to given value.

### HasAvailableBalanceAdjustment

`func (o *StatementEntry) HasAvailableBalanceAdjustment() bool`

HasAvailableBalanceAdjustment returns a boolean if a field has been set.

### GetActualBalanceAdjustment

`func (o *StatementEntry) GetActualBalanceAdjustment() CurrencyAmount`

GetActualBalanceAdjustment returns the ActualBalanceAdjustment field if non-nil, zero value otherwise.

### GetActualBalanceAdjustmentOk

`func (o *StatementEntry) GetActualBalanceAdjustmentOk() (*CurrencyAmount, bool)`

GetActualBalanceAdjustmentOk returns a tuple with the ActualBalanceAdjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualBalanceAdjustment

`func (o *StatementEntry) SetActualBalanceAdjustment(v CurrencyAmount)`

SetActualBalanceAdjustment sets ActualBalanceAdjustment field to given value.

### HasActualBalanceAdjustment

`func (o *StatementEntry) HasActualBalanceAdjustment() bool`

HasActualBalanceAdjustment returns a boolean if a field has been set.

### GetBalanceAfter

`func (o *StatementEntry) GetBalanceAfter() CurrencyAmount`

GetBalanceAfter returns the BalanceAfter field if non-nil, zero value otherwise.

### GetBalanceAfterOk

`func (o *StatementEntry) GetBalanceAfterOk() (*CurrencyAmount, bool)`

GetBalanceAfterOk returns a tuple with the BalanceAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAfter

`func (o *StatementEntry) SetBalanceAfter(v CurrencyAmount)`

SetBalanceAfter sets BalanceAfter field to given value.

### HasBalanceAfter

`func (o *StatementEntry) HasBalanceAfter() bool`

HasBalanceAfter returns a boolean if a field has been set.

### GetAvailableBalanceAfter

`func (o *StatementEntry) GetAvailableBalanceAfter() CurrencyAmount`

GetAvailableBalanceAfter returns the AvailableBalanceAfter field if non-nil, zero value otherwise.

### GetAvailableBalanceAfterOk

`func (o *StatementEntry) GetAvailableBalanceAfterOk() (*CurrencyAmount, bool)`

GetAvailableBalanceAfterOk returns a tuple with the AvailableBalanceAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalanceAfter

`func (o *StatementEntry) SetAvailableBalanceAfter(v CurrencyAmount)`

SetAvailableBalanceAfter sets AvailableBalanceAfter field to given value.

### HasAvailableBalanceAfter

`func (o *StatementEntry) HasAvailableBalanceAfter() bool`

HasAvailableBalanceAfter returns a boolean if a field has been set.

### GetActualBalanceAfter

`func (o *StatementEntry) GetActualBalanceAfter() CurrencyAmount`

GetActualBalanceAfter returns the ActualBalanceAfter field if non-nil, zero value otherwise.

### GetActualBalanceAfterOk

`func (o *StatementEntry) GetActualBalanceAfterOk() (*CurrencyAmount, bool)`

GetActualBalanceAfterOk returns a tuple with the ActualBalanceAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualBalanceAfter

`func (o *StatementEntry) SetActualBalanceAfter(v CurrencyAmount)`

SetActualBalanceAfter sets ActualBalanceAfter field to given value.

### HasActualBalanceAfter

`func (o *StatementEntry) HasActualBalanceAfter() bool`

HasActualBalanceAfter returns a boolean if a field has been set.

### GetTransactionFee

`func (o *StatementEntry) GetTransactionFee() CurrencyAmount`

GetTransactionFee returns the TransactionFee field if non-nil, zero value otherwise.

### GetTransactionFeeOk

`func (o *StatementEntry) GetTransactionFeeOk() (*CurrencyAmount, bool)`

GetTransactionFeeOk returns a tuple with the TransactionFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionFee

`func (o *StatementEntry) SetTransactionFee(v CurrencyAmount)`

SetTransactionFee sets TransactionFee field to given value.

### HasTransactionFee

`func (o *StatementEntry) HasTransactionFee() bool`

HasTransactionFee returns a boolean if a field has been set.

### GetCardholderFee

`func (o *StatementEntry) GetCardholderFee() CurrencyAmount`

GetCardholderFee returns the CardholderFee field if non-nil, zero value otherwise.

### GetCardholderFeeOk

`func (o *StatementEntry) GetCardholderFeeOk() (*CurrencyAmount, bool)`

GetCardholderFeeOk returns a tuple with the CardholderFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderFee

`func (o *StatementEntry) SetCardholderFee(v CurrencyAmount)`

SetCardholderFee sets CardholderFee field to given value.

### HasCardholderFee

`func (o *StatementEntry) HasCardholderFee() bool`

HasCardholderFee returns a boolean if a field has been set.

### GetProcessedTimestamp

`func (o *StatementEntry) GetProcessedTimestamp() interface{}`

GetProcessedTimestamp returns the ProcessedTimestamp field if non-nil, zero value otherwise.

### GetProcessedTimestampOk

`func (o *StatementEntry) GetProcessedTimestampOk() (*interface{}, bool)`

GetProcessedTimestampOk returns a tuple with the ProcessedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedTimestamp

`func (o *StatementEntry) SetProcessedTimestamp(v interface{})`

SetProcessedTimestamp sets ProcessedTimestamp field to given value.


### SetProcessedTimestampNil

`func (o *StatementEntry) SetProcessedTimestampNil(b bool)`

 SetProcessedTimestampNil sets the value for ProcessedTimestamp to be an explicit nil

### UnsetProcessedTimestamp
`func (o *StatementEntry) UnsetProcessedTimestamp()`

UnsetProcessedTimestamp ensures that no value is present for ProcessedTimestamp, not even an explicit nil
### GetSourceAmount

`func (o *StatementEntry) GetSourceAmount() CurrencyAmount`

GetSourceAmount returns the SourceAmount field if non-nil, zero value otherwise.

### GetSourceAmountOk

`func (o *StatementEntry) GetSourceAmountOk() (*CurrencyAmount, bool)`

GetSourceAmountOk returns a tuple with the SourceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAmount

`func (o *StatementEntry) SetSourceAmount(v CurrencyAmount)`

SetSourceAmount sets SourceAmount field to given value.

### HasSourceAmount

`func (o *StatementEntry) HasSourceAmount() bool`

HasSourceAmount returns a boolean if a field has been set.

### GetAdditionalFields

`func (o *StatementEntry) GetAdditionalFields() map[string]interface{}`

GetAdditionalFields returns the AdditionalFields field if non-nil, zero value otherwise.

### GetAdditionalFieldsOk

`func (o *StatementEntry) GetAdditionalFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalFieldsOk returns a tuple with the AdditionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFields

`func (o *StatementEntry) SetAdditionalFields(v map[string]interface{})`

SetAdditionalFields sets AdditionalFields field to given value.

### HasAdditionalFields

`func (o *StatementEntry) HasAdditionalFields() bool`

HasAdditionalFields returns a boolean if a field has been set.

### SetAdditionalFieldsNil

`func (o *StatementEntry) SetAdditionalFieldsNil(b bool)`

 SetAdditionalFieldsNil sets the value for AdditionalFields to be an explicit nil

### UnsetAdditionalFields
`func (o *StatementEntry) UnsetAdditionalFields()`

UnsetAdditionalFields ensures that no value is present for AdditionalFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


