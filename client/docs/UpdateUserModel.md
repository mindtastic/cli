# UpdateUserModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** |  | [optional] 
**Role** | Pointer to [**RoleEnum**](RoleEnum.md) |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to [**UserSettingsModel**](UserSettingsModel.md) |  | [optional] 

## Methods

### NewUpdateUserModel

`func NewUpdateUserModel() *UpdateUserModel`

NewUpdateUserModel instantiates a new UpdateUserModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserModelWithDefaults

`func NewUpdateUserModelWithDefaults() *UpdateUserModel`

NewUpdateUserModelWithDefaults instantiates a new UpdateUserModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UpdateUserModel) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateUserModel) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateUserModel) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateUserModel) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetRole

`func (o *UpdateUserModel) GetRole() RoleEnum`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UpdateUserModel) GetRoleOk() (*RoleEnum, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UpdateUserModel) SetRole(v RoleEnum)`

SetRole sets Role field to given value.

### HasRole

`func (o *UpdateUserModel) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateUserModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateUserModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateUserModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateUserModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSettings

`func (o *UpdateUserModel) GetSettings() UserSettingsModel`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *UpdateUserModel) GetSettingsOk() (*UserSettingsModel, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *UpdateUserModel) SetSettings(v UserSettingsModel)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *UpdateUserModel) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


