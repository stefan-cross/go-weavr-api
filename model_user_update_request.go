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

// checks if the UserUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserUpdateRequest{}

// UserUpdateRequest struct for UserUpdateRequest
type UserUpdateRequest struct {
	// The first name of the user.
	Name interface{} `json:"name,omitempty"`
	// The last name of the user.
	Surname interface{} `json:"surname,omitempty"`
	// E-mail Address of the user
	Email interface{} `json:"email,omitempty"`
	Mobile *Mobile `json:"mobile,omitempty"`
	// Date of birth of the authorised user.
	DateOfBirth *Date `json:"dateOfBirth,omitempty"`
}

// NewUserUpdateRequest instantiates a new UserUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserUpdateRequest() *UserUpdateRequest {
	this := UserUpdateRequest{}
	return &this
}

// NewUserUpdateRequestWithDefaults instantiates a new UserUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserUpdateRequestWithDefaults() *UserUpdateRequest {
	this := UserUpdateRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUpdateRequest) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUpdateRequest) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserUpdateRequest) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *UserUpdateRequest) SetName(v interface{}) {
	o.Name = v
}

// GetSurname returns the Surname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUpdateRequest) GetSurname() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Surname
}

// GetSurnameOk returns a tuple with the Surname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUpdateRequest) GetSurnameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Surname) {
		return nil, false
	}
	return &o.Surname, true
}

// HasSurname returns a boolean if a field has been set.
func (o *UserUpdateRequest) HasSurname() bool {
	if o != nil && IsNil(o.Surname) {
		return true
	}

	return false
}

// SetSurname gets a reference to the given interface{} and assigns it to the Surname field.
func (o *UserUpdateRequest) SetSurname(v interface{}) {
	o.Surname = v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUpdateRequest) GetEmail() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUpdateRequest) GetEmailOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return &o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserUpdateRequest) HasEmail() bool {
	if o != nil && IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given interface{} and assigns it to the Email field.
func (o *UserUpdateRequest) SetEmail(v interface{}) {
	o.Email = v
}

// GetMobile returns the Mobile field value if set, zero value otherwise.
func (o *UserUpdateRequest) GetMobile() Mobile {
	if o == nil || IsNil(o.Mobile) {
		var ret Mobile
		return ret
	}
	return *o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdateRequest) GetMobileOk() (*Mobile, bool) {
	if o == nil || IsNil(o.Mobile) {
		return nil, false
	}
	return o.Mobile, true
}

// HasMobile returns a boolean if a field has been set.
func (o *UserUpdateRequest) HasMobile() bool {
	if o != nil && !IsNil(o.Mobile) {
		return true
	}

	return false
}

// SetMobile gets a reference to the given Mobile and assigns it to the Mobile field.
func (o *UserUpdateRequest) SetMobile(v Mobile) {
	o.Mobile = &v
}

// GetDateOfBirth returns the DateOfBirth field value if set, zero value otherwise.
func (o *UserUpdateRequest) GetDateOfBirth() Date {
	if o == nil || IsNil(o.DateOfBirth) {
		var ret Date
		return ret
	}
	return *o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdateRequest) GetDateOfBirthOk() (*Date, bool) {
	if o == nil || IsNil(o.DateOfBirth) {
		return nil, false
	}
	return o.DateOfBirth, true
}

// HasDateOfBirth returns a boolean if a field has been set.
func (o *UserUpdateRequest) HasDateOfBirth() bool {
	if o != nil && !IsNil(o.DateOfBirth) {
		return true
	}

	return false
}

// SetDateOfBirth gets a reference to the given Date and assigns it to the DateOfBirth field.
func (o *UserUpdateRequest) SetDateOfBirth(v Date) {
	o.DateOfBirth = &v
}

func (o UserUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Surname != nil {
		toSerialize["surname"] = o.Surname
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Mobile) {
		toSerialize["mobile"] = o.Mobile
	}
	if !IsNil(o.DateOfBirth) {
		toSerialize["dateOfBirth"] = o.DateOfBirth
	}
	return toSerialize, nil
}

type NullableUserUpdateRequest struct {
	value *UserUpdateRequest
	isSet bool
}

func (v NullableUserUpdateRequest) Get() *UserUpdateRequest {
	return v.value
}

func (v *NullableUserUpdateRequest) Set(val *UserUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserUpdateRequest(val *UserUpdateRequest) *NullableUserUpdateRequest {
	return &NullableUserUpdateRequest{value: val, isSet: true}
}

func (v NullableUserUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

