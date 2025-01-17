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
)

// NewGetListCountryAgeRestrictionParams creates a new GetListCountryAgeRestrictionParams object
// with the default values initialized.
func NewGetListCountryAgeRestrictionParams() *GetListCountryAgeRestrictionParams {
	var ()
	return &GetListCountryAgeRestrictionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetListCountryAgeRestrictionParamsWithTimeout creates a new GetListCountryAgeRestrictionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetListCountryAgeRestrictionParamsWithTimeout(timeout time.Duration) *GetListCountryAgeRestrictionParams {
	var ()
	return &GetListCountryAgeRestrictionParams{

		timeout: timeout,
	}
}

// NewGetListCountryAgeRestrictionParamsWithContext creates a new GetListCountryAgeRestrictionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetListCountryAgeRestrictionParamsWithContext(ctx context.Context) *GetListCountryAgeRestrictionParams {
	var ()
	return &GetListCountryAgeRestrictionParams{

		Context: ctx,
	}
}

// NewGetListCountryAgeRestrictionParamsWithHTTPClient creates a new GetListCountryAgeRestrictionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetListCountryAgeRestrictionParamsWithHTTPClient(client *http.Client) *GetListCountryAgeRestrictionParams {
	var ()
	return &GetListCountryAgeRestrictionParams{
		HTTPClient: client,
	}
}

/*GetListCountryAgeRestrictionParams contains all the parameters to send to the API endpoint
for the get list country age restriction operation typically these are written to a http.Request
*/
type GetListCountryAgeRestrictionParams struct {

	/*Namespace
	  Namespace

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get list country age restriction params
func (o *GetListCountryAgeRestrictionParams) WithTimeout(timeout time.Duration) *GetListCountryAgeRestrictionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get list country age restriction params
func (o *GetListCountryAgeRestrictionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get list country age restriction params
func (o *GetListCountryAgeRestrictionParams) WithContext(ctx context.Context) *GetListCountryAgeRestrictionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get list country age restriction params
func (o *GetListCountryAgeRestrictionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get list country age restriction params
func (o *GetListCountryAgeRestrictionParams) WithHTTPClient(client *http.Client) *GetListCountryAgeRestrictionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get list country age restriction params
func (o *GetListCountryAgeRestrictionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get list country age restriction params
func (o *GetListCountryAgeRestrictionParams) WithNamespace(namespace string) *GetListCountryAgeRestrictionParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get list country age restriction params
func (o *GetListCountryAgeRestrictionParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *GetListCountryAgeRestrictionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
