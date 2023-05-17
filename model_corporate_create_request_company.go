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

// checks if the CorporateCreateRequestCompany type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CorporateCreateRequestCompany{}

// CorporateCreateRequestCompany The details associated with the company being on-boarded. The details provided need to match exactly with the details provided during KYB. 
type CorporateCreateRequestCompany struct {
	Type CompanyType `json:"type"`
	BusinessAddress *Address `json:"businessAddress,omitempty"`
	// The registered name of the company.
	Name interface{} `json:"name"`
	// The registration number of the company.
	RegistrationNumber interface{} `json:"registrationNumber,omitempty"`
	// The country of company registration, in ISO 3166 alpha-2 format.
	RegistrationCountry interface{} `json:"registrationCountry"`
}

// NewCorporateCreateRequestCompany instantiates a new CorporateCreateRequestCompany object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCorporateCreateRequestCompany(type_ CompanyType, name interface{}, registrationCountry interface{}) *CorporateCreateRequestCompany {
	this := CorporateCreateRequestCompany{}
	this.Type = type_
	this.Name = name
	this.RegistrationCountry = registrationCountry
	return &this
}

// NewCorporateCreateRequestCompanyWithDefaults instantiates a new CorporateCreateRequestCompany object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCorporateCreateRequestCompanyWithDefaults() *CorporateCreateRequestCompany {
	this := CorporateCreateRequestCompany{}
	return &this
}

// GetType returns the Type field value
func (o *CorporateCreateRequestCompany) GetType() CompanyType {
	if o == nil {
		var ret CompanyType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CorporateCreateRequestCompany) GetTypeOk() (*CompanyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CorporateCreateRequestCompany) SetType(v CompanyType) {
	o.Type = v
}

// GetBusinessAddress returns the BusinessAddress field value if set, zero value otherwise.
func (o *CorporateCreateRequestCompany) GetBusinessAddress() Address {
	if o == nil || IsNil(o.BusinessAddress) {
		var ret Address
		return ret
	}
	return *o.BusinessAddress
}

// GetBusinessAddressOk returns a tuple with the BusinessAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorporateCreateRequestCompany) GetBusinessAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.BusinessAddress) {
		return nil, false
	}
	return o.BusinessAddress, true
}

// HasBusinessAddress returns a boolean if a field has been set.
func (o *CorporateCreateRequestCompany) HasBusinessAddress() bool {
	if o != nil && !IsNil(o.BusinessAddress) {
		return true
	}

	return false
}

// SetBusinessAddress gets a reference to the given Address and assigns it to the BusinessAddress field.
func (o *CorporateCreateRequestCompany) SetBusinessAddress(v Address) {
	o.BusinessAddress = &v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CorporateCreateRequestCompany) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CorporateCreateRequestCompany) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CorporateCreateRequestCompany) SetName(v interface{}) {
	o.Name = v
}

// GetRegistrationNumber returns the RegistrationNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CorporateCreateRequestCompany) GetRegistrationNumber() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RegistrationNumber
}

// GetRegistrationNumberOk returns a tuple with the RegistrationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CorporateCreateRequestCompany) GetRegistrationNumberOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RegistrationNumber) {
		return nil, false
	}
	return &o.RegistrationNumber, true
}

// HasRegistrationNumber returns a boolean if a field has been set.
func (o *CorporateCreateRequestCompany) HasRegistrationNumber() bool {
	if o != nil && IsNil(o.RegistrationNumber) {
		return true
	}

	return false
}

// SetRegistrationNumber gets a reference to the given interface{} and assigns it to the RegistrationNumber field.
func (o *CorporateCreateRequestCompany) SetRegistrationNumber(v interface{}) {
	o.RegistrationNumber = v
}

// GetRegistrationCountry returns the RegistrationCountry field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CorporateCreateRequestCompany) GetRegistrationCountry() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.RegistrationCountry
}

// GetRegistrationCountryOk returns a tuple with the RegistrationCountry field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CorporateCreateRequestCompany) GetRegistrationCountryOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RegistrationCountry) {
		return nil, false
	}
	return &o.RegistrationCountry, true
}

// SetRegistrationCountry sets field value
func (o *CorporateCreateRequestCompany) SetRegistrationCountry(v interface{}) {
	o.RegistrationCountry = v
}

func (o CorporateCreateRequestCompany) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CorporateCreateRequestCompany) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.BusinessAddress) {
		toSerialize["businessAddress"] = o.BusinessAddress
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RegistrationNumber != nil {
		toSerialize["registrationNumber"] = o.RegistrationNumber
	}
	if o.RegistrationCountry != nil {
		toSerialize["registrationCountry"] = o.RegistrationCountry
	}
	return toSerialize, nil
}

type NullableCorporateCreateRequestCompany struct {
	value *CorporateCreateRequestCompany
	isSet bool
}

func (v NullableCorporateCreateRequestCompany) Get() *CorporateCreateRequestCompany {
	return v.value
}

func (v *NullableCorporateCreateRequestCompany) Set(val *CorporateCreateRequestCompany) {
	v.value = val
	v.isSet = true
}

func (v NullableCorporateCreateRequestCompany) IsSet() bool {
	return v.isSet
}

func (v *NullableCorporateCreateRequestCompany) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCorporateCreateRequestCompany(val *CorporateCreateRequestCompany) *NullableCorporateCreateRequestCompany {
	return &NullableCorporateCreateRequestCompany{value: val, isSet: true}
}

func (v NullableCorporateCreateRequestCompany) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCorporateCreateRequestCompany) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


