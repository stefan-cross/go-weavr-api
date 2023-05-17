# \ManagedCardsApi

All URIs are relative to *https://sandbox.weavr.io/multi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ManagedCardAssign**](ManagedCardsApi.md#ManagedCardAssign) | **Post** /managed_cards/assign | Assign a managed card
[**ManagedCardBlock**](ManagedCardsApi.md#ManagedCardBlock) | **Post** /managed_cards/{id}/block | Block a managed card
[**ManagedCardCreate**](ManagedCardsApi.md#ManagedCardCreate) | **Post** /managed_cards | Create a managed card
[**ManagedCardGet**](ManagedCardsApi.md#ManagedCardGet) | **Get** /managed_cards/{id} | Get a managed card
[**ManagedCardPhysicalActivate**](ManagedCardsApi.md#ManagedCardPhysicalActivate) | **Post** /managed_cards/{id}/physical/activate | Activate a physical card
[**ManagedCardPhysicalContactlessLimitReset**](ManagedCardsApi.md#ManagedCardPhysicalContactlessLimitReset) | **Post** /managed_cards/{id}/physical/contactless_limit/reset | Reset contactless limit for a physical card
[**ManagedCardPhysicalPinGet**](ManagedCardsApi.md#ManagedCardPhysicalPinGet) | **Get** /managed_cards/{id}/physical/pin | Get PIN for a physical card
[**ManagedCardPhysicalPinUnBlock**](ManagedCardsApi.md#ManagedCardPhysicalPinUnBlock) | **Patch** /managed_cards/{id}/physical/pin/unblock | Unblock PIN for a physical card
[**ManagedCardPhysicalReplaceDamaged**](ManagedCardsApi.md#ManagedCardPhysicalReplaceDamaged) | **Post** /managed_cards/{id}/physical/replace_damaged | Replace a damaged physical card
[**ManagedCardPhysicalReplaceLostOrStolen**](ManagedCardsApi.md#ManagedCardPhysicalReplaceLostOrStolen) | **Post** /managed_cards/{id}/physical/replace_lost_stolen | Replace a lost or stolen physical card
[**ManagedCardPhysicalReportLost**](ManagedCardsApi.md#ManagedCardPhysicalReportLost) | **Post** /managed_cards/{id}/physical/report_lost | Report a physical card as lost
[**ManagedCardPhysicalReportStolen**](ManagedCardsApi.md#ManagedCardPhysicalReportStolen) | **Post** /managed_cards/{id}/physical/report_stolen | Report a physical card as stolen
[**ManagedCardPhysicalUpgrade**](ManagedCardsApi.md#ManagedCardPhysicalUpgrade) | **Post** /managed_cards/{id}/physical | Upgrade a card to physical
[**ManagedCardRemove**](ManagedCardsApi.md#ManagedCardRemove) | **Post** /managed_cards/{id}/remove | Remove a managed card
[**ManagedCardSpendRulesCreate**](ManagedCardsApi.md#ManagedCardSpendRulesCreate) | **Post** /managed_cards/{id}/spend_rules | Create spend rules for a managed card
[**ManagedCardSpendRulesDelete**](ManagedCardsApi.md#ManagedCardSpendRulesDelete) | **Delete** /managed_cards/{id}/spend_rules | Delete all spend rules for a managed card
[**ManagedCardSpendRulesGet**](ManagedCardsApi.md#ManagedCardSpendRulesGet) | **Get** /managed_cards/{id}/spend_rules | Get all spend rules for a managed card
[**ManagedCardSpendRulesSet**](ManagedCardsApi.md#ManagedCardSpendRulesSet) | **Put** /managed_cards/{id}/spend_rules | Set spend rules for a managed card
[**ManagedCardSpendRulesUpdate**](ManagedCardsApi.md#ManagedCardSpendRulesUpdate) | **Patch** /managed_cards/{id}/spend_rules | Update spend rules for a managed card
[**ManagedCardStatement**](ManagedCardsApi.md#ManagedCardStatement) | **Get** /managed_cards/{id}/statement | Get a managed card statement
[**ManagedCardUnblock**](ManagedCardsApi.md#ManagedCardUnblock) | **Post** /managed_cards/{id}/unblock | Unblock a managed card
[**ManagedCardUpdate**](ManagedCardsApi.md#ManagedCardUpdate) | **Patch** /managed_cards/{id} | Update a managed card
[**ManagedCardsGet**](ManagedCardsApi.md#ManagedCardsGet) | **Get** /managed_cards | Get all managed cards



## ManagedCardAssign

> ManagedCardAssign(ctx).ManagedCardAssignRequest(managedCardAssignRequest).Execute()

Assign a managed card



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
    managedCardAssignRequest := *openapiclient.NewManagedCardAssignRequest(interface{}(123), interface{}(123), interface{}(123), interface{}(123), *openapiclient.NewAddress(interface{}(123), interface{}(123), interface{}(123), interface{}(123))) // ManagedCardAssignRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedCardsApi.ManagedCardAssign(context.Background()).ManagedCardAssignRequest(managedCardAssignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCardsApi.ManagedCardAssign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiManagedCardAssignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **managedCardAssignRequest** | [**ManagedCardAssignRequest**](ManagedCardAssignRequest.md) |  | 

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


## ManagedCardBlock

> ManagedCardBlock(ctx, id).Execute()

Block a managed card



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
    id := TODO // interface{} | The unique identifier of a managed card.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedCardsApi.ManagedCardBlock(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCardsApi.ManagedCardBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The unique identifier of a managed card. | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedCardBlockRequest struct via the builder pattern


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


## ManagedCardCreate

> ManagedCardCreate(ctx).ManagedCardRequest(managedCardRequest).IdempotencyRef(idempotencyRef).Execute()

Create a managed card



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
    managedCardRequest := *openapiclient.NewManagedCardRequest(interface{}(123), interface{}(123), interface{}(123), interface{}(123), *openapiclient.NewAddress(interface{}(123), interface{}(123), interface{}(123), interface{}(123)), interface{}(123)) // ManagedCardRequest | 
    idempotencyRef := TODO // interface{} | A unique call reference generated by the caller that, taking into consideration the payload as well as the operation itself, helps avoid duplicate operations. Idempotency reference uniqueness is maintained for at least 24 hours. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedCardsApi.ManagedCardCreate(context.Background()).ManagedCardRequest(managedCardRequest).IdempotencyRef(idempotencyRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCardsApi.ManagedCardCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiManagedCardCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **managedCardRequest** | [**ManagedCardRequest**](ManagedCardRequest.md) |  | 
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


## ManagedCardGet

> ManagedCardGet(ctx, id).Execute()

Get a managed card



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
    id := TODO // interface{} | The unique identifier of a card.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedCardsApi.ManagedCardGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCardsApi.ManagedCardGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The unique identifier of a card. | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedCardGetRequest struct via the builder pattern


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


## ManagedCardPhysicalActivate

> ManagedCardPhysicalActivate(ctx, id).ManagedCardPhysicalActivateRequest(managedCardPhysicalActivateRequest).Execute()

Activate a physical card



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
    id := TODO // interface{} | 
    managedCardPhysicalActivateRequest := *openapiclient.NewManagedCardPhysicalActivateRequest(interface{}(123)) // ManagedCardPhysicalActivateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedCardsApi.ManagedCardPhysicalActivate(context.Background(), id).ManagedCardPhysicalActivateRequest(managedCardPhysicalActivateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCardsApi.ManagedCardPhysicalActivate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedCardPhysicalActivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **managedCardPhysicalActivateRequest** | [**ManagedCardPhysicalActivateRequest**](ManagedCardPhysicalActivateRequest.md) |  | 

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


## ManagedCardPhysicalContactlessLimitReset

> ManagedCardPhysicalContactlessLimitReset(ctx, id).Execute()

Reset contactless limit for a physical card



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
    id := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedCardsApi.ManagedCardPhysicalContactlessLimitReset(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCardsApi.ManagedCardPhysicalContactlessLimitReset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedCardPhysicalContactlessLimitResetRequest struct via the builder pattern


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


## ManagedCardPhysicalPinGet

> ManagedCardPhysicalPinGet(ctx, id).Execute()

Get PIN for a physical card



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
    id := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedCardsApi.ManagedCardPhysicalPinGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCardsApi.ManagedCardPhysicalPinGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedCardPhysicalPinGetRequest struct via the builder pattern


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


## ManagedCardPhysicalPinUnBlock

> ManagedCardPhysicalPinUnBlock(ctx, id).Execute()

Unblock PIN for a physical card



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
    id := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedCardsApi.ManagedCardPhysicalPinUnBlock(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCardsApi.ManagedCardPhysicalPinUnBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedCardPhysicalPinUnBlockRequest struct via the builder pattern


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


## ManagedCardPhysicalReplaceDamaged

> ManagedCardPhysicalReplaceDamaged(ctx, id).ManagedCardPhysicalReplaceDamagedRequest(managedCardPhysicalReplaceDamagedRequest).Execute()

Replace a damaged physical card



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
    id := TODO // interface{} | 
    managedCardPhysicalReplaceDamagedRequest := *openapiclient.NewManagedCardPhysicalReplaceDamagedRequest(interface{}(123)) // ManagedCardPhysicalReplaceDamagedRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedCardsApi.ManagedCardPhysicalReplaceDamaged(context.Background(), id).ManagedCardPhysicalReplaceDamagedRequest(managedCardPhysicalReplaceDamagedRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCardsApi.ManagedCardPhysicalReplaceDamaged``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedCardPhysicalReplaceDamagedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **managedCardPhysicalReplaceDamagedRequest** | [**ManagedCardPhysicalReplaceDamagedRequest**](ManagedCardPhysicalReplaceDamagedRequest.md) |  | 

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


## ManagedCardPhysicalReplaceLostOrStolen

> ManagedCardPhysicalReplaceLostOrStolen(ctx, id).ManagedCardPhysicalReplaceLostOrStolenRequest(managedCardPhysicalReplaceLostOrStolenRequest).Execute()

Replace a lost or stolen physical card



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
    id := TODO // interface{} | 
    managedCardPhysicalReplaceLostOrStolenRequest := *openapiclient.NewManagedCardPhysicalReplaceLostOrStolenRequest(interface{}(123)) // ManagedCardPhysicalReplaceLostOrStolenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedCardsApi.ManagedCardPhysicalReplaceLostOrStolen(context.Background(), id).ManagedCardPhysicalReplaceLostOrStolenRequest(managedCardPhysicalReplaceLostOrStolenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCardsApi.ManagedCardPhysicalReplaceLostOrStolen``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedCardPhysicalReplaceLostOrStolenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **managedCardPhysicalReplaceLostOrStolenRequest** | [**ManagedCardPhysicalReplaceLostOrStolenRequest**](ManagedCardPhysicalReplaceLostOrStolenRequest.md) |  | 

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


## ManagedCardPhysicalReportLost

> ManagedCardPhysicalReportLost(ctx, id).Execute()

Report a physical card as lost



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
    id := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedCardsApi.ManagedCardPhysicalReportLost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCardsApi.ManagedCardPhysicalReportLost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedCardPhysicalReportLostRequest struct via the builder pattern


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


## ManagedCardPhysicalReportStolen

> ManagedCardPhysicalReportStolen(ctx, id).Execute()

Report a physical card as stolen



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
    id := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedCardsApi.ManagedCardPhysicalReportStolen(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCardsApi.ManagedCardPhysicalReportStolen``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedCardPhysicalReportStolenRequest struct via the builder pattern


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


## ManagedCardPhysicalUpgrade

> ManagedCardPhysicalUpgrade(ctx, id).ManagedCardPhysicalUpgradeRequest(managedCardPhysicalUpgradeRequest).Execute()

Upgrade a card to physical



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
    id := TODO // interface{} | 
    managedCardPhysicalUpgradeRequest := *openapiclient.NewManagedCardPhysicalUpgradeRequest(*openapiclient.NewDeliveryAddress(interface{}(123), interface{}(123), interface{}(123), interface{}(123), interface{}(123), interface{}(123)), interface{}(123)) // ManagedCardPhysicalUpgradeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedCardsApi.ManagedCardPhysicalUpgrade(context.Background(), id).ManagedCardPhysicalUpgradeRequest(managedCardPhysicalUpgradeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCardsApi.ManagedCardPhysicalUpgrade``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedCardPhysicalUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **managedCardPhysicalUpgradeRequest** | [**ManagedCardPhysicalUpgradeRequest**](ManagedCardPhysicalUpgradeRequest.md) |  | 

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


## ManagedCardRemove

> ManagedCardRemove(ctx, id).Execute()

Remove a managed card



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
    id := TODO // interface{} | The unique identifier of a managed card.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedCardsApi.ManagedCardRemove(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCardsApi.ManagedCardRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The unique identifier of a managed card. | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedCardRemoveRequest struct via the builder pattern


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


## ManagedCardSpendRulesCreate

> ManagedCardSpendRulesCreate(ctx, id).Body(body).Execute()

Create spend rules for a managed card



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
    id := TODO // interface{} | 
    body := interface{}(987) // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedCardsApi.ManagedCardSpendRulesCreate(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCardsApi.ManagedCardSpendRulesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedCardSpendRulesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** |  | 

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


## ManagedCardSpendRulesDelete

> ManagedCardSpendRulesDelete(ctx, id).Execute()

Delete all spend rules for a managed card



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
    id := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedCardsApi.ManagedCardSpendRulesDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCardsApi.ManagedCardSpendRulesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedCardSpendRulesDeleteRequest struct via the builder pattern


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


## ManagedCardSpendRulesGet

> ManagedCardSpendRulesGet(ctx, id).Execute()

Get all spend rules for a managed card



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
    id := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedCardsApi.ManagedCardSpendRulesGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCardsApi.ManagedCardSpendRulesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedCardSpendRulesGetRequest struct via the builder pattern


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


## ManagedCardSpendRulesSet

> ManagedCardSpendRulesSet(ctx, id).Body(body).Execute()

Set spend rules for a managed card



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
    id := TODO // interface{} | 
    body := interface{}(987) // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedCardsApi.ManagedCardSpendRulesSet(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCardsApi.ManagedCardSpendRulesSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedCardSpendRulesSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** |  | 

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


## ManagedCardSpendRulesUpdate

> ManagedCardSpendRulesUpdate(ctx, id).Body(body).Execute()

Update spend rules for a managed card



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
    id := TODO // interface{} | 
    body := interface{}(987) // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedCardsApi.ManagedCardSpendRulesUpdate(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCardsApi.ManagedCardSpendRulesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedCardSpendRulesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** |  | 

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


## ManagedCardStatement

> ManagedCardStatement(ctx, id).Accept(accept).Offset(offset).Limit(limit).OrderByTimestamp(orderByTimestamp).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).Execute()

Get a managed card statement



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
    id := TODO // interface{} | The unique identifier of a managed card.
    accept := TODO // interface{} | A request parameter specifying the type of response the client would like. Must be one of `application/json`, `application/pdf` or `text/csv`.  The default response type (`application/json`) will be returned if specified incorrectly or not specified.  (optional)
    offset := TODO // interface{} | The offset value for paging, indicating the initial item number to be returned from the data set satisfying the given criteria. Leave out to fetch the first page of results. (optional)
    limit := TODO // interface{} | The limit of the results for paging, starting at the offset. Limit is always capped at 100. (optional)
    orderByTimestamp := TODO // interface{} | Orders the result in ascending or descending order.   - ASC: Ascending order, oldest transactions first.   - DESC: Descending order, most recent transactions first.  If not specified, the transactions will be returned in descending order.  (optional)
    fromTimestamp := TODO // interface{} | Filter for transactions having transaction timestamp after `fromTimestamp`. Timestamp is expressed in Epoch timestamp using millisecond precision. (optional)
    toTimestamp := TODO // interface{} | Filter for transactions having transaction timestamp before the `toTimestamp`. Timestamp is expressed in Epoch timestamp using millisecond precision. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedCardsApi.ManagedCardStatement(context.Background(), id).Accept(accept).Offset(offset).Limit(limit).OrderByTimestamp(orderByTimestamp).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCardsApi.ManagedCardStatement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The unique identifier of a managed card. | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedCardStatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | [**interface{}**](interface{}.md) | A request parameter specifying the type of response the client would like. Must be one of &#x60;application/json&#x60;, &#x60;application/pdf&#x60; or &#x60;text/csv&#x60;.  The default response type (&#x60;application/json&#x60;) will be returned if specified incorrectly or not specified.  | 
 **offset** | [**interface{}**](interface{}.md) | The offset value for paging, indicating the initial item number to be returned from the data set satisfying the given criteria. Leave out to fetch the first page of results. | 
 **limit** | [**interface{}**](interface{}.md) | The limit of the results for paging, starting at the offset. Limit is always capped at 100. | 
 **orderByTimestamp** | [**interface{}**](interface{}.md) | Orders the result in ascending or descending order.   - ASC: Ascending order, oldest transactions first.   - DESC: Descending order, most recent transactions first.  If not specified, the transactions will be returned in descending order.  | 
 **fromTimestamp** | [**interface{}**](interface{}.md) | Filter for transactions having transaction timestamp after &#x60;fromTimestamp&#x60;. Timestamp is expressed in Epoch timestamp using millisecond precision. | 
 **toTimestamp** | [**interface{}**](interface{}.md) | Filter for transactions having transaction timestamp before the &#x60;toTimestamp&#x60;. Timestamp is expressed in Epoch timestamp using millisecond precision. | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [auth_token](../README.md#auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManagedCardUnblock

> ManagedCardUnblock(ctx, id).Execute()

Unblock a managed card



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
    id := TODO // interface{} | The unique identifier of the managed card.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedCardsApi.ManagedCardUnblock(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCardsApi.ManagedCardUnblock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The unique identifier of the managed card. | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedCardUnblockRequest struct via the builder pattern


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


## ManagedCardUpdate

> ManagedCardUpdate(ctx, id).ManagedCardUpdateRequest(managedCardUpdateRequest).Execute()

Update a managed card



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
    id := TODO // interface{} | The unique identifier of a card.
    managedCardUpdateRequest := *openapiclient.NewManagedCardUpdateRequest() // ManagedCardUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedCardsApi.ManagedCardUpdate(context.Background(), id).ManagedCardUpdateRequest(managedCardUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCardsApi.ManagedCardUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The unique identifier of a card. | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedCardUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **managedCardUpdateRequest** | [**ManagedCardUpdateRequest**](ManagedCardUpdateRequest.md) |  | 

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


## ManagedCardsGet

> ManagedCardsGet(ctx).Offset(offset).Limit(limit).ProfileId(profileId).FriendlyName(friendlyName).State(state).StateBlockedReason(stateBlockedReason).StateDestroyedReason(stateDestroyedReason).Currency(currency).Type_(type_).ExternalHandle(externalHandle).CardNumberFirstSix(cardNumberFirstSix).CardNumberLastFour(cardNumberLastFour).CreatedFrom(createdFrom).CreatedTo(createdTo).Mode(mode).Tag(tag).ParentManagedAccountId(parentManagedAccountId).ManufacturingState(manufacturingState).Execute()

Get all managed cards



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
    offset := TODO // interface{} | The offset value for paging, indicating the initial item number to be returned from the data set satisfying the given criteria. Leave out to fetch the first page of results. (optional)
    limit := TODO // interface{} | The limit of the results for paging, starting at the offset. Limit is always capped at 100. (optional)
    profileId := TODO // interface{} | Filter by the managed account/card profile. Leave out to fetch all managed accounts/card. (optional)
    friendlyName := TODO // interface{} | Filter by the managed account/card friendly name. Leave out to fetch all managed accounts/card.  The exact name must be provided, as wildcards are not supported.  (optional)
    state := TODO // interface{} |  (optional)
    stateBlockedReason := TODO // interface{} |  (optional)
    stateDestroyedReason := TODO // interface{} |  (optional)
    currency := TODO // interface{} | Filter by the managed account/card currency.  Currencies are expressed as an ISO 4217 code. Leave out to fetch all managed accounts/card.  (optional)
    type_ := TODO // interface{} | Filter by the type of the card. (optional)
    externalHandle := TODO // interface{} | Search by the card's `externalHandle`. (optional)
    cardNumberFirstSix := TODO // interface{} | Filter by first six digits of the card. (optional)
    cardNumberLastFour := TODO // interface{} | Filter by last four digits of the card. (optional)
    createdFrom := TODO // interface{} | Filter for managed accounts/cards created after `createdFrom` timestamp. Timestamp is expressed in Epoch timestamp using millisecond precision. Leave out to fetch all managed accounts/cards. (optional)
    createdTo := TODO // interface{} | Filter for managed accounts/cards created before `createdTo` timestamp. Timestamp is expressed in Epoch timestamp using millisecond precision. Leave out to fetch all managed accounts/cards. (optional)
    mode := TODO // interface{} | Filter by card mode (prepaid mode or debit mode). (optional)
    tag := TODO // interface{} | Filter by the managed account/card tag. The exact tag must be provided, as wildcards are not supported. Leave out to fetch all managed accounts/card. (optional)
    parentManagedAccountId := TODO // interface{} | Filter by the Id of the parent managed account associated with the card. This is applicable only for debit mode cards. (optional)
    manufacturingState := TODO // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedCardsApi.ManagedCardsGet(context.Background()).Offset(offset).Limit(limit).ProfileId(profileId).FriendlyName(friendlyName).State(state).StateBlockedReason(stateBlockedReason).StateDestroyedReason(stateDestroyedReason).Currency(currency).Type_(type_).ExternalHandle(externalHandle).CardNumberFirstSix(cardNumberFirstSix).CardNumberLastFour(cardNumberLastFour).CreatedFrom(createdFrom).CreatedTo(createdTo).Mode(mode).Tag(tag).ParentManagedAccountId(parentManagedAccountId).ManufacturingState(manufacturingState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedCardsApi.ManagedCardsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiManagedCardsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | [**interface{}**](interface{}.md) | The offset value for paging, indicating the initial item number to be returned from the data set satisfying the given criteria. Leave out to fetch the first page of results. | 
 **limit** | [**interface{}**](interface{}.md) | The limit of the results for paging, starting at the offset. Limit is always capped at 100. | 
 **profileId** | [**interface{}**](interface{}.md) | Filter by the managed account/card profile. Leave out to fetch all managed accounts/card. | 
 **friendlyName** | [**interface{}**](interface{}.md) | Filter by the managed account/card friendly name. Leave out to fetch all managed accounts/card.  The exact name must be provided, as wildcards are not supported.  | 
 **state** | [**interface{}**](interface{}.md) |  | 
 **stateBlockedReason** | [**interface{}**](interface{}.md) |  | 
 **stateDestroyedReason** | [**interface{}**](interface{}.md) |  | 
 **currency** | [**interface{}**](interface{}.md) | Filter by the managed account/card currency.  Currencies are expressed as an ISO 4217 code. Leave out to fetch all managed accounts/card.  | 
 **type_** | [**interface{}**](interface{}.md) | Filter by the type of the card. | 
 **externalHandle** | [**interface{}**](interface{}.md) | Search by the card&#39;s &#x60;externalHandle&#x60;. | 
 **cardNumberFirstSix** | [**interface{}**](interface{}.md) | Filter by first six digits of the card. | 
 **cardNumberLastFour** | [**interface{}**](interface{}.md) | Filter by last four digits of the card. | 
 **createdFrom** | [**interface{}**](interface{}.md) | Filter for managed accounts/cards created after &#x60;createdFrom&#x60; timestamp. Timestamp is expressed in Epoch timestamp using millisecond precision. Leave out to fetch all managed accounts/cards. | 
 **createdTo** | [**interface{}**](interface{}.md) | Filter for managed accounts/cards created before &#x60;createdTo&#x60; timestamp. Timestamp is expressed in Epoch timestamp using millisecond precision. Leave out to fetch all managed accounts/cards. | 
 **mode** | [**interface{}**](interface{}.md) | Filter by card mode (prepaid mode or debit mode). | 
 **tag** | [**interface{}**](interface{}.md) | Filter by the managed account/card tag. The exact tag must be provided, as wildcards are not supported. Leave out to fetch all managed accounts/card. | 
 **parentManagedAccountId** | [**interface{}**](interface{}.md) | Filter by the Id of the parent managed account associated with the card. This is applicable only for debit mode cards. | 
 **manufacturingState** | [**interface{}**](interface{}.md) |  | 

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

