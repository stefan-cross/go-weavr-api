# UserCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | The first name of the user. | 
**Surname** | **interface{}** | The last name of the user. | 
**Email** | **interface{}** | E-mail Address of the user | 
**Mobile** | Pointer to [**Mobile**](Mobile.md) |  | [optional] 
**DateOfBirth** | Pointer to [**Date**](Date.md) | Date of birth of the authorised user. | [optional] 

## Methods

### NewUserCreateRequest

`func NewUserCreateRequest(name interface{}, surname interface{}, email interface{}, ) *UserCreateRequest`

NewUserCreateRequest instantiates a new UserCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreateRequestWithDefaults

`func NewUserCreateRequestWithDefaults() *UserCreateRequest`

NewUserCreateRequestWithDefaults instantiates a new UserCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserCreateRequest) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserCreateRequest) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserCreateRequest) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *UserCreateRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UserCreateRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSurname

`func (o *UserCreateRequest) GetSurname() interface{}`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *UserCreateRequest) GetSurnameOk() (*interface{}, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *UserCreateRequest) SetSurname(v interface{})`

SetSurname sets Surname field to given value.


### SetSurnameNil

`func (o *UserCreateRequest) SetSurnameNil(b bool)`

 SetSurnameNil sets the value for Surname to be an explicit nil

### UnsetSurname
`func (o *UserCreateRequest) UnsetSurname()`

UnsetSurname ensures that no value is present for Surname, not even an explicit nil
### GetEmail

`func (o *UserCreateRequest) GetEmail() interface{}`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserCreateRequest) GetEmailOk() (*interface{}, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserCreateRequest) SetEmail(v interface{})`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *UserCreateRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UserCreateRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetMobile

`func (o *UserCreateRequest) GetMobile() Mobile`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *UserCreateRequest) GetMobileOk() (*Mobile, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *UserCreateRequest) SetMobile(v Mobile)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *UserCreateRequest) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *UserCreateRequest) GetDateOfBirth() Date`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *UserCreateRequest) GetDateOfBirthOk() (*Date, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *UserCreateRequest) SetDateOfBirth(v Date)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *UserCreateRequest) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


