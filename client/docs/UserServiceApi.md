# \UserServiceApi

All URIs are relative to *https://echo.api.live.mindtastic.lol*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserAdminCreate**](UserServiceApi.md#UserAdminCreate) | **Post** /admin/user | Create user.
[**UserAdminGetAll**](UserServiceApi.md#UserAdminGetAll) | **Get** /admin/user | Get all users.
[**UserDataDelete**](UserServiceApi.md#UserDataDelete) | **Delete** /user | Delete data of the authenticated user
[**UserDataGet**](UserServiceApi.md#UserDataGet) | **Get** /user | Get data of current user
[**UserDataUpdate**](UserServiceApi.md#UserDataUpdate) | **Put** /user | Update data of the authenticated user



## UserAdminCreate

> UserAdminCreate(ctx).XUserId(xUserId).UpdateUserModel(updateUserModel).Execute()

Create user.



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
    xUserId := "7f036688-36c2-4d73-80c7-a820fcf156a6" // string | The service specific user id of the user to create.  **Caution:** The user id MUST NOT be the users account key nor a authentication session token.
    updateUserModel := *openapiclient.NewUpdateUserModel() // UpdateUserModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserServiceApi.UserAdminCreate(context.Background()).XUserId(xUserId).UpdateUserModel(updateUserModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserServiceApi.UserAdminCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserAdminCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xUserId** | **string** | The service specific user id of the user to create.  **Caution:** The user id MUST NOT be the users account key nor a authentication session token. | 
 **updateUserModel** | [**UpdateUserModel**](UpdateUserModel.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserAdminGetAll

> []UserDataResponse UserAdminGetAll(ctx).UpdateUserModel(updateUserModel).Execute()

Get all users.



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
    updateUserModel := *openapiclient.NewUpdateUserModel() // UpdateUserModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserServiceApi.UserAdminGetAll(context.Background()).UpdateUserModel(updateUserModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserServiceApi.UserAdminGetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserAdminGetAll`: []UserDataResponse
    fmt.Fprintf(os.Stdout, "Response from `UserServiceApi.UserAdminGetAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserAdminGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateUserModel** | [**UpdateUserModel**](UpdateUserModel.md) |  | 

### Return type

[**[]UserDataResponse**](UserDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserDataDelete

> UserDataDelete(ctx).Execute()

Delete data of the authenticated user



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
    resp, r, err := apiClient.UserServiceApi.UserDataDelete(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserServiceApi.UserDataDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserDataDeleteRequest struct via the builder pattern


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


## UserDataGet

> UserDataResponse UserDataGet(ctx).Execute()

Get data of current user



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
    resp, r, err := apiClient.UserServiceApi.UserDataGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserServiceApi.UserDataGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserDataGet`: UserDataResponse
    fmt.Fprintf(os.Stdout, "Response from `UserServiceApi.UserDataGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserDataGetRequest struct via the builder pattern


### Return type

[**UserDataResponse**](UserDataResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserDataUpdate

> UserDataResponse UserDataUpdate(ctx).UpdateUserModel(updateUserModel).Execute()

Update data of the authenticated user

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
    updateUserModel := *openapiclient.NewUpdateUserModel() // UpdateUserModel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserServiceApi.UserDataUpdate(context.Background()).UpdateUserModel(updateUserModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserServiceApi.UserDataUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserDataUpdate`: UserDataResponse
    fmt.Fprintf(os.Stdout, "Response from `UserServiceApi.UserDataUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserDataUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateUserModel** | [**UpdateUserModel**](UpdateUserModel.md) |  | 

### Return type

[**UserDataResponse**](UserDataResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

