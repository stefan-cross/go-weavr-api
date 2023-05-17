/*
Weavr Multi Product API

Weavr Multi API provides a simple and flexible way to issue cards and accounts to your customers.  By integrating Weavr Multi API in your application you can embed banking capabilities within your app and provide a seamless experience for your customers.  # Authentication Each request to the Multi API must include an `api_key` that represents your account. You can obtain an API Key by registering for a Multi account [here](https://portal.weavr.io).  Almost all endpoints require a secondary authentication token `auth_token` that represents the user for whom the request is being executed. 

API version: 3.32.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package weavr

import (
	"encoding/json"
)

// checks if the CommonLevelSpendRules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonLevelSpendRules{}

// CommonLevelSpendRules struct for CommonLevelSpendRules
type CommonLevelSpendRules struct {
	// Whitelist MCC: A list of allowed merchant category codes (MCCs). If the MCC does not match, then the transaction will be declined. If an MCC is also in the blocked list, the blocked list will take precedence.
	AllowedMerchantCategories interface{} `json:"allowedMerchantCategories,omitempty"`
	// Blacklist MCC: A list of disallowed merchant category codes (MCCs). If the MCC matches, then the transaction will be declined. If an MCC is also in the allowed list, the blocked list will take precedence.
	BlockedMerchantCategories interface{} `json:"blockedMerchantCategories,omitempty"`
	// Whitelist Merchant Id: A list of allowed merchant IDs. If the Merchant Id does not match, then the transaction will be declined. If a Merchant Id is also provided in the blocked list, the blocked list will take precedence.
	AllowedMerchantIds interface{} `json:"allowedMerchantIds,omitempty"`
	// Blacklist Merchant Id: A list of disallowed merchant IDs. If the Merchant Id matches, then the transaction will be declined. If a Merchant Id is also in the allowed list, the blocked list will take precedence.
	BlockedMerchantIds interface{} `json:"blockedMerchantIds,omitempty"`
	// Whitelist Merchant Country: A list of allowed merchant countries, in ISO 3166-1 alpha-2 format. If the Merchant country does not match, then the transaction will be declined. If a Merchant Country is also provided in the blocked list, the blocked list will take precedence.
	AllowedMerchantCountries interface{} `json:"allowedMerchantCountries,omitempty"`
	// Blacklist Merchant Country: A list of disallowed merchant countries, in ISO 3166-1 alpha-2 format. If the Merchant country matches, then the transaction will be declined. If a Merchant Country is also in the allowed list, the blocked list will take precedence.
	BlockedMerchantCountries interface{} `json:"blockedMerchantCountries,omitempty"`
	// Indicates if a contactless transaction is allowed on the card.
	AllowContactless interface{} `json:"allowContactless,omitempty"`
	// Indicates if an ATM Withdrawal transaction is allowed on the card.
	AllowAtm interface{} `json:"allowAtm,omitempty"`
	// Indicates if an online transaction is allowed on the card.
	AllowECommerce interface{} `json:"allowECommerce,omitempty"`
	// Indicates if a cashback transaction at a physical terminal is allowed on the card.
	AllowCashback interface{} `json:"allowCashback,omitempty"`
	// Indicates if a the card can receive a credit transaction.
	AllowCreditAuthorisations interface{} `json:"allowCreditAuthorisations,omitempty"`
}

// NewCommonLevelSpendRules instantiates a new CommonLevelSpendRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonLevelSpendRules() *CommonLevelSpendRules {
	this := CommonLevelSpendRules{}
	return &this
}

// NewCommonLevelSpendRulesWithDefaults instantiates a new CommonLevelSpendRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonLevelSpendRulesWithDefaults() *CommonLevelSpendRules {
	this := CommonLevelSpendRules{}
	return &this
}

// GetAllowedMerchantCategories returns the AllowedMerchantCategories field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonLevelSpendRules) GetAllowedMerchantCategories() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AllowedMerchantCategories
}

// GetAllowedMerchantCategoriesOk returns a tuple with the AllowedMerchantCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonLevelSpendRules) GetAllowedMerchantCategoriesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AllowedMerchantCategories) {
		return nil, false
	}
	return &o.AllowedMerchantCategories, true
}

// HasAllowedMerchantCategories returns a boolean if a field has been set.
func (o *CommonLevelSpendRules) HasAllowedMerchantCategories() bool {
	if o != nil && IsNil(o.AllowedMerchantCategories) {
		return true
	}

	return false
}

// SetAllowedMerchantCategories gets a reference to the given interface{} and assigns it to the AllowedMerchantCategories field.
func (o *CommonLevelSpendRules) SetAllowedMerchantCategories(v interface{}) {
	o.AllowedMerchantCategories = v
}

// GetBlockedMerchantCategories returns the BlockedMerchantCategories field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonLevelSpendRules) GetBlockedMerchantCategories() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.BlockedMerchantCategories
}

// GetBlockedMerchantCategoriesOk returns a tuple with the BlockedMerchantCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonLevelSpendRules) GetBlockedMerchantCategoriesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.BlockedMerchantCategories) {
		return nil, false
	}
	return &o.BlockedMerchantCategories, true
}

// HasBlockedMerchantCategories returns a boolean if a field has been set.
func (o *CommonLevelSpendRules) HasBlockedMerchantCategories() bool {
	if o != nil && IsNil(o.BlockedMerchantCategories) {
		return true
	}

	return false
}

// SetBlockedMerchantCategories gets a reference to the given interface{} and assigns it to the BlockedMerchantCategories field.
func (o *CommonLevelSpendRules) SetBlockedMerchantCategories(v interface{}) {
	o.BlockedMerchantCategories = v
}

// GetAllowedMerchantIds returns the AllowedMerchantIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonLevelSpendRules) GetAllowedMerchantIds() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AllowedMerchantIds
}

// GetAllowedMerchantIdsOk returns a tuple with the AllowedMerchantIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonLevelSpendRules) GetAllowedMerchantIdsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AllowedMerchantIds) {
		return nil, false
	}
	return &o.AllowedMerchantIds, true
}

// HasAllowedMerchantIds returns a boolean if a field has been set.
func (o *CommonLevelSpendRules) HasAllowedMerchantIds() bool {
	if o != nil && IsNil(o.AllowedMerchantIds) {
		return true
	}

	return false
}

// SetAllowedMerchantIds gets a reference to the given interface{} and assigns it to the AllowedMerchantIds field.
func (o *CommonLevelSpendRules) SetAllowedMerchantIds(v interface{}) {
	o.AllowedMerchantIds = v
}

// GetBlockedMerchantIds returns the BlockedMerchantIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonLevelSpendRules) GetBlockedMerchantIds() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.BlockedMerchantIds
}

// GetBlockedMerchantIdsOk returns a tuple with the BlockedMerchantIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonLevelSpendRules) GetBlockedMerchantIdsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.BlockedMerchantIds) {
		return nil, false
	}
	return &o.BlockedMerchantIds, true
}

// HasBlockedMerchantIds returns a boolean if a field has been set.
func (o *CommonLevelSpendRules) HasBlockedMerchantIds() bool {
	if o != nil && IsNil(o.BlockedMerchantIds) {
		return true
	}

	return false
}

// SetBlockedMerchantIds gets a reference to the given interface{} and assigns it to the BlockedMerchantIds field.
func (o *CommonLevelSpendRules) SetBlockedMerchantIds(v interface{}) {
	o.BlockedMerchantIds = v
}

// GetAllowedMerchantCountries returns the AllowedMerchantCountries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonLevelSpendRules) GetAllowedMerchantCountries() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AllowedMerchantCountries
}

// GetAllowedMerchantCountriesOk returns a tuple with the AllowedMerchantCountries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonLevelSpendRules) GetAllowedMerchantCountriesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AllowedMerchantCountries) {
		return nil, false
	}
	return &o.AllowedMerchantCountries, true
}

// HasAllowedMerchantCountries returns a boolean if a field has been set.
func (o *CommonLevelSpendRules) HasAllowedMerchantCountries() bool {
	if o != nil && IsNil(o.AllowedMerchantCountries) {
		return true
	}

	return false
}

// SetAllowedMerchantCountries gets a reference to the given interface{} and assigns it to the AllowedMerchantCountries field.
func (o *CommonLevelSpendRules) SetAllowedMerchantCountries(v interface{}) {
	o.AllowedMerchantCountries = v
}

// GetBlockedMerchantCountries returns the BlockedMerchantCountries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonLevelSpendRules) GetBlockedMerchantCountries() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.BlockedMerchantCountries
}

// GetBlockedMerchantCountriesOk returns a tuple with the BlockedMerchantCountries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonLevelSpendRules) GetBlockedMerchantCountriesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.BlockedMerchantCountries) {
		return nil, false
	}
	return &o.BlockedMerchantCountries, true
}

// HasBlockedMerchantCountries returns a boolean if a field has been set.
func (o *CommonLevelSpendRules) HasBlockedMerchantCountries() bool {
	if o != nil && IsNil(o.BlockedMerchantCountries) {
		return true
	}

	return false
}

// SetBlockedMerchantCountries gets a reference to the given interface{} and assigns it to the BlockedMerchantCountries field.
func (o *CommonLevelSpendRules) SetBlockedMerchantCountries(v interface{}) {
	o.BlockedMerchantCountries = v
}

// GetAllowContactless returns the AllowContactless field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonLevelSpendRules) GetAllowContactless() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AllowContactless
}

// GetAllowContactlessOk returns a tuple with the AllowContactless field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonLevelSpendRules) GetAllowContactlessOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AllowContactless) {
		return nil, false
	}
	return &o.AllowContactless, true
}

// HasAllowContactless returns a boolean if a field has been set.
func (o *CommonLevelSpendRules) HasAllowContactless() bool {
	if o != nil && IsNil(o.AllowContactless) {
		return true
	}

	return false
}

// SetAllowContactless gets a reference to the given interface{} and assigns it to the AllowContactless field.
func (o *CommonLevelSpendRules) SetAllowContactless(v interface{}) {
	o.AllowContactless = v
}

// GetAllowAtm returns the AllowAtm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonLevelSpendRules) GetAllowAtm() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AllowAtm
}

// GetAllowAtmOk returns a tuple with the AllowAtm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonLevelSpendRules) GetAllowAtmOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AllowAtm) {
		return nil, false
	}
	return &o.AllowAtm, true
}

// HasAllowAtm returns a boolean if a field has been set.
func (o *CommonLevelSpendRules) HasAllowAtm() bool {
	if o != nil && IsNil(o.AllowAtm) {
		return true
	}

	return false
}

// SetAllowAtm gets a reference to the given interface{} and assigns it to the AllowAtm field.
func (o *CommonLevelSpendRules) SetAllowAtm(v interface{}) {
	o.AllowAtm = v
}

// GetAllowECommerce returns the AllowECommerce field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonLevelSpendRules) GetAllowECommerce() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AllowECommerce
}

// GetAllowECommerceOk returns a tuple with the AllowECommerce field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonLevelSpendRules) GetAllowECommerceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AllowECommerce) {
		return nil, false
	}
	return &o.AllowECommerce, true
}

// HasAllowECommerce returns a boolean if a field has been set.
func (o *CommonLevelSpendRules) HasAllowECommerce() bool {
	if o != nil && IsNil(o.AllowECommerce) {
		return true
	}

	return false
}

// SetAllowECommerce gets a reference to the given interface{} and assigns it to the AllowECommerce field.
func (o *CommonLevelSpendRules) SetAllowECommerce(v interface{}) {
	o.AllowECommerce = v
}

// GetAllowCashback returns the AllowCashback field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonLevelSpendRules) GetAllowCashback() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AllowCashback
}

// GetAllowCashbackOk returns a tuple with the AllowCashback field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonLevelSpendRules) GetAllowCashbackOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AllowCashback) {
		return nil, false
	}
	return &o.AllowCashback, true
}

// HasAllowCashback returns a boolean if a field has been set.
func (o *CommonLevelSpendRules) HasAllowCashback() bool {
	if o != nil && IsNil(o.AllowCashback) {
		return true
	}

	return false
}

// SetAllowCashback gets a reference to the given interface{} and assigns it to the AllowCashback field.
func (o *CommonLevelSpendRules) SetAllowCashback(v interface{}) {
	o.AllowCashback = v
}

// GetAllowCreditAuthorisations returns the AllowCreditAuthorisations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonLevelSpendRules) GetAllowCreditAuthorisations() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AllowCreditAuthorisations
}

// GetAllowCreditAuthorisationsOk returns a tuple with the AllowCreditAuthorisations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonLevelSpendRules) GetAllowCreditAuthorisationsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AllowCreditAuthorisations) {
		return nil, false
	}
	return &o.AllowCreditAuthorisations, true
}

// HasAllowCreditAuthorisations returns a boolean if a field has been set.
func (o *CommonLevelSpendRules) HasAllowCreditAuthorisations() bool {
	if o != nil && IsNil(o.AllowCreditAuthorisations) {
		return true
	}

	return false
}

// SetAllowCreditAuthorisations gets a reference to the given interface{} and assigns it to the AllowCreditAuthorisations field.
func (o *CommonLevelSpendRules) SetAllowCreditAuthorisations(v interface{}) {
	o.AllowCreditAuthorisations = v
}

func (o CommonLevelSpendRules) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonLevelSpendRules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowedMerchantCategories != nil {
		toSerialize["allowedMerchantCategories"] = o.AllowedMerchantCategories
	}
	if o.BlockedMerchantCategories != nil {
		toSerialize["blockedMerchantCategories"] = o.BlockedMerchantCategories
	}
	if o.AllowedMerchantIds != nil {
		toSerialize["allowedMerchantIds"] = o.AllowedMerchantIds
	}
	if o.BlockedMerchantIds != nil {
		toSerialize["blockedMerchantIds"] = o.BlockedMerchantIds
	}
	if o.AllowedMerchantCountries != nil {
		toSerialize["allowedMerchantCountries"] = o.AllowedMerchantCountries
	}
	if o.BlockedMerchantCountries != nil {
		toSerialize["blockedMerchantCountries"] = o.BlockedMerchantCountries
	}
	if o.AllowContactless != nil {
		toSerialize["allowContactless"] = o.AllowContactless
	}
	if o.AllowAtm != nil {
		toSerialize["allowAtm"] = o.AllowAtm
	}
	if o.AllowECommerce != nil {
		toSerialize["allowECommerce"] = o.AllowECommerce
	}
	if o.AllowCashback != nil {
		toSerialize["allowCashback"] = o.AllowCashback
	}
	if o.AllowCreditAuthorisations != nil {
		toSerialize["allowCreditAuthorisations"] = o.AllowCreditAuthorisations
	}
	return toSerialize, nil
}

type NullableCommonLevelSpendRules struct {
	value *CommonLevelSpendRules
	isSet bool
}

func (v NullableCommonLevelSpendRules) Get() *CommonLevelSpendRules {
	return v.value
}

func (v *NullableCommonLevelSpendRules) Set(val *CommonLevelSpendRules) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonLevelSpendRules) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonLevelSpendRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonLevelSpendRules(val *CommonLevelSpendRules) *NullableCommonLevelSpendRules {
	return &NullableCommonLevelSpendRules{value: val, isSet: true}
}

func (v NullableCommonLevelSpendRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonLevelSpendRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

