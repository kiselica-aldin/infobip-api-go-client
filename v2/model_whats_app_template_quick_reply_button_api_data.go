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

// WhatsAppTemplateQuickReplyButtonApiData struct for WhatsAppTemplateQuickReplyButtonApiData
type WhatsAppTemplateQuickReplyButtonApiData struct {
	WhatsAppTemplateButtonApiData
	// Button text.
	Text string `json:"text"`
}

// NewWhatsAppTemplateQuickReplyButtonApiData instantiates a new WhatsAppTemplateQuickReplyButtonApiData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsAppTemplateQuickReplyButtonApiData(text string) *WhatsAppTemplateQuickReplyButtonApiData {
	this := WhatsAppTemplateQuickReplyButtonApiData{}
	this.Text = text
	return &this
}

// NewWhatsAppTemplateQuickReplyButtonApiDataWithDefaults instantiates a new WhatsAppTemplateQuickReplyButtonApiData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsAppTemplateQuickReplyButtonApiDataWithDefaults() *WhatsAppTemplateQuickReplyButtonApiData {
	this := WhatsAppTemplateQuickReplyButtonApiData{}
	return &this
}

// GetText returns the Text field value
func (o *WhatsAppTemplateQuickReplyButtonApiData) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *WhatsAppTemplateQuickReplyButtonApiData) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *WhatsAppTemplateQuickReplyButtonApiData) SetText(v string) {
	o.Text = v
}

func (o WhatsAppTemplateQuickReplyButtonApiData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedWhatsAppTemplateButtonApiData, errWhatsAppTemplateButtonApiData := json.Marshal(o.WhatsAppTemplateButtonApiData)
	if errWhatsAppTemplateButtonApiData != nil {
		return []byte{}, errWhatsAppTemplateButtonApiData
	}
	errWhatsAppTemplateButtonApiData = json.Unmarshal([]byte(serializedWhatsAppTemplateButtonApiData), &toSerialize)
	if errWhatsAppTemplateButtonApiData != nil {
		return []byte{}, errWhatsAppTemplateButtonApiData
	}
	if true {
		toSerialize["text"] = o.Text
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsAppTemplateQuickReplyButtonApiData struct {
	value *WhatsAppTemplateQuickReplyButtonApiData
	isSet bool
}

func (v NullableWhatsAppTemplateQuickReplyButtonApiData) Get() *WhatsAppTemplateQuickReplyButtonApiData {
	return v.value
}

func (v *NullableWhatsAppTemplateQuickReplyButtonApiData) Set(val *WhatsAppTemplateQuickReplyButtonApiData) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsAppTemplateQuickReplyButtonApiData) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsAppTemplateQuickReplyButtonApiData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsAppTemplateQuickReplyButtonApiData(val *WhatsAppTemplateQuickReplyButtonApiData) *NullableWhatsAppTemplateQuickReplyButtonApiData {
	return &NullableWhatsAppTemplateQuickReplyButtonApiData{value: val, isSet: true}
}

func (v NullableWhatsAppTemplateQuickReplyButtonApiData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsAppTemplateQuickReplyButtonApiData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
