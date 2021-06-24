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

// WhatsAppTemplateDataWhatsAppApiContent Template data. Values have to be set as registered in the template.
type WhatsAppTemplateDataWhatsAppApiContent struct {
	Body   WhatsAppTemplateBodyWhatsAppApiContent    `json:"body"`
	Header *WhatsAppTemplateHeaderWhatsAppApiContent `json:"header,omitempty"`
	// Template buttons. Should be defined in correct order, only if `quick reply` or `dynamic URL` buttons have been registered. It can have up to three `quick reply` buttons or only one `dynamic URL` button.
	Buttons *[]WhatsAppTemplateButtonWhatsAppApiContent `json:"buttons,omitempty"`
}

// NewWhatsAppTemplateDataWhatsAppApiContent instantiates a new WhatsAppTemplateDataWhatsAppApiContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWhatsAppTemplateDataWhatsAppApiContent(body WhatsAppTemplateBodyWhatsAppApiContent) *WhatsAppTemplateDataWhatsAppApiContent {
	this := WhatsAppTemplateDataWhatsAppApiContent{}
	this.Body = body
	return &this
}

// NewWhatsAppTemplateDataWhatsAppApiContentWithDefaults instantiates a new WhatsAppTemplateDataWhatsAppApiContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWhatsAppTemplateDataWhatsAppApiContentWithDefaults() *WhatsAppTemplateDataWhatsAppApiContent {
	this := WhatsAppTemplateDataWhatsAppApiContent{}
	return &this
}

// GetBody returns the Body field value
func (o *WhatsAppTemplateDataWhatsAppApiContent) GetBody() WhatsAppTemplateBodyWhatsAppApiContent {
	if o == nil {
		var ret WhatsAppTemplateBodyWhatsAppApiContent
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *WhatsAppTemplateDataWhatsAppApiContent) GetBodyOk() (*WhatsAppTemplateBodyWhatsAppApiContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *WhatsAppTemplateDataWhatsAppApiContent) SetBody(v WhatsAppTemplateBodyWhatsAppApiContent) {
	o.Body = v
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *WhatsAppTemplateDataWhatsAppApiContent) GetHeader() WhatsAppTemplateHeaderWhatsAppApiContent {
	if o == nil || o.Header == nil {
		var ret WhatsAppTemplateHeaderWhatsAppApiContent
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsAppTemplateDataWhatsAppApiContent) GetHeaderOk() (*WhatsAppTemplateHeaderWhatsAppApiContent, bool) {
	if o == nil || o.Header == nil {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *WhatsAppTemplateDataWhatsAppApiContent) HasHeader() bool {
	if o != nil && o.Header != nil {
		return true
	}

	return false
}

// SetHeader gets a reference to the given WhatsAppTemplateHeaderWhatsAppApiContent and assigns it to the Header field.
func (o *WhatsAppTemplateDataWhatsAppApiContent) SetHeader(v WhatsAppTemplateHeaderWhatsAppApiContent) {
	o.Header = &v
}

// GetButtons returns the Buttons field value if set, zero value otherwise.
func (o *WhatsAppTemplateDataWhatsAppApiContent) GetButtons() []WhatsAppTemplateButtonWhatsAppApiContent {
	if o == nil || o.Buttons == nil {
		var ret []WhatsAppTemplateButtonWhatsAppApiContent
		return ret
	}
	return *o.Buttons
}

// GetButtonsOk returns a tuple with the Buttons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WhatsAppTemplateDataWhatsAppApiContent) GetButtonsOk() (*[]WhatsAppTemplateButtonWhatsAppApiContent, bool) {
	if o == nil || o.Buttons == nil {
		return nil, false
	}
	return o.Buttons, true
}

// HasButtons returns a boolean if a field has been set.
func (o *WhatsAppTemplateDataWhatsAppApiContent) HasButtons() bool {
	if o != nil && o.Buttons != nil {
		return true
	}

	return false
}

// SetButtons gets a reference to the given []WhatsAppTemplateButtonWhatsAppApiContent and assigns it to the Buttons field.
func (o *WhatsAppTemplateDataWhatsAppApiContent) SetButtons(v []WhatsAppTemplateButtonWhatsAppApiContent) {
	o.Buttons = &v
}

func (o WhatsAppTemplateDataWhatsAppApiContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["body"] = o.Body
	}
	if o.Header != nil {
		toSerialize["header"] = o.Header
	}
	if o.Buttons != nil {
		toSerialize["buttons"] = o.Buttons
	}
	return json.Marshal(toSerialize)
}

type NullableWhatsAppTemplateDataWhatsAppApiContent struct {
	value *WhatsAppTemplateDataWhatsAppApiContent
	isSet bool
}

func (v NullableWhatsAppTemplateDataWhatsAppApiContent) Get() *WhatsAppTemplateDataWhatsAppApiContent {
	return v.value
}

func (v *NullableWhatsAppTemplateDataWhatsAppApiContent) Set(val *WhatsAppTemplateDataWhatsAppApiContent) {
	v.value = val
	v.isSet = true
}

func (v NullableWhatsAppTemplateDataWhatsAppApiContent) IsSet() bool {
	return v.isSet
}

func (v *NullableWhatsAppTemplateDataWhatsAppApiContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWhatsAppTemplateDataWhatsAppApiContent(val *WhatsAppTemplateDataWhatsAppApiContent) *NullableWhatsAppTemplateDataWhatsAppApiContent {
	return &NullableWhatsAppTemplateDataWhatsAppApiContent{value: val, isSet: true}
}

func (v NullableWhatsAppTemplateDataWhatsAppApiContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWhatsAppTemplateDataWhatsAppApiContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}