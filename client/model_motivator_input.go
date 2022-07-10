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

// MotivatorInput struct for MotivatorInput
type MotivatorInput struct {
	Id *int32 `json:"id,omitempty"`
	Fields *map[string]MotivatorInputElement `json:"fields,omitempty"`
}

// NewMotivatorInput instantiates a new MotivatorInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMotivatorInput() *MotivatorInput {
	this := MotivatorInput{}
	return &this
}

// NewMotivatorInputWithDefaults instantiates a new MotivatorInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMotivatorInputWithDefaults() *MotivatorInput {
	this := MotivatorInput{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MotivatorInput) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MotivatorInput) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MotivatorInput) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MotivatorInput) SetId(v int32) {
	o.Id = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *MotivatorInput) GetFields() map[string]MotivatorInputElement {
	if o == nil || o.Fields == nil {
		var ret map[string]MotivatorInputElement
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MotivatorInput) GetFieldsOk() (*map[string]MotivatorInputElement, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *MotivatorInput) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string]MotivatorInputElement and assigns it to the Fields field.
func (o *MotivatorInput) SetFields(v map[string]MotivatorInputElement) {
	o.Fields = &v
}

func (o MotivatorInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	return json.Marshal(toSerialize)
}

type NullableMotivatorInput struct {
	value *MotivatorInput
	isSet bool
}

func (v NullableMotivatorInput) Get() *MotivatorInput {
	return v.value
}

func (v *NullableMotivatorInput) Set(val *MotivatorInput) {
	v.value = val
	v.isSet = true
}

func (v NullableMotivatorInput) IsSet() bool {
	return v.isSet
}

func (v *NullableMotivatorInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMotivatorInput(val *MotivatorInput) *NullableMotivatorInput {
	return &NullableMotivatorInput{value: val, isSet: true}
}

func (v NullableMotivatorInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMotivatorInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


