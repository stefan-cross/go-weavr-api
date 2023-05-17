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

// checks if the OutgoingWireTransferRejectedInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutgoingWireTransferRejectedInfo{}

// OutgoingWireTransferRejectedInfo Additional info when the outgoing wire transfer has been rejected, as follows:   - SYSTEM: The wire transfer was rejected by the system.   - USER: The wire transfer was rejected by the user. 
type OutgoingWireTransferRejectedInfo struct {
}

// NewOutgoingWireTransferRejectedInfo instantiates a new OutgoingWireTransferRejectedInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutgoingWireTransferRejectedInfo() *OutgoingWireTransferRejectedInfo {
	this := OutgoingWireTransferRejectedInfo{}
	return &this
}

// NewOutgoingWireTransferRejectedInfoWithDefaults instantiates a new OutgoingWireTransferRejectedInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutgoingWireTransferRejectedInfoWithDefaults() *OutgoingWireTransferRejectedInfo {
	this := OutgoingWireTransferRejectedInfo{}
	return &this
}

func (o OutgoingWireTransferRejectedInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutgoingWireTransferRejectedInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableOutgoingWireTransferRejectedInfo struct {
	value *OutgoingWireTransferRejectedInfo
	isSet bool
}

func (v NullableOutgoingWireTransferRejectedInfo) Get() *OutgoingWireTransferRejectedInfo {
	return v.value
}

func (v *NullableOutgoingWireTransferRejectedInfo) Set(val *OutgoingWireTransferRejectedInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOutgoingWireTransferRejectedInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOutgoingWireTransferRejectedInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutgoingWireTransferRejectedInfo(val *OutgoingWireTransferRejectedInfo) *NullableOutgoingWireTransferRejectedInfo {
	return &NullableOutgoingWireTransferRejectedInfo{value: val, isSet: true}
}

func (v NullableOutgoingWireTransferRejectedInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutgoingWireTransferRejectedInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

