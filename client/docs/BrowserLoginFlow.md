# BrowserLoginFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to [**IdentityCredentialsType**](IdentityCredentialsType.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ExpiresAt** | **time.Time** |  | 
**Id** | **string** |  | 
**IssuedAt** | **time.Time** | IssuedAt is the time (UTC) when the flow started. | 
**Refresh** | Pointer to **bool** | Refresh stores whether this login flow should enforce re-authentication. | [optional] 
**RequestUrl** | **string** | RequestURL is the initial URL that was requested from Ory Kratos. It can be used to forward information contained in the URL&#39;s path or query for example. | 
**RequestedAal** | Pointer to [**AuthenticatorAssuranceLevel**](AuthenticatorAssuranceLevel.md) |  | [optional] 
**ReturnTo** | Pointer to **string** | ReturnTo contains the requested return_to URL. | [optional] 
**Type** | **string** | The flow type can either be &#x60;api&#x60; or &#x60;browser&#x60;. | 
**Ui** | [**UiContainer**](UiContainer.md) |  | 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is a helper struct field for gobuffalo.pop. | [optional] 

## Methods

### NewBrowserLoginFlow

`func NewBrowserLoginFlow(expiresAt time.Time, id string, issuedAt time.Time, requestUrl string, type_ string, ui UiContainer, ) *BrowserLoginFlow`

NewBrowserLoginFlow instantiates a new BrowserLoginFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrowserLoginFlowWithDefaults

`func NewBrowserLoginFlowWithDefaults() *BrowserLoginFlow`

NewBrowserLoginFlowWithDefaults instantiates a new BrowserLoginFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *BrowserLoginFlow) GetActive() IdentityCredentialsType`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *BrowserLoginFlow) GetActiveOk() (*IdentityCredentialsType, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *BrowserLoginFlow) SetActive(v IdentityCredentialsType)`

SetActive sets Active field to given value.

### HasActive

`func (o *BrowserLoginFlow) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BrowserLoginFlow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BrowserLoginFlow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BrowserLoginFlow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BrowserLoginFlow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *BrowserLoginFlow) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *BrowserLoginFlow) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *BrowserLoginFlow) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetId

`func (o *BrowserLoginFlow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BrowserLoginFlow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BrowserLoginFlow) SetId(v string)`

SetId sets Id field to given value.


### GetIssuedAt

`func (o *BrowserLoginFlow) GetIssuedAt() time.Time`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *BrowserLoginFlow) GetIssuedAtOk() (*time.Time, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *BrowserLoginFlow) SetIssuedAt(v time.Time)`

SetIssuedAt sets IssuedAt field to given value.


### GetRefresh

`func (o *BrowserLoginFlow) GetRefresh() bool`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *BrowserLoginFlow) GetRefreshOk() (*bool, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *BrowserLoginFlow) SetRefresh(v bool)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *BrowserLoginFlow) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetRequestUrl

`func (o *BrowserLoginFlow) GetRequestUrl() string`

GetRequestUrl returns the RequestUrl field if non-nil, zero value otherwise.

### GetRequestUrlOk

`func (o *BrowserLoginFlow) GetRequestUrlOk() (*string, bool)`

GetRequestUrlOk returns a tuple with the RequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUrl

`func (o *BrowserLoginFlow) SetRequestUrl(v string)`

SetRequestUrl sets RequestUrl field to given value.


### GetRequestedAal

`func (o *BrowserLoginFlow) GetRequestedAal() AuthenticatorAssuranceLevel`

GetRequestedAal returns the RequestedAal field if non-nil, zero value otherwise.

### GetRequestedAalOk

`func (o *BrowserLoginFlow) GetRequestedAalOk() (*AuthenticatorAssuranceLevel, bool)`

GetRequestedAalOk returns a tuple with the RequestedAal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAal

`func (o *BrowserLoginFlow) SetRequestedAal(v AuthenticatorAssuranceLevel)`

SetRequestedAal sets RequestedAal field to given value.

### HasRequestedAal

`func (o *BrowserLoginFlow) HasRequestedAal() bool`

HasRequestedAal returns a boolean if a field has been set.

### GetReturnTo

`func (o *BrowserLoginFlow) GetReturnTo() string`

GetReturnTo returns the ReturnTo field if non-nil, zero value otherwise.

### GetReturnToOk

`func (o *BrowserLoginFlow) GetReturnToOk() (*string, bool)`

GetReturnToOk returns a tuple with the ReturnTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnTo

`func (o *BrowserLoginFlow) SetReturnTo(v string)`

SetReturnTo sets ReturnTo field to given value.

### HasReturnTo

`func (o *BrowserLoginFlow) HasReturnTo() bool`

HasReturnTo returns a boolean if a field has been set.

### GetType

`func (o *BrowserLoginFlow) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BrowserLoginFlow) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BrowserLoginFlow) SetType(v string)`

SetType sets Type field to given value.


### GetUi

`func (o *BrowserLoginFlow) GetUi() UiContainer`

GetUi returns the Ui field if non-nil, zero value otherwise.

### GetUiOk

`func (o *BrowserLoginFlow) GetUiOk() (*UiContainer, bool)`

GetUiOk returns a tuple with the Ui field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUi

`func (o *BrowserLoginFlow) SetUi(v UiContainer)`

SetUi sets Ui field to given value.


### GetUpdatedAt

`func (o *BrowserLoginFlow) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BrowserLoginFlow) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BrowserLoginFlow) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BrowserLoginFlow) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


