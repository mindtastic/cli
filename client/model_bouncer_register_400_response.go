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

// BouncerRegister400Response struct for BouncerRegister400Response
type BouncerRegister400Response struct {
	Error GenericError `json:"error"`
}

// NewBouncerRegister400Response instantiates a new BouncerRegister400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBouncerRegister400Response(error_ GenericError) *BouncerRegister400Response {
	this := BouncerRegister400Response{}
	this.Error = error_
	return &this
}

// NewBouncerRegister400ResponseWithDefaults instantiates a new BouncerRegister400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBouncerRegister400ResponseWithDefaults() *BouncerRegister400Response {
	this := BouncerRegister400Response{}
	return &this
}

// GetError returns the Error field value
func (o *BouncerRegister400Response) GetError() GenericError {
	if o == nil {
		var ret GenericError
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *BouncerRegister400Response) GetErrorOk() (*GenericError, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *BouncerRegister400Response) SetError(v GenericError) {
	o.Error = v
}

func (o BouncerRegister400Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableBouncerRegister400Response struct {
	value *BouncerRegister400Response
	isSet bool
}

func (v NullableBouncerRegister400Response) Get() *BouncerRegister400Response {
	return v.value
}

func (v *NullableBouncerRegister400Response) Set(val *BouncerRegister400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableBouncerRegister400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableBouncerRegister400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBouncerRegister400Response(val *BouncerRegister400Response) *NullableBouncerRegister400Response {
	return &NullableBouncerRegister400Response{value: val, isSet: true}
}

func (v NullableBouncerRegister400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBouncerRegister400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


