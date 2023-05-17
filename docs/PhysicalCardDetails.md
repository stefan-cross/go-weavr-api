# PhysicalCardDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductReference** | Pointer to **interface{}** | The unique reference indicating the configuration of the physical card. Example the card design. | [optional] 
**CarrierType** | Pointer to **interface{}** |  | [optional] 
**PendingActivation** | **interface{}** | Indicates if the physical card is activated for physical usage such as with physical terminals. | 
**PinBlocked** | Pointer to **interface{}** | Indicates if the physical card is blocked due to providing incorrect PINs. | [optional] 
**ManufacturingState** | Pointer to [**ManufacturingState**](ManufacturingState.md) |  | [optional] 
**Replacement** | Pointer to [**PhysicalCardDetailsReplacement**](PhysicalCardDetailsReplacement.md) |  | [optional] 
**DeliveryAddress** | Pointer to [**DeliveryAddress**](DeliveryAddress.md) |  | [optional] 
**DeliveryMethod** | Pointer to [**DeliveryMethod**](DeliveryMethod.md) |  | [optional] 
**DeliveryTrackingCode** | Pointer to **interface{}** | The delivery tracking code for tracking the physical card&#39;s delivery status. | [optional] 
**DeliveryTrackingMethod** | Pointer to **interface{}** | The delivery tracking method for tracking the physical card&#39;s delivery status. | [optional] 
**NameOnCardLine2** | Pointer to **interface{}** | Line 2 of the &#39;name on card&#39; field. For Physical cards, this field will be printed on the card. The maximum characters allowed will depend on the design chosen and will be provided to you by Weavr when setting up your plastic cards. This field is deprecated | [optional] 

## Methods

### NewPhysicalCardDetails

`func NewPhysicalCardDetails(pendingActivation interface{}, ) *PhysicalCardDetails`

NewPhysicalCardDetails instantiates a new PhysicalCardDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalCardDetailsWithDefaults

`func NewPhysicalCardDetailsWithDefaults() *PhysicalCardDetails`

NewPhysicalCardDetailsWithDefaults instantiates a new PhysicalCardDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductReference

`func (o *PhysicalCardDetails) GetProductReference() interface{}`

GetProductReference returns the ProductReference field if non-nil, zero value otherwise.

### GetProductReferenceOk

`func (o *PhysicalCardDetails) GetProductReferenceOk() (*interface{}, bool)`

GetProductReferenceOk returns a tuple with the ProductReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductReference

`func (o *PhysicalCardDetails) SetProductReference(v interface{})`

SetProductReference sets ProductReference field to given value.

### HasProductReference

`func (o *PhysicalCardDetails) HasProductReference() bool`

HasProductReference returns a boolean if a field has been set.

### SetProductReferenceNil

`func (o *PhysicalCardDetails) SetProductReferenceNil(b bool)`

 SetProductReferenceNil sets the value for ProductReference to be an explicit nil

### UnsetProductReference
`func (o *PhysicalCardDetails) UnsetProductReference()`

UnsetProductReference ensures that no value is present for ProductReference, not even an explicit nil
### GetCarrierType

`func (o *PhysicalCardDetails) GetCarrierType() interface{}`

GetCarrierType returns the CarrierType field if non-nil, zero value otherwise.

### GetCarrierTypeOk

`func (o *PhysicalCardDetails) GetCarrierTypeOk() (*interface{}, bool)`

GetCarrierTypeOk returns a tuple with the CarrierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierType

`func (o *PhysicalCardDetails) SetCarrierType(v interface{})`

SetCarrierType sets CarrierType field to given value.

### HasCarrierType

`func (o *PhysicalCardDetails) HasCarrierType() bool`

HasCarrierType returns a boolean if a field has been set.

### SetCarrierTypeNil

`func (o *PhysicalCardDetails) SetCarrierTypeNil(b bool)`

 SetCarrierTypeNil sets the value for CarrierType to be an explicit nil

### UnsetCarrierType
`func (o *PhysicalCardDetails) UnsetCarrierType()`

UnsetCarrierType ensures that no value is present for CarrierType, not even an explicit nil
### GetPendingActivation

`func (o *PhysicalCardDetails) GetPendingActivation() interface{}`

GetPendingActivation returns the PendingActivation field if non-nil, zero value otherwise.

### GetPendingActivationOk

`func (o *PhysicalCardDetails) GetPendingActivationOk() (*interface{}, bool)`

GetPendingActivationOk returns a tuple with the PendingActivation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingActivation

`func (o *PhysicalCardDetails) SetPendingActivation(v interface{})`

SetPendingActivation sets PendingActivation field to given value.


### SetPendingActivationNil

`func (o *PhysicalCardDetails) SetPendingActivationNil(b bool)`

 SetPendingActivationNil sets the value for PendingActivation to be an explicit nil

### UnsetPendingActivation
`func (o *PhysicalCardDetails) UnsetPendingActivation()`

UnsetPendingActivation ensures that no value is present for PendingActivation, not even an explicit nil
### GetPinBlocked

`func (o *PhysicalCardDetails) GetPinBlocked() interface{}`

GetPinBlocked returns the PinBlocked field if non-nil, zero value otherwise.

### GetPinBlockedOk

`func (o *PhysicalCardDetails) GetPinBlockedOk() (*interface{}, bool)`

GetPinBlockedOk returns a tuple with the PinBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinBlocked

`func (o *PhysicalCardDetails) SetPinBlocked(v interface{})`

SetPinBlocked sets PinBlocked field to given value.

### HasPinBlocked

`func (o *PhysicalCardDetails) HasPinBlocked() bool`

HasPinBlocked returns a boolean if a field has been set.

### SetPinBlockedNil

`func (o *PhysicalCardDetails) SetPinBlockedNil(b bool)`

 SetPinBlockedNil sets the value for PinBlocked to be an explicit nil

### UnsetPinBlocked
`func (o *PhysicalCardDetails) UnsetPinBlocked()`

UnsetPinBlocked ensures that no value is present for PinBlocked, not even an explicit nil
### GetManufacturingState

`func (o *PhysicalCardDetails) GetManufacturingState() ManufacturingState`

GetManufacturingState returns the ManufacturingState field if non-nil, zero value otherwise.

### GetManufacturingStateOk

`func (o *PhysicalCardDetails) GetManufacturingStateOk() (*ManufacturingState, bool)`

GetManufacturingStateOk returns a tuple with the ManufacturingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturingState

`func (o *PhysicalCardDetails) SetManufacturingState(v ManufacturingState)`

SetManufacturingState sets ManufacturingState field to given value.

### HasManufacturingState

`func (o *PhysicalCardDetails) HasManufacturingState() bool`

HasManufacturingState returns a boolean if a field has been set.

### GetReplacement

`func (o *PhysicalCardDetails) GetReplacement() PhysicalCardDetailsReplacement`

GetReplacement returns the Replacement field if non-nil, zero value otherwise.

### GetReplacementOk

`func (o *PhysicalCardDetails) GetReplacementOk() (*PhysicalCardDetailsReplacement, bool)`

GetReplacementOk returns a tuple with the Replacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacement

`func (o *PhysicalCardDetails) SetReplacement(v PhysicalCardDetailsReplacement)`

SetReplacement sets Replacement field to given value.

### HasReplacement

`func (o *PhysicalCardDetails) HasReplacement() bool`

HasReplacement returns a boolean if a field has been set.

### GetDeliveryAddress

`func (o *PhysicalCardDetails) GetDeliveryAddress() DeliveryAddress`

GetDeliveryAddress returns the DeliveryAddress field if non-nil, zero value otherwise.

### GetDeliveryAddressOk

`func (o *PhysicalCardDetails) GetDeliveryAddressOk() (*DeliveryAddress, bool)`

GetDeliveryAddressOk returns a tuple with the DeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddress

`func (o *PhysicalCardDetails) SetDeliveryAddress(v DeliveryAddress)`

SetDeliveryAddress sets DeliveryAddress field to given value.

### HasDeliveryAddress

`func (o *PhysicalCardDetails) HasDeliveryAddress() bool`

HasDeliveryAddress returns a boolean if a field has been set.

### GetDeliveryMethod

`func (o *PhysicalCardDetails) GetDeliveryMethod() DeliveryMethod`

GetDeliveryMethod returns the DeliveryMethod field if non-nil, zero value otherwise.

### GetDeliveryMethodOk

`func (o *PhysicalCardDetails) GetDeliveryMethodOk() (*DeliveryMethod, bool)`

GetDeliveryMethodOk returns a tuple with the DeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethod

`func (o *PhysicalCardDetails) SetDeliveryMethod(v DeliveryMethod)`

SetDeliveryMethod sets DeliveryMethod field to given value.

### HasDeliveryMethod

`func (o *PhysicalCardDetails) HasDeliveryMethod() bool`

HasDeliveryMethod returns a boolean if a field has been set.

### GetDeliveryTrackingCode

`func (o *PhysicalCardDetails) GetDeliveryTrackingCode() interface{}`

GetDeliveryTrackingCode returns the DeliveryTrackingCode field if non-nil, zero value otherwise.

### GetDeliveryTrackingCodeOk

`func (o *PhysicalCardDetails) GetDeliveryTrackingCodeOk() (*interface{}, bool)`

GetDeliveryTrackingCodeOk returns a tuple with the DeliveryTrackingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTrackingCode

`func (o *PhysicalCardDetails) SetDeliveryTrackingCode(v interface{})`

SetDeliveryTrackingCode sets DeliveryTrackingCode field to given value.

### HasDeliveryTrackingCode

`func (o *PhysicalCardDetails) HasDeliveryTrackingCode() bool`

HasDeliveryTrackingCode returns a boolean if a field has been set.

### SetDeliveryTrackingCodeNil

`func (o *PhysicalCardDetails) SetDeliveryTrackingCodeNil(b bool)`

 SetDeliveryTrackingCodeNil sets the value for DeliveryTrackingCode to be an explicit nil

### UnsetDeliveryTrackingCode
`func (o *PhysicalCardDetails) UnsetDeliveryTrackingCode()`

UnsetDeliveryTrackingCode ensures that no value is present for DeliveryTrackingCode, not even an explicit nil
### GetDeliveryTrackingMethod

`func (o *PhysicalCardDetails) GetDeliveryTrackingMethod() interface{}`

GetDeliveryTrackingMethod returns the DeliveryTrackingMethod field if non-nil, zero value otherwise.

### GetDeliveryTrackingMethodOk

`func (o *PhysicalCardDetails) GetDeliveryTrackingMethodOk() (*interface{}, bool)`

GetDeliveryTrackingMethodOk returns a tuple with the DeliveryTrackingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTrackingMethod

`func (o *PhysicalCardDetails) SetDeliveryTrackingMethod(v interface{})`

SetDeliveryTrackingMethod sets DeliveryTrackingMethod field to given value.

### HasDeliveryTrackingMethod

`func (o *PhysicalCardDetails) HasDeliveryTrackingMethod() bool`

HasDeliveryTrackingMethod returns a boolean if a field has been set.

### SetDeliveryTrackingMethodNil

`func (o *PhysicalCardDetails) SetDeliveryTrackingMethodNil(b bool)`

 SetDeliveryTrackingMethodNil sets the value for DeliveryTrackingMethod to be an explicit nil

### UnsetDeliveryTrackingMethod
`func (o *PhysicalCardDetails) UnsetDeliveryTrackingMethod()`

UnsetDeliveryTrackingMethod ensures that no value is present for DeliveryTrackingMethod, not even an explicit nil
### GetNameOnCardLine2

`func (o *PhysicalCardDetails) GetNameOnCardLine2() interface{}`

GetNameOnCardLine2 returns the NameOnCardLine2 field if non-nil, zero value otherwise.

### GetNameOnCardLine2Ok

`func (o *PhysicalCardDetails) GetNameOnCardLine2Ok() (*interface{}, bool)`

GetNameOnCardLine2Ok returns a tuple with the NameOnCardLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnCardLine2

`func (o *PhysicalCardDetails) SetNameOnCardLine2(v interface{})`

SetNameOnCardLine2 sets NameOnCardLine2 field to given value.

### HasNameOnCardLine2

`func (o *PhysicalCardDetails) HasNameOnCardLine2() bool`

HasNameOnCardLine2 returns a boolean if a field has been set.

### SetNameOnCardLine2Nil

`func (o *PhysicalCardDetails) SetNameOnCardLine2Nil(b bool)`

 SetNameOnCardLine2Nil sets the value for NameOnCardLine2 to be an explicit nil

### UnsetNameOnCardLine2
`func (o *PhysicalCardDetails) UnsetNameOnCardLine2()`

UnsetNameOnCardLine2 ensures that no value is present for NameOnCardLine2, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


