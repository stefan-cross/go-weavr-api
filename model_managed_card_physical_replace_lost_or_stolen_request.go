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

// checks if the ManagedCardPhysicalReplaceLostOrStolenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedCardPhysicalReplaceLostOrStolenRequest{}

// ManagedCardPhysicalReplaceLostOrStolenRequest struct for ManagedCardPhysicalReplaceLostOrStolenRequest
type ManagedCardPhysicalReplaceLostOrStolenRequest struct {
	// A unique code to be used to activate the replacement physical card.
	ActivationCode interface{} `json:"activationCode"`
}

// NewManagedCardPhysicalReplaceLostOrStolenRequest instantiates a new ManagedCardPhysicalReplaceLostOrStolenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedCardPhysicalReplaceLostOrStolenRequest(activationCode interface{}) *ManagedCardPhysicalReplaceLostOrStolenRequest {
	this := ManagedCardPhysicalReplaceLostOrStolenRequest{}
	this.ActivationCode = activationCode
	return &this
}

// NewManagedCardPhysicalReplaceLostOrStolenRequestWithDefaults instantiates a new ManagedCardPhysicalReplaceLostOrStolenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedCardPhysicalReplaceLostOrStolenRequestWithDefaults() *ManagedCardPhysicalReplaceLostOrStolenRequest {
	this := ManagedCardPhysicalReplaceLostOrStolenRequest{}
	return &this
}

// GetActivationCode returns the ActivationCode field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ManagedCardPhysicalReplaceLostOrStolenRequest) GetActivationCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.ActivationCode
}

// GetActivationCodeOk returns a tuple with the ActivationCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedCardPhysicalReplaceLostOrStolenRequest) GetActivationCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ActivationCode) {
		return nil, false
	}
	return &o.ActivationCode, true
}

// SetActivationCode sets field value
func (o *ManagedCardPhysicalReplaceLostOrStolenRequest) SetActivationCode(v interface{}) {
	o.ActivationCode = v
}

func (o ManagedCardPhysicalReplaceLostOrStolenRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedCardPhysicalReplaceLostOrStolenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ActivationCode != nil {
		toSerialize["activationCode"] = o.ActivationCode
	}
	return toSerialize, nil
}

type NullableManagedCardPhysicalReplaceLostOrStolenRequest struct {
	value *ManagedCardPhysicalReplaceLostOrStolenRequest
	isSet bool
}

func (v NullableManagedCardPhysicalReplaceLostOrStolenRequest) Get() *ManagedCardPhysicalReplaceLostOrStolenRequest {
	return v.value
}

func (v *NullableManagedCardPhysicalReplaceLostOrStolenRequest) Set(val *ManagedCardPhysicalReplaceLostOrStolenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedCardPhysicalReplaceLostOrStolenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedCardPhysicalReplaceLostOrStolenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedCardPhysicalReplaceLostOrStolenRequest(val *ManagedCardPhysicalReplaceLostOrStolenRequest) *NullableManagedCardPhysicalReplaceLostOrStolenRequest {
	return &NullableManagedCardPhysicalReplaceLostOrStolenRequest{value: val, isSet: true}
}

func (v NullableManagedCardPhysicalReplaceLostOrStolenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedCardPhysicalReplaceLostOrStolenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


