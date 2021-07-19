// Code generated by go-swagger; DO NOT EDIT.

package entitlement

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

// PublicGetUserAppEntitlementOwnershipByAppIDReader is a Reader for the PublicGetUserAppEntitlementOwnershipByAppID structure.
type PublicGetUserAppEntitlementOwnershipByAppIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicGetUserAppEntitlementOwnershipByAppIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicGetUserAppEntitlementOwnershipByAppIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /public/namespaces/{namespace}/users/{userId}/entitlements/ownership/byAppId returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicGetUserAppEntitlementOwnershipByAppIDOK creates a PublicGetUserAppEntitlementOwnershipByAppIDOK with default headers values
func NewPublicGetUserAppEntitlementOwnershipByAppIDOK() *PublicGetUserAppEntitlementOwnershipByAppIDOK {
	return &PublicGetUserAppEntitlementOwnershipByAppIDOK{}
}

/*PublicGetUserAppEntitlementOwnershipByAppIDOK handles this case with default header values.

  successful operation
*/
type PublicGetUserAppEntitlementOwnershipByAppIDOK struct {
	Payload *platformclientmodels.Ownership
}

func (o *PublicGetUserAppEntitlementOwnershipByAppIDOK) Error() string {
	return fmt.Sprintf("[GET /public/namespaces/{namespace}/users/{userId}/entitlements/ownership/byAppId][%d] publicGetUserAppEntitlementOwnershipByAppIdOK  %+v", 200, o.Payload)
}

func (o *PublicGetUserAppEntitlementOwnershipByAppIDOK) GetPayload() *platformclientmodels.Ownership {
	return o.Payload
}

func (o *PublicGetUserAppEntitlementOwnershipByAppIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.Ownership)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}