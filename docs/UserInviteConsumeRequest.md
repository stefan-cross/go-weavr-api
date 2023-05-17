# UserInviteConsumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InviteCode** | **interface{}** | A randomly generated one-time use code. | 
**Password** | [**SensitivePassword**](SensitivePassword.md) |  | 

## Methods

### NewUserInviteConsumeRequest

`func NewUserInviteConsumeRequest(inviteCode interface{}, password SensitivePassword, ) *UserInviteConsumeRequest`

NewUserInviteConsumeRequest instantiates a new UserInviteConsumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInviteConsumeRequestWithDefaults

`func NewUserInviteConsumeRequestWithDefaults() *UserInviteConsumeRequest`

NewUserInviteConsumeRequestWithDefaults instantiates a new UserInviteConsumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInviteCode

`func (o *UserInviteConsumeRequest) GetInviteCode() interface{}`

GetInviteCode returns the InviteCode field if non-nil, zero value otherwise.

### GetInviteCodeOk

`func (o *UserInviteConsumeRequest) GetInviteCodeOk() (*interface{}, bool)`

GetInviteCodeOk returns a tuple with the InviteCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteCode

`func (o *UserInviteConsumeRequest) SetInviteCode(v interface{})`

SetInviteCode sets InviteCode field to given value.


### SetInviteCodeNil

`func (o *UserInviteConsumeRequest) SetInviteCodeNil(b bool)`

 SetInviteCodeNil sets the value for InviteCode to be an explicit nil

### UnsetInviteCode
`func (o *UserInviteConsumeRequest) UnsetInviteCode()`

UnsetInviteCode ensures that no value is present for InviteCode, not even an explicit nil
### GetPassword

`func (o *UserInviteConsumeRequest) GetPassword() SensitivePassword`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserInviteConsumeRequest) GetPasswordOk() (*SensitivePassword, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserInviteConsumeRequest) SetPassword(v SensitivePassword)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


