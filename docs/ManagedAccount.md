# ManagedAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **interface{}** | The unique identifier of a Managed Account. | 
**ProfileId** | **interface{}** | The profile Id which a specific identity, instrument or transaction type is linked to.  Profiles contain configuration and determine behavioral aspects of the newly created transaction, for example, fees that may apply.  You can have one or more profiles linked to your application, and these can be used to drive different behaviors according to your product&#39;s needs.  Profile Ids can be found in the Multi Portal, in the API Credentials page.  | 
**Tag** | Pointer to **interface{}** | The tag field is a custom field that can be used to search and filter. | [optional] 
**FriendlyName** | **interface{}** | The friendly name given to the managed account. | 
**Currency** | **interface{}** | The currency expressed in ISO-4217 code. Example: GBP, EUR, USD. | 
**Balances** | [**ManagedInstrumentBalance**](ManagedInstrumentBalance.md) |  | 
**State** | [**ManagedInstrumentState**](ManagedInstrumentState.md) |  | 
**CreationTimestamp** | **interface{}** | The time when the managed account was created, expressed in Epoch timestamp using millisecond precision. | 

## Methods

### NewManagedAccount

`func NewManagedAccount(id interface{}, profileId interface{}, friendlyName interface{}, currency interface{}, balances ManagedInstrumentBalance, state ManagedInstrumentState, creationTimestamp interface{}, ) *ManagedAccount`

NewManagedAccount instantiates a new ManagedAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedAccountWithDefaults

`func NewManagedAccountWithDefaults() *ManagedAccount`

NewManagedAccountWithDefaults instantiates a new ManagedAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManagedAccount) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedAccount) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedAccount) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *ManagedAccount) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ManagedAccount) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetProfileId

`func (o *ManagedAccount) GetProfileId() interface{}`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *ManagedAccount) GetProfileIdOk() (*interface{}, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *ManagedAccount) SetProfileId(v interface{})`

SetProfileId sets ProfileId field to given value.


### SetProfileIdNil

`func (o *ManagedAccount) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *ManagedAccount) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
### GetTag

`func (o *ManagedAccount) GetTag() interface{}`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ManagedAccount) GetTagOk() (*interface{}, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ManagedAccount) SetTag(v interface{})`

SetTag sets Tag field to given value.

### HasTag

`func (o *ManagedAccount) HasTag() bool`

HasTag returns a boolean if a field has been set.

### SetTagNil

`func (o *ManagedAccount) SetTagNil(b bool)`

 SetTagNil sets the value for Tag to be an explicit nil

### UnsetTag
`func (o *ManagedAccount) UnsetTag()`

UnsetTag ensures that no value is present for Tag, not even an explicit nil
### GetFriendlyName

`func (o *ManagedAccount) GetFriendlyName() interface{}`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ManagedAccount) GetFriendlyNameOk() (*interface{}, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ManagedAccount) SetFriendlyName(v interface{})`

SetFriendlyName sets FriendlyName field to given value.


### SetFriendlyNameNil

`func (o *ManagedAccount) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ManagedAccount) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetCurrency

`func (o *ManagedAccount) GetCurrency() interface{}`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ManagedAccount) GetCurrencyOk() (*interface{}, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ManagedAccount) SetCurrency(v interface{})`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *ManagedAccount) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *ManagedAccount) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetBalances

`func (o *ManagedAccount) GetBalances() ManagedInstrumentBalance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *ManagedAccount) GetBalancesOk() (*ManagedInstrumentBalance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *ManagedAccount) SetBalances(v ManagedInstrumentBalance)`

SetBalances sets Balances field to given value.


### GetState

`func (o *ManagedAccount) GetState() ManagedInstrumentState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ManagedAccount) GetStateOk() (*ManagedInstrumentState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ManagedAccount) SetState(v ManagedInstrumentState)`

SetState sets State field to given value.


### GetCreationTimestamp

`func (o *ManagedAccount) GetCreationTimestamp() interface{}`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *ManagedAccount) GetCreationTimestampOk() (*interface{}, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *ManagedAccount) SetCreationTimestamp(v interface{})`

SetCreationTimestamp sets CreationTimestamp field to given value.


### SetCreationTimestampNil

`func (o *ManagedAccount) SetCreationTimestampNil(b bool)`

 SetCreationTimestampNil sets the value for CreationTimestamp to be an explicit nil

### UnsetCreationTimestamp
`func (o *ManagedAccount) UnsetCreationTimestamp()`

UnsetCreationTimestamp ensures that no value is present for CreationTimestamp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


