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

// WhatsAppImageWhatsAppApiContent Content of the message that will be sent.
type WhatsAppImageWhatsAppApiContent struct {
	// URL of an image sent in the WhatsApp message. It is expected to be a valid URL starting with `https://`, `http://` or `ftp://`. Supported image types are `JPG`, `JPEG`, `PNG`. Maximum image size is 5MB.
	MediaUrl string `json:"mediaUrl"`
	// Caption of the image.
	Caption *string `json:"caption,omitempty"`
}

// NewWhatsAppImageWhatsAppApiContent instantiates a new WhatsAppImageWhatsAppApiContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsAppImageWhatsAppApiContent(mediaUrl string) *WhatsAppImageWhatsAppApiContent {
	this := WhatsAppImageWhatsAppApiContent{}
	this.MediaUrl = mediaUrl
	return &this
}

// NewWhatsAppImageWhatsAppApiContentWithDefaults instantiates a new WhatsAppImageWhatsAppApiContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsAppImageWhatsAppApiContentWithDefaults() *WhatsAppImageWhatsAppApiContent {
	this := WhatsAppImageWhatsAppApiContent{}
	return &this
}

// GetMediaUrl returns the MediaUrl field value
func (o *WhatsAppImageWhatsAppApiContent) GetMediaUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MediaUrl
}

// GetMediaUrlOk returns a tuple with the MediaUrl field value
// and a boolean to check if the value has been set.
func (o *WhatsAppImageWhatsAppApiContent) GetMediaUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaUrl, true
}

// SetMediaUrl sets field value
func (o *WhatsAppImageWhatsAppApiContent) SetMediaUrl(v string) {
	o.MediaUrl = v
}

// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *WhatsAppImageWhatsAppApiContent) GetCaption() string {
	if o == nil || o.Caption == nil {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsAppImageWhatsAppApiContent) GetCaptionOk() (*string, bool) {
	if o == nil || o.Caption == nil {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *WhatsAppImageWhatsAppApiContent) HasCaption() bool {
	if o != nil && o.Caption != nil {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *WhatsAppImageWhatsAppApiContent) SetCaption(v string) {
	o.Caption = &v
}

func (o WhatsAppImageWhatsAppApiContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mediaUrl"] = o.MediaUrl
	}
	if o.Caption != nil {
		toSerialize["caption"] = o.Caption
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsAppImageWhatsAppApiContent struct {
	value *WhatsAppImageWhatsAppApiContent
	isSet bool
}

func (v NullableWhatsAppImageWhatsAppApiContent) Get() *WhatsAppImageWhatsAppApiContent {
	return v.value
}

func (v *NullableWhatsAppImageWhatsAppApiContent) Set(val *WhatsAppImageWhatsAppApiContent) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsAppImageWhatsAppApiContent) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsAppImageWhatsAppApiContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsAppImageWhatsAppApiContent(val *WhatsAppImageWhatsAppApiContent) *NullableWhatsAppImageWhatsAppApiContent {
	return &NullableWhatsAppImageWhatsAppApiContent{value: val, isSet: true}
}

func (v NullableWhatsAppImageWhatsAppApiContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsAppImageWhatsAppApiContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
