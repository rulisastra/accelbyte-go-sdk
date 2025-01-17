// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// NewCountServerDetailedParams creates a new CountServerDetailedParams object
// with the default values initialized.
func NewCountServerDetailedParams() *CountServerDetailedParams {
	var ()
	return &CountServerDetailedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCountServerDetailedParamsWithTimeout creates a new CountServerDetailedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCountServerDetailedParamsWithTimeout(timeout time.Duration) *CountServerDetailedParams {
	var ()
	return &CountServerDetailedParams{

		timeout: timeout,
	}
}

// NewCountServerDetailedParamsWithContext creates a new CountServerDetailedParams object
// with the default values initialized, and the ability to set a context for a request
func NewCountServerDetailedParamsWithContext(ctx context.Context) *CountServerDetailedParams {
	var ()
	return &CountServerDetailedParams{

		Context: ctx,
	}
}

// NewCountServerDetailedParamsWithHTTPClient creates a new CountServerDetailedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCountServerDetailedParamsWithHTTPClient(client *http.Client) *CountServerDetailedParams {
	var ()
	return &CountServerDetailedParams{
		HTTPClient: client,
	}
}

/*CountServerDetailedParams contains all the parameters to send to the API endpoint
for the count server detailed operation typically these are written to a http.Request
*/
type CountServerDetailedParams struct {

	/*Namespace
	  namespace of the game

	*/
	Namespace string
	/*Region
	  region where DS server is located.

	*/
	Region *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the count server detailed params
func (o *CountServerDetailedParams) WithTimeout(timeout time.Duration) *CountServerDetailedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the count server detailed params
func (o *CountServerDetailedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the count server detailed params
func (o *CountServerDetailedParams) WithContext(ctx context.Context) *CountServerDetailedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the count server detailed params
func (o *CountServerDetailedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the count server detailed params
func (o *CountServerDetailedParams) WithHTTPClient(client *http.Client) *CountServerDetailedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the count server detailed params
func (o *CountServerDetailedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the count server detailed params
func (o *CountServerDetailedParams) WithNamespace(namespace string) *CountServerDetailedParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the count server detailed params
func (o *CountServerDetailedParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithRegion adds the region to the count server detailed params
func (o *CountServerDetailedParams) WithRegion(region *string) *CountServerDetailedParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the count server detailed params
func (o *CountServerDetailedParams) SetRegion(region *string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *CountServerDetailedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.Region != nil {

		// query param region
		var qrRegion string
		if o.Region != nil {
			qrRegion = *o.Region
		}
		qRegion := qrRegion
		if qRegion != "" {
			if err := r.SetQueryParam("region", qRegion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
