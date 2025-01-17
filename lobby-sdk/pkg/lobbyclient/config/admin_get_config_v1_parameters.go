// Code generated by go-swagger; DO NOT EDIT.

package config

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
)

// NewAdminGetConfigV1Params creates a new AdminGetConfigV1Params object
// with the default values initialized.
func NewAdminGetConfigV1Params() *AdminGetConfigV1Params {
	var ()
	return &AdminGetConfigV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminGetConfigV1ParamsWithTimeout creates a new AdminGetConfigV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminGetConfigV1ParamsWithTimeout(timeout time.Duration) *AdminGetConfigV1Params {
	var ()
	return &AdminGetConfigV1Params{

		timeout: timeout,
	}
}

// NewAdminGetConfigV1ParamsWithContext creates a new AdminGetConfigV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewAdminGetConfigV1ParamsWithContext(ctx context.Context) *AdminGetConfigV1Params {
	var ()
	return &AdminGetConfigV1Params{

		Context: ctx,
	}
}

// NewAdminGetConfigV1ParamsWithHTTPClient creates a new AdminGetConfigV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminGetConfigV1ParamsWithHTTPClient(client *http.Client) *AdminGetConfigV1Params {
	var ()
	return &AdminGetConfigV1Params{
		HTTPClient: client,
	}
}

/*AdminGetConfigV1Params contains all the parameters to send to the API endpoint
for the admin get config v1 operation typically these are written to a http.Request
*/
type AdminGetConfigV1Params struct {

	/*Namespace
	  namespace

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin get config v1 params
func (o *AdminGetConfigV1Params) WithTimeout(timeout time.Duration) *AdminGetConfigV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin get config v1 params
func (o *AdminGetConfigV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin get config v1 params
func (o *AdminGetConfigV1Params) WithContext(ctx context.Context) *AdminGetConfigV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin get config v1 params
func (o *AdminGetConfigV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin get config v1 params
func (o *AdminGetConfigV1Params) WithHTTPClient(client *http.Client) *AdminGetConfigV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin get config v1 params
func (o *AdminGetConfigV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the admin get config v1 params
func (o *AdminGetConfigV1Params) WithNamespace(namespace string) *AdminGetConfigV1Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin get config v1 params
func (o *AdminGetConfigV1Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *AdminGetConfigV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
