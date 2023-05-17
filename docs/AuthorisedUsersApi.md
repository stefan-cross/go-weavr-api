# \AuthorisedUsersApi

All URIs are relative to *https://sandbox.weavr.io/multi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorisedUserEmailVerificationCodeSend**](AuthorisedUsersApi.md#AuthorisedUserEmailVerificationCodeSend) | **Post** /users/verification/email/send | Send an email verification code to the authorised user
[**AuthorisedUserEmailVerify**](AuthorisedUsersApi.md#AuthorisedUserEmailVerify) | **Post** /users/verification/email/verify | Verify email of the authorised user
[**UserActivate**](AuthorisedUsersApi.md#UserActivate) | **Post** /users/{user_id}/activate | Activate a user
[**UserCreate**](AuthorisedUsersApi.md#UserCreate) | **Post** /users | Create a user
[**UserDeactivate**](AuthorisedUsersApi.md#UserDeactivate) | **Post** /users/{user_id}/deactivate | Deactivate a user
[**UserGetById**](AuthorisedUsersApi.md#UserGetById) | **Get** /users/{user_id} | Get a user
[**UserInviteConsume**](AuthorisedUsersApi.md#UserInviteConsume) | **Post** /users/{user_id}/invite/consume | Consume a user invite
[**UserInviteSend**](AuthorisedUsersApi.md#UserInviteSend) | **Post** /users/{user_id}/invite | Send a user invite
[**UserInviteValidate**](AuthorisedUsersApi.md#UserInviteValidate) | **Post** /users/{user_id}/invite/validate | Validate a user invite
[**UserKyc**](AuthorisedUsersApi.md#UserKyc) | **Post** /users/kyc | Start KYC for the user
[**UserUpdate**](AuthorisedUsersApi.md#UserUpdate) | **Patch** /users/{user_id} | Update a user
[**UsersGet**](AuthorisedUsersApi.md#UsersGet) | **Get** /users | Get all users



## AuthorisedUserEmailVerificationCodeSend

> AuthorisedUserEmailVerificationCodeSend(ctx).LostPasswordInitiateRequest(lostPasswordInitiateRequest).Execute()

Send an email verification code to the authorised user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stefan-cross/go-weavr-api"
)

func main() {
    lostPasswordInitiateRequest := *openapiclient.NewLostPasswordInitiateRequest(interface{}(123)) // LostPasswordInitiateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorisedUsersApi.AuthorisedUserEmailVerificationCodeSend(context.Background()).LostPasswordInitiateRequest(lostPasswordInitiateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorisedUsersApi.AuthorisedUserEmailVerificationCodeSend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorisedUserEmailVerificationCodeSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lostPasswordInitiateRequest** | [**LostPasswordInitiateRequest**](LostPasswordInitiateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorisedUserEmailVerify

> AuthorisedUserEmailVerify(ctx).AuthorisedUserEmailVerifyRequest(authorisedUserEmailVerifyRequest).Execute()

Verify email of the authorised user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stefan-cross/go-weavr-api"
)

func main() {
    authorisedUserEmailVerifyRequest := *openapiclient.NewAuthorisedUserEmailVerifyRequest(interface{}(123), interface{}(123)) // AuthorisedUserEmailVerifyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorisedUsersApi.AuthorisedUserEmailVerify(context.Background()).AuthorisedUserEmailVerifyRequest(authorisedUserEmailVerifyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorisedUsersApi.AuthorisedUserEmailVerify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorisedUserEmailVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorisedUserEmailVerifyRequest** | [**AuthorisedUserEmailVerifyRequest**](AuthorisedUserEmailVerifyRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserActivate

> UserActivate(ctx, userId).Execute()

Activate a user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stefan-cross/go-weavr-api"
)

func main() {
    userId := TODO // interface{} | The user id for which this password is created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorisedUsersApi.UserActivate(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorisedUsersApi.UserActivate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**interface{}**](.md) | The user id for which this password is created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserActivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [auth_token](../README.md#auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserCreate

> UserCreate(ctx).UserCreateRequest(userCreateRequest).Execute()

Create a user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stefan-cross/go-weavr-api"
)

func main() {
    userCreateRequest := *openapiclient.NewUserCreateRequest(interface{}(123), interface{}(123), interface{}(123)) // UserCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorisedUsersApi.UserCreate(context.Background()).UserCreateRequest(userCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorisedUsersApi.UserCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userCreateRequest** | [**UserCreateRequest**](UserCreateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [auth_token](../README.md#auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserDeactivate

> UserDeactivate(ctx, userId).Execute()

Deactivate a user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stefan-cross/go-weavr-api"
)

func main() {
    userId := TODO // interface{} | The user id for which this password is created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorisedUsersApi.UserDeactivate(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorisedUsersApi.UserDeactivate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**interface{}**](.md) | The user id for which this password is created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserDeactivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [auth_token](../README.md#auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGetById

> UserGetById(ctx, userId).Execute()

Get a user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stefan-cross/go-weavr-api"
)

func main() {
    userId := TODO // interface{} | The user id for which this password is created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorisedUsersApi.UserGetById(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorisedUsersApi.UserGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**interface{}**](.md) | The user id for which this password is created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [auth_token](../README.md#auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserInviteConsume

> UserInviteConsume(ctx, userId).UserInviteConsumeRequest(userInviteConsumeRequest).Execute()

Consume a user invite



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stefan-cross/go-weavr-api"
)

func main() {
    userId := TODO // interface{} | The unique identifier for the user.
    userInviteConsumeRequest := *openapiclient.NewUserInviteConsumeRequest(interface{}(123), *openapiclient.NewSensitivePassword(interface{}(123))) // UserInviteConsumeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorisedUsersApi.UserInviteConsume(context.Background(), userId).UserInviteConsumeRequest(userInviteConsumeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorisedUsersApi.UserInviteConsume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**interface{}**](.md) | The unique identifier for the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserInviteConsumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userInviteConsumeRequest** | [**UserInviteConsumeRequest**](UserInviteConsumeRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserInviteSend

> UserInviteSend(ctx, userId).Execute()

Send a user invite



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stefan-cross/go-weavr-api"
)

func main() {
    userId := TODO // interface{} | The unique identifier for the user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorisedUsersApi.UserInviteSend(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorisedUsersApi.UserInviteSend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**interface{}**](.md) | The unique identifier for the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserInviteSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [auth_token](../README.md#auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserInviteValidate

> UserInviteValidate(ctx, userId).UserInviteValidateRequest(userInviteValidateRequest).Execute()

Validate a user invite



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stefan-cross/go-weavr-api"
)

func main() {
    userId := TODO // interface{} | The unique identifier for the user.
    userInviteValidateRequest := *openapiclient.NewUserInviteValidateRequest(interface{}(123)) // UserInviteValidateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorisedUsersApi.UserInviteValidate(context.Background(), userId).UserInviteValidateRequest(userInviteValidateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorisedUsersApi.UserInviteValidate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**interface{}**](.md) | The unique identifier for the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserInviteValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userInviteValidateRequest** | [**UserInviteValidateRequest**](UserInviteValidateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserKyc

> UserKyc200Response UserKyc(ctx).Execute()

Start KYC for the user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stefan-cross/go-weavr-api"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorisedUsersApi.UserKyc(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorisedUsersApi.UserKyc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserKyc`: UserKyc200Response
    fmt.Fprintf(os.Stdout, "Response from `AuthorisedUsersApi.UserKyc`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserKycRequest struct via the builder pattern


### Return type

[**UserKyc200Response**](UserKyc200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserUpdate

> UserUpdate(ctx, userId).UserUpdateRequest(userUpdateRequest).Execute()

Update a user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stefan-cross/go-weavr-api"
)

func main() {
    userId := TODO // interface{} | The user id for which this password is created.
    userUpdateRequest := *openapiclient.NewUserUpdateRequest() // UserUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorisedUsersApi.UserUpdate(context.Background(), userId).UserUpdateRequest(userUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorisedUsersApi.UserUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | [**interface{}**](.md) | The user id for which this password is created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userUpdateRequest** | [**UserUpdateRequest**](UserUpdateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [auth_token](../README.md#auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGet

> UsersGet(ctx).Offset(offset).Limit(limit).Active(active).Email(email).Execute()

Get all users



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/stefan-cross/go-weavr-api"
)

func main() {
    offset := TODO // interface{} | The offset value for paging, indicating the initial item number to be returned from the data set satisfying the given criteria. Leave out to fetch the first page of results. (optional)
    limit := TODO // interface{} | The limit of the results for paging, starting at the offset. Limit is always capped at 100. (optional)
    active := TODO // interface{} | Filter for active or deactivated users. Leave out to fetch all users. (optional)
    email := TODO // interface{} | Filter by the email address of the user. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorisedUsersApi.UsersGet(context.Background()).Offset(offset).Limit(limit).Active(active).Email(email).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorisedUsersApi.UsersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | [**interface{}**](interface{}.md) | The offset value for paging, indicating the initial item number to be returned from the data set satisfying the given criteria. Leave out to fetch the first page of results. | 
 **limit** | [**interface{}**](interface{}.md) | The limit of the results for paging, starting at the offset. Limit is always capped at 100. | 
 **active** | [**interface{}**](interface{}.md) | Filter for active or deactivated users. Leave out to fetch all users. | 
 **email** | [**interface{}**](interface{}.md) | Filter by the email address of the user. | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [auth_token](../README.md#auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

