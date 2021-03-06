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

// InputContent struct for InputContent
type InputContent struct {
	Type string `json:"type"`
	InputId int32 `json:"inputId"`
}

// NewInputContent instantiates a new InputContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputContent(type_ string, inputId int32) *InputContent {
	this := InputContent{}
	this.Type = type_
	this.InputId = inputId
	return &this
}

// NewInputContentWithDefaults instantiates a new InputContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputContentWithDefaults() *InputContent {
	this := InputContent{}
	return &this
}

// GetType returns the Type field value
func (o *InputContent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InputContent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InputContent) SetType(v string) {
	o.Type = v
}

// GetInputId returns the InputId field value
func (o *InputContent) GetInputId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InputId
}

// GetInputIdOk returns a tuple with the InputId field value
// and a boolean to check if the value has been set.
func (o *InputContent) GetInputIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputId, true
}

// SetInputId sets field value
func (o *InputContent) SetInputId(v int32) {
	o.InputId = v
}

func (o InputContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["inputId"] = o.InputId
	}
	return json.Marshal(toSerialize)
}

type NullableInputContent struct {
	value *InputContent
	isSet bool
}

func (v NullableInputContent) Get() *InputContent {
	return v.value
}

func (v *NullableInputContent) Set(val *InputContent) {
	v.value = val
	v.isSet = true
}

func (v NullableInputContent) IsSet() bool {
	return v.isSet
}

func (v *NullableInputContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputContent(val *InputContent) *NullableInputContent {
	return &NullableInputContent{value: val, isSet: true}
}

func (v NullableInputContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


