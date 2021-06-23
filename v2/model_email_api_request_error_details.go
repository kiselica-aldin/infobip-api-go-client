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

// EmailApiRequestErrorDetails struct for EmailApiRequestErrorDetails
type EmailApiRequestErrorDetails struct {
	// Identifier of the error.
	MessageId *string `json:"messageId,omitempty"`
	// Detailed error description.
	Text *string `json:"text,omitempty"`
}

// NewEmailApiRequestErrorDetails instantiates a new EmailApiRequestErrorDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailApiRequestErrorDetails() *EmailApiRequestErrorDetails {
	this := EmailApiRequestErrorDetails{}
	return &this
}

// NewEmailApiRequestErrorDetailsWithDefaults instantiates a new EmailApiRequestErrorDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailApiRequestErrorDetailsWithDefaults() *EmailApiRequestErrorDetails {
	this := EmailApiRequestErrorDetails{}
	return &this
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *EmailApiRequestErrorDetails) GetMessageId() string {
	if o == nil || o.MessageId == nil {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailApiRequestErrorDetails) GetMessageIdOk() (*string, bool) {
	if o == nil || o.MessageId == nil {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *EmailApiRequestErrorDetails) HasMessageId() bool {
	if o != nil && o.MessageId != nil {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *EmailApiRequestErrorDetails) SetMessageId(v string) {
	o.MessageId = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *EmailApiRequestErrorDetails) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailApiRequestErrorDetails) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *EmailApiRequestErrorDetails) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *EmailApiRequestErrorDetails) SetText(v string) {
	o.Text = &v
}

func (o EmailApiRequestErrorDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MessageId != nil {
		toSerialize["messageId"] = o.MessageId
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	return json.Marshal(toSerialize)
}

type NullableEmailApiRequestErrorDetails struct {
	value *EmailApiRequestErrorDetails
	isSet bool
}

func (v NullableEmailApiRequestErrorDetails) Get() *EmailApiRequestErrorDetails {
	return v.value
}

func (v *NullableEmailApiRequestErrorDetails) Set(val *EmailApiRequestErrorDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailApiRequestErrorDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailApiRequestErrorDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailApiRequestErrorDetails(val *EmailApiRequestErrorDetails) *NullableEmailApiRequestErrorDetails {
	return &NullableEmailApiRequestErrorDetails{value: val, isSet: true}
}

func (v NullableEmailApiRequestErrorDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailApiRequestErrorDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
