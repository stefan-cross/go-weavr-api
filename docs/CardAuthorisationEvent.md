# CardAuthorisationEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardId** | **interface{}** | The unique identifier of the card on which an authorisation is being performed | 
**TransactionId** | **interface{}** | The id of this transaction, for reference. | 
**AuthorisationType** | [**CardAuthorisationDetailsAuthorisationType**](CardAuthorisationDetailsAuthorisationType.md) | The type of authorisation (Debit or Credit) | 
**SourceAmount** | [**CurrencyAmount**](CurrencyAmount.md) | The amount in the currency of the merchant. | 
**TransactionAmount** | [**CurrencyAmount**](CurrencyAmount.md) | The amount in the currency of the card. | 
**TotalTransactionCost** | [**CurrencyAmount**](CurrencyAmount.md) | The total amount to be deducted in the currency of the card. This is the summation of the transaction amount, forex padding and forex fees. | 
**TransactionTimestamp** | **interface{}** | The timestamp of the transaction, using epoch timestamp with millisecond precision. | 
**MerchantData** | [**MerchantData**](MerchantData.md) | Merchant related information | 
**Owner** | [**IdentityId**](IdentityId.md) | The owner of the card | 
**CardholderPresent** | Pointer to [**CardHolderPresent**](CardHolderPresent.md) | Optional detail indicating if the card holder was present when the authorisation occurred. | [optional] 
**CardPresent** | Pointer to **interface{}** | Optional detail indicating if the card was present when the authorisation occurred. | [optional] 
**AuthCode** | Pointer to **interface{}** | The authorisation code associated with this authorisation. | [optional] 
**ForexPadding** | Pointer to [**CurrencyAmount**](CurrencyAmount.md) | The forex padding amount, if any, that has been included in the transactionAmount.  Forex padding is extra amount blocked to cater for currency fluctuation. | [optional] 
**ForexFee** | Pointer to [**CurrencyAmount**](CurrencyAmount.md) | The forex fee, if set, that is included in the transactionAmount. | [optional] 
**Mode** | Pointer to **interface{}** | The card can be created in prepaid mode or debit mode.  - A prepaid mode card has its own balance and can have funds transferred to or from it. - A debit mode card does not have its own balance but will be able to spend funds belonging to its parent managed account, subject to a configurable spend limit.  | [optional] 

## Methods

### NewCardAuthorisationEvent

`func NewCardAuthorisationEvent(cardId interface{}, transactionId interface{}, authorisationType CardAuthorisationDetailsAuthorisationType, sourceAmount CurrencyAmount, transactionAmount CurrencyAmount, totalTransactionCost CurrencyAmount, transactionTimestamp interface{}, merchantData MerchantData, owner IdentityId, ) *CardAuthorisationEvent`

NewCardAuthorisationEvent instantiates a new CardAuthorisationEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardAuthorisationEventWithDefaults

`func NewCardAuthorisationEventWithDefaults() *CardAuthorisationEvent`

NewCardAuthorisationEventWithDefaults instantiates a new CardAuthorisationEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardId

`func (o *CardAuthorisationEvent) GetCardId() interface{}`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *CardAuthorisationEvent) GetCardIdOk() (*interface{}, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *CardAuthorisationEvent) SetCardId(v interface{})`

SetCardId sets CardId field to given value.


### SetCardIdNil

`func (o *CardAuthorisationEvent) SetCardIdNil(b bool)`

 SetCardIdNil sets the value for CardId to be an explicit nil

### UnsetCardId
`func (o *CardAuthorisationEvent) UnsetCardId()`

UnsetCardId ensures that no value is present for CardId, not even an explicit nil
### GetTransactionId

`func (o *CardAuthorisationEvent) GetTransactionId() interface{}`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *CardAuthorisationEvent) GetTransactionIdOk() (*interface{}, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *CardAuthorisationEvent) SetTransactionId(v interface{})`

SetTransactionId sets TransactionId field to given value.


### SetTransactionIdNil

`func (o *CardAuthorisationEvent) SetTransactionIdNil(b bool)`

 SetTransactionIdNil sets the value for TransactionId to be an explicit nil

### UnsetTransactionId
`func (o *CardAuthorisationEvent) UnsetTransactionId()`

UnsetTransactionId ensures that no value is present for TransactionId, not even an explicit nil
### GetAuthorisationType

`func (o *CardAuthorisationEvent) GetAuthorisationType() CardAuthorisationDetailsAuthorisationType`

GetAuthorisationType returns the AuthorisationType field if non-nil, zero value otherwise.

### GetAuthorisationTypeOk

`func (o *CardAuthorisationEvent) GetAuthorisationTypeOk() (*CardAuthorisationDetailsAuthorisationType, bool)`

GetAuthorisationTypeOk returns a tuple with the AuthorisationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorisationType

`func (o *CardAuthorisationEvent) SetAuthorisationType(v CardAuthorisationDetailsAuthorisationType)`

SetAuthorisationType sets AuthorisationType field to given value.


### GetSourceAmount

`func (o *CardAuthorisationEvent) GetSourceAmount() CurrencyAmount`

GetSourceAmount returns the SourceAmount field if non-nil, zero value otherwise.

### GetSourceAmountOk

`func (o *CardAuthorisationEvent) GetSourceAmountOk() (*CurrencyAmount, bool)`

GetSourceAmountOk returns a tuple with the SourceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAmount

`func (o *CardAuthorisationEvent) SetSourceAmount(v CurrencyAmount)`

SetSourceAmount sets SourceAmount field to given value.


### GetTransactionAmount

`func (o *CardAuthorisationEvent) GetTransactionAmount() CurrencyAmount`

GetTransactionAmount returns the TransactionAmount field if non-nil, zero value otherwise.

### GetTransactionAmountOk

`func (o *CardAuthorisationEvent) GetTransactionAmountOk() (*CurrencyAmount, bool)`

GetTransactionAmountOk returns a tuple with the TransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAmount

`func (o *CardAuthorisationEvent) SetTransactionAmount(v CurrencyAmount)`

SetTransactionAmount sets TransactionAmount field to given value.


### GetTotalTransactionCost

`func (o *CardAuthorisationEvent) GetTotalTransactionCost() CurrencyAmount`

GetTotalTransactionCost returns the TotalTransactionCost field if non-nil, zero value otherwise.

### GetTotalTransactionCostOk

`func (o *CardAuthorisationEvent) GetTotalTransactionCostOk() (*CurrencyAmount, bool)`

GetTotalTransactionCostOk returns a tuple with the TotalTransactionCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTransactionCost

`func (o *CardAuthorisationEvent) SetTotalTransactionCost(v CurrencyAmount)`

SetTotalTransactionCost sets TotalTransactionCost field to given value.


### GetTransactionTimestamp

`func (o *CardAuthorisationEvent) GetTransactionTimestamp() interface{}`

GetTransactionTimestamp returns the TransactionTimestamp field if non-nil, zero value otherwise.

### GetTransactionTimestampOk

`func (o *CardAuthorisationEvent) GetTransactionTimestampOk() (*interface{}, bool)`

GetTransactionTimestampOk returns a tuple with the TransactionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTimestamp

`func (o *CardAuthorisationEvent) SetTransactionTimestamp(v interface{})`

SetTransactionTimestamp sets TransactionTimestamp field to given value.


### SetTransactionTimestampNil

`func (o *CardAuthorisationEvent) SetTransactionTimestampNil(b bool)`

 SetTransactionTimestampNil sets the value for TransactionTimestamp to be an explicit nil

### UnsetTransactionTimestamp
`func (o *CardAuthorisationEvent) UnsetTransactionTimestamp()`

UnsetTransactionTimestamp ensures that no value is present for TransactionTimestamp, not even an explicit nil
### GetMerchantData

`func (o *CardAuthorisationEvent) GetMerchantData() MerchantData`

GetMerchantData returns the MerchantData field if non-nil, zero value otherwise.

### GetMerchantDataOk

`func (o *CardAuthorisationEvent) GetMerchantDataOk() (*MerchantData, bool)`

GetMerchantDataOk returns a tuple with the MerchantData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantData

`func (o *CardAuthorisationEvent) SetMerchantData(v MerchantData)`

SetMerchantData sets MerchantData field to given value.


### GetOwner

`func (o *CardAuthorisationEvent) GetOwner() IdentityId`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CardAuthorisationEvent) GetOwnerOk() (*IdentityId, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CardAuthorisationEvent) SetOwner(v IdentityId)`

SetOwner sets Owner field to given value.


### GetCardholderPresent

`func (o *CardAuthorisationEvent) GetCardholderPresent() CardHolderPresent`

GetCardholderPresent returns the CardholderPresent field if non-nil, zero value otherwise.

### GetCardholderPresentOk

`func (o *CardAuthorisationEvent) GetCardholderPresentOk() (*CardHolderPresent, bool)`

GetCardholderPresentOk returns a tuple with the CardholderPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderPresent

`func (o *CardAuthorisationEvent) SetCardholderPresent(v CardHolderPresent)`

SetCardholderPresent sets CardholderPresent field to given value.

### HasCardholderPresent

`func (o *CardAuthorisationEvent) HasCardholderPresent() bool`

HasCardholderPresent returns a boolean if a field has been set.

### GetCardPresent

`func (o *CardAuthorisationEvent) GetCardPresent() interface{}`

GetCardPresent returns the CardPresent field if non-nil, zero value otherwise.

### GetCardPresentOk

`func (o *CardAuthorisationEvent) GetCardPresentOk() (*interface{}, bool)`

GetCardPresentOk returns a tuple with the CardPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardPresent

`func (o *CardAuthorisationEvent) SetCardPresent(v interface{})`

SetCardPresent sets CardPresent field to given value.

### HasCardPresent

`func (o *CardAuthorisationEvent) HasCardPresent() bool`

HasCardPresent returns a boolean if a field has been set.

### SetCardPresentNil

`func (o *CardAuthorisationEvent) SetCardPresentNil(b bool)`

 SetCardPresentNil sets the value for CardPresent to be an explicit nil

### UnsetCardPresent
`func (o *CardAuthorisationEvent) UnsetCardPresent()`

UnsetCardPresent ensures that no value is present for CardPresent, not even an explicit nil
### GetAuthCode

`func (o *CardAuthorisationEvent) GetAuthCode() interface{}`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *CardAuthorisationEvent) GetAuthCodeOk() (*interface{}, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *CardAuthorisationEvent) SetAuthCode(v interface{})`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *CardAuthorisationEvent) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### SetAuthCodeNil

`func (o *CardAuthorisationEvent) SetAuthCodeNil(b bool)`

 SetAuthCodeNil sets the value for AuthCode to be an explicit nil

### UnsetAuthCode
`func (o *CardAuthorisationEvent) UnsetAuthCode()`

UnsetAuthCode ensures that no value is present for AuthCode, not even an explicit nil
### GetForexPadding

`func (o *CardAuthorisationEvent) GetForexPadding() CurrencyAmount`

GetForexPadding returns the ForexPadding field if non-nil, zero value otherwise.

### GetForexPaddingOk

`func (o *CardAuthorisationEvent) GetForexPaddingOk() (*CurrencyAmount, bool)`

GetForexPaddingOk returns a tuple with the ForexPadding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexPadding

`func (o *CardAuthorisationEvent) SetForexPadding(v CurrencyAmount)`

SetForexPadding sets ForexPadding field to given value.

### HasForexPadding

`func (o *CardAuthorisationEvent) HasForexPadding() bool`

HasForexPadding returns a boolean if a field has been set.

### GetForexFee

`func (o *CardAuthorisationEvent) GetForexFee() CurrencyAmount`

GetForexFee returns the ForexFee field if non-nil, zero value otherwise.

### GetForexFeeOk

`func (o *CardAuthorisationEvent) GetForexFeeOk() (*CurrencyAmount, bool)`

GetForexFeeOk returns a tuple with the ForexFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForexFee

`func (o *CardAuthorisationEvent) SetForexFee(v CurrencyAmount)`

SetForexFee sets ForexFee field to given value.

### HasForexFee

`func (o *CardAuthorisationEvent) HasForexFee() bool`

HasForexFee returns a boolean if a field has been set.

### GetMode

`func (o *CardAuthorisationEvent) GetMode() interface{}`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CardAuthorisationEvent) GetModeOk() (*interface{}, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CardAuthorisationEvent) SetMode(v interface{})`

SetMode sets Mode field to given value.

### HasMode

`func (o *CardAuthorisationEvent) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *CardAuthorisationEvent) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *CardAuthorisationEvent) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


