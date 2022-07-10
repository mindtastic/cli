# SchemaSourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DataCategory** | Pointer to **string** |  | [optional] 
**Sources** | Pointer to [**[]SchemaSourcesInnerSourcesInner**](SchemaSourcesInnerSourcesInner.md) |  | [optional] 

## Methods

### NewSchemaSourcesInner

`func NewSchemaSourcesInner() *SchemaSourcesInner`

NewSchemaSourcesInner instantiates a new SchemaSourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaSourcesInnerWithDefaults

`func NewSchemaSourcesInnerWithDefaults() *SchemaSourcesInner`

NewSchemaSourcesInnerWithDefaults instantiates a new SchemaSourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SchemaSourcesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SchemaSourcesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SchemaSourcesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SchemaSourcesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDataCategory

`func (o *SchemaSourcesInner) GetDataCategory() string`

GetDataCategory returns the DataCategory field if non-nil, zero value otherwise.

### GetDataCategoryOk

`func (o *SchemaSourcesInner) GetDataCategoryOk() (*string, bool)`

GetDataCategoryOk returns a tuple with the DataCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCategory

`func (o *SchemaSourcesInner) SetDataCategory(v string)`

SetDataCategory sets DataCategory field to given value.

### HasDataCategory

`func (o *SchemaSourcesInner) HasDataCategory() bool`

HasDataCategory returns a boolean if a field has been set.

### GetSources

`func (o *SchemaSourcesInner) GetSources() []SchemaSourcesInnerSourcesInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *SchemaSourcesInner) GetSourcesOk() (*[]SchemaSourcesInnerSourcesInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *SchemaSourcesInner) SetSources(v []SchemaSourcesInnerSourcesInner)`

SetSources sets Sources field to given value.

### HasSources

`func (o *SchemaSourcesInner) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


