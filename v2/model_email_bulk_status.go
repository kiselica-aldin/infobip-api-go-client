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
	"fmt"
)

// EmailBulkStatus the model 'EmailBulkStatus'
type EmailBulkStatus string

// List of EmailBulkStatus
const (
	EMAILBULKSTATUS_PENDING    EmailBulkStatus = "PENDING"
	EMAILBULKSTATUS_PAUSED     EmailBulkStatus = "PAUSED"
	EMAILBULKSTATUS_PROCESSING EmailBulkStatus = "PROCESSING"
	EMAILBULKSTATUS_CANCELED   EmailBulkStatus = "CANCELED"
	EMAILBULKSTATUS_FINISHED   EmailBulkStatus = "FINISHED"
	EMAILBULKSTATUS_FAILED     EmailBulkStatus = "FAILED"
)

func (v *EmailBulkStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EmailBulkStatus(value)
	for _, existing := range []EmailBulkStatus{"PENDING", "PAUSED", "PROCESSING", "CANCELED", "FINISHED", "FAILED"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EmailBulkStatus", value)
}

// Ptr returns reference to EmailBulkStatus value
func (v EmailBulkStatus) Ptr() *EmailBulkStatus {
	return &v
}

type NullableEmailBulkStatus struct {
	value *EmailBulkStatus
	isSet bool
}

func (v NullableEmailBulkStatus) Get() *EmailBulkStatus {
	return v.value
}

func (v *NullableEmailBulkStatus) Set(val *EmailBulkStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailBulkStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailBulkStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailBulkStatus(val *EmailBulkStatus) *NullableEmailBulkStatus {
	return &NullableEmailBulkStatus{value: val, isSet: true}
}

func (v NullableEmailBulkStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailBulkStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
