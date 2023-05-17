# ManagedAccountIBAN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **interface{}** | The state of the Managed Account&#39;s IBAN as follows:   - UNALLOCATED: The Managed Account has never been assigned an IBAN. Use the _managedAccountsIBANUpgrade_ operation to assign an IBAN to a Managed Account.   - PENDING_ALLOCATION: The IBAN is being allocated to the Managed Account.   - ALLOCATED: An IBAN is allocated to the Managed Account.  | 
**BankAccountDetails** | **interface{}** | A list of bank account details associated with the IBAN. Multiple details can be provided if multiple IBAN providers are supported by your payment model. | 

## Methods

### NewManagedAccountIBAN

`func NewManagedAccountIBAN(state interface{}, bankAccountDetails interface{}, ) *ManagedAccountIBAN`

NewManagedAccountIBAN instantiates a new ManagedAccountIBAN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedAccountIBANWithDefaults

`func NewManagedAccountIBANWithDefaults() *ManagedAccountIBAN`

NewManagedAccountIBANWithDefaults instantiates a new ManagedAccountIBAN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ManagedAccountIBAN) GetState() interface{}`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ManagedAccountIBAN) GetStateOk() (*interface{}, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ManagedAccountIBAN) SetState(v interface{})`

SetState sets State field to given value.


### SetStateNil

`func (o *ManagedAccountIBAN) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *ManagedAccountIBAN) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetBankAccountDetails

`func (o *ManagedAccountIBAN) GetBankAccountDetails() interface{}`

GetBankAccountDetails returns the BankAccountDetails field if non-nil, zero value otherwise.

### GetBankAccountDetailsOk

`func (o *ManagedAccountIBAN) GetBankAccountDetailsOk() (*interface{}, bool)`

GetBankAccountDetailsOk returns a tuple with the BankAccountDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountDetails

`func (o *ManagedAccountIBAN) SetBankAccountDetails(v interface{})`

SetBankAccountDetails sets BankAccountDetails field to given value.


### SetBankAccountDetailsNil

`func (o *ManagedAccountIBAN) SetBankAccountDetailsNil(b bool)`

 SetBankAccountDetailsNil sets the value for BankAccountDetails to be an explicit nil

### UnsetBankAccountDetails
`func (o *ManagedAccountIBAN) UnsetBankAccountDetails()`

UnsetBankAccountDetails ensures that no value is present for BankAccountDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


