# \MotivatorServiceApi

All URIs are relative to *https://echo.api.live.mindtastic.lol*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MotivatorGet**](MotivatorServiceApi.md#MotivatorGet) | **Get** /motivator | Get motivators for authenticated user
[**MotivatorResultAdd**](MotivatorServiceApi.md#MotivatorResultAdd) | **Post** /motivator/{motivatorId}/result | Add a new result to user&#39;s current motivator.
[**MotivatorResultDeleteAll**](MotivatorServiceApi.md#MotivatorResultDeleteAll) | **Delete** /motivator/{motivatorId}/result | Deletes all results on a motivator for the user
[**MotivatorSafetyNetAdd**](MotivatorServiceApi.md#MotivatorSafetyNetAdd) | **Post** /safetyNet | Add safety net item to currently authenticated user safety net.
[**MotivatorSafetyNetDelete**](MotivatorServiceApi.md#MotivatorSafetyNetDelete) | **Delete** /safetyNet/{safetyNetItemId} | Deletes an item from users safety net
[**MotivatorSafetyNetGet**](MotivatorServiceApi.md#MotivatorSafetyNetGet) | **Get** /safetyNet | Get safety net for the current user
[**MotivatorSafetyNetReplace**](MotivatorServiceApi.md#MotivatorSafetyNetReplace) | **Put** /safetyNet/{safetyNetItemId} | Replaces a safety net item with a new one



## MotivatorGet

> []Motivator MotivatorGet(ctx).Execute()

Get motivators for authenticated user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MotivatorServiceApi.MotivatorGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MotivatorServiceApi.MotivatorGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MotivatorGet`: []Motivator
    fmt.Fprintf(os.Stdout, "Response from `MotivatorServiceApi.MotivatorGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMotivatorGetRequest struct via the builder pattern


### Return type

[**[]Motivator**](Motivator.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MotivatorResultAdd

> MotivatorResultCreatedResponse MotivatorResultAdd(ctx, motivatorId).MotivatorResult(motivatorResult).Execute()

Add a new result to user's current motivator.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    motivatorId := int32(56) // int32 | The `id` of the motivator you want to add a result for.
    motivatorResult := *openapiclient.NewMotivatorResult(time.Now(), map[string]MotivatorResultEntry{"key": *openapiclient.NewMotivatorResultEntry()}) // MotivatorResult | The result to submit on the motivator.   The `values` property acts as a dictionary, containing the IDs referenced in the `input` object in the motivator content.   The coressponding value, again is a dictionary of mapping the reference string to objects containing a single `value` property.   The type of `value` depends on the individual input type. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MotivatorServiceApi.MotivatorResultAdd(context.Background(), motivatorId).MotivatorResult(motivatorResult).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MotivatorServiceApi.MotivatorResultAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MotivatorResultAdd`: MotivatorResultCreatedResponse
    fmt.Fprintf(os.Stdout, "Response from `MotivatorServiceApi.MotivatorResultAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**motivatorId** | **int32** | The &#x60;id&#x60; of the motivator you want to add a result for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMotivatorResultAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **motivatorResult** | [**MotivatorResult**](MotivatorResult.md) | The result to submit on the motivator.   The &#x60;values&#x60; property acts as a dictionary, containing the IDs referenced in the &#x60;input&#x60; object in the motivator content.   The coressponding value, again is a dictionary of mapping the reference string to objects containing a single &#x60;value&#x60; property.   The type of &#x60;value&#x60; depends on the individual input type. | 

### Return type

[**MotivatorResultCreatedResponse**](MotivatorResultCreatedResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MotivatorResultDeleteAll

> MotivatorResultDeleteAll(ctx, motivatorId).Execute()

Deletes all results on a motivator for the user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    motivatorId := int32(56) // int32 | The ID of the motivator to fetch users results

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MotivatorServiceApi.MotivatorResultDeleteAll(context.Background(), motivatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MotivatorServiceApi.MotivatorResultDeleteAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**motivatorId** | **int32** | The ID of the motivator to fetch users results | 

### Other Parameters

Other parameters are passed through a pointer to a apiMotivatorResultDeleteAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MotivatorSafetyNetAdd

> SafetyNetItem MotivatorSafetyNetAdd(ctx).SafetyNetItemContent(safetyNetItemContent).Execute()

Add safety net item to currently authenticated user safety net.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    safetyNetItemContent := *openapiclient.NewSafetyNetItemContent("Name_example", "Type_example", []string{"Strategies_example"}) // SafetyNetItemContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MotivatorServiceApi.MotivatorSafetyNetAdd(context.Background()).SafetyNetItemContent(safetyNetItemContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MotivatorServiceApi.MotivatorSafetyNetAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MotivatorSafetyNetAdd`: SafetyNetItem
    fmt.Fprintf(os.Stdout, "Response from `MotivatorServiceApi.MotivatorSafetyNetAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMotivatorSafetyNetAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **safetyNetItemContent** | [**SafetyNetItemContent**](SafetyNetItemContent.md) |  | 

### Return type

[**SafetyNetItem**](SafetyNetItem.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MotivatorSafetyNetDelete

> MotivatorSafetyNetDelete(ctx, safetyNetItemId).Execute()

Deletes an item from users safety net



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    safetyNetItemId := int32(5421) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MotivatorServiceApi.MotivatorSafetyNetDelete(context.Background(), safetyNetItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MotivatorServiceApi.MotivatorSafetyNetDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**safetyNetItemId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMotivatorSafetyNetDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MotivatorSafetyNetGet

> []SafetyNetItem MotivatorSafetyNetGet(ctx).Execute()

Get safety net for the current user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MotivatorServiceApi.MotivatorSafetyNetGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MotivatorServiceApi.MotivatorSafetyNetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MotivatorSafetyNetGet`: []SafetyNetItem
    fmt.Fprintf(os.Stdout, "Response from `MotivatorServiceApi.MotivatorSafetyNetGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMotivatorSafetyNetGetRequest struct via the builder pattern


### Return type

[**[]SafetyNetItem**](SafetyNetItem.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MotivatorSafetyNetReplace

> SafetyNetItem MotivatorSafetyNetReplace(ctx, safetyNetItemId).SafetyNetItemContent(safetyNetItemContent).Execute()

Replaces a safety net item with a new one



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    safetyNetItemId := int32(5421) // int32 | 
    safetyNetItemContent := *openapiclient.NewSafetyNetItemContent("Name_example", "Type_example", []string{"Strategies_example"}) // SafetyNetItemContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MotivatorServiceApi.MotivatorSafetyNetReplace(context.Background(), safetyNetItemId).SafetyNetItemContent(safetyNetItemContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MotivatorServiceApi.MotivatorSafetyNetReplace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MotivatorSafetyNetReplace`: SafetyNetItem
    fmt.Fprintf(os.Stdout, "Response from `MotivatorServiceApi.MotivatorSafetyNetReplace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**safetyNetItemId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMotivatorSafetyNetReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **safetyNetItemContent** | [**SafetyNetItemContent**](SafetyNetItemContent.md) |  | 

### Return type

[**SafetyNetItem**](SafetyNetItem.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

