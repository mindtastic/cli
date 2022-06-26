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

// GenericError1 struct for GenericError1
type GenericError1 struct {
	// The HTTP status code of the reponse
	StatusCode *int32 `json:"status_code,omitempty"`
	// Error message
	Error *string `json:"error,omitempty"`
}

// NewGenericError1 instantiates a new GenericError1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericError1() *GenericError1 {
	this := GenericError1{}
	return &this
}

// NewGenericError1WithDefaults instantiates a new GenericError1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericError1WithDefaults() *GenericError1 {
	this := GenericError1{}
	return &this
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *GenericError1) GetStatusCode() int32 {
	if o == nil || o.StatusCode == nil {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericError1) GetStatusCodeOk() (*int32, bool) {
	if o == nil || o.StatusCode == nil {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *GenericError1) HasStatusCode() bool {
	if o != nil && o.StatusCode != nil {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *GenericError1) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GenericError1) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericError1) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GenericError1) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *GenericError1) SetError(v string) {
	o.Error = &v
}

func (o GenericError1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StatusCode != nil {
		toSerialize["status_code"] = o.StatusCode
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableGenericError1 struct {
	value *GenericError1
	isSet bool
}

func (v NullableGenericError1) Get() *GenericError1 {
	return v.value
}

func (v *NullableGenericError1) Set(val *GenericError1) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericError1) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericError1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericError1(val *GenericError1) *NullableGenericError1 {
	return &NullableGenericError1{value: val, isSet: true}
}

func (v NullableGenericError1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericError1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


