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

// WhatsAppTemplateQuickReplyButtonApiDataAllOf struct for WhatsAppTemplateQuickReplyButtonApiDataAllOf
type WhatsAppTemplateQuickReplyButtonApiDataAllOf struct {
	// Button text.
	Text *string `json:"text,omitempty"`
}

// NewWhatsAppTemplateQuickReplyButtonApiDataAllOf instantiates a new WhatsAppTemplateQuickReplyButtonApiDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsAppTemplateQuickReplyButtonApiDataAllOf() *WhatsAppTemplateQuickReplyButtonApiDataAllOf {
	this := WhatsAppTemplateQuickReplyButtonApiDataAllOf{}
	return &this
}

// NewWhatsAppTemplateQuickReplyButtonApiDataAllOfWithDefaults instantiates a new WhatsAppTemplateQuickReplyButtonApiDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsAppTemplateQuickReplyButtonApiDataAllOfWithDefaults() *WhatsAppTemplateQuickReplyButtonApiDataAllOf {
	this := WhatsAppTemplateQuickReplyButtonApiDataAllOf{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *WhatsAppTemplateQuickReplyButtonApiDataAllOf) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsAppTemplateQuickReplyButtonApiDataAllOf) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *WhatsAppTemplateQuickReplyButtonApiDataAllOf) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *WhatsAppTemplateQuickReplyButtonApiDataAllOf) SetText(v string) {
	o.Text = &v
}

func (o WhatsAppTemplateQuickReplyButtonApiDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsAppTemplateQuickReplyButtonApiDataAllOf struct {
	value *WhatsAppTemplateQuickReplyButtonApiDataAllOf
	isSet bool
}

func (v NullableWhatsAppTemplateQuickReplyButtonApiDataAllOf) Get() *WhatsAppTemplateQuickReplyButtonApiDataAllOf {
	return v.value
}

func (v *NullableWhatsAppTemplateQuickReplyButtonApiDataAllOf) Set(val *WhatsAppTemplateQuickReplyButtonApiDataAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsAppTemplateQuickReplyButtonApiDataAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsAppTemplateQuickReplyButtonApiDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsAppTemplateQuickReplyButtonApiDataAllOf(val *WhatsAppTemplateQuickReplyButtonApiDataAllOf) *NullableWhatsAppTemplateQuickReplyButtonApiDataAllOf {
	return &NullableWhatsAppTemplateQuickReplyButtonApiDataAllOf{value: val, isSet: true}
}

func (v NullableWhatsAppTemplateQuickReplyButtonApiDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsAppTemplateQuickReplyButtonApiDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
