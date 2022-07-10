# MotivatorInputElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputType** | **string** |  | 
**Label** | **string** | A string the range input shall be captioned with | 
**MaxLength** | Pointer to **int32** | Optional value for the maximum allowed length for the text input | [optional] 
**Options** | Pointer to [**MotivatorSliderInputElementOptions**](MotivatorSliderInputElementOptions.md) |  | [optional] 

## Methods

### NewMotivatorInputElement

`func NewMotivatorInputElement(inputType string, label string, ) *MotivatorInputElement`

NewMotivatorInputElement instantiates a new MotivatorInputElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMotivatorInputElementWithDefaults

`func NewMotivatorInputElementWithDefaults() *MotivatorInputElement`

NewMotivatorInputElementWithDefaults instantiates a new MotivatorInputElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputType

`func (o *MotivatorInputElement) GetInputType() string`

GetInputType returns the InputType field if non-nil, zero value otherwise.

### GetInputTypeOk

`func (o *MotivatorInputElement) GetInputTypeOk() (*string, bool)`

GetInputTypeOk returns a tuple with the InputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputType

`func (o *MotivatorInputElement) SetInputType(v string)`

SetInputType sets InputType field to given value.


### GetLabel

`func (o *MotivatorInputElement) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MotivatorInputElement) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MotivatorInputElement) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetMaxLength

`func (o *MotivatorInputElement) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *MotivatorInputElement) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *MotivatorInputElement) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *MotivatorInputElement) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetOptions

`func (o *MotivatorInputElement) GetOptions() MotivatorSliderInputElementOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *MotivatorInputElement) GetOptionsOk() (*MotivatorSliderInputElementOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *MotivatorInputElement) SetOptions(v MotivatorSliderInputElementOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *MotivatorInputElement) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


