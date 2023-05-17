# ConsumerCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **interface{}** | The profile Id which a specific identity, instrument or transaction type is linked to.  Profiles contain configuration and determine behavioral aspects of the newly created transaction, for example, fees that may apply.  You can have one or more profiles linked to your application, and these can be used to drive different behaviors according to your product&#39;s needs.  Profile Ids can be found in the Multi Portal, in the API Credentials page.  | 
**Tag** | Pointer to **interface{}** | The tag field is a custom field that can be used to search and filter. | [optional] 
**RootUser** | [**ConsumerCreateRequestRootUser**](ConsumerCreateRequestRootUser.md) |  | 
**IpAddress** | **interface{}** | The IP address of the user doing the registration. | 
**AcceptedTerms** | **interface{}** | Must be set to *true* to indicate that the consumer has accepted the terms and conditions. | 
**BaseCurrency** | Pointer to **interface{}** | The currency expressed in ISO-4217 code. Example: GBP, EUR, USD. | [optional] 
**FeeGroup** | Pointer to **interface{}** | The fee group which the consumer is bound to. Fee groups provide the possibility of different fees to users under the same profile. If fee groups are not required, ignore this field. | [optional] 
**SourceOfFunds** | Pointer to [**ConsumerSourceOfFunds**](ConsumerSourceOfFunds.md) |  | [optional] 
**SourceOfFundsOther** | Pointer to **interface{}** | Description of source of funds in case &#x60;OTHER&#x60; was chosen. | [optional] 

## Methods

### NewConsumerCreateRequest

`func NewConsumerCreateRequest(profileId interface{}, rootUser ConsumerCreateRequestRootUser, ipAddress interface{}, acceptedTerms interface{}, ) *ConsumerCreateRequest`

NewConsumerCreateRequest instantiates a new ConsumerCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerCreateRequestWithDefaults

`func NewConsumerCreateRequestWithDefaults() *ConsumerCreateRequest`

NewConsumerCreateRequestWithDefaults instantiates a new ConsumerCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *ConsumerCreateRequest) GetProfileId() interface{}`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *ConsumerCreateRequest) GetProfileIdOk() (*interface{}, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *ConsumerCreateRequest) SetProfileId(v interface{})`

SetProfileId sets ProfileId field to given value.


### SetProfileIdNil

`func (o *ConsumerCreateRequest) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *ConsumerCreateRequest) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
### GetTag

`func (o *ConsumerCreateRequest) GetTag() interface{}`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ConsumerCreateRequest) GetTagOk() (*interface{}, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ConsumerCreateRequest) SetTag(v interface{})`

SetTag sets Tag field to given value.

### HasTag

`func (o *ConsumerCreateRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *ConsumerCreateRequest) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *ConsumerCreateRequest) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil
### GetRootUser

`func (o *ConsumerCreateRequest) GetRootUser() ConsumerCreateRequestRootUser`

GetRootUser returns the RootUser field if non-nil, zero value otherwise.

### GetRootUserOk

`func (o *ConsumerCreateRequest) GetRootUserOk() (*ConsumerCreateRequestRootUser, bool)`

GetRootUserOk returns a tuple with the RootUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootUser

`func (o *ConsumerCreateRequest) SetRootUser(v ConsumerCreateRequestRootUser)`

SetRootUser sets RootUser field to given value.


### GetIpAddress

`func (o *ConsumerCreateRequest) GetIpAddress() interface{}`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ConsumerCreateRequest) GetIpAddressOk() (*interface{}, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ConsumerCreateRequest) SetIpAddress(v interface{})`

SetIpAddress sets IpAddress field to given value.


### SetIpAddressNil

`func (o *ConsumerCreateRequest) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *ConsumerCreateRequest) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetAcceptedTerms

`func (o *ConsumerCreateRequest) GetAcceptedTerms() interface{}`

GetAcceptedTerms returns the AcceptedTerms field if non-nil, zero value otherwise.

### GetAcceptedTermsOk

`func (o *ConsumerCreateRequest) GetAcceptedTermsOk() (*interface{}, bool)`

GetAcceptedTermsOk returns a tuple with the AcceptedTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedTerms

`func (o *ConsumerCreateRequest) SetAcceptedTerms(v interface{})`

SetAcceptedTerms sets AcceptedTerms field to given value.


### SetAcceptedTermsNil

`func (o *ConsumerCreateRequest) SetAcceptedTermsNil(b bool)`

 SetAcceptedTermsNil sets the value for AcceptedTerms to be an explicit nil

### UnsetAcceptedTerms
`func (o *ConsumerCreateRequest) UnsetAcceptedTerms()`

UnsetAcceptedTerms ensures that no value is present for AcceptedTerms, not even an explicit nil
### GetBaseCurrency

`func (o *ConsumerCreateRequest) GetBaseCurrency() interface{}`

GetBaseCurrency returns the BaseCurrency field if non-nil, zero value otherwise.

### GetBaseCurrencyOk

`func (o *ConsumerCreateRequest) GetBaseCurrencyOk() (*interface{}, bool)`

GetBaseCurrencyOk returns a tuple with the BaseCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrency

`func (o *ConsumerCreateRequest) SetBaseCurrency(v interface{})`

SetBaseCurrency sets BaseCurrency field to given value.

### HasBaseCurrency

`func (o *ConsumerCreateRequest) HasBaseCurrency() bool`

HasBaseCurrency returns a boolean if a field has been set.

### SetBaseCurrencyNil

`func (o *ConsumerCreateRequest) SetBaseCurrencyNil(b bool)`

 SetBaseCurrencyNil sets the value for BaseCurrency to be an explicit nil

### UnsetBaseCurrency
`func (o *ConsumerCreateRequest) UnsetBaseCurrency()`

UnsetBaseCurrency ensures that no value is present for BaseCurrency, not even an explicit nil
### GetFeeGroup

`func (o *ConsumerCreateRequest) GetFeeGroup() interface{}`

GetFeeGroup returns the FeeGroup field if non-nil, zero value otherwise.

### GetFeeGroupOk

`func (o *ConsumerCreateRequest) GetFeeGroupOk() (*interface{}, bool)`

GetFeeGroupOk returns a tuple with the FeeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeGroup

`func (o *ConsumerCreateRequest) SetFeeGroup(v interface{})`

SetFeeGroup sets FeeGroup field to given value.

### HasFeeGroup

`func (o *ConsumerCreateRequest) HasFeeGroup() bool`

HasFeeGroup returns a boolean if a field has been set.

### SetFeeGroupNil

`func (o *ConsumerCreateRequest) SetFeeGroupNil(b bool)`

 SetFeeGroupNil sets the value for FeeGroup to be an explicit nil

### UnsetFeeGroup
`func (o *ConsumerCreateRequest) UnsetFeeGroup()`

UnsetFeeGroup ensures that no value is present for FeeGroup, not even an explicit nil
### GetSourceOfFunds

`func (o *ConsumerCreateRequest) GetSourceOfFunds() ConsumerSourceOfFunds`

GetSourceOfFunds returns the SourceOfFunds field if non-nil, zero value otherwise.

### GetSourceOfFundsOk

`func (o *ConsumerCreateRequest) GetSourceOfFundsOk() (*ConsumerSourceOfFunds, bool)`

GetSourceOfFundsOk returns a tuple with the SourceOfFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfFunds

`func (o *ConsumerCreateRequest) SetSourceOfFunds(v ConsumerSourceOfFunds)`

SetSourceOfFunds sets SourceOfFunds field to given value.

### HasSourceOfFunds

`func (o *ConsumerCreateRequest) HasSourceOfFunds() bool`

HasSourceOfFunds returns a boolean if a field has been set.

### GetSourceOfFundsOther

`func (o *ConsumerCreateRequest) GetSourceOfFundsOther() interface{}`

GetSourceOfFundsOther returns the SourceOfFundsOther field if non-nil, zero value otherwise.

### GetSourceOfFundsOtherOk

`func (o *ConsumerCreateRequest) GetSourceOfFundsOtherOk() (*interface{}, bool)`

GetSourceOfFundsOtherOk returns a tuple with the SourceOfFundsOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfFundsOther

`func (o *ConsumerCreateRequest) SetSourceOfFundsOther(v interface{})`

SetSourceOfFundsOther sets SourceOfFundsOther field to given value.

### HasSourceOfFundsOther

`func (o *ConsumerCreateRequest) HasSourceOfFundsOther() bool`

HasSourceOfFundsOther returns a boolean if a field has been set.

### SetSourceOfFundsOtherNil

`func (o *ConsumerCreateRequest) SetSourceOfFundsOtherNil(b bool)`

 SetSourceOfFundsOtherNil sets the value for SourceOfFundsOther to be an explicit nil

### UnsetSourceOfFundsOther
`func (o *ConsumerCreateRequest) UnsetSourceOfFundsOther()`

UnsetSourceOfFundsOther ensures that no value is present for SourceOfFundsOther, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


