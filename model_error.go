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

// checks if the Error type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Error{}

// Error struct for Error
type Error struct {
	Code interface{} `json:"code,omitempty"`
	Message interface{} `json:"message,omitempty"`
}

// NewError instantiates a new Error object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewError() *Error {
	this := Error{}
	return &this
}

// NewErrorWithDefaults instantiates a new Error object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorWithDefaults() *Error {
	this := Error{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Error) GetCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Error) GetCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return &o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Error) HasCode() bool {
	if o != nil && IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given interface{} and assigns it to the Code field.
func (o *Error) SetCode(v interface{}) {
	o.Code = v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Error) GetMessage() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Error) GetMessageOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return &o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Error) HasMessage() bool {
	if o != nil && IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given interface{} and assigns it to the Message field.
func (o *Error) SetMessage(v interface{}) {
	o.Message = v
}

func (o Error) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Error) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableError struct {
	value *Error
	isSet bool
}

func (v NullableError) Get() *Error {
	return v.value
}

func (v *NullableError) Set(val *Error) {
	v.value = val
	v.isSet = true
}

func (v NullableError) IsSet() bool {
	return v.isSet
}

func (v *NullableError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableError(val *Error) *NullableError {
	return &NullableError{value: val, isSet: true}
}

func (v NullableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


