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

// EmailBulkUpdateStatusRequest struct for EmailBulkUpdateStatusRequest
type EmailBulkUpdateStatusRequest struct {
	Status EmailBulkStatus `json:"status"`
}

// NewEmailBulkUpdateStatusRequest instantiates a new EmailBulkUpdateStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailBulkUpdateStatusRequest(status EmailBulkStatus) *EmailBulkUpdateStatusRequest {
	this := EmailBulkUpdateStatusRequest{}
	this.Status = status
	return &this
}

// NewEmailBulkUpdateStatusRequestWithDefaults instantiates a new EmailBulkUpdateStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailBulkUpdateStatusRequestWithDefaults() *EmailBulkUpdateStatusRequest {
	this := EmailBulkUpdateStatusRequest{}
	return &this
}

// GetStatus returns the Status field value
func (o *EmailBulkUpdateStatusRequest) GetStatus() EmailBulkStatus {
	if o == nil {
		var ret EmailBulkStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EmailBulkUpdateStatusRequest) GetStatusOk() (*EmailBulkStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *EmailBulkUpdateStatusRequest) SetStatus(v EmailBulkStatus) {
	o.Status = v
}

func (o EmailBulkUpdateStatusRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableEmailBulkUpdateStatusRequest struct {
	value *EmailBulkUpdateStatusRequest
	isSet bool
}

func (v NullableEmailBulkUpdateStatusRequest) Get() *EmailBulkUpdateStatusRequest {
	return v.value
}

func (v *NullableEmailBulkUpdateStatusRequest) Set(val *EmailBulkUpdateStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailBulkUpdateStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailBulkUpdateStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailBulkUpdateStatusRequest(val *EmailBulkUpdateStatusRequest) *NullableEmailBulkUpdateStatusRequest {
	return &NullableEmailBulkUpdateStatusRequest{value: val, isSet: true}
}

func (v NullableEmailBulkUpdateStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailBulkUpdateStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
