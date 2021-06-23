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

// WhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf struct for WhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf
type WhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf struct {
	// Latitude of a location sent in the header.
	Latitude *float64 `json:"latitude,omitempty"`
	// Longitude of a location sent in the header.
	Longitude *float64 `json:"longitude,omitempty"`
}

// NewWhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf instantiates a new WhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf() *WhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf {
	this := WhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf{}
	return &this
}

// NewWhatsAppTemplateLocationHeaderWhatsAppApiContentAllOfWithDefaults instantiates a new WhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsAppTemplateLocationHeaderWhatsAppApiContentAllOfWithDefaults() *WhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf {
	this := WhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf{}
	return &this
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *WhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf) GetLatitude() float64 {
	if o == nil || o.Latitude == nil {
		var ret float64
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf) GetLatitudeOk() (*float64, bool) {
	if o == nil || o.Latitude == nil {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *WhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf) HasLatitude() bool {
	if o != nil && o.Latitude != nil {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given float64 and assigns it to the Latitude field.
func (o *WhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf) SetLatitude(v float64) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *WhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf) GetLongitude() float64 {
	if o == nil || o.Longitude == nil {
		var ret float64
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf) GetLongitudeOk() (*float64, bool) {
	if o == nil || o.Longitude == nil {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *WhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf) HasLongitude() bool {
	if o != nil && o.Longitude != nil {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given float64 and assigns it to the Longitude field.
func (o *WhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf) SetLongitude(v float64) {
	o.Longitude = &v
}

func (o WhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Latitude != nil {
		toSerialize["latitude"] = o.Latitude
	}
	if o.Longitude != nil {
		toSerialize["longitude"] = o.Longitude
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf struct {
	value *WhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf
	isSet bool
}

func (v NullableWhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf) Get() *WhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf {
	return v.value
}

func (v *NullableWhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf) Set(val *WhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf(val *WhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf) *NullableWhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf {
	return &NullableWhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf{value: val, isSet: true}
}

func (v NullableWhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsAppTemplateLocationHeaderWhatsAppApiContentAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
