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

// checks if the DestroyedReason type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DestroyedReason{}

// DestroyedReason The reason why the instrument has been destroyed:   - SYSTEM: The platform or an administrator of the platform has destroyed the instrument.   - USER: The root, or an authorised user, of the identity owning the instrument has destroyed the instrument.   - LOST: The instrument was automatically destroyed as it was marked as lost.   - STOLEN: The instrument was automatically destroyed as it was marked as stolen.   - EXPIRED: The instrument was automatically destroyed as it expired. 
type DestroyedReason struct {
}

// NewDestroyedReason instantiates a new DestroyedReason object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestroyedReason() *DestroyedReason {
	this := DestroyedReason{}
	return &this
}

// NewDestroyedReasonWithDefaults instantiates a new DestroyedReason object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestroyedReasonWithDefaults() *DestroyedReason {
	this := DestroyedReason{}
	return &this
}

func (o DestroyedReason) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DestroyedReason) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableDestroyedReason struct {
	value *DestroyedReason
	isSet bool
}

func (v NullableDestroyedReason) Get() *DestroyedReason {
	return v.value
}

func (v *NullableDestroyedReason) Set(val *DestroyedReason) {
	v.value = val
	v.isSet = true
}

func (v NullableDestroyedReason) IsSet() bool {
	return v.isSet
}

func (v *NullableDestroyedReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestroyedReason(val *DestroyedReason) *NullableDestroyedReason {
	return &NullableDestroyedReason{value: val, isSet: true}
}

func (v NullableDestroyedReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestroyedReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

