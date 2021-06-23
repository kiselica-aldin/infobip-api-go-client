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

// WhatsAppSingleMessageInfo struct for WhatsAppSingleMessageInfo
type WhatsAppSingleMessageInfo struct {
	// Message destination.
	To *string `json:"to,omitempty"`
	// Number of messages required to deliver.
	MessageCount *int32 `json:"messageCount,omitempty"`
	// The ID that uniquely identifies the message sent.
	MessageId *string                      `json:"messageId,omitempty"`
	Status    *WhatsAppSingleMessageStatus `json:"status,omitempty"`
}

// NewWhatsAppSingleMessageInfo instantiates a new WhatsAppSingleMessageInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsAppSingleMessageInfo() *WhatsAppSingleMessageInfo {
	this := WhatsAppSingleMessageInfo{}
	return &this
}

// NewWhatsAppSingleMessageInfoWithDefaults instantiates a new WhatsAppSingleMessageInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsAppSingleMessageInfoWithDefaults() *WhatsAppSingleMessageInfo {
	this := WhatsAppSingleMessageInfo{}
	return &this
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *WhatsAppSingleMessageInfo) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsAppSingleMessageInfo) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *WhatsAppSingleMessageInfo) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *WhatsAppSingleMessageInfo) SetTo(v string) {
	o.To = &v
}

// GetMessageCount returns the MessageCount field value if set, zero value otherwise.
func (o *WhatsAppSingleMessageInfo) GetMessageCount() int32 {
	if o == nil || o.MessageCount == nil {
		var ret int32
		return ret
	}
	return *o.MessageCount
}

// GetMessageCountOk returns a tuple with the MessageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsAppSingleMessageInfo) GetMessageCountOk() (*int32, bool) {
	if o == nil || o.MessageCount == nil {
		return nil, false
	}
	return o.MessageCount, true
}

// HasMessageCount returns a boolean if a field has been set.
func (o *WhatsAppSingleMessageInfo) HasMessageCount() bool {
	if o != nil && o.MessageCount != nil {
		return true
	}

	return false
}

// SetMessageCount gets a reference to the given int32 and assigns it to the MessageCount field.
func (o *WhatsAppSingleMessageInfo) SetMessageCount(v int32) {
	o.MessageCount = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *WhatsAppSingleMessageInfo) GetMessageId() string {
	if o == nil || o.MessageId == nil {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsAppSingleMessageInfo) GetMessageIdOk() (*string, bool) {
	if o == nil || o.MessageId == nil {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *WhatsAppSingleMessageInfo) HasMessageId() bool {
	if o != nil && o.MessageId != nil {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *WhatsAppSingleMessageInfo) SetMessageId(v string) {
	o.MessageId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WhatsAppSingleMessageInfo) GetStatus() WhatsAppSingleMessageStatus {
	if o == nil || o.Status == nil {
		var ret WhatsAppSingleMessageStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsAppSingleMessageInfo) GetStatusOk() (*WhatsAppSingleMessageStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WhatsAppSingleMessageInfo) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given WhatsAppSingleMessageStatus and assigns it to the Status field.
func (o *WhatsAppSingleMessageInfo) SetStatus(v WhatsAppSingleMessageStatus) {
	o.Status = &v
}

func (o WhatsAppSingleMessageInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	if o.MessageCount != nil {
		toSerialize["messageCount"] = o.MessageCount
	}
	if o.MessageId != nil {
		toSerialize["messageId"] = o.MessageId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsAppSingleMessageInfo struct {
	value *WhatsAppSingleMessageInfo
	isSet bool
}

func (v NullableWhatsAppSingleMessageInfo) Get() *WhatsAppSingleMessageInfo {
	return v.value
}

func (v *NullableWhatsAppSingleMessageInfo) Set(val *WhatsAppSingleMessageInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsAppSingleMessageInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsAppSingleMessageInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsAppSingleMessageInfo(val *WhatsAppSingleMessageInfo) *NullableWhatsAppSingleMessageInfo {
	return &NullableWhatsAppSingleMessageInfo{value: val, isSet: true}
}

func (v NullableWhatsAppSingleMessageInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsAppSingleMessageInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
