# LostPasswordResumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nonce** | **interface{}** | A randomly generated one-time use code. | 
**Email** | **interface{}** | E-mail Address of the user | 
**NewPassword** | [**SensitivePassword**](SensitivePassword.md) |  | 

## Methods

### NewLostPasswordResumeRequest

`func NewLostPasswordResumeRequest(nonce interface{}, email interface{}, newPassword SensitivePassword, ) *LostPasswordResumeRequest`

NewLostPasswordResumeRequest instantiates a new LostPasswordResumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLostPasswordResumeRequestWithDefaults

`func NewLostPasswordResumeRequestWithDefaults() *LostPasswordResumeRequest`

NewLostPasswordResumeRequestWithDefaults instantiates a new LostPasswordResumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNonce

`func (o *LostPasswordResumeRequest) GetNonce() interface{}`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *LostPasswordResumeRequest) GetNonceOk() (*interface{}, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *LostPasswordResumeRequest) SetNonce(v interface{})`

SetNonce sets Nonce field to given value.


### SetNonceNil

`func (o *LostPasswordResumeRequest) SetNonceNil(b bool)`

 SetNonceNil sets the value for Nonce to be an explicit nil

### UnsetNonce
`func (o *LostPasswordResumeRequest) UnsetNonce()`

UnsetNonce ensures that no value is present for Nonce, not even an explicit nil
### GetEmail

`func (o *LostPasswordResumeRequest) GetEmail() interface{}`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LostPasswordResumeRequest) GetEmailOk() (*interface{}, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LostPasswordResumeRequest) SetEmail(v interface{})`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *LostPasswordResumeRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *LostPasswordResumeRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetNewPassword

`func (o *LostPasswordResumeRequest) GetNewPassword() SensitivePassword`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *LostPasswordResumeRequest) GetNewPasswordOk() (*SensitivePassword, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *LostPasswordResumeRequest) SetNewPassword(v SensitivePassword)`

SetNewPassword sets NewPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


