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
	"fmt"
)

// AuthenticatorAssuranceLevel Internal value used by Ory Kratos. Not relevant for the Kopfsachen API
type AuthenticatorAssuranceLevel string

// List of authenticatorAssuranceLevel
const (
	AAL0 AuthenticatorAssuranceLevel = "aal0"
	AAL1 AuthenticatorAssuranceLevel = "aal1"
	AAL2 AuthenticatorAssuranceLevel = "aal2"
	AAL3 AuthenticatorAssuranceLevel = "aal3"
)

// All allowed values of AuthenticatorAssuranceLevel enum
var AllowedAuthenticatorAssuranceLevelEnumValues = []AuthenticatorAssuranceLevel{
	"aal0",
	"aal1",
	"aal2",
	"aal3",
}

func (v *AuthenticatorAssuranceLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AuthenticatorAssuranceLevel(value)
	for _, existing := range AllowedAuthenticatorAssuranceLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AuthenticatorAssuranceLevel", value)
}

// NewAuthenticatorAssuranceLevelFromValue returns a pointer to a valid AuthenticatorAssuranceLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAuthenticatorAssuranceLevelFromValue(v string) (*AuthenticatorAssuranceLevel, error) {
	ev := AuthenticatorAssuranceLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AuthenticatorAssuranceLevel: valid values are %v", v, AllowedAuthenticatorAssuranceLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AuthenticatorAssuranceLevel) IsValid() bool {
	for _, existing := range AllowedAuthenticatorAssuranceLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to authenticatorAssuranceLevel value
func (v AuthenticatorAssuranceLevel) Ptr() *AuthenticatorAssuranceLevel {
	return &v
}

type NullableAuthenticatorAssuranceLevel struct {
	value *AuthenticatorAssuranceLevel
	isSet bool
}

func (v NullableAuthenticatorAssuranceLevel) Get() *AuthenticatorAssuranceLevel {
	return v.value
}

func (v *NullableAuthenticatorAssuranceLevel) Set(val *AuthenticatorAssuranceLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorAssuranceLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorAssuranceLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorAssuranceLevel(val *AuthenticatorAssuranceLevel) *NullableAuthenticatorAssuranceLevel {
	return &NullableAuthenticatorAssuranceLevel{value: val, isSet: true}
}

func (v NullableAuthenticatorAssuranceLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorAssuranceLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

