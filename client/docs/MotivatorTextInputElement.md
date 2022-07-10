# MotivatorTextInputElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputType** | **string** |  | 
**Label** | **string** | A string the input field shall be captioned with. | 
**MaxLength** | Pointer to **int32** | Optional value for the maximum allowed length for the text input | [optional] 

## Methods

### NewMotivatorTextInputElement

`func NewMotivatorTextInputElement(inputType string, label string, ) *MotivatorTextInputElement`

NewMotivatorTextInputElement instantiates a new MotivatorTextInputElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMotivatorTextInputElementWithDefaults

`func NewMotivatorTextInputElementWithDefaults() *MotivatorTextInputElement`

NewMotivatorTextInputElementWithDefaults instantiates a new MotivatorTextInputElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputType

`func (o *MotivatorTextInputElement) GetInputType() string`

GetInputType returns the InputType field if non-nil, zero value otherwise.

### GetInputTypeOk

`func (o *MotivatorTextInputElement) GetInputTypeOk() (*string, bool)`

GetInputTypeOk returns a tuple with the InputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputType

`func (o *MotivatorTextInputElement) SetInputType(v string)`

SetInputType sets InputType field to given value.


### GetLabel

`func (o *MotivatorTextInputElement) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MotivatorTextInputElement) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MotivatorTextInputElement) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetMaxLength

`func (o *MotivatorTextInputElement) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *MotivatorTextInputElement) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *MotivatorTextInputElement) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *MotivatorTextInputElement) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


