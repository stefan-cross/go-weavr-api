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

// checks if the MerchantData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantData{}

// MerchantData struct for MerchantData
type MerchantData struct {
	// The name of the merchant where the authorisation has been made.
	MerchantName interface{} `json:"merchantName"`
	// The merchant category code.
	MerchantCategoryCode interface{} `json:"merchantCategoryCode"`
	// The merchant ID.
	MerchantId interface{} `json:"merchantId,omitempty"`
	// The merchant description
	MerchantDescription interface{} `json:"merchantDescription,omitempty"`
	// The merchant street address (if available)
	MerchantStreet interface{} `json:"merchantStreet,omitempty"`
	// The merchant city (if available)
	MerchantCity interface{} `json:"merchantCity,omitempty"`
	// The merchant state address (if available)
	MerchantState interface{} `json:"merchantState,omitempty"`
	// The merchant postal code (if available)
	MerchantPostalCode interface{} `json:"merchantPostalCode,omitempty"`
	// The merchant country address
	MerchantCountry interface{} `json:"merchantCountry,omitempty"`
	// The merchant telephone number
	MerchantTelephone interface{} `json:"merchantTelephone,omitempty"`
	// The merchant URL
	MerchantURL interface{} `json:"merchantURL,omitempty"`
	// The alternative merchant name
	MerchantNameOther interface{} `json:"merchantNameOther,omitempty"`
	// The merchant's network id
	MerchantNetworkId interface{} `json:"merchantNetworkId,omitempty"`
	// The merchant's contact (if available)
	MerchantContact interface{} `json:"merchantContact,omitempty"`
}

// NewMerchantData instantiates a new MerchantData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantData(merchantName interface{}, merchantCategoryCode interface{}) *MerchantData {
	this := MerchantData{}
	this.MerchantName = merchantName
	this.MerchantCategoryCode = merchantCategoryCode
	return &this
}

// NewMerchantDataWithDefaults instantiates a new MerchantData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantDataWithDefaults() *MerchantData {
	this := MerchantData{}
	return &this
}

// GetMerchantName returns the MerchantName field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *MerchantData) GetMerchantName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.MerchantName
}

// GetMerchantNameOk returns a tuple with the MerchantName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantData) GetMerchantNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MerchantName) {
		return nil, false
	}
	return &o.MerchantName, true
}

// SetMerchantName sets field value
func (o *MerchantData) SetMerchantName(v interface{}) {
	o.MerchantName = v
}

// GetMerchantCategoryCode returns the MerchantCategoryCode field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *MerchantData) GetMerchantCategoryCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.MerchantCategoryCode
}

// GetMerchantCategoryCodeOk returns a tuple with the MerchantCategoryCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantData) GetMerchantCategoryCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MerchantCategoryCode) {
		return nil, false
	}
	return &o.MerchantCategoryCode, true
}

// SetMerchantCategoryCode sets field value
func (o *MerchantData) SetMerchantCategoryCode(v interface{}) {
	o.MerchantCategoryCode = v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantData) GetMerchantId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantData) GetMerchantIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return &o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *MerchantData) HasMerchantId() bool {
	if o != nil && IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given interface{} and assigns it to the MerchantId field.
func (o *MerchantData) SetMerchantId(v interface{}) {
	o.MerchantId = v
}

// GetMerchantDescription returns the MerchantDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantData) GetMerchantDescription() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MerchantDescription
}

// GetMerchantDescriptionOk returns a tuple with the MerchantDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantData) GetMerchantDescriptionOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MerchantDescription) {
		return nil, false
	}
	return &o.MerchantDescription, true
}

// HasMerchantDescription returns a boolean if a field has been set.
func (o *MerchantData) HasMerchantDescription() bool {
	if o != nil && IsNil(o.MerchantDescription) {
		return true
	}

	return false
}

// SetMerchantDescription gets a reference to the given interface{} and assigns it to the MerchantDescription field.
func (o *MerchantData) SetMerchantDescription(v interface{}) {
	o.MerchantDescription = v
}

// GetMerchantStreet returns the MerchantStreet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantData) GetMerchantStreet() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MerchantStreet
}

// GetMerchantStreetOk returns a tuple with the MerchantStreet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantData) GetMerchantStreetOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MerchantStreet) {
		return nil, false
	}
	return &o.MerchantStreet, true
}

// HasMerchantStreet returns a boolean if a field has been set.
func (o *MerchantData) HasMerchantStreet() bool {
	if o != nil && IsNil(o.MerchantStreet) {
		return true
	}

	return false
}

// SetMerchantStreet gets a reference to the given interface{} and assigns it to the MerchantStreet field.
func (o *MerchantData) SetMerchantStreet(v interface{}) {
	o.MerchantStreet = v
}

// GetMerchantCity returns the MerchantCity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantData) GetMerchantCity() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MerchantCity
}

// GetMerchantCityOk returns a tuple with the MerchantCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantData) GetMerchantCityOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MerchantCity) {
		return nil, false
	}
	return &o.MerchantCity, true
}

// HasMerchantCity returns a boolean if a field has been set.
func (o *MerchantData) HasMerchantCity() bool {
	if o != nil && IsNil(o.MerchantCity) {
		return true
	}

	return false
}

// SetMerchantCity gets a reference to the given interface{} and assigns it to the MerchantCity field.
func (o *MerchantData) SetMerchantCity(v interface{}) {
	o.MerchantCity = v
}

// GetMerchantState returns the MerchantState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantData) GetMerchantState() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MerchantState
}

// GetMerchantStateOk returns a tuple with the MerchantState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantData) GetMerchantStateOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MerchantState) {
		return nil, false
	}
	return &o.MerchantState, true
}

// HasMerchantState returns a boolean if a field has been set.
func (o *MerchantData) HasMerchantState() bool {
	if o != nil && IsNil(o.MerchantState) {
		return true
	}

	return false
}

// SetMerchantState gets a reference to the given interface{} and assigns it to the MerchantState field.
func (o *MerchantData) SetMerchantState(v interface{}) {
	o.MerchantState = v
}

// GetMerchantPostalCode returns the MerchantPostalCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantData) GetMerchantPostalCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MerchantPostalCode
}

// GetMerchantPostalCodeOk returns a tuple with the MerchantPostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantData) GetMerchantPostalCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MerchantPostalCode) {
		return nil, false
	}
	return &o.MerchantPostalCode, true
}

// HasMerchantPostalCode returns a boolean if a field has been set.
func (o *MerchantData) HasMerchantPostalCode() bool {
	if o != nil && IsNil(o.MerchantPostalCode) {
		return true
	}

	return false
}

// SetMerchantPostalCode gets a reference to the given interface{} and assigns it to the MerchantPostalCode field.
func (o *MerchantData) SetMerchantPostalCode(v interface{}) {
	o.MerchantPostalCode = v
}

// GetMerchantCountry returns the MerchantCountry field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantData) GetMerchantCountry() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MerchantCountry
}

// GetMerchantCountryOk returns a tuple with the MerchantCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantData) GetMerchantCountryOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MerchantCountry) {
		return nil, false
	}
	return &o.MerchantCountry, true
}

// HasMerchantCountry returns a boolean if a field has been set.
func (o *MerchantData) HasMerchantCountry() bool {
	if o != nil && IsNil(o.MerchantCountry) {
		return true
	}

	return false
}

// SetMerchantCountry gets a reference to the given interface{} and assigns it to the MerchantCountry field.
func (o *MerchantData) SetMerchantCountry(v interface{}) {
	o.MerchantCountry = v
}

// GetMerchantTelephone returns the MerchantTelephone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantData) GetMerchantTelephone() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MerchantTelephone
}

// GetMerchantTelephoneOk returns a tuple with the MerchantTelephone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantData) GetMerchantTelephoneOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MerchantTelephone) {
		return nil, false
	}
	return &o.MerchantTelephone, true
}

// HasMerchantTelephone returns a boolean if a field has been set.
func (o *MerchantData) HasMerchantTelephone() bool {
	if o != nil && IsNil(o.MerchantTelephone) {
		return true
	}

	return false
}

// SetMerchantTelephone gets a reference to the given interface{} and assigns it to the MerchantTelephone field.
func (o *MerchantData) SetMerchantTelephone(v interface{}) {
	o.MerchantTelephone = v
}

// GetMerchantURL returns the MerchantURL field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantData) GetMerchantURL() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MerchantURL
}

// GetMerchantURLOk returns a tuple with the MerchantURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantData) GetMerchantURLOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MerchantURL) {
		return nil, false
	}
	return &o.MerchantURL, true
}

// HasMerchantURL returns a boolean if a field has been set.
func (o *MerchantData) HasMerchantURL() bool {
	if o != nil && IsNil(o.MerchantURL) {
		return true
	}

	return false
}

// SetMerchantURL gets a reference to the given interface{} and assigns it to the MerchantURL field.
func (o *MerchantData) SetMerchantURL(v interface{}) {
	o.MerchantURL = v
}

// GetMerchantNameOther returns the MerchantNameOther field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantData) GetMerchantNameOther() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MerchantNameOther
}

// GetMerchantNameOtherOk returns a tuple with the MerchantNameOther field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantData) GetMerchantNameOtherOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MerchantNameOther) {
		return nil, false
	}
	return &o.MerchantNameOther, true
}

// HasMerchantNameOther returns a boolean if a field has been set.
func (o *MerchantData) HasMerchantNameOther() bool {
	if o != nil && IsNil(o.MerchantNameOther) {
		return true
	}

	return false
}

// SetMerchantNameOther gets a reference to the given interface{} and assigns it to the MerchantNameOther field.
func (o *MerchantData) SetMerchantNameOther(v interface{}) {
	o.MerchantNameOther = v
}

// GetMerchantNetworkId returns the MerchantNetworkId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantData) GetMerchantNetworkId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MerchantNetworkId
}

// GetMerchantNetworkIdOk returns a tuple with the MerchantNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantData) GetMerchantNetworkIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MerchantNetworkId) {
		return nil, false
	}
	return &o.MerchantNetworkId, true
}

// HasMerchantNetworkId returns a boolean if a field has been set.
func (o *MerchantData) HasMerchantNetworkId() bool {
	if o != nil && IsNil(o.MerchantNetworkId) {
		return true
	}

	return false
}

// SetMerchantNetworkId gets a reference to the given interface{} and assigns it to the MerchantNetworkId field.
func (o *MerchantData) SetMerchantNetworkId(v interface{}) {
	o.MerchantNetworkId = v
}

// GetMerchantContact returns the MerchantContact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantData) GetMerchantContact() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.MerchantContact
}

// GetMerchantContactOk returns a tuple with the MerchantContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantData) GetMerchantContactOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MerchantContact) {
		return nil, false
	}
	return &o.MerchantContact, true
}

// HasMerchantContact returns a boolean if a field has been set.
func (o *MerchantData) HasMerchantContact() bool {
	if o != nil && IsNil(o.MerchantContact) {
		return true
	}

	return false
}

// SetMerchantContact gets a reference to the given interface{} and assigns it to the MerchantContact field.
func (o *MerchantData) SetMerchantContact(v interface{}) {
	o.MerchantContact = v
}

func (o MerchantData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MerchantName != nil {
		toSerialize["merchantName"] = o.MerchantName
	}
	if o.MerchantCategoryCode != nil {
		toSerialize["merchantCategoryCode"] = o.MerchantCategoryCode
	}
	if o.MerchantId != nil {
		toSerialize["merchantId"] = o.MerchantId
	}
	if o.MerchantDescription != nil {
		toSerialize["merchantDescription"] = o.MerchantDescription
	}
	if o.MerchantStreet != nil {
		toSerialize["merchantStreet"] = o.MerchantStreet
	}
	if o.MerchantCity != nil {
		toSerialize["merchantCity"] = o.MerchantCity
	}
	if o.MerchantState != nil {
		toSerialize["merchantState"] = o.MerchantState
	}
	if o.MerchantPostalCode != nil {
		toSerialize["merchantPostalCode"] = o.MerchantPostalCode
	}
	if o.MerchantCountry != nil {
		toSerialize["merchantCountry"] = o.MerchantCountry
	}
	if o.MerchantTelephone != nil {
		toSerialize["merchantTelephone"] = o.MerchantTelephone
	}
	if o.MerchantURL != nil {
		toSerialize["merchantURL"] = o.MerchantURL
	}
	if o.MerchantNameOther != nil {
		toSerialize["merchantNameOther"] = o.MerchantNameOther
	}
	if o.MerchantNetworkId != nil {
		toSerialize["merchantNetworkId"] = o.MerchantNetworkId
	}
	if o.MerchantContact != nil {
		toSerialize["merchantContact"] = o.MerchantContact
	}
	return toSerialize, nil
}

type NullableMerchantData struct {
	value *MerchantData
	isSet bool
}

func (v NullableMerchantData) Get() *MerchantData {
	return v.value
}

func (v *NullableMerchantData) Set(val *MerchantData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantData(val *MerchantData) *NullableMerchantData {
	return &NullableMerchantData{value: val, isSet: true}
}

func (v NullableMerchantData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


