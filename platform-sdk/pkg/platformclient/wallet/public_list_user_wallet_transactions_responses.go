// Code generated by go-swagger; DO NOT EDIT.

package wallet

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

// PublicListUserWalletTransactionsReader is a Reader for the PublicListUserWalletTransactions structure.
type PublicListUserWalletTransactionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicListUserWalletTransactionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicListUserWalletTransactionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /public/namespaces/{namespace}/users/{userId}/wallets/{currencyCode}/transactions returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicListUserWalletTransactionsOK creates a PublicListUserWalletTransactionsOK with default headers values
func NewPublicListUserWalletTransactionsOK() *PublicListUserWalletTransactionsOK {
	return &PublicListUserWalletTransactionsOK{}
}

/*PublicListUserWalletTransactionsOK handles this case with default header values.

  successful operation
*/
type PublicListUserWalletTransactionsOK struct {
	Payload *platformclientmodels.WalletTransactionPagingSlicedResult
}

func (o *PublicListUserWalletTransactionsOK) Error() string {
	return fmt.Sprintf("[GET /public/namespaces/{namespace}/users/{userId}/wallets/{currencyCode}/transactions][%d] publicListUserWalletTransactionsOK  %+v", 200, o.Payload)
}

func (o *PublicListUserWalletTransactionsOK) GetPayload() *platformclientmodels.WalletTransactionPagingSlicedResult {
	return o.Payload
}

func (o *PublicListUserWalletTransactionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.WalletTransactionPagingSlicedResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}