// Code generated by go-swagger; DO NOT EDIT.

package notification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/lobby-sdk/pkg/lobbyclientmodels"
)

// NotificationWithTemplateReader is a Reader for the NotificationWithTemplate structure.
type NotificationWithTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotificationWithTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewNotificationWithTemplateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewNotificationWithTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewNotificationWithTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewNotificationWithTemplateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewNotificationWithTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /notification/namespaces/{namespace}/templated returns an error %d: %s", response.Code(), string(data))
	}
}

// NewNotificationWithTemplateAccepted creates a NotificationWithTemplateAccepted with default headers values
func NewNotificationWithTemplateAccepted() *NotificationWithTemplateAccepted {
	return &NotificationWithTemplateAccepted{}
}

/*NotificationWithTemplateAccepted handles this case with default header values.

  Accepted
*/
type NotificationWithTemplateAccepted struct {
}

func (o *NotificationWithTemplateAccepted) Error() string {
	return fmt.Sprintf("[POST /notification/namespaces/{namespace}/templated][%d] notificationWithTemplateAccepted ", 202)
}

func (o *NotificationWithTemplateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNotificationWithTemplateBadRequest creates a NotificationWithTemplateBadRequest with default headers values
func NewNotificationWithTemplateBadRequest() *NotificationWithTemplateBadRequest {
	return &NotificationWithTemplateBadRequest{}
}

/*NotificationWithTemplateBadRequest handles this case with default header values.

  Bad Request
*/
type NotificationWithTemplateBadRequest struct {
	Payload *lobbyclientmodels.RestapiErrorResponseBody
}

func (o *NotificationWithTemplateBadRequest) Error() string {
	return fmt.Sprintf("[POST /notification/namespaces/{namespace}/templated][%d] notificationWithTemplateBadRequest  %+v", 400, o.Payload)
}

func (o *NotificationWithTemplateBadRequest) GetPayload() *lobbyclientmodels.RestapiErrorResponseBody {
	return o.Payload
}

func (o *NotificationWithTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationWithTemplateUnauthorized creates a NotificationWithTemplateUnauthorized with default headers values
func NewNotificationWithTemplateUnauthorized() *NotificationWithTemplateUnauthorized {
	return &NotificationWithTemplateUnauthorized{}
}

/*NotificationWithTemplateUnauthorized handles this case with default header values.

  Unauthorized
*/
type NotificationWithTemplateUnauthorized struct {
	Payload *lobbyclientmodels.RestapiErrorResponseBody
}

func (o *NotificationWithTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /notification/namespaces/{namespace}/templated][%d] notificationWithTemplateUnauthorized  %+v", 401, o.Payload)
}

func (o *NotificationWithTemplateUnauthorized) GetPayload() *lobbyclientmodels.RestapiErrorResponseBody {
	return o.Payload
}

func (o *NotificationWithTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationWithTemplateForbidden creates a NotificationWithTemplateForbidden with default headers values
func NewNotificationWithTemplateForbidden() *NotificationWithTemplateForbidden {
	return &NotificationWithTemplateForbidden{}
}

/*NotificationWithTemplateForbidden handles this case with default header values.

  Forbidden
*/
type NotificationWithTemplateForbidden struct {
	Payload *lobbyclientmodels.RestapiErrorResponseBody
}

func (o *NotificationWithTemplateForbidden) Error() string {
	return fmt.Sprintf("[POST /notification/namespaces/{namespace}/templated][%d] notificationWithTemplateForbidden  %+v", 403, o.Payload)
}

func (o *NotificationWithTemplateForbidden) GetPayload() *lobbyclientmodels.RestapiErrorResponseBody {
	return o.Payload
}

func (o *NotificationWithTemplateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotificationWithTemplateNotFound creates a NotificationWithTemplateNotFound with default headers values
func NewNotificationWithTemplateNotFound() *NotificationWithTemplateNotFound {
	return &NotificationWithTemplateNotFound{}
}

/*NotificationWithTemplateNotFound handles this case with default header values.

  Not Found
*/
type NotificationWithTemplateNotFound struct {
	Payload *lobbyclientmodels.RestapiErrorResponseBody
}

func (o *NotificationWithTemplateNotFound) Error() string {
	return fmt.Sprintf("[POST /notification/namespaces/{namespace}/templated][%d] notificationWithTemplateNotFound  %+v", 404, o.Payload)
}

func (o *NotificationWithTemplateNotFound) GetPayload() *lobbyclientmodels.RestapiErrorResponseBody {
	return o.Payload
}

func (o *NotificationWithTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(lobbyclientmodels.RestapiErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
