# TransferCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **interface{}** | The profile Id which a specific identity, instrument or transaction type is linked to.  Profiles contain configuration and determine behavioral aspects of the newly created transaction, for example, fees that may apply.  You can have one or more profiles linked to your application, and these can be used to drive different behaviors according to your product&#39;s needs.  Profile Ids can be found in the Multi Portal, in the API Credentials page.  | 
**Tag** | Pointer to **interface{}** | The tag field is a custom field that can be used to search and filter. | [optional] 
**Source** | [**InstrumentId**](InstrumentId.md) | The managed account or managed card from where the funds will be transferred. | 
**Destination** | [**InstrumentId**](InstrumentId.md) | The managed account or managed card to where the funds will be transferred. | 
**DestinationAmount** | [**CurrencyAmount**](CurrencyAmount.md) | The amount, in same currency as the source and destination instruments, to be transferred to the destination (exclusive of any fee amount that may be specified in the profile). | 

## Methods

### NewTransferCreateRequest

`func NewTransferCreateRequest(profileId interface{}, source InstrumentId, destination InstrumentId, destinationAmount CurrencyAmount, ) *TransferCreateRequest`

NewTransferCreateRequest instantiates a new TransferCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferCreateRequestWithDefaults

`func NewTransferCreateRequestWithDefaults() *TransferCreateRequest`

NewTransferCreateRequestWithDefaults instantiates a new TransferCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *TransferCreateRequest) GetProfileId() interface{}`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *TransferCreateRequest) GetProfileIdOk() (*interface{}, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *TransferCreateRequest) SetProfileId(v interface{})`

SetProfileId sets ProfileId field to given value.


### SetProfileIdNil

`func (o *TransferCreateRequest) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *TransferCreateRequest) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
### GetTag

`func (o *TransferCreateRequest) GetTag() interface{}`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *TransferCreateRequest) GetTagOk() (*interface{}, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *TransferCreateRequest) SetTag(v interface{})`

SetTag sets Tag field to given value.

### HasTag

`func (o *TransferCreateRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *TransferCreateRequest) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *TransferCreateRequest) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil
### GetSource

`func (o *TransferCreateRequest) GetSource() InstrumentId`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TransferCreateRequest) GetSourceOk() (*InstrumentId, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TransferCreateRequest) SetSource(v InstrumentId)`

SetSource sets Source field to given value.


### GetDestination

`func (o *TransferCreateRequest) GetDestination() InstrumentId`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *TransferCreateRequest) GetDestinationOk() (*InstrumentId, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *TransferCreateRequest) SetDestination(v InstrumentId)`

SetDestination sets Destination field to given value.


### GetDestinationAmount

`func (o *TransferCreateRequest) GetDestinationAmount() CurrencyAmount`

GetDestinationAmount returns the DestinationAmount field if non-nil, zero value otherwise.

### GetDestinationAmountOk

`func (o *TransferCreateRequest) GetDestinationAmountOk() (*CurrencyAmount, bool)`

GetDestinationAmountOk returns a tuple with the DestinationAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAmount

`func (o *TransferCreateRequest) SetDestinationAmount(v CurrencyAmount)`

SetDestinationAmount sets DestinationAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


