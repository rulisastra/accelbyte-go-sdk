// Code generated by go-swagger; DO NOT EDIT.

package stat_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/social-sdk/pkg/socialclientmodels"
)

// GetStatReader is a Reader for the GetStat structure.
type GetStatReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStatReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStatOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetStatNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /v1/admin/namespaces/{namespace}/stats/{statCode} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetStatOK creates a GetStatOK with default headers values
func NewGetStatOK() *GetStatOK {
	return &GetStatOK{}
}

/*GetStatOK handles this case with default header values.

  successful operation
*/
type GetStatOK struct {
	Payload *socialclientmodels.StatInfo
}

func (o *GetStatOK) Error() string {
	return fmt.Sprintf("[GET /v1/admin/namespaces/{namespace}/stats/{statCode}][%d] getStatOK  %+v", 200, o.Payload)
}

func (o *GetStatOK) GetPayload() *socialclientmodels.StatInfo {
	return o.Payload
}

func (o *GetStatOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(socialclientmodels.StatInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatNotFound creates a GetStatNotFound with default headers values
func NewGetStatNotFound() *GetStatNotFound {
	return &GetStatNotFound{}
}

/*GetStatNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>12241</td><td>Stat [{statCode}] cannot be found in namespace [{namespace}]</td></tr></table>
*/
type GetStatNotFound struct {
	Payload *socialclientmodels.ErrorEntity
}

func (o *GetStatNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/admin/namespaces/{namespace}/stats/{statCode}][%d] getStatNotFound  %+v", 404, o.Payload)
}

func (o *GetStatNotFound) GetPayload() *socialclientmodels.ErrorEntity {
	return o.Payload
}

func (o *GetStatNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(socialclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}