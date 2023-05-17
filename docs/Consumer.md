# Consumer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**IdentityId**](IdentityId.md) | The unique identifier of the Consumer Identity. | 
**ProfileId** | **interface{}** | The profile Id which a specific identity, instrument or transaction type is linked to.  Profiles contain configuration and determine behavioral aspects of the newly created transaction, for example, fees that may apply.  You can have one or more profiles linked to your application, and these can be used to drive different behaviors according to your product&#39;s needs.  Profile Ids can be found in the Multi Portal, in the API Credentials page.  | 
**Tag** | Pointer to **interface{}** | The tag field is a custom field that can be used to search and filter. | [optional] 
**RootUser** | [**ConsumerRootUser**](ConsumerRootUser.md) | The root user of the Consumer Identity. | 
**CreationTimestamp** | **interface{}** | The time when this consumer was created, expressed in Epoch timestamp using millisecond precision. | 
**IpAddress** | **interface{}** | The IP address of the consumer user doing the registration. | 
**AcceptedTerms** | **interface{}** | Must be set to *true* to indicate that the consumer root user has accepted the terms and conditions. | 
**BaseCurrency** | Pointer to **interface{}** | The currency expressed in ISO-4217 code. Example: GBP, EUR, USD. | [optional] 
**FeeGroup** | Pointer to **interface{}** | The fee group which the consumer is bound to. | [optional] 
**SourceOfFunds** | Pointer to [**ConsumerSourceOfFunds**](ConsumerSourceOfFunds.md) |  | [optional] 
**SourceOfFundsOther** | Pointer to **interface{}** | Description of source of funds in case &#x60;OTHER&#x60; was chosen. | [optional] 

## Methods

### NewConsumer

`func NewConsumer(id IdentityId, profileId interface{}, rootUser ConsumerRootUser, creationTimestamp interface{}, ipAddress interface{}, acceptedTerms interface{}, ) *Consumer`

NewConsumer instantiates a new Consumer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerWithDefaults

`func NewConsumerWithDefaults() *Consumer`

NewConsumerWithDefaults instantiates a new Consumer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Consumer) GetId() IdentityId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Consumer) GetIdOk() (*IdentityId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Consumer) SetId(v IdentityId)`

SetId sets Id field to given value.


### GetProfileId

`func (o *Consumer) GetProfileId() interface{}`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *Consumer) GetProfileIdOk() (*interface{}, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *Consumer) SetProfileId(v interface{})`

SetProfileId sets ProfileId field to given value.


### SetProfileIdNil

`func (o *Consumer) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *Consumer) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
### GetTag

`func (o *Consumer) GetTag() interface{}`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Consumer) GetTagOk() (*interface{}, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Consumer) SetTag(v interface{})`

SetTag sets Tag field to given value.

### HasTag

`func (o *Consumer) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *Consumer) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *Consumer) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil
### GetRootUser

`func (o *Consumer) GetRootUser() ConsumerRootUser`

GetRootUser returns the RootUser field if non-nil, zero value otherwise.

### GetRootUserOk

`func (o *Consumer) GetRootUserOk() (*ConsumerRootUser, bool)`

GetRootUserOk returns a tuple with the RootUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootUser

`func (o *Consumer) SetRootUser(v ConsumerRootUser)`

SetRootUser sets RootUser field to given value.


### GetCreationTimestamp

`func (o *Consumer) GetCreationTimestamp() interface{}`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *Consumer) GetCreationTimestampOk() (*interface{}, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *Consumer) SetCreationTimestamp(v interface{})`

SetCreationTimestamp sets CreationTimestamp field to given value.


### SetCreationTimestampNil

`func (o *Consumer) SetCreationTimestampNil(b bool)`

 SetCreationTimestampNil sets the value for CreationTimestamp to be an explicit nil

### UnsetCreationTimestamp
`func (o *Consumer) UnsetCreationTimestamp()`

UnsetCreationTimestamp ensures that no value is present for CreationTimestamp, not even an explicit nil
### GetIpAddress

`func (o *Consumer) GetIpAddress() interface{}`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Consumer) GetIpAddressOk() (*interface{}, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Consumer) SetIpAddress(v interface{})`

SetIpAddress sets IpAddress field to given value.


### SetIpAddressNil

`func (o *Consumer) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *Consumer) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetAcceptedTerms

`func (o *Consumer) GetAcceptedTerms() interface{}`

GetAcceptedTerms returns the AcceptedTerms field if non-nil, zero value otherwise.

### GetAcceptedTermsOk

`func (o *Consumer) GetAcceptedTermsOk() (*interface{}, bool)`

GetAcceptedTermsOk returns a tuple with the AcceptedTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedTerms

`func (o *Consumer) SetAcceptedTerms(v interface{})`

SetAcceptedTerms sets AcceptedTerms field to given value.


### SetAcceptedTermsNil

`func (o *Consumer) SetAcceptedTermsNil(b bool)`

 SetAcceptedTermsNil sets the value for AcceptedTerms to be an explicit nil

### UnsetAcceptedTerms
`func (o *Consumer) UnsetAcceptedTerms()`

UnsetAcceptedTerms ensures that no value is present for AcceptedTerms, not even an explicit nil
### GetBaseCurrency

`func (o *Consumer) GetBaseCurrency() interface{}`

GetBaseCurrency returns the BaseCurrency field if non-nil, zero value otherwise.

### GetBaseCurrencyOk

`func (o *Consumer) GetBaseCurrencyOk() (*interface{}, bool)`

GetBaseCurrencyOk returns a tuple with the BaseCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrency

`func (o *Consumer) SetBaseCurrency(v interface{})`

SetBaseCurrency sets BaseCurrency field to given value.

### HasBaseCurrency

`func (o *Consumer) HasBaseCurrency() bool`

HasBaseCurrency returns a boolean if a field has been set.

### SetBaseCurrencyNil

`func (o *Consumer) SetBaseCurrencyNil(b bool)`

 SetBaseCurrencyNil sets the value for BaseCurrency to be an explicit nil

### UnsetBaseCurrency
`func (o *Consumer) UnsetBaseCurrency()`

UnsetBaseCurrency ensures that no value is present for BaseCurrency, not even an explicit nil
### GetFeeGroup

`func (o *Consumer) GetFeeGroup() interface{}`

GetFeeGroup returns the FeeGroup field if non-nil, zero value otherwise.

### GetFeeGroupOk

`func (o *Consumer) GetFeeGroupOk() (*interface{}, bool)`

GetFeeGroupOk returns a tuple with the FeeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeGroup

`func (o *Consumer) SetFeeGroup(v interface{})`

SetFeeGroup sets FeeGroup field to given value.

### HasFeeGroup

`func (o *Consumer) HasFeeGroup() bool`

HasFeeGroup returns a boolean if a field has been set.

### SetFeeGroupNil

`func (o *Consumer) SetFeeGroupNil(b bool)`

 SetFeeGroupNil sets the value for FeeGroup to be an explicit nil

### UnsetFeeGroup
`func (o *Consumer) UnsetFeeGroup()`

UnsetFeeGroup ensures that no value is present for FeeGroup, not even an explicit nil
### GetSourceOfFunds

`func (o *Consumer) GetSourceOfFunds() ConsumerSourceOfFunds`

GetSourceOfFunds returns the SourceOfFunds field if non-nil, zero value otherwise.

### GetSourceOfFundsOk

`func (o *Consumer) GetSourceOfFundsOk() (*ConsumerSourceOfFunds, bool)`

GetSourceOfFundsOk returns a tuple with the SourceOfFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfFunds

`func (o *Consumer) SetSourceOfFunds(v ConsumerSourceOfFunds)`

SetSourceOfFunds sets SourceOfFunds field to given value.

### HasSourceOfFunds

`func (o *Consumer) HasSourceOfFunds() bool`

HasSourceOfFunds returns a boolean if a field has been set.

### GetSourceOfFundsOther

`func (o *Consumer) GetSourceOfFundsOther() interface{}`

GetSourceOfFundsOther returns the SourceOfFundsOther field if non-nil, zero value otherwise.

### GetSourceOfFundsOtherOk

`func (o *Consumer) GetSourceOfFundsOtherOk() (*interface{}, bool)`

GetSourceOfFundsOtherOk returns a tuple with the SourceOfFundsOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfFundsOther

`func (o *Consumer) SetSourceOfFundsOther(v interface{})`

SetSourceOfFundsOther sets SourceOfFundsOther field to given value.

### HasSourceOfFundsOther

`func (o *Consumer) HasSourceOfFundsOther() bool`

HasSourceOfFundsOther returns a boolean if a field has been set.

### SetSourceOfFundsOtherNil

`func (o *Consumer) SetSourceOfFundsOtherNil(b bool)`

 SetSourceOfFundsOtherNil sets the value for SourceOfFundsOther to be an explicit nil

### UnsetSourceOfFundsOther
`func (o *Consumer) UnsetSourceOfFundsOther()`

UnsetSourceOfFundsOther ensures that no value is present for SourceOfFundsOther, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


