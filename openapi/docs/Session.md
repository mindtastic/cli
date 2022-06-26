# Session

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionToken** | Pointer to **string** |  | [optional] 
**Session** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSession

`func NewSession() *Session`

NewSession instantiates a new Session object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionWithDefaults

`func NewSessionWithDefaults() *Session`

NewSessionWithDefaults instantiates a new Session object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionToken

`func (o *Session) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *Session) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *Session) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.

### HasSessionToken

`func (o *Session) HasSessionToken() bool`

HasSessionToken returns a boolean if a field has been set.

### GetSession

`func (o *Session) GetSession() map[string]interface{}`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *Session) GetSessionOk() (*map[string]interface{}, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *Session) SetSession(v map[string]interface{})`

SetSession sets Session field to given value.

### HasSession

`func (o *Session) HasSession() bool`

HasSession returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


