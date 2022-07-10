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

// BouncerRegisterRequest struct for BouncerRegisterRequest
type BouncerRegisterRequest struct {
	CsrfToken *string `json:"csrf_token,omitempty"`
}

// NewBouncerRegisterRequest instantiates a new BouncerRegisterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBouncerRegisterRequest() *BouncerRegisterRequest {
	this := BouncerRegisterRequest{}
	return &this
}

// NewBouncerRegisterRequestWithDefaults instantiates a new BouncerRegisterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBouncerRegisterRequestWithDefaults() *BouncerRegisterRequest {
	this := BouncerRegisterRequest{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *BouncerRegisterRequest) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BouncerRegisterRequest) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *BouncerRegisterRequest) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *BouncerRegisterRequest) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

func (o BouncerRegisterRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CsrfToken != nil {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	return json.Marshal(toSerialize)
}

type NullableBouncerRegisterRequest struct {
	value *BouncerRegisterRequest
	isSet bool
}

func (v NullableBouncerRegisterRequest) Get() *BouncerRegisterRequest {
	return v.value
}

func (v *NullableBouncerRegisterRequest) Set(val *BouncerRegisterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBouncerRegisterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBouncerRegisterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBouncerRegisterRequest(val *BouncerRegisterRequest) *NullableBouncerRegisterRequest {
	return &NullableBouncerRegisterRequest{value: val, isSet: true}
}

func (v NullableBouncerRegisterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBouncerRegisterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


