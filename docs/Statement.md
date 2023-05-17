# Statement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entry** | Pointer to **interface{}** |  | [optional] 
**Count** | Pointer to **interface{}** | The total number of records (excluding the paging limit). | [optional] 
**ResponseCount** | Pointer to **interface{}** | The total number of records returned in this response. | [optional] 

## Methods

### NewStatement

`func NewStatement() *Statement`

NewStatement instantiates a new Statement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementWithDefaults

`func NewStatementWithDefaults() *Statement`

NewStatementWithDefaults instantiates a new Statement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntry

`func (o *Statement) GetEntry() interface{}`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *Statement) GetEntryOk() (*interface{}, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *Statement) SetEntry(v interface{})`

SetEntry sets Entry field to given value.

### HasEntry

`func (o *Statement) HasEntry() bool`

HasEntry returns a boolean if a field has been set.

### SetEntryNil

`func (o *Statement) SetEntryNil(b bool)`

 SetEntryNil sets the value for Entry to be an explicit nil

### UnsetEntry
`func (o *Statement) UnsetEntry()`

UnsetEntry ensures that no value is present for Entry, not even an explicit nil
### GetCount

`func (o *Statement) GetCount() interface{}`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Statement) GetCountOk() (*interface{}, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Statement) SetCount(v interface{})`

SetCount sets Count field to given value.

### HasCount

`func (o *Statement) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *Statement) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *Statement) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetResponseCount

`func (o *Statement) GetResponseCount() interface{}`

GetResponseCount returns the ResponseCount field if non-nil, zero value otherwise.

### GetResponseCountOk

`func (o *Statement) GetResponseCountOk() (*interface{}, bool)`

GetResponseCountOk returns a tuple with the ResponseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCount

`func (o *Statement) SetResponseCount(v interface{})`

SetResponseCount sets ResponseCount field to given value.

### HasResponseCount

`func (o *Statement) HasResponseCount() bool`

HasResponseCount returns a boolean if a field has been set.

### SetResponseCountNil

`func (o *Statement) SetResponseCountNil(b bool)`

 SetResponseCountNil sets the value for ResponseCount to be an explicit nil

### UnsetResponseCount
`func (o *Statement) UnsetResponseCount()`

UnsetResponseCount ensures that no value is present for ResponseCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


