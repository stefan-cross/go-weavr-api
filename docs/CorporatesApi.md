# \CorporatesApi

All URIs are relative to *https://sandbox.weavr.io/multi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CorporateChargeFee**](CorporatesApi.md#CorporateChargeFee) | **Post** /corporates/fees/charge | Charge fee to a corporate
[**CorporateCreate**](CorporatesApi.md#CorporateCreate) | **Post** /corporates | Create a corporate
[**CorporateGet**](CorporatesApi.md#CorporateGet) | **Get** /corporates | Get a corporate
[**CorporateKybGet**](CorporatesApi.md#CorporateKybGet) | **Get** /corporates/kyb | Get KYB for a corporate
[**CorporateKybStart**](CorporatesApi.md#CorporateKybStart) | **Post** /corporates/kyb | Start KYB for a corporate
[**CorporateRootUserEmailVerificationCodeSend**](CorporatesApi.md#CorporateRootUserEmailVerificationCodeSend) | **Post** /corporates/verification/email/send | Send an email verification code to the root user
[**CorporateRootUserEmailVerify**](CorporatesApi.md#CorporateRootUserEmailVerify) | **Post** /corporates/verification/email/verify | Verify email of the root user
[**CorporateUpdate**](CorporatesApi.md#CorporateUpdate) | **Patch** /corporates | Update a corporate



## CorporateChargeFee

> CorporateChargeFee(ctx).Fee(fee).IdempotencyRef(idempotencyRef).Execute()

Charge fee to a corporate



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
    fee := *openapiclient.NewFee(interface{}(123), *openapiclient.NewInstrumentId(interface{}(123), *openapiclient.NewInstrumentType())) // Fee | 
    idempotencyRef := TODO // interface{} | A unique call reference generated by the caller that, taking into consideration the payload as well as the operation itself, helps avoid duplicate operations. Idempotency reference uniqueness is maintained for at least 24 hours. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorporatesApi.CorporateChargeFee(context.Background()).Fee(fee).IdempotencyRef(idempotencyRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorporatesApi.CorporateChargeFee``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCorporateChargeFeeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fee** | [**Fee**](Fee.md) |  | 
 **idempotencyRef** | [**interface{}**](interface{}.md) | A unique call reference generated by the caller that, taking into consideration the payload as well as the operation itself, helps avoid duplicate operations. Idempotency reference uniqueness is maintained for at least 24 hours. | 

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


## CorporateCreate

> CorporateCreate(ctx).CorporateCreateRequest(corporateCreateRequest).Execute()

Create a corporate



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
    corporateCreateRequest := *openapiclient.NewCorporateCreateRequest(interface{}(123), *openapiclient.NewCorporateCreateRequestRootUser(interface{}(123), interface{}(123), interface{}(123), *openapiclient.NewMobile(interface{}(123), interface{}(123)), *openapiclient.NewCompanyPosition()), *openapiclient.NewCorporateCreateRequestCompany(*openapiclient.NewCompanyType(), interface{}(123), interface{}(123)), interface{}(123), interface{}(123), interface{}(123)) // CorporateCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorporatesApi.CorporateCreate(context.Background()).CorporateCreateRequest(corporateCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorporatesApi.CorporateCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCorporateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **corporateCreateRequest** | [**CorporateCreateRequest**](CorporateCreateRequest.md) |  | 

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


## CorporateGet

> CorporateGet(ctx).Execute()

Get a corporate



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
    r, err := apiClient.CorporatesApi.CorporateGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorporatesApi.CorporateGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCorporateGetRequest struct via the builder pattern


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


## CorporateKybGet

> CorporateKybGet(ctx).Execute()

Get KYB for a corporate



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
    r, err := apiClient.CorporatesApi.CorporateKybGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorporatesApi.CorporateKybGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCorporateKybGetRequest struct via the builder pattern


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


## CorporateKybStart

> UserKyc200Response CorporateKybStart(ctx).Execute()

Start KYB for a corporate



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
    resp, r, err := apiClient.CorporatesApi.CorporateKybStart(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorporatesApi.CorporateKybStart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CorporateKybStart`: UserKyc200Response
    fmt.Fprintf(os.Stdout, "Response from `CorporatesApi.CorporateKybStart`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCorporateKybStartRequest struct via the builder pattern


### Return type

[**UserKyc200Response**](UserKyc200Response.md)

### Authorization

[api_key](../README.md#api_key), [auth_token](../README.md#auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CorporateRootUserEmailVerificationCodeSend

> CorporateRootUserEmailVerificationCodeSend(ctx).LostPasswordInitiateRequest(lostPasswordInitiateRequest).Execute()

Send an email verification code to the root user



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
    lostPasswordInitiateRequest := *openapiclient.NewLostPasswordInitiateRequest(interface{}(123)) // LostPasswordInitiateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorporatesApi.CorporateRootUserEmailVerificationCodeSend(context.Background()).LostPasswordInitiateRequest(lostPasswordInitiateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorporatesApi.CorporateRootUserEmailVerificationCodeSend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCorporateRootUserEmailVerificationCodeSendRequest struct via the builder pattern


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


## CorporateRootUserEmailVerify

> CorporateRootUserEmailVerify(ctx).AuthorisedUserEmailVerifyRequest(authorisedUserEmailVerifyRequest).Execute()

Verify email of the root user



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
    authorisedUserEmailVerifyRequest := *openapiclient.NewAuthorisedUserEmailVerifyRequest(interface{}(123), interface{}(123)) // AuthorisedUserEmailVerifyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorporatesApi.CorporateRootUserEmailVerify(context.Background()).AuthorisedUserEmailVerifyRequest(authorisedUserEmailVerifyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorporatesApi.CorporateRootUserEmailVerify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCorporateRootUserEmailVerifyRequest struct via the builder pattern


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


## CorporateUpdate

> CorporateUpdate(ctx).CorporateUpdateRequest(corporateUpdateRequest).Execute()

Update a corporate



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
    corporateUpdateRequest := *openapiclient.NewCorporateUpdateRequest() // CorporateUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CorporatesApi.CorporateUpdate(context.Background()).CorporateUpdateRequest(corporateUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorporatesApi.CorporateUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCorporateUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **corporateUpdateRequest** | [**CorporateUpdateRequest**](CorporateUpdateRequest.md) |  | 

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
