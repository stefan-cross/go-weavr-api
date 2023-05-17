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

// checks if the PasswordCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordCreateRequest{}

// PasswordCreateRequest struct for PasswordCreateRequest
type PasswordCreateRequest struct {
	Password SensitivePassword `json:"password"`
}

// NewPasswordCreateRequest instantiates a new PasswordCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordCreateRequest(password SensitivePassword) *PasswordCreateRequest {
	this := PasswordCreateRequest{}
	this.Password = password
	return &this
}

// NewPasswordCreateRequestWithDefaults instantiates a new PasswordCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordCreateRequestWithDefaults() *PasswordCreateRequest {
	this := PasswordCreateRequest{}
	return &this
}

// GetPassword returns the Password field value
func (o *PasswordCreateRequest) GetPassword() SensitivePassword {
	if o == nil {
		var ret SensitivePassword
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *PasswordCreateRequest) GetPasswordOk() (*SensitivePassword, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *PasswordCreateRequest) SetPassword(v SensitivePassword) {
	o.Password = v
}

func (o PasswordCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

type NullablePasswordCreateRequest struct {
	value *PasswordCreateRequest
	isSet bool
}

func (v NullablePasswordCreateRequest) Get() *PasswordCreateRequest {
	return v.value
}

func (v *NullablePasswordCreateRequest) Set(val *PasswordCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordCreateRequest(val *PasswordCreateRequest) *NullablePasswordCreateRequest {
	return &NullablePasswordCreateRequest{value: val, isSet: true}
}

func (v NullablePasswordCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


