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

// checks if the AvailableToSpend type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailableToSpend{}

// AvailableToSpend struct for AvailableToSpend
type AvailableToSpend struct {
	// The amount and currency that is available to spend, (for the given interval).
	Value *CurrencyAmount `json:"value,omitempty"`
	Interval *SpendLimitInterval `json:"interval,omitempty"`
}

// NewAvailableToSpend instantiates a new AvailableToSpend object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailableToSpend() *AvailableToSpend {
	this := AvailableToSpend{}
	return &this
}

// NewAvailableToSpendWithDefaults instantiates a new AvailableToSpend object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailableToSpendWithDefaults() *AvailableToSpend {
	this := AvailableToSpend{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AvailableToSpend) GetValue() CurrencyAmount {
	if o == nil || IsNil(o.Value) {
		var ret CurrencyAmount
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableToSpend) GetValueOk() (*CurrencyAmount, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AvailableToSpend) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given CurrencyAmount and assigns it to the Value field.
func (o *AvailableToSpend) SetValue(v CurrencyAmount) {
	o.Value = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *AvailableToSpend) GetInterval() SpendLimitInterval {
	if o == nil || IsNil(o.Interval) {
		var ret SpendLimitInterval
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableToSpend) GetIntervalOk() (*SpendLimitInterval, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *AvailableToSpend) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given SpendLimitInterval and assigns it to the Interval field.
func (o *AvailableToSpend) SetInterval(v SpendLimitInterval) {
	o.Interval = &v
}

func (o AvailableToSpend) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvailableToSpend) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	return toSerialize, nil
}

type NullableAvailableToSpend struct {
	value *AvailableToSpend
	isSet bool
}

func (v NullableAvailableToSpend) Get() *AvailableToSpend {
	return v.value
}

func (v *NullableAvailableToSpend) Set(val *AvailableToSpend) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailableToSpend) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailableToSpend) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailableToSpend(val *AvailableToSpend) *NullableAvailableToSpend {
	return &NullableAvailableToSpend{value: val, isSet: true}
}

func (v NullableAvailableToSpend) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailableToSpend) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


