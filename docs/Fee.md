# Fee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeType** | **interface{}** | The fee type as defined in the Multi Portal. | 
**Source** | [**InstrumentId**](InstrumentId.md) | The instrument from where the fee should be deducted. | 

## Methods

### NewFee

`func NewFee(feeType interface{}, source InstrumentId, ) *Fee`

NewFee instantiates a new Fee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeWithDefaults

`func NewFeeWithDefaults() *Fee`

NewFeeWithDefaults instantiates a new Fee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeType

`func (o *Fee) GetFeeType() interface{}`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *Fee) GetFeeTypeOk() (*interface{}, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *Fee) SetFeeType(v interface{})`

SetFeeType sets FeeType field to given value.


### SetFeeTypeNil

`func (o *Fee) SetFeeTypeNil(b bool)`

 SetFeeTypeNil sets the value for FeeType to be an explicit nil

### UnsetFeeType
`func (o *Fee) UnsetFeeType()`

UnsetFeeType ensures that no value is present for FeeType, not even an explicit nil
### GetSource

`func (o *Fee) GetSource() InstrumentId`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Fee) GetSourceOk() (*InstrumentId, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Fee) SetSource(v InstrumentId)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


