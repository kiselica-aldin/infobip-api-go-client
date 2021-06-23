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

// SmsError struct for SmsError
type SmsError struct {
	// Human-readable description of the error..
	Description *string `json:"description,omitempty"`
	// Error group ID.
	GroupId *int32 `json:"groupId,omitempty"`
	// Error group name.
	GroupName *string `json:"groupName,omitempty"`
	// Error ID.
	Id *int32 `json:"id,omitempty"`
	// Error name.
	Name *string `json:"name,omitempty"`
	// Tells if the error is permanent.
	Permanent *bool `json:"permanent,omitempty"`
}

// NewSmsError instantiates a new SmsError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsError() *SmsError {
	this := SmsError{}
	return &this
}

// NewSmsErrorWithDefaults instantiates a new SmsError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsErrorWithDefaults() *SmsError {
	this := SmsError{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SmsError) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsError) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SmsError) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SmsError) SetDescription(v string) {
	o.Description = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *SmsError) GetGroupId() int32 {
	if o == nil || o.GroupId == nil {
		var ret int32
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsError) GetGroupIdOk() (*int32, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *SmsError) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given int32 and assigns it to the GroupId field.
func (o *SmsError) SetGroupId(v int32) {
	o.GroupId = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *SmsError) GetGroupName() string {
	if o == nil || o.GroupName == nil {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsError) GetGroupNameOk() (*string, bool) {
	if o == nil || o.GroupName == nil {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *SmsError) HasGroupName() bool {
	if o != nil && o.GroupName != nil {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *SmsError) SetGroupName(v string) {
	o.GroupName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SmsError) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsError) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SmsError) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SmsError) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SmsError) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsError) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SmsError) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SmsError) SetName(v string) {
	o.Name = &v
}

// GetPermanent returns the Permanent field value if set, zero value otherwise.
func (o *SmsError) GetPermanent() bool {
	if o == nil || o.Permanent == nil {
		var ret bool
		return ret
	}
	return *o.Permanent
}

// GetPermanentOk returns a tuple with the Permanent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsError) GetPermanentOk() (*bool, bool) {
	if o == nil || o.Permanent == nil {
		return nil, false
	}
	return o.Permanent, true
}

// HasPermanent returns a boolean if a field has been set.
func (o *SmsError) HasPermanent() bool {
	if o != nil && o.Permanent != nil {
		return true
	}

	return false
}

// SetPermanent gets a reference to the given bool and assigns it to the Permanent field.
func (o *SmsError) SetPermanent(v bool) {
	o.Permanent = &v
}

func (o SmsError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	if o.GroupName != nil {
		toSerialize["groupName"] = o.GroupName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Permanent != nil {
		toSerialize["permanent"] = o.Permanent
	}
	return json.Marshal(toSerialize)
}

type NullableSmsError struct {
	value *SmsError
	isSet bool
}

func (v NullableSmsError) Get() *SmsError {
	return v.value
}

func (v *NullableSmsError) Set(val *SmsError) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsError) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsError(val *SmsError) *NullableSmsError {
	return &NullableSmsError{value: val, isSet: true}
}

func (v NullableSmsError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
