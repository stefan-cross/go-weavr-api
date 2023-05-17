# CorporateCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **interface{}** | The profile Id which a specific identity, instrument or transaction type is linked to.  Profiles contain configuration and determine behavioral aspects of the newly created transaction, for example, fees that may apply.  You can have one or more profiles linked to your application, and these can be used to drive different behaviors according to your product&#39;s needs.  Profile Ids can be found in the Multi Portal, in the API Credentials page.  | 
**Tag** | Pointer to **interface{}** | The tag field is a custom field that can be used to search and filter. | [optional] 
**RootUser** | [**CorporateCreateRequestRootUser**](CorporateCreateRequestRootUser.md) |  | 
**Company** | [**CorporateCreateRequestCompany**](CorporateCreateRequestCompany.md) |  | 
**Industry** | Pointer to [**Industry**](Industry.md) |  | [optional] 
**SourceOfFunds** | Pointer to [**CorporateSourceOfFunds**](CorporateSourceOfFunds.md) |  | [optional] 
**SourceOfFundsOther** | Pointer to **interface{}** | Description of source of funds in case &#x60;OTHER&#x60; was chosen. | [optional] 
**AcceptedTerms** | **interface{}** | Must be set to *true* to indicate that the corporate root user has accepted the terms and conditions. | 
**IpAddress** | **interface{}** | The IP address of the corporate user doing the registration. | 
**BaseCurrency** | **interface{}** | The currency expressed in ISO-4217 code. Example: GBP, EUR, USD. | 
**FeeGroup** | Pointer to **interface{}** | Fee groups allow the possibility to charge different fees to users under the same profile. If fee groups are not required, ignore this field. | [optional] 

## Methods

### NewCorporateCreateRequest

`func NewCorporateCreateRequest(profileId interface{}, rootUser CorporateCreateRequestRootUser, company CorporateCreateRequestCompany, acceptedTerms interface{}, ipAddress interface{}, baseCurrency interface{}, ) *CorporateCreateRequest`

NewCorporateCreateRequest instantiates a new CorporateCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporateCreateRequestWithDefaults

`func NewCorporateCreateRequestWithDefaults() *CorporateCreateRequest`

NewCorporateCreateRequestWithDefaults instantiates a new CorporateCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *CorporateCreateRequest) GetProfileId() interface{}`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *CorporateCreateRequest) GetProfileIdOk() (*interface{}, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *CorporateCreateRequest) SetProfileId(v interface{})`

SetProfileId sets ProfileId field to given value.


### SetProfileIdNil

`func (o *CorporateCreateRequest) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *CorporateCreateRequest) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
### GetTag

`func (o *CorporateCreateRequest) GetTag() interface{}`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CorporateCreateRequest) GetTagOk() (*interface{}, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CorporateCreateRequest) SetTag(v interface{})`

SetTag sets Tag field to given value.

### HasTag

`func (o *CorporateCreateRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *CorporateCreateRequest) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *CorporateCreateRequest) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil
### GetRootUser

`func (o *CorporateCreateRequest) GetRootUser() CorporateCreateRequestRootUser`

GetRootUser returns the RootUser field if non-nil, zero value otherwise.

### GetRootUserOk

`func (o *CorporateCreateRequest) GetRootUserOk() (*CorporateCreateRequestRootUser, bool)`

GetRootUserOk returns a tuple with the RootUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootUser

`func (o *CorporateCreateRequest) SetRootUser(v CorporateCreateRequestRootUser)`

SetRootUser sets RootUser field to given value.


### GetCompany

`func (o *CorporateCreateRequest) GetCompany() CorporateCreateRequestCompany`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CorporateCreateRequest) GetCompanyOk() (*CorporateCreateRequestCompany, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CorporateCreateRequest) SetCompany(v CorporateCreateRequestCompany)`

SetCompany sets Company field to given value.


### GetIndustry

`func (o *CorporateCreateRequest) GetIndustry() Industry`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *CorporateCreateRequest) GetIndustryOk() (*Industry, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *CorporateCreateRequest) SetIndustry(v Industry)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *CorporateCreateRequest) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetSourceOfFunds

`func (o *CorporateCreateRequest) GetSourceOfFunds() CorporateSourceOfFunds`

GetSourceOfFunds returns the SourceOfFunds field if non-nil, zero value otherwise.

### GetSourceOfFundsOk

`func (o *CorporateCreateRequest) GetSourceOfFundsOk() (*CorporateSourceOfFunds, bool)`

GetSourceOfFundsOk returns a tuple with the SourceOfFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfFunds

`func (o *CorporateCreateRequest) SetSourceOfFunds(v CorporateSourceOfFunds)`

SetSourceOfFunds sets SourceOfFunds field to given value.

### HasSourceOfFunds

`func (o *CorporateCreateRequest) HasSourceOfFunds() bool`

HasSourceOfFunds returns a boolean if a field has been set.

### GetSourceOfFundsOther

`func (o *CorporateCreateRequest) GetSourceOfFundsOther() interface{}`

GetSourceOfFundsOther returns the SourceOfFundsOther field if non-nil, zero value otherwise.

### GetSourceOfFundsOtherOk

`func (o *CorporateCreateRequest) GetSourceOfFundsOtherOk() (*interface{}, bool)`

GetSourceOfFundsOtherOk returns a tuple with the SourceOfFundsOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfFundsOther

`func (o *CorporateCreateRequest) SetSourceOfFundsOther(v interface{})`

SetSourceOfFundsOther sets SourceOfFundsOther field to given value.

### HasSourceOfFundsOther

`func (o *CorporateCreateRequest) HasSourceOfFundsOther() bool`

HasSourceOfFundsOther returns a boolean if a field has been set.

### SetSourceOfFundsOtherNil

`func (o *CorporateCreateRequest) SetSourceOfFundsOtherNil(b bool)`

 SetSourceOfFundsOtherNil sets the value for SourceOfFundsOther to be an explicit nil

### UnsetSourceOfFundsOther
`func (o *CorporateCreateRequest) UnsetSourceOfFundsOther()`

UnsetSourceOfFundsOther ensures that no value is present for SourceOfFundsOther, not even an explicit nil
### GetAcceptedTerms

`func (o *CorporateCreateRequest) GetAcceptedTerms() interface{}`

GetAcceptedTerms returns the AcceptedTerms field if non-nil, zero value otherwise.

### GetAcceptedTermsOk

`func (o *CorporateCreateRequest) GetAcceptedTermsOk() (*interface{}, bool)`

GetAcceptedTermsOk returns a tuple with the AcceptedTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedTerms

`func (o *CorporateCreateRequest) SetAcceptedTerms(v interface{})`

SetAcceptedTerms sets AcceptedTerms field to given value.


### SetAcceptedTermsNil

`func (o *CorporateCreateRequest) SetAcceptedTermsNil(b bool)`

 SetAcceptedTermsNil sets the value for AcceptedTerms to be an explicit nil

### UnsetAcceptedTerms
`func (o *CorporateCreateRequest) UnsetAcceptedTerms()`

UnsetAcceptedTerms ensures that no value is present for AcceptedTerms, not even an explicit nil
### GetIpAddress

`func (o *CorporateCreateRequest) GetIpAddress() interface{}`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *CorporateCreateRequest) GetIpAddressOk() (*interface{}, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *CorporateCreateRequest) SetIpAddress(v interface{})`

SetIpAddress sets IpAddress field to given value.


### SetIpAddressNil

`func (o *CorporateCreateRequest) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *CorporateCreateRequest) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetBaseCurrency

`func (o *CorporateCreateRequest) GetBaseCurrency() interface{}`

GetBaseCurrency returns the BaseCurrency field if non-nil, zero value otherwise.

### GetBaseCurrencyOk

`func (o *CorporateCreateRequest) GetBaseCurrencyOk() (*interface{}, bool)`

GetBaseCurrencyOk returns a tuple with the BaseCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrency

`func (o *CorporateCreateRequest) SetBaseCurrency(v interface{})`

SetBaseCurrency sets BaseCurrency field to given value.


### SetBaseCurrencyNil

`func (o *CorporateCreateRequest) SetBaseCurrencyNil(b bool)`

 SetBaseCurrencyNil sets the value for BaseCurrency to be an explicit nil

### UnsetBaseCurrency
`func (o *CorporateCreateRequest) UnsetBaseCurrency()`

UnsetBaseCurrency ensures that no value is present for BaseCurrency, not even an explicit nil
### GetFeeGroup

`func (o *CorporateCreateRequest) GetFeeGroup() interface{}`

GetFeeGroup returns the FeeGroup field if non-nil, zero value otherwise.

### GetFeeGroupOk

`func (o *CorporateCreateRequest) GetFeeGroupOk() (*interface{}, bool)`

GetFeeGroupOk returns a tuple with the FeeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeGroup

`func (o *CorporateCreateRequest) SetFeeGroup(v interface{})`

SetFeeGroup sets FeeGroup field to given value.

### HasFeeGroup

`func (o *CorporateCreateRequest) HasFeeGroup() bool`

HasFeeGroup returns a boolean if a field has been set.

### SetFeeGroupNil

`func (o *CorporateCreateRequest) SetFeeGroupNil(b bool)`

 SetFeeGroupNil sets the value for FeeGroup to be an explicit nil

### UnsetFeeGroup
`func (o *CorporateCreateRequest) UnsetFeeGroup()`

UnsetFeeGroup ensures that no value is present for FeeGroup, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


