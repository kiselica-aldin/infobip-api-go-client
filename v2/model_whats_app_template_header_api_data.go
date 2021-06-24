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

// WhatsAppTemplateHeaderApiData Template header. Can be `image`, `document`, `video`, `location` or `text`.
type WhatsAppTemplateHeaderApiData struct {
	Format *string `json:"format,omitempty"`
}

// NewWhatsAppTemplateHeaderApiData instantiates a new WhatsAppTemplateHeaderApiData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsAppTemplateHeaderApiData() *WhatsAppTemplateHeaderApiData {
	this := WhatsAppTemplateHeaderApiData{}
	return &this
}

// NewWhatsAppTemplateHeaderApiDataWithDefaults instantiates a new WhatsAppTemplateHeaderApiData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsAppTemplateHeaderApiDataWithDefaults() *WhatsAppTemplateHeaderApiData {
	this := WhatsAppTemplateHeaderApiData{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *WhatsAppTemplateHeaderApiData) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsAppTemplateHeaderApiData) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *WhatsAppTemplateHeaderApiData) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *WhatsAppTemplateHeaderApiData) SetFormat(v string) {
	o.Format = &v
}

func (o WhatsAppTemplateHeaderApiData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsAppTemplateHeaderApiData struct {
	value *WhatsAppTemplateHeaderApiData
	isSet bool
}

func (v NullableWhatsAppTemplateHeaderApiData) Get() *WhatsAppTemplateHeaderApiData {
	return v.value
}

func (v *NullableWhatsAppTemplateHeaderApiData) Set(val *WhatsAppTemplateHeaderApiData) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsAppTemplateHeaderApiData) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsAppTemplateHeaderApiData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsAppTemplateHeaderApiData(val *WhatsAppTemplateHeaderApiData) *NullableWhatsAppTemplateHeaderApiData {
	return &NullableWhatsAppTemplateHeaderApiData{value: val, isSet: true}
}

func (v NullableWhatsAppTemplateHeaderApiData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsAppTemplateHeaderApiData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}