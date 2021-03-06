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

// SchemaSourcesInnerSourcesInner struct for SchemaSourcesInnerSourcesInner
type SchemaSourcesInnerSourcesInner struct {
	Description *string `json:"description,omitempty"`
	Url *string `json:"url,omitempty"`
	PubliclyAvailable *bool `json:"publiclyAvailable,omitempty"`
}

// NewSchemaSourcesInnerSourcesInner instantiates a new SchemaSourcesInnerSourcesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaSourcesInnerSourcesInner() *SchemaSourcesInnerSourcesInner {
	this := SchemaSourcesInnerSourcesInner{}
	return &this
}

// NewSchemaSourcesInnerSourcesInnerWithDefaults instantiates a new SchemaSourcesInnerSourcesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaSourcesInnerSourcesInnerWithDefaults() *SchemaSourcesInnerSourcesInner {
	this := SchemaSourcesInnerSourcesInner{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SchemaSourcesInnerSourcesInner) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSourcesInnerSourcesInner) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SchemaSourcesInnerSourcesInner) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SchemaSourcesInnerSourcesInner) SetDescription(v string) {
	o.Description = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SchemaSourcesInnerSourcesInner) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSourcesInnerSourcesInner) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SchemaSourcesInnerSourcesInner) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SchemaSourcesInnerSourcesInner) SetUrl(v string) {
	o.Url = &v
}

// GetPubliclyAvailable returns the PubliclyAvailable field value if set, zero value otherwise.
func (o *SchemaSourcesInnerSourcesInner) GetPubliclyAvailable() bool {
	if o == nil || o.PubliclyAvailable == nil {
		var ret bool
		return ret
	}
	return *o.PubliclyAvailable
}

// GetPubliclyAvailableOk returns a tuple with the PubliclyAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSourcesInnerSourcesInner) GetPubliclyAvailableOk() (*bool, bool) {
	if o == nil || o.PubliclyAvailable == nil {
		return nil, false
	}
	return o.PubliclyAvailable, true
}

// HasPubliclyAvailable returns a boolean if a field has been set.
func (o *SchemaSourcesInnerSourcesInner) HasPubliclyAvailable() bool {
	if o != nil && o.PubliclyAvailable != nil {
		return true
	}

	return false
}

// SetPubliclyAvailable gets a reference to the given bool and assigns it to the PubliclyAvailable field.
func (o *SchemaSourcesInnerSourcesInner) SetPubliclyAvailable(v bool) {
	o.PubliclyAvailable = &v
}

func (o SchemaSourcesInnerSourcesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.PubliclyAvailable != nil {
		toSerialize["publiclyAvailable"] = o.PubliclyAvailable
	}
	return json.Marshal(toSerialize)
}

type NullableSchemaSourcesInnerSourcesInner struct {
	value *SchemaSourcesInnerSourcesInner
	isSet bool
}

func (v NullableSchemaSourcesInnerSourcesInner) Get() *SchemaSourcesInnerSourcesInner {
	return v.value
}

func (v *NullableSchemaSourcesInnerSourcesInner) Set(val *SchemaSourcesInnerSourcesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaSourcesInnerSourcesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaSourcesInnerSourcesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaSourcesInnerSourcesInner(val *SchemaSourcesInnerSourcesInner) *NullableSchemaSourcesInnerSourcesInner {
	return &NullableSchemaSourcesInnerSourcesInner{value: val, isSet: true}
}

func (v NullableSchemaSourcesInnerSourcesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaSourcesInnerSourcesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


