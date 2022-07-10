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

// WikiHeaders struct for WikiHeaders
type WikiHeaders struct {
	Entries []WikiEntryHeader `json:"entries,omitempty"`
}

// NewWikiHeaders instantiates a new WikiHeaders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWikiHeaders() *WikiHeaders {
	this := WikiHeaders{}
	return &this
}

// NewWikiHeadersWithDefaults instantiates a new WikiHeaders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWikiHeadersWithDefaults() *WikiHeaders {
	this := WikiHeaders{}
	return &this
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *WikiHeaders) GetEntries() []WikiEntryHeader {
	if o == nil || o.Entries == nil {
		var ret []WikiEntryHeader
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiHeaders) GetEntriesOk() ([]WikiEntryHeader, bool) {
	if o == nil || o.Entries == nil {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *WikiHeaders) HasEntries() bool {
	if o != nil && o.Entries != nil {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []WikiEntryHeader and assigns it to the Entries field.
func (o *WikiHeaders) SetEntries(v []WikiEntryHeader) {
	o.Entries = v
}

func (o WikiHeaders) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Entries != nil {
		toSerialize["entries"] = o.Entries
	}
	return json.Marshal(toSerialize)
}

type NullableWikiHeaders struct {
	value *WikiHeaders
	isSet bool
}

func (v NullableWikiHeaders) Get() *WikiHeaders {
	return v.value
}

func (v *NullableWikiHeaders) Set(val *WikiHeaders) {
	v.value = val
	v.isSet = true
}

func (v NullableWikiHeaders) IsSet() bool {
	return v.isSet
}

func (v *NullableWikiHeaders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWikiHeaders(val *WikiHeaders) *NullableWikiHeaders {
	return &NullableWikiHeaders{value: val, isSet: true}
}

func (v NullableWikiHeaders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWikiHeaders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


