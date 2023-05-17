# SepaBankDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iban** | **interface{}** | International Bank Account Number, required for wire transfer over SEPA. | 
**BankIdentifierCode** | **interface{}** | BIC, required for wire transfer over SEPA. | 

## Methods

### NewSepaBankDetails

`func NewSepaBankDetails(iban interface{}, bankIdentifierCode interface{}, ) *SepaBankDetails`

NewSepaBankDetails instantiates a new SepaBankDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSepaBankDetailsWithDefaults

`func NewSepaBankDetailsWithDefaults() *SepaBankDetails`

NewSepaBankDetailsWithDefaults instantiates a new SepaBankDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIban

`func (o *SepaBankDetails) GetIban() interface{}`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *SepaBankDetails) GetIbanOk() (*interface{}, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *SepaBankDetails) SetIban(v interface{})`

SetIban sets Iban field to given value.


### SetIbanNil

`func (o *SepaBankDetails) SetIbanNil(b bool)`

 SetIbanNil sets the value for Iban to be an explicit nil

### UnsetIban
`func (o *SepaBankDetails) UnsetIban()`

UnsetIban ensures that no value is present for Iban, not even an explicit nil
### GetBankIdentifierCode

`func (o *SepaBankDetails) GetBankIdentifierCode() interface{}`

GetBankIdentifierCode returns the BankIdentifierCode field if non-nil, zero value otherwise.

### GetBankIdentifierCodeOk

`func (o *SepaBankDetails) GetBankIdentifierCodeOk() (*interface{}, bool)`

GetBankIdentifierCodeOk returns a tuple with the BankIdentifierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankIdentifierCode

`func (o *SepaBankDetails) SetBankIdentifierCode(v interface{})`

SetBankIdentifierCode sets BankIdentifierCode field to given value.


### SetBankIdentifierCodeNil

`func (o *SepaBankDetails) SetBankIdentifierCodeNil(b bool)`

 SetBankIdentifierCodeNil sets the value for BankIdentifierCode to be an explicit nil

### UnsetBankIdentifierCode
`func (o *SepaBankDetails) UnsetBankIdentifierCode()`

UnsetBankIdentifierCode ensures that no value is present for BankIdentifierCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


