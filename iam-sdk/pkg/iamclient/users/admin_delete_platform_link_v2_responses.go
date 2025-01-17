// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AdminDeletePlatformLinkV2Reader is a Reader for the AdminDeletePlatformLinkV2 structure.
type AdminDeletePlatformLinkV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminDeletePlatformLinkV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAdminDeletePlatformLinkV2NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminDeletePlatformLinkV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminDeletePlatformLinkV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAdminDeletePlatformLinkV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAdminDeletePlatformLinkV2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewAdminDeletePlatformLinkV2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /iam/v2/admin/namespaces/{namespace}/users/{userId}/platforms/{platformId}/link returns an error %d: %s", response.Code(), string(data))
	}
}

// NewAdminDeletePlatformLinkV2NoContent creates a AdminDeletePlatformLinkV2NoContent with default headers values
func NewAdminDeletePlatformLinkV2NoContent() *AdminDeletePlatformLinkV2NoContent {
	return &AdminDeletePlatformLinkV2NoContent{}
}

/*AdminDeletePlatformLinkV2NoContent handles this case with default header values.

  Operation succeeded
*/
type AdminDeletePlatformLinkV2NoContent struct {
}

func (o *AdminDeletePlatformLinkV2NoContent) Error() string {
	return fmt.Sprintf("[DELETE /iam/v2/admin/namespaces/{namespace}/users/{userId}/platforms/{platformId}/link][%d] adminDeletePlatformLinkV2NoContent ", 204)
}

func (o *AdminDeletePlatformLinkV2NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminDeletePlatformLinkV2BadRequest creates a AdminDeletePlatformLinkV2BadRequest with default headers values
func NewAdminDeletePlatformLinkV2BadRequest() *AdminDeletePlatformLinkV2BadRequest {
	return &AdminDeletePlatformLinkV2BadRequest{}
}

/*AdminDeletePlatformLinkV2BadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20019</td><td>unable to parse request body</td></tr></table>
*/
type AdminDeletePlatformLinkV2BadRequest struct {
}

func (o *AdminDeletePlatformLinkV2BadRequest) Error() string {
	return fmt.Sprintf("[DELETE /iam/v2/admin/namespaces/{namespace}/users/{userId}/platforms/{platformId}/link][%d] adminDeletePlatformLinkV2BadRequest ", 400)
}

func (o *AdminDeletePlatformLinkV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminDeletePlatformLinkV2Unauthorized creates a AdminDeletePlatformLinkV2Unauthorized with default headers values
func NewAdminDeletePlatformLinkV2Unauthorized() *AdminDeletePlatformLinkV2Unauthorized {
	return &AdminDeletePlatformLinkV2Unauthorized{}
}

/*AdminDeletePlatformLinkV2Unauthorized handles this case with default header values.

  Unauthorized access
*/
type AdminDeletePlatformLinkV2Unauthorized struct {
}

func (o *AdminDeletePlatformLinkV2Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /iam/v2/admin/namespaces/{namespace}/users/{userId}/platforms/{platformId}/link][%d] adminDeletePlatformLinkV2Unauthorized ", 401)
}

func (o *AdminDeletePlatformLinkV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminDeletePlatformLinkV2Forbidden creates a AdminDeletePlatformLinkV2Forbidden with default headers values
func NewAdminDeletePlatformLinkV2Forbidden() *AdminDeletePlatformLinkV2Forbidden {
	return &AdminDeletePlatformLinkV2Forbidden{}
}

/*AdminDeletePlatformLinkV2Forbidden handles this case with default header values.

  Forbidden
*/
type AdminDeletePlatformLinkV2Forbidden struct {
}

func (o *AdminDeletePlatformLinkV2Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /iam/v2/admin/namespaces/{namespace}/users/{userId}/platforms/{platformId}/link][%d] adminDeletePlatformLinkV2Forbidden ", 403)
}

func (o *AdminDeletePlatformLinkV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminDeletePlatformLinkV2NotFound creates a AdminDeletePlatformLinkV2NotFound with default headers values
func NewAdminDeletePlatformLinkV2NotFound() *AdminDeletePlatformLinkV2NotFound {
	return &AdminDeletePlatformLinkV2NotFound{}
}

/*AdminDeletePlatformLinkV2NotFound handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20008</td><td>user not found</td></tr></table>
*/
type AdminDeletePlatformLinkV2NotFound struct {
}

func (o *AdminDeletePlatformLinkV2NotFound) Error() string {
	return fmt.Sprintf("[DELETE /iam/v2/admin/namespaces/{namespace}/users/{userId}/platforms/{platformId}/link][%d] adminDeletePlatformLinkV2NotFound ", 404)
}

func (o *AdminDeletePlatformLinkV2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminDeletePlatformLinkV2InternalServerError creates a AdminDeletePlatformLinkV2InternalServerError with default headers values
func NewAdminDeletePlatformLinkV2InternalServerError() *AdminDeletePlatformLinkV2InternalServerError {
	return &AdminDeletePlatformLinkV2InternalServerError{}
}

/*AdminDeletePlatformLinkV2InternalServerError handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20000</td><td>internal server error</td></tr></table>
*/
type AdminDeletePlatformLinkV2InternalServerError struct {
}

func (o *AdminDeletePlatformLinkV2InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /iam/v2/admin/namespaces/{namespace}/users/{userId}/platforms/{platformId}/link][%d] adminDeletePlatformLinkV2InternalServerError ", 500)
}

func (o *AdminDeletePlatformLinkV2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
