# Transfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **interface{}** | The unique identifier of a Transfer transaction. | 
**ProfileId** | **interface{}** | The profile Id which a specific identity, instrument or transaction type is linked to.  Profiles contain configuration and determine behavioral aspects of the newly created transaction, for example, fees that may apply.  You can have one or more profiles linked to your application, and these can be used to drive different behaviors according to your product&#39;s needs.  Profile Ids can be found in the Multi Portal, in the API Credentials page.  | 
**Tag** | Pointer to **interface{}** | The tag field is a custom field that can be used to search and filter. | [optional] 
**Source** | [**InstrumentId**](InstrumentId.md) | The source managed account or card from where the funds were transferred. | 
**Destination** | [**InstrumentId**](InstrumentId.md) | The destination managed account or card to where the funds were transferred. | 
**DestinationAmount** | [**CurrencyAmount**](CurrencyAmount.md) | The amount, in same currency as source and destination, that was transferred from the source to the destination instrument. | 
**State** | [**TransactionState**](TransactionState.md) |  | 
**CreationTimestamp** | **interface{}** | The time when the transaction was created, expressed in Epoch timestamp using millisecond precision. | 

## Methods

### NewTransfer

`func NewTransfer(id interface{}, profileId interface{}, source InstrumentId, destination InstrumentId, destinationAmount CurrencyAmount, state TransactionState, creationTimestamp interface{}, ) *Transfer`

NewTransfer instantiates a new Transfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferWithDefaults

`func NewTransferWithDefaults() *Transfer`

NewTransferWithDefaults instantiates a new Transfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Transfer) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transfer) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transfer) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *Transfer) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Transfer) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetProfileId

`func (o *Transfer) GetProfileId() interface{}`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *Transfer) GetProfileIdOk() (*interface{}, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *Transfer) SetProfileId(v interface{})`

SetProfileId sets ProfileId field to given value.


### SetProfileIdNil

`func (o *Transfer) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *Transfer) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
### GetTag

`func (o *Transfer) GetTag() interface{}`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Transfer) GetTagOk() (*interface{}, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Transfer) SetTag(v interface{})`

SetTag sets Tag field to given value.

### HasTag

`func (o *Transfer) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *Transfer) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *Transfer) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil
### GetSource

`func (o *Transfer) GetSource() InstrumentId`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Transfer) GetSourceOk() (*InstrumentId, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Transfer) SetSource(v InstrumentId)`

SetSource sets Source field to given value.


### GetDestination

`func (o *Transfer) GetDestination() InstrumentId`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *Transfer) GetDestinationOk() (*InstrumentId, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *Transfer) SetDestination(v InstrumentId)`

SetDestination sets Destination field to given value.


### GetDestinationAmount

`func (o *Transfer) GetDestinationAmount() CurrencyAmount`

GetDestinationAmount returns the DestinationAmount field if non-nil, zero value otherwise.

### GetDestinationAmountOk

`func (o *Transfer) GetDestinationAmountOk() (*CurrencyAmount, bool)`

GetDestinationAmountOk returns a tuple with the DestinationAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAmount

`func (o *Transfer) SetDestinationAmount(v CurrencyAmount)`

SetDestinationAmount sets DestinationAmount field to given value.


### GetState

`func (o *Transfer) GetState() TransactionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Transfer) GetStateOk() (*TransactionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Transfer) SetState(v TransactionState)`

SetState sets State field to given value.


### GetCreationTimestamp

`func (o *Transfer) GetCreationTimestamp() interface{}`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *Transfer) GetCreationTimestampOk() (*interface{}, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *Transfer) SetCreationTimestamp(v interface{})`

SetCreationTimestamp sets CreationTimestamp field to given value.


### SetCreationTimestampNil

`func (o *Transfer) SetCreationTimestampNil(b bool)`

 SetCreationTimestampNil sets the value for CreationTimestamp to be an explicit nil

### UnsetCreationTimestamp
`func (o *Transfer) UnsetCreationTimestamp()`

UnsetCreationTimestamp ensures that no value is present for CreationTimestamp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


