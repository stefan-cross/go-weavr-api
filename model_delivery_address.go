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

// checks if the DeliveryAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryAddress{}

// DeliveryAddress The delivery address where the physical card is delivered.
type DeliveryAddress struct {
	Name interface{} `json:"name"`
	Surname interface{} `json:"surname"`
	AddressLine1 interface{} `json:"addressLine1"`
	AddressLine2 interface{} `json:"addressLine2,omitempty"`
	City interface{} `json:"city"`
	PostCode interface{} `json:"postCode"`
	State interface{} `json:"state,omitempty"`
	// Country of the identity in ISO 3166 alpha-2 format.
	Country interface{} `json:"country"`
}

// NewDeliveryAddress instantiates a new DeliveryAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryAddress(name interface{}, surname interface{}, addressLine1 interface{}, city interface{}, postCode interface{}, country interface{}) *DeliveryAddress {
	this := DeliveryAddress{}
	this.Name = name
	this.Surname = surname
	this.AddressLine1 = addressLine1
	this.City = city
	this.PostCode = postCode
	this.Country = country
	return &this
}

// NewDeliveryAddressWithDefaults instantiates a new DeliveryAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryAddressWithDefaults() *DeliveryAddress {
	this := DeliveryAddress{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *DeliveryAddress) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryAddress) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeliveryAddress) SetName(v interface{}) {
	o.Name = v
}

// GetSurname returns the Surname field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *DeliveryAddress) GetSurname() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Surname
}

// GetSurnameOk returns a tuple with the Surname field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryAddress) GetSurnameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Surname) {
		return nil, false
	}
	return &o.Surname, true
}

// SetSurname sets field value
func (o *DeliveryAddress) SetSurname(v interface{}) {
	o.Surname = v
}

// GetAddressLine1 returns the AddressLine1 field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *DeliveryAddress) GetAddressLine1() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryAddress) GetAddressLine1Ok() (*interface{}, bool) {
	if o == nil || IsNil(o.AddressLine1) {
		return nil, false
	}
	return &o.AddressLine1, true
}

// SetAddressLine1 sets field value
func (o *DeliveryAddress) SetAddressLine1(v interface{}) {
	o.AddressLine1 = v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeliveryAddress) GetAddressLine2() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryAddress) GetAddressLine2Ok() (*interface{}, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return &o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *DeliveryAddress) HasAddressLine2() bool {
	if o != nil && IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given interface{} and assigns it to the AddressLine2 field.
func (o *DeliveryAddress) SetAddressLine2(v interface{}) {
	o.AddressLine2 = v
}

// GetCity returns the City field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *DeliveryAddress) GetCity() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryAddress) GetCityOk() (*interface{}, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *DeliveryAddress) SetCity(v interface{}) {
	o.City = v
}

// GetPostCode returns the PostCode field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *DeliveryAddress) GetPostCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.PostCode
}

// GetPostCodeOk returns a tuple with the PostCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryAddress) GetPostCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PostCode) {
		return nil, false
	}
	return &o.PostCode, true
}

// SetPostCode sets field value
func (o *DeliveryAddress) SetPostCode(v interface{}) {
	o.PostCode = v
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeliveryAddress) GetState() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryAddress) GetStateOk() (*interface{}, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return &o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DeliveryAddress) HasState() bool {
	if o != nil && IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given interface{} and assigns it to the State field.
func (o *DeliveryAddress) SetState(v interface{}) {
	o.State = v
}

// GetCountry returns the Country field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *DeliveryAddress) GetCountry() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryAddress) GetCountryOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *DeliveryAddress) SetCountry(v interface{}) {
	o.Country = v
}

func (o DeliveryAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Surname != nil {
		toSerialize["surname"] = o.Surname
	}
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

type NullableDeliveryAddress struct {
	value *DeliveryAddress
	isSet bool
}

func (v NullableDeliveryAddress) Get() *DeliveryAddress {
	return v.value
}

func (v *NullableDeliveryAddress) Set(val *DeliveryAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryAddress(val *DeliveryAddress) *NullableDeliveryAddress {
	return &NullableDeliveryAddress{value: val, isSet: true}
}

func (v NullableDeliveryAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

