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

// SchemaDataDisclosedInnerStorageInner struct for SchemaDataDisclosedInnerStorageInner
type SchemaDataDisclosedInnerStorageInner struct {
	Temporal []SchemaDataDisclosedInnerStorageInnerTemporalInner `json:"temporal,omitempty"`
	PurposeConditional []map[string]interface{} `json:"purposeConditional,omitempty"`
	LegalBasisConditional []map[string]interface{} `json:"legalBasisConditional,omitempty"`
	AggregationFunction *string `json:"aggregationFunction,omitempty"`
}

// NewSchemaDataDisclosedInnerStorageInner instantiates a new SchemaDataDisclosedInnerStorageInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaDataDisclosedInnerStorageInner() *SchemaDataDisclosedInnerStorageInner {
	this := SchemaDataDisclosedInnerStorageInner{}
	return &this
}

// NewSchemaDataDisclosedInnerStorageInnerWithDefaults instantiates a new SchemaDataDisclosedInnerStorageInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaDataDisclosedInnerStorageInnerWithDefaults() *SchemaDataDisclosedInnerStorageInner {
	this := SchemaDataDisclosedInnerStorageInner{}
	return &this
}

// GetTemporal returns the Temporal field value if set, zero value otherwise.
func (o *SchemaDataDisclosedInnerStorageInner) GetTemporal() []SchemaDataDisclosedInnerStorageInnerTemporalInner {
	if o == nil || o.Temporal == nil {
		var ret []SchemaDataDisclosedInnerStorageInnerTemporalInner
		return ret
	}
	return o.Temporal
}

// GetTemporalOk returns a tuple with the Temporal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaDataDisclosedInnerStorageInner) GetTemporalOk() ([]SchemaDataDisclosedInnerStorageInnerTemporalInner, bool) {
	if o == nil || o.Temporal == nil {
		return nil, false
	}
	return o.Temporal, true
}

// HasTemporal returns a boolean if a field has been set.
func (o *SchemaDataDisclosedInnerStorageInner) HasTemporal() bool {
	if o != nil && o.Temporal != nil {
		return true
	}

	return false
}

// SetTemporal gets a reference to the given []SchemaDataDisclosedInnerStorageInnerTemporalInner and assigns it to the Temporal field.
func (o *SchemaDataDisclosedInnerStorageInner) SetTemporal(v []SchemaDataDisclosedInnerStorageInnerTemporalInner) {
	o.Temporal = v
}

// GetPurposeConditional returns the PurposeConditional field value if set, zero value otherwise.
func (o *SchemaDataDisclosedInnerStorageInner) GetPurposeConditional() []map[string]interface{} {
	if o == nil || o.PurposeConditional == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.PurposeConditional
}

// GetPurposeConditionalOk returns a tuple with the PurposeConditional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaDataDisclosedInnerStorageInner) GetPurposeConditionalOk() ([]map[string]interface{}, bool) {
	if o == nil || o.PurposeConditional == nil {
		return nil, false
	}
	return o.PurposeConditional, true
}

// HasPurposeConditional returns a boolean if a field has been set.
func (o *SchemaDataDisclosedInnerStorageInner) HasPurposeConditional() bool {
	if o != nil && o.PurposeConditional != nil {
		return true
	}

	return false
}

// SetPurposeConditional gets a reference to the given []map[string]interface{} and assigns it to the PurposeConditional field.
func (o *SchemaDataDisclosedInnerStorageInner) SetPurposeConditional(v []map[string]interface{}) {
	o.PurposeConditional = v
}

// GetLegalBasisConditional returns the LegalBasisConditional field value if set, zero value otherwise.
func (o *SchemaDataDisclosedInnerStorageInner) GetLegalBasisConditional() []map[string]interface{} {
	if o == nil || o.LegalBasisConditional == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.LegalBasisConditional
}

// GetLegalBasisConditionalOk returns a tuple with the LegalBasisConditional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaDataDisclosedInnerStorageInner) GetLegalBasisConditionalOk() ([]map[string]interface{}, bool) {
	if o == nil || o.LegalBasisConditional == nil {
		return nil, false
	}
	return o.LegalBasisConditional, true
}

// HasLegalBasisConditional returns a boolean if a field has been set.
func (o *SchemaDataDisclosedInnerStorageInner) HasLegalBasisConditional() bool {
	if o != nil && o.LegalBasisConditional != nil {
		return true
	}

	return false
}

// SetLegalBasisConditional gets a reference to the given []map[string]interface{} and assigns it to the LegalBasisConditional field.
func (o *SchemaDataDisclosedInnerStorageInner) SetLegalBasisConditional(v []map[string]interface{}) {
	o.LegalBasisConditional = v
}

// GetAggregationFunction returns the AggregationFunction field value if set, zero value otherwise.
func (o *SchemaDataDisclosedInnerStorageInner) GetAggregationFunction() string {
	if o == nil || o.AggregationFunction == nil {
		var ret string
		return ret
	}
	return *o.AggregationFunction
}

// GetAggregationFunctionOk returns a tuple with the AggregationFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaDataDisclosedInnerStorageInner) GetAggregationFunctionOk() (*string, bool) {
	if o == nil || o.AggregationFunction == nil {
		return nil, false
	}
	return o.AggregationFunction, true
}

// HasAggregationFunction returns a boolean if a field has been set.
func (o *SchemaDataDisclosedInnerStorageInner) HasAggregationFunction() bool {
	if o != nil && o.AggregationFunction != nil {
		return true
	}

	return false
}

// SetAggregationFunction gets a reference to the given string and assigns it to the AggregationFunction field.
func (o *SchemaDataDisclosedInnerStorageInner) SetAggregationFunction(v string) {
	o.AggregationFunction = &v
}

func (o SchemaDataDisclosedInnerStorageInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Temporal != nil {
		toSerialize["temporal"] = o.Temporal
	}
	if o.PurposeConditional != nil {
		toSerialize["purposeConditional"] = o.PurposeConditional
	}
	if o.LegalBasisConditional != nil {
		toSerialize["legalBasisConditional"] = o.LegalBasisConditional
	}
	if o.AggregationFunction != nil {
		toSerialize["aggregationFunction"] = o.AggregationFunction
	}
	return json.Marshal(toSerialize)
}

type NullableSchemaDataDisclosedInnerStorageInner struct {
	value *SchemaDataDisclosedInnerStorageInner
	isSet bool
}

func (v NullableSchemaDataDisclosedInnerStorageInner) Get() *SchemaDataDisclosedInnerStorageInner {
	return v.value
}

func (v *NullableSchemaDataDisclosedInnerStorageInner) Set(val *SchemaDataDisclosedInnerStorageInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaDataDisclosedInnerStorageInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaDataDisclosedInnerStorageInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaDataDisclosedInnerStorageInner(val *SchemaDataDisclosedInnerStorageInner) *NullableSchemaDataDisclosedInnerStorageInner {
	return &NullableSchemaDataDisclosedInnerStorageInner{value: val, isSet: true}
}

func (v NullableSchemaDataDisclosedInnerStorageInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaDataDisclosedInnerStorageInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


