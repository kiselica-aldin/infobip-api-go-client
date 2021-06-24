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

// WhatsAppStickerWhatsAppApiContent Content of the message that will be sent.
type WhatsAppStickerWhatsAppApiContent struct {
	// URL of a sticker sent in the WhatsApp message. It is expected to be a valid URL starting with `https://`, `http://` or `ftp://`. Supported sticker type is `WebP`. Sticker file should be 512x512 pixels. Maximum sticker size is 100KB.
	MediaUrl string `json:"mediaUrl"`
}

// NewWhatsAppStickerWhatsAppApiContent instantiates a new WhatsAppStickerWhatsAppApiContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsAppStickerWhatsAppApiContent(mediaUrl string) *WhatsAppStickerWhatsAppApiContent {
	this := WhatsAppStickerWhatsAppApiContent{}
	this.MediaUrl = mediaUrl
	return &this
}

// NewWhatsAppStickerWhatsAppApiContentWithDefaults instantiates a new WhatsAppStickerWhatsAppApiContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsAppStickerWhatsAppApiContentWithDefaults() *WhatsAppStickerWhatsAppApiContent {
	this := WhatsAppStickerWhatsAppApiContent{}
	return &this
}

// GetMediaUrl returns the MediaUrl field value
func (o *WhatsAppStickerWhatsAppApiContent) GetMediaUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MediaUrl
}

// GetMediaUrlOk returns a tuple with the MediaUrl field value
// and a boolean to check if the value has been set.
func (o *WhatsAppStickerWhatsAppApiContent) GetMediaUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaUrl, true
}

// SetMediaUrl sets field value
func (o *WhatsAppStickerWhatsAppApiContent) SetMediaUrl(v string) {
	o.MediaUrl = v
}

func (o WhatsAppStickerWhatsAppApiContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mediaUrl"] = o.MediaUrl
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsAppStickerWhatsAppApiContent struct {
	value *WhatsAppStickerWhatsAppApiContent
	isSet bool
}

func (v NullableWhatsAppStickerWhatsAppApiContent) Get() *WhatsAppStickerWhatsAppApiContent {
	return v.value
}

func (v *NullableWhatsAppStickerWhatsAppApiContent) Set(val *WhatsAppStickerWhatsAppApiContent) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsAppStickerWhatsAppApiContent) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsAppStickerWhatsAppApiContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsAppStickerWhatsAppApiContent(val *WhatsAppStickerWhatsAppApiContent) *NullableWhatsAppStickerWhatsAppApiContent {
	return &NullableWhatsAppStickerWhatsAppApiContent{value: val, isSet: true}
}

func (v NullableWhatsAppStickerWhatsAppApiContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsAppStickerWhatsAppApiContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}