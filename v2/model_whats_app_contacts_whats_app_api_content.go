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

// WhatsAppContactsWhatsAppApiContent Content of the message that will be sent.
type WhatsAppContactsWhatsAppApiContent struct {
	// Array of contacts sent in the WhatsApp message.
	Contacts []WhatsAppContactWhatsAppApiContent `json:"contacts"`
}

// NewWhatsAppContactsWhatsAppApiContent instantiates a new WhatsAppContactsWhatsAppApiContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsAppContactsWhatsAppApiContent(contacts []WhatsAppContactWhatsAppApiContent) *WhatsAppContactsWhatsAppApiContent {
	this := WhatsAppContactsWhatsAppApiContent{}
	this.Contacts = contacts
	return &this
}

// NewWhatsAppContactsWhatsAppApiContentWithDefaults instantiates a new WhatsAppContactsWhatsAppApiContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsAppContactsWhatsAppApiContentWithDefaults() *WhatsAppContactsWhatsAppApiContent {
	this := WhatsAppContactsWhatsAppApiContent{}
	return &this
}

// GetContacts returns the Contacts field value
func (o *WhatsAppContactsWhatsAppApiContent) GetContacts() []WhatsAppContactWhatsAppApiContent {
	if o == nil {
		var ret []WhatsAppContactWhatsAppApiContent
		return ret
	}

	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value
// and a boolean to check if the value has been set.
func (o *WhatsAppContactsWhatsAppApiContent) GetContactsOk() (*[]WhatsAppContactWhatsAppApiContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contacts, true
}

// SetContacts sets field value
func (o *WhatsAppContactsWhatsAppApiContent) SetContacts(v []WhatsAppContactWhatsAppApiContent) {
	o.Contacts = v
}

func (o WhatsAppContactsWhatsAppApiContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contacts"] = o.Contacts
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsAppContactsWhatsAppApiContent struct {
	value *WhatsAppContactsWhatsAppApiContent
	isSet bool
}

func (v NullableWhatsAppContactsWhatsAppApiContent) Get() *WhatsAppContactsWhatsAppApiContent {
	return v.value
}

func (v *NullableWhatsAppContactsWhatsAppApiContent) Set(val *WhatsAppContactsWhatsAppApiContent) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsAppContactsWhatsAppApiContent) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsAppContactsWhatsAppApiContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsAppContactsWhatsAppApiContent(val *WhatsAppContactsWhatsAppApiContent) *NullableWhatsAppContactsWhatsAppApiContent {
	return &NullableWhatsAppContactsWhatsAppApiContent{value: val, isSet: true}
}

func (v NullableWhatsAppContactsWhatsAppApiContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsAppContactsWhatsAppApiContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
