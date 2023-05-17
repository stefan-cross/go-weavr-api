# Corporate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**IdentityId**](IdentityId.md) | The unique identifier of the Corporate Identity. | 
**ProfileId** | **interface{}** | The profile Id which a specific identity, instrument or transaction type is linked to.  Profiles contain configuration and determine behavioral aspects of the newly created transaction, for example, fees that may apply.  You can have one or more profiles linked to your application, and these can be used to drive different behaviors according to your product&#39;s needs.  Profile Ids can be found in the Multi Portal, in the API Credentials page.  | 
**Tag** | Pointer to **interface{}** | The tag field is a custom field that can be used to search and filter. | [optional] 
**RootUser** | [**CorporateRootUser**](CorporateRootUser.md) | The root user of the Corporate Identity. | 
**Company** | [**CorporateCompany**](CorporateCompany.md) |  | 
**Industry** | Pointer to [**Industry**](Industry.md) |  | [optional] 
**SourceOfFunds** | Pointer to [**CorporateSourceOfFunds**](CorporateSourceOfFunds.md) |  | [optional] 
**SourceOfFundsOther** | Pointer to **interface{}** | Description of source of funds in case &#x60;OTHER&#x60; was chosen. | [optional] 
**AcceptedTerms** | **interface{}** | Must be set to *true* to indicate that the root user has accepted the terms and conditions. | 
**IpAddress** | **interface{}** | The IP address of the user doing the registration. | 
**BaseCurrency** | **interface{}** | The currency expressed in ISO-4217 code. Example: GBP, EUR, USD. | 
**FeeGroup** | Pointer to **interface{}** | Fee groups allow the possibility to charge different fees to users under the same profile. If fee groups are not required, ignore this field. | [optional] 
**CreationTimestamp** | **interface{}** | The time when the Corporate was created, expressed in Epoch timestamp using millisecond precision. | 

## Methods

### NewCorporate

`func NewCorporate(id IdentityId, profileId interface{}, rootUser CorporateRootUser, company CorporateCompany, acceptedTerms interface{}, ipAddress interface{}, baseCurrency interface{}, creationTimestamp interface{}, ) *Corporate`

NewCorporate instantiates a new Corporate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporateWithDefaults

`func NewCorporateWithDefaults() *Corporate`

NewCorporateWithDefaults instantiates a new Corporate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Corporate) GetId() IdentityId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Corporate) GetIdOk() (*IdentityId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Corporate) SetId(v IdentityId)`

SetId sets Id field to given value.


### GetProfileId

`func (o *Corporate) GetProfileId() interface{}`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *Corporate) GetProfileIdOk() (*interface{}, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *Corporate) SetProfileId(v interface{})`

SetProfileId sets ProfileId field to given value.


### SetProfileIdNil

`func (o *Corporate) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *Corporate) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
### GetTag

`func (o *Corporate) GetTag() interface{}`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Corporate) GetTagOk() (*interface{}, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Corporate) SetTag(v interface{})`

SetTag sets Tag field to given value.

### HasTag

`func (o *Corporate) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *Corporate) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *Corporate) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil
### GetRootUser

`func (o *Corporate) GetRootUser() CorporateRootUser`

GetRootUser returns the RootUser field if non-nil, zero value otherwise.

### GetRootUserOk

`func (o *Corporate) GetRootUserOk() (*CorporateRootUser, bool)`

GetRootUserOk returns a tuple with the RootUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootUser

`func (o *Corporate) SetRootUser(v CorporateRootUser)`

SetRootUser sets RootUser field to given value.


### GetCompany

`func (o *Corporate) GetCompany() CorporateCompany`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Corporate) GetCompanyOk() (*CorporateCompany, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Corporate) SetCompany(v CorporateCompany)`

SetCompany sets Company field to given value.


### GetIndustry

`func (o *Corporate) GetIndustry() Industry`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *Corporate) GetIndustryOk() (*Industry, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *Corporate) SetIndustry(v Industry)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *Corporate) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetSourceOfFunds

`func (o *Corporate) GetSourceOfFunds() CorporateSourceOfFunds`

GetSourceOfFunds returns the SourceOfFunds field if non-nil, zero value otherwise.

### GetSourceOfFundsOk

`func (o *Corporate) GetSourceOfFundsOk() (*CorporateSourceOfFunds, bool)`

GetSourceOfFundsOk returns a tuple with the SourceOfFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfFunds

`func (o *Corporate) SetSourceOfFunds(v CorporateSourceOfFunds)`

SetSourceOfFunds sets SourceOfFunds field to given value.

### HasSourceOfFunds

`func (o *Corporate) HasSourceOfFunds() bool`

HasSourceOfFunds returns a boolean if a field has been set.

### GetSourceOfFundsOther

`func (o *Corporate) GetSourceOfFundsOther() interface{}`

GetSourceOfFundsOther returns the SourceOfFundsOther field if non-nil, zero value otherwise.

### GetSourceOfFundsOtherOk

`func (o *Corporate) GetSourceOfFundsOtherOk() (*interface{}, bool)`

GetSourceOfFundsOtherOk returns a tuple with the SourceOfFundsOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfFundsOther

`func (o *Corporate) SetSourceOfFundsOther(v interface{})`

SetSourceOfFundsOther sets SourceOfFundsOther field to given value.

### HasSourceOfFundsOther

`func (o *Corporate) HasSourceOfFundsOther() bool`

HasSourceOfFundsOther returns a boolean if a field has been set.

### SetSourceOfFundsOtherNil

`func (o *Corporate) SetSourceOfFundsOtherNil(b bool)`

 SetSourceOfFundsOtherNil sets the value for SourceOfFundsOther to be an explicit nil

### UnsetSourceOfFundsOther
`func (o *Corporate) UnsetSourceOfFundsOther()`

UnsetSourceOfFundsOther ensures that no value is present for SourceOfFundsOther, not even an explicit nil
### GetAcceptedTerms

`func (o *Corporate) GetAcceptedTerms() interface{}`

GetAcceptedTerms returns the AcceptedTerms field if non-nil, zero value otherwise.

### GetAcceptedTermsOk

`func (o *Corporate) GetAcceptedTermsOk() (*interface{}, bool)`

GetAcceptedTermsOk returns a tuple with the AcceptedTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedTerms

`func (o *Corporate) SetAcceptedTerms(v interface{})`

SetAcceptedTerms sets AcceptedTerms field to given value.


### SetAcceptedTermsNil

`func (o *Corporate) SetAcceptedTermsNil(b bool)`

 SetAcceptedTermsNil sets the value for AcceptedTerms to be an explicit nil

### UnsetAcceptedTerms
`func (o *Corporate) UnsetAcceptedTerms()`

UnsetAcceptedTerms ensures that no value is present for AcceptedTerms, not even an explicit nil
### GetIpAddress

`func (o *Corporate) GetIpAddress() interface{}`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Corporate) GetIpAddressOk() (*interface{}, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Corporate) SetIpAddress(v interface{})`

SetIpAddress sets IpAddress field to given value.


### SetIpAddressNil

`func (o *Corporate) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *Corporate) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetBaseCurrency

`func (o *Corporate) GetBaseCurrency() interface{}`

GetBaseCurrency returns the BaseCurrency field if non-nil, zero value otherwise.

### GetBaseCurrencyOk

`func (o *Corporate) GetBaseCurrencyOk() (*interface{}, bool)`

GetBaseCurrencyOk returns a tuple with the BaseCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrency

`func (o *Corporate) SetBaseCurrency(v interface{})`

SetBaseCurrency sets BaseCurrency field to given value.


### SetBaseCurrencyNil

`func (o *Corporate) SetBaseCurrencyNil(b bool)`

 SetBaseCurrencyNil sets the value for BaseCurrency to be an explicit nil

### UnsetBaseCurrency
`func (o *Corporate) UnsetBaseCurrency()`

UnsetBaseCurrency ensures that no value is present for BaseCurrency, not even an explicit nil
### GetFeeGroup

`func (o *Corporate) GetFeeGroup() interface{}`

GetFeeGroup returns the FeeGroup field if non-nil, zero value otherwise.

### GetFeeGroupOk

`func (o *Corporate) GetFeeGroupOk() (*interface{}, bool)`

GetFeeGroupOk returns a tuple with the FeeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeGroup

`func (o *Corporate) SetFeeGroup(v interface{})`

SetFeeGroup sets FeeGroup field to given value.

### HasFeeGroup

`func (o *Corporate) HasFeeGroup() bool`

HasFeeGroup returns a boolean if a field has been set.

### SetFeeGroupNil

`func (o *Corporate) SetFeeGroupNil(b bool)`

 SetFeeGroupNil sets the value for FeeGroup to be an explicit nil

### UnsetFeeGroup
`func (o *Corporate) UnsetFeeGroup()`

UnsetFeeGroup ensures that no value is present for FeeGroup, not even an explicit nil
### GetCreationTimestamp

`func (o *Corporate) GetCreationTimestamp() interface{}`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *Corporate) GetCreationTimestampOk() (*interface{}, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *Corporate) SetCreationTimestamp(v interface{})`

SetCreationTimestamp sets CreationTimestamp field to given value.


### SetCreationTimestampNil

`func (o *Corporate) SetCreationTimestampNil(b bool)`

 SetCreationTimestampNil sets the value for CreationTimestamp to be an explicit nil

### UnsetCreationTimestamp
`func (o *Corporate) UnsetCreationTimestamp()`

UnsetCreationTimestamp ensures that no value is present for CreationTimestamp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


