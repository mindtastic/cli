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

// WikiEntryHeader struct for WikiEntryHeader
type WikiEntryHeader struct {
	Id *string `json:"id,omitempty"`
	Title string `json:"title"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewWikiEntryHeader instantiates a new WikiEntryHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWikiEntryHeader(title string) *WikiEntryHeader {
	this := WikiEntryHeader{}
	this.Title = title
	return &this
}

// NewWikiEntryHeaderWithDefaults instantiates a new WikiEntryHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWikiEntryHeaderWithDefaults() *WikiEntryHeader {
	this := WikiEntryHeader{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WikiEntryHeader) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiEntryHeader) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WikiEntryHeader) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WikiEntryHeader) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value
func (o *WikiEntryHeader) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *WikiEntryHeader) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *WikiEntryHeader) SetTitle(v string) {
	o.Title = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *WikiEntryHeader) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiEntryHeader) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WikiEntryHeader) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *WikiEntryHeader) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *WikiEntryHeader) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiEntryHeader) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *WikiEntryHeader) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *WikiEntryHeader) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o WikiEntryHeader) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableWikiEntryHeader struct {
	value *WikiEntryHeader
	isSet bool
}

func (v NullableWikiEntryHeader) Get() *WikiEntryHeader {
	return v.value
}

func (v *NullableWikiEntryHeader) Set(val *WikiEntryHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableWikiEntryHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableWikiEntryHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWikiEntryHeader(val *WikiEntryHeader) *NullableWikiEntryHeader {
	return &NullableWikiEntryHeader{value: val, isSet: true}
}

func (v NullableWikiEntryHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWikiEntryHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

