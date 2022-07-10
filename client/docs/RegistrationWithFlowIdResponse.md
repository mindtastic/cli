# RegistrationWithFlowIdResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | Pointer to [**Identity**](Identity.md) |  | [optional] 
**Session** | Pointer to [**Session1**](Session1.md) |  | [optional] 
**SessionToken** | Pointer to **string** | The Session Token  A session token is equivalent to a session cookie, but it can be sent in the HTTP Authorization  Header: Authorization: bearer ${session-token} The session token is only issued for API flows, not for Browser flows! | [optional] 

## Methods

### NewRegistrationWithFlowIdResponse

`func NewRegistrationWithFlowIdResponse() *RegistrationWithFlowIdResponse`

NewRegistrationWithFlowIdResponse instantiates a new RegistrationWithFlowIdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationWithFlowIdResponseWithDefaults

`func NewRegistrationWithFlowIdResponseWithDefaults() *RegistrationWithFlowIdResponse`

NewRegistrationWithFlowIdResponseWithDefaults instantiates a new RegistrationWithFlowIdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *RegistrationWithFlowIdResponse) GetIdentity() Identity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *RegistrationWithFlowIdResponse) GetIdentityOk() (*Identity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *RegistrationWithFlowIdResponse) SetIdentity(v Identity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *RegistrationWithFlowIdResponse) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetSession

`func (o *RegistrationWithFlowIdResponse) GetSession() Session1`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *RegistrationWithFlowIdResponse) GetSessionOk() (*Session1, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *RegistrationWithFlowIdResponse) SetSession(v Session1)`

SetSession sets Session field to given value.

### HasSession

`func (o *RegistrationWithFlowIdResponse) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetSessionToken

`func (o *RegistrationWithFlowIdResponse) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *RegistrationWithFlowIdResponse) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *RegistrationWithFlowIdResponse) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.

### HasSessionToken

`func (o *RegistrationWithFlowIdResponse) HasSessionToken() bool`

HasSessionToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


