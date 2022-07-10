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

// PageBreakContent struct for PageBreakContent
type PageBreakContent struct {
	Type string `json:"type"`
}

// NewPageBreakContent instantiates a new PageBreakContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageBreakContent(type_ string) *PageBreakContent {
	this := PageBreakContent{}
	this.Type = type_
	return &this
}

// NewPageBreakContentWithDefaults instantiates a new PageBreakContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageBreakContentWithDefaults() *PageBreakContent {
	this := PageBreakContent{}
	return &this
}

// GetType returns the Type field value
func (o *PageBreakContent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PageBreakContent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PageBreakContent) SetType(v string) {
	o.Type = v
}

func (o PageBreakContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullablePageBreakContent struct {
	value *PageBreakContent
	isSet bool
}

func (v NullablePageBreakContent) Get() *PageBreakContent {
	return v.value
}

func (v *NullablePageBreakContent) Set(val *PageBreakContent) {
	v.value = val
	v.isSet = true
}

func (v NullablePageBreakContent) IsSet() bool {
	return v.isSet
}

func (v *NullablePageBreakContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageBreakContent(val *PageBreakContent) *NullablePageBreakContent {
	return &NullablePageBreakContent{value: val, isSet: true}
}

func (v NullablePageBreakContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageBreakContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


