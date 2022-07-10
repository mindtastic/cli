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

// BrowserLoginFlowExample struct for BrowserLoginFlowExample
type BrowserLoginFlowExample struct {
}

// NewBrowserLoginFlowExample instantiates a new BrowserLoginFlowExample object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrowserLoginFlowExample() *BrowserLoginFlowExample {
	this := BrowserLoginFlowExample{}
	return &this
}

// NewBrowserLoginFlowExampleWithDefaults instantiates a new BrowserLoginFlowExample object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrowserLoginFlowExampleWithDefaults() *BrowserLoginFlowExample {
	this := BrowserLoginFlowExample{}
	return &this
}

func (o BrowserLoginFlowExample) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableBrowserLoginFlowExample struct {
	value *BrowserLoginFlowExample
	isSet bool
}

func (v NullableBrowserLoginFlowExample) Get() *BrowserLoginFlowExample {
	return v.value
}

func (v *NullableBrowserLoginFlowExample) Set(val *BrowserLoginFlowExample) {
	v.value = val
	v.isSet = true
}

func (v NullableBrowserLoginFlowExample) IsSet() bool {
	return v.isSet
}

func (v *NullableBrowserLoginFlowExample) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrowserLoginFlowExample(val *BrowserLoginFlowExample) *NullableBrowserLoginFlowExample {
	return &NullableBrowserLoginFlowExample{value: val, isSet: true}
}

func (v NullableBrowserLoginFlowExample) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrowserLoginFlowExample) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


