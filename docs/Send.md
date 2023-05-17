# Send

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **interface{}** | The unique identifier of a send transaction. | 
**ProfileId** | **interface{}** | The profile Id which a specific identity, instrument or transaction type is linked to.  Profiles contain configuration and determine behavioral aspects of the newly created transaction, for example, fees that may apply.  You can have one or more profiles linked to your application, and these can be used to drive different behaviors according to your product&#39;s needs.  Profile Ids can be found in the Multi Portal, in the API Credentials page.  | 
**Tag** | Pointer to **interface{}** | The tag field is a custom field that can be used to search and filter. | [optional] 
**Source** | [**InstrumentId**](InstrumentId.md) | The source managed account or card from where the funds were sent. | 
**Destination** | [**InstrumentId**](InstrumentId.md) | The destination managed account or card to where the funds were sent. | 
**DestinationAmount** | [**CurrencyAmount**](CurrencyAmount.md) | The amount, in same currency as source and destination, that was sent from the source to the destination instrument. | 
**State** | [**SendState**](SendState.md) |  | 
**CreationTimestamp** | **interface{}** | The time when the transaction was created, expressed in Epoch timestamp using millisecond precision. | 

## Methods

### NewSend

`func NewSend(id interface{}, profileId interface{}, source InstrumentId, destination InstrumentId, destinationAmount CurrencyAmount, state SendState, creationTimestamp interface{}, ) *Send`

NewSend instantiates a new Send object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendWithDefaults

`func NewSendWithDefaults() *Send`

NewSendWithDefaults instantiates a new Send object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Send) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Send) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Send) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *Send) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Send) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetProfileId

`func (o *Send) GetProfileId() interface{}`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *Send) GetProfileIdOk() (*interface{}, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *Send) SetProfileId(v interface{})`

SetProfileId sets ProfileId field to given value.


### SetProfileIdNil

`func (o *Send) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *Send) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
### GetTag

`func (o *Send) GetTag() interface{}`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Send) GetTagOk() (*interface{}, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Send) SetTag(v interface{})`

SetTag sets Tag field to given value.

### HasTag

`func (o *Send) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *Send) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *Send) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil
### GetSource

`func (o *Send) GetSource() InstrumentId`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Send) GetSourceOk() (*InstrumentId, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Send) SetSource(v InstrumentId)`

SetSource sets Source field to given value.


### GetDestination

`func (o *Send) GetDestination() InstrumentId`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *Send) GetDestinationOk() (*InstrumentId, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *Send) SetDestination(v InstrumentId)`

SetDestination sets Destination field to given value.


### GetDestinationAmount

`func (o *Send) GetDestinationAmount() CurrencyAmount`

GetDestinationAmount returns the DestinationAmount field if non-nil, zero value otherwise.

### GetDestinationAmountOk

`func (o *Send) GetDestinationAmountOk() (*CurrencyAmount, bool)`

GetDestinationAmountOk returns a tuple with the DestinationAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAmount

`func (o *Send) SetDestinationAmount(v CurrencyAmount)`

SetDestinationAmount sets DestinationAmount field to given value.


### GetState

`func (o *Send) GetState() SendState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Send) GetStateOk() (*SendState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Send) SetState(v SendState)`

SetState sets State field to given value.


### GetCreationTimestamp

`func (o *Send) GetCreationTimestamp() interface{}`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *Send) GetCreationTimestampOk() (*interface{}, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *Send) SetCreationTimestamp(v interface{})`

SetCreationTimestamp sets CreationTimestamp field to given value.


### SetCreationTimestampNil

`func (o *Send) SetCreationTimestampNil(b bool)`

 SetCreationTimestampNil sets the value for CreationTimestamp to be an explicit nil

### UnsetCreationTimestamp
`func (o *Send) UnsetCreationTimestamp()`

UnsetCreationTimestamp ensures that no value is present for CreationTimestamp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


