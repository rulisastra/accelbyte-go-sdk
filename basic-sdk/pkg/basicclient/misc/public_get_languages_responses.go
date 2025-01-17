// Code generated by go-swagger; DO NOT EDIT.

package misc

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

// PublicGetLanguagesReader is a Reader for the PublicGetLanguages structure.
type PublicGetLanguagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicGetLanguagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicGetLanguagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPublicGetLanguagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /v1/public/namespaces/{namespace}/misc/languages returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicGetLanguagesOK creates a PublicGetLanguagesOK with default headers values
func NewPublicGetLanguagesOK() *PublicGetLanguagesOK {
	return &PublicGetLanguagesOK{}
}

/*PublicGetLanguagesOK handles this case with default header values.

  successful operation
*/
type PublicGetLanguagesOK struct {
	Payload map[string]interface{}
}

func (o *PublicGetLanguagesOK) Error() string {
	return fmt.Sprintf("[GET /v1/public/namespaces/{namespace}/misc/languages][%d] publicGetLanguagesOK  %+v", 200, o.Payload)
}

func (o *PublicGetLanguagesOK) GetPayload() map[string]interface{} {
	return o.Payload
}

func (o *PublicGetLanguagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicGetLanguagesBadRequest creates a PublicGetLanguagesBadRequest with default headers values
func NewPublicGetLanguagesBadRequest() *PublicGetLanguagesBadRequest {
	return &PublicGetLanguagesBadRequest{}
}

/*PublicGetLanguagesBadRequest handles this case with default header values.

  <table><tr><td>errorCode</td><td>errorMessage</td></tr><tr><td>20002</td><td>validation error</td></tr></table>
*/
type PublicGetLanguagesBadRequest struct {
	Payload *basicclientmodels.ValidationErrorEntity
}

func (o *PublicGetLanguagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/public/namespaces/{namespace}/misc/languages][%d] publicGetLanguagesBadRequest  %+v", 400, o.Payload)
}

func (o *PublicGetLanguagesBadRequest) GetPayload() *basicclientmodels.ValidationErrorEntity {
	return o.Payload
}

func (o *PublicGetLanguagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(basicclientmodels.ValidationErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
