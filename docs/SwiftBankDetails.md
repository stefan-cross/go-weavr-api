# SwiftBankDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iban** | **interface{}** | International Bank Account Number, required for wire transfer over SWIFT. | 
**Code** | **interface{}** | SWIFT code, identifying a particular bank or branch, required for wire transfer over SWIFT. | 

## Methods

### NewSwiftBankDetails

`func NewSwiftBankDetails(iban interface{}, code interface{}, ) *SwiftBankDetails`

NewSwiftBankDetails instantiates a new SwiftBankDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwiftBankDetailsWithDefaults

`func NewSwiftBankDetailsWithDefaults() *SwiftBankDetails`

NewSwiftBankDetailsWithDefaults instantiates a new SwiftBankDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIban

`func (o *SwiftBankDetails) GetIban() interface{}`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *SwiftBankDetails) GetIbanOk() (*interface{}, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *SwiftBankDetails) SetIban(v interface{})`

SetIban sets Iban field to given value.


### SetIbanNil

`func (o *SwiftBankDetails) SetIbanNil(b bool)`

 SetIbanNil sets the value for Iban to be an explicit nil

### UnsetIban
`func (o *SwiftBankDetails) UnsetIban()`

UnsetIban ensures that no value is present for Iban, not even an explicit nil
### GetCode

`func (o *SwiftBankDetails) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SwiftBankDetails) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SwiftBankDetails) SetCode(v interface{})`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *SwiftBankDetails) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *SwiftBankDetails) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


