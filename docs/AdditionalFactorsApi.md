# \AdditionalFactorsApi

All URIs are relative to *https://sandbox.weavr.io/multi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthFactorsGet**](AdditionalFactorsApi.md#AuthFactorsGet) | **Get** /authentication_factors | Get user authentication factors
[**EnrolDeviceUsingOtpStepOne**](AdditionalFactorsApi.md#EnrolDeviceUsingOtpStepOne) | **Post** /authentication_factors/otp/{channel} | Enrol a user device for authentication using one-time passwords
[**EnrolDeviceUsingOtpStepTwo**](AdditionalFactorsApi.md#EnrolDeviceUsingOtpStepTwo) | **Post** /authentication_factors/otp/{channel}/verify | Verify enrolment of a user device for authentication using one-time passwords
[**EnrolDeviceUsingPush**](AdditionalFactorsApi.md#EnrolDeviceUsingPush) | **Post** /authentication_factors/push/{channel} | Enrol a user device for authentication using push notifications
[**UnlinkDeviceUsingPush**](AdditionalFactorsApi.md#UnlinkDeviceUsingPush) | **Delete** /authentication_factors/push/{channel} | Unlink a user device for authentication using push notifications



## AuthFactorsGet

> AuthFactorsGet(ctx).Execute()

Get user authentication factors



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
    r, err := apiClient.AdditionalFactorsApi.AuthFactorsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdditionalFactorsApi.AuthFactorsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthFactorsGetRequest struct via the builder pattern


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


## EnrolDeviceUsingOtpStepOne

> EnrolDeviceUsingOtpStepOne(ctx, channel).Execute()

Enrol a user device for authentication using one-time passwords



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
    r, err := apiClient.AdditionalFactorsApi.EnrolDeviceUsingOtpStepOne(context.Background(), channel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdditionalFactorsApi.EnrolDeviceUsingOtpStepOne``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnrolDeviceUsingOtpStepOneRequest struct via the builder pattern


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


## EnrolDeviceUsingOtpStepTwo

> EnrolDeviceUsingOtpStepTwo(ctx, channel).StepupSCAVerifyRequest(stepupSCAVerifyRequest).Execute()

Verify enrolment of a user device for authentication using one-time passwords



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
    r, err := apiClient.AdditionalFactorsApi.EnrolDeviceUsingOtpStepTwo(context.Background(), channel).StepupSCAVerifyRequest(stepupSCAVerifyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdditionalFactorsApi.EnrolDeviceUsingOtpStepTwo``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnrolDeviceUsingOtpStepTwoRequest struct via the builder pattern


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


## EnrolDeviceUsingPush

> EnrolDeviceUsingPush(ctx, channel).Execute()

Enrol a user device for authentication using push notifications



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
    r, err := apiClient.AdditionalFactorsApi.EnrolDeviceUsingPush(context.Background(), channel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdditionalFactorsApi.EnrolDeviceUsingPush``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnrolDeviceUsingPushRequest struct via the builder pattern


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


## UnlinkDeviceUsingPush

> UnlinkDeviceUsingPush(ctx, channel).Execute()

Unlink a user device for authentication using push notifications



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
    r, err := apiClient.AdditionalFactorsApi.UnlinkDeviceUsingPush(context.Background(), channel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdditionalFactorsApi.UnlinkDeviceUsingPush``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUnlinkDeviceUsingPushRequest struct via the builder pattern


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

