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

// SchemaDataDisclosedInnerRecipientsInnerRepresentative struct for SchemaDataDisclosedInnerRecipientsInnerRepresentative
type SchemaDataDisclosedInnerRecipientsInnerRepresentative struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// NewSchemaDataDisclosedInnerRecipientsInnerRepresentative instantiates a new SchemaDataDisclosedInnerRecipientsInnerRepresentative object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaDataDisclosedInnerRecipientsInnerRepresentative(name string, email string, phone string) *SchemaDataDisclosedInnerRecipientsInnerRepresentative {
	this := SchemaDataDisclosedInnerRecipientsInnerRepresentative{}
	this.Name = name
	this.Email = email
	this.Phone = phone
	return &this
}

// NewSchemaDataDisclosedInnerRecipientsInnerRepresentativeWithDefaults instantiates a new SchemaDataDisclosedInnerRecipientsInnerRepresentative object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaDataDisclosedInnerRecipientsInnerRepresentativeWithDefaults() *SchemaDataDisclosedInnerRecipientsInnerRepresentative {
	this := SchemaDataDisclosedInnerRecipientsInnerRepresentative{}
	return &this
}

// GetName returns the Name field value
func (o *SchemaDataDisclosedInnerRecipientsInnerRepresentative) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SchemaDataDisclosedInnerRecipientsInnerRepresentative) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SchemaDataDisclosedInnerRecipientsInnerRepresentative) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value
func (o *SchemaDataDisclosedInnerRecipientsInnerRepresentative) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SchemaDataDisclosedInnerRecipientsInnerRepresentative) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SchemaDataDisclosedInnerRecipientsInnerRepresentative) SetEmail(v string) {
	o.Email = v
}

// GetPhone returns the Phone field value
func (o *SchemaDataDisclosedInnerRecipientsInnerRepresentative) GetPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *SchemaDataDisclosedInnerRecipientsInnerRepresentative) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *SchemaDataDisclosedInnerRecipientsInnerRepresentative) SetPhone(v string) {
	o.Phone = v
}

func (o SchemaDataDisclosedInnerRecipientsInnerRepresentative) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["phone"] = o.Phone
	}
	return json.Marshal(toSerialize)
}

type NullableSchemaDataDisclosedInnerRecipientsInnerRepresentative struct {
	value *SchemaDataDisclosedInnerRecipientsInnerRepresentative
	isSet bool
}

func (v NullableSchemaDataDisclosedInnerRecipientsInnerRepresentative) Get() *SchemaDataDisclosedInnerRecipientsInnerRepresentative {
	return v.value
}

func (v *NullableSchemaDataDisclosedInnerRecipientsInnerRepresentative) Set(val *SchemaDataDisclosedInnerRecipientsInnerRepresentative) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaDataDisclosedInnerRecipientsInnerRepresentative) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaDataDisclosedInnerRecipientsInnerRepresentative) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaDataDisclosedInnerRecipientsInnerRepresentative(val *SchemaDataDisclosedInnerRecipientsInnerRepresentative) *NullableSchemaDataDisclosedInnerRecipientsInnerRepresentative {
	return &NullableSchemaDataDisclosedInnerRecipientsInnerRepresentative{value: val, isSet: true}
}

func (v NullableSchemaDataDisclosedInnerRecipientsInnerRepresentative) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaDataDisclosedInnerRecipientsInnerRepresentative) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


