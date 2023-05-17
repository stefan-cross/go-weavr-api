# AuthorisedUserEmailVerifyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **interface{}** | E-mail Address of the user | 
**VerificationCode** | **interface{}** | A randomly generated one-time use code used to verify the user&#39;s email address or mobile number. | 

## Methods

### NewAuthorisedUserEmailVerifyRequest

`func NewAuthorisedUserEmailVerifyRequest(email interface{}, verificationCode interface{}, ) *AuthorisedUserEmailVerifyRequest`

NewAuthorisedUserEmailVerifyRequest instantiates a new AuthorisedUserEmailVerifyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorisedUserEmailVerifyRequestWithDefaults

`func NewAuthorisedUserEmailVerifyRequestWithDefaults() *AuthorisedUserEmailVerifyRequest`

NewAuthorisedUserEmailVerifyRequestWithDefaults instantiates a new AuthorisedUserEmailVerifyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AuthorisedUserEmailVerifyRequest) GetEmail() interface{}`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AuthorisedUserEmailVerifyRequest) GetEmailOk() (*interface{}, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AuthorisedUserEmailVerifyRequest) SetEmail(v interface{})`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *AuthorisedUserEmailVerifyRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *AuthorisedUserEmailVerifyRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetVerificationCode

`func (o *AuthorisedUserEmailVerifyRequest) GetVerificationCode() interface{}`

GetVerificationCode returns the VerificationCode field if non-nil, zero value otherwise.

### GetVerificationCodeOk

`func (o *AuthorisedUserEmailVerifyRequest) GetVerificationCodeOk() (*interface{}, bool)`

GetVerificationCodeOk returns a tuple with the VerificationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationCode

`func (o *AuthorisedUserEmailVerifyRequest) SetVerificationCode(v interface{})`

SetVerificationCode sets VerificationCode field to given value.


### SetVerificationCodeNil

`func (o *AuthorisedUserEmailVerifyRequest) SetVerificationCodeNil(b bool)`

 SetVerificationCodeNil sets the value for VerificationCode to be an explicit nil

### UnsetVerificationCode
`func (o *AuthorisedUserEmailVerifyRequest) UnsetVerificationCode()`

UnsetVerificationCode ensures that no value is present for VerificationCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


