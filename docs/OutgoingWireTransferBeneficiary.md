# OutgoingWireTransferBeneficiary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | The beneficiary&#39;s full name. | 
**Address** | Pointer to **interface{}** | The beneficiary&#39;s address. | [optional] 
**BankName** | Pointer to **interface{}** | The beneficiary&#39;s bank name. | [optional] 
**BankAddress** | Pointer to **interface{}** | The beneficiary&#39;s bank address. | [optional] 
**BankCountry** | Pointer to **interface{}** | The beneficiary&#39;s bank country. | [optional] 
**BankAccountDetails** | **interface{}** | Details of the beneficiary bank account, depending on the type of transfer chosen. | 

## Methods

### NewOutgoingWireTransferBeneficiary

`func NewOutgoingWireTransferBeneficiary(name interface{}, bankAccountDetails interface{}, ) *OutgoingWireTransferBeneficiary`

NewOutgoingWireTransferBeneficiary instantiates a new OutgoingWireTransferBeneficiary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutgoingWireTransferBeneficiaryWithDefaults

`func NewOutgoingWireTransferBeneficiaryWithDefaults() *OutgoingWireTransferBeneficiary`

NewOutgoingWireTransferBeneficiaryWithDefaults instantiates a new OutgoingWireTransferBeneficiary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OutgoingWireTransferBeneficiary) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OutgoingWireTransferBeneficiary) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OutgoingWireTransferBeneficiary) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *OutgoingWireTransferBeneficiary) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OutgoingWireTransferBeneficiary) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAddress

`func (o *OutgoingWireTransferBeneficiary) GetAddress() interface{}`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *OutgoingWireTransferBeneficiary) GetAddressOk() (*interface{}, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *OutgoingWireTransferBeneficiary) SetAddress(v interface{})`

SetAddress sets Address field to given value.

### HasAddress

`func (o *OutgoingWireTransferBeneficiary) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *OutgoingWireTransferBeneficiary) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *OutgoingWireTransferBeneficiary) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetBankName

`func (o *OutgoingWireTransferBeneficiary) GetBankName() interface{}`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *OutgoingWireTransferBeneficiary) GetBankNameOk() (*interface{}, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *OutgoingWireTransferBeneficiary) SetBankName(v interface{})`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *OutgoingWireTransferBeneficiary) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### SetBankNameNil

`func (o *OutgoingWireTransferBeneficiary) SetBankNameNil(b bool)`

 SetBankNameNil sets the value for BankName to be an explicit nil

### UnsetBankName
`func (o *OutgoingWireTransferBeneficiary) UnsetBankName()`

UnsetBankName ensures that no value is present for BankName, not even an explicit nil
### GetBankAddress

`func (o *OutgoingWireTransferBeneficiary) GetBankAddress() interface{}`

GetBankAddress returns the BankAddress field if non-nil, zero value otherwise.

### GetBankAddressOk

`func (o *OutgoingWireTransferBeneficiary) GetBankAddressOk() (*interface{}, bool)`

GetBankAddressOk returns a tuple with the BankAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAddress

`func (o *OutgoingWireTransferBeneficiary) SetBankAddress(v interface{})`

SetBankAddress sets BankAddress field to given value.

### HasBankAddress

`func (o *OutgoingWireTransferBeneficiary) HasBankAddress() bool`

HasBankAddress returns a boolean if a field has been set.

### SetBankAddressNil

`func (o *OutgoingWireTransferBeneficiary) SetBankAddressNil(b bool)`

 SetBankAddressNil sets the value for BankAddress to be an explicit nil

### UnsetBankAddress
`func (o *OutgoingWireTransferBeneficiary) UnsetBankAddress()`

UnsetBankAddress ensures that no value is present for BankAddress, not even an explicit nil
### GetBankCountry

`func (o *OutgoingWireTransferBeneficiary) GetBankCountry() interface{}`

GetBankCountry returns the BankCountry field if non-nil, zero value otherwise.

### GetBankCountryOk

`func (o *OutgoingWireTransferBeneficiary) GetBankCountryOk() (*interface{}, bool)`

GetBankCountryOk returns a tuple with the BankCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCountry

`func (o *OutgoingWireTransferBeneficiary) SetBankCountry(v interface{})`

SetBankCountry sets BankCountry field to given value.

### HasBankCountry

`func (o *OutgoingWireTransferBeneficiary) HasBankCountry() bool`

HasBankCountry returns a boolean if a field has been set.

### SetBankCountryNil

`func (o *OutgoingWireTransferBeneficiary) SetBankCountryNil(b bool)`

 SetBankCountryNil sets the value for BankCountry to be an explicit nil

### UnsetBankCountry
`func (o *OutgoingWireTransferBeneficiary) UnsetBankCountry()`

UnsetBankCountry ensures that no value is present for BankCountry, not even an explicit nil
### GetBankAccountDetails

`func (o *OutgoingWireTransferBeneficiary) GetBankAccountDetails() interface{}`

GetBankAccountDetails returns the BankAccountDetails field if non-nil, zero value otherwise.

### GetBankAccountDetailsOk

`func (o *OutgoingWireTransferBeneficiary) GetBankAccountDetailsOk() (*interface{}, bool)`

GetBankAccountDetailsOk returns a tuple with the BankAccountDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountDetails

`func (o *OutgoingWireTransferBeneficiary) SetBankAccountDetails(v interface{})`

SetBankAccountDetails sets BankAccountDetails field to given value.


### SetBankAccountDetailsNil

`func (o *OutgoingWireTransferBeneficiary) SetBankAccountDetailsNil(b bool)`

 SetBankAccountDetailsNil sets the value for BankAccountDetails to be an explicit nil

### UnsetBankAccountDetails
`func (o *OutgoingWireTransferBeneficiary) UnsetBankAccountDetails()`

UnsetBankAccountDetails ensures that no value is present for BankAccountDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


