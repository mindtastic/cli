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

// SafetyNetItem struct for SafetyNetItem
type SafetyNetItem struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	// Which kind of item is this (i.a., influences symbol in UI)?
	Type string `json:"type"`
	// How can this item help the user to feel better? An array of strings describing the strategies.
	Strategies []string `json:"strategies"`
	Feedback []SafetyNetItemFeedback `json:"feedback,omitempty"`
}

// NewSafetyNetItem instantiates a new SafetyNetItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSafetyNetItem(id int32, name string, type_ string, strategies []string) *SafetyNetItem {
	this := SafetyNetItem{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.Strategies = strategies
	return &this
}

// NewSafetyNetItemWithDefaults instantiates a new SafetyNetItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSafetyNetItemWithDefaults() *SafetyNetItem {
	this := SafetyNetItem{}
	return &this
}

// GetId returns the Id field value
func (o *SafetyNetItem) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SafetyNetItem) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SafetyNetItem) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *SafetyNetItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SafetyNetItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SafetyNetItem) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *SafetyNetItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SafetyNetItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SafetyNetItem) SetType(v string) {
	o.Type = v
}

// GetStrategies returns the Strategies field value
func (o *SafetyNetItem) GetStrategies() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Strategies
}

// GetStrategiesOk returns a tuple with the Strategies field value
// and a boolean to check if the value has been set.
func (o *SafetyNetItem) GetStrategiesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Strategies, true
}

// SetStrategies sets field value
func (o *SafetyNetItem) SetStrategies(v []string) {
	o.Strategies = v
}

// GetFeedback returns the Feedback field value if set, zero value otherwise.
func (o *SafetyNetItem) GetFeedback() []SafetyNetItemFeedback {
	if o == nil || o.Feedback == nil {
		var ret []SafetyNetItemFeedback
		return ret
	}
	return o.Feedback
}

// GetFeedbackOk returns a tuple with the Feedback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafetyNetItem) GetFeedbackOk() ([]SafetyNetItemFeedback, bool) {
	if o == nil || o.Feedback == nil {
		return nil, false
	}
	return o.Feedback, true
}

// HasFeedback returns a boolean if a field has been set.
func (o *SafetyNetItem) HasFeedback() bool {
	if o != nil && o.Feedback != nil {
		return true
	}

	return false
}

// SetFeedback gets a reference to the given []SafetyNetItemFeedback and assigns it to the Feedback field.
func (o *SafetyNetItem) SetFeedback(v []SafetyNetItemFeedback) {
	o.Feedback = v
}

func (o SafetyNetItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["strategies"] = o.Strategies
	}
	if o.Feedback != nil {
		toSerialize["feedback"] = o.Feedback
	}
	return json.Marshal(toSerialize)
}

type NullableSafetyNetItem struct {
	value *SafetyNetItem
	isSet bool
}

func (v NullableSafetyNetItem) Get() *SafetyNetItem {
	return v.value
}

func (v *NullableSafetyNetItem) Set(val *SafetyNetItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSafetyNetItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSafetyNetItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSafetyNetItem(val *SafetyNetItem) *NullableSafetyNetItem {
	return &NullableSafetyNetItem{value: val, isSet: true}
}

func (v NullableSafetyNetItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSafetyNetItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


