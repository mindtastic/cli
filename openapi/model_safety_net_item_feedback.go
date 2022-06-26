/*
Kopfsachen

Kopfsachen e. V. is an association for the promotion of young people's mental health. The goal is to teach the basics of mental health literacy in various educational formats.

API version: 0.2
Contact: mail@kopfsachen.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// SafetyNetItemFeedback 
type SafetyNetItemFeedback struct {
	ItHelped *bool `json:"itHelped,omitempty"`
	Comment *string `json:"comment,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// NewSafetyNetItemFeedback instantiates a new SafetyNetItemFeedback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSafetyNetItemFeedback() *SafetyNetItemFeedback {
	this := SafetyNetItemFeedback{}
	return &this
}

// NewSafetyNetItemFeedbackWithDefaults instantiates a new SafetyNetItemFeedback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSafetyNetItemFeedbackWithDefaults() *SafetyNetItemFeedback {
	this := SafetyNetItemFeedback{}
	return &this
}

// GetItHelped returns the ItHelped field value if set, zero value otherwise.
func (o *SafetyNetItemFeedback) GetItHelped() bool {
	if o == nil || o.ItHelped == nil {
		var ret bool
		return ret
	}
	return *o.ItHelped
}

// GetItHelpedOk returns a tuple with the ItHelped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafetyNetItemFeedback) GetItHelpedOk() (*bool, bool) {
	if o == nil || o.ItHelped == nil {
		return nil, false
	}
	return o.ItHelped, true
}

// HasItHelped returns a boolean if a field has been set.
func (o *SafetyNetItemFeedback) HasItHelped() bool {
	if o != nil && o.ItHelped != nil {
		return true
	}

	return false
}

// SetItHelped gets a reference to the given bool and assigns it to the ItHelped field.
func (o *SafetyNetItemFeedback) SetItHelped(v bool) {
	o.ItHelped = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *SafetyNetItemFeedback) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafetyNetItemFeedback) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *SafetyNetItemFeedback) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *SafetyNetItemFeedback) SetComment(v string) {
	o.Comment = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *SafetyNetItemFeedback) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SafetyNetItemFeedback) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *SafetyNetItemFeedback) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *SafetyNetItemFeedback) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o SafetyNetItemFeedback) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ItHelped != nil {
		toSerialize["itHelped"] = o.ItHelped
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableSafetyNetItemFeedback struct {
	value *SafetyNetItemFeedback
	isSet bool
}

func (v NullableSafetyNetItemFeedback) Get() *SafetyNetItemFeedback {
	return v.value
}

func (v *NullableSafetyNetItemFeedback) Set(val *SafetyNetItemFeedback) {
	v.value = val
	v.isSet = true
}

func (v NullableSafetyNetItemFeedback) IsSet() bool {
	return v.isSet
}

func (v *NullableSafetyNetItemFeedback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSafetyNetItemFeedback(val *SafetyNetItemFeedback) *NullableSafetyNetItemFeedback {
	return &NullableSafetyNetItemFeedback{value: val, isSet: true}
}

func (v NullableSafetyNetItemFeedback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSafetyNetItemFeedback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

