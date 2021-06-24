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

// WhatsAppEmailWhatsAppApiContent Array of emails information.
type WhatsAppEmailWhatsAppApiContent struct {
	// Contact's email.
	Email *string `json:"email,omitempty"`
	// Type of the email. Can be `HOME` or `WORK`.
	Type *string `json:"type,omitempty"`
}

// NewWhatsAppEmailWhatsAppApiContent instantiates a new WhatsAppEmailWhatsAppApiContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsAppEmailWhatsAppApiContent() *WhatsAppEmailWhatsAppApiContent {
	this := WhatsAppEmailWhatsAppApiContent{}
	return &this
}

// NewWhatsAppEmailWhatsAppApiContentWithDefaults instantiates a new WhatsAppEmailWhatsAppApiContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsAppEmailWhatsAppApiContentWithDefaults() *WhatsAppEmailWhatsAppApiContent {
	this := WhatsAppEmailWhatsAppApiContent{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *WhatsAppEmailWhatsAppApiContent) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsAppEmailWhatsAppApiContent) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *WhatsAppEmailWhatsAppApiContent) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *WhatsAppEmailWhatsAppApiContent) SetEmail(v string) {
	o.Email = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WhatsAppEmailWhatsAppApiContent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsAppEmailWhatsAppApiContent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WhatsAppEmailWhatsAppApiContent) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WhatsAppEmailWhatsAppApiContent) SetType(v string) {
	o.Type = &v
}

func (o WhatsAppEmailWhatsAppApiContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsAppEmailWhatsAppApiContent struct {
	value *WhatsAppEmailWhatsAppApiContent
	isSet bool
}

func (v NullableWhatsAppEmailWhatsAppApiContent) Get() *WhatsAppEmailWhatsAppApiContent {
	return v.value
}

func (v *NullableWhatsAppEmailWhatsAppApiContent) Set(val *WhatsAppEmailWhatsAppApiContent) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsAppEmailWhatsAppApiContent) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsAppEmailWhatsAppApiContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsAppEmailWhatsAppApiContent(val *WhatsAppEmailWhatsAppApiContent) *NullableWhatsAppEmailWhatsAppApiContent {
	return &NullableWhatsAppEmailWhatsAppApiContent{value: val, isSet: true}
}

func (v NullableWhatsAppEmailWhatsAppApiContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsAppEmailWhatsAppApiContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}