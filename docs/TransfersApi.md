# \TransfersApi

All URIs are relative to *https://sandbox.weavr.io/multi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransferCreate**](TransfersApi.md#TransferCreate) | **Post** /transfers | Create a transfer transaction
[**TransferGet**](TransfersApi.md#TransferGet) | **Get** /transfers/{id} | Get a transfer transaction
[**TransfersGet**](TransfersApi.md#TransfersGet) | **Get** /transfers | Get all transfer transactions



## TransferCreate

> TransferCreate(ctx).TransferCreateRequest(transferCreateRequest).IdempotencyRef(idempotencyRef).Execute()

Create a transfer transaction



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
    transferCreateRequest := *openapiclient.NewTransferCreateRequest(interface{}(123), *openapiclient.NewInstrumentId(interface{}(123), *openapiclient.NewInstrumentType()), *openapiclient.NewInstrumentId(interface{}(123), *openapiclient.NewInstrumentType()), *openapiclient.NewCurrencyAmount(interface{}(123), interface{}(123))) // TransferCreateRequest | 
    idempotencyRef := TODO // interface{} | A unique call reference generated by the caller that, taking into consideration the payload as well as the operation itself, helps avoid duplicate operations. Idempotency reference uniqueness is maintained for at least 24 hours. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TransfersApi.TransferCreate(context.Background()).TransferCreateRequest(transferCreateRequest).IdempotencyRef(idempotencyRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransfersApi.TransferCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferCreateRequest** | [**TransferCreateRequest**](TransferCreateRequest.md) |  | 
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


## TransferGet

> TransferGet(ctx, id).Execute()

Get a transfer transaction



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
    id := TODO // interface{} | The unique identifier of the Transfer transaction.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TransfersApi.TransferGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransfersApi.TransferGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**interface{}**](.md) | The unique identifier of the Transfer transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferGetRequest struct via the builder pattern


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


## TransfersGet

> TransfersGet(ctx).Offset(offset).Limit(limit).ProfileId(profileId).InstrumentId(instrumentId).State(state).CreatedFrom(createdFrom).CreatedTo(createdTo).Tag(tag).Execute()

Get all transfer transactions



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
    profileId := TODO // interface{} | Filter by the transfer profile. Leave out to fetch all transfer transactions. (optional)
    instrumentId := *openapiclient.NewInstrumentId(interface{}(123), *openapiclient.NewInstrumentType()) // InstrumentId | Filter by the source instrument id. (optional)
    state := TODO // interface{} | Filter by the transfer transaction state. Leave out to fetch all states. (optional)
    createdFrom := TODO // interface{} | Filter for transfer transactions created after `createdFrom` timestamp. Timestamp is expressed in Epoch timestamp using millisecond precision. Leave out to fetch all transfers. (optional)
    createdTo := TODO // interface{} | Filter for transfer transactions created before `createdTo` timestamp. Timestamp is expressed in Epoch timestamp using millisecond precision. Leave out to fetch all transfers. (optional)
    tag := TODO // interface{} | Filter by the transfer tag. The exact tag must be provided, as wildcards are not supported. Leave out to fetch all entries. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TransfersApi.TransfersGet(context.Background()).Offset(offset).Limit(limit).ProfileId(profileId).InstrumentId(instrumentId).State(state).CreatedFrom(createdFrom).CreatedTo(createdTo).Tag(tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransfersApi.TransfersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransfersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | [**interface{}**](interface{}.md) | The offset value for paging, indicating the initial item number to be returned from the data set satisfying the given criteria. Leave out to fetch the first page of results. | 
 **limit** | [**interface{}**](interface{}.md) | The limit of the results for paging, starting at the offset. Limit is always capped at 100. | 
 **profileId** | [**interface{}**](interface{}.md) | Filter by the transfer profile. Leave out to fetch all transfer transactions. | 
 **instrumentId** | [**InstrumentId**](InstrumentId.md) | Filter by the source instrument id. | 
 **state** | [**interface{}**](interface{}.md) | Filter by the transfer transaction state. Leave out to fetch all states. | 
 **createdFrom** | [**interface{}**](interface{}.md) | Filter for transfer transactions created after &#x60;createdFrom&#x60; timestamp. Timestamp is expressed in Epoch timestamp using millisecond precision. Leave out to fetch all transfers. | 
 **createdTo** | [**interface{}**](interface{}.md) | Filter for transfer transactions created before &#x60;createdTo&#x60; timestamp. Timestamp is expressed in Epoch timestamp using millisecond precision. Leave out to fetch all transfers. | 
 **tag** | [**interface{}**](interface{}.md) | Filter by the transfer tag. The exact tag must be provided, as wildcards are not supported. Leave out to fetch all entries. | 

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

