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
	"time"
)

// Motivator struct for Motivator
type Motivator struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Headline *string `json:"headline,omitempty"`
	Description *string `json:"description,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	Content []MotivatorContent `json:"content,omitempty"`
	Inputs []MotivatorInput `json:"inputs,omitempty"`
}

// NewMotivator instantiates a new Motivator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMotivator() *Motivator {
	this := Motivator{}
	return &this
}

// NewMotivatorWithDefaults instantiates a new Motivator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMotivatorWithDefaults() *Motivator {
	this := Motivator{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Motivator) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Motivator) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Motivator) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Motivator) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Motivator) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Motivator) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Motivator) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Motivator) SetName(v string) {
	o.Name = &v
}

// GetHeadline returns the Headline field value if set, zero value otherwise.
func (o *Motivator) GetHeadline() string {
	if o == nil || o.Headline == nil {
		var ret string
		return ret
	}
	return *o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Motivator) GetHeadlineOk() (*string, bool) {
	if o == nil || o.Headline == nil {
		return nil, false
	}
	return o.Headline, true
}

// HasHeadline returns a boolean if a field has been set.
func (o *Motivator) HasHeadline() bool {
	if o != nil && o.Headline != nil {
		return true
	}

	return false
}

// SetHeadline gets a reference to the given string and assigns it to the Headline field.
func (o *Motivator) SetHeadline(v string) {
	o.Headline = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Motivator) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Motivator) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Motivator) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Motivator) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Motivator) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Motivator) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Motivator) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Motivator) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *Motivator) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Motivator) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *Motivator) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *Motivator) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *Motivator) GetContent() []MotivatorContent {
	if o == nil || o.Content == nil {
		var ret []MotivatorContent
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Motivator) GetContentOk() ([]MotivatorContent, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *Motivator) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given []MotivatorContent and assigns it to the Content field.
func (o *Motivator) SetContent(v []MotivatorContent) {
	o.Content = v
}

// GetInputs returns the Inputs field value if set, zero value otherwise.
func (o *Motivator) GetInputs() []MotivatorInput {
	if o == nil || o.Inputs == nil {
		var ret []MotivatorInput
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Motivator) GetInputsOk() ([]MotivatorInput, bool) {
	if o == nil || o.Inputs == nil {
		return nil, false
	}
	return o.Inputs, true
}

// HasInputs returns a boolean if a field has been set.
func (o *Motivator) HasInputs() bool {
	if o != nil && o.Inputs != nil {
		return true
	}

	return false
}

// SetInputs gets a reference to the given []MotivatorInput and assigns it to the Inputs field.
func (o *Motivator) SetInputs(v []MotivatorInput) {
	o.Inputs = v
}

func (o Motivator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Headline != nil {
		toSerialize["headline"] = o.Headline
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.Inputs != nil {
		toSerialize["inputs"] = o.Inputs
	}
	return json.Marshal(toSerialize)
}

type NullableMotivator struct {
	value *Motivator
	isSet bool
}

func (v NullableMotivator) Get() *Motivator {
	return v.value
}

func (v *NullableMotivator) Set(val *Motivator) {
	v.value = val
	v.isSet = true
}

func (v NullableMotivator) IsSet() bool {
	return v.isSet
}

func (v *NullableMotivator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMotivator(val *Motivator) *NullableMotivator {
	return &NullableMotivator{value: val, isSet: true}
}

func (v NullableMotivator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMotivator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


