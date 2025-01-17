// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclientmodels"
)

// DeleteLocalServerReader is a Reader for the DeleteLocalServer structure.
type DeleteLocalServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLocalServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLocalServerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteLocalServerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewDeleteLocalServerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /dsmcontroller/admin/namespaces/{namespace}/servers/local/{name} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewDeleteLocalServerNoContent creates a DeleteLocalServerNoContent with default headers values
func NewDeleteLocalServerNoContent() *DeleteLocalServerNoContent {
	return &DeleteLocalServerNoContent{}
}

/*DeleteLocalServerNoContent handles this case with default header values.

  server deleted
*/
type DeleteLocalServerNoContent struct {
}

func (o *DeleteLocalServerNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dsmcontroller/admin/namespaces/{namespace}/servers/local/{name}][%d] deleteLocalServerNoContent ", 204)
}

func (o *DeleteLocalServerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLocalServerUnauthorized creates a DeleteLocalServerUnauthorized with default headers values
func NewDeleteLocalServerUnauthorized() *DeleteLocalServerUnauthorized {
	return &DeleteLocalServerUnauthorized{}
}

/*DeleteLocalServerUnauthorized handles this case with default header values.

  Unauthorized
*/
type DeleteLocalServerUnauthorized struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *DeleteLocalServerUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /dsmcontroller/admin/namespaces/{namespace}/servers/local/{name}][%d] deleteLocalServerUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteLocalServerUnauthorized) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *DeleteLocalServerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLocalServerInternalServerError creates a DeleteLocalServerInternalServerError with default headers values
func NewDeleteLocalServerInternalServerError() *DeleteLocalServerInternalServerError {
	return &DeleteLocalServerInternalServerError{}
}

/*DeleteLocalServerInternalServerError handles this case with default header values.

  Internal Server Error
*/
type DeleteLocalServerInternalServerError struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *DeleteLocalServerInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /dsmcontroller/admin/namespaces/{namespace}/servers/local/{name}][%d] deleteLocalServerInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteLocalServerInternalServerError) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *DeleteLocalServerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
