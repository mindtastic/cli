# \MoodDiaryServiceApi

All URIs are relative to *https://echo.api.live.mindtastic.lol*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMood**](MoodDiaryServiceApi.md#AddMood) | **Post** /diary | Add mood to diary
[**DeleteMoodDiary**](MoodDiaryServiceApi.md#DeleteMoodDiary) | **Delete** /diary | Delete current user&#39;s mood diary.
[**GetMoodDiary**](MoodDiaryServiceApi.md#GetMoodDiary) | **Get** /diary | Get mood diary by userId.



## AddMood

> AddMood(ctx).Mood(mood).Execute()

Add mood to diary



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
    mood := *openapiclient.NewMood() // Mood |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoodDiaryServiceApi.AddMood(context.Background()).Mood(mood).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoodDiaryServiceApi.AddMood``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddMoodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mood** | [**Mood**](Mood.md) |  | 

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


## DeleteMoodDiary

> DeleteMoodDiary(ctx).Execute()

Delete current user's mood diary.



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
    resp, r, err := apiClient.MoodDiaryServiceApi.DeleteMoodDiary(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoodDiaryServiceApi.DeleteMoodDiary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMoodDiaryRequest struct via the builder pattern


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


## GetMoodDiary

> []Mood GetMoodDiary(ctx).Execute()

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
    resp, r, err := apiClient.MoodDiaryServiceApi.GetMoodDiary(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoodDiaryServiceApi.GetMoodDiary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMoodDiary`: []Mood
    fmt.Fprintf(os.Stdout, "Response from `MoodDiaryServiceApi.GetMoodDiary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMoodDiaryRequest struct via the builder pattern


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

