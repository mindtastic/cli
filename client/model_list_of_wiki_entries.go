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

// ListOfWikiEntries struct for ListOfWikiEntries
type ListOfWikiEntries struct {
	EntryCount int32 `json:"entry_count"`
	Entries []WikiEntry `json:"entries"`
}

// NewListOfWikiEntries instantiates a new ListOfWikiEntries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOfWikiEntries(entryCount int32, entries []WikiEntry) *ListOfWikiEntries {
	this := ListOfWikiEntries{}
	this.EntryCount = entryCount
	this.Entries = entries
	return &this
}

// NewListOfWikiEntriesWithDefaults instantiates a new ListOfWikiEntries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOfWikiEntriesWithDefaults() *ListOfWikiEntries {
	this := ListOfWikiEntries{}
	return &this
}

// GetEntryCount returns the EntryCount field value
func (o *ListOfWikiEntries) GetEntryCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EntryCount
}

// GetEntryCountOk returns a tuple with the EntryCount field value
// and a boolean to check if the value has been set.
func (o *ListOfWikiEntries) GetEntryCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryCount, true
}

// SetEntryCount sets field value
func (o *ListOfWikiEntries) SetEntryCount(v int32) {
	o.EntryCount = v
}

// GetEntries returns the Entries field value
func (o *ListOfWikiEntries) GetEntries() []WikiEntry {
	if o == nil {
		var ret []WikiEntry
		return ret
	}

	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value
// and a boolean to check if the value has been set.
func (o *ListOfWikiEntries) GetEntriesOk() ([]WikiEntry, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entries, true
}

// SetEntries sets field value
func (o *ListOfWikiEntries) SetEntries(v []WikiEntry) {
	o.Entries = v
}

func (o ListOfWikiEntries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entry_count"] = o.EntryCount
	}
	if true {
		toSerialize["entries"] = o.Entries
	}
	return json.Marshal(toSerialize)
}

type NullableListOfWikiEntries struct {
	value *ListOfWikiEntries
	isSet bool
}

func (v NullableListOfWikiEntries) Get() *ListOfWikiEntries {
	return v.value
}

func (v *NullableListOfWikiEntries) Set(val *ListOfWikiEntries) {
	v.value = val
	v.isSet = true
}

func (v NullableListOfWikiEntries) IsSet() bool {
	return v.isSet
}

func (v *NullableListOfWikiEntries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOfWikiEntries(val *ListOfWikiEntries) *NullableListOfWikiEntries {
	return &NullableListOfWikiEntries{value: val, isSet: true}
}

func (v NullableListOfWikiEntries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOfWikiEntries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

