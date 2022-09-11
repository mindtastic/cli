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

// MotivatorResultItemInputValues struct for MotivatorResultItemInputValues
type MotivatorResultItemInputValues struct {
	// The actual value of the input to store. It MUST BE a valid json string.
	Value string `json:"value"`
}

// NewMotivatorResultItemInputValues instantiates a new MotivatorResultItemInputValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMotivatorResultItemInputValues(value string) *MotivatorResultItemInputValues {
	this := MotivatorResultItemInputValues{}
	this.Value = value
	return &this
}

// NewMotivatorResultItemInputValuesWithDefaults instantiates a new MotivatorResultItemInputValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMotivatorResultItemInputValuesWithDefaults() *MotivatorResultItemInputValues {
	this := MotivatorResultItemInputValues{}
	return &this
}

// GetValue returns the Value field value
func (o *MotivatorResultItemInputValues) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *MotivatorResultItemInputValues) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *MotivatorResultItemInputValues) SetValue(v string) {
	o.Value = v
}

func (o MotivatorResultItemInputValues) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableMotivatorResultItemInputValues struct {
	value *MotivatorResultItemInputValues
	isSet bool
}

func (v NullableMotivatorResultItemInputValues) Get() *MotivatorResultItemInputValues {
	return v.value
}

func (v *NullableMotivatorResultItemInputValues) Set(val *MotivatorResultItemInputValues) {
	v.value = val
	v.isSet = true
}

func (v NullableMotivatorResultItemInputValues) IsSet() bool {
	return v.isSet
}

func (v *NullableMotivatorResultItemInputValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMotivatorResultItemInputValues(val *MotivatorResultItemInputValues) *NullableMotivatorResultItemInputValues {
	return &NullableMotivatorResultItemInputValues{value: val, isSet: true}
}

func (v NullableMotivatorResultItemInputValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMotivatorResultItemInputValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

