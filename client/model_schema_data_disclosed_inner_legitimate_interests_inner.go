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

// SchemaDataDisclosedInnerLegitimateInterestsInner struct for SchemaDataDisclosedInnerLegitimateInterestsInner
type SchemaDataDisclosedInnerLegitimateInterestsInner struct {
	Exists *bool `json:"exists,omitempty"`
	Reasoning *string `json:"reasoning,omitempty"`
}

// NewSchemaDataDisclosedInnerLegitimateInterestsInner instantiates a new SchemaDataDisclosedInnerLegitimateInterestsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaDataDisclosedInnerLegitimateInterestsInner() *SchemaDataDisclosedInnerLegitimateInterestsInner {
	this := SchemaDataDisclosedInnerLegitimateInterestsInner{}
	return &this
}

// NewSchemaDataDisclosedInnerLegitimateInterestsInnerWithDefaults instantiates a new SchemaDataDisclosedInnerLegitimateInterestsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaDataDisclosedInnerLegitimateInterestsInnerWithDefaults() *SchemaDataDisclosedInnerLegitimateInterestsInner {
	this := SchemaDataDisclosedInnerLegitimateInterestsInner{}
	return &this
}

// GetExists returns the Exists field value if set, zero value otherwise.
func (o *SchemaDataDisclosedInnerLegitimateInterestsInner) GetExists() bool {
	if o == nil || o.Exists == nil {
		var ret bool
		return ret
	}
	return *o.Exists
}

// GetExistsOk returns a tuple with the Exists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaDataDisclosedInnerLegitimateInterestsInner) GetExistsOk() (*bool, bool) {
	if o == nil || o.Exists == nil {
		return nil, false
	}
	return o.Exists, true
}

// HasExists returns a boolean if a field has been set.
func (o *SchemaDataDisclosedInnerLegitimateInterestsInner) HasExists() bool {
	if o != nil && o.Exists != nil {
		return true
	}

	return false
}

// SetExists gets a reference to the given bool and assigns it to the Exists field.
func (o *SchemaDataDisclosedInnerLegitimateInterestsInner) SetExists(v bool) {
	o.Exists = &v
}

// GetReasoning returns the Reasoning field value if set, zero value otherwise.
func (o *SchemaDataDisclosedInnerLegitimateInterestsInner) GetReasoning() string {
	if o == nil || o.Reasoning == nil {
		var ret string
		return ret
	}
	return *o.Reasoning
}

// GetReasoningOk returns a tuple with the Reasoning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaDataDisclosedInnerLegitimateInterestsInner) GetReasoningOk() (*string, bool) {
	if o == nil || o.Reasoning == nil {
		return nil, false
	}
	return o.Reasoning, true
}

// HasReasoning returns a boolean if a field has been set.
func (o *SchemaDataDisclosedInnerLegitimateInterestsInner) HasReasoning() bool {
	if o != nil && o.Reasoning != nil {
		return true
	}

	return false
}

// SetReasoning gets a reference to the given string and assigns it to the Reasoning field.
func (o *SchemaDataDisclosedInnerLegitimateInterestsInner) SetReasoning(v string) {
	o.Reasoning = &v
}

func (o SchemaDataDisclosedInnerLegitimateInterestsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Exists != nil {
		toSerialize["exists"] = o.Exists
	}
	if o.Reasoning != nil {
		toSerialize["reasoning"] = o.Reasoning
	}
	return json.Marshal(toSerialize)
}

type NullableSchemaDataDisclosedInnerLegitimateInterestsInner struct {
	value *SchemaDataDisclosedInnerLegitimateInterestsInner
	isSet bool
}

func (v NullableSchemaDataDisclosedInnerLegitimateInterestsInner) Get() *SchemaDataDisclosedInnerLegitimateInterestsInner {
	return v.value
}

func (v *NullableSchemaDataDisclosedInnerLegitimateInterestsInner) Set(val *SchemaDataDisclosedInnerLegitimateInterestsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaDataDisclosedInnerLegitimateInterestsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaDataDisclosedInnerLegitimateInterestsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaDataDisclosedInnerLegitimateInterestsInner(val *SchemaDataDisclosedInnerLegitimateInterestsInner) *NullableSchemaDataDisclosedInnerLegitimateInterestsInner {
	return &NullableSchemaDataDisclosedInnerLegitimateInterestsInner{value: val, isSet: true}
}

func (v NullableSchemaDataDisclosedInnerLegitimateInterestsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaDataDisclosedInnerLegitimateInterestsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

