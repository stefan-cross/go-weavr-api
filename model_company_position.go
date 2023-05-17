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

// checks if the CompanyPosition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyPosition{}

// CompanyPosition The company position of the Corporate Root User.
type CompanyPosition struct {
}

// NewCompanyPosition instantiates a new CompanyPosition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyPosition() *CompanyPosition {
	this := CompanyPosition{}
	return &this
}

// NewCompanyPositionWithDefaults instantiates a new CompanyPosition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyPositionWithDefaults() *CompanyPosition {
	this := CompanyPosition{}
	return &this
}

func (o CompanyPosition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyPosition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableCompanyPosition struct {
	value *CompanyPosition
	isSet bool
}

func (v NullableCompanyPosition) Get() *CompanyPosition {
	return v.value
}

func (v *NullableCompanyPosition) Set(val *CompanyPosition) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyPosition) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyPosition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyPosition(val *CompanyPosition) *NullableCompanyPosition {
	return &NullableCompanyPosition{value: val, isSet: true}
}

func (v NullableCompanyPosition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyPosition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

