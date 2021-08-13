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

// GetUserJusticePlatformAccountReader is a Reader for the GetUserJusticePlatformAccount structure.
type GetUserJusticePlatformAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserJusticePlatformAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserJusticePlatformAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserJusticePlatformAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUserJusticePlatformAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetUserJusticePlatformAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /iam/namespaces/{namespace}/users/{userId}/platforms/justice/{targetNamespace} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetUserJusticePlatformAccountOK creates a GetUserJusticePlatformAccountOK with default headers values
func NewGetUserJusticePlatformAccountOK() *GetUserJusticePlatformAccountOK {
	return &GetUserJusticePlatformAccountOK{}
}

/*GetUserJusticePlatformAccountOK handles this case with default header values.

  OK
*/
type GetUserJusticePlatformAccountOK struct {
	Payload *iamclientmodels.ModelGetUserJusticePlatformAccountResponse
}

func (o *GetUserJusticePlatformAccountOK) Error() string {
	return fmt.Sprintf("[POST /iam/namespaces/{namespace}/users/{userId}/platforms/justice/{targetNamespace}][%d] getUserJusticePlatformAccountOK  %+v", 200, o.Payload)
}

func (o *GetUserJusticePlatformAccountOK) GetPayload() *iamclientmodels.ModelGetUserJusticePlatformAccountResponse {
	return o.Payload
}

func (o *GetUserJusticePlatformAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.ModelGetUserJusticePlatformAccountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserJusticePlatformAccountBadRequest creates a GetUserJusticePlatformAccountBadRequest with default headers values
func NewGetUserJusticePlatformAccountBadRequest() *GetUserJusticePlatformAccountBadRequest {
	return &GetUserJusticePlatformAccountBadRequest{}
}

/*GetUserJusticePlatformAccountBadRequest handles this case with default header values.

  Invalid request
*/
type GetUserJusticePlatformAccountBadRequest struct {
}

func (o *GetUserJusticePlatformAccountBadRequest) Error() string {
	return fmt.Sprintf("[POST /iam/namespaces/{namespace}/users/{userId}/platforms/justice/{targetNamespace}][%d] getUserJusticePlatformAccountBadRequest ", 400)
}

func (o *GetUserJusticePlatformAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserJusticePlatformAccountUnauthorized creates a GetUserJusticePlatformAccountUnauthorized with default headers values
func NewGetUserJusticePlatformAccountUnauthorized() *GetUserJusticePlatformAccountUnauthorized {
	return &GetUserJusticePlatformAccountUnauthorized{}
}

/*GetUserJusticePlatformAccountUnauthorized handles this case with default header values.

  Unauthorized access
*/
type GetUserJusticePlatformAccountUnauthorized struct {
}

func (o *GetUserJusticePlatformAccountUnauthorized) Error() string {
	return fmt.Sprintf("[POST /iam/namespaces/{namespace}/users/{userId}/platforms/justice/{targetNamespace}][%d] getUserJusticePlatformAccountUnauthorized ", 401)
}

func (o *GetUserJusticePlatformAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserJusticePlatformAccountForbidden creates a GetUserJusticePlatformAccountForbidden with default headers values
func NewGetUserJusticePlatformAccountForbidden() *GetUserJusticePlatformAccountForbidden {
	return &GetUserJusticePlatformAccountForbidden{}
}

/*GetUserJusticePlatformAccountForbidden handles this case with default header values.

  Forbidden
*/
type GetUserJusticePlatformAccountForbidden struct {
}

func (o *GetUserJusticePlatformAccountForbidden) Error() string {
	return fmt.Sprintf("[POST /iam/namespaces/{namespace}/users/{userId}/platforms/justice/{targetNamespace}][%d] getUserJusticePlatformAccountForbidden ", 403)
}

func (o *GetUserJusticePlatformAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}