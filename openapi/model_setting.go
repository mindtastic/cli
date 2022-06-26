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

// Setting Key-value pairs of settings in the app.
type Setting struct {
	Lang *string `json:"lang,omitempty"`
}

// NewSetting instantiates a new Setting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetting() *Setting {
	this := Setting{}
	var lang string = "de"
	this.Lang = &lang
	return &this
}

// NewSettingWithDefaults instantiates a new Setting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingWithDefaults() *Setting {
	this := Setting{}
	var lang string = "de"
	this.Lang = &lang
	return &this
}

// GetLang returns the Lang field value if set, zero value otherwise.
func (o *Setting) GetLang() string {
	if o == nil || o.Lang == nil {
		var ret string
		return ret
	}
	return *o.Lang
}

// GetLangOk returns a tuple with the Lang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Setting) GetLangOk() (*string, bool) {
	if o == nil || o.Lang == nil {
		return nil, false
	}
	return o.Lang, true
}

// HasLang returns a boolean if a field has been set.
func (o *Setting) HasLang() bool {
	if o != nil && o.Lang != nil {
		return true
	}

	return false
}

// SetLang gets a reference to the given string and assigns it to the Lang field.
func (o *Setting) SetLang(v string) {
	o.Lang = &v
}

func (o Setting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Lang != nil {
		toSerialize["lang"] = o.Lang
	}
	return json.Marshal(toSerialize)
}

type NullableSetting struct {
	value *Setting
	isSet bool
}

func (v NullableSetting) Get() *Setting {
	return v.value
}

func (v *NullableSetting) Set(val *Setting) {
	v.value = val
	v.isSet = true
}

func (v NullableSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetting(val *Setting) *NullableSetting {
	return &NullableSetting{value: val, isSet: true}
}

func (v NullableSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


