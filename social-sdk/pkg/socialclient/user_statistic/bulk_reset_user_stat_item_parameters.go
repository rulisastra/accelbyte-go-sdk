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

// NewBulkResetUserStatItemParams creates a new BulkResetUserStatItemParams object
// with the default values initialized.
func NewBulkResetUserStatItemParams() *BulkResetUserStatItemParams {
	var ()
	return &BulkResetUserStatItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBulkResetUserStatItemParamsWithTimeout creates a new BulkResetUserStatItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBulkResetUserStatItemParamsWithTimeout(timeout time.Duration) *BulkResetUserStatItemParams {
	var ()
	return &BulkResetUserStatItemParams{

		timeout: timeout,
	}
}

// NewBulkResetUserStatItemParamsWithContext creates a new BulkResetUserStatItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewBulkResetUserStatItemParamsWithContext(ctx context.Context) *BulkResetUserStatItemParams {
	var ()
	return &BulkResetUserStatItemParams{

		Context: ctx,
	}
}

// NewBulkResetUserStatItemParamsWithHTTPClient creates a new BulkResetUserStatItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBulkResetUserStatItemParamsWithHTTPClient(client *http.Client) *BulkResetUserStatItemParams {
	var ()
	return &BulkResetUserStatItemParams{
		HTTPClient: client,
	}
}

/*BulkResetUserStatItemParams contains all the parameters to send to the API endpoint
for the bulk reset user stat item operation typically these are written to a http.Request
*/
type BulkResetUserStatItemParams struct {

	/*Body*/
	Body []*socialclientmodels.BulkUserStatItemReset
	/*Namespace
	  namespace

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the bulk reset user stat item params
func (o *BulkResetUserStatItemParams) WithTimeout(timeout time.Duration) *BulkResetUserStatItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bulk reset user stat item params
func (o *BulkResetUserStatItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bulk reset user stat item params
func (o *BulkResetUserStatItemParams) WithContext(ctx context.Context) *BulkResetUserStatItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bulk reset user stat item params
func (o *BulkResetUserStatItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bulk reset user stat item params
func (o *BulkResetUserStatItemParams) WithHTTPClient(client *http.Client) *BulkResetUserStatItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bulk reset user stat item params
func (o *BulkResetUserStatItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the bulk reset user stat item params
func (o *BulkResetUserStatItemParams) WithBody(body []*socialclientmodels.BulkUserStatItemReset) *BulkResetUserStatItemParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the bulk reset user stat item params
func (o *BulkResetUserStatItemParams) SetBody(body []*socialclientmodels.BulkUserStatItemReset) {
	o.Body = body
}

// WithNamespace adds the namespace to the bulk reset user stat item params
func (o *BulkResetUserStatItemParams) WithNamespace(namespace string) *BulkResetUserStatItemParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the bulk reset user stat item params
func (o *BulkResetUserStatItemParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *BulkResetUserStatItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
