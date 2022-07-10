# Motivator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Headline** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ModifiedAt** | Pointer to **time.Time** |  | [optional] 
**Content** | Pointer to [**[]MotivatorContent**](MotivatorContent.md) |  | [optional] 
**Inputs** | Pointer to [**[]MotivatorInput**](MotivatorInput.md) |  | [optional] 

## Methods

### NewMotivator

`func NewMotivator() *Motivator`

NewMotivator instantiates a new Motivator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMotivatorWithDefaults

`func NewMotivatorWithDefaults() *Motivator`

NewMotivatorWithDefaults instantiates a new Motivator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Motivator) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Motivator) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Motivator) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Motivator) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Motivator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Motivator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Motivator) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Motivator) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHeadline

`func (o *Motivator) GetHeadline() string`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *Motivator) GetHeadlineOk() (*string, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *Motivator) SetHeadline(v string)`

SetHeadline sets Headline field to given value.

### HasHeadline

`func (o *Motivator) HasHeadline() bool`

HasHeadline returns a boolean if a field has been set.

### GetDescription

`func (o *Motivator) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Motivator) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Motivator) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Motivator) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Motivator) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Motivator) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Motivator) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Motivator) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Motivator) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Motivator) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Motivator) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Motivator) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetContent

`func (o *Motivator) GetContent() []MotivatorContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Motivator) GetContentOk() (*[]MotivatorContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Motivator) SetContent(v []MotivatorContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *Motivator) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetInputs

`func (o *Motivator) GetInputs() []MotivatorInput`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *Motivator) GetInputsOk() (*[]MotivatorInput, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *Motivator) SetInputs(v []MotivatorInput)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *Motivator) HasInputs() bool`

HasInputs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


