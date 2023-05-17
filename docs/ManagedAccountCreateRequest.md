# ManagedAccountCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **interface{}** | The profile Id which a specific identity, instrument or transaction type is linked to.  Profiles contain configuration and determine behavioral aspects of the newly created transaction, for example, fees that may apply.  You can have one or more profiles linked to your application, and these can be used to drive different behaviors according to your product&#39;s needs.  Profile Ids can be found in the Multi Portal, in the API Credentials page.  | 
**FriendlyName** | **interface{}** | The friendly name to be given to the managed account. | 
**Currency** | **interface{}** | The currency expressed in ISO-4217 code. Example: GBP, EUR, USD. | 
**Tag** | Pointer to **interface{}** | The tag field is a custom field that can be used to search and filter. | [optional] 

## Methods

### NewManagedAccountCreateRequest

`func NewManagedAccountCreateRequest(profileId interface{}, friendlyName interface{}, currency interface{}, ) *ManagedAccountCreateRequest`

NewManagedAccountCreateRequest instantiates a new ManagedAccountCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedAccountCreateRequestWithDefaults

`func NewManagedAccountCreateRequestWithDefaults() *ManagedAccountCreateRequest`

NewManagedAccountCreateRequestWithDefaults instantiates a new ManagedAccountCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *ManagedAccountCreateRequest) GetProfileId() interface{}`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *ManagedAccountCreateRequest) GetProfileIdOk() (*interface{}, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *ManagedAccountCreateRequest) SetProfileId(v interface{})`

SetProfileId sets ProfileId field to given value.


### SetProfileIdNil

`func (o *ManagedAccountCreateRequest) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *ManagedAccountCreateRequest) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
### GetFriendlyName

`func (o *ManagedAccountCreateRequest) GetFriendlyName() interface{}`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ManagedAccountCreateRequest) GetFriendlyNameOk() (*interface{}, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ManagedAccountCreateRequest) SetFriendlyName(v interface{})`

SetFriendlyName sets FriendlyName field to given value.


### SetFriendlyNameNil

`func (o *ManagedAccountCreateRequest) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ManagedAccountCreateRequest) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetCurrency

`func (o *ManagedAccountCreateRequest) GetCurrency() interface{}`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ManagedAccountCreateRequest) GetCurrencyOk() (*interface{}, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ManagedAccountCreateRequest) SetCurrency(v interface{})`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *ManagedAccountCreateRequest) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *ManagedAccountCreateRequest) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetTag

`func (o *ManagedAccountCreateRequest) GetTag() interface{}`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ManagedAccountCreateRequest) GetTagOk() (*interface{}, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ManagedAccountCreateRequest) SetTag(v interface{})`

SetTag sets Tag field to given value.

### HasTag

`func (o *ManagedAccountCreateRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *ManagedAccountCreateRequest) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *ManagedAccountCreateRequest) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


