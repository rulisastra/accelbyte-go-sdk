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

// NewPublicFulfillGoogleIAPItemParams creates a new PublicFulfillGoogleIAPItemParams object
// with the default values initialized.
func NewPublicFulfillGoogleIAPItemParams() *PublicFulfillGoogleIAPItemParams {
	var ()
	return &PublicFulfillGoogleIAPItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicFulfillGoogleIAPItemParamsWithTimeout creates a new PublicFulfillGoogleIAPItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicFulfillGoogleIAPItemParamsWithTimeout(timeout time.Duration) *PublicFulfillGoogleIAPItemParams {
	var ()
	return &PublicFulfillGoogleIAPItemParams{

		timeout: timeout,
	}
}

// NewPublicFulfillGoogleIAPItemParamsWithContext creates a new PublicFulfillGoogleIAPItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicFulfillGoogleIAPItemParamsWithContext(ctx context.Context) *PublicFulfillGoogleIAPItemParams {
	var ()
	return &PublicFulfillGoogleIAPItemParams{

		Context: ctx,
	}
}

// NewPublicFulfillGoogleIAPItemParamsWithHTTPClient creates a new PublicFulfillGoogleIAPItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicFulfillGoogleIAPItemParamsWithHTTPClient(client *http.Client) *PublicFulfillGoogleIAPItemParams {
	var ()
	return &PublicFulfillGoogleIAPItemParams{
		HTTPClient: client,
	}
}

/*PublicFulfillGoogleIAPItemParams contains all the parameters to send to the API endpoint
for the public fulfill google i a p item operation typically these are written to a http.Request
*/
type PublicFulfillGoogleIAPItemParams struct {

	/*Body*/
	Body *platformclientmodels.GoogleIAPReceipt
	/*Namespace*/
	Namespace string
	/*UserID*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the public fulfill google i a p item params
func (o *PublicFulfillGoogleIAPItemParams) WithTimeout(timeout time.Duration) *PublicFulfillGoogleIAPItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public fulfill google i a p item params
func (o *PublicFulfillGoogleIAPItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public fulfill google i a p item params
func (o *PublicFulfillGoogleIAPItemParams) WithContext(ctx context.Context) *PublicFulfillGoogleIAPItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public fulfill google i a p item params
func (o *PublicFulfillGoogleIAPItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public fulfill google i a p item params
func (o *PublicFulfillGoogleIAPItemParams) WithHTTPClient(client *http.Client) *PublicFulfillGoogleIAPItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public fulfill google i a p item params
func (o *PublicFulfillGoogleIAPItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the public fulfill google i a p item params
func (o *PublicFulfillGoogleIAPItemParams) WithBody(body *platformclientmodels.GoogleIAPReceipt) *PublicFulfillGoogleIAPItemParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the public fulfill google i a p item params
func (o *PublicFulfillGoogleIAPItemParams) SetBody(body *platformclientmodels.GoogleIAPReceipt) {
	o.Body = body
}

// WithNamespace adds the namespace to the public fulfill google i a p item params
func (o *PublicFulfillGoogleIAPItemParams) WithNamespace(namespace string) *PublicFulfillGoogleIAPItemParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public fulfill google i a p item params
func (o *PublicFulfillGoogleIAPItemParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the public fulfill google i a p item params
func (o *PublicFulfillGoogleIAPItemParams) WithUserID(userID string) *PublicFulfillGoogleIAPItemParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the public fulfill google i a p item params
func (o *PublicFulfillGoogleIAPItemParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PublicFulfillGoogleIAPItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
