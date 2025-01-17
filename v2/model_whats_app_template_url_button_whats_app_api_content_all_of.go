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

// WhatsAppTemplateUrlButtonWhatsAppApiContentAllOf struct for WhatsAppTemplateUrlButtonWhatsAppApiContentAllOf
type WhatsAppTemplateUrlButtonWhatsAppApiContentAllOf struct {
	// URL extension of a `dynamic URL` defined in the registered template.
	Parameter *string `json:"parameter,omitempty"`
}

// NewWhatsAppTemplateUrlButtonWhatsAppApiContentAllOf instantiates a new WhatsAppTemplateUrlButtonWhatsAppApiContentAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsAppTemplateUrlButtonWhatsAppApiContentAllOf() *WhatsAppTemplateUrlButtonWhatsAppApiContentAllOf {
	this := WhatsAppTemplateUrlButtonWhatsAppApiContentAllOf{}
	return &this
}

// NewWhatsAppTemplateUrlButtonWhatsAppApiContentAllOfWithDefaults instantiates a new WhatsAppTemplateUrlButtonWhatsAppApiContentAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsAppTemplateUrlButtonWhatsAppApiContentAllOfWithDefaults() *WhatsAppTemplateUrlButtonWhatsAppApiContentAllOf {
	this := WhatsAppTemplateUrlButtonWhatsAppApiContentAllOf{}
	return &this
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *WhatsAppTemplateUrlButtonWhatsAppApiContentAllOf) GetParameter() string {
	if o == nil || o.Parameter == nil {
		var ret string
		return ret
	}
	return *o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsAppTemplateUrlButtonWhatsAppApiContentAllOf) GetParameterOk() (*string, bool) {
	if o == nil || o.Parameter == nil {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *WhatsAppTemplateUrlButtonWhatsAppApiContentAllOf) HasParameter() bool {
	if o != nil && o.Parameter != nil {
		return true
	}

	return false
}

// SetParameter gets a reference to the given string and assigns it to the Parameter field.
func (o *WhatsAppTemplateUrlButtonWhatsAppApiContentAllOf) SetParameter(v string) {
	o.Parameter = &v
}

func (o WhatsAppTemplateUrlButtonWhatsAppApiContentAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Parameter != nil {
		toSerialize["parameter"] = o.Parameter
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsAppTemplateUrlButtonWhatsAppApiContentAllOf struct {
	value *WhatsAppTemplateUrlButtonWhatsAppApiContentAllOf
	isSet bool
}

func (v NullableWhatsAppTemplateUrlButtonWhatsAppApiContentAllOf) Get() *WhatsAppTemplateUrlButtonWhatsAppApiContentAllOf {
	return v.value
}

func (v *NullableWhatsAppTemplateUrlButtonWhatsAppApiContentAllOf) Set(val *WhatsAppTemplateUrlButtonWhatsAppApiContentAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsAppTemplateUrlButtonWhatsAppApiContentAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsAppTemplateUrlButtonWhatsAppApiContentAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsAppTemplateUrlButtonWhatsAppApiContentAllOf(val *WhatsAppTemplateUrlButtonWhatsAppApiContentAllOf) *NullableWhatsAppTemplateUrlButtonWhatsAppApiContentAllOf {
	return &NullableWhatsAppTemplateUrlButtonWhatsAppApiContentAllOf{value: val, isSet: true}
}

func (v NullableWhatsAppTemplateUrlButtonWhatsAppApiContentAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsAppTemplateUrlButtonWhatsAppApiContentAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
