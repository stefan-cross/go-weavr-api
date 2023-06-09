# \ManagedAccountsApi

All URIs are relative to *https://sandbox.weavr.io/multi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ManagedAccountBlock**](ManagedAccountsApi.md#ManagedAccountBlock) | **Post** /managed_accounts/{id}/block | Block a managed account
[**ManagedAccountCreate**](ManagedAccountsApi.md#ManagedAccountCreate) | **Post** /managed_accounts | Create a managed account
[**ManagedAccountGet**](ManagedAccountsApi.md#ManagedAccountGet) | **Get** /managed_accounts/{id} | Get a managed account
[**ManagedAccountIBANGet**](ManagedAccountsApi.md#ManagedAccountIBANGet) | **Get** /managed_accounts/{id}/iban | Get a managed account IBAN
[**ManagedAccountRemove**](ManagedAccountsApi.md#ManagedAccountRemove) | **Post** /managed_accounts/{id}/remove | Remove a managed account
[**ManagedAccountStatement**](ManagedAccountsApi.md#ManagedAccountStatement) | **Get** /managed_accounts/{id}/statement | Get a managed account statement
[**ManagedAccountUnblock**](ManagedAccountsApi.md#ManagedAccountUnblock) | **Post** /managed_accounts/{id}/unblock | Unblock a managed account
[**ManagedAccountUpdate**](ManagedAccountsApi.md#ManagedAccountUpdate) | **Patch** /managed_accounts/{id} | Update a managed account
[**ManagedAccountsGet**](ManagedAccountsApi.md#ManagedAccountsGet) | **Get** /managed_accounts | Get all managed accounts
[**ManagedAccountsIBANUpgrade**](ManagedAccountsApi.md#ManagedAccountsIBANUpgrade) | **Post** /managed_accounts/{id}/iban | Upgrade a managed account with IBAN



## ManagedAccountBlock

> ManagedAccountBlock(ctx, id).Execute()

Block a managed account



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
    id := TODO // interface{} | The unique identifier of the managed account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedAccountsApi.ManagedAccountBlock(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsApi.ManagedAccountBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The unique identifier of the managed account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedAccountBlockRequest struct via the builder pattern


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


## ManagedAccountCreate

> ManagedAccountCreate(ctx).ManagedAccountCreateRequest(managedAccountCreateRequest).IdempotencyRef(idempotencyRef).Execute()

Create a managed account



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
    managedAccountCreateRequest := *openapiclient.NewManagedAccountCreateRequest(interface{}(123), interface{}(123), interface{}(123)) // ManagedAccountCreateRequest | 
    idempotencyRef := TODO // interface{} | A unique call reference generated by the caller that, taking into consideration the payload as well as the operation itself, helps avoid duplicate operations. Idempotency reference uniqueness is maintained for at least 24 hours. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedAccountsApi.ManagedAccountCreate(context.Background()).ManagedAccountCreateRequest(managedAccountCreateRequest).IdempotencyRef(idempotencyRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsApi.ManagedAccountCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiManagedAccountCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **managedAccountCreateRequest** | [**ManagedAccountCreateRequest**](ManagedAccountCreateRequest.md) |  | 
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


## ManagedAccountGet

> ManagedAccountGet(ctx, id).Execute()

Get a managed account



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
    id := TODO // interface{} | The unique identifier of the Managed Account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedAccountsApi.ManagedAccountGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsApi.ManagedAccountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The unique identifier of the Managed Account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedAccountGetRequest struct via the builder pattern


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


## ManagedAccountIBANGet

> ManagedAccountIBANGet(ctx, id).Execute()

Get a managed account IBAN



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
    id := TODO // interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedAccountsApi.ManagedAccountIBANGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsApi.ManagedAccountIBANGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiManagedAccountIBANGetRequest struct via the builder pattern


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


## ManagedAccountRemove

> ManagedAccountRemove(ctx, id).Execute()

Remove a managed account



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
    id := TODO // interface{} | The unique identifier of the managed account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedAccountsApi.ManagedAccountRemove(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsApi.ManagedAccountRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The unique identifier of the managed account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedAccountRemoveRequest struct via the builder pattern


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


## ManagedAccountStatement

> ManagedAccountStatement(ctx, id).Accept(accept).Offset(offset).Limit(limit).OrderByTimestamp(orderByTimestamp).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).ShowFundMovementsOnly(showFundMovementsOnly).Execute()

Get a managed account statement



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
    id := TODO // interface{} | The unique identifier of the managed account.
    accept := TODO // interface{} | A request parameter specifying the type of response the client would like. Must be one of `application/json`, `application/pdf` or `text/csv`.  The default response type (`application/json`) will be returned if specified incorrectly or not specified.  (optional)
    offset := TODO // interface{} | The offset value for paging, indicating the initial item number to be returned from the data set satisfying the given criteria. Leave out to fetch the first page of results. (optional)
    limit := TODO // interface{} | The limit of the results for paging, starting at the offset. Limit is always capped at 100. (optional)
    orderByTimestamp := TODO // interface{} | Orders the result in ascending or descending order.   - ASC: Ascending order, oldest transactions first.   - DESC: Descending order, most recent transactions first.  If not specified, the transactions will be returned in descending order.  (optional)
    fromTimestamp := TODO // interface{} | Filter for transactions having transaction timestamp after `fromTimestamp`. Timestamp is expressed in Epoch timestamp using millisecond precision. Leave out to fetch all managed accounts. (optional)
    toTimestamp := TODO // interface{} | Filter for transactions having transaction timestamp before `toTimestamp`. Timestamp is expressed in Epoch timestamp using millisecond precision. Leave out to fetch all managed accounts. (optional)
    showFundMovementsOnly := TODO // interface{} | Returns only the entries which contain fund movements. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedAccountsApi.ManagedAccountStatement(context.Background(), id).Accept(accept).Offset(offset).Limit(limit).OrderByTimestamp(orderByTimestamp).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).ShowFundMovementsOnly(showFundMovementsOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsApi.ManagedAccountStatement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The unique identifier of the managed account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedAccountStatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | [**interface{}**](interface{}.md) | A request parameter specifying the type of response the client would like. Must be one of &#x60;application/json&#x60;, &#x60;application/pdf&#x60; or &#x60;text/csv&#x60;.  The default response type (&#x60;application/json&#x60;) will be returned if specified incorrectly or not specified.  | 
 **offset** | [**interface{}**](interface{}.md) | The offset value for paging, indicating the initial item number to be returned from the data set satisfying the given criteria. Leave out to fetch the first page of results. | 
 **limit** | [**interface{}**](interface{}.md) | The limit of the results for paging, starting at the offset. Limit is always capped at 100. | 
 **orderByTimestamp** | [**interface{}**](interface{}.md) | Orders the result in ascending or descending order.   - ASC: Ascending order, oldest transactions first.   - DESC: Descending order, most recent transactions first.  If not specified, the transactions will be returned in descending order.  | 
 **fromTimestamp** | [**interface{}**](interface{}.md) | Filter for transactions having transaction timestamp after &#x60;fromTimestamp&#x60;. Timestamp is expressed in Epoch timestamp using millisecond precision. Leave out to fetch all managed accounts. | 
 **toTimestamp** | [**interface{}**](interface{}.md) | Filter for transactions having transaction timestamp before &#x60;toTimestamp&#x60;. Timestamp is expressed in Epoch timestamp using millisecond precision. Leave out to fetch all managed accounts. | 
 **showFundMovementsOnly** | [**interface{}**](interface{}.md) | Returns only the entries which contain fund movements. | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [auth_token](../README.md#auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/pdf, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManagedAccountUnblock

> ManagedAccountUnblock(ctx, id).Execute()

Unblock a managed account



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
    id := TODO // interface{} | The unique identifier of the managed account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedAccountsApi.ManagedAccountUnblock(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsApi.ManagedAccountUnblock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The unique identifier of the managed account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedAccountUnblockRequest struct via the builder pattern


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


## ManagedAccountUpdate

> ManagedAccountUpdate(ctx, id).ManagedAccountUpdateRequest(managedAccountUpdateRequest).Execute()

Update a managed account



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
    id := TODO // interface{} | The unique identifier of the managed account.
    managedAccountUpdateRequest := *openapiclient.NewManagedAccountUpdateRequest() // ManagedAccountUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedAccountsApi.ManagedAccountUpdate(context.Background(), id).ManagedAccountUpdateRequest(managedAccountUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsApi.ManagedAccountUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The unique identifier of the managed account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedAccountUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **managedAccountUpdateRequest** | [**ManagedAccountUpdateRequest**](ManagedAccountUpdateRequest.md) |  | 

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


## ManagedAccountsGet

> ManagedAccountsGet(ctx).Offset(offset).Limit(limit).ProfileId(profileId).FriendlyName(friendlyName).State(state).StateBlockedReason(stateBlockedReason).StateDestroyedReason(stateDestroyedReason).Currency(currency).CreatedFrom(createdFrom).CreatedTo(createdTo).Tag(tag).Execute()

Get all managed accounts



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
    profileId := TODO // interface{} | Filter by the managed account/card profile. Leave out to fetch all managed accounts/card. (optional)
    friendlyName := TODO // interface{} | Filter by the managed account/card friendly name. Leave out to fetch all managed accounts/card.  The exact name must be provided, as wildcards are not supported.  (optional)
    state := TODO // interface{} |  (optional)
    stateBlockedReason := TODO // interface{} |  (optional)
    stateDestroyedReason := TODO // interface{} |  (optional)
    currency := TODO // interface{} | Filter by the managed account/card currency.  Currencies are expressed as an ISO 4217 code. Leave out to fetch all managed accounts/card.  (optional)
    createdFrom := TODO // interface{} | Filter for managed accounts/cards created after `createdFrom` timestamp. Timestamp is expressed in Epoch timestamp using millisecond precision. Leave out to fetch all managed accounts/cards. (optional)
    createdTo := TODO // interface{} | Filter for managed accounts/cards created before `createdTo` timestamp. Timestamp is expressed in Epoch timestamp using millisecond precision. Leave out to fetch all managed accounts/cards. (optional)
    tag := TODO // interface{} | Filter by the managed account/card tag. The exact tag must be provided, as wildcards are not supported. Leave out to fetch all managed accounts/card. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedAccountsApi.ManagedAccountsGet(context.Background()).Offset(offset).Limit(limit).ProfileId(profileId).FriendlyName(friendlyName).State(state).StateBlockedReason(stateBlockedReason).StateDestroyedReason(stateDestroyedReason).Currency(currency).CreatedFrom(createdFrom).CreatedTo(createdTo).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsApi.ManagedAccountsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiManagedAccountsGetRequest struct via the builder pattern


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
 **createdFrom** | [**interface{}**](interface{}.md) | Filter for managed accounts/cards created after &#x60;createdFrom&#x60; timestamp. Timestamp is expressed in Epoch timestamp using millisecond precision. Leave out to fetch all managed accounts/cards. | 
 **createdTo** | [**interface{}**](interface{}.md) | Filter for managed accounts/cards created before &#x60;createdTo&#x60; timestamp. Timestamp is expressed in Epoch timestamp using millisecond precision. Leave out to fetch all managed accounts/cards. | 
 **tag** | [**interface{}**](interface{}.md) | Filter by the managed account/card tag. The exact tag must be provided, as wildcards are not supported. Leave out to fetch all managed accounts/card. | 

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


## ManagedAccountsIBANUpgrade

> ManagedAccountsIBANUpgrade(ctx, id).Execute()

Upgrade a managed account with IBAN



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
    id := TODO // interface{} | The unique identifier of the Managed Account.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ManagedAccountsApi.ManagedAccountsIBANUpgrade(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsApi.ManagedAccountsIBANUpgrade``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The unique identifier of the Managed Account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiManagedAccountsIBANUpgradeRequest struct via the builder pattern


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

