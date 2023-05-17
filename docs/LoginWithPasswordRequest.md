# LoginWithPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **interface{}** | E-mail Address of the user | 
**Password** | [**SensitivePassword**](SensitivePassword.md) |  | 

## Methods

### NewLoginWithPasswordRequest

`func NewLoginWithPasswordRequest(email interface{}, password SensitivePassword, ) *LoginWithPasswordRequest`

NewLoginWithPasswordRequest instantiates a new LoginWithPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginWithPasswordRequestWithDefaults

`func NewLoginWithPasswordRequestWithDefaults() *LoginWithPasswordRequest`

NewLoginWithPasswordRequestWithDefaults instantiates a new LoginWithPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *LoginWithPasswordRequest) GetEmail() interface{}`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LoginWithPasswordRequest) GetEmailOk() (*interface{}, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LoginWithPasswordRequest) SetEmail(v interface{})`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *LoginWithPasswordRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *LoginWithPasswordRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPassword

`func (o *LoginWithPasswordRequest) GetPassword() SensitivePassword`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LoginWithPasswordRequest) GetPasswordOk() (*SensitivePassword, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LoginWithPasswordRequest) SetPassword(v SensitivePassword)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


