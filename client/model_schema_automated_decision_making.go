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

// SchemaAutomatedDecisionMaking struct for SchemaAutomatedDecisionMaking
type SchemaAutomatedDecisionMaking struct {
	InUse *bool `json:"inUse,omitempty"`
	LogicInvolved *string `json:"logicInvolved,omitempty"`
	ScopeAndIntendedEffects *string `json:"scopeAndIntendedEffects,omitempty"`
}

// NewSchemaAutomatedDecisionMaking instantiates a new SchemaAutomatedDecisionMaking object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaAutomatedDecisionMaking() *SchemaAutomatedDecisionMaking {
	this := SchemaAutomatedDecisionMaking{}
	return &this
}

// NewSchemaAutomatedDecisionMakingWithDefaults instantiates a new SchemaAutomatedDecisionMaking object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaAutomatedDecisionMakingWithDefaults() *SchemaAutomatedDecisionMaking {
	this := SchemaAutomatedDecisionMaking{}
	return &this
}

// GetInUse returns the InUse field value if set, zero value otherwise.
func (o *SchemaAutomatedDecisionMaking) GetInUse() bool {
	if o == nil || o.InUse == nil {
		var ret bool
		return ret
	}
	return *o.InUse
}

// GetInUseOk returns a tuple with the InUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaAutomatedDecisionMaking) GetInUseOk() (*bool, bool) {
	if o == nil || o.InUse == nil {
		return nil, false
	}
	return o.InUse, true
}

// HasInUse returns a boolean if a field has been set.
func (o *SchemaAutomatedDecisionMaking) HasInUse() bool {
	if o != nil && o.InUse != nil {
		return true
	}

	return false
}

// SetInUse gets a reference to the given bool and assigns it to the InUse field.
func (o *SchemaAutomatedDecisionMaking) SetInUse(v bool) {
	o.InUse = &v
}

// GetLogicInvolved returns the LogicInvolved field value if set, zero value otherwise.
func (o *SchemaAutomatedDecisionMaking) GetLogicInvolved() string {
	if o == nil || o.LogicInvolved == nil {
		var ret string
		return ret
	}
	return *o.LogicInvolved
}

// GetLogicInvolvedOk returns a tuple with the LogicInvolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaAutomatedDecisionMaking) GetLogicInvolvedOk() (*string, bool) {
	if o == nil || o.LogicInvolved == nil {
		return nil, false
	}
	return o.LogicInvolved, true
}

// HasLogicInvolved returns a boolean if a field has been set.
func (o *SchemaAutomatedDecisionMaking) HasLogicInvolved() bool {
	if o != nil && o.LogicInvolved != nil {
		return true
	}

	return false
}

// SetLogicInvolved gets a reference to the given string and assigns it to the LogicInvolved field.
func (o *SchemaAutomatedDecisionMaking) SetLogicInvolved(v string) {
	o.LogicInvolved = &v
}

// GetScopeAndIntendedEffects returns the ScopeAndIntendedEffects field value if set, zero value otherwise.
func (o *SchemaAutomatedDecisionMaking) GetScopeAndIntendedEffects() string {
	if o == nil || o.ScopeAndIntendedEffects == nil {
		var ret string
		return ret
	}
	return *o.ScopeAndIntendedEffects
}

// GetScopeAndIntendedEffectsOk returns a tuple with the ScopeAndIntendedEffects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaAutomatedDecisionMaking) GetScopeAndIntendedEffectsOk() (*string, bool) {
	if o == nil || o.ScopeAndIntendedEffects == nil {
		return nil, false
	}
	return o.ScopeAndIntendedEffects, true
}

// HasScopeAndIntendedEffects returns a boolean if a field has been set.
func (o *SchemaAutomatedDecisionMaking) HasScopeAndIntendedEffects() bool {
	if o != nil && o.ScopeAndIntendedEffects != nil {
		return true
	}

	return false
}

// SetScopeAndIntendedEffects gets a reference to the given string and assigns it to the ScopeAndIntendedEffects field.
func (o *SchemaAutomatedDecisionMaking) SetScopeAndIntendedEffects(v string) {
	o.ScopeAndIntendedEffects = &v
}

func (o SchemaAutomatedDecisionMaking) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InUse != nil {
		toSerialize["inUse"] = o.InUse
	}
	if o.LogicInvolved != nil {
		toSerialize["logicInvolved"] = o.LogicInvolved
	}
	if o.ScopeAndIntendedEffects != nil {
		toSerialize["scopeAndIntendedEffects"] = o.ScopeAndIntendedEffects
	}
	return json.Marshal(toSerialize)
}

type NullableSchemaAutomatedDecisionMaking struct {
	value *SchemaAutomatedDecisionMaking
	isSet bool
}

func (v NullableSchemaAutomatedDecisionMaking) Get() *SchemaAutomatedDecisionMaking {
	return v.value
}

func (v *NullableSchemaAutomatedDecisionMaking) Set(val *SchemaAutomatedDecisionMaking) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaAutomatedDecisionMaking) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaAutomatedDecisionMaking) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaAutomatedDecisionMaking(val *SchemaAutomatedDecisionMaking) *NullableSchemaAutomatedDecisionMaking {
	return &NullableSchemaAutomatedDecisionMaking{value: val, isSet: true}
}

func (v NullableSchemaAutomatedDecisionMaking) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaAutomatedDecisionMaking) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


