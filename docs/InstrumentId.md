# InstrumentId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **interface{}** |  | 
**Type** | [**InstrumentType**](InstrumentType.md) |  | 

## Methods

### NewInstrumentId

`func NewInstrumentId(id interface{}, type_ InstrumentType, ) *InstrumentId`

NewInstrumentId instantiates a new InstrumentId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstrumentIdWithDefaults

`func NewInstrumentIdWithDefaults() *InstrumentId`

NewInstrumentIdWithDefaults instantiates a new InstrumentId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstrumentId) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstrumentId) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstrumentId) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *InstrumentId) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *InstrumentId) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetType

`func (o *InstrumentId) GetType() InstrumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstrumentId) GetTypeOk() (*InstrumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstrumentId) SetType(v InstrumentType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


