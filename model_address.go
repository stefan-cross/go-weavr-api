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

// checks if the Address type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Address{}

// Address struct for Address
type Address struct {
	// The first line of the address.
	AddressLine1 interface{} `json:"addressLine1"`
	// The second line of the address.
	AddressLine2 interface{} `json:"addressLine2,omitempty"`
	// The city of the address.
	City interface{} `json:"city"`
	// The post cost associated with the address.
	PostCode interface{} `json:"postCode"`
	// The state of the address.
	State interface{} `json:"state,omitempty"`
	// The country of the address expressed in ISO 3166 alpha-2 format.
	Country interface{} `json:"country"`
}

// NewAddress instantiates a new Address object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddress(addressLine1 interface{}, city interface{}, postCode interface{}, country interface{}) *Address {
	this := Address{}
	this.AddressLine1 = addressLine1
	this.City = city
	this.PostCode = postCode
	this.Country = country
	return &this
}

// NewAddressWithDefaults instantiates a new Address object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressWithDefaults() *Address {
	this := Address{}
	return &this
}

// GetAddressLine1 returns the AddressLine1 field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *Address) GetAddressLine1() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetAddressLine1Ok() (*interface{}, bool) {
	if o == nil || IsNil(o.AddressLine1) {
		return nil, false
	}
	return &o.AddressLine1, true
}

// SetAddressLine1 sets field value
func (o *Address) SetAddressLine1(v interface{}) {
	o.AddressLine1 = v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetAddressLine2() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetAddressLine2Ok() (*interface{}, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return &o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *Address) HasAddressLine2() bool {
	if o != nil && IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given interface{} and assigns it to the AddressLine2 field.
func (o *Address) SetAddressLine2(v interface{}) {
	o.AddressLine2 = v
}

// GetCity returns the City field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *Address) GetCity() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetCityOk() (*interface{}, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *Address) SetCity(v interface{}) {
	o.City = v
}

// GetPostCode returns the PostCode field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *Address) GetPostCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.PostCode
}

// GetPostCodeOk returns a tuple with the PostCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetPostCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PostCode) {
		return nil, false
	}
	return &o.PostCode, true
}

// SetPostCode sets field value
func (o *Address) SetPostCode(v interface{}) {
	o.PostCode = v
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetState() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetStateOk() (*interface{}, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return &o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Address) HasState() bool {
	if o != nil && IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given interface{} and assigns it to the State field.
func (o *Address) SetState(v interface{}) {
	o.State = v
}

// GetCountry returns the Country field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *Address) GetCountry() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetCountryOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *Address) SetCountry(v interface{}) {
	o.Country = v
}

func (o Address) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Address) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AddressLine1 != nil {
		toSerialize["addressLine1"] = o.AddressLine1
	}
	if o.AddressLine2 != nil {
		toSerialize["addressLine2"] = o.AddressLine2
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.PostCode != nil {
		toSerialize["postCode"] = o.PostCode
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	return toSerialize, nil
}

type NullableAddress struct {
	value *Address
	isSet bool
}

func (v NullableAddress) Get() *Address {
	return v.value
}

func (v *NullableAddress) Set(val *Address) {
	v.value = val
	v.isSet = true
}

func (v NullableAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddress(val *Address) *NullableAddress {
	return &NullableAddress{value: val, isSet: true}
}

func (v NullableAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


