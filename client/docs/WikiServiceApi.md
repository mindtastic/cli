# \WikiServiceApi

All URIs are relative to *https://echo.api.live.mindtastic.lol*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WikiCreateEntry**](WikiServiceApi.md#WikiCreateEntry) | **Post** /admin/wiki | Add new wiki entries
[**WikiDeleteEntry**](WikiServiceApi.md#WikiDeleteEntry) | **Delete** /admin/wiki/{articleId} | Delete wiki entry by id.
[**WikiGet**](WikiServiceApi.md#WikiGet) | **Get** /wiki/{articleId} | Fetch single wiki entry
[**WikiList**](WikiServiceApi.md#WikiList) | **Get** /wiki | Fetch wiki entries



## WikiCreateEntry

> WikiEntry WikiCreateEntry(ctx).CreateEntryBody(createEntryBody).Execute()

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
    createEntryBody := *openapiclient.NewCreateEntryBody() // CreateEntryBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WikiServiceApi.WikiCreateEntry(context.Background()).CreateEntryBody(createEntryBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WikiServiceApi.WikiCreateEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WikiCreateEntry`: WikiEntry
    fmt.Fprintf(os.Stdout, "Response from `WikiServiceApi.WikiCreateEntry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWikiCreateEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEntryBody** | [**CreateEntryBody**](CreateEntryBody.md) |  | 

### Return type

[**WikiEntry**](WikiEntry.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WikiDeleteEntry

> WikiDeleteEntry(ctx, articleId).Execute()

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
    articleId := "62bda938ada05ab5872edea7" // string | ID of a wiki entry

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WikiServiceApi.WikiDeleteEntry(context.Background(), articleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WikiServiceApi.WikiDeleteEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **string** | ID of a wiki entry | 

### Other Parameters

Other parameters are passed through a pointer to a apiWikiDeleteEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WikiGet

> WikiEntry WikiGet(ctx, articleId).Execute()

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
    articleId := "62bda938ada05ab5872edea7" // string | ID of a wiki entry

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WikiServiceApi.WikiGet(context.Background(), articleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WikiServiceApi.WikiGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WikiGet`: WikiEntry
    fmt.Fprintf(os.Stdout, "Response from `WikiServiceApi.WikiGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**articleId** | **string** | ID of a wiki entry | 

### Other Parameters

Other parameters are passed through a pointer to a apiWikiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WikiEntry**](WikiEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WikiList

> WikiList200Response WikiList(ctx).Query(query).Limit(limit).Offset(offset).WithContent(withContent).Execute()

Fetch wiki entries



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
    query := "query_example" // string |  (optional)
    limit := int32(56) // int32 | The maximum number of articles the response will include (Use together with `offset` for api pagination). (optional) (default to 20)
    offset := int32(56) // int32 | The number of articles that will we skipped in response (Use together wuth `limit` for api pagination). (optional) (default to 0)
    withContent := true // bool | If the with content attribute is not `true`, the response will not return the content of wiki articles, but only the title and metadata attributes.  *Note*: This behaviour is not implemented yet, but It will be later. This endpoint can be used to download an index of all articles, without having to download the article content. The metadata might also contain a `preview` field with the first few character of each article in a later persion of the spec.   As adding the parameter later would be a breaking spec, always send the with_content parameter if you actually want to fetch article contents. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WikiServiceApi.WikiList(context.Background()).Query(query).Limit(limit).Offset(offset).WithContent(withContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WikiServiceApi.WikiList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WikiList`: WikiList200Response
    fmt.Fprintf(os.Stdout, "Response from `WikiServiceApi.WikiList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWikiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
 **limit** | **int32** | The maximum number of articles the response will include (Use together with &#x60;offset&#x60; for api pagination). | [default to 20]
 **offset** | **int32** | The number of articles that will we skipped in response (Use together wuth &#x60;limit&#x60; for api pagination). | [default to 0]
 **withContent** | **bool** | If the with content attribute is not &#x60;true&#x60;, the response will not return the content of wiki articles, but only the title and metadata attributes.  *Note*: This behaviour is not implemented yet, but It will be later. This endpoint can be used to download an index of all articles, without having to download the article content. The metadata might also contain a &#x60;preview&#x60; field with the first few character of each article in a later persion of the spec.   As adding the parameter later would be a breaking spec, always send the with_content parameter if you actually want to fetch article contents. | [default to false]

### Return type

[**WikiList200Response**](WikiList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

