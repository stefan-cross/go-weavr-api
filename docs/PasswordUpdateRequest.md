# PasswordUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OldPassword** | [**SensitivePassword**](SensitivePassword.md) |  | 
**NewPassword** | [**SensitivePassword**](SensitivePassword.md) |  | 

## Methods

### NewPasswordUpdateRequest

`func NewPasswordUpdateRequest(oldPassword SensitivePassword, newPassword SensitivePassword, ) *PasswordUpdateRequest`

NewPasswordUpdateRequest instantiates a new PasswordUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordUpdateRequestWithDefaults

`func NewPasswordUpdateRequestWithDefaults() *PasswordUpdateRequest`

NewPasswordUpdateRequestWithDefaults instantiates a new PasswordUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOldPassword

`func (o *PasswordUpdateRequest) GetOldPassword() SensitivePassword`

GetOldPassword returns the OldPassword field if non-nil, zero value otherwise.

### GetOldPasswordOk

`func (o *PasswordUpdateRequest) GetOldPasswordOk() (*SensitivePassword, bool)`

GetOldPasswordOk returns a tuple with the OldPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPassword

`func (o *PasswordUpdateRequest) SetOldPassword(v SensitivePassword)`

SetOldPassword sets OldPassword field to given value.


### GetNewPassword

`func (o *PasswordUpdateRequest) GetNewPassword() SensitivePassword`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *PasswordUpdateRequest) GetNewPasswordOk() (*SensitivePassword, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *PasswordUpdateRequest) SetNewPassword(v SensitivePassword)`

SetNewPassword sets NewPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


