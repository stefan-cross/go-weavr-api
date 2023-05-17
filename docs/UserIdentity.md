# UserIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**IdentityId**](IdentityId.md) | The unique identifier of the Corporate Identity. | 
**Name** | **interface{}** | The name of the company. | 

## Methods

### NewUserIdentity

`func NewUserIdentity(id IdentityId, name interface{}, ) *UserIdentity`

NewUserIdentity instantiates a new UserIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserIdentityWithDefaults

`func NewUserIdentityWithDefaults() *UserIdentity`

NewUserIdentityWithDefaults instantiates a new UserIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserIdentity) GetId() IdentityId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserIdentity) GetIdOk() (*IdentityId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserIdentity) SetId(v IdentityId)`

SetId sets Id field to given value.


### GetName

`func (o *UserIdentity) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserIdentity) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserIdentity) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *UserIdentity) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UserIdentity) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


