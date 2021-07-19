// Code generated by go-swagger; DO NOT EDIT.

package i_a_p

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

// UpdateGoogleP12FileReader is a Reader for the UpdateGoogleP12File structure.
type UpdateGoogleP12FileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateGoogleP12FileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateGoogleP12FileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /admin/namespaces/{namespace}/iap/config/google/cert returns an error %d: %s", response.Code(), string(data))
	}
}

// NewUpdateGoogleP12FileOK creates a UpdateGoogleP12FileOK with default headers values
func NewUpdateGoogleP12FileOK() *UpdateGoogleP12FileOK {
	return &UpdateGoogleP12FileOK{}
}

/*UpdateGoogleP12FileOK handles this case with default header values.

  successful operation
*/
type UpdateGoogleP12FileOK struct {
	Payload *platformclientmodels.GoogleIAPConfigInfo
}

func (o *UpdateGoogleP12FileOK) Error() string {
	return fmt.Sprintf("[PUT /admin/namespaces/{namespace}/iap/config/google/cert][%d] updateGoogleP12FileOK  %+v", 200, o.Payload)
}

func (o *UpdateGoogleP12FileOK) GetPayload() *platformclientmodels.GoogleIAPConfigInfo {
	return o.Payload
}

func (o *UpdateGoogleP12FileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.GoogleIAPConfigInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}