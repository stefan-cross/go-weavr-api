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

// checks if the PhysicalCardDetailsReplacement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PhysicalCardDetailsReplacement{}

// PhysicalCardDetailsReplacement Indicates if the physical card is replaced by another card.
type PhysicalCardDetailsReplacement struct {
	// The reason why the physical card was replaced.   - DAMAGED: The physical card was damaged and cannot be used at a physical terminal.   - LOST_STOLEN: The physical card was either lost or stolen and cannot be used.   - EXPIRED: The physical card expired. 
	ReplacementReason interface{} `json:"replacementReason,omitempty"`
	// The unique identifier of the new card that replaces this card.
	ReplacementId interface{} `json:"replacementId,omitempty"`
}

// NewPhysicalCardDetailsReplacement instantiates a new PhysicalCardDetailsReplacement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhysicalCardDetailsReplacement() *PhysicalCardDetailsReplacement {
	this := PhysicalCardDetailsReplacement{}
	return &this
}

// NewPhysicalCardDetailsReplacementWithDefaults instantiates a new PhysicalCardDetailsReplacement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhysicalCardDetailsReplacementWithDefaults() *PhysicalCardDetailsReplacement {
	this := PhysicalCardDetailsReplacement{}
	return &this
}

// GetReplacementReason returns the ReplacementReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhysicalCardDetailsReplacement) GetReplacementReason() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReplacementReason
}

// GetReplacementReasonOk returns a tuple with the ReplacementReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalCardDetailsReplacement) GetReplacementReasonOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReplacementReason) {
		return nil, false
	}
	return &o.ReplacementReason, true
}

// HasReplacementReason returns a boolean if a field has been set.
func (o *PhysicalCardDetailsReplacement) HasReplacementReason() bool {
	if o != nil && IsNil(o.ReplacementReason) {
		return true
	}

	return false
}

// SetReplacementReason gets a reference to the given interface{} and assigns it to the ReplacementReason field.
func (o *PhysicalCardDetailsReplacement) SetReplacementReason(v interface{}) {
	o.ReplacementReason = v
}

// GetReplacementId returns the ReplacementId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PhysicalCardDetailsReplacement) GetReplacementId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReplacementId
}

// GetReplacementIdOk returns a tuple with the ReplacementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PhysicalCardDetailsReplacement) GetReplacementIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReplacementId) {
		return nil, false
	}
	return &o.ReplacementId, true
}

// HasReplacementId returns a boolean if a field has been set.
func (o *PhysicalCardDetailsReplacement) HasReplacementId() bool {
	if o != nil && IsNil(o.ReplacementId) {
		return true
	}

	return false
}

// SetReplacementId gets a reference to the given interface{} and assigns it to the ReplacementId field.
func (o *PhysicalCardDetailsReplacement) SetReplacementId(v interface{}) {
	o.ReplacementId = v
}

func (o PhysicalCardDetailsReplacement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PhysicalCardDetailsReplacement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ReplacementReason != nil {
		toSerialize["replacementReason"] = o.ReplacementReason
	}
	if o.ReplacementId != nil {
		toSerialize["replacementId"] = o.ReplacementId
	}
	return toSerialize, nil
}

type NullablePhysicalCardDetailsReplacement struct {
	value *PhysicalCardDetailsReplacement
	isSet bool
}

func (v NullablePhysicalCardDetailsReplacement) Get() *PhysicalCardDetailsReplacement {
	return v.value
}

func (v *NullablePhysicalCardDetailsReplacement) Set(val *PhysicalCardDetailsReplacement) {
	v.value = val
	v.isSet = true
}

func (v NullablePhysicalCardDetailsReplacement) IsSet() bool {
	return v.isSet
}

func (v *NullablePhysicalCardDetailsReplacement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhysicalCardDetailsReplacement(val *PhysicalCardDetailsReplacement) *NullablePhysicalCardDetailsReplacement {
	return &NullablePhysicalCardDetailsReplacement{value: val, isSet: true}
}

func (v NullablePhysicalCardDetailsReplacement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhysicalCardDetailsReplacement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

