// Code generated by go-swagger; DO NOT EDIT.

package user_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/basic-sdk/pkg/basicclientmodels"
)

// UpdateCustomAttributesPartially1Reader is a Reader for the UpdateCustomAttributesPartially1 structure.
type UpdateCustomAttributesPartially1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCustomAttributesPartially1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCustomAttributesPartially1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateCustomAttributesPartially1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateCustomAttributesPartially1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateCustomAttributesPartially1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateCustomAttributesPartially1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /v1/public/namespaces/{namespace}/users/{userId}/profiles/customAttributes returns an error %d: %s", response.Code(), string(data))
	}
}

// NewUpdateCustomAttributesPartially1OK creates a UpdateCustomAttributesPartially1OK with default headers values
func NewUpdateCustomAttributesPartially1OK() *UpdateCustomAttributesPartially1OK {
	return &UpdateCustomAttributesPartially1OK{}
}

/*UpdateCustomAttributesPartially1OK handles this case with default header values.

  Successful operation
*/
type UpdateCustomAttributesPartially1OK struct {
	Payload map[string]interface{}
}

func (o *UpdateCustomAttributesPartially1OK) Error() string {
	return fmt.Sprintf("[PUT /v1/public/namespaces/{namespace}/users/{userId}/profiles/customAttributes][%d] updateCustomAttributesPartially1OK  %+v", 200, o.Payload)
}

func (o *UpdateCustomAttributesPartially1OK) GetPayload() map[string]interface{} {
	return o.Payload
}

func (o *UpdateCustomAttributesPartially1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCustomAttributesPartially1BadRequest creates a UpdateCustomAttributesPartially1BadRequest with default headers values
func NewUpdateCustomAttributesPartially1BadRequest() *UpdateCustomAttributesPartially1BadRequest {
	return &UpdateCustomAttributesPartially1BadRequest{}
}

/*UpdateCustomAttributesPartially1BadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr><tr><td>20019</td><td>unable to parse request body</td></tr></table>
*/
type UpdateCustomAttributesPartially1BadRequest struct {
	Payload *basicclientmodels.ValidationErrorEntity
}

func (o *UpdateCustomAttributesPartially1BadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/public/namespaces/{namespace}/users/{userId}/profiles/customAttributes][%d] updateCustomAttributesPartially1BadRequest  %+v", 400, o.Payload)
}

func (o *UpdateCustomAttributesPartially1BadRequest) GetPayload() *basicclientmodels.ValidationErrorEntity {
	return o.Payload
}

func (o *UpdateCustomAttributesPartially1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(basicclientmodels.ValidationErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCustomAttributesPartially1Unauthorized creates a UpdateCustomAttributesPartially1Unauthorized with default headers values
func NewUpdateCustomAttributesPartially1Unauthorized() *UpdateCustomAttributesPartially1Unauthorized {
	return &UpdateCustomAttributesPartially1Unauthorized{}
}

/*UpdateCustomAttributesPartially1Unauthorized handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20001</td><td>unauthorized</td></tr></table>
*/
type UpdateCustomAttributesPartially1Unauthorized struct {
	Payload *basicclientmodels.ErrorEntity
}

func (o *UpdateCustomAttributesPartially1Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/public/namespaces/{namespace}/users/{userId}/profiles/customAttributes][%d] updateCustomAttributesPartially1Unauthorized  %+v", 401, o.Payload)
}

func (o *UpdateCustomAttributesPartially1Unauthorized) GetPayload() *basicclientmodels.ErrorEntity {
	return o.Payload
}

func (o *UpdateCustomAttributesPartially1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(basicclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCustomAttributesPartially1Forbidden creates a UpdateCustomAttributesPartially1Forbidden with default headers values
func NewUpdateCustomAttributesPartially1Forbidden() *UpdateCustomAttributesPartially1Forbidden {
	return &UpdateCustomAttributesPartially1Forbidden{}
}

/*UpdateCustomAttributesPartially1Forbidden handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20013</td><td>insufficient permission</td></tr></table>
*/
type UpdateCustomAttributesPartially1Forbidden struct {
	Payload *basicclientmodels.ErrorEntity
}

func (o *UpdateCustomAttributesPartially1Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/public/namespaces/{namespace}/users/{userId}/profiles/customAttributes][%d] updateCustomAttributesPartially1Forbidden  %+v", 403, o.Payload)
}

func (o *UpdateCustomAttributesPartially1Forbidden) GetPayload() *basicclientmodels.ErrorEntity {
	return o.Payload
}

func (o *UpdateCustomAttributesPartially1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(basicclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCustomAttributesPartially1NotFound creates a UpdateCustomAttributesPartially1NotFound with default headers values
func NewUpdateCustomAttributesPartially1NotFound() *UpdateCustomAttributesPartially1NotFound {
	return &UpdateCustomAttributesPartially1NotFound{}
}

/*UpdateCustomAttributesPartially1NotFound handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>11440</td><td>user profile not found</td></tr></table>
*/
type UpdateCustomAttributesPartially1NotFound struct {
	Payload *basicclientmodels.ErrorEntity
}

func (o *UpdateCustomAttributesPartially1NotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/public/namespaces/{namespace}/users/{userId}/profiles/customAttributes][%d] updateCustomAttributesPartially1NotFound  %+v", 404, o.Payload)
}

func (o *UpdateCustomAttributesPartially1NotFound) GetPayload() *basicclientmodels.ErrorEntity {
	return o.Payload
}

func (o *UpdateCustomAttributesPartially1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(basicclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
