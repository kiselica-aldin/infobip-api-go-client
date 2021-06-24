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

// SmsTracking struct for SmsTracking
type SmsTracking struct {
	// Custom base url used for shortening links from SMS text in `URL` Conversion rate tracking use-case.
	BaseUrl *string `json:"baseUrl,omitempty"`
	// Key that uniquely identifies Conversion tracking process.
	ProcessKey *string `json:"processKey,omitempty"`
	// Indicates if the message has to be tracked for Conversion rates. Possible values: `SMS` and `URL`
	Track *string `json:"track,omitempty"`
	// User-defined type of the Conversion tracking process or flow type or message type, etc. Example: `ONE_TIME_PIN or SOCIAL_INVITES`.
	Type *string `json:"type,omitempty"`
}

// NewSmsTracking instantiates a new SmsTracking object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsTracking() *SmsTracking {
	this := SmsTracking{}
	return &this
}

// NewSmsTrackingWithDefaults instantiates a new SmsTracking object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsTrackingWithDefaults() *SmsTracking {
	this := SmsTracking{}
	return &this
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *SmsTracking) GetBaseUrl() string {
	if o == nil || o.BaseUrl == nil {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsTracking) GetBaseUrlOk() (*string, bool) {
	if o == nil || o.BaseUrl == nil {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *SmsTracking) HasBaseUrl() bool {
	if o != nil && o.BaseUrl != nil {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *SmsTracking) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

// GetProcessKey returns the ProcessKey field value if set, zero value otherwise.
func (o *SmsTracking) GetProcessKey() string {
	if o == nil || o.ProcessKey == nil {
		var ret string
		return ret
	}
	return *o.ProcessKey
}

// GetProcessKeyOk returns a tuple with the ProcessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsTracking) GetProcessKeyOk() (*string, bool) {
	if o == nil || o.ProcessKey == nil {
		return nil, false
	}
	return o.ProcessKey, true
}

// HasProcessKey returns a boolean if a field has been set.
func (o *SmsTracking) HasProcessKey() bool {
	if o != nil && o.ProcessKey != nil {
		return true
	}

	return false
}

// SetProcessKey gets a reference to the given string and assigns it to the ProcessKey field.
func (o *SmsTracking) SetProcessKey(v string) {
	o.ProcessKey = &v
}

// GetTrack returns the Track field value if set, zero value otherwise.
func (o *SmsTracking) GetTrack() string {
	if o == nil || o.Track == nil {
		var ret string
		return ret
	}
	return *o.Track
}

// GetTrackOk returns a tuple with the Track field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsTracking) GetTrackOk() (*string, bool) {
	if o == nil || o.Track == nil {
		return nil, false
	}
	return o.Track, true
}

// HasTrack returns a boolean if a field has been set.
func (o *SmsTracking) HasTrack() bool {
	if o != nil && o.Track != nil {
		return true
	}

	return false
}

// SetTrack gets a reference to the given string and assigns it to the Track field.
func (o *SmsTracking) SetTrack(v string) {
	o.Track = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SmsTracking) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsTracking) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SmsTracking) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SmsTracking) SetType(v string) {
	o.Type = &v
}

func (o SmsTracking) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BaseUrl != nil {
		toSerialize["baseUrl"] = o.BaseUrl
	}
	if o.ProcessKey != nil {
		toSerialize["processKey"] = o.ProcessKey
	}
	if o.Track != nil {
		toSerialize["track"] = o.Track
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSmsTracking struct {
	value *SmsTracking
	isSet bool
}

func (v NullableSmsTracking) Get() *SmsTracking {
	return v.value
}

func (v *NullableSmsTracking) Set(val *SmsTracking) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsTracking) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsTracking) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsTracking(val *SmsTracking) *NullableSmsTracking {
	return &NullableSmsTracking{value: val, isSet: true}
}

func (v NullableSmsTracking) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsTracking) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}