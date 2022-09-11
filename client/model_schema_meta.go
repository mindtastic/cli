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

// SchemaMeta struct for SchemaMeta
type SchemaMeta struct {
	Id *string `json:"_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Created *string `json:"created,omitempty"`
	Modified *string `json:"modified,omitempty"`
	Version *float32 `json:"version,omitempty"`
	Language *string `json:"language,omitempty"`
	Status *string `json:"status,omitempty"`
	Url *string `json:"url,omitempty"`
	Hash *string `json:"_hash,omitempty"`
}

// NewSchemaMeta instantiates a new SchemaMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaMeta() *SchemaMeta {
	this := SchemaMeta{}
	return &this
}

// NewSchemaMetaWithDefaults instantiates a new SchemaMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaMetaWithDefaults() *SchemaMeta {
	this := SchemaMeta{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SchemaMeta) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaMeta) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SchemaMeta) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SchemaMeta) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SchemaMeta) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaMeta) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SchemaMeta) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SchemaMeta) SetName(v string) {
	o.Name = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *SchemaMeta) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaMeta) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *SchemaMeta) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *SchemaMeta) SetCreated(v string) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *SchemaMeta) GetModified() string {
	if o == nil || o.Modified == nil {
		var ret string
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaMeta) GetModifiedOk() (*string, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *SchemaMeta) HasModified() bool {
	if o != nil && o.Modified != nil {
		return true
	}

	return false
}

// SetModified gets a reference to the given string and assigns it to the Modified field.
func (o *SchemaMeta) SetModified(v string) {
	o.Modified = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SchemaMeta) GetVersion() float32 {
	if o == nil || o.Version == nil {
		var ret float32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaMeta) GetVersionOk() (*float32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SchemaMeta) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given float32 and assigns it to the Version field.
func (o *SchemaMeta) SetVersion(v float32) {
	o.Version = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *SchemaMeta) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaMeta) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *SchemaMeta) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *SchemaMeta) SetLanguage(v string) {
	o.Language = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SchemaMeta) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaMeta) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SchemaMeta) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SchemaMeta) SetStatus(v string) {
	o.Status = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SchemaMeta) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaMeta) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SchemaMeta) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SchemaMeta) SetUrl(v string) {
	o.Url = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *SchemaMeta) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaMeta) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *SchemaMeta) HasHash() bool {
	if o != nil && o.Hash != nil {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *SchemaMeta) SetHash(v string) {
	o.Hash = &v
}

func (o SchemaMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["_id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Modified != nil {
		toSerialize["modified"] = o.Modified
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Hash != nil {
		toSerialize["_hash"] = o.Hash
	}
	return json.Marshal(toSerialize)
}

type NullableSchemaMeta struct {
	value *SchemaMeta
	isSet bool
}

func (v NullableSchemaMeta) Get() *SchemaMeta {
	return v.value
}

func (v *NullableSchemaMeta) Set(val *SchemaMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaMeta(val *SchemaMeta) *NullableSchemaMeta {
	return &NullableSchemaMeta{value: val, isSet: true}
}

func (v NullableSchemaMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

