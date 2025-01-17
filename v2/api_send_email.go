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
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"os"
	"reflect"
)

// Linger please
var (
	_ _context.Context
)

// SendEmailApiService SendEmailApi service
type SendEmailApiService service

type ApiGetEmailDeliveryReportsRequest struct {
	ctx        _context.Context
	ApiService *SendEmailApiService
	bulkId     *string
	messageId  *string
	limit      *int32
}

func (r ApiGetEmailDeliveryReportsRequest) BulkId(bulkId string) ApiGetEmailDeliveryReportsRequest {
	r.bulkId = &bulkId
	return r
}
func (r ApiGetEmailDeliveryReportsRequest) MessageId(messageId string) ApiGetEmailDeliveryReportsRequest {
	r.messageId = &messageId
	return r
}
func (r ApiGetEmailDeliveryReportsRequest) Limit(limit int32) ApiGetEmailDeliveryReportsRequest {
	r.limit = &limit
	return r
}

func (r ApiGetEmailDeliveryReportsRequest) Execute() (EmailReportsResult, *_nethttp.Response, error) {
	return r.ApiService.GetEmailDeliveryReportsExecute(r)
}

/*
 * GetEmailDeliveryReports Email delivery reports
 * Get one-time delivery reports for all sent emails.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetEmailDeliveryReportsRequest
 */
func (a *SendEmailApiService) GetEmailDeliveryReports(ctx _context.Context) ApiGetEmailDeliveryReportsRequest {
	return ApiGetEmailDeliveryReportsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return EmailReportsResult
 */
func (a *SendEmailApiService) GetEmailDeliveryReportsExecute(r ApiGetEmailDeliveryReportsRequest) (EmailReportsResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EmailReportsResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SendEmailApiService.GetEmailDeliveryReports")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/email/1/reports"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.bulkId != nil {
		localVarQueryParams.Add("bulkId", parameterToString(*r.bulkId, ""))
	}
	if r.messageId != nil {
		localVarQueryParams.Add("messageId", parameterToString(*r.messageId, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode <= 499 {
			var v EmailApiException
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 && localVarHTTPResponse.StatusCode <= 599 {
			var v EmailApiException
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v EmailReportsResult
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetEmailLogsRequest struct {
	ctx           _context.Context
	ApiService    *SendEmailApiService
	from          *string
	to            *string
	bulkId        *[]string
	messageId     *[]string
	generalStatus *string
	sentSince     *Time
	sentUntil     *Time
	limit         *int32
}

func (r ApiGetEmailLogsRequest) From(from string) ApiGetEmailLogsRequest {
	r.from = &from
	return r
}
func (r ApiGetEmailLogsRequest) To(to string) ApiGetEmailLogsRequest {
	r.to = &to
	return r
}
func (r ApiGetEmailLogsRequest) BulkId(bulkId []string) ApiGetEmailLogsRequest {
	r.bulkId = &bulkId
	return r
}
func (r ApiGetEmailLogsRequest) MessageId(messageId []string) ApiGetEmailLogsRequest {
	r.messageId = &messageId
	return r
}
func (r ApiGetEmailLogsRequest) GeneralStatus(generalStatus string) ApiGetEmailLogsRequest {
	r.generalStatus = &generalStatus
	return r
}
func (r ApiGetEmailLogsRequest) SentSince(sentSince Time) ApiGetEmailLogsRequest {
	r.sentSince = &sentSince
	return r
}
func (r ApiGetEmailLogsRequest) SentUntil(sentUntil Time) ApiGetEmailLogsRequest {
	r.sentUntil = &sentUntil
	return r
}
func (r ApiGetEmailLogsRequest) Limit(limit int32) ApiGetEmailLogsRequest {
	r.limit = &limit
	return r
}

func (r ApiGetEmailLogsRequest) Execute() (EmailLogsResult, *_nethttp.Response, error) {
	return r.ApiService.GetEmailLogsExecute(r)
}

/*
 * GetEmailLogs Email messages logs
 * Get the logs to all email communications, including statuses.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetEmailLogsRequest
 */
func (a *SendEmailApiService) GetEmailLogs(ctx _context.Context) ApiGetEmailLogsRequest {
	return ApiGetEmailLogsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return EmailLogsResult
 */
func (a *SendEmailApiService) GetEmailLogsExecute(r ApiGetEmailLogsRequest) (EmailLogsResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EmailLogsResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SendEmailApiService.GetEmailLogs")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/email/1/logs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.to != nil {
		localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	}
	if r.bulkId != nil {
		t := *r.bulkId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("bulkId", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("bulkId", parameterToString(t, "multi"))
		}
	}
	if r.messageId != nil {
		t := *r.messageId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("messageId", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("messageId", parameterToString(t, "multi"))
		}
	}
	if r.generalStatus != nil {
		localVarQueryParams.Add("generalStatus", parameterToString(*r.generalStatus, ""))
	}
	if r.sentSince != nil {
		localVarQueryParams.Add("sentSince", parameterToString(*r.sentSince, ""))
	}
	if r.sentUntil != nil {
		localVarQueryParams.Add("sentUntil", parameterToString(*r.sentUntil, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode <= 499 {
			var v EmailApiException
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 && localVarHTTPResponse.StatusCode <= 599 {
			var v EmailApiException
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v EmailLogsResult
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSendEmailRequest struct {
	ctx                 _context.Context
	ApiService          *SendEmailApiService
	from                *string
	to                  *string
	subject             *string
	cc                  *string
	bcc                 *string
	text                *string
	bulkId              *string
	messageId           *string
	templateid          *int32
	attachment          **os.File
	inlineImage         **os.File
	hTML                *string
	replyto             *string
	defaultplaceholders *string
	preserverecipients  *bool
	trackingUrl         *string
	trackclicks         *bool
	trackopens          *bool
	track               *bool
	callbackData        *string
	intermediateReport  *bool
	notifyUrl           *string
	notifyContentType   *string
}

func (r ApiSendEmailRequest) From(from string) ApiSendEmailRequest {
	r.from = &from
	return r
}
func (r ApiSendEmailRequest) To(to string) ApiSendEmailRequest {
	r.to = &to
	return r
}
func (r ApiSendEmailRequest) Subject(subject string) ApiSendEmailRequest {
	r.subject = &subject
	return r
}
func (r ApiSendEmailRequest) Cc(cc string) ApiSendEmailRequest {
	r.cc = &cc
	return r
}
func (r ApiSendEmailRequest) Bcc(bcc string) ApiSendEmailRequest {
	r.bcc = &bcc
	return r
}
func (r ApiSendEmailRequest) Text(text string) ApiSendEmailRequest {
	r.text = &text
	return r
}
func (r ApiSendEmailRequest) BulkId(bulkId string) ApiSendEmailRequest {
	r.bulkId = &bulkId
	return r
}
func (r ApiSendEmailRequest) MessageId(messageId string) ApiSendEmailRequest {
	r.messageId = &messageId
	return r
}
func (r ApiSendEmailRequest) Templateid(templateid int32) ApiSendEmailRequest {
	r.templateid = &templateid
	return r
}
func (r ApiSendEmailRequest) Attachment(attachment *os.File) ApiSendEmailRequest {
	r.attachment = &attachment
	return r
}
func (r ApiSendEmailRequest) InlineImage(inlineImage *os.File) ApiSendEmailRequest {
	r.inlineImage = &inlineImage
	return r
}
func (r ApiSendEmailRequest) HTML(hTML string) ApiSendEmailRequest {
	r.hTML = &hTML
	return r
}
func (r ApiSendEmailRequest) Replyto(replyto string) ApiSendEmailRequest {
	r.replyto = &replyto
	return r
}
func (r ApiSendEmailRequest) Defaultplaceholders(defaultplaceholders string) ApiSendEmailRequest {
	r.defaultplaceholders = &defaultplaceholders
	return r
}
func (r ApiSendEmailRequest) Preserverecipients(preserverecipients bool) ApiSendEmailRequest {
	r.preserverecipients = &preserverecipients
	return r
}
func (r ApiSendEmailRequest) TrackingUrl(trackingUrl string) ApiSendEmailRequest {
	r.trackingUrl = &trackingUrl
	return r
}
func (r ApiSendEmailRequest) Trackclicks(trackclicks bool) ApiSendEmailRequest {
	r.trackclicks = &trackclicks
	return r
}
func (r ApiSendEmailRequest) Trackopens(trackopens bool) ApiSendEmailRequest {
	r.trackopens = &trackopens
	return r
}
func (r ApiSendEmailRequest) Track(track bool) ApiSendEmailRequest {
	r.track = &track
	return r
}
func (r ApiSendEmailRequest) CallbackData(callbackData string) ApiSendEmailRequest {
	r.callbackData = &callbackData
	return r
}
func (r ApiSendEmailRequest) IntermediateReport(intermediateReport bool) ApiSendEmailRequest {
	r.intermediateReport = &intermediateReport
	return r
}
func (r ApiSendEmailRequest) NotifyUrl(notifyUrl string) ApiSendEmailRequest {
	r.notifyUrl = &notifyUrl
	return r
}
func (r ApiSendEmailRequest) NotifyContentType(notifyContentType string) ApiSendEmailRequest {
	r.notifyContentType = &notifyContentType
	return r
}

func (r ApiSendEmailRequest) Execute() (EmailSendResponse, *_nethttp.Response, error) {
	return r.ApiService.SendEmailExecute(r)
}

/*
 * SendEmail Fully featured email
 * Send an email or multiple emails to a recipient or multiple recipients with CC/BCC enabled.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiSendEmailRequest
 */
func (a *SendEmailApiService) SendEmail(ctx _context.Context) ApiSendEmailRequest {
	return ApiSendEmailRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

/*
 * Execute executes the request
 * @return EmailSendResponse
 */
func (a *SendEmailApiService) SendEmailExecute(r ApiSendEmailRequest) (EmailSendResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EmailSendResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SendEmailApiService.SendEmail")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/email/2/send"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.from == nil {
		return localVarReturnValue, nil, reportError("from is required and must be specified")
	}
	if r.to == nil {
		return localVarReturnValue, nil, reportError("to is required and must be specified")
	}
	if r.subject == nil {
		return localVarReturnValue, nil, reportError("subject is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarFormParams.Add("from", parameterToString(*r.from, ""))
	localVarFormParams.Add("to", parameterToString(*r.to, ""))
	if r.cc != nil {
		localVarFormParams.Add("cc", parameterToString(*r.cc, ""))
	}
	if r.bcc != nil {
		localVarFormParams.Add("bcc", parameterToString(*r.bcc, ""))
	}
	localVarFormParams.Add("subject", parameterToString(*r.subject, ""))
	if r.text != nil {
		localVarFormParams.Add("text", parameterToString(*r.text, ""))
	}
	if r.bulkId != nil {
		localVarFormParams.Add("bulkId", parameterToString(*r.bulkId, ""))
	}
	if r.messageId != nil {
		localVarFormParams.Add("messageId", parameterToString(*r.messageId, ""))
	}
	if r.templateid != nil {
		localVarFormParams.Add("templateid", parameterToString(*r.templateid, ""))
	}
	localVarFormFileName = "attachment"
	var attachment *os.File
	if r.attachment != nil {
		attachment = *r.attachment
	}
	if attachment != nil {
		fbs, _ := _ioutil.ReadAll(attachment)
		localVarFileBytes = fbs
		localVarFileName = attachment.Name()
		attachment.Close()
	}
	localVarFormFileName = "inlineImage"
	var inlineImage *os.File
	if r.inlineImage != nil {
		inlineImage = *r.inlineImage
	}
	if inlineImage != nil {
		fbs, _ := _ioutil.ReadAll(inlineImage)
		localVarFileBytes = fbs
		localVarFileName = inlineImage.Name()
		inlineImage.Close()
	}
	if r.hTML != nil {
		localVarFormParams.Add("HTML", parameterToString(*r.hTML, ""))
	}
	if r.replyto != nil {
		localVarFormParams.Add("replyto", parameterToString(*r.replyto, ""))
	}
	if r.defaultplaceholders != nil {
		localVarFormParams.Add("defaultplaceholders", parameterToString(*r.defaultplaceholders, ""))
	}
	if r.preserverecipients != nil {
		localVarFormParams.Add("preserverecipients", parameterToString(*r.preserverecipients, ""))
	}
	if r.trackingUrl != nil {
		localVarFormParams.Add("trackingUrl", parameterToString(*r.trackingUrl, ""))
	}
	if r.trackclicks != nil {
		localVarFormParams.Add("trackclicks", parameterToString(*r.trackclicks, ""))
	}
	if r.trackopens != nil {
		localVarFormParams.Add("trackopens", parameterToString(*r.trackopens, ""))
	}
	if r.track != nil {
		localVarFormParams.Add("track", parameterToString(*r.track, ""))
	}
	if r.callbackData != nil {
		localVarFormParams.Add("callbackData", parameterToString(*r.callbackData, ""))
	}
	if r.intermediateReport != nil {
		localVarFormParams.Add("intermediateReport", parameterToString(*r.intermediateReport, ""))
	}
	if r.notifyUrl != nil {
		localVarFormParams.Add("notifyUrl", parameterToString(*r.notifyUrl, ""))
	}
	if r.notifyContentType != nil {
		localVarFormParams.Add("notifyContentType", parameterToString(*r.notifyContentType, ""))
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode <= 499 {
			var v EmailApiException
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 && localVarHTTPResponse.StatusCode <= 599 {
			var v EmailApiException
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v EmailSendResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
