# CorporateUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | Pointer to **interface{}** | The tag field is a custom field that can be used to search and filter. | [optional] 
**Industry** | Pointer to [**Industry**](Industry.md) |  | [optional] 
**SourceOfFunds** | Pointer to [**CorporateSourceOfFunds**](CorporateSourceOfFunds.md) |  | [optional] 
**SourceOfFundsOther** | Pointer to **interface{}** | Description of source of funds in case &#x60;OTHER&#x60; was chosen. | [optional] 
**CompanyBusinessAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**FeeGroup** | Pointer to **interface{}** | The fee group which the Corporate will be bound to. Do not specify this if you are not using fee groups. | [optional] 
**BaseCurrency** | Pointer to **interface{}** | The currency expressed in ISO-4217 code. Example: GBP, EUR, USD. | [optional] 
**Name** | Pointer to **interface{}** | The first name of the Corporate root user. | [optional] 
**Surname** | Pointer to **interface{}** | The last name of the Corporate root user. | [optional] 
**Email** | Pointer to **interface{}** | E-mail Address of the user | [optional] 
**Mobile** | Pointer to [**Mobile**](Mobile.md) |  | [optional] 
**DateOfBirth** | Pointer to [**Date**](Date.md) | Date of birth of the authorised user. | [optional] 
**ResetMobileCounter** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewCorporateUpdateRequest

`func NewCorporateUpdateRequest() *CorporateUpdateRequest`

NewCorporateUpdateRequest instantiates a new CorporateUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporateUpdateRequestWithDefaults

`func NewCorporateUpdateRequestWithDefaults() *CorporateUpdateRequest`

NewCorporateUpdateRequestWithDefaults instantiates a new CorporateUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *CorporateUpdateRequest) GetTag() interface{}`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CorporateUpdateRequest) GetTagOk() (*interface{}, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CorporateUpdateRequest) SetTag(v interface{})`

SetTag sets Tag field to given value.

### HasTag

`func (o *CorporateUpdateRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *CorporateUpdateRequest) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *CorporateUpdateRequest) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil
### GetIndustry

`func (o *CorporateUpdateRequest) GetIndustry() Industry`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *CorporateUpdateRequest) GetIndustryOk() (*Industry, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *CorporateUpdateRequest) SetIndustry(v Industry)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *CorporateUpdateRequest) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetSourceOfFunds

`func (o *CorporateUpdateRequest) GetSourceOfFunds() CorporateSourceOfFunds`

GetSourceOfFunds returns the SourceOfFunds field if non-nil, zero value otherwise.

### GetSourceOfFundsOk

`func (o *CorporateUpdateRequest) GetSourceOfFundsOk() (*CorporateSourceOfFunds, bool)`

GetSourceOfFundsOk returns a tuple with the SourceOfFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfFunds

`func (o *CorporateUpdateRequest) SetSourceOfFunds(v CorporateSourceOfFunds)`

SetSourceOfFunds sets SourceOfFunds field to given value.

### HasSourceOfFunds

`func (o *CorporateUpdateRequest) HasSourceOfFunds() bool`

HasSourceOfFunds returns a boolean if a field has been set.

### GetSourceOfFundsOther

`func (o *CorporateUpdateRequest) GetSourceOfFundsOther() interface{}`

GetSourceOfFundsOther returns the SourceOfFundsOther field if non-nil, zero value otherwise.

### GetSourceOfFundsOtherOk

`func (o *CorporateUpdateRequest) GetSourceOfFundsOtherOk() (*interface{}, bool)`

GetSourceOfFundsOtherOk returns a tuple with the SourceOfFundsOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfFundsOther

`func (o *CorporateUpdateRequest) SetSourceOfFundsOther(v interface{})`

SetSourceOfFundsOther sets SourceOfFundsOther field to given value.

### HasSourceOfFundsOther

`func (o *CorporateUpdateRequest) HasSourceOfFundsOther() bool`

HasSourceOfFundsOther returns a boolean if a field has been set.

### SetSourceOfFundsOtherNil

`func (o *CorporateUpdateRequest) SetSourceOfFundsOtherNil(b bool)`

 SetSourceOfFundsOtherNil sets the value for SourceOfFundsOther to be an explicit nil

### UnsetSourceOfFundsOther
`func (o *CorporateUpdateRequest) UnsetSourceOfFundsOther()`

UnsetSourceOfFundsOther ensures that no value is present for SourceOfFundsOther, not even an explicit nil
### GetCompanyBusinessAddress

`func (o *CorporateUpdateRequest) GetCompanyBusinessAddress() Address`

GetCompanyBusinessAddress returns the CompanyBusinessAddress field if non-nil, zero value otherwise.

### GetCompanyBusinessAddressOk

`func (o *CorporateUpdateRequest) GetCompanyBusinessAddressOk() (*Address, bool)`

GetCompanyBusinessAddressOk returns a tuple with the CompanyBusinessAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyBusinessAddress

`func (o *CorporateUpdateRequest) SetCompanyBusinessAddress(v Address)`

SetCompanyBusinessAddress sets CompanyBusinessAddress field to given value.

### HasCompanyBusinessAddress

`func (o *CorporateUpdateRequest) HasCompanyBusinessAddress() bool`

HasCompanyBusinessAddress returns a boolean if a field has been set.

### GetFeeGroup

`func (o *CorporateUpdateRequest) GetFeeGroup() interface{}`

GetFeeGroup returns the FeeGroup field if non-nil, zero value otherwise.

### GetFeeGroupOk

`func (o *CorporateUpdateRequest) GetFeeGroupOk() (*interface{}, bool)`

GetFeeGroupOk returns a tuple with the FeeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeGroup

`func (o *CorporateUpdateRequest) SetFeeGroup(v interface{})`

SetFeeGroup sets FeeGroup field to given value.

### HasFeeGroup

`func (o *CorporateUpdateRequest) HasFeeGroup() bool`

HasFeeGroup returns a boolean if a field has been set.

### SetFeeGroupNil

`func (o *CorporateUpdateRequest) SetFeeGroupNil(b bool)`

 SetFeeGroupNil sets the value for FeeGroup to be an explicit nil

### UnsetFeeGroup
`func (o *CorporateUpdateRequest) UnsetFeeGroup()`

UnsetFeeGroup ensures that no value is present for FeeGroup, not even an explicit nil
### GetBaseCurrency

`func (o *CorporateUpdateRequest) GetBaseCurrency() interface{}`

GetBaseCurrency returns the BaseCurrency field if non-nil, zero value otherwise.

### GetBaseCurrencyOk

`func (o *CorporateUpdateRequest) GetBaseCurrencyOk() (*interface{}, bool)`

GetBaseCurrencyOk returns a tuple with the BaseCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrency

`func (o *CorporateUpdateRequest) SetBaseCurrency(v interface{})`

SetBaseCurrency sets BaseCurrency field to given value.

### HasBaseCurrency

`func (o *CorporateUpdateRequest) HasBaseCurrency() bool`

HasBaseCurrency returns a boolean if a field has been set.

### SetBaseCurrencyNil

`func (o *CorporateUpdateRequest) SetBaseCurrencyNil(b bool)`

 SetBaseCurrencyNil sets the value for BaseCurrency to be an explicit nil

### UnsetBaseCurrency
`func (o *CorporateUpdateRequest) UnsetBaseCurrency()`

UnsetBaseCurrency ensures that no value is present for BaseCurrency, not even an explicit nil
### GetName

`func (o *CorporateUpdateRequest) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CorporateUpdateRequest) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CorporateUpdateRequest) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *CorporateUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CorporateUpdateRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CorporateUpdateRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSurname

`func (o *CorporateUpdateRequest) GetSurname() interface{}`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *CorporateUpdateRequest) GetSurnameOk() (*interface{}, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *CorporateUpdateRequest) SetSurname(v interface{})`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *CorporateUpdateRequest) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### SetSurnameNil

`func (o *CorporateUpdateRequest) SetSurnameNil(b bool)`

 SetSurnameNil sets the value for Surname to be an explicit nil

### UnsetSurname
`func (o *CorporateUpdateRequest) UnsetSurname()`

UnsetSurname ensures that no value is present for Surname, not even an explicit nil
### GetEmail

`func (o *CorporateUpdateRequest) GetEmail() interface{}`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CorporateUpdateRequest) GetEmailOk() (*interface{}, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CorporateUpdateRequest) SetEmail(v interface{})`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CorporateUpdateRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CorporateUpdateRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CorporateUpdateRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetMobile

`func (o *CorporateUpdateRequest) GetMobile() Mobile`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *CorporateUpdateRequest) GetMobileOk() (*Mobile, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *CorporateUpdateRequest) SetMobile(v Mobile)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *CorporateUpdateRequest) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *CorporateUpdateRequest) GetDateOfBirth() Date`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *CorporateUpdateRequest) GetDateOfBirthOk() (*Date, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *CorporateUpdateRequest) SetDateOfBirth(v Date)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *CorporateUpdateRequest) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetResetMobileCounter

`func (o *CorporateUpdateRequest) GetResetMobileCounter() interface{}`

GetResetMobileCounter returns the ResetMobileCounter field if non-nil, zero value otherwise.

### GetResetMobileCounterOk

`func (o *CorporateUpdateRequest) GetResetMobileCounterOk() (*interface{}, bool)`

GetResetMobileCounterOk returns a tuple with the ResetMobileCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetMobileCounter

`func (o *CorporateUpdateRequest) SetResetMobileCounter(v interface{})`

SetResetMobileCounter sets ResetMobileCounter field to given value.

### HasResetMobileCounter

`func (o *CorporateUpdateRequest) HasResetMobileCounter() bool`

HasResetMobileCounter returns a boolean if a field has been set.

### SetResetMobileCounterNil

`func (o *CorporateUpdateRequest) SetResetMobileCounterNil(b bool)`

 SetResetMobileCounterNil sets the value for ResetMobileCounter to be an explicit nil

### UnsetResetMobileCounter
`func (o *CorporateUpdateRequest) UnsetResetMobileCounter()`

UnsetResetMobileCounter ensures that no value is present for ResetMobileCounter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


