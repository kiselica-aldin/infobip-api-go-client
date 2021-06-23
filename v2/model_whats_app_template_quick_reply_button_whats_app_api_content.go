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

// WhatsAppTemplateQuickReplyButtonWhatsAppApiContent struct for WhatsAppTemplateQuickReplyButtonWhatsAppApiContent
type WhatsAppTemplateQuickReplyButtonWhatsAppApiContent struct {
	WhatsAppTemplateButtonWhatsAppApiContent
	// Payload of a `quick reply` button.
	Parameter string `json:"parameter"`
}

// NewWhatsAppTemplateQuickReplyButtonWhatsAppApiContent instantiates a new WhatsAppTemplateQuickReplyButtonWhatsAppApiContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsAppTemplateQuickReplyButtonWhatsAppApiContent(parameter string, type_ string) *WhatsAppTemplateQuickReplyButtonWhatsAppApiContent {
	this := WhatsAppTemplateQuickReplyButtonWhatsAppApiContent{}
	this.Type = type_
	this.Parameter = parameter
	return &this
}

// NewWhatsAppTemplateQuickReplyButtonWhatsAppApiContentWithDefaults instantiates a new WhatsAppTemplateQuickReplyButtonWhatsAppApiContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsAppTemplateQuickReplyButtonWhatsAppApiContentWithDefaults() *WhatsAppTemplateQuickReplyButtonWhatsAppApiContent {
	this := WhatsAppTemplateQuickReplyButtonWhatsAppApiContent{}
	return &this
}

// GetParameter returns the Parameter field value
func (o *WhatsAppTemplateQuickReplyButtonWhatsAppApiContent) GetParameter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value
// and a boolean to check if the value has been set.
func (o *WhatsAppTemplateQuickReplyButtonWhatsAppApiContent) GetParameterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameter, true
}

// SetParameter sets field value
func (o *WhatsAppTemplateQuickReplyButtonWhatsAppApiContent) SetParameter(v string) {
	o.Parameter = v
}

func (o WhatsAppTemplateQuickReplyButtonWhatsAppApiContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedWhatsAppTemplateButtonWhatsAppApiContent, errWhatsAppTemplateButtonWhatsAppApiContent := json.Marshal(o.WhatsAppTemplateButtonWhatsAppApiContent)
	if errWhatsAppTemplateButtonWhatsAppApiContent != nil {
		return []byte{}, errWhatsAppTemplateButtonWhatsAppApiContent
	}
	errWhatsAppTemplateButtonWhatsAppApiContent = json.Unmarshal([]byte(serializedWhatsAppTemplateButtonWhatsAppApiContent), &toSerialize)
	if errWhatsAppTemplateButtonWhatsAppApiContent != nil {
		return []byte{}, errWhatsAppTemplateButtonWhatsAppApiContent
	}
	if true {
		toSerialize["parameter"] = o.Parameter
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsAppTemplateQuickReplyButtonWhatsAppApiContent struct {
	value *WhatsAppTemplateQuickReplyButtonWhatsAppApiContent
	isSet bool
}

func (v NullableWhatsAppTemplateQuickReplyButtonWhatsAppApiContent) Get() *WhatsAppTemplateQuickReplyButtonWhatsAppApiContent {
	return v.value
}

func (v *NullableWhatsAppTemplateQuickReplyButtonWhatsAppApiContent) Set(val *WhatsAppTemplateQuickReplyButtonWhatsAppApiContent) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsAppTemplateQuickReplyButtonWhatsAppApiContent) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsAppTemplateQuickReplyButtonWhatsAppApiContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsAppTemplateQuickReplyButtonWhatsAppApiContent(val *WhatsAppTemplateQuickReplyButtonWhatsAppApiContent) *NullableWhatsAppTemplateQuickReplyButtonWhatsAppApiContent {
	return &NullableWhatsAppTemplateQuickReplyButtonWhatsAppApiContent{value: val, isSet: true}
}

func (v NullableWhatsAppTemplateQuickReplyButtonWhatsAppApiContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsAppTemplateQuickReplyButtonWhatsAppApiContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
