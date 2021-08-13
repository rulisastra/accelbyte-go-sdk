// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclientmodels"
)

// GetUserMappingReader is a Reader for the GetUserMapping structure.
type GetUserMappingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserMappingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserMappingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserMappingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUserMappingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetUserMappingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetUserMappingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /iam/namespaces/{namespace}/users/{userId}/platforms/justice/{targetNamespace} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetUserMappingOK creates a GetUserMappingOK with default headers values
func NewGetUserMappingOK() *GetUserMappingOK {
	return &GetUserMappingOK{}
}

/*GetUserMappingOK handles this case with default header values.

  OK
*/
type GetUserMappingOK struct {
	Payload *iamclientmodels.ModelGetUserMapping
}

func (o *GetUserMappingOK) Error() string {
	return fmt.Sprintf("[GET /iam/namespaces/{namespace}/users/{userId}/platforms/justice/{targetNamespace}][%d] getUserMappingOK  %+v", 200, o.Payload)
}

func (o *GetUserMappingOK) GetPayload() *iamclientmodels.ModelGetUserMapping {
	return o.Payload
}

func (o *GetUserMappingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.ModelGetUserMapping)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserMappingBadRequest creates a GetUserMappingBadRequest with default headers values
func NewGetUserMappingBadRequest() *GetUserMappingBadRequest {
	return &GetUserMappingBadRequest{}
}

/*GetUserMappingBadRequest handles this case with default header values.

  Invalid request
*/
type GetUserMappingBadRequest struct {
}

func (o *GetUserMappingBadRequest) Error() string {
	return fmt.Sprintf("[GET /iam/namespaces/{namespace}/users/{userId}/platforms/justice/{targetNamespace}][%d] getUserMappingBadRequest ", 400)
}

func (o *GetUserMappingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserMappingUnauthorized creates a GetUserMappingUnauthorized with default headers values
func NewGetUserMappingUnauthorized() *GetUserMappingUnauthorized {
	return &GetUserMappingUnauthorized{}
}

/*GetUserMappingUnauthorized handles this case with default header values.

  Unauthorized access
*/
type GetUserMappingUnauthorized struct {
}

func (o *GetUserMappingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /iam/namespaces/{namespace}/users/{userId}/platforms/justice/{targetNamespace}][%d] getUserMappingUnauthorized ", 401)
}

func (o *GetUserMappingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserMappingForbidden creates a GetUserMappingForbidden with default headers values
func NewGetUserMappingForbidden() *GetUserMappingForbidden {
	return &GetUserMappingForbidden{}
}

/*GetUserMappingForbidden handles this case with default header values.

  Forbidden
*/
type GetUserMappingForbidden struct {
}

func (o *GetUserMappingForbidden) Error() string {
	return fmt.Sprintf("[GET /iam/namespaces/{namespace}/users/{userId}/platforms/justice/{targetNamespace}][%d] getUserMappingForbidden ", 403)
}

func (o *GetUserMappingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserMappingNotFound creates a GetUserMappingNotFound with default headers values
func NewGetUserMappingNotFound() *GetUserMappingNotFound {
	return &GetUserMappingNotFound{}
}

/*GetUserMappingNotFound handles this case with default header values.

  Data not found
*/
type GetUserMappingNotFound struct {
}

func (o *GetUserMappingNotFound) Error() string {
	return fmt.Sprintf("[GET /iam/namespaces/{namespace}/users/{userId}/platforms/justice/{targetNamespace}][%d] getUserMappingNotFound ", 404)
}

func (o *GetUserMappingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}