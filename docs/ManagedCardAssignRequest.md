# ManagedCardAssignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalReference** | **interface{}** | Unique code identifying a card. | 
**ActivationCode** | **interface{}** | The code to be used to activate the physical card. Depending on how the cards were created, this may be a code set upon creation, or it may be the same as the &#x60;externalReference&#x60;. | 
**FriendlyName** | **interface{}** | The friendly name given to the card. | 
**NameOnCard** | **interface{}** | The card holder’s name for the card. This may be verified by merchants when the card is used online. For Physical cards, this field will be printed on the card. The maximum characters allowed will depend on the design chosen and will be provided to you by Weavr when setting up your plastic cards. | 
**NameOnCardLine2** | Pointer to **interface{}** | Line 2 of the &#39;name on card&#39; field. For Physical cards, this field will be printed on the card. The maximum characters allowed will depend on the design chosen and will be provided to you by Weavr when setting up your plastic cards. | [optional] 
**BillingAddress** | [**Address**](Address.md) | The billing address set for the card holder, required for AVS checking. | 
**CardholderMobileNumber** | Pointer to **interface{}** | The mobile number including country code of the card holder, needed for 3DS challenge. | [optional] 
**AuthForwardingDefaultTimeoutDecision** | Pointer to **interface{}** | Default decision for auth forwarding on timeout | [optional] 

## Methods

### NewManagedCardAssignRequest

`func NewManagedCardAssignRequest(externalReference interface{}, activationCode interface{}, friendlyName interface{}, nameOnCard interface{}, billingAddress Address, ) *ManagedCardAssignRequest`

NewManagedCardAssignRequest instantiates a new ManagedCardAssignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedCardAssignRequestWithDefaults

`func NewManagedCardAssignRequestWithDefaults() *ManagedCardAssignRequest`

NewManagedCardAssignRequestWithDefaults instantiates a new ManagedCardAssignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalReference

`func (o *ManagedCardAssignRequest) GetExternalReference() interface{}`

GetExternalReference returns the ExternalReference field if non-nil, zero value otherwise.

### GetExternalReferenceOk

`func (o *ManagedCardAssignRequest) GetExternalReferenceOk() (*interface{}, bool)`

GetExternalReferenceOk returns a tuple with the ExternalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReference

`func (o *ManagedCardAssignRequest) SetExternalReference(v interface{})`

SetExternalReference sets ExternalReference field to given value.


### SetExternalReferenceNil

`func (o *ManagedCardAssignRequest) SetExternalReferenceNil(b bool)`

 SetExternalReferenceNil sets the value for ExternalReference to be an explicit nil

### UnsetExternalReference
`func (o *ManagedCardAssignRequest) UnsetExternalReference()`

UnsetExternalReference ensures that no value is present for ExternalReference, not even an explicit nil
### GetActivationCode

`func (o *ManagedCardAssignRequest) GetActivationCode() interface{}`

GetActivationCode returns the ActivationCode field if non-nil, zero value otherwise.

### GetActivationCodeOk

`func (o *ManagedCardAssignRequest) GetActivationCodeOk() (*interface{}, bool)`

GetActivationCodeOk returns a tuple with the ActivationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationCode

`func (o *ManagedCardAssignRequest) SetActivationCode(v interface{})`

SetActivationCode sets ActivationCode field to given value.


### SetActivationCodeNil

`func (o *ManagedCardAssignRequest) SetActivationCodeNil(b bool)`

 SetActivationCodeNil sets the value for ActivationCode to be an explicit nil

### UnsetActivationCode
`func (o *ManagedCardAssignRequest) UnsetActivationCode()`

UnsetActivationCode ensures that no value is present for ActivationCode, not even an explicit nil
### GetFriendlyName

`func (o *ManagedCardAssignRequest) GetFriendlyName() interface{}`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ManagedCardAssignRequest) GetFriendlyNameOk() (*interface{}, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ManagedCardAssignRequest) SetFriendlyName(v interface{})`

SetFriendlyName sets FriendlyName field to given value.


### SetFriendlyNameNil

`func (o *ManagedCardAssignRequest) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ManagedCardAssignRequest) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetNameOnCard

`func (o *ManagedCardAssignRequest) GetNameOnCard() interface{}`

GetNameOnCard returns the NameOnCard field if non-nil, zero value otherwise.

### GetNameOnCardOk

`func (o *ManagedCardAssignRequest) GetNameOnCardOk() (*interface{}, bool)`

GetNameOnCardOk returns a tuple with the NameOnCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnCard

`func (o *ManagedCardAssignRequest) SetNameOnCard(v interface{})`

SetNameOnCard sets NameOnCard field to given value.


### SetNameOnCardNil

`func (o *ManagedCardAssignRequest) SetNameOnCardNil(b bool)`

 SetNameOnCardNil sets the value for NameOnCard to be an explicit nil

### UnsetNameOnCard
`func (o *ManagedCardAssignRequest) UnsetNameOnCard()`

UnsetNameOnCard ensures that no value is present for NameOnCard, not even an explicit nil
### GetNameOnCardLine2

`func (o *ManagedCardAssignRequest) GetNameOnCardLine2() interface{}`

GetNameOnCardLine2 returns the NameOnCardLine2 field if non-nil, zero value otherwise.

### GetNameOnCardLine2Ok

`func (o *ManagedCardAssignRequest) GetNameOnCardLine2Ok() (*interface{}, bool)`

GetNameOnCardLine2Ok returns a tuple with the NameOnCardLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnCardLine2

`func (o *ManagedCardAssignRequest) SetNameOnCardLine2(v interface{})`

SetNameOnCardLine2 sets NameOnCardLine2 field to given value.

### HasNameOnCardLine2

`func (o *ManagedCardAssignRequest) HasNameOnCardLine2() bool`

HasNameOnCardLine2 returns a boolean if a field has been set.

### SetNameOnCardLine2Nil

`func (o *ManagedCardAssignRequest) SetNameOnCardLine2Nil(b bool)`

 SetNameOnCardLine2Nil sets the value for NameOnCardLine2 to be an explicit nil

### UnsetNameOnCardLine2
`func (o *ManagedCardAssignRequest) UnsetNameOnCardLine2()`

UnsetNameOnCardLine2 ensures that no value is present for NameOnCardLine2, not even an explicit nil
### GetBillingAddress

`func (o *ManagedCardAssignRequest) GetBillingAddress() Address`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *ManagedCardAssignRequest) GetBillingAddressOk() (*Address, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *ManagedCardAssignRequest) SetBillingAddress(v Address)`

SetBillingAddress sets BillingAddress field to given value.


### GetCardholderMobileNumber

`func (o *ManagedCardAssignRequest) GetCardholderMobileNumber() interface{}`

GetCardholderMobileNumber returns the CardholderMobileNumber field if non-nil, zero value otherwise.

### GetCardholderMobileNumberOk

`func (o *ManagedCardAssignRequest) GetCardholderMobileNumberOk() (*interface{}, bool)`

GetCardholderMobileNumberOk returns a tuple with the CardholderMobileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderMobileNumber

`func (o *ManagedCardAssignRequest) SetCardholderMobileNumber(v interface{})`

SetCardholderMobileNumber sets CardholderMobileNumber field to given value.

### HasCardholderMobileNumber

`func (o *ManagedCardAssignRequest) HasCardholderMobileNumber() bool`

HasCardholderMobileNumber returns a boolean if a field has been set.

### SetCardholderMobileNumberNil

`func (o *ManagedCardAssignRequest) SetCardholderMobileNumberNil(b bool)`

 SetCardholderMobileNumberNil sets the value for CardholderMobileNumber to be an explicit nil

### UnsetCardholderMobileNumber
`func (o *ManagedCardAssignRequest) UnsetCardholderMobileNumber()`

UnsetCardholderMobileNumber ensures that no value is present for CardholderMobileNumber, not even an explicit nil
### GetAuthForwardingDefaultTimeoutDecision

`func (o *ManagedCardAssignRequest) GetAuthForwardingDefaultTimeoutDecision() interface{}`

GetAuthForwardingDefaultTimeoutDecision returns the AuthForwardingDefaultTimeoutDecision field if non-nil, zero value otherwise.

### GetAuthForwardingDefaultTimeoutDecisionOk

`func (o *ManagedCardAssignRequest) GetAuthForwardingDefaultTimeoutDecisionOk() (*interface{}, bool)`

GetAuthForwardingDefaultTimeoutDecisionOk returns a tuple with the AuthForwardingDefaultTimeoutDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthForwardingDefaultTimeoutDecision

`func (o *ManagedCardAssignRequest) SetAuthForwardingDefaultTimeoutDecision(v interface{})`

SetAuthForwardingDefaultTimeoutDecision sets AuthForwardingDefaultTimeoutDecision field to given value.

### HasAuthForwardingDefaultTimeoutDecision

`func (o *ManagedCardAssignRequest) HasAuthForwardingDefaultTimeoutDecision() bool`

HasAuthForwardingDefaultTimeoutDecision returns a boolean if a field has been set.

### SetAuthForwardingDefaultTimeoutDecisionNil

`func (o *ManagedCardAssignRequest) SetAuthForwardingDefaultTimeoutDecisionNil(b bool)`

 SetAuthForwardingDefaultTimeoutDecisionNil sets the value for AuthForwardingDefaultTimeoutDecision to be an explicit nil

### UnsetAuthForwardingDefaultTimeoutDecision
`func (o *ManagedCardAssignRequest) UnsetAuthForwardingDefaultTimeoutDecision()`

UnsetAuthForwardingDefaultTimeoutDecision ensures that no value is present for AuthForwardingDefaultTimeoutDecision, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


