# \MotivatorServiceApi

All URIs are relative to *https://echo.api.live.mindtastic.lol*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMotivatorUserId**](MotivatorServiceApi.md#DeleteMotivatorUserId) | **Delete** /motivator | Delete a motivator from user&#39;s motivators.
[**DeleteMotivatorUserIdResultMotivatorId**](MotivatorServiceApi.md#DeleteMotivatorUserIdResultMotivatorId) | **Delete** /motivator/result/{motivatorId} | Delete result from user&#39;s current motivator by timestamp.
[**GetMotivatorUserId**](MotivatorServiceApi.md#GetMotivatorUserId) | **Get** /motivator | Get current motivators by userId.
[**GetSafetynet**](MotivatorServiceApi.md#GetSafetynet) | **Get** /safetyNet | Get currently authenticated users safety-net
[**PostMotivatorUserId**](MotivatorServiceApi.md#PostMotivatorUserId) | **Post** /motivator | Add motivator to user&#39;s motivators.
[**PostMotivatorUserIdFeedbackMotivatorId**](MotivatorServiceApi.md#PostMotivatorUserIdFeedbackMotivatorId) | **Post** /motivator/result/{motivatorId} | Add a new result to user&#39;s current motivator.
[**SafetyNetPost**](MotivatorServiceApi.md#SafetyNetPost) | **Post** /safetyNet | Add safety net item to currently authenticated user safety net.



## DeleteMotivatorUserId

> DeleteMotivatorUserId(ctx).Execute()

Delete a motivator from user's motivators.



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
    resp, r, err := apiClient.MotivatorServiceApi.DeleteMotivatorUserId(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MotivatorServiceApi.DeleteMotivatorUserId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMotivatorUserIdRequest struct via the builder pattern


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


## DeleteMotivatorUserIdResultMotivatorId

> DeleteMotivatorUserIdResultMotivatorId(ctx, motivatorId).Execute()

Delete result from user's current motivator by timestamp.



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
    resp, r, err := apiClient.MotivatorServiceApi.DeleteMotivatorUserIdResultMotivatorId(context.Background(), motivatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MotivatorServiceApi.DeleteMotivatorUserIdResultMotivatorId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteMotivatorUserIdResultMotivatorIdRequest struct via the builder pattern


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


## GetMotivatorUserId

> []GetMotivatorUserId200ResponseInner GetMotivatorUserId(ctx).Execute()

Get current motivators by userId.



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
    resp, r, err := apiClient.MotivatorServiceApi.GetMotivatorUserId(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MotivatorServiceApi.GetMotivatorUserId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMotivatorUserId`: []GetMotivatorUserId200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `MotivatorServiceApi.GetMotivatorUserId`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMotivatorUserIdRequest struct via the builder pattern


### Return type

[**[]GetMotivatorUserId200ResponseInner**](GetMotivatorUserId200ResponseInner.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSafetynet

> []SafetyNetItem GetSafetynet(ctx).Execute()

Get currently authenticated users safety-net



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
    resp, r, err := apiClient.MotivatorServiceApi.GetSafetynet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MotivatorServiceApi.GetSafetynet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSafetynet`: []SafetyNetItem
    fmt.Fprintf(os.Stdout, "Response from `MotivatorServiceApi.GetSafetynet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSafetynetRequest struct via the builder pattern


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


## PostMotivatorUserId

> PostMotivatorUserId(ctx).Motivator(motivator).Execute()

Add motivator to user's motivators.



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
    motivator := *openapiclient.NewMotivator() // Motivator |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MotivatorServiceApi.PostMotivatorUserId(context.Background()).Motivator(motivator).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MotivatorServiceApi.PostMotivatorUserId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMotivatorUserIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **motivator** | [**Motivator**](Motivator.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMotivatorUserIdFeedbackMotivatorId

> PostMotivatorUserIdFeedbackMotivatorId(ctx, motivatorId).MotivatorResult(motivatorResult).Execute()

Add a new result to user's current motivator.



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
    motivatorResult := *openapiclient.NewMotivatorResult() // MotivatorResult |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MotivatorServiceApi.PostMotivatorUserIdFeedbackMotivatorId(context.Background(), motivatorId).MotivatorResult(motivatorResult).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MotivatorServiceApi.PostMotivatorUserIdFeedbackMotivatorId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostMotivatorUserIdFeedbackMotivatorIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **motivatorResult** | [**MotivatorResult**](MotivatorResult.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SafetyNetPost

> SafetyNetPost(ctx).SafetyNetItem(safetyNetItem).Execute()

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
    safetyNetItem := *openapiclient.NewSafetyNetItem("Name_example", "Type_example", []string{"Strategies_example"}) // SafetyNetItem |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MotivatorServiceApi.SafetyNetPost(context.Background()).SafetyNetItem(safetyNetItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MotivatorServiceApi.SafetyNetPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSafetyNetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **safetyNetItem** | [**SafetyNetItem**](SafetyNetItem.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

