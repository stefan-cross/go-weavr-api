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

// checks if the SCAPushChannel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SCAPushChannel{}

// SCAPushChannel - \"AUTHY\": The push notification is sent on the user's device using [Twilio Authy](https://www.twilio.com/authy) - \"BIOMETRIC\": The push notification is sent to the user's device 
type SCAPushChannel struct {
}

// NewSCAPushChannel instantiates a new SCAPushChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSCAPushChannel() *SCAPushChannel {
	this := SCAPushChannel{}
	return &this
}

// NewSCAPushChannelWithDefaults instantiates a new SCAPushChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSCAPushChannelWithDefaults() *SCAPushChannel {
	this := SCAPushChannel{}
	return &this
}

func (o SCAPushChannel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SCAPushChannel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableSCAPushChannel struct {
	value *SCAPushChannel
	isSet bool
}

func (v NullableSCAPushChannel) Get() *SCAPushChannel {
	return v.value
}

func (v *NullableSCAPushChannel) Set(val *SCAPushChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableSCAPushChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableSCAPushChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSCAPushChannel(val *SCAPushChannel) *NullableSCAPushChannel {
	return &NullableSCAPushChannel{value: val, isSet: true}
}

func (v NullableSCAPushChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSCAPushChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


