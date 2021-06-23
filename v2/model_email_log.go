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

// EmailLog struct for EmailLog
type EmailLog struct {
	BulkId       *string      `json:"bulkId,omitempty"`
	DoneAt       *Time        `json:"doneAt,omitempty"`
	From         *string      `json:"from,omitempty"`
	MccMnc       *string      `json:"mccMnc,omitempty"`
	MessageCount *int32       `json:"messageCount,omitempty"`
	MessageId    *string      `json:"messageId,omitempty"`
	Price        *EmailPrice  `json:"price,omitempty"`
	SentAt       *Time        `json:"sentAt,omitempty"`
	Status       *EmailStatus `json:"status,omitempty"`
	Text         *string      `json:"text,omitempty"`
	To           *string      `json:"to,omitempty"`
}

// NewEmailLog instantiates a new EmailLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailLog() *EmailLog {
	this := EmailLog{}
	return &this
}

// NewEmailLogWithDefaults instantiates a new EmailLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailLogWithDefaults() *EmailLog {
	this := EmailLog{}
	return &this
}

// GetBulkId returns the BulkId field value if set, zero value otherwise.
func (o *EmailLog) GetBulkId() string {
	if o == nil || o.BulkId == nil {
		var ret string
		return ret
	}
	return *o.BulkId
}

// GetBulkIdOk returns a tuple with the BulkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailLog) GetBulkIdOk() (*string, bool) {
	if o == nil || o.BulkId == nil {
		return nil, false
	}
	return o.BulkId, true
}

// HasBulkId returns a boolean if a field has been set.
func (o *EmailLog) HasBulkId() bool {
	if o != nil && o.BulkId != nil {
		return true
	}

	return false
}

// SetBulkId gets a reference to the given string and assigns it to the BulkId field.
func (o *EmailLog) SetBulkId(v string) {
	o.BulkId = &v
}

// GetDoneAt returns the DoneAt field value if set, zero value otherwise.
func (o *EmailLog) GetDoneAt() Time {
	if o == nil || o.DoneAt == nil {
		var ret Time
		return ret
	}
	return *o.DoneAt
}

// GetDoneAtOk returns a tuple with the DoneAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailLog) GetDoneAtOk() (*Time, bool) {
	if o == nil || o.DoneAt == nil {
		return nil, false
	}
	return o.DoneAt, true
}

// HasDoneAt returns a boolean if a field has been set.
func (o *EmailLog) HasDoneAt() bool {
	if o != nil && o.DoneAt != nil {
		return true
	}

	return false
}

// SetDoneAt gets a reference to the given Time and assigns it to the DoneAt field.
func (o *EmailLog) SetDoneAt(v Time) {
	o.DoneAt = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *EmailLog) GetFrom() string {
	if o == nil || o.From == nil {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailLog) GetFromOk() (*string, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *EmailLog) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *EmailLog) SetFrom(v string) {
	o.From = &v
}

// GetMccMnc returns the MccMnc field value if set, zero value otherwise.
func (o *EmailLog) GetMccMnc() string {
	if o == nil || o.MccMnc == nil {
		var ret string
		return ret
	}
	return *o.MccMnc
}

// GetMccMncOk returns a tuple with the MccMnc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailLog) GetMccMncOk() (*string, bool) {
	if o == nil || o.MccMnc == nil {
		return nil, false
	}
	return o.MccMnc, true
}

// HasMccMnc returns a boolean if a field has been set.
func (o *EmailLog) HasMccMnc() bool {
	if o != nil && o.MccMnc != nil {
		return true
	}

	return false
}

// SetMccMnc gets a reference to the given string and assigns it to the MccMnc field.
func (o *EmailLog) SetMccMnc(v string) {
	o.MccMnc = &v
}

// GetMessageCount returns the MessageCount field value if set, zero value otherwise.
func (o *EmailLog) GetMessageCount() int32 {
	if o == nil || o.MessageCount == nil {
		var ret int32
		return ret
	}
	return *o.MessageCount
}

// GetMessageCountOk returns a tuple with the MessageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailLog) GetMessageCountOk() (*int32, bool) {
	if o == nil || o.MessageCount == nil {
		return nil, false
	}
	return o.MessageCount, true
}

// HasMessageCount returns a boolean if a field has been set.
func (o *EmailLog) HasMessageCount() bool {
	if o != nil && o.MessageCount != nil {
		return true
	}

	return false
}

// SetMessageCount gets a reference to the given int32 and assigns it to the MessageCount field.
func (o *EmailLog) SetMessageCount(v int32) {
	o.MessageCount = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *EmailLog) GetMessageId() string {
	if o == nil || o.MessageId == nil {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailLog) GetMessageIdOk() (*string, bool) {
	if o == nil || o.MessageId == nil {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *EmailLog) HasMessageId() bool {
	if o != nil && o.MessageId != nil {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *EmailLog) SetMessageId(v string) {
	o.MessageId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *EmailLog) GetPrice() EmailPrice {
	if o == nil || o.Price == nil {
		var ret EmailPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailLog) GetPriceOk() (*EmailPrice, bool) {
	if o == nil || o.Price == nil {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *EmailLog) HasPrice() bool {
	if o != nil && o.Price != nil {
		return true
	}

	return false
}

// SetPrice gets a reference to the given EmailPrice and assigns it to the Price field.
func (o *EmailLog) SetPrice(v EmailPrice) {
	o.Price = &v
}

// GetSentAt returns the SentAt field value if set, zero value otherwise.
func (o *EmailLog) GetSentAt() Time {
	if o == nil || o.SentAt == nil {
		var ret Time
		return ret
	}
	return *o.SentAt
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailLog) GetSentAtOk() (*Time, bool) {
	if o == nil || o.SentAt == nil {
		return nil, false
	}
	return o.SentAt, true
}

// HasSentAt returns a boolean if a field has been set.
func (o *EmailLog) HasSentAt() bool {
	if o != nil && o.SentAt != nil {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given Time and assigns it to the SentAt field.
func (o *EmailLog) SetSentAt(v Time) {
	o.SentAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EmailLog) GetStatus() EmailStatus {
	if o == nil || o.Status == nil {
		var ret EmailStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailLog) GetStatusOk() (*EmailStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EmailLog) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given EmailStatus and assigns it to the Status field.
func (o *EmailLog) SetStatus(v EmailStatus) {
	o.Status = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *EmailLog) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailLog) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *EmailLog) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *EmailLog) SetText(v string) {
	o.Text = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *EmailLog) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailLog) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *EmailLog) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *EmailLog) SetTo(v string) {
	o.To = &v
}

func (o EmailLog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BulkId != nil {
		toSerialize["bulkId"] = o.BulkId
	}
	if o.DoneAt != nil {
		toSerialize["doneAt"] = o.DoneAt
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.MccMnc != nil {
		toSerialize["mccMnc"] = o.MccMnc
	}
	if o.MessageCount != nil {
		toSerialize["messageCount"] = o.MessageCount
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

type NullableEmailLog struct {
	value *EmailLog
	isSet bool
}

func (v NullableEmailLog) Get() *EmailLog {
	return v.value
}

func (v *NullableEmailLog) Set(val *EmailLog) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailLog) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailLog(val *EmailLog) *NullableEmailLog {
	return &NullableEmailLog{value: val, isSet: true}
}

func (v NullableEmailLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
