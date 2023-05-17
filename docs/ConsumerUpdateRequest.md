# ConsumerUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | Pointer to **interface{}** | The tag field is a custom field that can be used to search and filter. | [optional] 
**Name** | Pointer to **interface{}** | The first name of the Consumer root user. | [optional] 
**Surname** | Pointer to **interface{}** | The last name of the Consumer root user. | [optional] 
**Email** | Pointer to **interface{}** | E-mail Address of the user | [optional] 
**Mobile** | Pointer to [**Mobile**](Mobile.md) |  | [optional] 
**DateOfBirth** | Pointer to [**Date**](Date.md) | Date of birth of the consumer root user. | [optional] 
**Address** | Pointer to [**AddressWithNoRequiredFields**](AddressWithNoRequiredFields.md) | Address of the consumer root user. | [optional] 
**FeeGroup** | Pointer to **interface{}** | The fee group which the consumer will be bound to. Do not specify this if you are not using fee groups. | [optional] 
**BaseCurrency** | Pointer to **interface{}** | The currency expressed in ISO-4217 code. Example: GBP, EUR, USD. | [optional] 
**Occupation** | Pointer to [**Occupation**](Occupation.md) |  | [optional] 
**SourceOfFunds** | Pointer to [**ConsumerSourceOfFunds**](ConsumerSourceOfFunds.md) |  | [optional] 
**SourceOfFundsOther** | Pointer to **interface{}** | Description of source of funds in case &#x60;OTHER&#x60; was chosen. | [optional] 
**PlaceOfBirth** | Pointer to **interface{}** | The place of birth of the consumer root user. | [optional] 
**Nationality** | Pointer to **interface{}** | Nationality of the user - using ISO 3166 alpha-2. | [optional] 
**ResetMobileCounter** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewConsumerUpdateRequest

`func NewConsumerUpdateRequest() *ConsumerUpdateRequest`

NewConsumerUpdateRequest instantiates a new ConsumerUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerUpdateRequestWithDefaults

`func NewConsumerUpdateRequestWithDefaults() *ConsumerUpdateRequest`

NewConsumerUpdateRequestWithDefaults instantiates a new ConsumerUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *ConsumerUpdateRequest) GetTag() interface{}`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ConsumerUpdateRequest) GetTagOk() (*interface{}, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ConsumerUpdateRequest) SetTag(v interface{})`

SetTag sets Tag field to given value.

### HasTag

`func (o *ConsumerUpdateRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *ConsumerUpdateRequest) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *ConsumerUpdateRequest) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil
### GetName

`func (o *ConsumerUpdateRequest) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsumerUpdateRequest) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsumerUpdateRequest) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *ConsumerUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ConsumerUpdateRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConsumerUpdateRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSurname

`func (o *ConsumerUpdateRequest) GetSurname() interface{}`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *ConsumerUpdateRequest) GetSurnameOk() (*interface{}, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *ConsumerUpdateRequest) SetSurname(v interface{})`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *ConsumerUpdateRequest) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### SetSurnameNil

`func (o *ConsumerUpdateRequest) SetSurnameNil(b bool)`

 SetSurnameNil sets the value for Surname to be an explicit nil

### UnsetSurname
`func (o *ConsumerUpdateRequest) UnsetSurname()`

UnsetSurname ensures that no value is present for Surname, not even an explicit nil
### GetEmail

`func (o *ConsumerUpdateRequest) GetEmail() interface{}`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ConsumerUpdateRequest) GetEmailOk() (*interface{}, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ConsumerUpdateRequest) SetEmail(v interface{})`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ConsumerUpdateRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ConsumerUpdateRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ConsumerUpdateRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetMobile

`func (o *ConsumerUpdateRequest) GetMobile() Mobile`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *ConsumerUpdateRequest) GetMobileOk() (*Mobile, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *ConsumerUpdateRequest) SetMobile(v Mobile)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *ConsumerUpdateRequest) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *ConsumerUpdateRequest) GetDateOfBirth() Date`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *ConsumerUpdateRequest) GetDateOfBirthOk() (*Date, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *ConsumerUpdateRequest) SetDateOfBirth(v Date)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *ConsumerUpdateRequest) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetAddress

`func (o *ConsumerUpdateRequest) GetAddress() AddressWithNoRequiredFields`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ConsumerUpdateRequest) GetAddressOk() (*AddressWithNoRequiredFields, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ConsumerUpdateRequest) SetAddress(v AddressWithNoRequiredFields)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ConsumerUpdateRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetFeeGroup

`func (o *ConsumerUpdateRequest) GetFeeGroup() interface{}`

GetFeeGroup returns the FeeGroup field if non-nil, zero value otherwise.

### GetFeeGroupOk

`func (o *ConsumerUpdateRequest) GetFeeGroupOk() (*interface{}, bool)`

GetFeeGroupOk returns a tuple with the FeeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeGroup

`func (o *ConsumerUpdateRequest) SetFeeGroup(v interface{})`

SetFeeGroup sets FeeGroup field to given value.

### HasFeeGroup

`func (o *ConsumerUpdateRequest) HasFeeGroup() bool`

HasFeeGroup returns a boolean if a field has been set.

### SetFeeGroupNil

`func (o *ConsumerUpdateRequest) SetFeeGroupNil(b bool)`

 SetFeeGroupNil sets the value for FeeGroup to be an explicit nil

### UnsetFeeGroup
`func (o *ConsumerUpdateRequest) UnsetFeeGroup()`

UnsetFeeGroup ensures that no value is present for FeeGroup, not even an explicit nil
### GetBaseCurrency

`func (o *ConsumerUpdateRequest) GetBaseCurrency() interface{}`

GetBaseCurrency returns the BaseCurrency field if non-nil, zero value otherwise.

### GetBaseCurrencyOk

`func (o *ConsumerUpdateRequest) GetBaseCurrencyOk() (*interface{}, bool)`

GetBaseCurrencyOk returns a tuple with the BaseCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrency

`func (o *ConsumerUpdateRequest) SetBaseCurrency(v interface{})`

SetBaseCurrency sets BaseCurrency field to given value.

### HasBaseCurrency

`func (o *ConsumerUpdateRequest) HasBaseCurrency() bool`

HasBaseCurrency returns a boolean if a field has been set.

### SetBaseCurrencyNil

`func (o *ConsumerUpdateRequest) SetBaseCurrencyNil(b bool)`

 SetBaseCurrencyNil sets the value for BaseCurrency to be an explicit nil

### UnsetBaseCurrency
`func (o *ConsumerUpdateRequest) UnsetBaseCurrency()`

UnsetBaseCurrency ensures that no value is present for BaseCurrency, not even an explicit nil
### GetOccupation

`func (o *ConsumerUpdateRequest) GetOccupation() Occupation`

GetOccupation returns the Occupation field if non-nil, zero value otherwise.

### GetOccupationOk

`func (o *ConsumerUpdateRequest) GetOccupationOk() (*Occupation, bool)`

GetOccupationOk returns a tuple with the Occupation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupation

`func (o *ConsumerUpdateRequest) SetOccupation(v Occupation)`

SetOccupation sets Occupation field to given value.

### HasOccupation

`func (o *ConsumerUpdateRequest) HasOccupation() bool`

HasOccupation returns a boolean if a field has been set.

### GetSourceOfFunds

`func (o *ConsumerUpdateRequest) GetSourceOfFunds() ConsumerSourceOfFunds`

GetSourceOfFunds returns the SourceOfFunds field if non-nil, zero value otherwise.

### GetSourceOfFundsOk

`func (o *ConsumerUpdateRequest) GetSourceOfFundsOk() (*ConsumerSourceOfFunds, bool)`

GetSourceOfFundsOk returns a tuple with the SourceOfFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfFunds

`func (o *ConsumerUpdateRequest) SetSourceOfFunds(v ConsumerSourceOfFunds)`

SetSourceOfFunds sets SourceOfFunds field to given value.

### HasSourceOfFunds

`func (o *ConsumerUpdateRequest) HasSourceOfFunds() bool`

HasSourceOfFunds returns a boolean if a field has been set.

### GetSourceOfFundsOther

`func (o *ConsumerUpdateRequest) GetSourceOfFundsOther() interface{}`

GetSourceOfFundsOther returns the SourceOfFundsOther field if non-nil, zero value otherwise.

### GetSourceOfFundsOtherOk

`func (o *ConsumerUpdateRequest) GetSourceOfFundsOtherOk() (*interface{}, bool)`

GetSourceOfFundsOtherOk returns a tuple with the SourceOfFundsOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfFundsOther

`func (o *ConsumerUpdateRequest) SetSourceOfFundsOther(v interface{})`

SetSourceOfFundsOther sets SourceOfFundsOther field to given value.

### HasSourceOfFundsOther

`func (o *ConsumerUpdateRequest) HasSourceOfFundsOther() bool`

HasSourceOfFundsOther returns a boolean if a field has been set.

### SetSourceOfFundsOtherNil

`func (o *ConsumerUpdateRequest) SetSourceOfFundsOtherNil(b bool)`

 SetSourceOfFundsOtherNil sets the value for SourceOfFundsOther to be an explicit nil

### UnsetSourceOfFundsOther
`func (o *ConsumerUpdateRequest) UnsetSourceOfFundsOther()`

UnsetSourceOfFundsOther ensures that no value is present for SourceOfFundsOther, not even an explicit nil
### GetPlaceOfBirth

`func (o *ConsumerUpdateRequest) GetPlaceOfBirth() interface{}`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *ConsumerUpdateRequest) GetPlaceOfBirthOk() (*interface{}, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *ConsumerUpdateRequest) SetPlaceOfBirth(v interface{})`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *ConsumerUpdateRequest) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### SetPlaceOfBirthNil

`func (o *ConsumerUpdateRequest) SetPlaceOfBirthNil(b bool)`

 SetPlaceOfBirthNil sets the value for PlaceOfBirth to be an explicit nil

### UnsetPlaceOfBirth
`func (o *ConsumerUpdateRequest) UnsetPlaceOfBirth()`

UnsetPlaceOfBirth ensures that no value is present for PlaceOfBirth, not even an explicit nil
### GetNationality

`func (o *ConsumerUpdateRequest) GetNationality() interface{}`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *ConsumerUpdateRequest) GetNationalityOk() (*interface{}, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *ConsumerUpdateRequest) SetNationality(v interface{})`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *ConsumerUpdateRequest) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### SetNationalityNil

`func (o *ConsumerUpdateRequest) SetNationalityNil(b bool)`

 SetNationalityNil sets the value for Nationality to be an explicit nil

### UnsetNationality
`func (o *ConsumerUpdateRequest) UnsetNationality()`

UnsetNationality ensures that no value is present for Nationality, not even an explicit nil
### GetResetMobileCounter

`func (o *ConsumerUpdateRequest) GetResetMobileCounter() interface{}`

GetResetMobileCounter returns the ResetMobileCounter field if non-nil, zero value otherwise.

### GetResetMobileCounterOk

`func (o *ConsumerUpdateRequest) GetResetMobileCounterOk() (*interface{}, bool)`

GetResetMobileCounterOk returns a tuple with the ResetMobileCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetMobileCounter

`func (o *ConsumerUpdateRequest) SetResetMobileCounter(v interface{})`

SetResetMobileCounter sets ResetMobileCounter field to given value.

### HasResetMobileCounter

`func (o *ConsumerUpdateRequest) HasResetMobileCounter() bool`

HasResetMobileCounter returns a boolean if a field has been set.

### SetResetMobileCounterNil

`func (o *ConsumerUpdateRequest) SetResetMobileCounterNil(b bool)`

 SetResetMobileCounterNil sets the value for ResetMobileCounter to be an explicit nil

### UnsetResetMobileCounter
`func (o *ConsumerUpdateRequest) UnsetResetMobileCounter()`

UnsetResetMobileCounter ensures that no value is present for ResetMobileCounter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


