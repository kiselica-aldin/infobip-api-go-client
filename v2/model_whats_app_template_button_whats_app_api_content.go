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

// WhatsAppTemplateButtonWhatsAppApiContent Template buttons. Should be defined in correct order, only if `quick reply` or `dynamic URL` buttons have been registered. It can have up to three `quick reply` buttons or only one `dynamic URL` button.
type WhatsAppTemplateButtonWhatsAppApiContent struct {
	Type string `json:"type"`
}

// NewWhatsAppTemplateButtonWhatsAppApiContent instantiates a new WhatsAppTemplateButtonWhatsAppApiContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsAppTemplateButtonWhatsAppApiContent(type_ string) *WhatsAppTemplateButtonWhatsAppApiContent {
	this := WhatsAppTemplateButtonWhatsAppApiContent{}
	this.Type = type_
	return &this
}

// NewWhatsAppTemplateButtonWhatsAppApiContentWithDefaults instantiates a new WhatsAppTemplateButtonWhatsAppApiContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsAppTemplateButtonWhatsAppApiContentWithDefaults() *WhatsAppTemplateButtonWhatsAppApiContent {
	this := WhatsAppTemplateButtonWhatsAppApiContent{}
	return &this
}

// GetType returns the Type field value
func (o *WhatsAppTemplateButtonWhatsAppApiContent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WhatsAppTemplateButtonWhatsAppApiContent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WhatsAppTemplateButtonWhatsAppApiContent) SetType(v string) {
	o.Type = v
}

func (o WhatsAppTemplateButtonWhatsAppApiContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsAppTemplateButtonWhatsAppApiContent struct {
	value *WhatsAppTemplateButtonWhatsAppApiContent
	isSet bool
}

func (v NullableWhatsAppTemplateButtonWhatsAppApiContent) Get() *WhatsAppTemplateButtonWhatsAppApiContent {
	return v.value
}

func (v *NullableWhatsAppTemplateButtonWhatsAppApiContent) Set(val *WhatsAppTemplateButtonWhatsAppApiContent) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsAppTemplateButtonWhatsAppApiContent) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsAppTemplateButtonWhatsAppApiContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsAppTemplateButtonWhatsAppApiContent(val *WhatsAppTemplateButtonWhatsAppApiContent) *NullableWhatsAppTemplateButtonWhatsAppApiContent {
	return &NullableWhatsAppTemplateButtonWhatsAppApiContent{value: val, isSet: true}
}

func (v NullableWhatsAppTemplateButtonWhatsAppApiContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsAppTemplateButtonWhatsAppApiContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}