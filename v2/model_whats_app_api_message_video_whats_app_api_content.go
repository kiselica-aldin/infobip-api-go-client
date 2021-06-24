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

// WhatsAppApiMessageVideoWhatsAppApiContent struct for WhatsAppApiMessageVideoWhatsAppApiContent
type WhatsAppApiMessageVideoWhatsAppApiContent struct {
	// Registered WhatsApp sender number. Must be in international format.
	From string `json:"from"`
	// Message recipient number. Must be in international format.
	To string `json:"to"`
	// The ID that uniquely identifies the message sent.
	MessageId *string                         `json:"messageId,omitempty"`
	Content   WhatsAppVideoWhatsAppApiContent `json:"content"`
	// Custom client data that will be included in Delivery Report.
	CallbackData *string `json:"callbackData,omitempty"`
}

// NewWhatsAppApiMessageVideoWhatsAppApiContent instantiates a new WhatsAppApiMessageVideoWhatsAppApiContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsAppApiMessageVideoWhatsAppApiContent(from string, to string, content WhatsAppVideoWhatsAppApiContent) *WhatsAppApiMessageVideoWhatsAppApiContent {
	this := WhatsAppApiMessageVideoWhatsAppApiContent{}
	this.From = from
	this.To = to
	this.Content = content
	return &this
}

// NewWhatsAppApiMessageVideoWhatsAppApiContentWithDefaults instantiates a new WhatsAppApiMessageVideoWhatsAppApiContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsAppApiMessageVideoWhatsAppApiContentWithDefaults() *WhatsAppApiMessageVideoWhatsAppApiContent {
	this := WhatsAppApiMessageVideoWhatsAppApiContent{}
	return &this
}

// GetFrom returns the From field value
func (o *WhatsAppApiMessageVideoWhatsAppApiContent) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *WhatsAppApiMessageVideoWhatsAppApiContent) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *WhatsAppApiMessageVideoWhatsAppApiContent) SetFrom(v string) {
	o.From = v
}

// GetTo returns the To field value
func (o *WhatsAppApiMessageVideoWhatsAppApiContent) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *WhatsAppApiMessageVideoWhatsAppApiContent) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *WhatsAppApiMessageVideoWhatsAppApiContent) SetTo(v string) {
	o.To = v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *WhatsAppApiMessageVideoWhatsAppApiContent) GetMessageId() string {
	if o == nil || o.MessageId == nil {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsAppApiMessageVideoWhatsAppApiContent) GetMessageIdOk() (*string, bool) {
	if o == nil || o.MessageId == nil {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *WhatsAppApiMessageVideoWhatsAppApiContent) HasMessageId() bool {
	if o != nil && o.MessageId != nil {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *WhatsAppApiMessageVideoWhatsAppApiContent) SetMessageId(v string) {
	o.MessageId = &v
}

// GetContent returns the Content field value
func (o *WhatsAppApiMessageVideoWhatsAppApiContent) GetContent() WhatsAppVideoWhatsAppApiContent {
	if o == nil {
		var ret WhatsAppVideoWhatsAppApiContent
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *WhatsAppApiMessageVideoWhatsAppApiContent) GetContentOk() (*WhatsAppVideoWhatsAppApiContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *WhatsAppApiMessageVideoWhatsAppApiContent) SetContent(v WhatsAppVideoWhatsAppApiContent) {
	o.Content = v
}

// GetCallbackData returns the CallbackData field value if set, zero value otherwise.
func (o *WhatsAppApiMessageVideoWhatsAppApiContent) GetCallbackData() string {
	if o == nil || o.CallbackData == nil {
		var ret string
		return ret
	}
	return *o.CallbackData
}

// GetCallbackDataOk returns a tuple with the CallbackData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsAppApiMessageVideoWhatsAppApiContent) GetCallbackDataOk() (*string, bool) {
	if o == nil || o.CallbackData == nil {
		return nil, false
	}
	return o.CallbackData, true
}

// HasCallbackData returns a boolean if a field has been set.
func (o *WhatsAppApiMessageVideoWhatsAppApiContent) HasCallbackData() bool {
	if o != nil && o.CallbackData != nil {
		return true
	}

	return false
}

// SetCallbackData gets a reference to the given string and assigns it to the CallbackData field.
func (o *WhatsAppApiMessageVideoWhatsAppApiContent) SetCallbackData(v string) {
	o.CallbackData = &v
}

func (o WhatsAppApiMessageVideoWhatsAppApiContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["from"] = o.From
	}
	if true {
		toSerialize["to"] = o.To
	}
	if o.MessageId != nil {
		toSerialize["messageId"] = o.MessageId
	}
	if true {
		toSerialize["content"] = o.Content
	}
	if o.CallbackData != nil {
		toSerialize["callbackData"] = o.CallbackData
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsAppApiMessageVideoWhatsAppApiContent struct {
	value *WhatsAppApiMessageVideoWhatsAppApiContent
	isSet bool
}

func (v NullableWhatsAppApiMessageVideoWhatsAppApiContent) Get() *WhatsAppApiMessageVideoWhatsAppApiContent {
	return v.value
}

func (v *NullableWhatsAppApiMessageVideoWhatsAppApiContent) Set(val *WhatsAppApiMessageVideoWhatsAppApiContent) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsAppApiMessageVideoWhatsAppApiContent) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsAppApiMessageVideoWhatsAppApiContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsAppApiMessageVideoWhatsAppApiContent(val *WhatsAppApiMessageVideoWhatsAppApiContent) *NullableWhatsAppApiMessageVideoWhatsAppApiContent {
	return &NullableWhatsAppApiMessageVideoWhatsAppApiContent{value: val, isSet: true}
}

func (v NullableWhatsAppApiMessageVideoWhatsAppApiContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsAppApiMessageVideoWhatsAppApiContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}