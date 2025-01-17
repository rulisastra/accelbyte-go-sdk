// Code generated by go-swagger; DO NOT EDIT.

package session

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

	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclientmodels"
)

// NewClaimServerParams creates a new ClaimServerParams object
// with the default values initialized.
func NewClaimServerParams() *ClaimServerParams {
	var ()
	return &ClaimServerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewClaimServerParamsWithTimeout creates a new ClaimServerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewClaimServerParamsWithTimeout(timeout time.Duration) *ClaimServerParams {
	var ()
	return &ClaimServerParams{

		timeout: timeout,
	}
}

// NewClaimServerParamsWithContext creates a new ClaimServerParams object
// with the default values initialized, and the ability to set a context for a request
func NewClaimServerParamsWithContext(ctx context.Context) *ClaimServerParams {
	var ()
	return &ClaimServerParams{

		Context: ctx,
	}
}

// NewClaimServerParamsWithHTTPClient creates a new ClaimServerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewClaimServerParamsWithHTTPClient(client *http.Client) *ClaimServerParams {
	var ()
	return &ClaimServerParams{
		HTTPClient: client,
	}
}

/*ClaimServerParams contains all the parameters to send to the API endpoint
for the claim server operation typically these are written to a http.Request
*/
type ClaimServerParams struct {

	/*Body*/
	Body *dsmcclientmodels.ModelsClaimSessionRequest
	/*Namespace
	  namespace of the game

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the claim server params
func (o *ClaimServerParams) WithTimeout(timeout time.Duration) *ClaimServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the claim server params
func (o *ClaimServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the claim server params
func (o *ClaimServerParams) WithContext(ctx context.Context) *ClaimServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the claim server params
func (o *ClaimServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the claim server params
func (o *ClaimServerParams) WithHTTPClient(client *http.Client) *ClaimServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the claim server params
func (o *ClaimServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the claim server params
func (o *ClaimServerParams) WithBody(body *dsmcclientmodels.ModelsClaimSessionRequest) *ClaimServerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the claim server params
func (o *ClaimServerParams) SetBody(body *dsmcclientmodels.ModelsClaimSessionRequest) {
	o.Body = body
}

// WithNamespace adds the namespace to the claim server params
func (o *ClaimServerParams) WithNamespace(namespace string) *ClaimServerParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the claim server params
func (o *ClaimServerParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *ClaimServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
