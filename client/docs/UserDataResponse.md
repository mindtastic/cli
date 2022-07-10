# UserDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** |  | [optional] 
**Role** | Pointer to [**RoleEnum**](RoleEnum.md) |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to [**UserSettingsModel**](UserSettingsModel.md) |  | [optional] 

## Methods

### NewUserDataResponse

`func NewUserDataResponse() *UserDataResponse`

NewUserDataResponse instantiates a new UserDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDataResponseWithDefaults

`func NewUserDataResponseWithDefaults() *UserDataResponse`

NewUserDataResponseWithDefaults instantiates a new UserDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserDataResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserDataResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserDataResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserDataResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetRole

`func (o *UserDataResponse) GetRole() RoleEnum`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserDataResponse) GetRoleOk() (*RoleEnum, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserDataResponse) SetRole(v RoleEnum)`

SetRole sets Role field to given value.

### HasRole

`func (o *UserDataResponse) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetEmail

`func (o *UserDataResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserDataResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserDataResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserDataResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSettings

`func (o *UserDataResponse) GetSettings() UserSettingsModel`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *UserDataResponse) GetSettingsOk() (*UserSettingsModel, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *UserDataResponse) SetSettings(v UserSettingsModel)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *UserDataResponse) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


