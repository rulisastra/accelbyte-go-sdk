// Code generated by go-swagger; DO NOT EDIT.

package concurrent_record

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/cloudsave-sdk/pkg/cloudsaveclientmodels"
)

// PutGameRecordConcurrentHandlerV1Reader is a Reader for the PutGameRecordConcurrentHandlerV1 structure.
type PutGameRecordConcurrentHandlerV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutGameRecordConcurrentHandlerV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutGameRecordConcurrentHandlerV1NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutGameRecordConcurrentHandlerV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 412:
		result := NewPutGameRecordConcurrentHandlerV1PreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPutGameRecordConcurrentHandlerV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /cloudsave/v1/namespaces/{namespace}/concurrent/records/{key} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPutGameRecordConcurrentHandlerV1NoContent creates a PutGameRecordConcurrentHandlerV1NoContent with default headers values
func NewPutGameRecordConcurrentHandlerV1NoContent() *PutGameRecordConcurrentHandlerV1NoContent {
	return &PutGameRecordConcurrentHandlerV1NoContent{}
}

/*PutGameRecordConcurrentHandlerV1NoContent handles this case with default header values.

  Record saved
*/
type PutGameRecordConcurrentHandlerV1NoContent struct {
}

func (o *PutGameRecordConcurrentHandlerV1NoContent) Error() string {
	return fmt.Sprintf("[PUT /cloudsave/v1/namespaces/{namespace}/concurrent/records/{key}][%d] putGameRecordConcurrentHandlerV1NoContent ", 204)
}

func (o *PutGameRecordConcurrentHandlerV1NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutGameRecordConcurrentHandlerV1BadRequest creates a PutGameRecordConcurrentHandlerV1BadRequest with default headers values
func NewPutGameRecordConcurrentHandlerV1BadRequest() *PutGameRecordConcurrentHandlerV1BadRequest {
	return &PutGameRecordConcurrentHandlerV1BadRequest{}
}

/*PutGameRecordConcurrentHandlerV1BadRequest handles this case with default header values.

  Bad Request
*/
type PutGameRecordConcurrentHandlerV1BadRequest struct {
	Payload *cloudsaveclientmodels.ResponseError
}

func (o *PutGameRecordConcurrentHandlerV1BadRequest) Error() string {
	return fmt.Sprintf("[PUT /cloudsave/v1/namespaces/{namespace}/concurrent/records/{key}][%d] putGameRecordConcurrentHandlerV1BadRequest  %+v", 400, o.Payload)
}

func (o *PutGameRecordConcurrentHandlerV1BadRequest) GetPayload() *cloudsaveclientmodels.ResponseError {
	return o.Payload
}

func (o *PutGameRecordConcurrentHandlerV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloudsaveclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGameRecordConcurrentHandlerV1PreconditionFailed creates a PutGameRecordConcurrentHandlerV1PreconditionFailed with default headers values
func NewPutGameRecordConcurrentHandlerV1PreconditionFailed() *PutGameRecordConcurrentHandlerV1PreconditionFailed {
	return &PutGameRecordConcurrentHandlerV1PreconditionFailed{}
}

/*PutGameRecordConcurrentHandlerV1PreconditionFailed handles this case with default header values.

  Precondition Failed
*/
type PutGameRecordConcurrentHandlerV1PreconditionFailed struct {
	Payload *cloudsaveclientmodels.ResponseError
}

func (o *PutGameRecordConcurrentHandlerV1PreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /cloudsave/v1/namespaces/{namespace}/concurrent/records/{key}][%d] putGameRecordConcurrentHandlerV1PreconditionFailed  %+v", 412, o.Payload)
}

func (o *PutGameRecordConcurrentHandlerV1PreconditionFailed) GetPayload() *cloudsaveclientmodels.ResponseError {
	return o.Payload
}

func (o *PutGameRecordConcurrentHandlerV1PreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloudsaveclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGameRecordConcurrentHandlerV1InternalServerError creates a PutGameRecordConcurrentHandlerV1InternalServerError with default headers values
func NewPutGameRecordConcurrentHandlerV1InternalServerError() *PutGameRecordConcurrentHandlerV1InternalServerError {
	return &PutGameRecordConcurrentHandlerV1InternalServerError{}
}

/*PutGameRecordConcurrentHandlerV1InternalServerError handles this case with default header values.

  Internal Server Error
*/
type PutGameRecordConcurrentHandlerV1InternalServerError struct {
	Payload *cloudsaveclientmodels.ResponseError
}

func (o *PutGameRecordConcurrentHandlerV1InternalServerError) Error() string {
	return fmt.Sprintf("[PUT /cloudsave/v1/namespaces/{namespace}/concurrent/records/{key}][%d] putGameRecordConcurrentHandlerV1InternalServerError  %+v", 500, o.Payload)
}

func (o *PutGameRecordConcurrentHandlerV1InternalServerError) GetPayload() *cloudsaveclientmodels.ResponseError {
	return o.Payload
}

func (o *PutGameRecordConcurrentHandlerV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloudsaveclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
