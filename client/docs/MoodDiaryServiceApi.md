# \MoodDiaryServiceApi

All URIs are relative to *https://echo.api.live.mindtastic.lol*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiaryMoodCreate**](MoodDiaryServiceApi.md#DiaryMoodCreate) | **Post** /diary | Add mood to diary
[**DiaryMoodDelete**](MoodDiaryServiceApi.md#DiaryMoodDelete) | **Delete** /diary/{id} | Delete single mood diary entry
[**DiaryMoodGet**](MoodDiaryServiceApi.md#DiaryMoodGet) | **Get** /diary/{id} | Fetch single mood diary entry
[**DiaryMoodGetMany**](MoodDiaryServiceApi.md#DiaryMoodGetMany) | **Get** /diary | Get mood diary by userId.
[**DiaryMoodUpdate**](MoodDiaryServiceApi.md#DiaryMoodUpdate) | **Put** /diary/{id} | Update a single mood diary entry



## DiaryMoodCreate

> Mood DiaryMoodCreate(ctx).MoodData(moodData).Execute()

Add mood to diary



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
    moodData := *openapiclient.NewMoodData(time.Now(), "MoodType_example") // MoodData |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoodDiaryServiceApi.DiaryMoodCreate(context.Background()).MoodData(moodData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoodDiaryServiceApi.DiaryMoodCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DiaryMoodCreate`: Mood
    fmt.Fprintf(os.Stdout, "Response from `MoodDiaryServiceApi.DiaryMoodCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiaryMoodCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **moodData** | [**MoodData**](MoodData.md) |  | 

### Return type

[**Mood**](Mood.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiaryMoodDelete

> DiaryMoodDelete(ctx, id).Execute()

Delete single mood diary entry



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
    id := int32(2143) // int32 | ID of the mood diary entry the user wants to perform a request on

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoodDiaryServiceApi.DiaryMoodDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoodDiaryServiceApi.DiaryMoodDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the mood diary entry the user wants to perform a request on | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiaryMoodDeleteRequest struct via the builder pattern


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


## DiaryMoodGet

> Mood DiaryMoodGet(ctx, id).Execute()

Fetch single mood diary entry



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
    id := int32(2143) // int32 | ID of the mood diary entry the user wants to perform a request on

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoodDiaryServiceApi.DiaryMoodGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoodDiaryServiceApi.DiaryMoodGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DiaryMoodGet`: Mood
    fmt.Fprintf(os.Stdout, "Response from `MoodDiaryServiceApi.DiaryMoodGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the mood diary entry the user wants to perform a request on | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiaryMoodGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Mood**](Mood.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiaryMoodGetMany

> []Mood DiaryMoodGetMany(ctx).Execute()

Get mood diary by userId.



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
    resp, r, err := apiClient.MoodDiaryServiceApi.DiaryMoodGetMany(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoodDiaryServiceApi.DiaryMoodGetMany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DiaryMoodGetMany`: []Mood
    fmt.Fprintf(os.Stdout, "Response from `MoodDiaryServiceApi.DiaryMoodGetMany`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDiaryMoodGetManyRequest struct via the builder pattern


### Return type

[**[]Mood**](Mood.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiaryMoodUpdate

> DiaryMoodUpdate(ctx, id).MoodData(moodData).Execute()

Update a single mood diary entry



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
    id := int32(2143) // int32 | ID of the mood diary entry the user wants to perform a request on
    moodData := *openapiclient.NewMoodData(time.Now(), "MoodType_example") // MoodData |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoodDiaryServiceApi.DiaryMoodUpdate(context.Background(), id).MoodData(moodData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoodDiaryServiceApi.DiaryMoodUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the mood diary entry the user wants to perform a request on | 

### Other Parameters

Other parameters are passed through a pointer to a apiDiaryMoodUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moodData** | [**MoodData**](MoodData.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

