# ManagedAccountUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | Pointer to **interface{}** | The tag field is a custom field that can be used to search and filter. | [optional] 
**FriendlyName** | Pointer to **interface{}** | Updates the friendly name of the managed account. Leave blank if no change is needed. | [optional] 

## Methods

### NewManagedAccountUpdateRequest

`func NewManagedAccountUpdateRequest() *ManagedAccountUpdateRequest`

NewManagedAccountUpdateRequest instantiates a new ManagedAccountUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedAccountUpdateRequestWithDefaults

`func NewManagedAccountUpdateRequestWithDefaults() *ManagedAccountUpdateRequest`

NewManagedAccountUpdateRequestWithDefaults instantiates a new ManagedAccountUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *ManagedAccountUpdateRequest) GetTag() interface{}`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ManagedAccountUpdateRequest) GetTagOk() (*interface{}, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ManagedAccountUpdateRequest) SetTag(v interface{})`

SetTag sets Tag field to given value.

### HasTag

`func (o *ManagedAccountUpdateRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *ManagedAccountUpdateRequest) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *ManagedAccountUpdateRequest) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil
### GetFriendlyName

`func (o *ManagedAccountUpdateRequest) GetFriendlyName() interface{}`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ManagedAccountUpdateRequest) GetFriendlyNameOk() (*interface{}, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ManagedAccountUpdateRequest) SetFriendlyName(v interface{})`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ManagedAccountUpdateRequest) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *ManagedAccountUpdateRequest) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ManagedAccountUpdateRequest) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


