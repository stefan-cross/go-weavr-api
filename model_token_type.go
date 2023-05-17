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

// checks if the TokenType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenType{}

// TokenType The auth token received can only be used to access the following endpoints:   - `/identities`   - `/access_token` 
type TokenType struct {
}

// NewTokenType instantiates a new TokenType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenType() *TokenType {
	this := TokenType{}
	return &this
}

// NewTokenTypeWithDefaults instantiates a new TokenType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenTypeWithDefaults() *TokenType {
	this := TokenType{}
	return &this
}

func (o TokenType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableTokenType struct {
	value *TokenType
	isSet bool
}

func (v NullableTokenType) Get() *TokenType {
	return v.value
}

func (v *NullableTokenType) Set(val *TokenType) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenType) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenType(val *TokenType) *NullableTokenType {
	return &NullableTokenType{value: val, isSet: true}
}

func (v NullableTokenType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

