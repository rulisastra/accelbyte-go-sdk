// Code generated by go-swagger; DO NOT EDIT.

package o_auth2_0_extension

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

// NewGetCountryLocationV3Params creates a new GetCountryLocationV3Params object
// with the default values initialized.
func NewGetCountryLocationV3Params() *GetCountryLocationV3Params {

	return &GetCountryLocationV3Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCountryLocationV3ParamsWithTimeout creates a new GetCountryLocationV3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCountryLocationV3ParamsWithTimeout(timeout time.Duration) *GetCountryLocationV3Params {

	return &GetCountryLocationV3Params{

		timeout: timeout,
	}
}

// NewGetCountryLocationV3ParamsWithContext creates a new GetCountryLocationV3Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetCountryLocationV3ParamsWithContext(ctx context.Context) *GetCountryLocationV3Params {

	return &GetCountryLocationV3Params{

		Context: ctx,
	}
}

// NewGetCountryLocationV3ParamsWithHTTPClient creates a new GetCountryLocationV3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCountryLocationV3ParamsWithHTTPClient(client *http.Client) *GetCountryLocationV3Params {

	return &GetCountryLocationV3Params{
		HTTPClient: client,
	}
}

/*GetCountryLocationV3Params contains all the parameters to send to the API endpoint
for the get country location v3 operation typically these are written to a http.Request
*/
type GetCountryLocationV3Params struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get country location v3 params
func (o *GetCountryLocationV3Params) WithTimeout(timeout time.Duration) *GetCountryLocationV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get country location v3 params
func (o *GetCountryLocationV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get country location v3 params
func (o *GetCountryLocationV3Params) WithContext(ctx context.Context) *GetCountryLocationV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get country location v3 params
func (o *GetCountryLocationV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get country location v3 params
func (o *GetCountryLocationV3Params) WithHTTPClient(client *http.Client) *GetCountryLocationV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get country location v3 params
func (o *GetCountryLocationV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetCountryLocationV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
