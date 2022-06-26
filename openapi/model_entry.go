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
)

// Entry struct for Entry
type Entry struct {
	Title *string `json:"title,omitempty"`
	Content []EntryContentInner `json:"content,omitempty"`
}

// NewEntry instantiates a new Entry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntry() *Entry {
	this := Entry{}
	return &this
}

// NewEntryWithDefaults instantiates a new Entry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntryWithDefaults() *Entry {
	this := Entry{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Entry) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entry) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Entry) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Entry) SetTitle(v string) {
	o.Title = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *Entry) GetContent() []EntryContentInner {
	if o == nil || o.Content == nil {
		var ret []EntryContentInner
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entry) GetContentOk() ([]EntryContentInner, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *Entry) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given []EntryContentInner and assigns it to the Content field.
func (o *Entry) SetContent(v []EntryContentInner) {
	o.Content = v
}

func (o Entry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullableEntry struct {
	value *Entry
	isSet bool
}

func (v NullableEntry) Get() *Entry {
	return v.value
}

func (v *NullableEntry) Set(val *Entry) {
	v.value = val
	v.isSet = true
}

func (v NullableEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntry(val *Entry) *NullableEntry {
	return &NullableEntry{value: val, isSet: true}
}

func (v NullableEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

