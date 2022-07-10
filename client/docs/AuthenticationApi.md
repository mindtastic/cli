# \AuthenticationApi

All URIs are relative to *https://echo.api.live.mindtastic.lol*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BouncerInitLoginApi**](AuthenticationApi.md#BouncerInitLoginApi) | **Get** /self-service/login/api | Initiate API login flow
[**BouncerInitLoginBrowser**](AuthenticationApi.md#BouncerInitLoginBrowser) | **Get** /self-service/login/browser | Initiate browser login flow
[**BouncerInitLogoutBrowser**](AuthenticationApi.md#BouncerInitLogoutBrowser) | **Get** /self-service/logout/browser | Initiate logout flow for browser
[**BouncerInitRegistrationApi**](AuthenticationApi.md#BouncerInitRegistrationApi) | **Get** /self-service/registration/api | Initate registration flow for APIs, Services, ...
[**BouncerInitRegistrationBrowser**](AuthenticationApi.md#BouncerInitRegistrationBrowser) | **Get** /self-service/registration/browser | Initialize Registration Flow for Browsers
[**BouncerLogin**](AuthenticationApi.md#BouncerLogin) | **Post** /self-service/login | Submit Login Flow
[**BouncerLogoutApi**](AuthenticationApi.md#BouncerLogoutApi) | **Delete** /self-service/logout/api | API user logout
[**BouncerLogoutBrowser**](AuthenticationApi.md#BouncerLogoutBrowser) | **Get** /self-service/logout | Submit Browser Logout
[**BouncerRegister**](AuthenticationApi.md#BouncerRegister) | **Post** /self-service/registration | Complete registration flow



## BouncerInitLoginApi

> ApiLoginFlow BouncerInitLoginApi(ctx).Refresh(refresh).XSessionToken(xSessionToken).Execute()

Initiate API login flow



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
    refresh := true // bool | Refresh a login.  If set to true, this will refresh an existing login session by asking the user to sign in again. This will reset the authenticated_at time of the session. (optional)
    xSessionToken := "xSessionToken_example" // string | The Session token of the Session performing the flow. Required when using the refresh parameter (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApi.BouncerInitLoginApi(context.Background()).Refresh(refresh).XSessionToken(xSessionToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.BouncerInitLoginApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BouncerInitLoginApi`: ApiLoginFlow
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.BouncerInitLoginApi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBouncerInitLoginApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refresh** | **bool** | Refresh a login.  If set to true, this will refresh an existing login session by asking the user to sign in again. This will reset the authenticated_at time of the session. | 
 **xSessionToken** | **string** | The Session token of the Session performing the flow. Required when using the refresh parameter | 

### Return type

[**ApiLoginFlow**](ApiLoginFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BouncerInitLoginBrowser

> BrowserLoginFlow BouncerInitLoginBrowser(ctx).ReturnTo(returnTo).Execute()

Initiate browser login flow



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
    returnTo := "returnTo_example" // string | The URL to return the browser after the flow was completed. For Kopfsachen you might never need that option. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApi.BouncerInitLoginBrowser(context.Background()).ReturnTo(returnTo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.BouncerInitLoginBrowser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BouncerInitLoginBrowser`: BrowserLoginFlow
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.BouncerInitLoginBrowser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBouncerInitLoginBrowserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **returnTo** | **string** | The URL to return the browser after the flow was completed. For Kopfsachen you might never need that option. | 

### Return type

[**BrowserLoginFlow**](BrowserLoginFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BouncerInitLogoutBrowser

> BrowserLogoutUrl BouncerInitLogoutBrowser(ctx).Cookie(cookie).Execute()

Initiate logout flow for browser



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
    cookie := "cookie_example" // string | HTTP Cookies If you call this endpoint from a backend, please include the original Cookie header in the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApi.BouncerInitLogoutBrowser(context.Background()).Cookie(cookie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.BouncerInitLogoutBrowser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BouncerInitLogoutBrowser`: BrowserLogoutUrl
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.BouncerInitLogoutBrowser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBouncerInitLogoutBrowserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cookie** | **string** | HTTP Cookies If you call this endpoint from a backend, please include the original Cookie header in the request. | 

### Return type

[**BrowserLogoutUrl**](BrowserLogoutUrl.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BouncerInitRegistrationApi

> ApiRegistrationFlow BouncerInitRegistrationApi(ctx).Execute()

Initate registration flow for APIs, Services, ...



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
    resp, r, err := apiClient.AuthenticationApi.BouncerInitRegistrationApi(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.BouncerInitRegistrationApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BouncerInitRegistrationApi`: ApiRegistrationFlow
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.BouncerInitRegistrationApi`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBouncerInitRegistrationApiRequest struct via the builder pattern


### Return type

[**ApiRegistrationFlow**](ApiRegistrationFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BouncerInitRegistrationBrowser

> BrowserRegistrationFlow BouncerInitRegistrationBrowser(ctx).ReturnTo(returnTo).Execute()

Initialize Registration Flow for Browsers



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
    returnTo := "returnTo_example" // string | The URL to return the browser after the flow was completed. For Kopfsachen you might never need that option. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApi.BouncerInitRegistrationBrowser(context.Background()).ReturnTo(returnTo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.BouncerInitRegistrationBrowser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BouncerInitRegistrationBrowser`: BrowserRegistrationFlow
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.BouncerInitRegistrationBrowser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBouncerInitRegistrationBrowserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **returnTo** | **string** | The URL to return the browser after the flow was completed. For Kopfsachen you might never need that option. | 

### Return type

[**BrowserRegistrationFlow**](BrowserRegistrationFlow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BouncerLogin

> Session BouncerLogin(ctx).Flow(flow).BouncerRegisterRequest(bouncerRegisterRequest).Execute()

Submit Login Flow



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
    flow := "flow_example" // string | The Login flow ID as an URL Query parameters
    bouncerRegisterRequest := *openapiclient.NewBouncerRegisterRequest() // BouncerRegisterRequest | Providing `csrf_token` is only required for browser flows. For API flows, request body can left empty. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApi.BouncerLogin(context.Background()).Flow(flow).BouncerRegisterRequest(bouncerRegisterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.BouncerLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BouncerLogin`: Session
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.BouncerLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBouncerLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flow** | **string** | The Login flow ID as an URL Query parameters | 
 **bouncerRegisterRequest** | [**BouncerRegisterRequest**](BouncerRegisterRequest.md) | Providing &#x60;csrf_token&#x60; is only required for browser flows. For API flows, request body can left empty. | 

### Return type

[**Session**](Session.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BouncerLogoutApi

> BouncerLogoutApi(ctx).SubmitSelfServiceLogoutFlowWithoutBrowserBody(submitSelfServiceLogoutFlowWithoutBrowserBody).Execute()

API user logout



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
    submitSelfServiceLogoutFlowWithoutBrowserBody := *openapiclient.NewSubmitSelfServiceLogoutFlowWithoutBrowserBody("SessionToken_example") // SubmitSelfServiceLogoutFlowWithoutBrowserBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApi.BouncerLogoutApi(context.Background()).SubmitSelfServiceLogoutFlowWithoutBrowserBody(submitSelfServiceLogoutFlowWithoutBrowserBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.BouncerLogoutApi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBouncerLogoutApiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitSelfServiceLogoutFlowWithoutBrowserBody** | [**SubmitSelfServiceLogoutFlowWithoutBrowserBody**](SubmitSelfServiceLogoutFlowWithoutBrowserBody.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BouncerLogoutBrowser

> BouncerLogoutBrowser(ctx).Token(token).ReturnTo(returnTo).Execute()

Submit Browser Logout



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
    token := "token_example" // string | A Valid Logout Token  If you do not have a logout token because you only have a session cookie, call `/self-service/logout/browser` (optional)
    returnTo := "returnTo_example" // string | The URL to return to after the logout was completed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApi.BouncerLogoutBrowser(context.Background()).Token(token).ReturnTo(returnTo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.BouncerLogoutBrowser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBouncerLogoutBrowserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | A Valid Logout Token  If you do not have a logout token because you only have a session cookie, call &#x60;/self-service/logout/browser&#x60; | 
 **returnTo** | **string** | The URL to return to after the logout was completed. | 

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


## BouncerRegister

> RegistrationWithFlowIdResponse BouncerRegister(ctx).Flow(flow).Cookie(cookie).BouncerRegisterRequest(bouncerRegisterRequest).Execute()

Complete registration flow



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
    flow := "flow_example" // string | The Registration flow ID as an URL Query parameters
    cookie := "cookie_example" // string | When using the browser flow, the HTTP cookie (encodes session and CSRF token) must be present (optional)
    bouncerRegisterRequest := *openapiclient.NewBouncerRegisterRequest() // BouncerRegisterRequest | Providing `csrf_token` is only required for browser flows. For API flows, request body can left empty. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationApi.BouncerRegister(context.Background()).Flow(flow).Cookie(cookie).BouncerRegisterRequest(bouncerRegisterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.BouncerRegister``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BouncerRegister`: RegistrationWithFlowIdResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationApi.BouncerRegister`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBouncerRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flow** | **string** | The Registration flow ID as an URL Query parameters | 
 **cookie** | **string** | When using the browser flow, the HTTP cookie (encodes session and CSRF token) must be present | 
 **bouncerRegisterRequest** | [**BouncerRegisterRequest**](BouncerRegisterRequest.md) | Providing &#x60;csrf_token&#x60; is only required for browser flows. For API flows, request body can left empty. | 

### Return type

[**RegistrationWithFlowIdResponse**](RegistrationWithFlowIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

