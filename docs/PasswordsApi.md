# \PasswordsApi

All URIs are relative to *https://sandbox.weavr.io/multi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LostPasswordInitiate**](PasswordsApi.md#LostPasswordInitiate) | **Post** /passwords/lost_password/start | Initiate lost password process
[**LostPasswordResume**](PasswordsApi.md#LostPasswordResume) | **Post** /passwords/lost_password/resume | Resume lost password process
[**PasswordCreate**](PasswordsApi.md#PasswordCreate) | **Post** /passwords/{user_id}/create | Create a password
[**PasswordUpdate**](PasswordsApi.md#PasswordUpdate) | **Post** /passwords/update | Update a password
[**PasswordValidate**](PasswordsApi.md#PasswordValidate) | **Post** /passwords/validate | Validate a password



## LostPasswordInitiate

> LostPasswordInitiate(ctx).LostPasswordInitiateRequest(lostPasswordInitiateRequest).Execute()

Initiate lost password process



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
    r, err := apiClient.PasswordsApi.LostPasswordInitiate(context.Background()).LostPasswordInitiateRequest(lostPasswordInitiateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordsApi.LostPasswordInitiate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLostPasswordInitiateRequest struct via the builder pattern


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


## LostPasswordResume

> LostPasswordResume(ctx).LostPasswordResumeRequest(lostPasswordResumeRequest).Execute()

Resume lost password process



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
    lostPasswordResumeRequest := *openapiclient.NewLostPasswordResumeRequest(interface{}(123), interface{}(123), *openapiclient.NewSensitivePassword(interface{}(123))) // LostPasswordResumeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PasswordsApi.LostPasswordResume(context.Background()).LostPasswordResumeRequest(lostPasswordResumeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordsApi.LostPasswordResume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLostPasswordResumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lostPasswordResumeRequest** | [**LostPasswordResumeRequest**](LostPasswordResumeRequest.md) |  | 

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


## PasswordCreate

> PasswordCreate(ctx, userId).PasswordCreateRequest(passwordCreateRequest).Execute()

Create a password



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
    passwordCreateRequest := *openapiclient.NewPasswordCreateRequest(*openapiclient.NewSensitivePassword(interface{}(123))) // PasswordCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PasswordsApi.PasswordCreate(context.Background(), userId).PasswordCreateRequest(passwordCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordsApi.PasswordCreate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPasswordCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **passwordCreateRequest** | [**PasswordCreateRequest**](PasswordCreateRequest.md) |  | 

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


## PasswordUpdate

> PasswordUpdate(ctx).PasswordUpdateRequest(passwordUpdateRequest).Execute()

Update a password



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
    passwordUpdateRequest := *openapiclient.NewPasswordUpdateRequest(*openapiclient.NewSensitivePassword(interface{}(123)), *openapiclient.NewSensitivePassword(interface{}(123))) // PasswordUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PasswordsApi.PasswordUpdate(context.Background()).PasswordUpdateRequest(passwordUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordsApi.PasswordUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPasswordUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordUpdateRequest** | [**PasswordUpdateRequest**](PasswordUpdateRequest.md) |  | 

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


## PasswordValidate

> PasswordValidate(ctx).PasswordCreateRequest(passwordCreateRequest).Execute()

Validate a password



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
    passwordCreateRequest := *openapiclient.NewPasswordCreateRequest(*openapiclient.NewSensitivePassword(interface{}(123))) // PasswordCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PasswordsApi.PasswordValidate(context.Background()).PasswordCreateRequest(passwordCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordsApi.PasswordValidate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPasswordValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordCreateRequest** | [**PasswordCreateRequest**](PasswordCreateRequest.md) |  | 

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

