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

// WhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf struct for WhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf
type WhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf struct {
	// URL of a document sent in the header. It is expected to be a valid URL starting with `https://`, `http://` or `ftp://`. Supported document type is `PDF`. Maximum document size is 100MB.
	MediaUrl *string `json:"mediaUrl,omitempty"`
	// Filename of the document.
	Filename *string `json:"filename,omitempty"`
}

// NewWhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf instantiates a new WhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf() *WhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf {
	this := WhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf{}
	return &this
}

// NewWhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOfWithDefaults instantiates a new WhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOfWithDefaults() *WhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf {
	this := WhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf{}
	return &this
}

// GetMediaUrl returns the MediaUrl field value if set, zero value otherwise.
func (o *WhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf) GetMediaUrl() string {
	if o == nil || o.MediaUrl == nil {
		var ret string
		return ret
	}
	return *o.MediaUrl
}

// GetMediaUrlOk returns a tuple with the MediaUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf) GetMediaUrlOk() (*string, bool) {
	if o == nil || o.MediaUrl == nil {
		return nil, false
	}
	return o.MediaUrl, true
}

// HasMediaUrl returns a boolean if a field has been set.
func (o *WhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf) HasMediaUrl() bool {
	if o != nil && o.MediaUrl != nil {
		return true
	}

	return false
}

// SetMediaUrl gets a reference to the given string and assigns it to the MediaUrl field.
func (o *WhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf) SetMediaUrl(v string) {
	o.MediaUrl = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *WhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *WhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *WhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf) SetFilename(v string) {
	o.Filename = &v
}

func (o WhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MediaUrl != nil {
		toSerialize["mediaUrl"] = o.MediaUrl
	}
	if o.Filename != nil {
		toSerialize["filename"] = o.Filename
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf struct {
	value *WhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf
	isSet bool
}

func (v NullableWhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf) Get() *WhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf {
	return v.value
}

func (v *NullableWhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf) Set(val *WhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf(val *WhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf) *NullableWhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf {
	return &NullableWhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf{value: val, isSet: true}
}

func (v NullableWhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsAppTemplateDocumentHeaderWhatsAppApiContentAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
