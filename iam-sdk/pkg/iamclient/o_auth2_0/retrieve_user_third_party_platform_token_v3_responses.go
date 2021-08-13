// Code generated by go-swagger; DO NOT EDIT.

package o_auth2_0

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

// RetrieveUserThirdPartyPlatformTokenV3Reader is a Reader for the RetrieveUserThirdPartyPlatformTokenV3 structure.
type RetrieveUserThirdPartyPlatformTokenV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveUserThirdPartyPlatformTokenV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetrieveUserThirdPartyPlatformTokenV3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRetrieveUserThirdPartyPlatformTokenV3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRetrieveUserThirdPartyPlatformTokenV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRetrieveUserThirdPartyPlatformTokenV3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /iam/v3/oauth/namespaces/{namespace}/users/{userId}/platforms/{platformId}/platformToken returns an error %d: %s", response.Code(), string(data))
	}
}

// NewRetrieveUserThirdPartyPlatformTokenV3OK creates a RetrieveUserThirdPartyPlatformTokenV3OK with default headers values
func NewRetrieveUserThirdPartyPlatformTokenV3OK() *RetrieveUserThirdPartyPlatformTokenV3OK {
	return &RetrieveUserThirdPartyPlatformTokenV3OK{}
}

/*RetrieveUserThirdPartyPlatformTokenV3OK handles this case with default header values.

  Token returned
*/
type RetrieveUserThirdPartyPlatformTokenV3OK struct {
	Payload *iamclientmodels.OauthmodelTokenThirdPartyResponse
}

func (o *RetrieveUserThirdPartyPlatformTokenV3OK) Error() string {
	return fmt.Sprintf("[GET /iam/v3/oauth/namespaces/{namespace}/users/{userId}/platforms/{platformId}/platformToken][%d] retrieveUserThirdPartyPlatformTokenV3OK  %+v", 200, o.Payload)
}

func (o *RetrieveUserThirdPartyPlatformTokenV3OK) GetPayload() *iamclientmodels.OauthmodelTokenThirdPartyResponse {
	return o.Payload
}

func (o *RetrieveUserThirdPartyPlatformTokenV3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.OauthmodelTokenThirdPartyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveUserThirdPartyPlatformTokenV3Unauthorized creates a RetrieveUserThirdPartyPlatformTokenV3Unauthorized with default headers values
func NewRetrieveUserThirdPartyPlatformTokenV3Unauthorized() *RetrieveUserThirdPartyPlatformTokenV3Unauthorized {
	return &RetrieveUserThirdPartyPlatformTokenV3Unauthorized{}
}

/*RetrieveUserThirdPartyPlatformTokenV3Unauthorized handles this case with default header values.

  Client authentication failed
*/
type RetrieveUserThirdPartyPlatformTokenV3Unauthorized struct {
	Payload *iamclientmodels.OauthmodelErrorResponse
}

func (o *RetrieveUserThirdPartyPlatformTokenV3Unauthorized) Error() string {
	return fmt.Sprintf("[GET /iam/v3/oauth/namespaces/{namespace}/users/{userId}/platforms/{platformId}/platformToken][%d] retrieveUserThirdPartyPlatformTokenV3Unauthorized  %+v", 401, o.Payload)
}

func (o *RetrieveUserThirdPartyPlatformTokenV3Unauthorized) GetPayload() *iamclientmodels.OauthmodelErrorResponse {
	return o.Payload
}

func (o *RetrieveUserThirdPartyPlatformTokenV3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.OauthmodelErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveUserThirdPartyPlatformTokenV3Forbidden creates a RetrieveUserThirdPartyPlatformTokenV3Forbidden with default headers values
func NewRetrieveUserThirdPartyPlatformTokenV3Forbidden() *RetrieveUserThirdPartyPlatformTokenV3Forbidden {
	return &RetrieveUserThirdPartyPlatformTokenV3Forbidden{}
}

/*RetrieveUserThirdPartyPlatformTokenV3Forbidden handles this case with default header values.

  Unauthorized access
*/
type RetrieveUserThirdPartyPlatformTokenV3Forbidden struct {
	Payload *iamclientmodels.OauthmodelErrorResponse
}

func (o *RetrieveUserThirdPartyPlatformTokenV3Forbidden) Error() string {
	return fmt.Sprintf("[GET /iam/v3/oauth/namespaces/{namespace}/users/{userId}/platforms/{platformId}/platformToken][%d] retrieveUserThirdPartyPlatformTokenV3Forbidden  %+v", 403, o.Payload)
}

func (o *RetrieveUserThirdPartyPlatformTokenV3Forbidden) GetPayload() *iamclientmodels.OauthmodelErrorResponse {
	return o.Payload
}

func (o *RetrieveUserThirdPartyPlatformTokenV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.OauthmodelErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveUserThirdPartyPlatformTokenV3NotFound creates a RetrieveUserThirdPartyPlatformTokenV3NotFound with default headers values
func NewRetrieveUserThirdPartyPlatformTokenV3NotFound() *RetrieveUserThirdPartyPlatformTokenV3NotFound {
	return &RetrieveUserThirdPartyPlatformTokenV3NotFound{}
}

/*RetrieveUserThirdPartyPlatformTokenV3NotFound handles this case with default header values.

  Platform Token Not Found
*/
type RetrieveUserThirdPartyPlatformTokenV3NotFound struct {
	Payload *iamclientmodels.OauthmodelErrorResponse
}

func (o *RetrieveUserThirdPartyPlatformTokenV3NotFound) Error() string {
	return fmt.Sprintf("[GET /iam/v3/oauth/namespaces/{namespace}/users/{userId}/platforms/{platformId}/platformToken][%d] retrieveUserThirdPartyPlatformTokenV3NotFound  %+v", 404, o.Payload)
}

func (o *RetrieveUserThirdPartyPlatformTokenV3NotFound) GetPayload() *iamclientmodels.OauthmodelErrorResponse {
	return o.Payload
}

func (o *RetrieveUserThirdPartyPlatformTokenV3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(iamclientmodels.OauthmodelErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}