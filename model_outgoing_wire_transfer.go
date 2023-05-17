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

// checks if the OutgoingWireTransfer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutgoingWireTransfer{}

// OutgoingWireTransfer struct for OutgoingWireTransfer
type OutgoingWireTransfer struct {
	// The unique identifier of an Outgoing Wire Transfer transaction.
	Id interface{} `json:"id"`
	// The profile Id which a specific identity, instrument or transaction type is linked to.  Profiles contain configuration and determine behavioral aspects of the newly created transaction, for example, fees that may apply.  You can have one or more profiles linked to your application, and these can be used to drive different behaviors according to your product's needs.  Profile Ids can be found in the Multi Portal, in the API Credentials page. 
	ProfileId interface{} `json:"profileId"`
	// The tag field is a custom field that can be used to search and filter.
	Tag interface{} `json:"tag,omitempty"`
	SourceInstrument OutgoingWireTransferSourceInstrument `json:"sourceInstrument"`
	// The amount, in same currency as source and destination, that was transferred from the source.
	TransferAmount CurrencyAmount `json:"transferAmount"`
	//  Transaction description.
	Description interface{} `json:"description,omitempty"`
	// The wire transfer rails used, from SEPA, Faster Payments or SWIFT.
	Type interface{} `json:"type"`
	Destination OutgoingWireTransferBeneficiary `json:"destination"`
	State OutgoingWireTransferState `json:"state"`
	RejectedInfo *OutgoingWireTransferRejectedInfo `json:"rejectedInfo,omitempty"`
	// The time when the transaction was created, expressed in Epoch timestamp using millisecond precision.
	CreationTimestamp interface{} `json:"creationTimestamp"`
}

// NewOutgoingWireTransfer instantiates a new OutgoingWireTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutgoingWireTransfer(id interface{}, profileId interface{}, sourceInstrument OutgoingWireTransferSourceInstrument, transferAmount CurrencyAmount, type_ interface{}, destination OutgoingWireTransferBeneficiary, state OutgoingWireTransferState, creationTimestamp interface{}) *OutgoingWireTransfer {
	this := OutgoingWireTransfer{}
	this.Id = id
	this.ProfileId = profileId
	this.SourceInstrument = sourceInstrument
	this.TransferAmount = transferAmount
	this.Type = type_
	this.Destination = destination
	this.State = state
	this.CreationTimestamp = creationTimestamp
	return &this
}

// NewOutgoingWireTransferWithDefaults instantiates a new OutgoingWireTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutgoingWireTransferWithDefaults() *OutgoingWireTransfer {
	this := OutgoingWireTransfer{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *OutgoingWireTransfer) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OutgoingWireTransfer) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OutgoingWireTransfer) SetId(v interface{}) {
	o.Id = v
}

// GetProfileId returns the ProfileId field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *OutgoingWireTransfer) GetProfileId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OutgoingWireTransfer) GetProfileIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return &o.ProfileId, true
}

// SetProfileId sets field value
func (o *OutgoingWireTransfer) SetProfileId(v interface{}) {
	o.ProfileId = v
}

// GetTag returns the Tag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OutgoingWireTransfer) GetTag() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OutgoingWireTransfer) GetTagOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return &o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *OutgoingWireTransfer) HasTag() bool {
	if o != nil && IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given interface{} and assigns it to the Tag field.
func (o *OutgoingWireTransfer) SetTag(v interface{}) {
	o.Tag = v
}

// GetSourceInstrument returns the SourceInstrument field value
func (o *OutgoingWireTransfer) GetSourceInstrument() OutgoingWireTransferSourceInstrument {
	if o == nil {
		var ret OutgoingWireTransferSourceInstrument
		return ret
	}

	return o.SourceInstrument
}

// GetSourceInstrumentOk returns a tuple with the SourceInstrument field value
// and a boolean to check if the value has been set.
func (o *OutgoingWireTransfer) GetSourceInstrumentOk() (*OutgoingWireTransferSourceInstrument, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceInstrument, true
}

// SetSourceInstrument sets field value
func (o *OutgoingWireTransfer) SetSourceInstrument(v OutgoingWireTransferSourceInstrument) {
	o.SourceInstrument = v
}

// GetTransferAmount returns the TransferAmount field value
func (o *OutgoingWireTransfer) GetTransferAmount() CurrencyAmount {
	if o == nil {
		var ret CurrencyAmount
		return ret
	}

	return o.TransferAmount
}

// GetTransferAmountOk returns a tuple with the TransferAmount field value
// and a boolean to check if the value has been set.
func (o *OutgoingWireTransfer) GetTransferAmountOk() (*CurrencyAmount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransferAmount, true
}

// SetTransferAmount sets field value
func (o *OutgoingWireTransfer) SetTransferAmount(v CurrencyAmount) {
	o.TransferAmount = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OutgoingWireTransfer) GetDescription() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OutgoingWireTransfer) GetDescriptionOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return &o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OutgoingWireTransfer) HasDescription() bool {
	if o != nil && IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given interface{} and assigns it to the Description field.
func (o *OutgoingWireTransfer) SetDescription(v interface{}) {
	o.Description = v
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *OutgoingWireTransfer) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OutgoingWireTransfer) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OutgoingWireTransfer) SetType(v interface{}) {
	o.Type = v
}

// GetDestination returns the Destination field value
func (o *OutgoingWireTransfer) GetDestination() OutgoingWireTransferBeneficiary {
	if o == nil {
		var ret OutgoingWireTransferBeneficiary
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *OutgoingWireTransfer) GetDestinationOk() (*OutgoingWireTransferBeneficiary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *OutgoingWireTransfer) SetDestination(v OutgoingWireTransferBeneficiary) {
	o.Destination = v
}

// GetState returns the State field value
func (o *OutgoingWireTransfer) GetState() OutgoingWireTransferState {
	if o == nil {
		var ret OutgoingWireTransferState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *OutgoingWireTransfer) GetStateOk() (*OutgoingWireTransferState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *OutgoingWireTransfer) SetState(v OutgoingWireTransferState) {
	o.State = v
}

// GetRejectedInfo returns the RejectedInfo field value if set, zero value otherwise.
func (o *OutgoingWireTransfer) GetRejectedInfo() OutgoingWireTransferRejectedInfo {
	if o == nil || IsNil(o.RejectedInfo) {
		var ret OutgoingWireTransferRejectedInfo
		return ret
	}
	return *o.RejectedInfo
}

// GetRejectedInfoOk returns a tuple with the RejectedInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutgoingWireTransfer) GetRejectedInfoOk() (*OutgoingWireTransferRejectedInfo, bool) {
	if o == nil || IsNil(o.RejectedInfo) {
		return nil, false
	}
	return o.RejectedInfo, true
}

// HasRejectedInfo returns a boolean if a field has been set.
func (o *OutgoingWireTransfer) HasRejectedInfo() bool {
	if o != nil && !IsNil(o.RejectedInfo) {
		return true
	}

	return false
}

// SetRejectedInfo gets a reference to the given OutgoingWireTransferRejectedInfo and assigns it to the RejectedInfo field.
func (o *OutgoingWireTransfer) SetRejectedInfo(v OutgoingWireTransferRejectedInfo) {
	o.RejectedInfo = &v
}

// GetCreationTimestamp returns the CreationTimestamp field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *OutgoingWireTransfer) GetCreationTimestamp() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.CreationTimestamp
}

// GetCreationTimestampOk returns a tuple with the CreationTimestamp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OutgoingWireTransfer) GetCreationTimestampOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreationTimestamp) {
		return nil, false
	}
	return &o.CreationTimestamp, true
}

// SetCreationTimestamp sets field value
func (o *OutgoingWireTransfer) SetCreationTimestamp(v interface{}) {
	o.CreationTimestamp = v
}

func (o OutgoingWireTransfer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutgoingWireTransfer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ProfileId != nil {
		toSerialize["profileId"] = o.ProfileId
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	toSerialize["sourceInstrument"] = o.SourceInstrument
	toSerialize["transferAmount"] = o.TransferAmount
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["destination"] = o.Destination
	toSerialize["state"] = o.State
	if !IsNil(o.RejectedInfo) {
		toSerialize["rejectedInfo"] = o.RejectedInfo
	}
	if o.CreationTimestamp != nil {
		toSerialize["creationTimestamp"] = o.CreationTimestamp
	}
	return toSerialize, nil
}

type NullableOutgoingWireTransfer struct {
	value *OutgoingWireTransfer
	isSet bool
}

func (v NullableOutgoingWireTransfer) Get() *OutgoingWireTransfer {
	return v.value
}

func (v *NullableOutgoingWireTransfer) Set(val *OutgoingWireTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableOutgoingWireTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableOutgoingWireTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutgoingWireTransfer(val *OutgoingWireTransfer) *NullableOutgoingWireTransfer {
	return &NullableOutgoingWireTransfer{value: val, isSet: true}
}

func (v NullableOutgoingWireTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutgoingWireTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

