// Code generated by go-swagger; DO NOT EDIT.

package user_statistic

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

	"github.com/AccelByte/accelbyte-go-sdk/social-sdk/pkg/socialclientmodels"
)

// NewPublicBulkIncUserStatItemValueParams creates a new PublicBulkIncUserStatItemValueParams object
// with the default values initialized.
func NewPublicBulkIncUserStatItemValueParams() *PublicBulkIncUserStatItemValueParams {
	var ()
	return &PublicBulkIncUserStatItemValueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicBulkIncUserStatItemValueParamsWithTimeout creates a new PublicBulkIncUserStatItemValueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicBulkIncUserStatItemValueParamsWithTimeout(timeout time.Duration) *PublicBulkIncUserStatItemValueParams {
	var ()
	return &PublicBulkIncUserStatItemValueParams{

		timeout: timeout,
	}
}

// NewPublicBulkIncUserStatItemValueParamsWithContext creates a new PublicBulkIncUserStatItemValueParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublicBulkIncUserStatItemValueParamsWithContext(ctx context.Context) *PublicBulkIncUserStatItemValueParams {
	var ()
	return &PublicBulkIncUserStatItemValueParams{

		Context: ctx,
	}
}

// NewPublicBulkIncUserStatItemValueParamsWithHTTPClient creates a new PublicBulkIncUserStatItemValueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicBulkIncUserStatItemValueParamsWithHTTPClient(client *http.Client) *PublicBulkIncUserStatItemValueParams {
	var ()
	return &PublicBulkIncUserStatItemValueParams{
		HTTPClient: client,
	}
}

/*PublicBulkIncUserStatItemValueParams contains all the parameters to send to the API endpoint
for the public bulk inc user stat item value operation typically these are written to a http.Request
*/
type PublicBulkIncUserStatItemValueParams struct {

	/*Body*/
	Body []*socialclientmodels.BulkUserStatItemInc
	/*Namespace
	  namespace

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the public bulk inc user stat item value params
func (o *PublicBulkIncUserStatItemValueParams) WithTimeout(timeout time.Duration) *PublicBulkIncUserStatItemValueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public bulk inc user stat item value params
func (o *PublicBulkIncUserStatItemValueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public bulk inc user stat item value params
func (o *PublicBulkIncUserStatItemValueParams) WithContext(ctx context.Context) *PublicBulkIncUserStatItemValueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public bulk inc user stat item value params
func (o *PublicBulkIncUserStatItemValueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public bulk inc user stat item value params
func (o *PublicBulkIncUserStatItemValueParams) WithHTTPClient(client *http.Client) *PublicBulkIncUserStatItemValueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public bulk inc user stat item value params
func (o *PublicBulkIncUserStatItemValueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the public bulk inc user stat item value params
func (o *PublicBulkIncUserStatItemValueParams) WithBody(body []*socialclientmodels.BulkUserStatItemInc) *PublicBulkIncUserStatItemValueParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the public bulk inc user stat item value params
func (o *PublicBulkIncUserStatItemValueParams) SetBody(body []*socialclientmodels.BulkUserStatItemInc) {
	o.Body = body
}

// WithNamespace adds the namespace to the public bulk inc user stat item value params
func (o *PublicBulkIncUserStatItemValueParams) WithNamespace(namespace string) *PublicBulkIncUserStatItemValueParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public bulk inc user stat item value params
func (o *PublicBulkIncUserStatItemValueParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *PublicBulkIncUserStatItemValueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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