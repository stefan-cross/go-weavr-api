# ManagedCardUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | Pointer to **interface{}** | The tag field is a custom field that can be used to search and filter. | [optional] 
**FriendlyName** | Pointer to **interface{}** | Updates the friendly name of the card. Leave blank if no change is needed. | [optional] 
**NameOnCard** | Pointer to **interface{}** | The card holder&#39;s name for the card.  This may be verified by merchants when the card is used online. For Physical cards, this field will be printed on the card. The maximum characters allowed will depend on the design chosen and will be provided to you by Weavr when setting up your plastic cards.  | [optional] 
**NameOnCardLine2** | Pointer to **interface{}** | Line 2 of the &#39;name on card&#39; field. For Physical cards, this field will be printed on the card. The maximum characters allowed will depend on the design chosen and will be provided to you by Weavr when setting up your plastic cards. | [optional] 
**CardholderMobileNumber** | Pointer to **interface{}** | The mobile number including country code of the card holder, needed in case a 3DS challenge is required. | [optional] 
**BillingAddress** | Pointer to [**Address**](Address.md) | The billing address of the card holder. Merchants may request the billing address to be checked for online purchases. | [optional] 
**DeliveryAddress** | Pointer to [**DeliveryAddress**](DeliveryAddress.md) | The delivery address set for the card holder. This is only applicable for physical cards. | [optional] 
**DeliveryMethod** | Pointer to [**DeliveryMethod**](DeliveryMethod.md) |  | [optional] 
**DigitalWallets** | Pointer to [**DigitalWallets**](DigitalWallets.md) |  | [optional] 
**AuthForwardingDefaultTimeoutDecision** | Pointer to **interface{}** | Default decision for auth forwarding on timeout | [optional] 

## Methods

### NewManagedCardUpdateRequest

`func NewManagedCardUpdateRequest() *ManagedCardUpdateRequest`

NewManagedCardUpdateRequest instantiates a new ManagedCardUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedCardUpdateRequestWithDefaults

`func NewManagedCardUpdateRequestWithDefaults() *ManagedCardUpdateRequest`

NewManagedCardUpdateRequestWithDefaults instantiates a new ManagedCardUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *ManagedCardUpdateRequest) GetTag() interface{}`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ManagedCardUpdateRequest) GetTagOk() (*interface{}, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ManagedCardUpdateRequest) SetTag(v interface{})`

SetTag sets Tag field to given value.

### HasTag

`func (o *ManagedCardUpdateRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *ManagedCardUpdateRequest) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *ManagedCardUpdateRequest) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil
### GetFriendlyName

`func (o *ManagedCardUpdateRequest) GetFriendlyName() interface{}`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ManagedCardUpdateRequest) GetFriendlyNameOk() (*interface{}, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ManagedCardUpdateRequest) SetFriendlyName(v interface{})`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ManagedCardUpdateRequest) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *ManagedCardUpdateRequest) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ManagedCardUpdateRequest) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetNameOnCard

`func (o *ManagedCardUpdateRequest) GetNameOnCard() interface{}`

GetNameOnCard returns the NameOnCard field if non-nil, zero value otherwise.

### GetNameOnCardOk

`func (o *ManagedCardUpdateRequest) GetNameOnCardOk() (*interface{}, bool)`

GetNameOnCardOk returns a tuple with the NameOnCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnCard

`func (o *ManagedCardUpdateRequest) SetNameOnCard(v interface{})`

SetNameOnCard sets NameOnCard field to given value.

### HasNameOnCard

`func (o *ManagedCardUpdateRequest) HasNameOnCard() bool`

HasNameOnCard returns a boolean if a field has been set.

### SetNameOnCardNil

`func (o *ManagedCardUpdateRequest) SetNameOnCardNil(b bool)`

 SetNameOnCardNil sets the value for NameOnCard to be an explicit nil

### UnsetNameOnCard
`func (o *ManagedCardUpdateRequest) UnsetNameOnCard()`

UnsetNameOnCard ensures that no value is present for NameOnCard, not even an explicit nil
### GetNameOnCardLine2

`func (o *ManagedCardUpdateRequest) GetNameOnCardLine2() interface{}`

GetNameOnCardLine2 returns the NameOnCardLine2 field if non-nil, zero value otherwise.

### GetNameOnCardLine2Ok

`func (o *ManagedCardUpdateRequest) GetNameOnCardLine2Ok() (*interface{}, bool)`

GetNameOnCardLine2Ok returns a tuple with the NameOnCardLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnCardLine2

`func (o *ManagedCardUpdateRequest) SetNameOnCardLine2(v interface{})`

SetNameOnCardLine2 sets NameOnCardLine2 field to given value.

### HasNameOnCardLine2

`func (o *ManagedCardUpdateRequest) HasNameOnCardLine2() bool`

HasNameOnCardLine2 returns a boolean if a field has been set.

### SetNameOnCardLine2Nil

`func (o *ManagedCardUpdateRequest) SetNameOnCardLine2Nil(b bool)`

 SetNameOnCardLine2Nil sets the value for NameOnCardLine2 to be an explicit nil

### UnsetNameOnCardLine2
`func (o *ManagedCardUpdateRequest) UnsetNameOnCardLine2()`

UnsetNameOnCardLine2 ensures that no value is present for NameOnCardLine2, not even an explicit nil
### GetCardholderMobileNumber

`func (o *ManagedCardUpdateRequest) GetCardholderMobileNumber() interface{}`

GetCardholderMobileNumber returns the CardholderMobileNumber field if non-nil, zero value otherwise.

### GetCardholderMobileNumberOk

`func (o *ManagedCardUpdateRequest) GetCardholderMobileNumberOk() (*interface{}, bool)`

GetCardholderMobileNumberOk returns a tuple with the CardholderMobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderMobileNumber

`func (o *ManagedCardUpdateRequest) SetCardholderMobileNumber(v interface{})`

SetCardholderMobileNumber sets CardholderMobileNumber field to given value.

### HasCardholderMobileNumber

`func (o *ManagedCardUpdateRequest) HasCardholderMobileNumber() bool`

HasCardholderMobileNumber returns a boolean if a field has been set.

### SetCardholderMobileNumberNil

`func (o *ManagedCardUpdateRequest) SetCardholderMobileNumberNil(b bool)`

 SetCardholderMobileNumberNil sets the value for CardholderMobileNumber to be an explicit nil

### UnsetCardholderMobileNumber
`func (o *ManagedCardUpdateRequest) UnsetCardholderMobileNumber()`

UnsetCardholderMobileNumber ensures that no value is present for CardholderMobileNumber, not even an explicit nil
### GetBillingAddress

`func (o *ManagedCardUpdateRequest) GetBillingAddress() Address`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *ManagedCardUpdateRequest) GetBillingAddressOk() (*Address, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *ManagedCardUpdateRequest) SetBillingAddress(v Address)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *ManagedCardUpdateRequest) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetDeliveryAddress

`func (o *ManagedCardUpdateRequest) GetDeliveryAddress() DeliveryAddress`

GetDeliveryAddress returns the DeliveryAddress field if non-nil, zero value otherwise.

### GetDeliveryAddressOk

`func (o *ManagedCardUpdateRequest) GetDeliveryAddressOk() (*DeliveryAddress, bool)`

GetDeliveryAddressOk returns a tuple with the DeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddress

`func (o *ManagedCardUpdateRequest) SetDeliveryAddress(v DeliveryAddress)`

SetDeliveryAddress sets DeliveryAddress field to given value.

### HasDeliveryAddress

`func (o *ManagedCardUpdateRequest) HasDeliveryAddress() bool`

HasDeliveryAddress returns a boolean if a field has been set.

### GetDeliveryMethod

`func (o *ManagedCardUpdateRequest) GetDeliveryMethod() DeliveryMethod`

GetDeliveryMethod returns the DeliveryMethod field if non-nil, zero value otherwise.

### GetDeliveryMethodOk

`func (o *ManagedCardUpdateRequest) GetDeliveryMethodOk() (*DeliveryMethod, bool)`

GetDeliveryMethodOk returns a tuple with the DeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethod

`func (o *ManagedCardUpdateRequest) SetDeliveryMethod(v DeliveryMethod)`

SetDeliveryMethod sets DeliveryMethod field to given value.

### HasDeliveryMethod

`func (o *ManagedCardUpdateRequest) HasDeliveryMethod() bool`

HasDeliveryMethod returns a boolean if a field has been set.

### GetDigitalWallets

`func (o *ManagedCardUpdateRequest) GetDigitalWallets() DigitalWallets`

GetDigitalWallets returns the DigitalWallets field if non-nil, zero value otherwise.

### GetDigitalWalletsOk

`func (o *ManagedCardUpdateRequest) GetDigitalWalletsOk() (*DigitalWallets, bool)`

GetDigitalWalletsOk returns a tuple with the DigitalWallets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalWallets

`func (o *ManagedCardUpdateRequest) SetDigitalWallets(v DigitalWallets)`

SetDigitalWallets sets DigitalWallets field to given value.

### HasDigitalWallets

`func (o *ManagedCardUpdateRequest) HasDigitalWallets() bool`

HasDigitalWallets returns a boolean if a field has been set.

### GetAuthForwardingDefaultTimeoutDecision

`func (o *ManagedCardUpdateRequest) GetAuthForwardingDefaultTimeoutDecision() interface{}`

GetAuthForwardingDefaultTimeoutDecision returns the AuthForwardingDefaultTimeoutDecision field if non-nil, zero value otherwise.

### GetAuthForwardingDefaultTimeoutDecisionOk

`func (o *ManagedCardUpdateRequest) GetAuthForwardingDefaultTimeoutDecisionOk() (*interface{}, bool)`

GetAuthForwardingDefaultTimeoutDecisionOk returns a tuple with the AuthForwardingDefaultTimeoutDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthForwardingDefaultTimeoutDecision

`func (o *ManagedCardUpdateRequest) SetAuthForwardingDefaultTimeoutDecision(v interface{})`

SetAuthForwardingDefaultTimeoutDecision sets AuthForwardingDefaultTimeoutDecision field to given value.

### HasAuthForwardingDefaultTimeoutDecision

`func (o *ManagedCardUpdateRequest) HasAuthForwardingDefaultTimeoutDecision() bool`

HasAuthForwardingDefaultTimeoutDecision returns a boolean if a field has been set.

### SetAuthForwardingDefaultTimeoutDecisionNil

`func (o *ManagedCardUpdateRequest) SetAuthForwardingDefaultTimeoutDecisionNil(b bool)`

 SetAuthForwardingDefaultTimeoutDecisionNil sets the value for AuthForwardingDefaultTimeoutDecision to be an explicit nil

### UnsetAuthForwardingDefaultTimeoutDecision
`func (o *ManagedCardUpdateRequest) UnsetAuthForwardingDefaultTimeoutDecision()`

UnsetAuthForwardingDefaultTimeoutDecision ensures that no value is present for AuthForwardingDefaultTimeoutDecision, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


