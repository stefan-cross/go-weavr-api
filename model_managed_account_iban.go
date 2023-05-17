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

// checks if the ManagedAccountIBAN type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedAccountIBAN{}

// ManagedAccountIBAN struct for ManagedAccountIBAN
type ManagedAccountIBAN struct {
	// The state of the Managed Account's IBAN as follows:   - UNALLOCATED: The Managed Account has never been assigned an IBAN. Use the _managedAccountsIBANUpgrade_ operation to assign an IBAN to a Managed Account.   - PENDING_ALLOCATION: The IBAN is being allocated to the Managed Account.   - ALLOCATED: An IBAN is allocated to the Managed Account. 
	State interface{} `json:"state"`
	// A list of bank account details associated with the IBAN. Multiple details can be provided if multiple IBAN providers are supported by your payment model.
	BankAccountDetails interface{} `json:"bankAccountDetails"`
}

// NewManagedAccountIBAN instantiates a new ManagedAccountIBAN object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedAccountIBAN(state interface{}, bankAccountDetails interface{}) *ManagedAccountIBAN {
	this := ManagedAccountIBAN{}
	this.State = state
	this.BankAccountDetails = bankAccountDetails
	return &this
}

// NewManagedAccountIBANWithDefaults instantiates a new ManagedAccountIBAN object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedAccountIBANWithDefaults() *ManagedAccountIBAN {
	this := ManagedAccountIBAN{}
	return &this
}

// GetState returns the State field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ManagedAccountIBAN) GetState() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedAccountIBAN) GetStateOk() (*interface{}, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ManagedAccountIBAN) SetState(v interface{}) {
	o.State = v
}

// GetBankAccountDetails returns the BankAccountDetails field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ManagedAccountIBAN) GetBankAccountDetails() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.BankAccountDetails
}

// GetBankAccountDetailsOk returns a tuple with the BankAccountDetails field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedAccountIBAN) GetBankAccountDetailsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.BankAccountDetails) {
		return nil, false
	}
	return &o.BankAccountDetails, true
}

// SetBankAccountDetails sets field value
func (o *ManagedAccountIBAN) SetBankAccountDetails(v interface{}) {
	o.BankAccountDetails = v
}

func (o ManagedAccountIBAN) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedAccountIBAN) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.BankAccountDetails != nil {
		toSerialize["bankAccountDetails"] = o.BankAccountDetails
	}
	return toSerialize, nil
}

type NullableManagedAccountIBAN struct {
	value *ManagedAccountIBAN
	isSet bool
}

func (v NullableManagedAccountIBAN) Get() *ManagedAccountIBAN {
	return v.value
}

func (v *NullableManagedAccountIBAN) Set(val *ManagedAccountIBAN) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedAccountIBAN) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedAccountIBAN) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedAccountIBAN(val *ManagedAccountIBAN) *NullableManagedAccountIBAN {
	return &NullableManagedAccountIBAN{value: val, isSet: true}
}

func (v NullableManagedAccountIBAN) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedAccountIBAN) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

