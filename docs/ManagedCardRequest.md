# ManagedCardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **interface{}** | The profile Id which a specific identity, instrument or transaction type is linked to.  Profiles contain configuration and determine behavioral aspects of the newly created transaction, for example, fees that may apply.  You can have one or more profiles linked to your application, and these can be used to drive different behaviors according to your product&#39;s needs.  Profile Ids can be found in the Multi Portal, in the API Credentials page.  | 
**Tag** | Pointer to **interface{}** | The tag field is a custom field that can be used to search and filter. | [optional] 
**FriendlyName** | **interface{}** | The friendly name for the card. | 
**NameOnCard** | **interface{}** | The card holder&#39;s name for the card.  This may be verified by merchants when the card is used online. For Physical cards, this field will be printed on the card. The maximum characters allowed will depend on the design chosen and will be provided to you by Weavr when setting up your plastic cards.  | 
**NameOnCardLine2** | Pointer to **interface{}** | Line 2 of the &#39;name on card&#39; field. For Physical cards, this field will be printed on the card. The maximum characters allowed will depend on the design chosen and will be provided to you by Weavr when setting up your plastic cards. | [optional] 
**CardholderMobileNumber** | **interface{}** | The mobile number including country code of the card holder.  For transactions that require a 3DS challenge, an SMS with a code will be sent on this number, to be entered during an online purchase.  | 
**BillingAddress** | [**Address**](Address.md) | The billing address set for the card holder. This may be verified by merchants when the card is used online. | 
**DigitalWallets** | Pointer to [**DigitalWallets**](DigitalWallets.md) | The Card Tokenisation details | [optional] 
**AuthForwardingDefaultTimeoutDecision** | Pointer to **interface{}** | Default decision for auth forwarding on timeout | [optional] 
**Mode** | **interface{}** | The card can be created in prepaid mode or debit mode.  - A prepaid mode card has its own balance and can have funds transferred to or from it. - A debit mode card does not have its own balance but will be able to spend funds belonging to its parent managed account, subject to a configurable spend limit.  | 

## Methods

### NewManagedCardRequest

`func NewManagedCardRequest(profileId interface{}, friendlyName interface{}, nameOnCard interface{}, cardholderMobileNumber interface{}, billingAddress Address, mode interface{}, ) *ManagedCardRequest`

NewManagedCardRequest instantiates a new ManagedCardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedCardRequestWithDefaults

`func NewManagedCardRequestWithDefaults() *ManagedCardRequest`

NewManagedCardRequestWithDefaults instantiates a new ManagedCardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *ManagedCardRequest) GetProfileId() interface{}`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *ManagedCardRequest) GetProfileIdOk() (*interface{}, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *ManagedCardRequest) SetProfileId(v interface{})`

SetProfileId sets ProfileId field to given value.


### SetProfileIdNil

`func (o *ManagedCardRequest) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *ManagedCardRequest) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
### GetTag

`func (o *ManagedCardRequest) GetTag() interface{}`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ManagedCardRequest) GetTagOk() (*interface{}, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ManagedCardRequest) SetTag(v interface{})`

SetTag sets Tag field to given value.

### HasTag

`func (o *ManagedCardRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *ManagedCardRequest) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *ManagedCardRequest) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil
### GetFriendlyName

`func (o *ManagedCardRequest) GetFriendlyName() interface{}`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ManagedCardRequest) GetFriendlyNameOk() (*interface{}, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ManagedCardRequest) SetFriendlyName(v interface{})`

SetFriendlyName sets FriendlyName field to given value.


### SetFriendlyNameNil

`func (o *ManagedCardRequest) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ManagedCardRequest) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetNameOnCard

`func (o *ManagedCardRequest) GetNameOnCard() interface{}`

GetNameOnCard returns the NameOnCard field if non-nil, zero value otherwise.

### GetNameOnCardOk

`func (o *ManagedCardRequest) GetNameOnCardOk() (*interface{}, bool)`

GetNameOnCardOk returns a tuple with the NameOnCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnCard

`func (o *ManagedCardRequest) SetNameOnCard(v interface{})`

SetNameOnCard sets NameOnCard field to given value.


### SetNameOnCardNil

`func (o *ManagedCardRequest) SetNameOnCardNil(b bool)`

 SetNameOnCardNil sets the value for NameOnCard to be an explicit nil

### UnsetNameOnCard
`func (o *ManagedCardRequest) UnsetNameOnCard()`

UnsetNameOnCard ensures that no value is present for NameOnCard, not even an explicit nil
### GetNameOnCardLine2

`func (o *ManagedCardRequest) GetNameOnCardLine2() interface{}`

GetNameOnCardLine2 returns the NameOnCardLine2 field if non-nil, zero value otherwise.

### GetNameOnCardLine2Ok

`func (o *ManagedCardRequest) GetNameOnCardLine2Ok() (*interface{}, bool)`

GetNameOnCardLine2Ok returns a tuple with the NameOnCardLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnCardLine2

`func (o *ManagedCardRequest) SetNameOnCardLine2(v interface{})`

SetNameOnCardLine2 sets NameOnCardLine2 field to given value.

### HasNameOnCardLine2

`func (o *ManagedCardRequest) HasNameOnCardLine2() bool`

HasNameOnCardLine2 returns a boolean if a field has been set.

### SetNameOnCardLine2Nil

`func (o *ManagedCardRequest) SetNameOnCardLine2Nil(b bool)`

 SetNameOnCardLine2Nil sets the value for NameOnCardLine2 to be an explicit nil

### UnsetNameOnCardLine2
`func (o *ManagedCardRequest) UnsetNameOnCardLine2()`

UnsetNameOnCardLine2 ensures that no value is present for NameOnCardLine2, not even an explicit nil
### GetCardholderMobileNumber

`func (o *ManagedCardRequest) GetCardholderMobileNumber() interface{}`

GetCardholderMobileNumber returns the CardholderMobileNumber field if non-nil, zero value otherwise.

### GetCardholderMobileNumberOk

`func (o *ManagedCardRequest) GetCardholderMobileNumberOk() (*interface{}, bool)`

GetCardholderMobileNumberOk returns a tuple with the CardholderMobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderMobileNumber

`func (o *ManagedCardRequest) SetCardholderMobileNumber(v interface{})`

SetCardholderMobileNumber sets CardholderMobileNumber field to given value.


### SetCardholderMobileNumberNil

`func (o *ManagedCardRequest) SetCardholderMobileNumberNil(b bool)`

 SetCardholderMobileNumberNil sets the value for CardholderMobileNumber to be an explicit nil

### UnsetCardholderMobileNumber
`func (o *ManagedCardRequest) UnsetCardholderMobileNumber()`

UnsetCardholderMobileNumber ensures that no value is present for CardholderMobileNumber, not even an explicit nil
### GetBillingAddress

`func (o *ManagedCardRequest) GetBillingAddress() Address`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *ManagedCardRequest) GetBillingAddressOk() (*Address, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *ManagedCardRequest) SetBillingAddress(v Address)`

SetBillingAddress sets BillingAddress field to given value.


### GetDigitalWallets

`func (o *ManagedCardRequest) GetDigitalWallets() DigitalWallets`

GetDigitalWallets returns the DigitalWallets field if non-nil, zero value otherwise.

### GetDigitalWalletsOk

`func (o *ManagedCardRequest) GetDigitalWalletsOk() (*DigitalWallets, bool)`

GetDigitalWalletsOk returns a tuple with the DigitalWallets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalWallets

`func (o *ManagedCardRequest) SetDigitalWallets(v DigitalWallets)`

SetDigitalWallets sets DigitalWallets field to given value.

### HasDigitalWallets

`func (o *ManagedCardRequest) HasDigitalWallets() bool`

HasDigitalWallets returns a boolean if a field has been set.

### GetAuthForwardingDefaultTimeoutDecision

`func (o *ManagedCardRequest) GetAuthForwardingDefaultTimeoutDecision() interface{}`

GetAuthForwardingDefaultTimeoutDecision returns the AuthForwardingDefaultTimeoutDecision field if non-nil, zero value otherwise.

### GetAuthForwardingDefaultTimeoutDecisionOk

`func (o *ManagedCardRequest) GetAuthForwardingDefaultTimeoutDecisionOk() (*interface{}, bool)`

GetAuthForwardingDefaultTimeoutDecisionOk returns a tuple with the AuthForwardingDefaultTimeoutDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthForwardingDefaultTimeoutDecision

`func (o *ManagedCardRequest) SetAuthForwardingDefaultTimeoutDecision(v interface{})`

SetAuthForwardingDefaultTimeoutDecision sets AuthForwardingDefaultTimeoutDecision field to given value.

### HasAuthForwardingDefaultTimeoutDecision

`func (o *ManagedCardRequest) HasAuthForwardingDefaultTimeoutDecision() bool`

HasAuthForwardingDefaultTimeoutDecision returns a boolean if a field has been set.

### SetAuthForwardingDefaultTimeoutDecisionNil

`func (o *ManagedCardRequest) SetAuthForwardingDefaultTimeoutDecisionNil(b bool)`

 SetAuthForwardingDefaultTimeoutDecisionNil sets the value for AuthForwardingDefaultTimeoutDecision to be an explicit nil

### UnsetAuthForwardingDefaultTimeoutDecision
`func (o *ManagedCardRequest) UnsetAuthForwardingDefaultTimeoutDecision()`

UnsetAuthForwardingDefaultTimeoutDecision ensures that no value is present for AuthForwardingDefaultTimeoutDecision, not even an explicit nil
### GetMode

`func (o *ManagedCardRequest) GetMode() interface{}`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ManagedCardRequest) GetModeOk() (*interface{}, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ManagedCardRequest) SetMode(v interface{})`

SetMode sets Mode field to given value.


### SetModeNil

`func (o *ManagedCardRequest) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *ManagedCardRequest) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


