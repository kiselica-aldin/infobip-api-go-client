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

// TfaCreateMessageRequest struct for TfaCreateMessageRequest
type TfaCreateMessageRequest struct {
	// Language code of language in which message text is written. It is used for reading the message when it is sent via voice. If no language is set, message will be read in `English`.
	Language *TfaLanguage `json:"language,omitempty"`
	// Text of a message that will be sent. It can contain placeholders that will be replaced upon sending. Placeholder format is `{{placeholderName}}`. Message text must contain `{{pin}}` placeholder.
	MessageText string `json:"messageText"`
	// PIN code length.
	PinLength *int32 `json:"pinLength,omitempty"`
	// Type of PIN code that will be generated and sent as part of 2FA message.
	PinType TfaPinType `json:"pinType"`
	// Region specific parameters, often specified by local laws. Use this if country or region that you are sending SMS to requires some extra parameters.
	Regional *TfaRegionalOptions `json:"regional,omitempty"`
	// In case PIN message is sent by Voice, DTMF code will enable replaying the message.
	RepeatDTMF *string `json:"repeatDTMF,omitempty"`
	// The name that will appear as the sender of the 2FA message (Example: CompanyName).
	SenderId *string `json:"senderId,omitempty"`
	// In case PIN message is sent by Voice, the speed of speech can be set for the message. Supported range is from `0.5` to `2`.
	SpeechRate *float64 `json:"speechRate,omitempty"`
}

// NewTfaCreateMessageRequest instantiates a new TfaCreateMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTfaCreateMessageRequest(messageText string, pinType TfaPinType) *TfaCreateMessageRequest {
	this := TfaCreateMessageRequest{}
	this.MessageText = messageText
	this.PinType = pinType
	return &this
}

// NewTfaCreateMessageRequestWithDefaults instantiates a new TfaCreateMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTfaCreateMessageRequestWithDefaults() *TfaCreateMessageRequest {
	this := TfaCreateMessageRequest{}
	return &this
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *TfaCreateMessageRequest) GetLanguage() TfaLanguage {
	if o == nil || o.Language == nil {
		var ret TfaLanguage
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaCreateMessageRequest) GetLanguageOk() (*TfaLanguage, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *TfaCreateMessageRequest) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given TfaLanguage and assigns it to the Language field.
func (o *TfaCreateMessageRequest) SetLanguage(v TfaLanguage) {
	o.Language = &v
}

// GetMessageText returns the MessageText field value
func (o *TfaCreateMessageRequest) GetMessageText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageText
}

// GetMessageTextOk returns a tuple with the MessageText field value
// and a boolean to check if the value has been set.
func (o *TfaCreateMessageRequest) GetMessageTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageText, true
}

// SetMessageText sets field value
func (o *TfaCreateMessageRequest) SetMessageText(v string) {
	o.MessageText = v
}

// GetPinLength returns the PinLength field value if set, zero value otherwise.
func (o *TfaCreateMessageRequest) GetPinLength() int32 {
	if o == nil || o.PinLength == nil {
		var ret int32
		return ret
	}
	return *o.PinLength
}

// GetPinLengthOk returns a tuple with the PinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaCreateMessageRequest) GetPinLengthOk() (*int32, bool) {
	if o == nil || o.PinLength == nil {
		return nil, false
	}
	return o.PinLength, true
}

// HasPinLength returns a boolean if a field has been set.
func (o *TfaCreateMessageRequest) HasPinLength() bool {
	if o != nil && o.PinLength != nil {
		return true
	}

	return false
}

// SetPinLength gets a reference to the given int32 and assigns it to the PinLength field.
func (o *TfaCreateMessageRequest) SetPinLength(v int32) {
	o.PinLength = &v
}

// GetPinType returns the PinType field value
func (o *TfaCreateMessageRequest) GetPinType() TfaPinType {
	if o == nil {
		var ret TfaPinType
		return ret
	}

	return o.PinType
}

// GetPinTypeOk returns a tuple with the PinType field value
// and a boolean to check if the value has been set.
func (o *TfaCreateMessageRequest) GetPinTypeOk() (*TfaPinType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PinType, true
}

// SetPinType sets field value
func (o *TfaCreateMessageRequest) SetPinType(v TfaPinType) {
	o.PinType = v
}

// GetRegional returns the Regional field value if set, zero value otherwise.
func (o *TfaCreateMessageRequest) GetRegional() TfaRegionalOptions {
	if o == nil || o.Regional == nil {
		var ret TfaRegionalOptions
		return ret
	}
	return *o.Regional
}

// GetRegionalOk returns a tuple with the Regional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaCreateMessageRequest) GetRegionalOk() (*TfaRegionalOptions, bool) {
	if o == nil || o.Regional == nil {
		return nil, false
	}
	return o.Regional, true
}

// HasRegional returns a boolean if a field has been set.
func (o *TfaCreateMessageRequest) HasRegional() bool {
	if o != nil && o.Regional != nil {
		return true
	}

	return false
}

// SetRegional gets a reference to the given TfaRegionalOptions and assigns it to the Regional field.
func (o *TfaCreateMessageRequest) SetRegional(v TfaRegionalOptions) {
	o.Regional = &v
}

// GetRepeatDTMF returns the RepeatDTMF field value if set, zero value otherwise.
func (o *TfaCreateMessageRequest) GetRepeatDTMF() string {
	if o == nil || o.RepeatDTMF == nil {
		var ret string
		return ret
	}
	return *o.RepeatDTMF
}

// GetRepeatDTMFOk returns a tuple with the RepeatDTMF field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaCreateMessageRequest) GetRepeatDTMFOk() (*string, bool) {
	if o == nil || o.RepeatDTMF == nil {
		return nil, false
	}
	return o.RepeatDTMF, true
}

// HasRepeatDTMF returns a boolean if a field has been set.
func (o *TfaCreateMessageRequest) HasRepeatDTMF() bool {
	if o != nil && o.RepeatDTMF != nil {
		return true
	}

	return false
}

// SetRepeatDTMF gets a reference to the given string and assigns it to the RepeatDTMF field.
func (o *TfaCreateMessageRequest) SetRepeatDTMF(v string) {
	o.RepeatDTMF = &v
}

// GetSenderId returns the SenderId field value if set, zero value otherwise.
func (o *TfaCreateMessageRequest) GetSenderId() string {
	if o == nil || o.SenderId == nil {
		var ret string
		return ret
	}
	return *o.SenderId
}

// GetSenderIdOk returns a tuple with the SenderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaCreateMessageRequest) GetSenderIdOk() (*string, bool) {
	if o == nil || o.SenderId == nil {
		return nil, false
	}
	return o.SenderId, true
}

// HasSenderId returns a boolean if a field has been set.
func (o *TfaCreateMessageRequest) HasSenderId() bool {
	if o != nil && o.SenderId != nil {
		return true
	}

	return false
}

// SetSenderId gets a reference to the given string and assigns it to the SenderId field.
func (o *TfaCreateMessageRequest) SetSenderId(v string) {
	o.SenderId = &v
}

// GetSpeechRate returns the SpeechRate field value if set, zero value otherwise.
func (o *TfaCreateMessageRequest) GetSpeechRate() float64 {
	if o == nil || o.SpeechRate == nil {
		var ret float64
		return ret
	}
	return *o.SpeechRate
}

// GetSpeechRateOk returns a tuple with the SpeechRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TfaCreateMessageRequest) GetSpeechRateOk() (*float64, bool) {
	if o == nil || o.SpeechRate == nil {
		return nil, false
	}
	return o.SpeechRate, true
}

// HasSpeechRate returns a boolean if a field has been set.
func (o *TfaCreateMessageRequest) HasSpeechRate() bool {
	if o != nil && o.SpeechRate != nil {
		return true
	}

	return false
}

// SetSpeechRate gets a reference to the given float64 and assigns it to the SpeechRate field.
func (o *TfaCreateMessageRequest) SetSpeechRate(v float64) {
	o.SpeechRate = &v
}

func (o TfaCreateMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if true {
		toSerialize["messageText"] = o.MessageText
	}
	if o.PinLength != nil {
		toSerialize["pinLength"] = o.PinLength
	}
	if true {
		toSerialize["pinType"] = o.PinType
	}
	if o.Regional != nil {
		toSerialize["regional"] = o.Regional
	}
	if o.RepeatDTMF != nil {
		toSerialize["repeatDTMF"] = o.RepeatDTMF
	}
	if o.SenderId != nil {
		toSerialize["senderId"] = o.SenderId
	}
	if o.SpeechRate != nil {
		toSerialize["speechRate"] = o.SpeechRate
	}
	return json.Marshal(toSerialize)
}

type NullableTfaCreateMessageRequest struct {
	value *TfaCreateMessageRequest
	isSet bool
}

func (v NullableTfaCreateMessageRequest) Get() *TfaCreateMessageRequest {
	return v.value
}

func (v *NullableTfaCreateMessageRequest) Set(val *TfaCreateMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTfaCreateMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTfaCreateMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTfaCreateMessageRequest(val *TfaCreateMessageRequest) *NullableTfaCreateMessageRequest {
	return &NullableTfaCreateMessageRequest{value: val, isSet: true}
}

func (v NullableTfaCreateMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTfaCreateMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}