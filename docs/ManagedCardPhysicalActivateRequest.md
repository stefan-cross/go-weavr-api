# ManagedCardPhysicalActivateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationCode** | **interface{}** | The code, set up when upgrading the card, that allows the physical card to be activated. | 

## Methods

### NewManagedCardPhysicalActivateRequest

`func NewManagedCardPhysicalActivateRequest(activationCode interface{}, ) *ManagedCardPhysicalActivateRequest`

NewManagedCardPhysicalActivateRequest instantiates a new ManagedCardPhysicalActivateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedCardPhysicalActivateRequestWithDefaults

`func NewManagedCardPhysicalActivateRequestWithDefaults() *ManagedCardPhysicalActivateRequest`

NewManagedCardPhysicalActivateRequestWithDefaults instantiates a new ManagedCardPhysicalActivateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationCode

`func (o *ManagedCardPhysicalActivateRequest) GetActivationCode() interface{}`

GetActivationCode returns the ActivationCode field if non-nil, zero value otherwise.

### GetActivationCodeOk

`func (o *ManagedCardPhysicalActivateRequest) GetActivationCodeOk() (*interface{}, bool)`

GetActivationCodeOk returns a tuple with the ActivationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationCode

`func (o *ManagedCardPhysicalActivateRequest) SetActivationCode(v interface{})`

SetActivationCode sets ActivationCode field to given value.


### SetActivationCodeNil

`func (o *ManagedCardPhysicalActivateRequest) SetActivationCodeNil(b bool)`

 SetActivationCodeNil sets the value for ActivationCode to be an explicit nil

### UnsetActivationCode
`func (o *ManagedCardPhysicalActivateRequest) UnsetActivationCode()`

UnsetActivationCode ensures that no value is present for ActivationCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


