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

// checks if the SepaBankDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SepaBankDetails{}

// SepaBankDetails Bank details used in case the wire transfer is executed over SEPA.
type SepaBankDetails struct {
	// International Bank Account Number, required for wire transfer over SEPA.
	Iban interface{} `json:"iban"`
	// BIC, required for wire transfer over SEPA.
	BankIdentifierCode interface{} `json:"bankIdentifierCode"`
}

// NewSepaBankDetails instantiates a new SepaBankDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSepaBankDetails(iban interface{}, bankIdentifierCode interface{}) *SepaBankDetails {
	this := SepaBankDetails{}
	this.Iban = iban
	this.BankIdentifierCode = bankIdentifierCode
	return &this
}

// NewSepaBankDetailsWithDefaults instantiates a new SepaBankDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSepaBankDetailsWithDefaults() *SepaBankDetails {
	this := SepaBankDetails{}
	return &this
}

// GetIban returns the Iban field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *SepaBankDetails) GetIban() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Iban
}

// GetIbanOk returns a tuple with the Iban field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SepaBankDetails) GetIbanOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Iban) {
		return nil, false
	}
	return &o.Iban, true
}

// SetIban sets field value
func (o *SepaBankDetails) SetIban(v interface{}) {
	o.Iban = v
}

// GetBankIdentifierCode returns the BankIdentifierCode field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *SepaBankDetails) GetBankIdentifierCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.BankIdentifierCode
}

// GetBankIdentifierCodeOk returns a tuple with the BankIdentifierCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SepaBankDetails) GetBankIdentifierCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.BankIdentifierCode) {
		return nil, false
	}
	return &o.BankIdentifierCode, true
}

// SetBankIdentifierCode sets field value
func (o *SepaBankDetails) SetBankIdentifierCode(v interface{}) {
	o.BankIdentifierCode = v
}

func (o SepaBankDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SepaBankDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Iban != nil {
		toSerialize["iban"] = o.Iban
	}
	if o.BankIdentifierCode != nil {
		toSerialize["bankIdentifierCode"] = o.BankIdentifierCode
	}
	return toSerialize, nil
}

type NullableSepaBankDetails struct {
	value *SepaBankDetails
	isSet bool
}

func (v NullableSepaBankDetails) Get() *SepaBankDetails {
	return v.value
}

func (v *NullableSepaBankDetails) Set(val *SepaBankDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSepaBankDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSepaBankDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSepaBankDetails(val *SepaBankDetails) *NullableSepaBankDetails {
	return &NullableSepaBankDetails{value: val, isSet: true}
}

func (v NullableSepaBankDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSepaBankDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

