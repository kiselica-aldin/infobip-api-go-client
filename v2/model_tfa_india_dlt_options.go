/**
 * Infobip Client API Libraries OpenAPI Specification
 *
 * OpenAPI specification containing public endpoints supported in client API libraries.
 *
 * Contact: support@infobip.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * Do not edit the class manually.
 */

package infobip

import (
	"encoding/json"
)

// TfaIndiaDltOptions struct for TfaIndiaDltOptions
type TfaIndiaDltOptions struct {
	// Id of your registered DTL content template that matches this message's text.
	ContentTemplateId *string `json:"contentTemplateId,omitempty"`
	// Your assigned DTL principal entity id.
	PrincipalEntityId string `json:"principalEntityId"`
}

// NewTfaIndiaDltOptions instantiates a new TfaIndiaDltOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTfaIndiaDltOptions(principalEntityId string) *TfaIndiaDltOptions {
	this := TfaIndiaDltOptions{}
	this.PrincipalEntityId = principalEntityId
	return &this
}

// NewTfaIndiaDltOptionsWithDefaults instantiates a new TfaIndiaDltOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTfaIndiaDltOptionsWithDefaults() *TfaIndiaDltOptions {
	this := TfaIndiaDltOptions{}
	return &this
}

// GetContentTemplateId returns the ContentTemplateId field value if set, zero value otherwise.
func (o *TfaIndiaDltOptions) GetContentTemplateId() string {
	if o == nil || o.ContentTemplateId == nil {
		var ret string
		return ret
	}
	return *o.ContentTemplateId
}

// GetContentTemplateIdOk returns a tuple with the ContentTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaIndiaDltOptions) GetContentTemplateIdOk() (*string, bool) {
	if o == nil || o.ContentTemplateId == nil {
		return nil, false
	}
	return o.ContentTemplateId, true
}

// HasContentTemplateId returns a boolean if a field has been set.
func (o *TfaIndiaDltOptions) HasContentTemplateId() bool {
	if o != nil && o.ContentTemplateId != nil {
		return true
	}

	return false
}

// SetContentTemplateId gets a reference to the given string and assigns it to the ContentTemplateId field.
func (o *TfaIndiaDltOptions) SetContentTemplateId(v string) {
	o.ContentTemplateId = &v
}

// GetPrincipalEntityId returns the PrincipalEntityId field value
func (o *TfaIndiaDltOptions) GetPrincipalEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrincipalEntityId
}

// GetPrincipalEntityIdOk returns a tuple with the PrincipalEntityId field value
// and a boolean to check if the value has been set.
func (o *TfaIndiaDltOptions) GetPrincipalEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrincipalEntityId, true
}

// SetPrincipalEntityId sets field value
func (o *TfaIndiaDltOptions) SetPrincipalEntityId(v string) {
	o.PrincipalEntityId = v
}

func (o TfaIndiaDltOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContentTemplateId != nil {
		toSerialize["contentTemplateId"] = o.ContentTemplateId
	}
	if true {
		toSerialize["principalEntityId"] = o.PrincipalEntityId
	}
	return json.Marshal(toSerialize)
}

type NullableTfaIndiaDltOptions struct {
	value *TfaIndiaDltOptions
	isSet bool
}

func (v NullableTfaIndiaDltOptions) Get() *TfaIndiaDltOptions {
	return v.value
}

func (v *NullableTfaIndiaDltOptions) Set(val *TfaIndiaDltOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableTfaIndiaDltOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableTfaIndiaDltOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTfaIndiaDltOptions(val *TfaIndiaDltOptions) *NullableTfaIndiaDltOptions {
	return &NullableTfaIndiaDltOptions{value: val, isSet: true}
}

func (v NullableTfaIndiaDltOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTfaIndiaDltOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
