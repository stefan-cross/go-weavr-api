# FasterPaymentsBankDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **interface{}** | Account number, required for wire transfer over Faster Payments. | 
**SortCode** | **interface{}** | Sort code, required for wire transfer over Faster Payments. | 

## Methods

### NewFasterPaymentsBankDetails

`func NewFasterPaymentsBankDetails(accountNumber interface{}, sortCode interface{}, ) *FasterPaymentsBankDetails`

NewFasterPaymentsBankDetails instantiates a new FasterPaymentsBankDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFasterPaymentsBankDetailsWithDefaults

`func NewFasterPaymentsBankDetailsWithDefaults() *FasterPaymentsBankDetails`

NewFasterPaymentsBankDetailsWithDefaults instantiates a new FasterPaymentsBankDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *FasterPaymentsBankDetails) GetAccountNumber() interface{}`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *FasterPaymentsBankDetails) GetAccountNumberOk() (*interface{}, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *FasterPaymentsBankDetails) SetAccountNumber(v interface{})`

SetAccountNumber sets AccountNumber field to given value.


### SetAccountNumberNil

`func (o *FasterPaymentsBankDetails) SetAccountNumberNil(b bool)`

 SetAccountNumberNil sets the value for AccountNumber to be an explicit nil

### UnsetAccountNumber
`func (o *FasterPaymentsBankDetails) UnsetAccountNumber()`

UnsetAccountNumber ensures that no value is present for AccountNumber, not even an explicit nil
### GetSortCode

`func (o *FasterPaymentsBankDetails) GetSortCode() interface{}`

GetSortCode returns the SortCode field if non-nil, zero value otherwise.

### GetSortCodeOk

`func (o *FasterPaymentsBankDetails) GetSortCodeOk() (*interface{}, bool)`

GetSortCodeOk returns a tuple with the SortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortCode

`func (o *FasterPaymentsBankDetails) SetSortCode(v interface{})`

SetSortCode sets SortCode field to given value.


### SetSortCodeNil

`func (o *FasterPaymentsBankDetails) SetSortCodeNil(b bool)`

 SetSortCodeNil sets the value for SortCode to be an explicit nil

### UnsetSortCode
`func (o *FasterPaymentsBankDetails) UnsetSortCode()`

UnsetSortCode ensures that no value is present for SortCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


