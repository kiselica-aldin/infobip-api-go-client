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

// SmsUpdateStatusRequest struct for SmsUpdateStatusRequest
type SmsUpdateStatusRequest struct {
	Status SmsBulkStatus `json:"status"`
}

// NewSmsUpdateStatusRequest instantiates a new SmsUpdateStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsUpdateStatusRequest(status SmsBulkStatus) *SmsUpdateStatusRequest {
	this := SmsUpdateStatusRequest{}
	this.Status = status
	return &this
}

// NewSmsUpdateStatusRequestWithDefaults instantiates a new SmsUpdateStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsUpdateStatusRequestWithDefaults() *SmsUpdateStatusRequest {
	this := SmsUpdateStatusRequest{}
	return &this
}

// GetStatus returns the Status field value
func (o *SmsUpdateStatusRequest) GetStatus() SmsBulkStatus {
	if o == nil {
		var ret SmsBulkStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SmsUpdateStatusRequest) GetStatusOk() (*SmsBulkStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SmsUpdateStatusRequest) SetStatus(v SmsBulkStatus) {
	o.Status = v
}

func (o SmsUpdateStatusRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableSmsUpdateStatusRequest struct {
	value *SmsUpdateStatusRequest
	isSet bool
}

func (v NullableSmsUpdateStatusRequest) Get() *SmsUpdateStatusRequest {
	return v.value
}

func (v *NullableSmsUpdateStatusRequest) Set(val *SmsUpdateStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsUpdateStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsUpdateStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsUpdateStatusRequest(val *SmsUpdateStatusRequest) *NullableSmsUpdateStatusRequest {
	return &NullableSmsUpdateStatusRequest{value: val, isSet: true}
}

func (v NullableSmsUpdateStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsUpdateStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
