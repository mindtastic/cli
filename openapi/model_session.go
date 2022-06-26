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

// Session struct for Session
type Session struct {
	SessionToken *string `json:"session_token,omitempty"`
	Session map[string]interface{} `json:"session,omitempty"`
}

// NewSession instantiates a new Session object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSession() *Session {
	this := Session{}
	return &this
}

// NewSessionWithDefaults instantiates a new Session object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionWithDefaults() *Session {
	this := Session{}
	return &this
}

// GetSessionToken returns the SessionToken field value if set, zero value otherwise.
func (o *Session) GetSessionToken() string {
	if o == nil || o.SessionToken == nil {
		var ret string
		return ret
	}
	return *o.SessionToken
}

// GetSessionTokenOk returns a tuple with the SessionToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetSessionTokenOk() (*string, bool) {
	if o == nil || o.SessionToken == nil {
		return nil, false
	}
	return o.SessionToken, true
}

// HasSessionToken returns a boolean if a field has been set.
func (o *Session) HasSessionToken() bool {
	if o != nil && o.SessionToken != nil {
		return true
	}

	return false
}

// SetSessionToken gets a reference to the given string and assigns it to the SessionToken field.
func (o *Session) SetSessionToken(v string) {
	o.SessionToken = &v
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *Session) GetSession() map[string]interface{} {
	if o == nil || o.Session == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetSessionOk() (map[string]interface{}, bool) {
	if o == nil || o.Session == nil {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *Session) HasSession() bool {
	if o != nil && o.Session != nil {
		return true
	}

	return false
}

// SetSession gets a reference to the given map[string]interface{} and assigns it to the Session field.
func (o *Session) SetSession(v map[string]interface{}) {
	o.Session = v
}

func (o Session) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SessionToken != nil {
		toSerialize["session_token"] = o.SessionToken
	}
	if o.Session != nil {
		toSerialize["session"] = o.Session
	}
	return json.Marshal(toSerialize)
}

type NullableSession struct {
	value *Session
	isSet bool
}

func (v NullableSession) Get() *Session {
	return v.value
}

func (v *NullableSession) Set(val *Session) {
	v.value = val
	v.isSet = true
}

func (v NullableSession) IsSet() bool {
	return v.isSet
}

func (v *NullableSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSession(val *Session) *NullableSession {
	return &NullableSession{value: val, isSet: true}
}

func (v NullableSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

