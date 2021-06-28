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

// SmsLog struct for SmsLog
type SmsLog struct {
	// The ID that uniquely identifies the request.
	BulkId *string `json:"bulkId,omitempty"`
	// Tells when the SMS was finished processing by Infobip (i.e. delivered to the destination, delivered to the destination network, etc.). Has the following format: `yyyy-MM-dd'T'HH:mm:ss.SSSZ`.
	DoneAt *Time     `json:"doneAt,omitempty"`
	Error  *SmsError `json:"error,omitempty"`
	// Sender ID that can be alphanumeric or numeric.
	From *string `json:"from,omitempty"`
	// Mobile country and network codes.
	MccMnc *string `json:"mccMnc,omitempty"`
	// The ID that uniquely identifies the message sent.
	MessageId *string   `json:"messageId,omitempty"`
	Price     *SmsPrice `json:"price,omitempty"`
	// Tells when the SMS was sent. Has the following format: `yyyy-MM-dd'T'HH:mm:ss.SSSZ`.
	SentAt *Time `json:"sentAt,omitempty"`
	// The number of sent message segments.
	SmsCount *int32     `json:"smsCount,omitempty"`
	Status   *SmsStatus `json:"status,omitempty"`
	// Text of the message that was sent.
	Text *string `json:"text,omitempty"`
	// The message destination address.
	To *string `json:"to,omitempty"`
}

// NewSmsLog instantiates a new SmsLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsLog() *SmsLog {
	this := SmsLog{}
	return &this
}

// NewSmsLogWithDefaults instantiates a new SmsLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsLogWithDefaults() *SmsLog {
	this := SmsLog{}
	return &this
}

// GetBulkId returns the BulkId field value if set, zero value otherwise.
func (o *SmsLog) GetBulkId() string {
	if o == nil || o.BulkId == nil {
		var ret string
		return ret
	}
	return *o.BulkId
}

// GetBulkIdOk returns a tuple with the BulkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsLog) GetBulkIdOk() (*string, bool) {
	if o == nil || o.BulkId == nil {
		return nil, false
	}
	return o.BulkId, true
}

// HasBulkId returns a boolean if a field has been set.
func (o *SmsLog) HasBulkId() bool {
	if o != nil && o.BulkId != nil {
		return true
	}

	return false
}

// SetBulkId gets a reference to the given string and assigns it to the BulkId field.
func (o *SmsLog) SetBulkId(v string) {
	o.BulkId = &v
}

// GetDoneAt returns the DoneAt field value if set, zero value otherwise.
func (o *SmsLog) GetDoneAt() Time {
	if o == nil || o.DoneAt == nil {
		var ret Time
		return ret
	}
	return *o.DoneAt
}

// GetDoneAtOk returns a tuple with the DoneAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsLog) GetDoneAtOk() (*Time, bool) {
	if o == nil || o.DoneAt == nil {
		return nil, false
	}
	return o.DoneAt, true
}

// HasDoneAt returns a boolean if a field has been set.
func (o *SmsLog) HasDoneAt() bool {
	if o != nil && o.DoneAt != nil {
		return true
	}

	return false
}

// SetDoneAt gets a reference to the given Time and assigns it to the DoneAt field.
func (o *SmsLog) SetDoneAt(v Time) {
	o.DoneAt = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SmsLog) GetError() SmsError {
	if o == nil || o.Error == nil {
		var ret SmsError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsLog) GetErrorOk() (*SmsError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SmsLog) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given SmsError and assigns it to the Error field.
func (o *SmsLog) SetError(v SmsError) {
	o.Error = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *SmsLog) GetFrom() string {
	if o == nil || o.From == nil {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsLog) GetFromOk() (*string, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *SmsLog) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *SmsLog) SetFrom(v string) {
	o.From = &v
}

// GetMccMnc returns the MccMnc field value if set, zero value otherwise.
func (o *SmsLog) GetMccMnc() string {
	if o == nil || o.MccMnc == nil {
		var ret string
		return ret
	}
	return *o.MccMnc
}

// GetMccMncOk returns a tuple with the MccMnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsLog) GetMccMncOk() (*string, bool) {
	if o == nil || o.MccMnc == nil {
		return nil, false
	}
	return o.MccMnc, true
}

// HasMccMnc returns a boolean if a field has been set.
func (o *SmsLog) HasMccMnc() bool {
	if o != nil && o.MccMnc != nil {
		return true
	}

	return false
}

// SetMccMnc gets a reference to the given string and assigns it to the MccMnc field.
func (o *SmsLog) SetMccMnc(v string) {
	o.MccMnc = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *SmsLog) GetMessageId() string {
	if o == nil || o.MessageId == nil {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsLog) GetMessageIdOk() (*string, bool) {
	if o == nil || o.MessageId == nil {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *SmsLog) HasMessageId() bool {
	if o != nil && o.MessageId != nil {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *SmsLog) SetMessageId(v string) {
	o.MessageId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SmsLog) GetPrice() SmsPrice {
	if o == nil || o.Price == nil {
		var ret SmsPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsLog) GetPriceOk() (*SmsPrice, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SmsLog) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given SmsPrice and assigns it to the Price field.
func (o *SmsLog) SetPrice(v SmsPrice) {
	o.Price = &v
}

// GetSentAt returns the SentAt field value if set, zero value otherwise.
func (o *SmsLog) GetSentAt() Time {
	if o == nil || o.SentAt == nil {
		var ret Time
		return ret
	}
	return *o.SentAt
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsLog) GetSentAtOk() (*Time, bool) {
	if o == nil || o.SentAt == nil {
		return nil, false
	}
	return o.SentAt, true
}

// HasSentAt returns a boolean if a field has been set.
func (o *SmsLog) HasSentAt() bool {
	if o != nil && o.SentAt != nil {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given Time and assigns it to the SentAt field.
func (o *SmsLog) SetSentAt(v Time) {
	o.SentAt = &v
}

// GetSmsCount returns the SmsCount field value if set, zero value otherwise.
func (o *SmsLog) GetSmsCount() int32 {
	if o == nil || o.SmsCount == nil {
		var ret int32
		return ret
	}
	return *o.SmsCount
}

// GetSmsCountOk returns a tuple with the SmsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsLog) GetSmsCountOk() (*int32, bool) {
	if o == nil || o.SmsCount == nil {
		return nil, false
	}
	return o.SmsCount, true
}

// HasSmsCount returns a boolean if a field has been set.
func (o *SmsLog) HasSmsCount() bool {
	if o != nil && o.SmsCount != nil {
		return true
	}

	return false
}

// SetSmsCount gets a reference to the given int32 and assigns it to the SmsCount field.
func (o *SmsLog) SetSmsCount(v int32) {
	o.SmsCount = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SmsLog) GetStatus() SmsStatus {
	if o == nil || o.Status == nil {
		var ret SmsStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsLog) GetStatusOk() (*SmsStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SmsLog) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given SmsStatus and assigns it to the Status field.
func (o *SmsLog) SetStatus(v SmsStatus) {
	o.Status = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *SmsLog) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsLog) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *SmsLog) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *SmsLog) SetText(v string) {
	o.Text = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *SmsLog) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsLog) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *SmsLog) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *SmsLog) SetTo(v string) {
	o.To = &v
}

func (o SmsLog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BulkId != nil {
		toSerialize["bulkId"] = o.BulkId
	}
	if o.DoneAt != nil {
		toSerialize["doneAt"] = o.DoneAt
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.MccMnc != nil {
		toSerialize["mccMnc"] = o.MccMnc
	}
	if o.MessageId != nil {
		toSerialize["messageId"] = o.MessageId
	}
	if o.Price != nil {
		toSerialize["price"] = o.Price
	}
	if o.SentAt != nil {
		toSerialize["sentAt"] = o.SentAt
	}
	if o.SmsCount != nil {
		toSerialize["smsCount"] = o.SmsCount
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableSmsLog struct {
	value *SmsLog
	isSet bool
}

func (v NullableSmsLog) Get() *SmsLog {
	return v.value
}

func (v *NullableSmsLog) Set(val *SmsLog) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsLog) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsLog(val *SmsLog) *NullableSmsLog {
	return &NullableSmsLog{value: val, isSet: true}
}

func (v NullableSmsLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
