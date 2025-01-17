// Code generated by go-swagger; DO NOT EDIT.

package slot_config

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

// NewGetNamespaceSlotConfigParams creates a new GetNamespaceSlotConfigParams object
// with the default values initialized.
func NewGetNamespaceSlotConfigParams() *GetNamespaceSlotConfigParams {
	var ()
	return &GetNamespaceSlotConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNamespaceSlotConfigParamsWithTimeout creates a new GetNamespaceSlotConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNamespaceSlotConfigParamsWithTimeout(timeout time.Duration) *GetNamespaceSlotConfigParams {
	var ()
	return &GetNamespaceSlotConfigParams{

		timeout: timeout,
	}
}

// NewGetNamespaceSlotConfigParamsWithContext creates a new GetNamespaceSlotConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNamespaceSlotConfigParamsWithContext(ctx context.Context) *GetNamespaceSlotConfigParams {
	var ()
	return &GetNamespaceSlotConfigParams{

		Context: ctx,
	}
}

// NewGetNamespaceSlotConfigParamsWithHTTPClient creates a new GetNamespaceSlotConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNamespaceSlotConfigParamsWithHTTPClient(client *http.Client) *GetNamespaceSlotConfigParams {
	var ()
	return &GetNamespaceSlotConfigParams{
		HTTPClient: client,
	}
}

/*GetNamespaceSlotConfigParams contains all the parameters to send to the API endpoint
for the get namespace slot config operation typically these are written to a http.Request
*/
type GetNamespaceSlotConfigParams struct {

	/*Namespace
	  Namespace

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get namespace slot config params
func (o *GetNamespaceSlotConfigParams) WithTimeout(timeout time.Duration) *GetNamespaceSlotConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get namespace slot config params
func (o *GetNamespaceSlotConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get namespace slot config params
func (o *GetNamespaceSlotConfigParams) WithContext(ctx context.Context) *GetNamespaceSlotConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get namespace slot config params
func (o *GetNamespaceSlotConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get namespace slot config params
func (o *GetNamespaceSlotConfigParams) WithHTTPClient(client *http.Client) *GetNamespaceSlotConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get namespace slot config params
func (o *GetNamespaceSlotConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get namespace slot config params
func (o *GetNamespaceSlotConfigParams) WithNamespace(namespace string) *GetNamespaceSlotConfigParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get namespace slot config params
func (o *GetNamespaceSlotConfigParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *GetNamespaceSlotConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
