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

// WhatsAppTemplateTextHeaderWhatsAppApiContentAllOf struct for WhatsAppTemplateTextHeaderWhatsAppApiContentAllOf
type WhatsAppTemplateTextHeaderWhatsAppApiContentAllOf struct {
	// Value of a placeholder in the text header.
	Placeholder *string `json:"placeholder,omitempty"`
}

// NewWhatsAppTemplateTextHeaderWhatsAppApiContentAllOf instantiates a new WhatsAppTemplateTextHeaderWhatsAppApiContentAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsAppTemplateTextHeaderWhatsAppApiContentAllOf() *WhatsAppTemplateTextHeaderWhatsAppApiContentAllOf {
	this := WhatsAppTemplateTextHeaderWhatsAppApiContentAllOf{}
	return &this
}

// NewWhatsAppTemplateTextHeaderWhatsAppApiContentAllOfWithDefaults instantiates a new WhatsAppTemplateTextHeaderWhatsAppApiContentAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsAppTemplateTextHeaderWhatsAppApiContentAllOfWithDefaults() *WhatsAppTemplateTextHeaderWhatsAppApiContentAllOf {
	this := WhatsAppTemplateTextHeaderWhatsAppApiContentAllOf{}
	return &this
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *WhatsAppTemplateTextHeaderWhatsAppApiContentAllOf) GetPlaceholder() string {
	if o == nil || o.Placeholder == nil {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsAppTemplateTextHeaderWhatsAppApiContentAllOf) GetPlaceholderOk() (*string, bool) {
	if o == nil || o.Placeholder == nil {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *WhatsAppTemplateTextHeaderWhatsAppApiContentAllOf) HasPlaceholder() bool {
	if o != nil && o.Placeholder != nil {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *WhatsAppTemplateTextHeaderWhatsAppApiContentAllOf) SetPlaceholder(v string) {
	o.Placeholder = &v
}

func (o WhatsAppTemplateTextHeaderWhatsAppApiContentAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Placeholder != nil {
		toSerialize["placeholder"] = o.Placeholder
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsAppTemplateTextHeaderWhatsAppApiContentAllOf struct {
	value *WhatsAppTemplateTextHeaderWhatsAppApiContentAllOf
	isSet bool
}

func (v NullableWhatsAppTemplateTextHeaderWhatsAppApiContentAllOf) Get() *WhatsAppTemplateTextHeaderWhatsAppApiContentAllOf {
	return v.value
}

func (v *NullableWhatsAppTemplateTextHeaderWhatsAppApiContentAllOf) Set(val *WhatsAppTemplateTextHeaderWhatsAppApiContentAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsAppTemplateTextHeaderWhatsAppApiContentAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsAppTemplateTextHeaderWhatsAppApiContentAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsAppTemplateTextHeaderWhatsAppApiContentAllOf(val *WhatsAppTemplateTextHeaderWhatsAppApiContentAllOf) *NullableWhatsAppTemplateTextHeaderWhatsAppApiContentAllOf {
	return &NullableWhatsAppTemplateTextHeaderWhatsAppApiContentAllOf{value: val, isSet: true}
}

func (v NullableWhatsAppTemplateTextHeaderWhatsAppApiContentAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsAppTemplateTextHeaderWhatsAppApiContentAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}