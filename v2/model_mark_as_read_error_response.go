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

// MarkAsReadErrorResponse struct for MarkAsReadErrorResponse
type MarkAsReadErrorResponse struct {
	// Error details
	Error *string `json:"error,omitempty"`
}

// NewMarkAsReadErrorResponse instantiates a new MarkAsReadErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarkAsReadErrorResponse() *MarkAsReadErrorResponse {
	this := MarkAsReadErrorResponse{}
	return &this
}

// NewMarkAsReadErrorResponseWithDefaults instantiates a new MarkAsReadErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarkAsReadErrorResponseWithDefaults() *MarkAsReadErrorResponse {
	this := MarkAsReadErrorResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *MarkAsReadErrorResponse) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkAsReadErrorResponse) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *MarkAsReadErrorResponse) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *MarkAsReadErrorResponse) SetError(v string) {
	o.Error = &v
}

func (o MarkAsReadErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableMarkAsReadErrorResponse struct {
	value *MarkAsReadErrorResponse
	isSet bool
}

func (v NullableMarkAsReadErrorResponse) Get() *MarkAsReadErrorResponse {
	return v.value
}

func (v *NullableMarkAsReadErrorResponse) Set(val *MarkAsReadErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMarkAsReadErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMarkAsReadErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarkAsReadErrorResponse(val *MarkAsReadErrorResponse) *NullableMarkAsReadErrorResponse {
	return &NullableMarkAsReadErrorResponse{value: val, isSet: true}
}

func (v NullableMarkAsReadErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarkAsReadErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
