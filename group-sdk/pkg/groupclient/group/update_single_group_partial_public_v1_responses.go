// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/group-sdk/pkg/groupclientmodels"
)

// UpdateSingleGroupPartialPublicV1Reader is a Reader for the UpdateSingleGroupPartialPublicV1 structure.
type UpdateSingleGroupPartialPublicV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSingleGroupPartialPublicV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSingleGroupPartialPublicV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateSingleGroupPartialPublicV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateSingleGroupPartialPublicV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateSingleGroupPartialPublicV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateSingleGroupPartialPublicV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewUpdateSingleGroupPartialPublicV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PATCH /group/v1/public/namespaces/{namespace}/groups/{groupId} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewUpdateSingleGroupPartialPublicV1OK creates a UpdateSingleGroupPartialPublicV1OK with default headers values
func NewUpdateSingleGroupPartialPublicV1OK() *UpdateSingleGroupPartialPublicV1OK {
	return &UpdateSingleGroupPartialPublicV1OK{}
}

/*UpdateSingleGroupPartialPublicV1OK handles this case with default header values.

  OK
*/
type UpdateSingleGroupPartialPublicV1OK struct {
	Payload *groupclientmodels.ModelsGroupResponseV1
}

func (o *UpdateSingleGroupPartialPublicV1OK) Error() string {
	return fmt.Sprintf("[PATCH /group/v1/public/namespaces/{namespace}/groups/{groupId}][%d] updateSingleGroupPartialPublicV1OK  %+v", 200, o.Payload)
}

func (o *UpdateSingleGroupPartialPublicV1OK) GetPayload() *groupclientmodels.ModelsGroupResponseV1 {
	return o.Payload
}

func (o *UpdateSingleGroupPartialPublicV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(groupclientmodels.ModelsGroupResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSingleGroupPartialPublicV1BadRequest creates a UpdateSingleGroupPartialPublicV1BadRequest with default headers values
func NewUpdateSingleGroupPartialPublicV1BadRequest() *UpdateSingleGroupPartialPublicV1BadRequest {
	return &UpdateSingleGroupPartialPublicV1BadRequest{}
}

/*UpdateSingleGroupPartialPublicV1BadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type UpdateSingleGroupPartialPublicV1BadRequest struct {
	Payload *groupclientmodels.ResponseErrorResponse
}

func (o *UpdateSingleGroupPartialPublicV1BadRequest) Error() string {
	return fmt.Sprintf("[PATCH /group/v1/public/namespaces/{namespace}/groups/{groupId}][%d] updateSingleGroupPartialPublicV1BadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSingleGroupPartialPublicV1BadRequest) GetPayload() *groupclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *UpdateSingleGroupPartialPublicV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(groupclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSingleGroupPartialPublicV1Unauthorized creates a UpdateSingleGroupPartialPublicV1Unauthorized with default headers values
func NewUpdateSingleGroupPartialPublicV1Unauthorized() *UpdateSingleGroupPartialPublicV1Unauthorized {
	return &UpdateSingleGroupPartialPublicV1Unauthorized{}
}

/*UpdateSingleGroupPartialPublicV1Unauthorized handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20001</td><td>unauthorized access</td></tr></table>
*/
type UpdateSingleGroupPartialPublicV1Unauthorized struct {
	Payload *groupclientmodels.ResponseErrorResponse
}

func (o *UpdateSingleGroupPartialPublicV1Unauthorized) Error() string {
	return fmt.Sprintf("[PATCH /group/v1/public/namespaces/{namespace}/groups/{groupId}][%d] updateSingleGroupPartialPublicV1Unauthorized  %+v", 401, o.Payload)
}

func (o *UpdateSingleGroupPartialPublicV1Unauthorized) GetPayload() *groupclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *UpdateSingleGroupPartialPublicV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(groupclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSingleGroupPartialPublicV1Forbidden creates a UpdateSingleGroupPartialPublicV1Forbidden with default headers values
func NewUpdateSingleGroupPartialPublicV1Forbidden() *UpdateSingleGroupPartialPublicV1Forbidden {
	return &UpdateSingleGroupPartialPublicV1Forbidden{}
}

/*UpdateSingleGroupPartialPublicV1Forbidden handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20013</td><td>insufficient permissions</td></tr><tr><td>20022</td><td>token is not user token</td></tr><tr><td>73036</td><td>insufficient member role permission</td></tr></table>
*/
type UpdateSingleGroupPartialPublicV1Forbidden struct {
	Payload *groupclientmodels.ResponseErrorResponse
}

func (o *UpdateSingleGroupPartialPublicV1Forbidden) Error() string {
	return fmt.Sprintf("[PATCH /group/v1/public/namespaces/{namespace}/groups/{groupId}][%d] updateSingleGroupPartialPublicV1Forbidden  %+v", 403, o.Payload)
}

func (o *UpdateSingleGroupPartialPublicV1Forbidden) GetPayload() *groupclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *UpdateSingleGroupPartialPublicV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(groupclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSingleGroupPartialPublicV1NotFound creates a UpdateSingleGroupPartialPublicV1NotFound with default headers values
func NewUpdateSingleGroupPartialPublicV1NotFound() *UpdateSingleGroupPartialPublicV1NotFound {
	return &UpdateSingleGroupPartialPublicV1NotFound{}
}

/*UpdateSingleGroupPartialPublicV1NotFound handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>73333</td><td>group not found</td></tr></table>
*/
type UpdateSingleGroupPartialPublicV1NotFound struct {
	Payload *groupclientmodels.ResponseErrorResponse
}

func (o *UpdateSingleGroupPartialPublicV1NotFound) Error() string {
	return fmt.Sprintf("[PATCH /group/v1/public/namespaces/{namespace}/groups/{groupId}][%d] updateSingleGroupPartialPublicV1NotFound  %+v", 404, o.Payload)
}

func (o *UpdateSingleGroupPartialPublicV1NotFound) GetPayload() *groupclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *UpdateSingleGroupPartialPublicV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(groupclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSingleGroupPartialPublicV1InternalServerError creates a UpdateSingleGroupPartialPublicV1InternalServerError with default headers values
func NewUpdateSingleGroupPartialPublicV1InternalServerError() *UpdateSingleGroupPartialPublicV1InternalServerError {
	return &UpdateSingleGroupPartialPublicV1InternalServerError{}
}

/*UpdateSingleGroupPartialPublicV1InternalServerError handles this case with default header values.

  Internal Server Error
*/
type UpdateSingleGroupPartialPublicV1InternalServerError struct {
	Payload *groupclientmodels.ResponseErrorResponse
}

func (o *UpdateSingleGroupPartialPublicV1InternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /group/v1/public/namespaces/{namespace}/groups/{groupId}][%d] updateSingleGroupPartialPublicV1InternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateSingleGroupPartialPublicV1InternalServerError) GetPayload() *groupclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *UpdateSingleGroupPartialPublicV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(groupclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
