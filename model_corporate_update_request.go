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

// checks if the CorporateUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CorporateUpdateRequest{}

// CorporateUpdateRequest struct for CorporateUpdateRequest
type CorporateUpdateRequest struct {
	// The tag field is a custom field that can be used to search and filter.
	Tag interface{} `json:"tag,omitempty"`
	// Deprecated
	Industry *Industry `json:"industry,omitempty"`
	// Deprecated
	SourceOfFunds *CorporateSourceOfFunds `json:"sourceOfFunds,omitempty"`
	// Description of source of funds in case `OTHER` was chosen.
	// Deprecated
	SourceOfFundsOther interface{} `json:"sourceOfFundsOther,omitempty"`
	CompanyBusinessAddress *Address `json:"companyBusinessAddress,omitempty"`
	// The fee group which the Corporate will be bound to. Do not specify this if you are not using fee groups.
	FeeGroup interface{} `json:"feeGroup,omitempty"`
	// The currency expressed in ISO-4217 code. Example: GBP, EUR, USD.
	BaseCurrency interface{} `json:"baseCurrency,omitempty"`
	// The first name of the Corporate root user.
	Name interface{} `json:"name,omitempty"`
	// The last name of the Corporate root user.
	Surname interface{} `json:"surname,omitempty"`
	// E-mail Address of the user
	Email interface{} `json:"email,omitempty"`
	Mobile *Mobile `json:"mobile,omitempty"`
	// Date of birth of the authorised user.
	DateOfBirth *Date `json:"dateOfBirth,omitempty"`
	ResetMobileCounter interface{} `json:"resetMobileCounter,omitempty"`
}

// NewCorporateUpdateRequest instantiates a new CorporateUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCorporateUpdateRequest() *CorporateUpdateRequest {
	this := CorporateUpdateRequest{}
	return &this
}

// NewCorporateUpdateRequestWithDefaults instantiates a new CorporateUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCorporateUpdateRequestWithDefaults() *CorporateUpdateRequest {
	this := CorporateUpdateRequest{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CorporateUpdateRequest) GetTag() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CorporateUpdateRequest) GetTagOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return &o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *CorporateUpdateRequest) HasTag() bool {
	if o != nil && IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given interface{} and assigns it to the Tag field.
func (o *CorporateUpdateRequest) SetTag(v interface{}) {
	o.Tag = v
}

// GetIndustry returns the Industry field value if set, zero value otherwise.
// Deprecated
func (o *CorporateUpdateRequest) GetIndustry() Industry {
	if o == nil || IsNil(o.Industry) {
		var ret Industry
		return ret
	}
	return *o.Industry
}

// GetIndustryOk returns a tuple with the Industry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CorporateUpdateRequest) GetIndustryOk() (*Industry, bool) {
	if o == nil || IsNil(o.Industry) {
		return nil, false
	}
	return o.Industry, true
}

// HasIndustry returns a boolean if a field has been set.
func (o *CorporateUpdateRequest) HasIndustry() bool {
	if o != nil && !IsNil(o.Industry) {
		return true
	}

	return false
}

// SetIndustry gets a reference to the given Industry and assigns it to the Industry field.
// Deprecated
func (o *CorporateUpdateRequest) SetIndustry(v Industry) {
	o.Industry = &v
}

// GetSourceOfFunds returns the SourceOfFunds field value if set, zero value otherwise.
// Deprecated
func (o *CorporateUpdateRequest) GetSourceOfFunds() CorporateSourceOfFunds {
	if o == nil || IsNil(o.SourceOfFunds) {
		var ret CorporateSourceOfFunds
		return ret
	}
	return *o.SourceOfFunds
}

// GetSourceOfFundsOk returns a tuple with the SourceOfFunds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CorporateUpdateRequest) GetSourceOfFundsOk() (*CorporateSourceOfFunds, bool) {
	if o == nil || IsNil(o.SourceOfFunds) {
		return nil, false
	}
	return o.SourceOfFunds, true
}

// HasSourceOfFunds returns a boolean if a field has been set.
func (o *CorporateUpdateRequest) HasSourceOfFunds() bool {
	if o != nil && !IsNil(o.SourceOfFunds) {
		return true
	}

	return false
}

// SetSourceOfFunds gets a reference to the given CorporateSourceOfFunds and assigns it to the SourceOfFunds field.
// Deprecated
func (o *CorporateUpdateRequest) SetSourceOfFunds(v CorporateSourceOfFunds) {
	o.SourceOfFunds = &v
}

// GetSourceOfFundsOther returns the SourceOfFundsOther field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *CorporateUpdateRequest) GetSourceOfFundsOther() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SourceOfFundsOther
}

// GetSourceOfFundsOtherOk returns a tuple with the SourceOfFundsOther field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *CorporateUpdateRequest) GetSourceOfFundsOtherOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SourceOfFundsOther) {
		return nil, false
	}
	return &o.SourceOfFundsOther, true
}

// HasSourceOfFundsOther returns a boolean if a field has been set.
func (o *CorporateUpdateRequest) HasSourceOfFundsOther() bool {
	if o != nil && IsNil(o.SourceOfFundsOther) {
		return true
	}

	return false
}

// SetSourceOfFundsOther gets a reference to the given interface{} and assigns it to the SourceOfFundsOther field.
// Deprecated
func (o *CorporateUpdateRequest) SetSourceOfFundsOther(v interface{}) {
	o.SourceOfFundsOther = v
}

// GetCompanyBusinessAddress returns the CompanyBusinessAddress field value if set, zero value otherwise.
func (o *CorporateUpdateRequest) GetCompanyBusinessAddress() Address {
	if o == nil || IsNil(o.CompanyBusinessAddress) {
		var ret Address
		return ret
	}
	return *o.CompanyBusinessAddress
}

// GetCompanyBusinessAddressOk returns a tuple with the CompanyBusinessAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorporateUpdateRequest) GetCompanyBusinessAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.CompanyBusinessAddress) {
		return nil, false
	}
	return o.CompanyBusinessAddress, true
}

// HasCompanyBusinessAddress returns a boolean if a field has been set.
func (o *CorporateUpdateRequest) HasCompanyBusinessAddress() bool {
	if o != nil && !IsNil(o.CompanyBusinessAddress) {
		return true
	}

	return false
}

// SetCompanyBusinessAddress gets a reference to the given Address and assigns it to the CompanyBusinessAddress field.
func (o *CorporateUpdateRequest) SetCompanyBusinessAddress(v Address) {
	o.CompanyBusinessAddress = &v
}

// GetFeeGroup returns the FeeGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CorporateUpdateRequest) GetFeeGroup() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FeeGroup
}

// GetFeeGroupOk returns a tuple with the FeeGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CorporateUpdateRequest) GetFeeGroupOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FeeGroup) {
		return nil, false
	}
	return &o.FeeGroup, true
}

// HasFeeGroup returns a boolean if a field has been set.
func (o *CorporateUpdateRequest) HasFeeGroup() bool {
	if o != nil && IsNil(o.FeeGroup) {
		return true
	}

	return false
}

// SetFeeGroup gets a reference to the given interface{} and assigns it to the FeeGroup field.
func (o *CorporateUpdateRequest) SetFeeGroup(v interface{}) {
	o.FeeGroup = v
}

// GetBaseCurrency returns the BaseCurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CorporateUpdateRequest) GetBaseCurrency() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.BaseCurrency
}

// GetBaseCurrencyOk returns a tuple with the BaseCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CorporateUpdateRequest) GetBaseCurrencyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.BaseCurrency) {
		return nil, false
	}
	return &o.BaseCurrency, true
}

// HasBaseCurrency returns a boolean if a field has been set.
func (o *CorporateUpdateRequest) HasBaseCurrency() bool {
	if o != nil && IsNil(o.BaseCurrency) {
		return true
	}

	return false
}

// SetBaseCurrency gets a reference to the given interface{} and assigns it to the BaseCurrency field.
func (o *CorporateUpdateRequest) SetBaseCurrency(v interface{}) {
	o.BaseCurrency = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CorporateUpdateRequest) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CorporateUpdateRequest) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CorporateUpdateRequest) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *CorporateUpdateRequest) SetName(v interface{}) {
	o.Name = v
}

// GetSurname returns the Surname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CorporateUpdateRequest) GetSurname() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Surname
}

// GetSurnameOk returns a tuple with the Surname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CorporateUpdateRequest) GetSurnameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Surname) {
		return nil, false
	}
	return &o.Surname, true
}

// HasSurname returns a boolean if a field has been set.
func (o *CorporateUpdateRequest) HasSurname() bool {
	if o != nil && IsNil(o.Surname) {
		return true
	}

	return false
}

// SetSurname gets a reference to the given interface{} and assigns it to the Surname field.
func (o *CorporateUpdateRequest) SetSurname(v interface{}) {
	o.Surname = v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CorporateUpdateRequest) GetEmail() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CorporateUpdateRequest) GetEmailOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return &o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CorporateUpdateRequest) HasEmail() bool {
	if o != nil && IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given interface{} and assigns it to the Email field.
func (o *CorporateUpdateRequest) SetEmail(v interface{}) {
	o.Email = v
}

// GetMobile returns the Mobile field value if set, zero value otherwise.
func (o *CorporateUpdateRequest) GetMobile() Mobile {
	if o == nil || IsNil(o.Mobile) {
		var ret Mobile
		return ret
	}
	return *o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorporateUpdateRequest) GetMobileOk() (*Mobile, bool) {
	if o == nil || IsNil(o.Mobile) {
		return nil, false
	}
	return o.Mobile, true
}

// HasMobile returns a boolean if a field has been set.
func (o *CorporateUpdateRequest) HasMobile() bool {
	if o != nil && !IsNil(o.Mobile) {
		return true
	}

	return false
}

// SetMobile gets a reference to the given Mobile and assigns it to the Mobile field.
func (o *CorporateUpdateRequest) SetMobile(v Mobile) {
	o.Mobile = &v
}

// GetDateOfBirth returns the DateOfBirth field value if set, zero value otherwise.
func (o *CorporateUpdateRequest) GetDateOfBirth() Date {
	if o == nil || IsNil(o.DateOfBirth) {
		var ret Date
		return ret
	}
	return *o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorporateUpdateRequest) GetDateOfBirthOk() (*Date, bool) {
	if o == nil || IsNil(o.DateOfBirth) {
		return nil, false
	}
	return o.DateOfBirth, true
}

// HasDateOfBirth returns a boolean if a field has been set.
func (o *CorporateUpdateRequest) HasDateOfBirth() bool {
	if o != nil && !IsNil(o.DateOfBirth) {
		return true
	}

	return false
}

// SetDateOfBirth gets a reference to the given Date and assigns it to the DateOfBirth field.
func (o *CorporateUpdateRequest) SetDateOfBirth(v Date) {
	o.DateOfBirth = &v
}

// GetResetMobileCounter returns the ResetMobileCounter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CorporateUpdateRequest) GetResetMobileCounter() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ResetMobileCounter
}

// GetResetMobileCounterOk returns a tuple with the ResetMobileCounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CorporateUpdateRequest) GetResetMobileCounterOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ResetMobileCounter) {
		return nil, false
	}
	return &o.ResetMobileCounter, true
}

// HasResetMobileCounter returns a boolean if a field has been set.
func (o *CorporateUpdateRequest) HasResetMobileCounter() bool {
	if o != nil && IsNil(o.ResetMobileCounter) {
		return true
	}

	return false
}

// SetResetMobileCounter gets a reference to the given interface{} and assigns it to the ResetMobileCounter field.
func (o *CorporateUpdateRequest) SetResetMobileCounter(v interface{}) {
	o.ResetMobileCounter = v
}

func (o CorporateUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CorporateUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Industry) {
		toSerialize["industry"] = o.Industry
	}
	if !IsNil(o.SourceOfFunds) {
		toSerialize["sourceOfFunds"] = o.SourceOfFunds
	}
	if o.SourceOfFundsOther != nil {
		toSerialize["sourceOfFundsOther"] = o.SourceOfFundsOther
	}
	if !IsNil(o.CompanyBusinessAddress) {
		toSerialize["companyBusinessAddress"] = o.CompanyBusinessAddress
	}
	if o.FeeGroup != nil {
		toSerialize["feeGroup"] = o.FeeGroup
	}
	if o.BaseCurrency != nil {
		toSerialize["baseCurrency"] = o.BaseCurrency
	}
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
	if o.ResetMobileCounter != nil {
		toSerialize["resetMobileCounter"] = o.ResetMobileCounter
	}
	return toSerialize, nil
}

type NullableCorporateUpdateRequest struct {
	value *CorporateUpdateRequest
	isSet bool
}

func (v NullableCorporateUpdateRequest) Get() *CorporateUpdateRequest {
	return v.value
}

func (v *NullableCorporateUpdateRequest) Set(val *CorporateUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCorporateUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCorporateUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCorporateUpdateRequest(val *CorporateUpdateRequest) *NullableCorporateUpdateRequest {
	return &NullableCorporateUpdateRequest{value: val, isSet: true}
}

func (v NullableCorporateUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCorporateUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


