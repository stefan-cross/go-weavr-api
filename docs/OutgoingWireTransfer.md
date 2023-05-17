# OutgoingWireTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **interface{}** | The unique identifier of an Outgoing Wire Transfer transaction. | 
**ProfileId** | **interface{}** | The profile Id which a specific identity, instrument or transaction type is linked to.  Profiles contain configuration and determine behavioral aspects of the newly created transaction, for example, fees that may apply.  You can have one or more profiles linked to your application, and these can be used to drive different behaviors according to your product&#39;s needs.  Profile Ids can be found in the Multi Portal, in the API Credentials page.  | 
**Tag** | Pointer to **interface{}** | The tag field is a custom field that can be used to search and filter. | [optional] 
**SourceInstrument** | [**OutgoingWireTransferSourceInstrument**](OutgoingWireTransferSourceInstrument.md) |  | 
**TransferAmount** | [**CurrencyAmount**](CurrencyAmount.md) | The amount, in same currency as source and destination, that was transferred from the source. | 
**Description** | Pointer to **interface{}** |  Transaction description. | [optional] 
**Type** | **interface{}** | The wire transfer rails used, from SEPA, Faster Payments or SWIFT. | 
**Destination** | [**OutgoingWireTransferBeneficiary**](OutgoingWireTransferBeneficiary.md) |  | 
**State** | [**OutgoingWireTransferState**](OutgoingWireTransferState.md) |  | 
**RejectedInfo** | Pointer to [**OutgoingWireTransferRejectedInfo**](OutgoingWireTransferRejectedInfo.md) |  | [optional] 
**CreationTimestamp** | **interface{}** | The time when the transaction was created, expressed in Epoch timestamp using millisecond precision. | 

## Methods

### NewOutgoingWireTransfer

`func NewOutgoingWireTransfer(id interface{}, profileId interface{}, sourceInstrument OutgoingWireTransferSourceInstrument, transferAmount CurrencyAmount, type_ interface{}, destination OutgoingWireTransferBeneficiary, state OutgoingWireTransferState, creationTimestamp interface{}, ) *OutgoingWireTransfer`

NewOutgoingWireTransfer instantiates a new OutgoingWireTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutgoingWireTransferWithDefaults

`func NewOutgoingWireTransferWithDefaults() *OutgoingWireTransfer`

NewOutgoingWireTransferWithDefaults instantiates a new OutgoingWireTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OutgoingWireTransfer) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutgoingWireTransfer) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutgoingWireTransfer) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *OutgoingWireTransfer) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *OutgoingWireTransfer) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetProfileId

`func (o *OutgoingWireTransfer) GetProfileId() interface{}`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *OutgoingWireTransfer) GetProfileIdOk() (*interface{}, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *OutgoingWireTransfer) SetProfileId(v interface{})`

SetProfileId sets ProfileId field to given value.


### SetProfileIdNil

`func (o *OutgoingWireTransfer) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *OutgoingWireTransfer) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
### GetTag

`func (o *OutgoingWireTransfer) GetTag() interface{}`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *OutgoingWireTransfer) GetTagOk() (*interface{}, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *OutgoingWireTransfer) SetTag(v interface{})`

SetTag sets Tag field to given value.

### HasTag

`func (o *OutgoingWireTransfer) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *OutgoingWireTransfer) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *OutgoingWireTransfer) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil
### GetSourceInstrument

`func (o *OutgoingWireTransfer) GetSourceInstrument() OutgoingWireTransferSourceInstrument`

GetSourceInstrument returns the SourceInstrument field if non-nil, zero value otherwise.

### GetSourceInstrumentOk

`func (o *OutgoingWireTransfer) GetSourceInstrumentOk() (*OutgoingWireTransferSourceInstrument, bool)`

GetSourceInstrumentOk returns a tuple with the SourceInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInstrument

`func (o *OutgoingWireTransfer) SetSourceInstrument(v OutgoingWireTransferSourceInstrument)`

SetSourceInstrument sets SourceInstrument field to given value.


### GetTransferAmount

`func (o *OutgoingWireTransfer) GetTransferAmount() CurrencyAmount`

GetTransferAmount returns the TransferAmount field if non-nil, zero value otherwise.

### GetTransferAmountOk

`func (o *OutgoingWireTransfer) GetTransferAmountOk() (*CurrencyAmount, bool)`

GetTransferAmountOk returns a tuple with the TransferAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAmount

`func (o *OutgoingWireTransfer) SetTransferAmount(v CurrencyAmount)`

SetTransferAmount sets TransferAmount field to given value.


### GetDescription

`func (o *OutgoingWireTransfer) GetDescription() interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OutgoingWireTransfer) GetDescriptionOk() (*interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OutgoingWireTransfer) SetDescription(v interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OutgoingWireTransfer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OutgoingWireTransfer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OutgoingWireTransfer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *OutgoingWireTransfer) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutgoingWireTransfer) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutgoingWireTransfer) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *OutgoingWireTransfer) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *OutgoingWireTransfer) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetDestination

`func (o *OutgoingWireTransfer) GetDestination() OutgoingWireTransferBeneficiary`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *OutgoingWireTransfer) GetDestinationOk() (*OutgoingWireTransferBeneficiary, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *OutgoingWireTransfer) SetDestination(v OutgoingWireTransferBeneficiary)`

SetDestination sets Destination field to given value.


### GetState

`func (o *OutgoingWireTransfer) GetState() OutgoingWireTransferState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OutgoingWireTransfer) GetStateOk() (*OutgoingWireTransferState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OutgoingWireTransfer) SetState(v OutgoingWireTransferState)`

SetState sets State field to given value.


### GetRejectedInfo

`func (o *OutgoingWireTransfer) GetRejectedInfo() OutgoingWireTransferRejectedInfo`

GetRejectedInfo returns the RejectedInfo field if non-nil, zero value otherwise.

### GetRejectedInfoOk

`func (o *OutgoingWireTransfer) GetRejectedInfoOk() (*OutgoingWireTransferRejectedInfo, bool)`

GetRejectedInfoOk returns a tuple with the RejectedInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedInfo

`func (o *OutgoingWireTransfer) SetRejectedInfo(v OutgoingWireTransferRejectedInfo)`

SetRejectedInfo sets RejectedInfo field to given value.

### HasRejectedInfo

`func (o *OutgoingWireTransfer) HasRejectedInfo() bool`

HasRejectedInfo returns a boolean if a field has been set.

### GetCreationTimestamp

`func (o *OutgoingWireTransfer) GetCreationTimestamp() interface{}`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *OutgoingWireTransfer) GetCreationTimestampOk() (*interface{}, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *OutgoingWireTransfer) SetCreationTimestamp(v interface{})`

SetCreationTimestamp sets CreationTimestamp field to given value.


### SetCreationTimestampNil

`func (o *OutgoingWireTransfer) SetCreationTimestampNil(b bool)`

 SetCreationTimestampNil sets the value for CreationTimestamp to be an explicit nil

### UnsetCreationTimestamp
`func (o *OutgoingWireTransfer) UnsetCreationTimestamp()`

UnsetCreationTimestamp ensures that no value is present for CreationTimestamp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


