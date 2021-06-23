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

// EmailApiRequestError struct for EmailApiRequestError
type EmailApiRequestError struct {
	ServiceException *EmailApiRequestErrorDetails `json:"serviceException,omitempty"`
}

// NewEmailApiRequestError instantiates a new EmailApiRequestError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailApiRequestError() *EmailApiRequestError {
	this := EmailApiRequestError{}
	return &this
}

// NewEmailApiRequestErrorWithDefaults instantiates a new EmailApiRequestError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailApiRequestErrorWithDefaults() *EmailApiRequestError {
	this := EmailApiRequestError{}
	return &this
}

// GetServiceException returns the ServiceException field value if set, zero value otherwise.
func (o *EmailApiRequestError) GetServiceException() EmailApiRequestErrorDetails {
	if o == nil || o.ServiceException == nil {
		var ret EmailApiRequestErrorDetails
		return ret
	}
	return *o.ServiceException
}

// GetServiceExceptionOk returns a tuple with the ServiceException field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailApiRequestError) GetServiceExceptionOk() (*EmailApiRequestErrorDetails, bool) {
	if o == nil || o.ServiceException == nil {
		return nil, false
	}
	return o.ServiceException, true
}

// HasServiceException returns a boolean if a field has been set.
func (o *EmailApiRequestError) HasServiceException() bool {
	if o != nil && o.ServiceException != nil {
		return true
	}

	return false
}

// SetServiceException gets a reference to the given EmailApiRequestErrorDetails and assigns it to the ServiceException field.
func (o *EmailApiRequestError) SetServiceException(v EmailApiRequestErrorDetails) {
	o.ServiceException = &v
}

func (o EmailApiRequestError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceException != nil {
		toSerialize["serviceException"] = o.ServiceException
	}
	return json.Marshal(toSerialize)
}

type NullableEmailApiRequestError struct {
	value *EmailApiRequestError
	isSet bool
}

func (v NullableEmailApiRequestError) Get() *EmailApiRequestError {
	return v.value
}

func (v *NullableEmailApiRequestError) Set(val *EmailApiRequestError) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailApiRequestError) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailApiRequestError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailApiRequestError(val *EmailApiRequestError) *NullableEmailApiRequestError {
	return &NullableEmailApiRequestError{value: val, isSet: true}
}

func (v NullableEmailApiRequestError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailApiRequestError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
