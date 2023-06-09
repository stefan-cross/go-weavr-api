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

// checks if the UserKyc200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserKyc200Response{}

// UserKyc200Response struct for UserKyc200Response
type UserKyc200Response struct {
	// The reference required to initialise the KYB UI Component.
	Reference interface{} `json:"reference"`
}

// NewUserKyc200Response instantiates a new UserKyc200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserKyc200Response(reference interface{}) *UserKyc200Response {
	this := UserKyc200Response{}
	this.Reference = reference
	return &this
}

// NewUserKyc200ResponseWithDefaults instantiates a new UserKyc200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserKyc200ResponseWithDefaults() *UserKyc200Response {
	this := UserKyc200Response{}
	return &this
}

// GetReference returns the Reference field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *UserKyc200Response) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserKyc200Response) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// SetReference sets field value
func (o *UserKyc200Response) SetReference(v interface{}) {
	o.Reference = v
}

func (o UserKyc200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserKyc200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	return toSerialize, nil
}

type NullableUserKyc200Response struct {
	value *UserKyc200Response
	isSet bool
}

func (v NullableUserKyc200Response) Get() *UserKyc200Response {
	return v.value
}

func (v *NullableUserKyc200Response) Set(val *UserKyc200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUserKyc200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUserKyc200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserKyc200Response(val *UserKyc200Response) *NullableUserKyc200Response {
	return &NullableUserKyc200Response{value: val, isSet: true}
}

func (v NullableUserKyc200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserKyc200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


