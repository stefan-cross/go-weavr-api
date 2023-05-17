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

// checks if the ManagedCardPhysicalReplaceDamagedRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedCardPhysicalReplaceDamagedRequest{}

// ManagedCardPhysicalReplaceDamagedRequest struct for ManagedCardPhysicalReplaceDamagedRequest
type ManagedCardPhysicalReplaceDamagedRequest struct {
	// The code that will be used to activate the physical card replacement.
	ActivationCode interface{} `json:"activationCode"`
}

// NewManagedCardPhysicalReplaceDamagedRequest instantiates a new ManagedCardPhysicalReplaceDamagedRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedCardPhysicalReplaceDamagedRequest(activationCode interface{}) *ManagedCardPhysicalReplaceDamagedRequest {
	this := ManagedCardPhysicalReplaceDamagedRequest{}
	this.ActivationCode = activationCode
	return &this
}

// NewManagedCardPhysicalReplaceDamagedRequestWithDefaults instantiates a new ManagedCardPhysicalReplaceDamagedRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedCardPhysicalReplaceDamagedRequestWithDefaults() *ManagedCardPhysicalReplaceDamagedRequest {
	this := ManagedCardPhysicalReplaceDamagedRequest{}
	return &this
}

// GetActivationCode returns the ActivationCode field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ManagedCardPhysicalReplaceDamagedRequest) GetActivationCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.ActivationCode
}

// GetActivationCodeOk returns a tuple with the ActivationCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedCardPhysicalReplaceDamagedRequest) GetActivationCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ActivationCode) {
		return nil, false
	}
	return &o.ActivationCode, true
}

// SetActivationCode sets field value
func (o *ManagedCardPhysicalReplaceDamagedRequest) SetActivationCode(v interface{}) {
	o.ActivationCode = v
}

func (o ManagedCardPhysicalReplaceDamagedRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedCardPhysicalReplaceDamagedRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ActivationCode != nil {
		toSerialize["activationCode"] = o.ActivationCode
	}
	return toSerialize, nil
}

type NullableManagedCardPhysicalReplaceDamagedRequest struct {
	value *ManagedCardPhysicalReplaceDamagedRequest
	isSet bool
}

func (v NullableManagedCardPhysicalReplaceDamagedRequest) Get() *ManagedCardPhysicalReplaceDamagedRequest {
	return v.value
}

func (v *NullableManagedCardPhysicalReplaceDamagedRequest) Set(val *ManagedCardPhysicalReplaceDamagedRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedCardPhysicalReplaceDamagedRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedCardPhysicalReplaceDamagedRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedCardPhysicalReplaceDamagedRequest(val *ManagedCardPhysicalReplaceDamagedRequest) *NullableManagedCardPhysicalReplaceDamagedRequest {
	return &NullableManagedCardPhysicalReplaceDamagedRequest{value: val, isSet: true}
}

func (v NullableManagedCardPhysicalReplaceDamagedRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedCardPhysicalReplaceDamagedRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

