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

// checks if the UserIdentity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserIdentity{}

// UserIdentity struct for UserIdentity
type UserIdentity struct {
	// The unique identifier of the Corporate Identity.
	Id IdentityId `json:"id"`
	// The name of the company.
	Name interface{} `json:"name"`
}

// NewUserIdentity instantiates a new UserIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserIdentity(id IdentityId, name interface{}) *UserIdentity {
	this := UserIdentity{}
	this.Id = id
	this.Name = name
	return &this
}

// NewUserIdentityWithDefaults instantiates a new UserIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserIdentityWithDefaults() *UserIdentity {
	this := UserIdentity{}
	return &this
}

// GetId returns the Id field value
func (o *UserIdentity) GetId() IdentityId {
	if o == nil {
		var ret IdentityId
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserIdentity) GetIdOk() (*IdentityId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserIdentity) SetId(v IdentityId) {
	o.Id = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *UserIdentity) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserIdentity) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserIdentity) SetName(v interface{}) {
	o.Name = v
}

func (o UserIdentity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserIdentity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableUserIdentity struct {
	value *UserIdentity
	isSet bool
}

func (v NullableUserIdentity) Get() *UserIdentity {
	return v.value
}

func (v *NullableUserIdentity) Set(val *UserIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableUserIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableUserIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserIdentity(val *UserIdentity) *NullableUserIdentity {
	return &NullableUserIdentity{value: val, isSet: true}
}

func (v NullableUserIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


