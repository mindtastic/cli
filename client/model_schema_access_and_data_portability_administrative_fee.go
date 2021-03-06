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

// SchemaAccessAndDataPortabilityAdministrativeFee struct for SchemaAccessAndDataPortabilityAdministrativeFee
type SchemaAccessAndDataPortabilityAdministrativeFee struct {
	Amount *float32 `json:"amount,omitempty"`
	Currency *string `json:"currency,omitempty"`
}

// NewSchemaAccessAndDataPortabilityAdministrativeFee instantiates a new SchemaAccessAndDataPortabilityAdministrativeFee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaAccessAndDataPortabilityAdministrativeFee() *SchemaAccessAndDataPortabilityAdministrativeFee {
	this := SchemaAccessAndDataPortabilityAdministrativeFee{}
	return &this
}

// NewSchemaAccessAndDataPortabilityAdministrativeFeeWithDefaults instantiates a new SchemaAccessAndDataPortabilityAdministrativeFee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaAccessAndDataPortabilityAdministrativeFeeWithDefaults() *SchemaAccessAndDataPortabilityAdministrativeFee {
	this := SchemaAccessAndDataPortabilityAdministrativeFee{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *SchemaAccessAndDataPortabilityAdministrativeFee) GetAmount() float32 {
	if o == nil || o.Amount == nil {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaAccessAndDataPortabilityAdministrativeFee) GetAmountOk() (*float32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *SchemaAccessAndDataPortabilityAdministrativeFee) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *SchemaAccessAndDataPortabilityAdministrativeFee) SetAmount(v float32) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *SchemaAccessAndDataPortabilityAdministrativeFee) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaAccessAndDataPortabilityAdministrativeFee) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *SchemaAccessAndDataPortabilityAdministrativeFee) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *SchemaAccessAndDataPortabilityAdministrativeFee) SetCurrency(v string) {
	o.Currency = &v
}

func (o SchemaAccessAndDataPortabilityAdministrativeFee) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	return json.Marshal(toSerialize)
}

type NullableSchemaAccessAndDataPortabilityAdministrativeFee struct {
	value *SchemaAccessAndDataPortabilityAdministrativeFee
	isSet bool
}

func (v NullableSchemaAccessAndDataPortabilityAdministrativeFee) Get() *SchemaAccessAndDataPortabilityAdministrativeFee {
	return v.value
}

func (v *NullableSchemaAccessAndDataPortabilityAdministrativeFee) Set(val *SchemaAccessAndDataPortabilityAdministrativeFee) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaAccessAndDataPortabilityAdministrativeFee) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaAccessAndDataPortabilityAdministrativeFee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaAccessAndDataPortabilityAdministrativeFee(val *SchemaAccessAndDataPortabilityAdministrativeFee) *NullableSchemaAccessAndDataPortabilityAdministrativeFee {
	return &NullableSchemaAccessAndDataPortabilityAdministrativeFee{value: val, isSet: true}
}

func (v NullableSchemaAccessAndDataPortabilityAdministrativeFee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaAccessAndDataPortabilityAdministrativeFee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


