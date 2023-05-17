# ManagedCardPhysicalUpgradeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductReference** | Pointer to **interface{}** | The product design reference to be used for this physical card.  This reference will be given to you once your physical card programme has been set up. If only one card programme is configured, this field can be left empty.  | [optional] 
**CarrierType** | Pointer to **interface{}** | The carrier type to be used for this physical card.  This reference will be given to you once your physical card carrier has been set up. If only one carrier is configured, this field can be left empty.  | [optional] 
**DeliveryMethod** | Pointer to [**DeliveryMethod**](DeliveryMethod.md) |  | [optional] 
**DeliveryAddress** | [**DeliveryAddress**](DeliveryAddress.md) |  | 
**ActivationCode** | **interface{}** | The unique code to be used to activate the physical card. | 
**Pin** | Pointer to [**SensitivePin**](SensitivePin.md) |  | [optional] 
**NameOnCardLine2** | Pointer to **interface{}** | Line 2 of the &#39;name on card&#39; field. For Physical cards, this field will be printed on the card. The maximum characters allowed will depend on the design chosen and will be provided to you by Weavr when setting up your plastic cards. This field is deprecated. | [optional] 

## Methods

### NewManagedCardPhysicalUpgradeRequest

`func NewManagedCardPhysicalUpgradeRequest(deliveryAddress DeliveryAddress, activationCode interface{}, ) *ManagedCardPhysicalUpgradeRequest`

NewManagedCardPhysicalUpgradeRequest instantiates a new ManagedCardPhysicalUpgradeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedCardPhysicalUpgradeRequestWithDefaults

`func NewManagedCardPhysicalUpgradeRequestWithDefaults() *ManagedCardPhysicalUpgradeRequest`

NewManagedCardPhysicalUpgradeRequestWithDefaults instantiates a new ManagedCardPhysicalUpgradeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductReference

`func (o *ManagedCardPhysicalUpgradeRequest) GetProductReference() interface{}`

GetProductReference returns the ProductReference field if non-nil, zero value otherwise.

### GetProductReferenceOk

`func (o *ManagedCardPhysicalUpgradeRequest) GetProductReferenceOk() (*interface{}, bool)`

GetProductReferenceOk returns a tuple with the ProductReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductReference

`func (o *ManagedCardPhysicalUpgradeRequest) SetProductReference(v interface{})`

SetProductReference sets ProductReference field to given value.

### HasProductReference

`func (o *ManagedCardPhysicalUpgradeRequest) HasProductReference() bool`

HasProductReference returns a boolean if a field has been set.

### SetProductReferenceNil

`func (o *ManagedCardPhysicalUpgradeRequest) SetProductReferenceNil(b bool)`

 SetProductReferenceNil sets the value for ProductReference to be an explicit nil

### UnsetProductReference
`func (o *ManagedCardPhysicalUpgradeRequest) UnsetProductReference()`

UnsetProductReference ensures that no value is present for ProductReference, not even an explicit nil
### GetCarrierType

`func (o *ManagedCardPhysicalUpgradeRequest) GetCarrierType() interface{}`

GetCarrierType returns the CarrierType field if non-nil, zero value otherwise.

### GetCarrierTypeOk

`func (o *ManagedCardPhysicalUpgradeRequest) GetCarrierTypeOk() (*interface{}, bool)`

GetCarrierTypeOk returns a tuple with the CarrierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierType

`func (o *ManagedCardPhysicalUpgradeRequest) SetCarrierType(v interface{})`

SetCarrierType sets CarrierType field to given value.

### HasCarrierType

`func (o *ManagedCardPhysicalUpgradeRequest) HasCarrierType() bool`

HasCarrierType returns a boolean if a field has been set.

### SetCarrierTypeNil

`func (o *ManagedCardPhysicalUpgradeRequest) SetCarrierTypeNil(b bool)`

 SetCarrierTypeNil sets the value for CarrierType to be an explicit nil

### UnsetCarrierType
`func (o *ManagedCardPhysicalUpgradeRequest) UnsetCarrierType()`

UnsetCarrierType ensures that no value is present for CarrierType, not even an explicit nil
### GetDeliveryMethod

`func (o *ManagedCardPhysicalUpgradeRequest) GetDeliveryMethod() DeliveryMethod`

GetDeliveryMethod returns the DeliveryMethod field if non-nil, zero value otherwise.

### GetDeliveryMethodOk

`func (o *ManagedCardPhysicalUpgradeRequest) GetDeliveryMethodOk() (*DeliveryMethod, bool)`

GetDeliveryMethodOk returns a tuple with the DeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethod

`func (o *ManagedCardPhysicalUpgradeRequest) SetDeliveryMethod(v DeliveryMethod)`

SetDeliveryMethod sets DeliveryMethod field to given value.

### HasDeliveryMethod

`func (o *ManagedCardPhysicalUpgradeRequest) HasDeliveryMethod() bool`

HasDeliveryMethod returns a boolean if a field has been set.

### GetDeliveryAddress

`func (o *ManagedCardPhysicalUpgradeRequest) GetDeliveryAddress() DeliveryAddress`

GetDeliveryAddress returns the DeliveryAddress field if non-nil, zero value otherwise.

### GetDeliveryAddressOk

`func (o *ManagedCardPhysicalUpgradeRequest) GetDeliveryAddressOk() (*DeliveryAddress, bool)`

GetDeliveryAddressOk returns a tuple with the DeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddress

`func (o *ManagedCardPhysicalUpgradeRequest) SetDeliveryAddress(v DeliveryAddress)`

SetDeliveryAddress sets DeliveryAddress field to given value.


### GetActivationCode

`func (o *ManagedCardPhysicalUpgradeRequest) GetActivationCode() interface{}`

GetActivationCode returns the ActivationCode field if non-nil, zero value otherwise.

### GetActivationCodeOk

`func (o *ManagedCardPhysicalUpgradeRequest) GetActivationCodeOk() (*interface{}, bool)`

GetActivationCodeOk returns a tuple with the ActivationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationCode

`func (o *ManagedCardPhysicalUpgradeRequest) SetActivationCode(v interface{})`

SetActivationCode sets ActivationCode field to given value.


### SetActivationCodeNil

`func (o *ManagedCardPhysicalUpgradeRequest) SetActivationCodeNil(b bool)`

 SetActivationCodeNil sets the value for ActivationCode to be an explicit nil

### UnsetActivationCode
`func (o *ManagedCardPhysicalUpgradeRequest) UnsetActivationCode()`

UnsetActivationCode ensures that no value is present for ActivationCode, not even an explicit nil
### GetPin

`func (o *ManagedCardPhysicalUpgradeRequest) GetPin() SensitivePin`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *ManagedCardPhysicalUpgradeRequest) GetPinOk() (*SensitivePin, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *ManagedCardPhysicalUpgradeRequest) SetPin(v SensitivePin)`

SetPin sets Pin field to given value.

### HasPin

`func (o *ManagedCardPhysicalUpgradeRequest) HasPin() bool`

HasPin returns a boolean if a field has been set.

### GetNameOnCardLine2

`func (o *ManagedCardPhysicalUpgradeRequest) GetNameOnCardLine2() interface{}`

GetNameOnCardLine2 returns the NameOnCardLine2 field if non-nil, zero value otherwise.

### GetNameOnCardLine2Ok

`func (o *ManagedCardPhysicalUpgradeRequest) GetNameOnCardLine2Ok() (*interface{}, bool)`

GetNameOnCardLine2Ok returns a tuple with the NameOnCardLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnCardLine2

`func (o *ManagedCardPhysicalUpgradeRequest) SetNameOnCardLine2(v interface{})`

SetNameOnCardLine2 sets NameOnCardLine2 field to given value.

### HasNameOnCardLine2

`func (o *ManagedCardPhysicalUpgradeRequest) HasNameOnCardLine2() bool`

HasNameOnCardLine2 returns a boolean if a field has been set.

### SetNameOnCardLine2Nil

`func (o *ManagedCardPhysicalUpgradeRequest) SetNameOnCardLine2Nil(b bool)`

 SetNameOnCardLine2Nil sets the value for NameOnCardLine2 to be an explicit nil

### UnsetNameOnCardLine2
`func (o *ManagedCardPhysicalUpgradeRequest) UnsetNameOnCardLine2()`

UnsetNameOnCardLine2 ensures that no value is present for NameOnCardLine2, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


