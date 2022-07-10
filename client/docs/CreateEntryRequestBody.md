# CreateEntryRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateEntryRequestBody

`func NewCreateEntryRequestBody() *CreateEntryRequestBody`

NewCreateEntryRequestBody instantiates a new CreateEntryRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEntryRequestBodyWithDefaults

`func NewCreateEntryRequestBodyWithDefaults() *CreateEntryRequestBody`

NewCreateEntryRequestBodyWithDefaults instantiates a new CreateEntryRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateEntryRequestBody) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateEntryRequestBody) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateEntryRequestBody) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateEntryRequestBody) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetContent

`func (o *CreateEntryRequestBody) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreateEntryRequestBody) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreateEntryRequestBody) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CreateEntryRequestBody) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


