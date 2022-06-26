# Motivator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Headline** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TextContents** | Pointer to **[]string** |  | [optional] 
**MediaContents** | Pointer to [**[]MotivatorMediaContentsInner**](MotivatorMediaContentsInner.md) |  | [optional] 
**UserInputForm** | Pointer to **string** | May describe a form for input by the user (format to be defined, e.g., JSON Schema). | [optional] 

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

`func (o *Motivator) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Motivator) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Motivator) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Motivator) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Motivator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Motivator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Motivator) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Motivator) HasType() bool`

HasType returns a boolean if a field has been set.

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

### GetTextContents

`func (o *Motivator) GetTextContents() []string`

GetTextContents returns the TextContents field if non-nil, zero value otherwise.

### GetTextContentsOk

`func (o *Motivator) GetTextContentsOk() (*[]string, bool)`

GetTextContentsOk returns a tuple with the TextContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextContents

`func (o *Motivator) SetTextContents(v []string)`

SetTextContents sets TextContents field to given value.

### HasTextContents

`func (o *Motivator) HasTextContents() bool`

HasTextContents returns a boolean if a field has been set.

### GetMediaContents

`func (o *Motivator) GetMediaContents() []MotivatorMediaContentsInner`

GetMediaContents returns the MediaContents field if non-nil, zero value otherwise.

### GetMediaContentsOk

`func (o *Motivator) GetMediaContentsOk() (*[]MotivatorMediaContentsInner, bool)`

GetMediaContentsOk returns a tuple with the MediaContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaContents

`func (o *Motivator) SetMediaContents(v []MotivatorMediaContentsInner)`

SetMediaContents sets MediaContents field to given value.

### HasMediaContents

`func (o *Motivator) HasMediaContents() bool`

HasMediaContents returns a boolean if a field has been set.

### GetUserInputForm

`func (o *Motivator) GetUserInputForm() string`

GetUserInputForm returns the UserInputForm field if non-nil, zero value otherwise.

### GetUserInputFormOk

`func (o *Motivator) GetUserInputFormOk() (*string, bool)`

GetUserInputFormOk returns a tuple with the UserInputForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInputForm

`func (o *Motivator) SetUserInputForm(v string)`

SetUserInputForm sets UserInputForm field to given value.

### HasUserInputForm

`func (o *Motivator) HasUserInputForm() bool`

HasUserInputForm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


