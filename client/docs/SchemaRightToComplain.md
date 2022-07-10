# SchemaRightToComplain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**IdentificationEvidences** | Pointer to **[]map[string]interface{}** |  | [optional] 
**SupervisoryAuthority** | Pointer to [**SchemaDataProtectionOfficer**](SchemaDataProtectionOfficer.md) |  | [optional] 

## Methods

### NewSchemaRightToComplain

`func NewSchemaRightToComplain() *SchemaRightToComplain`

NewSchemaRightToComplain instantiates a new SchemaRightToComplain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaRightToComplainWithDefaults

`func NewSchemaRightToComplainWithDefaults() *SchemaRightToComplain`

NewSchemaRightToComplainWithDefaults instantiates a new SchemaRightToComplain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *SchemaRightToComplain) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *SchemaRightToComplain) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *SchemaRightToComplain) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *SchemaRightToComplain) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetDescription

`func (o *SchemaRightToComplain) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SchemaRightToComplain) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SchemaRightToComplain) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SchemaRightToComplain) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *SchemaRightToComplain) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SchemaRightToComplain) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SchemaRightToComplain) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SchemaRightToComplain) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEmail

`func (o *SchemaRightToComplain) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SchemaRightToComplain) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SchemaRightToComplain) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SchemaRightToComplain) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetIdentificationEvidences

`func (o *SchemaRightToComplain) GetIdentificationEvidences() []map[string]interface{}`

GetIdentificationEvidences returns the IdentificationEvidences field if non-nil, zero value otherwise.

### GetIdentificationEvidencesOk

`func (o *SchemaRightToComplain) GetIdentificationEvidencesOk() (*[]map[string]interface{}, bool)`

GetIdentificationEvidencesOk returns a tuple with the IdentificationEvidences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationEvidences

`func (o *SchemaRightToComplain) SetIdentificationEvidences(v []map[string]interface{})`

SetIdentificationEvidences sets IdentificationEvidences field to given value.

### HasIdentificationEvidences

`func (o *SchemaRightToComplain) HasIdentificationEvidences() bool`

HasIdentificationEvidences returns a boolean if a field has been set.

### GetSupervisoryAuthority

`func (o *SchemaRightToComplain) GetSupervisoryAuthority() SchemaDataProtectionOfficer`

GetSupervisoryAuthority returns the SupervisoryAuthority field if non-nil, zero value otherwise.

### GetSupervisoryAuthorityOk

`func (o *SchemaRightToComplain) GetSupervisoryAuthorityOk() (*SchemaDataProtectionOfficer, bool)`

GetSupervisoryAuthorityOk returns a tuple with the SupervisoryAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisoryAuthority

`func (o *SchemaRightToComplain) SetSupervisoryAuthority(v SchemaDataProtectionOfficer)`

SetSupervisoryAuthority sets SupervisoryAuthority field to given value.

### HasSupervisoryAuthority

`func (o *SchemaRightToComplain) HasSupervisoryAuthority() bool`

HasSupervisoryAuthority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


