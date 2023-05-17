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

// checks if the CardAuthorisationDetailsAuthorisationType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardAuthorisationDetailsAuthorisationType{}

// CardAuthorisationDetailsAuthorisationType struct for CardAuthorisationDetailsAuthorisationType
type CardAuthorisationDetailsAuthorisationType struct {
}

// NewCardAuthorisationDetailsAuthorisationType instantiates a new CardAuthorisationDetailsAuthorisationType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardAuthorisationDetailsAuthorisationType() *CardAuthorisationDetailsAuthorisationType {
	this := CardAuthorisationDetailsAuthorisationType{}
	return &this
}

// NewCardAuthorisationDetailsAuthorisationTypeWithDefaults instantiates a new CardAuthorisationDetailsAuthorisationType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardAuthorisationDetailsAuthorisationTypeWithDefaults() *CardAuthorisationDetailsAuthorisationType {
	this := CardAuthorisationDetailsAuthorisationType{}
	return &this
}

func (o CardAuthorisationDetailsAuthorisationType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardAuthorisationDetailsAuthorisationType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableCardAuthorisationDetailsAuthorisationType struct {
	value *CardAuthorisationDetailsAuthorisationType
	isSet bool
}

func (v NullableCardAuthorisationDetailsAuthorisationType) Get() *CardAuthorisationDetailsAuthorisationType {
	return v.value
}

func (v *NullableCardAuthorisationDetailsAuthorisationType) Set(val *CardAuthorisationDetailsAuthorisationType) {
	v.value = val
	v.isSet = true
}

func (v NullableCardAuthorisationDetailsAuthorisationType) IsSet() bool {
	return v.isSet
}

func (v *NullableCardAuthorisationDetailsAuthorisationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardAuthorisationDetailsAuthorisationType(val *CardAuthorisationDetailsAuthorisationType) *NullableCardAuthorisationDetailsAuthorisationType {
	return &NullableCardAuthorisationDetailsAuthorisationType{value: val, isSet: true}
}

func (v NullableCardAuthorisationDetailsAuthorisationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardAuthorisationDetailsAuthorisationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

