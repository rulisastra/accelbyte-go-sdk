// Code generated by go-swagger; DO NOT EDIT.

package payment_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
)

// TestWxPayConfigByIDReader is a Reader for the TestWxPayConfigByID structure.
type TestWxPayConfigByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestWxPayConfigByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTestWxPayConfigByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewTestWxPayConfigByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /admin/payment/config/merchant/{id}/wxpayconfig/test returns an error %d: %s", response.Code(), string(data))
	}
}

// NewTestWxPayConfigByIDOK creates a TestWxPayConfigByIDOK with default headers values
func NewTestWxPayConfigByIDOK() *TestWxPayConfigByIDOK {
	return &TestWxPayConfigByIDOK{}
}

/*TestWxPayConfigByIDOK handles this case with default header values.

  successful operation
*/
type TestWxPayConfigByIDOK struct {
	Payload *platformclientmodels.TestResult
}

func (o *TestWxPayConfigByIDOK) Error() string {
	return fmt.Sprintf("[GET /admin/payment/config/merchant/{id}/wxpayconfig/test][%d] testWxPayConfigByIdOK  %+v", 200, o.Payload)
}

func (o *TestWxPayConfigByIDOK) GetPayload() *platformclientmodels.TestResult {
	return o.Payload
}

func (o *TestWxPayConfigByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.TestResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestWxPayConfigByIDNotFound creates a TestWxPayConfigByIDNotFound with default headers values
func NewTestWxPayConfigByIDNotFound() *TestWxPayConfigByIDNotFound {
	return &TestWxPayConfigByIDNotFound{}
}

/*TestWxPayConfigByIDNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>33242</td><td>Payment merchant config [{id}] does not exist</td></tr></table>
*/
type TestWxPayConfigByIDNotFound struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *TestWxPayConfigByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /admin/payment/config/merchant/{id}/wxpayconfig/test][%d] testWxPayConfigByIdNotFound  %+v", 404, o.Payload)
}

func (o *TestWxPayConfigByIDNotFound) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *TestWxPayConfigByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}