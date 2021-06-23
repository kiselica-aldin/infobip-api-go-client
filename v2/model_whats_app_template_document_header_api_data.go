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

// WhatsAppTemplateDocumentHeaderApiData struct for WhatsAppTemplateDocumentHeaderApiData
type WhatsAppTemplateDocumentHeaderApiData struct {
	WhatsAppTemplateHeaderApiData
}

// NewWhatsAppTemplateDocumentHeaderApiData instantiates a new WhatsAppTemplateDocumentHeaderApiData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsAppTemplateDocumentHeaderApiData() *WhatsAppTemplateDocumentHeaderApiData {
	this := WhatsAppTemplateDocumentHeaderApiData{}
	return &this
}

// NewWhatsAppTemplateDocumentHeaderApiDataWithDefaults instantiates a new WhatsAppTemplateDocumentHeaderApiData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsAppTemplateDocumentHeaderApiDataWithDefaults() *WhatsAppTemplateDocumentHeaderApiData {
	this := WhatsAppTemplateDocumentHeaderApiData{}
	return &this
}

func (o WhatsAppTemplateDocumentHeaderApiData) MarshalJSON() ([]byte, error) {
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

type NullableWhatsAppTemplateDocumentHeaderApiData struct {
	value *WhatsAppTemplateDocumentHeaderApiData
	isSet bool
}

func (v NullableWhatsAppTemplateDocumentHeaderApiData) Get() *WhatsAppTemplateDocumentHeaderApiData {
	return v.value
}

func (v *NullableWhatsAppTemplateDocumentHeaderApiData) Set(val *WhatsAppTemplateDocumentHeaderApiData) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsAppTemplateDocumentHeaderApiData) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsAppTemplateDocumentHeaderApiData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsAppTemplateDocumentHeaderApiData(val *WhatsAppTemplateDocumentHeaderApiData) *NullableWhatsAppTemplateDocumentHeaderApiData {
	return &NullableWhatsAppTemplateDocumentHeaderApiData{value: val, isSet: true}
}

func (v NullableWhatsAppTemplateDocumentHeaderApiData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsAppTemplateDocumentHeaderApiData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
