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

// WhatsAppTemplateTextHeaderApiDataAllOf struct for WhatsAppTemplateTextHeaderApiDataAllOf
type WhatsAppTemplateTextHeaderApiDataAllOf struct {
	// Template header text. Can contain up to 60 characters, with one placeholder {{1}}.
	Text *string `json:"text,omitempty"`
}

// NewWhatsAppTemplateTextHeaderApiDataAllOf instantiates a new WhatsAppTemplateTextHeaderApiDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsAppTemplateTextHeaderApiDataAllOf() *WhatsAppTemplateTextHeaderApiDataAllOf {
	this := WhatsAppTemplateTextHeaderApiDataAllOf{}
	return &this
}

// NewWhatsAppTemplateTextHeaderApiDataAllOfWithDefaults instantiates a new WhatsAppTemplateTextHeaderApiDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsAppTemplateTextHeaderApiDataAllOfWithDefaults() *WhatsAppTemplateTextHeaderApiDataAllOf {
	this := WhatsAppTemplateTextHeaderApiDataAllOf{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *WhatsAppTemplateTextHeaderApiDataAllOf) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsAppTemplateTextHeaderApiDataAllOf) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *WhatsAppTemplateTextHeaderApiDataAllOf) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *WhatsAppTemplateTextHeaderApiDataAllOf) SetText(v string) {
	o.Text = &v
}

func (o WhatsAppTemplateTextHeaderApiDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsAppTemplateTextHeaderApiDataAllOf struct {
	value *WhatsAppTemplateTextHeaderApiDataAllOf
	isSet bool
}

func (v NullableWhatsAppTemplateTextHeaderApiDataAllOf) Get() *WhatsAppTemplateTextHeaderApiDataAllOf {
	return v.value
}

func (v *NullableWhatsAppTemplateTextHeaderApiDataAllOf) Set(val *WhatsAppTemplateTextHeaderApiDataAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsAppTemplateTextHeaderApiDataAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsAppTemplateTextHeaderApiDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsAppTemplateTextHeaderApiDataAllOf(val *WhatsAppTemplateTextHeaderApiDataAllOf) *NullableWhatsAppTemplateTextHeaderApiDataAllOf {
	return &NullableWhatsAppTemplateTextHeaderApiDataAllOf{value: val, isSet: true}
}

func (v NullableWhatsAppTemplateTextHeaderApiDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsAppTemplateTextHeaderApiDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
