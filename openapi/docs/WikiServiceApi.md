# \WikiServiceApi

All URIs are relative to *https://echo.api.live.mindtastic.lol*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWiki**](WikiServiceApi.md#DeleteWiki) | **Delete** /wiki/{wikiEntryId}/admin | Delete wiki entry by id.
[**GetWiki**](WikiServiceApi.md#GetWiki) | **Get** /wiki | Get all available wiki entries
[**PostWiki**](WikiServiceApi.md#PostWiki) | **Post** /wiki/admin | Add new wiki entries
[**WikiWikiEntryIdGet**](WikiServiceApi.md#WikiWikiEntryIdGet) | **Get** /wiki/{wikiEntryId} | Fetch single wiki entry



## DeleteWiki

> DeleteWiki(ctx, wikiEntryId).Execute()

Delete wiki entry by id.

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
    wikiEntryId := "wikiEntryId_example" // string | ID of a wiki entry

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WikiServiceApi.DeleteWiki(context.Background(), wikiEntryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WikiServiceApi.DeleteWiki``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wikiEntryId** | **string** | ID of a wiki entry | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWikiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWiki

> []Entry GetWiki(ctx).Execute()

Get all available wiki entries

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
    resp, r, err := apiClient.WikiServiceApi.GetWiki(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WikiServiceApi.GetWiki``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWiki`: []Entry
    fmt.Fprintf(os.Stdout, "Response from `WikiServiceApi.GetWiki`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWikiRequest struct via the builder pattern


### Return type

[**[]Entry**](Entry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWiki

> InsertionResponse PostWiki(ctx).Entry(entry).Execute()

Add new wiki entries



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
    entry := []openapiclient.Entry{*openapiclient.NewEntry()} // []Entry |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WikiServiceApi.PostWiki(context.Background()).Entry(entry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WikiServiceApi.PostWiki``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostWiki`: InsertionResponse
    fmt.Fprintf(os.Stdout, "Response from `WikiServiceApi.PostWiki`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostWikiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entry** | [**[]Entry**](Entry.md) |  | 

### Return type

[**InsertionResponse**](InsertionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WikiWikiEntryIdGet

> Entry WikiWikiEntryIdGet(ctx, wikiEntryId).Execute()

Fetch single wiki entry

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
    wikiEntryId := "wikiEntryId_example" // string | ID of a wiki entry

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WikiServiceApi.WikiWikiEntryIdGet(context.Background(), wikiEntryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WikiServiceApi.WikiWikiEntryIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WikiWikiEntryIdGet`: Entry
    fmt.Fprintf(os.Stdout, "Response from `WikiServiceApi.WikiWikiEntryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wikiEntryId** | **string** | ID of a wiki entry | 

### Other Parameters

Other parameters are passed through a pointer to a apiWikiWikiEntryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Entry**](Entry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

