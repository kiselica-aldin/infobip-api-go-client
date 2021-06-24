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

// EmailSendResponse struct for EmailSendResponse
type EmailSendResponse struct {
	BulkId   *string             `json:"bulkId,omitempty"`
	Messages *[]EmailMessageInfo `json:"messages,omitempty"`
}

// NewEmailSendResponse instantiates a new EmailSendResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailSendResponse() *EmailSendResponse {
	this := EmailSendResponse{}
	return &this
}

// NewEmailSendResponseWithDefaults instantiates a new EmailSendResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailSendResponseWithDefaults() *EmailSendResponse {
	this := EmailSendResponse{}
	return &this
}

// GetBulkId returns the BulkId field value if set, zero value otherwise.
func (o *EmailSendResponse) GetBulkId() string {
	if o == nil || o.BulkId == nil {
		var ret string
		return ret
	}
	return *o.BulkId
}

// GetBulkIdOk returns a tuple with the BulkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSendResponse) GetBulkIdOk() (*string, bool) {
	if o == nil || o.BulkId == nil {
		return nil, false
	}
	return o.BulkId, true
}

// HasBulkId returns a boolean if a field has been set.
func (o *EmailSendResponse) HasBulkId() bool {
	if o != nil && o.BulkId != nil {
		return true
	}

	return false
}

// SetBulkId gets a reference to the given string and assigns it to the BulkId field.
func (o *EmailSendResponse) SetBulkId(v string) {
	o.BulkId = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *EmailSendResponse) GetMessages() []EmailMessageInfo {
	if o == nil || o.Messages == nil {
		var ret []EmailMessageInfo
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailSendResponse) GetMessagesOk() (*[]EmailMessageInfo, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *EmailSendResponse) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []EmailMessageInfo and assigns it to the Messages field.
func (o *EmailSendResponse) SetMessages(v []EmailMessageInfo) {
	o.Messages = &v
}

func (o EmailSendResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BulkId != nil {
		toSerialize["bulkId"] = o.BulkId
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	return json.Marshal(toSerialize)
}

type NullableEmailSendResponse struct {
	value *EmailSendResponse
	isSet bool
}

func (v NullableEmailSendResponse) Get() *EmailSendResponse {
	return v.value
}

func (v *NullableEmailSendResponse) Set(val *EmailSendResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailSendResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailSendResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailSendResponse(val *EmailSendResponse) *NullableEmailSendResponse {
	return &NullableEmailSendResponse{value: val, isSet: true}
}

func (v NullableEmailSendResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailSendResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}