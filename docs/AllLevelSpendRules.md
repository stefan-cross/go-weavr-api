# AllLevelSpendRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedMerchantCategories** | Pointer to **interface{}** | Whitelist MCC: A list of allowed merchant category codes (MCCs). If the MCC does not match, then the transaction will be declined. If an MCC is also in the blocked list, the blocked list will take precedence. This field is deprecated - use \&quot;cardLevelSpendRules\&quot;.\&quot;allowedMerchantCategories\&quot; instead. | [optional] 
**BlockedMerchantCategories** | Pointer to **interface{}** | Blacklist MCC: A list of disallowed merchant category codes (MCCs). If the MCC matches, then the transaction will be declined. If an MCC is also in the allowed list, the blocked list will take precedence. This field is deprecated - use \&quot;cardLevelSpendRules\&quot;.\&quot;blockedMerchantCategories\&quot; instead. | [optional] 
**AllowedMerchantIds** | Pointer to **interface{}** | Whitelist Merchant Id: A list of allowed merchant IDs. If the Merchant Id does not match, then the transaction will be declined. If a Merchant Id is also provided in the blocked list, the blocked list will take precedence. This field is deprecated - use \&quot;cardLevelSpendRules\&quot;.\&quot;allowedMerchantIds\&quot; instead. | [optional] 
**BlockedMerchantIds** | Pointer to **interface{}** | Blacklist Merchant Id: A list of disallowed merchant IDs. If the Merchant Id matches, then the transaction will be declined. If a Merchant Id is also in the allowed list, the blocked list will take precedence. This field is deprecated - use \&quot;cardLevelSpendRules\&quot;.\&quot;blockedMerchantIds\&quot; instead. | [optional] 
**AllowedMerchantCountries** | Pointer to **interface{}** | Whitelist Merchant Country: A list of allowed merchant countries, in ISO 3166-1 alpha-2 format. If the Merchant country does not match, then the transaction will be declined. If a Merchant Country is also provided in the blocked list, the blocked list will take precedence. This field is deprecated - use \&quot;cardLevelSpendRules\&quot;.\&quot;allowedMerchantCountries\&quot; instead. | [optional] 
**BlockedMerchantCountries** | Pointer to **interface{}** | Blacklist Merchant Country: A list of disallowed merchant countries, in ISO 3166-1 alpha-2 format. If the Merchant country matches, then the transaction will be declined. If a Merchant Country is also in the allowed list, the blocked list will take precedence. This field is deprecated - use \&quot;cardLevelSpendRules\&quot;.\&quot;blockedMerchantCountries\&quot; instead. | [optional] 
**AllowContactless** | Pointer to **interface{}** | Indicates if a contactless transaction is allowed on the card. This field is deprecated - use \&quot;cardLevelSpendRules\&quot;.\&quot;allowContactless\&quot; instead. | [optional] 
**AllowAtm** | Pointer to **interface{}** | Indicates if an ATM Withdrawal transaction is allowed on the card. This field is deprecated - use \&quot;cardLevelSpendRules\&quot;.\&quot;allowAtm\&quot; instead. | [optional] 
**AllowECommerce** | Pointer to **interface{}** | Indicates if an online transaction is allowed on the card. This field is deprecated - use \&quot;cardLevelSpendRules\&quot;.\&quot;allowECommerce\&quot; instead. | [optional] 
**AllowCashback** | Pointer to **interface{}** | Indicates if a cashback transaction at a physical terminal is allowed on the card. This field is deprecated - use \&quot;cardLevelSpendRules\&quot;.\&quot;allowCashback\&quot; instead. | [optional] 
**AllowCreditAuthorisations** | Pointer to **interface{}** | Indicates if a the card can receive a credit transaction. This field is deprecated - use \&quot;cardLevelSpendRules\&quot;.\&quot;allowCashback\&quot; instead. | [optional] 
**MinTransactionAmount** | Pointer to **interface{}** | The minimum transaction amount, in card currency, that is allowed. If the transaction amount is less than this value, then the transaction will be declined. Omit this, or set to 0 if no maximum transaction amount is to be set. This field is deprecated - use \&quot;cardLevelSpendRules\&quot;.\&quot;minTransactionAmount\&quot; instead. | [optional] 
**MaxTransactionAmount** | Pointer to **interface{}** | The maximum transaction amount, in card currency, that is allowed. If the transaction amount is greater than this value, then the transaction will be declined. Omit this, or set to 0 if no maximum transaction amount is to be set. This field is deprecated - use \&quot;cardLevelSpendRules\&quot;.\&quot;maxTransactionAmount\&quot; instead. | [optional] 
**SpendLimit** | Pointer to **interface{}** | The total amount of funds that can be spent using this card, in a given time interval. This field is deprecated - use \&quot;cardLevelSpendRules\&quot;.\&quot;spendLimit\&quot; instead. | [optional] 
**CardLevelSpendRules** | Pointer to **interface{}** |  | [optional] 
**ProfileLevelSpendRules** | Pointer to **interface{}** |  | [optional] 
**IdentityLevelSpendRules** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewAllLevelSpendRules

`func NewAllLevelSpendRules() *AllLevelSpendRules`

NewAllLevelSpendRules instantiates a new AllLevelSpendRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllLevelSpendRulesWithDefaults

`func NewAllLevelSpendRulesWithDefaults() *AllLevelSpendRules`

NewAllLevelSpendRulesWithDefaults instantiates a new AllLevelSpendRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedMerchantCategories

`func (o *AllLevelSpendRules) GetAllowedMerchantCategories() interface{}`

GetAllowedMerchantCategories returns the AllowedMerchantCategories field if non-nil, zero value otherwise.

### GetAllowedMerchantCategoriesOk

`func (o *AllLevelSpendRules) GetAllowedMerchantCategoriesOk() (*interface{}, bool)`

GetAllowedMerchantCategoriesOk returns a tuple with the AllowedMerchantCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMerchantCategories

`func (o *AllLevelSpendRules) SetAllowedMerchantCategories(v interface{})`

SetAllowedMerchantCategories sets AllowedMerchantCategories field to given value.

### HasAllowedMerchantCategories

`func (o *AllLevelSpendRules) HasAllowedMerchantCategories() bool`

HasAllowedMerchantCategories returns a boolean if a field has been set.

### SetAllowedMerchantCategoriesNil

`func (o *AllLevelSpendRules) SetAllowedMerchantCategoriesNil(b bool)`

 SetAllowedMerchantCategoriesNil sets the value for AllowedMerchantCategories to be an explicit nil

### UnsetAllowedMerchantCategories
`func (o *AllLevelSpendRules) UnsetAllowedMerchantCategories()`

UnsetAllowedMerchantCategories ensures that no value is present for AllowedMerchantCategories, not even an explicit nil
### GetBlockedMerchantCategories

`func (o *AllLevelSpendRules) GetBlockedMerchantCategories() interface{}`

GetBlockedMerchantCategories returns the BlockedMerchantCategories field if non-nil, zero value otherwise.

### GetBlockedMerchantCategoriesOk

`func (o *AllLevelSpendRules) GetBlockedMerchantCategoriesOk() (*interface{}, bool)`

GetBlockedMerchantCategoriesOk returns a tuple with the BlockedMerchantCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedMerchantCategories

`func (o *AllLevelSpendRules) SetBlockedMerchantCategories(v interface{})`

SetBlockedMerchantCategories sets BlockedMerchantCategories field to given value.

### HasBlockedMerchantCategories

`func (o *AllLevelSpendRules) HasBlockedMerchantCategories() bool`

HasBlockedMerchantCategories returns a boolean if a field has been set.

### SetBlockedMerchantCategoriesNil

`func (o *AllLevelSpendRules) SetBlockedMerchantCategoriesNil(b bool)`

 SetBlockedMerchantCategoriesNil sets the value for BlockedMerchantCategories to be an explicit nil

### UnsetBlockedMerchantCategories
`func (o *AllLevelSpendRules) UnsetBlockedMerchantCategories()`

UnsetBlockedMerchantCategories ensures that no value is present for BlockedMerchantCategories, not even an explicit nil
### GetAllowedMerchantIds

`func (o *AllLevelSpendRules) GetAllowedMerchantIds() interface{}`

GetAllowedMerchantIds returns the AllowedMerchantIds field if non-nil, zero value otherwise.

### GetAllowedMerchantIdsOk

`func (o *AllLevelSpendRules) GetAllowedMerchantIdsOk() (*interface{}, bool)`

GetAllowedMerchantIdsOk returns a tuple with the AllowedMerchantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMerchantIds

`func (o *AllLevelSpendRules) SetAllowedMerchantIds(v interface{})`

SetAllowedMerchantIds sets AllowedMerchantIds field to given value.

### HasAllowedMerchantIds

`func (o *AllLevelSpendRules) HasAllowedMerchantIds() bool`

HasAllowedMerchantIds returns a boolean if a field has been set.

### SetAllowedMerchantIdsNil

`func (o *AllLevelSpendRules) SetAllowedMerchantIdsNil(b bool)`

 SetAllowedMerchantIdsNil sets the value for AllowedMerchantIds to be an explicit nil

### UnsetAllowedMerchantIds
`func (o *AllLevelSpendRules) UnsetAllowedMerchantIds()`

UnsetAllowedMerchantIds ensures that no value is present for AllowedMerchantIds, not even an explicit nil
### GetBlockedMerchantIds

`func (o *AllLevelSpendRules) GetBlockedMerchantIds() interface{}`

GetBlockedMerchantIds returns the BlockedMerchantIds field if non-nil, zero value otherwise.

### GetBlockedMerchantIdsOk

`func (o *AllLevelSpendRules) GetBlockedMerchantIdsOk() (*interface{}, bool)`

GetBlockedMerchantIdsOk returns a tuple with the BlockedMerchantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedMerchantIds

`func (o *AllLevelSpendRules) SetBlockedMerchantIds(v interface{})`

SetBlockedMerchantIds sets BlockedMerchantIds field to given value.

### HasBlockedMerchantIds

`func (o *AllLevelSpendRules) HasBlockedMerchantIds() bool`

HasBlockedMerchantIds returns a boolean if a field has been set.

### SetBlockedMerchantIdsNil

`func (o *AllLevelSpendRules) SetBlockedMerchantIdsNil(b bool)`

 SetBlockedMerchantIdsNil sets the value for BlockedMerchantIds to be an explicit nil

### UnsetBlockedMerchantIds
`func (o *AllLevelSpendRules) UnsetBlockedMerchantIds()`

UnsetBlockedMerchantIds ensures that no value is present for BlockedMerchantIds, not even an explicit nil
### GetAllowedMerchantCountries

`func (o *AllLevelSpendRules) GetAllowedMerchantCountries() interface{}`

GetAllowedMerchantCountries returns the AllowedMerchantCountries field if non-nil, zero value otherwise.

### GetAllowedMerchantCountriesOk

`func (o *AllLevelSpendRules) GetAllowedMerchantCountriesOk() (*interface{}, bool)`

GetAllowedMerchantCountriesOk returns a tuple with the AllowedMerchantCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMerchantCountries

`func (o *AllLevelSpendRules) SetAllowedMerchantCountries(v interface{})`

SetAllowedMerchantCountries sets AllowedMerchantCountries field to given value.

### HasAllowedMerchantCountries

`func (o *AllLevelSpendRules) HasAllowedMerchantCountries() bool`

HasAllowedMerchantCountries returns a boolean if a field has been set.

### SetAllowedMerchantCountriesNil

`func (o *AllLevelSpendRules) SetAllowedMerchantCountriesNil(b bool)`

 SetAllowedMerchantCountriesNil sets the value for AllowedMerchantCountries to be an explicit nil

### UnsetAllowedMerchantCountries
`func (o *AllLevelSpendRules) UnsetAllowedMerchantCountries()`

UnsetAllowedMerchantCountries ensures that no value is present for AllowedMerchantCountries, not even an explicit nil
### GetBlockedMerchantCountries

`func (o *AllLevelSpendRules) GetBlockedMerchantCountries() interface{}`

GetBlockedMerchantCountries returns the BlockedMerchantCountries field if non-nil, zero value otherwise.

### GetBlockedMerchantCountriesOk

`func (o *AllLevelSpendRules) GetBlockedMerchantCountriesOk() (*interface{}, bool)`

GetBlockedMerchantCountriesOk returns a tuple with the BlockedMerchantCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedMerchantCountries

`func (o *AllLevelSpendRules) SetBlockedMerchantCountries(v interface{})`

SetBlockedMerchantCountries sets BlockedMerchantCountries field to given value.

### HasBlockedMerchantCountries

`func (o *AllLevelSpendRules) HasBlockedMerchantCountries() bool`

HasBlockedMerchantCountries returns a boolean if a field has been set.

### SetBlockedMerchantCountriesNil

`func (o *AllLevelSpendRules) SetBlockedMerchantCountriesNil(b bool)`

 SetBlockedMerchantCountriesNil sets the value for BlockedMerchantCountries to be an explicit nil

### UnsetBlockedMerchantCountries
`func (o *AllLevelSpendRules) UnsetBlockedMerchantCountries()`

UnsetBlockedMerchantCountries ensures that no value is present for BlockedMerchantCountries, not even an explicit nil
### GetAllowContactless

`func (o *AllLevelSpendRules) GetAllowContactless() interface{}`

GetAllowContactless returns the AllowContactless field if non-nil, zero value otherwise.

### GetAllowContactlessOk

`func (o *AllLevelSpendRules) GetAllowContactlessOk() (*interface{}, bool)`

GetAllowContactlessOk returns a tuple with the AllowContactless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowContactless

`func (o *AllLevelSpendRules) SetAllowContactless(v interface{})`

SetAllowContactless sets AllowContactless field to given value.

### HasAllowContactless

`func (o *AllLevelSpendRules) HasAllowContactless() bool`

HasAllowContactless returns a boolean if a field has been set.

### SetAllowContactlessNil

`func (o *AllLevelSpendRules) SetAllowContactlessNil(b bool)`

 SetAllowContactlessNil sets the value for AllowContactless to be an explicit nil

### UnsetAllowContactless
`func (o *AllLevelSpendRules) UnsetAllowContactless()`

UnsetAllowContactless ensures that no value is present for AllowContactless, not even an explicit nil
### GetAllowAtm

`func (o *AllLevelSpendRules) GetAllowAtm() interface{}`

GetAllowAtm returns the AllowAtm field if non-nil, zero value otherwise.

### GetAllowAtmOk

`func (o *AllLevelSpendRules) GetAllowAtmOk() (*interface{}, bool)`

GetAllowAtmOk returns a tuple with the AllowAtm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAtm

`func (o *AllLevelSpendRules) SetAllowAtm(v interface{})`

SetAllowAtm sets AllowAtm field to given value.

### HasAllowAtm

`func (o *AllLevelSpendRules) HasAllowAtm() bool`

HasAllowAtm returns a boolean if a field has been set.

### SetAllowAtmNil

`func (o *AllLevelSpendRules) SetAllowAtmNil(b bool)`

 SetAllowAtmNil sets the value for AllowAtm to be an explicit nil

### UnsetAllowAtm
`func (o *AllLevelSpendRules) UnsetAllowAtm()`

UnsetAllowAtm ensures that no value is present for AllowAtm, not even an explicit nil
### GetAllowECommerce

`func (o *AllLevelSpendRules) GetAllowECommerce() interface{}`

GetAllowECommerce returns the AllowECommerce field if non-nil, zero value otherwise.

### GetAllowECommerceOk

`func (o *AllLevelSpendRules) GetAllowECommerceOk() (*interface{}, bool)`

GetAllowECommerceOk returns a tuple with the AllowECommerce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowECommerce

`func (o *AllLevelSpendRules) SetAllowECommerce(v interface{})`

SetAllowECommerce sets AllowECommerce field to given value.

### HasAllowECommerce

`func (o *AllLevelSpendRules) HasAllowECommerce() bool`

HasAllowECommerce returns a boolean if a field has been set.

### SetAllowECommerceNil

`func (o *AllLevelSpendRules) SetAllowECommerceNil(b bool)`

 SetAllowECommerceNil sets the value for AllowECommerce to be an explicit nil

### UnsetAllowECommerce
`func (o *AllLevelSpendRules) UnsetAllowECommerce()`

UnsetAllowECommerce ensures that no value is present for AllowECommerce, not even an explicit nil
### GetAllowCashback

`func (o *AllLevelSpendRules) GetAllowCashback() interface{}`

GetAllowCashback returns the AllowCashback field if non-nil, zero value otherwise.

### GetAllowCashbackOk

`func (o *AllLevelSpendRules) GetAllowCashbackOk() (*interface{}, bool)`

GetAllowCashbackOk returns a tuple with the AllowCashback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCashback

`func (o *AllLevelSpendRules) SetAllowCashback(v interface{})`

SetAllowCashback sets AllowCashback field to given value.

### HasAllowCashback

`func (o *AllLevelSpendRules) HasAllowCashback() bool`

HasAllowCashback returns a boolean if a field has been set.

### SetAllowCashbackNil

`func (o *AllLevelSpendRules) SetAllowCashbackNil(b bool)`

 SetAllowCashbackNil sets the value for AllowCashback to be an explicit nil

### UnsetAllowCashback
`func (o *AllLevelSpendRules) UnsetAllowCashback()`

UnsetAllowCashback ensures that no value is present for AllowCashback, not even an explicit nil
### GetAllowCreditAuthorisations

`func (o *AllLevelSpendRules) GetAllowCreditAuthorisations() interface{}`

GetAllowCreditAuthorisations returns the AllowCreditAuthorisations field if non-nil, zero value otherwise.

### GetAllowCreditAuthorisationsOk

`func (o *AllLevelSpendRules) GetAllowCreditAuthorisationsOk() (*interface{}, bool)`

GetAllowCreditAuthorisationsOk returns a tuple with the AllowCreditAuthorisations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCreditAuthorisations

`func (o *AllLevelSpendRules) SetAllowCreditAuthorisations(v interface{})`

SetAllowCreditAuthorisations sets AllowCreditAuthorisations field to given value.

### HasAllowCreditAuthorisations

`func (o *AllLevelSpendRules) HasAllowCreditAuthorisations() bool`

HasAllowCreditAuthorisations returns a boolean if a field has been set.

### SetAllowCreditAuthorisationsNil

`func (o *AllLevelSpendRules) SetAllowCreditAuthorisationsNil(b bool)`

 SetAllowCreditAuthorisationsNil sets the value for AllowCreditAuthorisations to be an explicit nil

### UnsetAllowCreditAuthorisations
`func (o *AllLevelSpendRules) UnsetAllowCreditAuthorisations()`

UnsetAllowCreditAuthorisations ensures that no value is present for AllowCreditAuthorisations, not even an explicit nil
### GetMinTransactionAmount

`func (o *AllLevelSpendRules) GetMinTransactionAmount() interface{}`

GetMinTransactionAmount returns the MinTransactionAmount field if non-nil, zero value otherwise.

### GetMinTransactionAmountOk

`func (o *AllLevelSpendRules) GetMinTransactionAmountOk() (*interface{}, bool)`

GetMinTransactionAmountOk returns a tuple with the MinTransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTransactionAmount

`func (o *AllLevelSpendRules) SetMinTransactionAmount(v interface{})`

SetMinTransactionAmount sets MinTransactionAmount field to given value.

### HasMinTransactionAmount

`func (o *AllLevelSpendRules) HasMinTransactionAmount() bool`

HasMinTransactionAmount returns a boolean if a field has been set.

### SetMinTransactionAmountNil

`func (o *AllLevelSpendRules) SetMinTransactionAmountNil(b bool)`

 SetMinTransactionAmountNil sets the value for MinTransactionAmount to be an explicit nil

### UnsetMinTransactionAmount
`func (o *AllLevelSpendRules) UnsetMinTransactionAmount()`

UnsetMinTransactionAmount ensures that no value is present for MinTransactionAmount, not even an explicit nil
### GetMaxTransactionAmount

`func (o *AllLevelSpendRules) GetMaxTransactionAmount() interface{}`

GetMaxTransactionAmount returns the MaxTransactionAmount field if non-nil, zero value otherwise.

### GetMaxTransactionAmountOk

`func (o *AllLevelSpendRules) GetMaxTransactionAmountOk() (*interface{}, bool)`

GetMaxTransactionAmountOk returns a tuple with the MaxTransactionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTransactionAmount

`func (o *AllLevelSpendRules) SetMaxTransactionAmount(v interface{})`

SetMaxTransactionAmount sets MaxTransactionAmount field to given value.

### HasMaxTransactionAmount

`func (o *AllLevelSpendRules) HasMaxTransactionAmount() bool`

HasMaxTransactionAmount returns a boolean if a field has been set.

### SetMaxTransactionAmountNil

`func (o *AllLevelSpendRules) SetMaxTransactionAmountNil(b bool)`

 SetMaxTransactionAmountNil sets the value for MaxTransactionAmount to be an explicit nil

### UnsetMaxTransactionAmount
`func (o *AllLevelSpendRules) UnsetMaxTransactionAmount()`

UnsetMaxTransactionAmount ensures that no value is present for MaxTransactionAmount, not even an explicit nil
### GetSpendLimit

`func (o *AllLevelSpendRules) GetSpendLimit() interface{}`

GetSpendLimit returns the SpendLimit field if non-nil, zero value otherwise.

### GetSpendLimitOk

`func (o *AllLevelSpendRules) GetSpendLimitOk() (*interface{}, bool)`

GetSpendLimitOk returns a tuple with the SpendLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendLimit

`func (o *AllLevelSpendRules) SetSpendLimit(v interface{})`

SetSpendLimit sets SpendLimit field to given value.

### HasSpendLimit

`func (o *AllLevelSpendRules) HasSpendLimit() bool`

HasSpendLimit returns a boolean if a field has been set.

### SetSpendLimitNil

`func (o *AllLevelSpendRules) SetSpendLimitNil(b bool)`

 SetSpendLimitNil sets the value for SpendLimit to be an explicit nil

### UnsetSpendLimit
`func (o *AllLevelSpendRules) UnsetSpendLimit()`

UnsetSpendLimit ensures that no value is present for SpendLimit, not even an explicit nil
### GetCardLevelSpendRules

`func (o *AllLevelSpendRules) GetCardLevelSpendRules() interface{}`

GetCardLevelSpendRules returns the CardLevelSpendRules field if non-nil, zero value otherwise.

### GetCardLevelSpendRulesOk

`func (o *AllLevelSpendRules) GetCardLevelSpendRulesOk() (*interface{}, bool)`

GetCardLevelSpendRulesOk returns a tuple with the CardLevelSpendRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardLevelSpendRules

`func (o *AllLevelSpendRules) SetCardLevelSpendRules(v interface{})`

SetCardLevelSpendRules sets CardLevelSpendRules field to given value.

### HasCardLevelSpendRules

`func (o *AllLevelSpendRules) HasCardLevelSpendRules() bool`

HasCardLevelSpendRules returns a boolean if a field has been set.

### SetCardLevelSpendRulesNil

`func (o *AllLevelSpendRules) SetCardLevelSpendRulesNil(b bool)`

 SetCardLevelSpendRulesNil sets the value for CardLevelSpendRules to be an explicit nil

### UnsetCardLevelSpendRules
`func (o *AllLevelSpendRules) UnsetCardLevelSpendRules()`

UnsetCardLevelSpendRules ensures that no value is present for CardLevelSpendRules, not even an explicit nil
### GetProfileLevelSpendRules

`func (o *AllLevelSpendRules) GetProfileLevelSpendRules() interface{}`

GetProfileLevelSpendRules returns the ProfileLevelSpendRules field if non-nil, zero value otherwise.

### GetProfileLevelSpendRulesOk

`func (o *AllLevelSpendRules) GetProfileLevelSpendRulesOk() (*interface{}, bool)`

GetProfileLevelSpendRulesOk returns a tuple with the ProfileLevelSpendRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileLevelSpendRules

`func (o *AllLevelSpendRules) SetProfileLevelSpendRules(v interface{})`

SetProfileLevelSpendRules sets ProfileLevelSpendRules field to given value.

### HasProfileLevelSpendRules

`func (o *AllLevelSpendRules) HasProfileLevelSpendRules() bool`

HasProfileLevelSpendRules returns a boolean if a field has been set.

### SetProfileLevelSpendRulesNil

`func (o *AllLevelSpendRules) SetProfileLevelSpendRulesNil(b bool)`

 SetProfileLevelSpendRulesNil sets the value for ProfileLevelSpendRules to be an explicit nil

### UnsetProfileLevelSpendRules
`func (o *AllLevelSpendRules) UnsetProfileLevelSpendRules()`

UnsetProfileLevelSpendRules ensures that no value is present for ProfileLevelSpendRules, not even an explicit nil
### GetIdentityLevelSpendRules

`func (o *AllLevelSpendRules) GetIdentityLevelSpendRules() interface{}`

GetIdentityLevelSpendRules returns the IdentityLevelSpendRules field if non-nil, zero value otherwise.

### GetIdentityLevelSpendRulesOk

`func (o *AllLevelSpendRules) GetIdentityLevelSpendRulesOk() (*interface{}, bool)`

GetIdentityLevelSpendRulesOk returns a tuple with the IdentityLevelSpendRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityLevelSpendRules

`func (o *AllLevelSpendRules) SetIdentityLevelSpendRules(v interface{})`

SetIdentityLevelSpendRules sets IdentityLevelSpendRules field to given value.

### HasIdentityLevelSpendRules

`func (o *AllLevelSpendRules) HasIdentityLevelSpendRules() bool`

HasIdentityLevelSpendRules returns a boolean if a field has been set.

### SetIdentityLevelSpendRulesNil

`func (o *AllLevelSpendRules) SetIdentityLevelSpendRulesNil(b bool)`

 SetIdentityLevelSpendRulesNil sets the value for IdentityLevelSpendRules to be an explicit nil

### UnsetIdentityLevelSpendRules
`func (o *AllLevelSpendRules) UnsetIdentityLevelSpendRules()`

UnsetIdentityLevelSpendRules ensures that no value is present for IdentityLevelSpendRules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


