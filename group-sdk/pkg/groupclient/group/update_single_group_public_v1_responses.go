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

// UpdateSingleGroupPublicV1Reader is a Reader for the UpdateSingleGroupPublicV1 structure.
type UpdateSingleGroupPublicV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSingleGroupPublicV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSingleGroupPublicV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateSingleGroupPublicV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateSingleGroupPublicV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateSingleGroupPublicV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateSingleGroupPublicV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewUpdateSingleGroupPublicV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /group/v1/public/namespaces/{namespace}/groups/{groupId} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewUpdateSingleGroupPublicV1OK creates a UpdateSingleGroupPublicV1OK with default headers values
func NewUpdateSingleGroupPublicV1OK() *UpdateSingleGroupPublicV1OK {
	return &UpdateSingleGroupPublicV1OK{}
}

/*UpdateSingleGroupPublicV1OK handles this case with default header values.

  OK
*/
type UpdateSingleGroupPublicV1OK struct {
	Payload *groupclientmodels.ModelsGroupResponseV1
}

func (o *UpdateSingleGroupPublicV1OK) Error() string {
	return fmt.Sprintf("[PUT /group/v1/public/namespaces/{namespace}/groups/{groupId}][%d] updateSingleGroupPublicV1OK  %+v", 200, o.Payload)
}

func (o *UpdateSingleGroupPublicV1OK) GetPayload() *groupclientmodels.ModelsGroupResponseV1 {
	return o.Payload
}

func (o *UpdateSingleGroupPublicV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(groupclientmodels.ModelsGroupResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSingleGroupPublicV1BadRequest creates a UpdateSingleGroupPublicV1BadRequest with default headers values
func NewUpdateSingleGroupPublicV1BadRequest() *UpdateSingleGroupPublicV1BadRequest {
	return &UpdateSingleGroupPublicV1BadRequest{}
}

/*UpdateSingleGroupPublicV1BadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type UpdateSingleGroupPublicV1BadRequest struct {
	Payload *groupclientmodels.ResponseErrorResponse
}

func (o *UpdateSingleGroupPublicV1BadRequest) Error() string {
	return fmt.Sprintf("[PUT /group/v1/public/namespaces/{namespace}/groups/{groupId}][%d] updateSingleGroupPublicV1BadRequest  %+v", 400, o.Payload)
}

func (o *UpdateSingleGroupPublicV1BadRequest) GetPayload() *groupclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *UpdateSingleGroupPublicV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(groupclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSingleGroupPublicV1Unauthorized creates a UpdateSingleGroupPublicV1Unauthorized with default headers values
func NewUpdateSingleGroupPublicV1Unauthorized() *UpdateSingleGroupPublicV1Unauthorized {
	return &UpdateSingleGroupPublicV1Unauthorized{}
}

/*UpdateSingleGroupPublicV1Unauthorized handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20001</td><td>unauthorized access</td></tr></table>
*/
type UpdateSingleGroupPublicV1Unauthorized struct {
	Payload *groupclientmodels.ResponseErrorResponse
}

func (o *UpdateSingleGroupPublicV1Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /group/v1/public/namespaces/{namespace}/groups/{groupId}][%d] updateSingleGroupPublicV1Unauthorized  %+v", 401, o.Payload)
}

func (o *UpdateSingleGroupPublicV1Unauthorized) GetPayload() *groupclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *UpdateSingleGroupPublicV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(groupclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSingleGroupPublicV1Forbidden creates a UpdateSingleGroupPublicV1Forbidden with default headers values
func NewUpdateSingleGroupPublicV1Forbidden() *UpdateSingleGroupPublicV1Forbidden {
	return &UpdateSingleGroupPublicV1Forbidden{}
}

/*UpdateSingleGroupPublicV1Forbidden handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20013</td><td>insufficient permissions</td></tr><tr><td>20022</td><td>token is not user token</td></tr><tr><td>73036</td><td>insufficient member role permission</td></tr></table>
*/
type UpdateSingleGroupPublicV1Forbidden struct {
	Payload *groupclientmodels.ResponseErrorResponse
}

func (o *UpdateSingleGroupPublicV1Forbidden) Error() string {
	return fmt.Sprintf("[PUT /group/v1/public/namespaces/{namespace}/groups/{groupId}][%d] updateSingleGroupPublicV1Forbidden  %+v", 403, o.Payload)
}

func (o *UpdateSingleGroupPublicV1Forbidden) GetPayload() *groupclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *UpdateSingleGroupPublicV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(groupclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSingleGroupPublicV1NotFound creates a UpdateSingleGroupPublicV1NotFound with default headers values
func NewUpdateSingleGroupPublicV1NotFound() *UpdateSingleGroupPublicV1NotFound {
	return &UpdateSingleGroupPublicV1NotFound{}
}

/*UpdateSingleGroupPublicV1NotFound handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>73333</td><td>group not found</td></tr></table>
*/
type UpdateSingleGroupPublicV1NotFound struct {
	Payload *groupclientmodels.ResponseErrorResponse
}

func (o *UpdateSingleGroupPublicV1NotFound) Error() string {
	return fmt.Sprintf("[PUT /group/v1/public/namespaces/{namespace}/groups/{groupId}][%d] updateSingleGroupPublicV1NotFound  %+v", 404, o.Payload)
}

func (o *UpdateSingleGroupPublicV1NotFound) GetPayload() *groupclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *UpdateSingleGroupPublicV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(groupclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSingleGroupPublicV1InternalServerError creates a UpdateSingleGroupPublicV1InternalServerError with default headers values
func NewUpdateSingleGroupPublicV1InternalServerError() *UpdateSingleGroupPublicV1InternalServerError {
	return &UpdateSingleGroupPublicV1InternalServerError{}
}

/*UpdateSingleGroupPublicV1InternalServerError handles this case with default header values.

  Internal Server Error
*/
type UpdateSingleGroupPublicV1InternalServerError struct {
	Payload *groupclientmodels.ResponseErrorResponse
}

func (o *UpdateSingleGroupPublicV1InternalServerError) Error() string {
	return fmt.Sprintf("[PUT /group/v1/public/namespaces/{namespace}/groups/{groupId}][%d] updateSingleGroupPublicV1InternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateSingleGroupPublicV1InternalServerError) GetPayload() *groupclientmodels.ResponseErrorResponse {
	return o.Payload
}

func (o *UpdateSingleGroupPublicV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(groupclientmodels.ResponseErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
