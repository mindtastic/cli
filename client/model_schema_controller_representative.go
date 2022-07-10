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

// SchemaControllerRepresentative struct for SchemaControllerRepresentative
type SchemaControllerRepresentative struct {
	Name *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

// NewSchemaControllerRepresentative instantiates a new SchemaControllerRepresentative object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaControllerRepresentative() *SchemaControllerRepresentative {
	this := SchemaControllerRepresentative{}
	return &this
}

// NewSchemaControllerRepresentativeWithDefaults instantiates a new SchemaControllerRepresentative object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaControllerRepresentativeWithDefaults() *SchemaControllerRepresentative {
	this := SchemaControllerRepresentative{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SchemaControllerRepresentative) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaControllerRepresentative) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SchemaControllerRepresentative) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SchemaControllerRepresentative) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SchemaControllerRepresentative) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaControllerRepresentative) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SchemaControllerRepresentative) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *SchemaControllerRepresentative) SetEmail(v string) {
	o.Email = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *SchemaControllerRepresentative) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaControllerRepresentative) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *SchemaControllerRepresentative) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *SchemaControllerRepresentative) SetPhone(v string) {
	o.Phone = &v
}

func (o SchemaControllerRepresentative) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	return json.Marshal(toSerialize)
}

type NullableSchemaControllerRepresentative struct {
	value *SchemaControllerRepresentative
	isSet bool
}

func (v NullableSchemaControllerRepresentative) Get() *SchemaControllerRepresentative {
	return v.value
}

func (v *NullableSchemaControllerRepresentative) Set(val *SchemaControllerRepresentative) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaControllerRepresentative) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaControllerRepresentative) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaControllerRepresentative(val *SchemaControllerRepresentative) *NullableSchemaControllerRepresentative {
	return &NullableSchemaControllerRepresentative{value: val, isSet: true}
}

func (v NullableSchemaControllerRepresentative) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaControllerRepresentative) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


