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

// WhatsAppTemplateImageHeaderApiData struct for WhatsAppTemplateImageHeaderApiData
type WhatsAppTemplateImageHeaderApiData struct {
	WhatsAppTemplateHeaderApiData
}

// NewWhatsAppTemplateImageHeaderApiData instantiates a new WhatsAppTemplateImageHeaderApiData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsAppTemplateImageHeaderApiData() *WhatsAppTemplateImageHeaderApiData {
	this := WhatsAppTemplateImageHeaderApiData{}
	return &this
}

// NewWhatsAppTemplateImageHeaderApiDataWithDefaults instantiates a new WhatsAppTemplateImageHeaderApiData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsAppTemplateImageHeaderApiDataWithDefaults() *WhatsAppTemplateImageHeaderApiData {
	this := WhatsAppTemplateImageHeaderApiData{}
	return &this
}

func (o WhatsAppTemplateImageHeaderApiData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedWhatsAppTemplateHeaderApiData, errWhatsAppTemplateHeaderApiData := json.Marshal(o.WhatsAppTemplateHeaderApiData)
	if errWhatsAppTemplateHeaderApiData != nil {
		return []byte{}, errWhatsAppTemplateHeaderApiData
	}
	errWhatsAppTemplateHeaderApiData = json.Unmarshal([]byte(serializedWhatsAppTemplateHeaderApiData), &toSerialize)
	if errWhatsAppTemplateHeaderApiData != nil {
		return []byte{}, errWhatsAppTemplateHeaderApiData
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsAppTemplateImageHeaderApiData struct {
	value *WhatsAppTemplateImageHeaderApiData
	isSet bool
}

func (v NullableWhatsAppTemplateImageHeaderApiData) Get() *WhatsAppTemplateImageHeaderApiData {
	return v.value
}

func (v *NullableWhatsAppTemplateImageHeaderApiData) Set(val *WhatsAppTemplateImageHeaderApiData) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsAppTemplateImageHeaderApiData) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsAppTemplateImageHeaderApiData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsAppTemplateImageHeaderApiData(val *WhatsAppTemplateImageHeaderApiData) *NullableWhatsAppTemplateImageHeaderApiData {
	return &NullableWhatsAppTemplateImageHeaderApiData{value: val, isSet: true}
}

func (v NullableWhatsAppTemplateImageHeaderApiData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsAppTemplateImageHeaderApiData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
