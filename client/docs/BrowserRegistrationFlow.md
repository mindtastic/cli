# BrowserRegistrationFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to [**IdentityCredentialsType**](IdentityCredentialsType.md) |  | [optional] 
**ExpiresAt** | **time.Time** | ExpiresAt is the time (UTC) when the flow expires. If the user still wishes to log in, a new flow has to be initiated. | 
**Id** | **string** |  | 
**IssuedAt** | **time.Time** | IssuedAt is the time (UTC) when the flow occurred. | 
**RequestUrl** | **string** | RequestURL is the initial URL that was requested from Ory Kratos. It can be used to forward information contained in the URL&#39;s path or query for example. | 
**ReturnTo** | Pointer to **string** | ReturnTo contains the requested return_to URL. | [optional] 
**Type** | **string** | The flow type can either be &#x60;api&#x60; or &#x60;browser&#x60;. | 
**Ui** | [**UiContainer**](UiContainer.md) |  | 

## Methods

### NewBrowserRegistrationFlow

`func NewBrowserRegistrationFlow(expiresAt time.Time, id string, issuedAt time.Time, requestUrl string, type_ string, ui UiContainer, ) *BrowserRegistrationFlow`

NewBrowserRegistrationFlow instantiates a new BrowserRegistrationFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrowserRegistrationFlowWithDefaults

`func NewBrowserRegistrationFlowWithDefaults() *BrowserRegistrationFlow`

NewBrowserRegistrationFlowWithDefaults instantiates a new BrowserRegistrationFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *BrowserRegistrationFlow) GetActive() IdentityCredentialsType`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *BrowserRegistrationFlow) GetActiveOk() (*IdentityCredentialsType, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *BrowserRegistrationFlow) SetActive(v IdentityCredentialsType)`

SetActive sets Active field to given value.

### HasActive

`func (o *BrowserRegistrationFlow) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetExpiresAt

`func (o *BrowserRegistrationFlow) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *BrowserRegistrationFlow) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *BrowserRegistrationFlow) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetId

`func (o *BrowserRegistrationFlow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BrowserRegistrationFlow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BrowserRegistrationFlow) SetId(v string)`

SetId sets Id field to given value.


### GetIssuedAt

`func (o *BrowserRegistrationFlow) GetIssuedAt() time.Time`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *BrowserRegistrationFlow) GetIssuedAtOk() (*time.Time, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *BrowserRegistrationFlow) SetIssuedAt(v time.Time)`

SetIssuedAt sets IssuedAt field to given value.


### GetRequestUrl

`func (o *BrowserRegistrationFlow) GetRequestUrl() string`

GetRequestUrl returns the RequestUrl field if non-nil, zero value otherwise.

### GetRequestUrlOk

`func (o *BrowserRegistrationFlow) GetRequestUrlOk() (*string, bool)`

GetRequestUrlOk returns a tuple with the RequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUrl

`func (o *BrowserRegistrationFlow) SetRequestUrl(v string)`

SetRequestUrl sets RequestUrl field to given value.


### GetReturnTo

`func (o *BrowserRegistrationFlow) GetReturnTo() string`

GetReturnTo returns the ReturnTo field if non-nil, zero value otherwise.

### GetReturnToOk

`func (o *BrowserRegistrationFlow) GetReturnToOk() (*string, bool)`

GetReturnToOk returns a tuple with the ReturnTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnTo

`func (o *BrowserRegistrationFlow) SetReturnTo(v string)`

SetReturnTo sets ReturnTo field to given value.

### HasReturnTo

`func (o *BrowserRegistrationFlow) HasReturnTo() bool`

HasReturnTo returns a boolean if a field has been set.

### GetType

`func (o *BrowserRegistrationFlow) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BrowserRegistrationFlow) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BrowserRegistrationFlow) SetType(v string)`

SetType sets Type field to given value.


### GetUi

`func (o *BrowserRegistrationFlow) GetUi() UiContainer`

GetUi returns the Ui field if non-nil, zero value otherwise.

### GetUiOk

`func (o *BrowserRegistrationFlow) GetUiOk() (*UiContainer, bool)`

GetUiOk returns a tuple with the Ui field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUi

`func (o *BrowserRegistrationFlow) SetUi(v UiContainer)`

SetUi sets Ui field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


