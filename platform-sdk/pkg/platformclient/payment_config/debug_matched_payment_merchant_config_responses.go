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

// DebugMatchedPaymentMerchantConfigReader is a Reader for the DebugMatchedPaymentMerchantConfig structure.
type DebugMatchedPaymentMerchantConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DebugMatchedPaymentMerchantConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDebugMatchedPaymentMerchantConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDebugMatchedPaymentMerchantConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /admin/payment/config/merchant/matched returns an error %d: %s", response.Code(), string(data))
	}
}

// NewDebugMatchedPaymentMerchantConfigOK creates a DebugMatchedPaymentMerchantConfigOK with default headers values
func NewDebugMatchedPaymentMerchantConfigOK() *DebugMatchedPaymentMerchantConfigOK {
	return &DebugMatchedPaymentMerchantConfigOK{}
}

/*DebugMatchedPaymentMerchantConfigOK handles this case with default header values.

  successful operation
*/
type DebugMatchedPaymentMerchantConfigOK struct {
	Payload *platformclientmodels.PaymentMerchantConfigInfo
}

func (o *DebugMatchedPaymentMerchantConfigOK) Error() string {
	return fmt.Sprintf("[GET /admin/payment/config/merchant/matched][%d] debugMatchedPaymentMerchantConfigOK  %+v", 200, o.Payload)
}

func (o *DebugMatchedPaymentMerchantConfigOK) GetPayload() *platformclientmodels.PaymentMerchantConfigInfo {
	return o.Payload
}

func (o *DebugMatchedPaymentMerchantConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.PaymentMerchantConfigInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDebugMatchedPaymentMerchantConfigNotFound creates a DebugMatchedPaymentMerchantConfigNotFound with default headers values
func NewDebugMatchedPaymentMerchantConfigNotFound() *DebugMatchedPaymentMerchantConfigNotFound {
	return &DebugMatchedPaymentMerchantConfigNotFound{}
}

/*DebugMatchedPaymentMerchantConfigNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>33242</td><td>Payment merchant config [{id}] does not exist</td></tr></table>
*/
type DebugMatchedPaymentMerchantConfigNotFound struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *DebugMatchedPaymentMerchantConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /admin/payment/config/merchant/matched][%d] debugMatchedPaymentMerchantConfigNotFound  %+v", 404, o.Payload)
}

func (o *DebugMatchedPaymentMerchantConfigNotFound) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *DebugMatchedPaymentMerchantConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}