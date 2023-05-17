# \AccessApi

All URIs are relative to *https://sandbox.weavr.io/multi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoginWithPassword**](AccessApi.md#LoginWithPassword) | **Post** /login_with_password | Login with password
[**Logout**](AccessApi.md#Logout) | **Post** /logout | Logout
[**RequestAccessToken**](AccessApi.md#RequestAccessToken) | **Post** /access_token | Acquire a new access token
[**StepupSCAChallenge**](AccessApi.md#StepupSCAChallenge) | **Post** /stepup/challenges/otp/{channel} | Issue a one-time password that can be used to step-up a token
[**StepupSCAChallengePush**](AccessApi.md#StepupSCAChallengePush) | **Post** /stepup/challenges/push/{channel} | Issue a push notification that can be used to step-up a token
[**StepupSCAVerify**](AccessApi.md#StepupSCAVerify) | **Post** /stepup/challenges/otp/{channel}/verify | Verify a step-up token using a one-time password
[**UserIdentities**](AccessApi.md#UserIdentities) | **Get** /identities | Get user identities



## LoginWithPassword

> LoginWithPassword(ctx).LoginWithPasswordRequest(loginWithPasswordRequest).Execute()

Login with password



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    loginWithPasswordRequest := *openapiclient.NewLoginWithPasswordRequest(interface{}(123), *openapiclient.NewSensitivePassword(interface{}(123))) // LoginWithPasswordRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccessApi.LoginWithPassword(context.Background()).LoginWithPasswordRequest(loginWithPasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.LoginWithPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginWithPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginWithPasswordRequest** | [**LoginWithPasswordRequest**](LoginWithPasswordRequest.md) |  | 

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


## Logout

> Logout(ctx).Execute()

Logout



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccessApi.Logout(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.Logout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutRequest struct via the builder pattern


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


## RequestAccessToken

> RequestAccessToken(ctx).RequestAccessTokenRequest(requestAccessTokenRequest).Execute()

Acquire a new access token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    requestAccessTokenRequest := *openapiclient.NewRequestAccessTokenRequest(*openapiclient.NewIdentityId(interface{}(123), interface{}(123))) // RequestAccessTokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccessApi.RequestAccessToken(context.Background()).RequestAccessTokenRequest(requestAccessTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.RequestAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestAccessTokenRequest** | [**RequestAccessTokenRequest**](RequestAccessTokenRequest.md) |  | 

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


## StepupSCAChallenge

> StepupSCAChallenge(ctx, channel).Execute()

Issue a one-time password that can be used to step-up a token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    channel := *openapiclient.NewSCAOtpChannel() // SCAOtpChannel | The unique identifier for the channel.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccessApi.StepupSCAChallenge(context.Background(), channel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.StepupSCAChallenge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channel** | [**SCAOtpChannel**](.md) | The unique identifier for the channel. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStepupSCAChallengeRequest struct via the builder pattern


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


## StepupSCAChallengePush

> StepupSCAChallengePush(ctx, channel).Execute()

Issue a push notification that can be used to step-up a token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    channel := *openapiclient.NewSCAPushChannel() // SCAPushChannel | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccessApi.StepupSCAChallengePush(context.Background(), channel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.StepupSCAChallengePush``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channel** | [**SCAPushChannel**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStepupSCAChallengePushRequest struct via the builder pattern


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


## StepupSCAVerify

> StepupSCAVerify(ctx, channel).StepupSCAVerifyRequest(stepupSCAVerifyRequest).Execute()

Verify a step-up token using a one-time password



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    channel := *openapiclient.NewSCAOtpChannel() // SCAOtpChannel | The unique identifier for the channel.
    stepupSCAVerifyRequest := *openapiclient.NewStepupSCAVerifyRequest(interface{}(123)) // StepupSCAVerifyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccessApi.StepupSCAVerify(context.Background(), channel).StepupSCAVerifyRequest(stepupSCAVerifyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.StepupSCAVerify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channel** | [**SCAOtpChannel**](.md) | The unique identifier for the channel. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStepupSCAVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stepupSCAVerifyRequest** | [**StepupSCAVerifyRequest**](StepupSCAVerifyRequest.md) |  | 

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


## UserIdentities

> UserIdentities(ctx).Execute()

Get user identities



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccessApi.UserIdentities(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.UserIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserIdentitiesRequest struct via the builder pattern


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

