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

// checks if the SCAFactor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SCAFactor{}

// SCAFactor struct for SCAFactor
type SCAFactor struct {
	Type *SCAChallengeType `json:"type,omitempty"`
	Status *SCAFactorStatus `json:"status,omitempty"`
	Channel *SCAChannel `json:"channel,omitempty"`
}

// NewSCAFactor instantiates a new SCAFactor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSCAFactor() *SCAFactor {
	this := SCAFactor{}
	return &this
}

// NewSCAFactorWithDefaults instantiates a new SCAFactor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSCAFactorWithDefaults() *SCAFactor {
	this := SCAFactor{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SCAFactor) GetType() SCAChallengeType {
	if o == nil || IsNil(o.Type) {
		var ret SCAChallengeType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCAFactor) GetTypeOk() (*SCAChallengeType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SCAFactor) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given SCAChallengeType and assigns it to the Type field.
func (o *SCAFactor) SetType(v SCAChallengeType) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SCAFactor) GetStatus() SCAFactorStatus {
	if o == nil || IsNil(o.Status) {
		var ret SCAFactorStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCAFactor) GetStatusOk() (*SCAFactorStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SCAFactor) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given SCAFactorStatus and assigns it to the Status field.
func (o *SCAFactor) SetStatus(v SCAFactorStatus) {
	o.Status = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *SCAFactor) GetChannel() SCAChannel {
	if o == nil || IsNil(o.Channel) {
		var ret SCAChannel
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SCAFactor) GetChannelOk() (*SCAChannel, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *SCAFactor) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given SCAChannel and assigns it to the Channel field.
func (o *SCAFactor) SetChannel(v SCAChannel) {
	o.Channel = &v
}

func (o SCAFactor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SCAFactor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	return toSerialize, nil
}

type NullableSCAFactor struct {
	value *SCAFactor
	isSet bool
}

func (v NullableSCAFactor) Get() *SCAFactor {
	return v.value
}

func (v *NullableSCAFactor) Set(val *SCAFactor) {
	v.value = val
	v.isSet = true
}

func (v NullableSCAFactor) IsSet() bool {
	return v.isSet
}

func (v *NullableSCAFactor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSCAFactor(val *SCAFactor) *NullableSCAFactor {
	return &NullableSCAFactor{value: val, isSet: true}
}

func (v NullableSCAFactor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSCAFactor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


