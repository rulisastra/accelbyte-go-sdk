// Code generated by go-swagger; DO NOT EDIT.

package i_a_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
)

// NewPublicFulfillAppleIAPItemParams creates a new PublicFulfillAppleIAPItemParams object
// with the default values initialized.
func NewPublicFulfillAppleIAPItemParams() *PublicFulfillAppleIAPItemParams {
	var ()
	return &PublicFulfillAppleIAPItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicFulfillAppleIAPItemParamsWithTimeout creates a new PublicFulfillAppleIAPItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicFulfillAppleIAPItemParamsWithTimeout(timeout time.Duration) *PublicFulfillAppleIAPItemParams {
	var ()
	return &PublicFulfillAppleIAPItemParams{

		timeout: timeout,
	}
}

// NewPublicFulfillAppleIAPItemParamsWithContext creates a new PublicFulfillAppleIAPItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicFulfillAppleIAPItemParamsWithContext(ctx context.Context) *PublicFulfillAppleIAPItemParams {
	var ()
	return &PublicFulfillAppleIAPItemParams{

		Context: ctx,
	}
}

// NewPublicFulfillAppleIAPItemParamsWithHTTPClient creates a new PublicFulfillAppleIAPItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicFulfillAppleIAPItemParamsWithHTTPClient(client *http.Client) *PublicFulfillAppleIAPItemParams {
	var ()
	return &PublicFulfillAppleIAPItemParams{
		HTTPClient: client,
	}
}

/*PublicFulfillAppleIAPItemParams contains all the parameters to send to the API endpoint
for the public fulfill apple i a p item operation typically these are written to a http.Request
*/
type PublicFulfillAppleIAPItemParams struct {

	/*Body*/
	Body *platformclientmodels.AppleIAPReceipt
	/*Namespace*/
	Namespace string
	/*UserID*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the public fulfill apple i a p item params
func (o *PublicFulfillAppleIAPItemParams) WithTimeout(timeout time.Duration) *PublicFulfillAppleIAPItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public fulfill apple i a p item params
func (o *PublicFulfillAppleIAPItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public fulfill apple i a p item params
func (o *PublicFulfillAppleIAPItemParams) WithContext(ctx context.Context) *PublicFulfillAppleIAPItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public fulfill apple i a p item params
func (o *PublicFulfillAppleIAPItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public fulfill apple i a p item params
func (o *PublicFulfillAppleIAPItemParams) WithHTTPClient(client *http.Client) *PublicFulfillAppleIAPItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public fulfill apple i a p item params
func (o *PublicFulfillAppleIAPItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the public fulfill apple i a p item params
func (o *PublicFulfillAppleIAPItemParams) WithBody(body *platformclientmodels.AppleIAPReceipt) *PublicFulfillAppleIAPItemParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the public fulfill apple i a p item params
func (o *PublicFulfillAppleIAPItemParams) SetBody(body *platformclientmodels.AppleIAPReceipt) {
	o.Body = body
}

// WithNamespace adds the namespace to the public fulfill apple i a p item params
func (o *PublicFulfillAppleIAPItemParams) WithNamespace(namespace string) *PublicFulfillAppleIAPItemParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public fulfill apple i a p item params
func (o *PublicFulfillAppleIAPItemParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the public fulfill apple i a p item params
func (o *PublicFulfillAppleIAPItemParams) WithUserID(userID string) *PublicFulfillAppleIAPItemParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the public fulfill apple i a p item params
func (o *PublicFulfillAppleIAPItemParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PublicFulfillAppleIAPItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
