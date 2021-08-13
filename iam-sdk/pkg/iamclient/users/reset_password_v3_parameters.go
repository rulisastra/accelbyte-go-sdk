// Code generated by go-swagger; DO NOT EDIT.

package users

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

	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclientmodels"
)

// NewResetPasswordV3Params creates a new ResetPasswordV3Params object
// with the default values initialized.
func NewResetPasswordV3Params() *ResetPasswordV3Params {
	var ()
	return &ResetPasswordV3Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewResetPasswordV3ParamsWithTimeout creates a new ResetPasswordV3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewResetPasswordV3ParamsWithTimeout(timeout time.Duration) *ResetPasswordV3Params {
	var ()
	return &ResetPasswordV3Params{

		timeout: timeout,
	}
}

// NewResetPasswordV3ParamsWithContext creates a new ResetPasswordV3Params object
// with the default values initialized, and the ability to set a context for a request
func NewResetPasswordV3ParamsWithContext(ctx context.Context) *ResetPasswordV3Params {
	var ()
	return &ResetPasswordV3Params{

		Context: ctx,
	}
}

// NewResetPasswordV3ParamsWithHTTPClient creates a new ResetPasswordV3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewResetPasswordV3ParamsWithHTTPClient(client *http.Client) *ResetPasswordV3Params {
	var ()
	return &ResetPasswordV3Params{
		HTTPClient: client,
	}
}

/*ResetPasswordV3Params contains all the parameters to send to the API endpoint
for the reset password v3 operation typically these are written to a http.Request
*/
type ResetPasswordV3Params struct {

	/*Body*/
	Body *iamclientmodels.ModelResetPasswordRequestV3
	/*Namespace
	  Namespace, only accept alphabet and numeric

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the reset password v3 params
func (o *ResetPasswordV3Params) WithTimeout(timeout time.Duration) *ResetPasswordV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reset password v3 params
func (o *ResetPasswordV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reset password v3 params
func (o *ResetPasswordV3Params) WithContext(ctx context.Context) *ResetPasswordV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reset password v3 params
func (o *ResetPasswordV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reset password v3 params
func (o *ResetPasswordV3Params) WithHTTPClient(client *http.Client) *ResetPasswordV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reset password v3 params
func (o *ResetPasswordV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the reset password v3 params
func (o *ResetPasswordV3Params) WithBody(body *iamclientmodels.ModelResetPasswordRequestV3) *ResetPasswordV3Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the reset password v3 params
func (o *ResetPasswordV3Params) SetBody(body *iamclientmodels.ModelResetPasswordRequestV3) {
	o.Body = body
}

// WithNamespace adds the namespace to the reset password v3 params
func (o *ResetPasswordV3Params) WithNamespace(namespace string) *ResetPasswordV3Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the reset password v3 params
func (o *ResetPasswordV3Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *ResetPasswordV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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