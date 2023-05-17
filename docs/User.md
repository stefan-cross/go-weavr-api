# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **interface{}** | The unique identifier of the user. | 
**Identity** | [**IdentityId**](IdentityId.md) | The identity that the user belongs to. | 
**Name** | **interface{}** | The first name of the user. | 
**Surname** | **interface{}** | The last name of the user. | 
**Email** | **interface{}** | E-mail Address of the user | 
**Mobile** | Pointer to [**Mobile**](Mobile.md) |  | [optional] 
**Active** | **interface{}** | The state of the user. If the &#x60;active&#x60; attribute is false, then the user will not be able to log in. | 
**DateOfBirth** | Pointer to [**Date**](Date.md) | Date of birth of the authorised user. | [optional] 

## Methods

### NewUser

`func NewUser(id interface{}, identity IdentityId, name interface{}, surname interface{}, email interface{}, active interface{}, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *User) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *User) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIdentity

`func (o *User) GetIdentity() IdentityId`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *User) GetIdentityOk() (*IdentityId, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *User) SetIdentity(v IdentityId)`

SetIdentity sets Identity field to given value.


### GetName

`func (o *User) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *User) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *User) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSurname

`func (o *User) GetSurname() interface{}`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *User) GetSurnameOk() (*interface{}, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *User) SetSurname(v interface{})`

SetSurname sets Surname field to given value.


### SetSurnameNil

`func (o *User) SetSurnameNil(b bool)`

 SetSurnameNil sets the value for Surname to be an explicit nil

### UnsetSurname
`func (o *User) UnsetSurname()`

UnsetSurname ensures that no value is present for Surname, not even an explicit nil
### GetEmail

`func (o *User) GetEmail() interface{}`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*interface{}, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v interface{})`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *User) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *User) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetMobile

`func (o *User) GetMobile() Mobile`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *User) GetMobileOk() (*Mobile, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *User) SetMobile(v Mobile)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *User) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetActive

`func (o *User) GetActive() interface{}`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *User) GetActiveOk() (*interface{}, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *User) SetActive(v interface{})`

SetActive sets Active field to given value.


### SetActiveNil

`func (o *User) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *User) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetDateOfBirth

`func (o *User) GetDateOfBirth() Date`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *User) GetDateOfBirthOk() (*Date, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *User) SetDateOfBirth(v Date)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *User) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


