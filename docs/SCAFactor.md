# SCAFactor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**SCAChallengeType**](SCAChallengeType.md) |  | [optional] 
**Status** | Pointer to [**SCAFactorStatus**](SCAFactorStatus.md) |  | [optional] 
**Channel** | Pointer to [**SCAChannel**](SCAChannel.md) |  | [optional] 

## Methods

### NewSCAFactor

`func NewSCAFactor() *SCAFactor`

NewSCAFactor instantiates a new SCAFactor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSCAFactorWithDefaults

`func NewSCAFactorWithDefaults() *SCAFactor`

NewSCAFactorWithDefaults instantiates a new SCAFactor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SCAFactor) GetType() SCAChallengeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SCAFactor) GetTypeOk() (*SCAChallengeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SCAFactor) SetType(v SCAChallengeType)`

SetType sets Type field to given value.

### HasType

`func (o *SCAFactor) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *SCAFactor) GetStatus() SCAFactorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SCAFactor) GetStatusOk() (*SCAFactorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SCAFactor) SetStatus(v SCAFactorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SCAFactor) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetChannel

`func (o *SCAFactor) GetChannel() SCAChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *SCAFactor) GetChannelOk() (*SCAChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *SCAFactor) SetChannel(v SCAChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *SCAFactor) HasChannel() bool`

HasChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


