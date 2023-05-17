# PasswordCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | [**SensitivePassword**](SensitivePassword.md) |  | 

## Methods

### NewPasswordCreateRequest

`func NewPasswordCreateRequest(password SensitivePassword, ) *PasswordCreateRequest`

NewPasswordCreateRequest instantiates a new PasswordCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordCreateRequestWithDefaults

`func NewPasswordCreateRequestWithDefaults() *PasswordCreateRequest`

NewPasswordCreateRequestWithDefaults instantiates a new PasswordCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *PasswordCreateRequest) GetPassword() SensitivePassword`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PasswordCreateRequest) GetPasswordOk() (*SensitivePassword, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PasswordCreateRequest) SetPassword(v SensitivePassword)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


