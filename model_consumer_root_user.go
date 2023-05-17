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

// checks if the ConsumerRootUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConsumerRootUser{}

// ConsumerRootUser struct for ConsumerRootUser
type ConsumerRootUser struct {
	Id IdentityId `json:"id"`
	// First name of the root user.
	Name interface{} `json:"name"`
	// Last name of the root user.
	Surname interface{} `json:"surname"`
	// E-mail Address of the user
	Email interface{} `json:"email"`
	Mobile Mobile `json:"mobile"`
	// Deprecated
	Occupation *Occupation `json:"occupation,omitempty"`
	// The state of the root user. If false, then the user will not be able to log in.
	Active interface{} `json:"active"`
	// Indicates if the root user's email has been verified.
	EmailVerified interface{} `json:"emailVerified"`
	// Indicates if the root user's mobile number has been verified.
	MobileNumberVerified interface{} `json:"mobileNumberVerified"`
	// Date of birth of the consumer root user.
	DateOfBirth *Date `json:"dateOfBirth,omitempty"`
	// Address of the consumer root user.
	Address *AddressWithNoRequiredFields `json:"address,omitempty"`
	// Nationality of the user - using ISO 3166 alpha-2.
	Nationality interface{} `json:"nationality,omitempty"`
	// The place of birth of the consumer root user.
	PlaceOfBirth interface{} `json:"placeOfBirth,omitempty"`
}

// NewConsumerRootUser instantiates a new ConsumerRootUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumerRootUser(id IdentityId, name interface{}, surname interface{}, email interface{}, mobile Mobile, active interface{}, emailVerified interface{}, mobileNumberVerified interface{}) *ConsumerRootUser {
	this := ConsumerRootUser{}
	this.Id = id
	this.Name = name
	this.Surname = surname
	this.Email = email
	this.Mobile = mobile
	this.Active = active
	this.EmailVerified = emailVerified
	this.MobileNumberVerified = mobileNumberVerified
	return &this
}

// NewConsumerRootUserWithDefaults instantiates a new ConsumerRootUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumerRootUserWithDefaults() *ConsumerRootUser {
	this := ConsumerRootUser{}
	return &this
}

// GetId returns the Id field value
func (o *ConsumerRootUser) GetId() IdentityId {
	if o == nil {
		var ret IdentityId
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConsumerRootUser) GetIdOk() (*IdentityId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConsumerRootUser) SetId(v IdentityId) {
	o.Id = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ConsumerRootUser) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsumerRootUser) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConsumerRootUser) SetName(v interface{}) {
	o.Name = v
}

// GetSurname returns the Surname field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ConsumerRootUser) GetSurname() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Surname
}

// GetSurnameOk returns a tuple with the Surname field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsumerRootUser) GetSurnameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Surname) {
		return nil, false
	}
	return &o.Surname, true
}

// SetSurname sets field value
func (o *ConsumerRootUser) SetSurname(v interface{}) {
	o.Surname = v
}

// GetEmail returns the Email field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ConsumerRootUser) GetEmail() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsumerRootUser) GetEmailOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ConsumerRootUser) SetEmail(v interface{}) {
	o.Email = v
}

// GetMobile returns the Mobile field value
func (o *ConsumerRootUser) GetMobile() Mobile {
	if o == nil {
		var ret Mobile
		return ret
	}

	return o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value
// and a boolean to check if the value has been set.
func (o *ConsumerRootUser) GetMobileOk() (*Mobile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mobile, true
}

// SetMobile sets field value
func (o *ConsumerRootUser) SetMobile(v Mobile) {
	o.Mobile = v
}

// GetOccupation returns the Occupation field value if set, zero value otherwise.
// Deprecated
func (o *ConsumerRootUser) GetOccupation() Occupation {
	if o == nil || IsNil(o.Occupation) {
		var ret Occupation
		return ret
	}
	return *o.Occupation
}

// GetOccupationOk returns a tuple with the Occupation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ConsumerRootUser) GetOccupationOk() (*Occupation, bool) {
	if o == nil || IsNil(o.Occupation) {
		return nil, false
	}
	return o.Occupation, true
}

// HasOccupation returns a boolean if a field has been set.
func (o *ConsumerRootUser) HasOccupation() bool {
	if o != nil && !IsNil(o.Occupation) {
		return true
	}

	return false
}

// SetOccupation gets a reference to the given Occupation and assigns it to the Occupation field.
// Deprecated
func (o *ConsumerRootUser) SetOccupation(v Occupation) {
	o.Occupation = &v
}

// GetActive returns the Active field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ConsumerRootUser) GetActive() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsumerRootUser) GetActiveOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *ConsumerRootUser) SetActive(v interface{}) {
	o.Active = v
}

// GetEmailVerified returns the EmailVerified field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ConsumerRootUser) GetEmailVerified() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.EmailVerified
}

// GetEmailVerifiedOk returns a tuple with the EmailVerified field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsumerRootUser) GetEmailVerifiedOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EmailVerified) {
		return nil, false
	}
	return &o.EmailVerified, true
}

// SetEmailVerified sets field value
func (o *ConsumerRootUser) SetEmailVerified(v interface{}) {
	o.EmailVerified = v
}

// GetMobileNumberVerified returns the MobileNumberVerified field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ConsumerRootUser) GetMobileNumberVerified() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.MobileNumberVerified
}

// GetMobileNumberVerifiedOk returns a tuple with the MobileNumberVerified field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsumerRootUser) GetMobileNumberVerifiedOk() (*interface{}, bool) {
	if o == nil || IsNil(o.MobileNumberVerified) {
		return nil, false
	}
	return &o.MobileNumberVerified, true
}

// SetMobileNumberVerified sets field value
func (o *ConsumerRootUser) SetMobileNumberVerified(v interface{}) {
	o.MobileNumberVerified = v
}

// GetDateOfBirth returns the DateOfBirth field value if set, zero value otherwise.
func (o *ConsumerRootUser) GetDateOfBirth() Date {
	if o == nil || IsNil(o.DateOfBirth) {
		var ret Date
		return ret
	}
	return *o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerRootUser) GetDateOfBirthOk() (*Date, bool) {
	if o == nil || IsNil(o.DateOfBirth) {
		return nil, false
	}
	return o.DateOfBirth, true
}

// HasDateOfBirth returns a boolean if a field has been set.
func (o *ConsumerRootUser) HasDateOfBirth() bool {
	if o != nil && !IsNil(o.DateOfBirth) {
		return true
	}

	return false
}

// SetDateOfBirth gets a reference to the given Date and assigns it to the DateOfBirth field.
func (o *ConsumerRootUser) SetDateOfBirth(v Date) {
	o.DateOfBirth = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ConsumerRootUser) GetAddress() AddressWithNoRequiredFields {
	if o == nil || IsNil(o.Address) {
		var ret AddressWithNoRequiredFields
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerRootUser) GetAddressOk() (*AddressWithNoRequiredFields, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ConsumerRootUser) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given AddressWithNoRequiredFields and assigns it to the Address field.
func (o *ConsumerRootUser) SetAddress(v AddressWithNoRequiredFields) {
	o.Address = &v
}

// GetNationality returns the Nationality field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsumerRootUser) GetNationality() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Nationality
}

// GetNationalityOk returns a tuple with the Nationality field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsumerRootUser) GetNationalityOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Nationality) {
		return nil, false
	}
	return &o.Nationality, true
}

// HasNationality returns a boolean if a field has been set.
func (o *ConsumerRootUser) HasNationality() bool {
	if o != nil && IsNil(o.Nationality) {
		return true
	}

	return false
}

// SetNationality gets a reference to the given interface{} and assigns it to the Nationality field.
func (o *ConsumerRootUser) SetNationality(v interface{}) {
	o.Nationality = v
}

// GetPlaceOfBirth returns the PlaceOfBirth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsumerRootUser) GetPlaceOfBirth() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PlaceOfBirth
}

// GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsumerRootUser) GetPlaceOfBirthOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PlaceOfBirth) {
		return nil, false
	}
	return &o.PlaceOfBirth, true
}

// HasPlaceOfBirth returns a boolean if a field has been set.
func (o *ConsumerRootUser) HasPlaceOfBirth() bool {
	if o != nil && IsNil(o.PlaceOfBirth) {
		return true
	}

	return false
}

// SetPlaceOfBirth gets a reference to the given interface{} and assigns it to the PlaceOfBirth field.
func (o *ConsumerRootUser) SetPlaceOfBirth(v interface{}) {
	o.PlaceOfBirth = v
}

func (o ConsumerRootUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConsumerRootUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Surname != nil {
		toSerialize["surname"] = o.Surname
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	toSerialize["mobile"] = o.Mobile
	if !IsNil(o.Occupation) {
		toSerialize["occupation"] = o.Occupation
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.EmailVerified != nil {
		toSerialize["emailVerified"] = o.EmailVerified
	}
	if o.MobileNumberVerified != nil {
		toSerialize["mobileNumberVerified"] = o.MobileNumberVerified
	}
	if !IsNil(o.DateOfBirth) {
		toSerialize["dateOfBirth"] = o.DateOfBirth
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if o.Nationality != nil {
		toSerialize["nationality"] = o.Nationality
	}
	if o.PlaceOfBirth != nil {
		toSerialize["placeOfBirth"] = o.PlaceOfBirth
	}
	return toSerialize, nil
}

type NullableConsumerRootUser struct {
	value *ConsumerRootUser
	isSet bool
}

func (v NullableConsumerRootUser) Get() *ConsumerRootUser {
	return v.value
}

func (v *NullableConsumerRootUser) Set(val *ConsumerRootUser) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumerRootUser) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumerRootUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumerRootUser(val *ConsumerRootUser) *NullableConsumerRootUser {
	return &NullableConsumerRootUser{value: val, isSet: true}
}

func (v NullableConsumerRootUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumerRootUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


