# RequestAccessTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | [**IdentityId**](IdentityId.md) |  | 

## Methods

### NewRequestAccessTokenRequest

`func NewRequestAccessTokenRequest(identity IdentityId, ) *RequestAccessTokenRequest`

NewRequestAccessTokenRequest instantiates a new RequestAccessTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestAccessTokenRequestWithDefaults

`func NewRequestAccessTokenRequestWithDefaults() *RequestAccessTokenRequest`

NewRequestAccessTokenRequestWithDefaults instantiates a new RequestAccessTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *RequestAccessTokenRequest) GetIdentity() IdentityId`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *RequestAccessTokenRequest) GetIdentityOk() (*IdentityId, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *RequestAccessTokenRequest) SetIdentity(v IdentityId)`

SetIdentity sets Identity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


