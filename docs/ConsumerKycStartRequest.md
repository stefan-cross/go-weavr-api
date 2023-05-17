# ConsumerKycStartRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KycLevel** | Pointer to [**KycLevel**](KycLevel.md) | The KYC level that the consumer will be assigned to, which determines the due diligence details that the user will need to provide. | [optional] 

## Methods

### NewConsumerKycStartRequest

`func NewConsumerKycStartRequest() *ConsumerKycStartRequest`

NewConsumerKycStartRequest instantiates a new ConsumerKycStartRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerKycStartRequestWithDefaults

`func NewConsumerKycStartRequestWithDefaults() *ConsumerKycStartRequest`

NewConsumerKycStartRequestWithDefaults instantiates a new ConsumerKycStartRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKycLevel

`func (o *ConsumerKycStartRequest) GetKycLevel() KycLevel`

GetKycLevel returns the KycLevel field if non-nil, zero value otherwise.

### GetKycLevelOk

`func (o *ConsumerKycStartRequest) GetKycLevelOk() (*KycLevel, bool)`

GetKycLevelOk returns a tuple with the KycLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycLevel

`func (o *ConsumerKycStartRequest) SetKycLevel(v KycLevel)`

SetKycLevel sets KycLevel field to given value.

### HasKycLevel

`func (o *ConsumerKycStartRequest) HasKycLevel() bool`

HasKycLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


