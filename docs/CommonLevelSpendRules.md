# CommonLevelSpendRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedMerchantCategories** | Pointer to **interface{}** | Whitelist MCC: A list of allowed merchant category codes (MCCs). If the MCC does not match, then the transaction will be declined. If an MCC is also in the blocked list, the blocked list will take precedence. | [optional] 
**BlockedMerchantCategories** | Pointer to **interface{}** | Blacklist MCC: A list of disallowed merchant category codes (MCCs). If the MCC matches, then the transaction will be declined. If an MCC is also in the allowed list, the blocked list will take precedence. | [optional] 
**AllowedMerchantIds** | Pointer to **interface{}** | Whitelist Merchant Id: A list of allowed merchant IDs. If the Merchant Id does not match, then the transaction will be declined. If a Merchant Id is also provided in the blocked list, the blocked list will take precedence. | [optional] 
**BlockedMerchantIds** | Pointer to **interface{}** | Blacklist Merchant Id: A list of disallowed merchant IDs. If the Merchant Id matches, then the transaction will be declined. If a Merchant Id is also in the allowed list, the blocked list will take precedence. | [optional] 
**AllowedMerchantCountries** | Pointer to **interface{}** | Whitelist Merchant Country: A list of allowed merchant countries, in ISO 3166-1 alpha-2 format. If the Merchant country does not match, then the transaction will be declined. If a Merchant Country is also provided in the blocked list, the blocked list will take precedence. | [optional] 
**BlockedMerchantCountries** | Pointer to **interface{}** | Blacklist Merchant Country: A list of disallowed merchant countries, in ISO 3166-1 alpha-2 format. If the Merchant country matches, then the transaction will be declined. If a Merchant Country is also in the allowed list, the blocked list will take precedence. | [optional] 
**AllowContactless** | Pointer to **interface{}** | Indicates if a contactless transaction is allowed on the card. | [optional] 
**AllowAtm** | Pointer to **interface{}** | Indicates if an ATM Withdrawal transaction is allowed on the card. | [optional] 
**AllowECommerce** | Pointer to **interface{}** | Indicates if an online transaction is allowed on the card. | [optional] 
**AllowCashback** | Pointer to **interface{}** | Indicates if a cashback transaction at a physical terminal is allowed on the card. | [optional] 
**AllowCreditAuthorisations** | Pointer to **interface{}** | Indicates if a the card can receive a credit transaction. | [optional] 

## Methods

### NewCommonLevelSpendRules

`func NewCommonLevelSpendRules() *CommonLevelSpendRules`

NewCommonLevelSpendRules instantiates a new CommonLevelSpendRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonLevelSpendRulesWithDefaults

`func NewCommonLevelSpendRulesWithDefaults() *CommonLevelSpendRules`

NewCommonLevelSpendRulesWithDefaults instantiates a new CommonLevelSpendRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedMerchantCategories

`func (o *CommonLevelSpendRules) GetAllowedMerchantCategories() interface{}`

GetAllowedMerchantCategories returns the AllowedMerchantCategories field if non-nil, zero value otherwise.

### GetAllowedMerchantCategoriesOk

`func (o *CommonLevelSpendRules) GetAllowedMerchantCategoriesOk() (*interface{}, bool)`

GetAllowedMerchantCategoriesOk returns a tuple with the AllowedMerchantCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMerchantCategories

`func (o *CommonLevelSpendRules) SetAllowedMerchantCategories(v interface{})`

SetAllowedMerchantCategories sets AllowedMerchantCategories field to given value.

### HasAllowedMerchantCategories

`func (o *CommonLevelSpendRules) HasAllowedMerchantCategories() bool`

HasAllowedMerchantCategories returns a boolean if a field has been set.

### SetAllowedMerchantCategoriesNil

`func (o *CommonLevelSpendRules) SetAllowedMerchantCategoriesNil(b bool)`

 SetAllowedMerchantCategoriesNil sets the value for AllowedMerchantCategories to be an explicit nil

### UnsetAllowedMerchantCategories
`func (o *CommonLevelSpendRules) UnsetAllowedMerchantCategories()`

UnsetAllowedMerchantCategories ensures that no value is present for AllowedMerchantCategories, not even an explicit nil
### GetBlockedMerchantCategories

`func (o *CommonLevelSpendRules) GetBlockedMerchantCategories() interface{}`

GetBlockedMerchantCategories returns the BlockedMerchantCategories field if non-nil, zero value otherwise.

### GetBlockedMerchantCategoriesOk

`func (o *CommonLevelSpendRules) GetBlockedMerchantCategoriesOk() (*interface{}, bool)`

GetBlockedMerchantCategoriesOk returns a tuple with the BlockedMerchantCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedMerchantCategories

`func (o *CommonLevelSpendRules) SetBlockedMerchantCategories(v interface{})`

SetBlockedMerchantCategories sets BlockedMerchantCategories field to given value.

### HasBlockedMerchantCategories

`func (o *CommonLevelSpendRules) HasBlockedMerchantCategories() bool`

HasBlockedMerchantCategories returns a boolean if a field has been set.

### SetBlockedMerchantCategoriesNil

`func (o *CommonLevelSpendRules) SetBlockedMerchantCategoriesNil(b bool)`

 SetBlockedMerchantCategoriesNil sets the value for BlockedMerchantCategories to be an explicit nil

### UnsetBlockedMerchantCategories
`func (o *CommonLevelSpendRules) UnsetBlockedMerchantCategories()`

UnsetBlockedMerchantCategories ensures that no value is present for BlockedMerchantCategories, not even an explicit nil
### GetAllowedMerchantIds

`func (o *CommonLevelSpendRules) GetAllowedMerchantIds() interface{}`

GetAllowedMerchantIds returns the AllowedMerchantIds field if non-nil, zero value otherwise.

### GetAllowedMerchantIdsOk

`func (o *CommonLevelSpendRules) GetAllowedMerchantIdsOk() (*interface{}, bool)`

GetAllowedMerchantIdsOk returns a tuple with the AllowedMerchantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMerchantIds

`func (o *CommonLevelSpendRules) SetAllowedMerchantIds(v interface{})`

SetAllowedMerchantIds sets AllowedMerchantIds field to given value.

### HasAllowedMerchantIds

`func (o *CommonLevelSpendRules) HasAllowedMerchantIds() bool`

HasAllowedMerchantIds returns a boolean if a field has been set.

### SetAllowedMerchantIdsNil

`func (o *CommonLevelSpendRules) SetAllowedMerchantIdsNil(b bool)`

 SetAllowedMerchantIdsNil sets the value for AllowedMerchantIds to be an explicit nil

### UnsetAllowedMerchantIds
`func (o *CommonLevelSpendRules) UnsetAllowedMerchantIds()`

UnsetAllowedMerchantIds ensures that no value is present for AllowedMerchantIds, not even an explicit nil
### GetBlockedMerchantIds

`func (o *CommonLevelSpendRules) GetBlockedMerchantIds() interface{}`

GetBlockedMerchantIds returns the BlockedMerchantIds field if non-nil, zero value otherwise.

### GetBlockedMerchantIdsOk

`func (o *CommonLevelSpendRules) GetBlockedMerchantIdsOk() (*interface{}, bool)`

GetBlockedMerchantIdsOk returns a tuple with the BlockedMerchantIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedMerchantIds

`func (o *CommonLevelSpendRules) SetBlockedMerchantIds(v interface{})`

SetBlockedMerchantIds sets BlockedMerchantIds field to given value.

### HasBlockedMerchantIds

`func (o *CommonLevelSpendRules) HasBlockedMerchantIds() bool`

HasBlockedMerchantIds returns a boolean if a field has been set.

### SetBlockedMerchantIdsNil

`func (o *CommonLevelSpendRules) SetBlockedMerchantIdsNil(b bool)`

 SetBlockedMerchantIdsNil sets the value for BlockedMerchantIds to be an explicit nil

### UnsetBlockedMerchantIds
`func (o *CommonLevelSpendRules) UnsetBlockedMerchantIds()`

UnsetBlockedMerchantIds ensures that no value is present for BlockedMerchantIds, not even an explicit nil
### GetAllowedMerchantCountries

`func (o *CommonLevelSpendRules) GetAllowedMerchantCountries() interface{}`

GetAllowedMerchantCountries returns the AllowedMerchantCountries field if non-nil, zero value otherwise.

### GetAllowedMerchantCountriesOk

`func (o *CommonLevelSpendRules) GetAllowedMerchantCountriesOk() (*interface{}, bool)`

GetAllowedMerchantCountriesOk returns a tuple with the AllowedMerchantCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMerchantCountries

`func (o *CommonLevelSpendRules) SetAllowedMerchantCountries(v interface{})`

SetAllowedMerchantCountries sets AllowedMerchantCountries field to given value.

### HasAllowedMerchantCountries

`func (o *CommonLevelSpendRules) HasAllowedMerchantCountries() bool`

HasAllowedMerchantCountries returns a boolean if a field has been set.

### SetAllowedMerchantCountriesNil

`func (o *CommonLevelSpendRules) SetAllowedMerchantCountriesNil(b bool)`

 SetAllowedMerchantCountriesNil sets the value for AllowedMerchantCountries to be an explicit nil

### UnsetAllowedMerchantCountries
`func (o *CommonLevelSpendRules) UnsetAllowedMerchantCountries()`

UnsetAllowedMerchantCountries ensures that no value is present for AllowedMerchantCountries, not even an explicit nil
### GetBlockedMerchantCountries

`func (o *CommonLevelSpendRules) GetBlockedMerchantCountries() interface{}`

GetBlockedMerchantCountries returns the BlockedMerchantCountries field if non-nil, zero value otherwise.

### GetBlockedMerchantCountriesOk

`func (o *CommonLevelSpendRules) GetBlockedMerchantCountriesOk() (*interface{}, bool)`

GetBlockedMerchantCountriesOk returns a tuple with the BlockedMerchantCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedMerchantCountries

`func (o *CommonLevelSpendRules) SetBlockedMerchantCountries(v interface{})`

SetBlockedMerchantCountries sets BlockedMerchantCountries field to given value.

### HasBlockedMerchantCountries

`func (o *CommonLevelSpendRules) HasBlockedMerchantCountries() bool`

HasBlockedMerchantCountries returns a boolean if a field has been set.

### SetBlockedMerchantCountriesNil

`func (o *CommonLevelSpendRules) SetBlockedMerchantCountriesNil(b bool)`

 SetBlockedMerchantCountriesNil sets the value for BlockedMerchantCountries to be an explicit nil

### UnsetBlockedMerchantCountries
`func (o *CommonLevelSpendRules) UnsetBlockedMerchantCountries()`

UnsetBlockedMerchantCountries ensures that no value is present for BlockedMerchantCountries, not even an explicit nil
### GetAllowContactless

`func (o *CommonLevelSpendRules) GetAllowContactless() interface{}`

GetAllowContactless returns the AllowContactless field if non-nil, zero value otherwise.

### GetAllowContactlessOk

`func (o *CommonLevelSpendRules) GetAllowContactlessOk() (*interface{}, bool)`

GetAllowContactlessOk returns a tuple with the AllowContactless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowContactless

`func (o *CommonLevelSpendRules) SetAllowContactless(v interface{})`

SetAllowContactless sets AllowContactless field to given value.

### HasAllowContactless

`func (o *CommonLevelSpendRules) HasAllowContactless() bool`

HasAllowContactless returns a boolean if a field has been set.

### SetAllowContactlessNil

`func (o *CommonLevelSpendRules) SetAllowContactlessNil(b bool)`

 SetAllowContactlessNil sets the value for AllowContactless to be an explicit nil

### UnsetAllowContactless
`func (o *CommonLevelSpendRules) UnsetAllowContactless()`

UnsetAllowContactless ensures that no value is present for AllowContactless, not even an explicit nil
### GetAllowAtm

`func (o *CommonLevelSpendRules) GetAllowAtm() interface{}`

GetAllowAtm returns the AllowAtm field if non-nil, zero value otherwise.

### GetAllowAtmOk

`func (o *CommonLevelSpendRules) GetAllowAtmOk() (*interface{}, bool)`

GetAllowAtmOk returns a tuple with the AllowAtm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAtm

`func (o *CommonLevelSpendRules) SetAllowAtm(v interface{})`

SetAllowAtm sets AllowAtm field to given value.

### HasAllowAtm

`func (o *CommonLevelSpendRules) HasAllowAtm() bool`

HasAllowAtm returns a boolean if a field has been set.

### SetAllowAtmNil

`func (o *CommonLevelSpendRules) SetAllowAtmNil(b bool)`

 SetAllowAtmNil sets the value for AllowAtm to be an explicit nil

### UnsetAllowAtm
`func (o *CommonLevelSpendRules) UnsetAllowAtm()`

UnsetAllowAtm ensures that no value is present for AllowAtm, not even an explicit nil
### GetAllowECommerce

`func (o *CommonLevelSpendRules) GetAllowECommerce() interface{}`

GetAllowECommerce returns the AllowECommerce field if non-nil, zero value otherwise.

### GetAllowECommerceOk

`func (o *CommonLevelSpendRules) GetAllowECommerceOk() (*interface{}, bool)`

GetAllowECommerceOk returns a tuple with the AllowECommerce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowECommerce

`func (o *CommonLevelSpendRules) SetAllowECommerce(v interface{})`

SetAllowECommerce sets AllowECommerce field to given value.

### HasAllowECommerce

`func (o *CommonLevelSpendRules) HasAllowECommerce() bool`

HasAllowECommerce returns a boolean if a field has been set.

### SetAllowECommerceNil

`func (o *CommonLevelSpendRules) SetAllowECommerceNil(b bool)`

 SetAllowECommerceNil sets the value for AllowECommerce to be an explicit nil

### UnsetAllowECommerce
`func (o *CommonLevelSpendRules) UnsetAllowECommerce()`

UnsetAllowECommerce ensures that no value is present for AllowECommerce, not even an explicit nil
### GetAllowCashback

`func (o *CommonLevelSpendRules) GetAllowCashback() interface{}`

GetAllowCashback returns the AllowCashback field if non-nil, zero value otherwise.

### GetAllowCashbackOk

`func (o *CommonLevelSpendRules) GetAllowCashbackOk() (*interface{}, bool)`

GetAllowCashbackOk returns a tuple with the AllowCashback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCashback

`func (o *CommonLevelSpendRules) SetAllowCashback(v interface{})`

SetAllowCashback sets AllowCashback field to given value.

### HasAllowCashback

`func (o *CommonLevelSpendRules) HasAllowCashback() bool`

HasAllowCashback returns a boolean if a field has been set.

### SetAllowCashbackNil

`func (o *CommonLevelSpendRules) SetAllowCashbackNil(b bool)`

 SetAllowCashbackNil sets the value for AllowCashback to be an explicit nil

### UnsetAllowCashback
`func (o *CommonLevelSpendRules) UnsetAllowCashback()`

UnsetAllowCashback ensures that no value is present for AllowCashback, not even an explicit nil
### GetAllowCreditAuthorisations

`func (o *CommonLevelSpendRules) GetAllowCreditAuthorisations() interface{}`

GetAllowCreditAuthorisations returns the AllowCreditAuthorisations field if non-nil, zero value otherwise.

### GetAllowCreditAuthorisationsOk

`func (o *CommonLevelSpendRules) GetAllowCreditAuthorisationsOk() (*interface{}, bool)`

GetAllowCreditAuthorisationsOk returns a tuple with the AllowCreditAuthorisations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCreditAuthorisations

`func (o *CommonLevelSpendRules) SetAllowCreditAuthorisations(v interface{})`

SetAllowCreditAuthorisations sets AllowCreditAuthorisations field to given value.

### HasAllowCreditAuthorisations

`func (o *CommonLevelSpendRules) HasAllowCreditAuthorisations() bool`

HasAllowCreditAuthorisations returns a boolean if a field has been set.

### SetAllowCreditAuthorisationsNil

`func (o *CommonLevelSpendRules) SetAllowCreditAuthorisationsNil(b bool)`

 SetAllowCreditAuthorisationsNil sets the value for AllowCreditAuthorisations to be an explicit nil

### UnsetAllowCreditAuthorisations
`func (o *CommonLevelSpendRules) UnsetAllowCreditAuthorisations()`

UnsetAllowCreditAuthorisations ensures that no value is present for AllowCreditAuthorisations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


