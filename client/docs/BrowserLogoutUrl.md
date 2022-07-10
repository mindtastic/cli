# BrowserLogoutUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogoutToken** | **string** | LogoutToken can be used to perform logout using AJAX. | 
**LogoutUrl** | **string** | LogoutURL can be opened in a browser to sign the user out. format: uri | 

## Methods

### NewBrowserLogoutUrl

`func NewBrowserLogoutUrl(logoutToken string, logoutUrl string, ) *BrowserLogoutUrl`

NewBrowserLogoutUrl instantiates a new BrowserLogoutUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrowserLogoutUrlWithDefaults

`func NewBrowserLogoutUrlWithDefaults() *BrowserLogoutUrl`

NewBrowserLogoutUrlWithDefaults instantiates a new BrowserLogoutUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogoutToken

`func (o *BrowserLogoutUrl) GetLogoutToken() string`

GetLogoutToken returns the LogoutToken field if non-nil, zero value otherwise.

### GetLogoutTokenOk

`func (o *BrowserLogoutUrl) GetLogoutTokenOk() (*string, bool)`

GetLogoutTokenOk returns a tuple with the LogoutToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutToken

`func (o *BrowserLogoutUrl) SetLogoutToken(v string)`

SetLogoutToken sets LogoutToken field to given value.


### GetLogoutUrl

`func (o *BrowserLogoutUrl) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *BrowserLogoutUrl) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *BrowserLogoutUrl) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


