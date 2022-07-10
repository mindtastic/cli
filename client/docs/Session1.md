# Session1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Active** | Pointer to **bool** | Active state. If false the session is no longer active. | [optional] 
**ExpiresAt** | Pointer to **time.Time** | The Session Expiry When this session expires at. | [optional] 
**AuthenticatedAt** | Pointer to **time.Time** | The Session Authentication Timestamp When this session was authenticated at. If multi-factor authentication was used this is the time when the last factor was authenticated (e.g. the TOTP code challenge was completed). | [optional] 
**AuthenticatorAssuranceLevel** | Pointer to **string** | internal value | [optional] 
**AuthenticationMethods** | Pointer to [**[]SessionAuthenticationMethod**](SessionAuthenticationMethod.md) |  | [optional] 
**IssuedAt** | Pointer to **time.Time** | The Session Issuance Timestamp When this session was issued at. Usually equal or close to &#x60;authenticated_at&#x60;. | [optional] 
**Identity** | [**Identity**](Identity.md) |  | 

## Methods

### NewSession1

`func NewSession1(id string, identity Identity, ) *Session1`

NewSession1 instantiates a new Session1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSession1WithDefaults

`func NewSession1WithDefaults() *Session1`

NewSession1WithDefaults instantiates a new Session1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Session1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Session1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Session1) SetId(v string)`

SetId sets Id field to given value.


### GetActive

`func (o *Session1) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Session1) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Session1) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Session1) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetExpiresAt

`func (o *Session1) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Session1) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Session1) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *Session1) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetAuthenticatedAt

`func (o *Session1) GetAuthenticatedAt() time.Time`

GetAuthenticatedAt returns the AuthenticatedAt field if non-nil, zero value otherwise.

### GetAuthenticatedAtOk

`func (o *Session1) GetAuthenticatedAtOk() (*time.Time, bool)`

GetAuthenticatedAtOk returns a tuple with the AuthenticatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatedAt

`func (o *Session1) SetAuthenticatedAt(v time.Time)`

SetAuthenticatedAt sets AuthenticatedAt field to given value.

### HasAuthenticatedAt

`func (o *Session1) HasAuthenticatedAt() bool`

HasAuthenticatedAt returns a boolean if a field has been set.

### GetAuthenticatorAssuranceLevel

`func (o *Session1) GetAuthenticatorAssuranceLevel() string`

GetAuthenticatorAssuranceLevel returns the AuthenticatorAssuranceLevel field if non-nil, zero value otherwise.

### GetAuthenticatorAssuranceLevelOk

`func (o *Session1) GetAuthenticatorAssuranceLevelOk() (*string, bool)`

GetAuthenticatorAssuranceLevelOk returns a tuple with the AuthenticatorAssuranceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorAssuranceLevel

`func (o *Session1) SetAuthenticatorAssuranceLevel(v string)`

SetAuthenticatorAssuranceLevel sets AuthenticatorAssuranceLevel field to given value.

### HasAuthenticatorAssuranceLevel

`func (o *Session1) HasAuthenticatorAssuranceLevel() bool`

HasAuthenticatorAssuranceLevel returns a boolean if a field has been set.

### GetAuthenticationMethods

`func (o *Session1) GetAuthenticationMethods() []SessionAuthenticationMethod`

GetAuthenticationMethods returns the AuthenticationMethods field if non-nil, zero value otherwise.

### GetAuthenticationMethodsOk

`func (o *Session1) GetAuthenticationMethodsOk() (*[]SessionAuthenticationMethod, bool)`

GetAuthenticationMethodsOk returns a tuple with the AuthenticationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethods

`func (o *Session1) SetAuthenticationMethods(v []SessionAuthenticationMethod)`

SetAuthenticationMethods sets AuthenticationMethods field to given value.

### HasAuthenticationMethods

`func (o *Session1) HasAuthenticationMethods() bool`

HasAuthenticationMethods returns a boolean if a field has been set.

### GetIssuedAt

`func (o *Session1) GetIssuedAt() time.Time`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *Session1) GetIssuedAtOk() (*time.Time, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *Session1) SetIssuedAt(v time.Time)`

SetIssuedAt sets IssuedAt field to given value.

### HasIssuedAt

`func (o *Session1) HasIssuedAt() bool`

HasIssuedAt returns a boolean if a field has been set.

### GetIdentity

`func (o *Session1) GetIdentity() Identity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *Session1) GetIdentityOk() (*Identity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *Session1) SetIdentity(v Identity)`

SetIdentity sets Identity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


