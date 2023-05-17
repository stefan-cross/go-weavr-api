# OutgoingWireTransferCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **interface{}** | The profile Id which a specific identity, instrument or transaction type is linked to.  Profiles contain configuration and determine behavioral aspects of the newly created transaction, for example, fees that may apply.  You can have one or more profiles linked to your application, and these can be used to drive different behaviors according to your product&#39;s needs.  Profile Ids can be found in the Multi Portal, in the API Credentials page.  | 
**Tag** | Pointer to **interface{}** | The tag field is a custom field that can be used to search and filter. | [optional] 
**SourceInstrument** | [**OutgoingWireTransferCreateRequestSourceInstrument**](OutgoingWireTransferCreateRequestSourceInstrument.md) |  | 
**TransferAmount** | [**CurrencyAmount**](CurrencyAmount.md) | The amount, in same currency as source and destination, that was transferred from the source. | 
**Description** | Pointer to **interface{}** | These details are passed to the beneficiary as the &#x60;reference&#x60; and the allowed length is dependent on the payment type:  &#x60;SEPA&#x60; and &#x60;SWIFT&#x60; &lt;&#x3D; &#x60;35&#x60; characters  &#x60;Faster Payments&#x60; &lt;&#x3D; &#x60;18&#x60; characters  | [optional] 
**DestinationBeneficiary** | [**OutgoingWireTransferBeneficiary**](OutgoingWireTransferBeneficiary.md) |  | 

## Methods

### NewOutgoingWireTransferCreateRequest

`func NewOutgoingWireTransferCreateRequest(profileId interface{}, sourceInstrument OutgoingWireTransferCreateRequestSourceInstrument, transferAmount CurrencyAmount, destinationBeneficiary OutgoingWireTransferBeneficiary, ) *OutgoingWireTransferCreateRequest`

NewOutgoingWireTransferCreateRequest instantiates a new OutgoingWireTransferCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutgoingWireTransferCreateRequestWithDefaults

`func NewOutgoingWireTransferCreateRequestWithDefaults() *OutgoingWireTransferCreateRequest`

NewOutgoingWireTransferCreateRequestWithDefaults instantiates a new OutgoingWireTransferCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *OutgoingWireTransferCreateRequest) GetProfileId() interface{}`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *OutgoingWireTransferCreateRequest) GetProfileIdOk() (*interface{}, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *OutgoingWireTransferCreateRequest) SetProfileId(v interface{})`

SetProfileId sets ProfileId field to given value.


### SetProfileIdNil

`func (o *OutgoingWireTransferCreateRequest) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *OutgoingWireTransferCreateRequest) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
### GetTag

`func (o *OutgoingWireTransferCreateRequest) GetTag() interface{}`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *OutgoingWireTransferCreateRequest) GetTagOk() (*interface{}, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *OutgoingWireTransferCreateRequest) SetTag(v interface{})`

SetTag sets Tag field to given value.

### HasTag

`func (o *OutgoingWireTransferCreateRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *OutgoingWireTransferCreateRequest) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *OutgoingWireTransferCreateRequest) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil
### GetSourceInstrument

`func (o *OutgoingWireTransferCreateRequest) GetSourceInstrument() OutgoingWireTransferCreateRequestSourceInstrument`

GetSourceInstrument returns the SourceInstrument field if non-nil, zero value otherwise.

### GetSourceInstrumentOk

`func (o *OutgoingWireTransferCreateRequest) GetSourceInstrumentOk() (*OutgoingWireTransferCreateRequestSourceInstrument, bool)`

GetSourceInstrumentOk returns a tuple with the SourceInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInstrument

`func (o *OutgoingWireTransferCreateRequest) SetSourceInstrument(v OutgoingWireTransferCreateRequestSourceInstrument)`

SetSourceInstrument sets SourceInstrument field to given value.


### GetTransferAmount

`func (o *OutgoingWireTransferCreateRequest) GetTransferAmount() CurrencyAmount`

GetTransferAmount returns the TransferAmount field if non-nil, zero value otherwise.

### GetTransferAmountOk

`func (o *OutgoingWireTransferCreateRequest) GetTransferAmountOk() (*CurrencyAmount, bool)`

GetTransferAmountOk returns a tuple with the TransferAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAmount

`func (o *OutgoingWireTransferCreateRequest) SetTransferAmount(v CurrencyAmount)`

SetTransferAmount sets TransferAmount field to given value.


### GetDescription

`func (o *OutgoingWireTransferCreateRequest) GetDescription() interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OutgoingWireTransferCreateRequest) GetDescriptionOk() (*interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OutgoingWireTransferCreateRequest) SetDescription(v interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OutgoingWireTransferCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OutgoingWireTransferCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OutgoingWireTransferCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDestinationBeneficiary

`func (o *OutgoingWireTransferCreateRequest) GetDestinationBeneficiary() OutgoingWireTransferBeneficiary`

GetDestinationBeneficiary returns the DestinationBeneficiary field if non-nil, zero value otherwise.

### GetDestinationBeneficiaryOk

`func (o *OutgoingWireTransferCreateRequest) GetDestinationBeneficiaryOk() (*OutgoingWireTransferBeneficiary, bool)`

GetDestinationBeneficiaryOk returns a tuple with the DestinationBeneficiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationBeneficiary

`func (o *OutgoingWireTransferCreateRequest) SetDestinationBeneficiary(v OutgoingWireTransferBeneficiary)`

SetDestinationBeneficiary sets DestinationBeneficiary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


