/*
Kopfsachen

Kopfsachen e. V. is an association for the promotion of young people's mental health. The goal is to teach the basics of mental health literacy in various educational formats.

API version: 1.0.2
Contact: mail@kopfsachen.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// MotivatorSliderInputElementOptions An object that configures the range input
type MotivatorSliderInputElementOptions struct {
	MinValue *int32 `json:"minValue,omitempty"`
	MaxValue *int32 `json:"maxValue,omitempty"`
	StepSize *int32 `json:"stepSize,omitempty"`
}

// NewMotivatorSliderInputElementOptions instantiates a new MotivatorSliderInputElementOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMotivatorSliderInputElementOptions() *MotivatorSliderInputElementOptions {
	this := MotivatorSliderInputElementOptions{}
	var minValue int32 = 0
	this.MinValue = &minValue
	var maxValue int32 = 100
	this.MaxValue = &maxValue
	var stepSize int32 = 1
	this.StepSize = &stepSize
	return &this
}

// NewMotivatorSliderInputElementOptionsWithDefaults instantiates a new MotivatorSliderInputElementOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMotivatorSliderInputElementOptionsWithDefaults() *MotivatorSliderInputElementOptions {
	this := MotivatorSliderInputElementOptions{}
	var minValue int32 = 0
	this.MinValue = &minValue
	var maxValue int32 = 100
	this.MaxValue = &maxValue
	var stepSize int32 = 1
	this.StepSize = &stepSize
	return &this
}

// GetMinValue returns the MinValue field value if set, zero value otherwise.
func (o *MotivatorSliderInputElementOptions) GetMinValue() int32 {
	if o == nil || o.MinValue == nil {
		var ret int32
		return ret
	}
	return *o.MinValue
}

// GetMinValueOk returns a tuple with the MinValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MotivatorSliderInputElementOptions) GetMinValueOk() (*int32, bool) {
	if o == nil || o.MinValue == nil {
		return nil, false
	}
	return o.MinValue, true
}

// HasMinValue returns a boolean if a field has been set.
func (o *MotivatorSliderInputElementOptions) HasMinValue() bool {
	if o != nil && o.MinValue != nil {
		return true
	}

	return false
}

// SetMinValue gets a reference to the given int32 and assigns it to the MinValue field.
func (o *MotivatorSliderInputElementOptions) SetMinValue(v int32) {
	o.MinValue = &v
}

// GetMaxValue returns the MaxValue field value if set, zero value otherwise.
func (o *MotivatorSliderInputElementOptions) GetMaxValue() int32 {
	if o == nil || o.MaxValue == nil {
		var ret int32
		return ret
	}
	return *o.MaxValue
}

// GetMaxValueOk returns a tuple with the MaxValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MotivatorSliderInputElementOptions) GetMaxValueOk() (*int32, bool) {
	if o == nil || o.MaxValue == nil {
		return nil, false
	}
	return o.MaxValue, true
}

// HasMaxValue returns a boolean if a field has been set.
func (o *MotivatorSliderInputElementOptions) HasMaxValue() bool {
	if o != nil && o.MaxValue != nil {
		return true
	}

	return false
}

// SetMaxValue gets a reference to the given int32 and assigns it to the MaxValue field.
func (o *MotivatorSliderInputElementOptions) SetMaxValue(v int32) {
	o.MaxValue = &v
}

// GetStepSize returns the StepSize field value if set, zero value otherwise.
func (o *MotivatorSliderInputElementOptions) GetStepSize() int32 {
	if o == nil || o.StepSize == nil {
		var ret int32
		return ret
	}
	return *o.StepSize
}

// GetStepSizeOk returns a tuple with the StepSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MotivatorSliderInputElementOptions) GetStepSizeOk() (*int32, bool) {
	if o == nil || o.StepSize == nil {
		return nil, false
	}
	return o.StepSize, true
}

// HasStepSize returns a boolean if a field has been set.
func (o *MotivatorSliderInputElementOptions) HasStepSize() bool {
	if o != nil && o.StepSize != nil {
		return true
	}

	return false
}

// SetStepSize gets a reference to the given int32 and assigns it to the StepSize field.
func (o *MotivatorSliderInputElementOptions) SetStepSize(v int32) {
	o.StepSize = &v
}

func (o MotivatorSliderInputElementOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MinValue != nil {
		toSerialize["minValue"] = o.MinValue
	}
	if o.MaxValue != nil {
		toSerialize["maxValue"] = o.MaxValue
	}
	if o.StepSize != nil {
		toSerialize["stepSize"] = o.StepSize
	}
	return json.Marshal(toSerialize)
}

type NullableMotivatorSliderInputElementOptions struct {
	value *MotivatorSliderInputElementOptions
	isSet bool
}

func (v NullableMotivatorSliderInputElementOptions) Get() *MotivatorSliderInputElementOptions {
	return v.value
}

func (v *NullableMotivatorSliderInputElementOptions) Set(val *MotivatorSliderInputElementOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableMotivatorSliderInputElementOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableMotivatorSliderInputElementOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMotivatorSliderInputElementOptions(val *MotivatorSliderInputElementOptions) *NullableMotivatorSliderInputElementOptions {
	return &NullableMotivatorSliderInputElementOptions{value: val, isSet: true}
}

func (v NullableMotivatorSliderInputElementOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMotivatorSliderInputElementOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


