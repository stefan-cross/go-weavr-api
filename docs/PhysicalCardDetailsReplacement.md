# PhysicalCardDetailsReplacement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReplacementReason** | Pointer to **interface{}** | The reason why the physical card was replaced.   - DAMAGED: The physical card was damaged and cannot be used at a physical terminal.   - LOST_STOLEN: The physical card was either lost or stolen and cannot be used.   - EXPIRED: The physical card expired.  | [optional] 
**ReplacementId** | Pointer to **interface{}** | The unique identifier of the new card that replaces this card. | [optional] 

## Methods

### NewPhysicalCardDetailsReplacement

`func NewPhysicalCardDetailsReplacement() *PhysicalCardDetailsReplacement`

NewPhysicalCardDetailsReplacement instantiates a new PhysicalCardDetailsReplacement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalCardDetailsReplacementWithDefaults

`func NewPhysicalCardDetailsReplacementWithDefaults() *PhysicalCardDetailsReplacement`

NewPhysicalCardDetailsReplacementWithDefaults instantiates a new PhysicalCardDetailsReplacement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplacementReason

`func (o *PhysicalCardDetailsReplacement) GetReplacementReason() interface{}`

GetReplacementReason returns the ReplacementReason field if non-nil, zero value otherwise.

### GetReplacementReasonOk

`func (o *PhysicalCardDetailsReplacement) GetReplacementReasonOk() (*interface{}, bool)`

GetReplacementReasonOk returns a tuple with the ReplacementReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementReason

`func (o *PhysicalCardDetailsReplacement) SetReplacementReason(v interface{})`

SetReplacementReason sets ReplacementReason field to given value.

### HasReplacementReason

`func (o *PhysicalCardDetailsReplacement) HasReplacementReason() bool`

HasReplacementReason returns a boolean if a field has been set.

### SetReplacementReasonNil

`func (o *PhysicalCardDetailsReplacement) SetReplacementReasonNil(b bool)`

 SetReplacementReasonNil sets the value for ReplacementReason to be an explicit nil

### UnsetReplacementReason
`func (o *PhysicalCardDetailsReplacement) UnsetReplacementReason()`

UnsetReplacementReason ensures that no value is present for ReplacementReason, not even an explicit nil
### GetReplacementId

`func (o *PhysicalCardDetailsReplacement) GetReplacementId() interface{}`

GetReplacementId returns the ReplacementId field if non-nil, zero value otherwise.

### GetReplacementIdOk

`func (o *PhysicalCardDetailsReplacement) GetReplacementIdOk() (*interface{}, bool)`

GetReplacementIdOk returns a tuple with the ReplacementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementId

`func (o *PhysicalCardDetailsReplacement) SetReplacementId(v interface{})`

SetReplacementId sets ReplacementId field to given value.

### HasReplacementId

`func (o *PhysicalCardDetailsReplacement) HasReplacementId() bool`

HasReplacementId returns a boolean if a field has been set.

### SetReplacementIdNil

`func (o *PhysicalCardDetailsReplacement) SetReplacementIdNil(b bool)`

 SetReplacementIdNil sets the value for ReplacementId to be an explicit nil

### UnsetReplacementId
`func (o *PhysicalCardDetailsReplacement) UnsetReplacementId()`

UnsetReplacementId ensures that no value is present for ReplacementId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


