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

// WhatsAppTemplateTextHeaderWhatsAppApiContent struct for WhatsAppTemplateTextHeaderWhatsAppApiContent
type WhatsAppTemplateTextHeaderWhatsAppApiContent struct {
	WhatsAppTemplateHeaderWhatsAppApiContent
	// Value of a placeholder in the text header.
	Placeholder string `json:"placeholder"`
}

// NewWhatsAppTemplateTextHeaderWhatsAppApiContent instantiates a new WhatsAppTemplateTextHeaderWhatsAppApiContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsAppTemplateTextHeaderWhatsAppApiContent(placeholder string, type_ string) *WhatsAppTemplateTextHeaderWhatsAppApiContent {
	this := WhatsAppTemplateTextHeaderWhatsAppApiContent{}
	this.Type = type_
	this.Placeholder = placeholder
	return &this
}

// NewWhatsAppTemplateTextHeaderWhatsAppApiContentWithDefaults instantiates a new WhatsAppTemplateTextHeaderWhatsAppApiContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsAppTemplateTextHeaderWhatsAppApiContentWithDefaults() *WhatsAppTemplateTextHeaderWhatsAppApiContent {
	this := WhatsAppTemplateTextHeaderWhatsAppApiContent{}
	return &this
}

// GetPlaceholder returns the Placeholder field value
func (o *WhatsAppTemplateTextHeaderWhatsAppApiContent) GetPlaceholder() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value
// and a boolean to check if the value has been set.
func (o *WhatsAppTemplateTextHeaderWhatsAppApiContent) GetPlaceholderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Placeholder, true
}

// SetPlaceholder sets field value
func (o *WhatsAppTemplateTextHeaderWhatsAppApiContent) SetPlaceholder(v string) {
	o.Placeholder = v
}

func (o WhatsAppTemplateTextHeaderWhatsAppApiContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedWhatsAppTemplateHeaderWhatsAppApiContent, errWhatsAppTemplateHeaderWhatsAppApiContent := json.Marshal(o.WhatsAppTemplateHeaderWhatsAppApiContent)
	if errWhatsAppTemplateHeaderWhatsAppApiContent != nil {
		return []byte{}, errWhatsAppTemplateHeaderWhatsAppApiContent
	}
	errWhatsAppTemplateHeaderWhatsAppApiContent = json.Unmarshal([]byte(serializedWhatsAppTemplateHeaderWhatsAppApiContent), &toSerialize)
	if errWhatsAppTemplateHeaderWhatsAppApiContent != nil {
		return []byte{}, errWhatsAppTemplateHeaderWhatsAppApiContent
	}
	if true {
		toSerialize["placeholder"] = o.Placeholder
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsAppTemplateTextHeaderWhatsAppApiContent struct {
	value *WhatsAppTemplateTextHeaderWhatsAppApiContent
	isSet bool
}

func (v NullableWhatsAppTemplateTextHeaderWhatsAppApiContent) Get() *WhatsAppTemplateTextHeaderWhatsAppApiContent {
	return v.value
}

func (v *NullableWhatsAppTemplateTextHeaderWhatsAppApiContent) Set(val *WhatsAppTemplateTextHeaderWhatsAppApiContent) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsAppTemplateTextHeaderWhatsAppApiContent) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsAppTemplateTextHeaderWhatsAppApiContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsAppTemplateTextHeaderWhatsAppApiContent(val *WhatsAppTemplateTextHeaderWhatsAppApiContent) *NullableWhatsAppTemplateTextHeaderWhatsAppApiContent {
	return &NullableWhatsAppTemplateTextHeaderWhatsAppApiContent{value: val, isSet: true}
}

func (v NullableWhatsAppTemplateTextHeaderWhatsAppApiContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsAppTemplateTextHeaderWhatsAppApiContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
