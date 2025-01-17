// Code generated by go-swagger; DO NOT EDIT.

package wallet

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

// NewGetWalletParams creates a new GetWalletParams object
// with the default values initialized.
func NewGetWalletParams() *GetWalletParams {
	var ()
	return &GetWalletParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWalletParamsWithTimeout creates a new GetWalletParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWalletParamsWithTimeout(timeout time.Duration) *GetWalletParams {
	var ()
	return &GetWalletParams{

		timeout: timeout,
	}
}

// NewGetWalletParamsWithContext creates a new GetWalletParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWalletParamsWithContext(ctx context.Context) *GetWalletParams {
	var ()
	return &GetWalletParams{

		Context: ctx,
	}
}

// NewGetWalletParamsWithHTTPClient creates a new GetWalletParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWalletParamsWithHTTPClient(client *http.Client) *GetWalletParams {
	var ()
	return &GetWalletParams{
		HTTPClient: client,
	}
}

/*GetWalletParams contains all the parameters to send to the API endpoint
for the get wallet operation typically these are written to a http.Request
*/
type GetWalletParams struct {

	/*Namespace
	  Namespace

	*/
	Namespace string
	/*WalletID*/
	WalletID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get wallet params
func (o *GetWalletParams) WithTimeout(timeout time.Duration) *GetWalletParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get wallet params
func (o *GetWalletParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get wallet params
func (o *GetWalletParams) WithContext(ctx context.Context) *GetWalletParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get wallet params
func (o *GetWalletParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get wallet params
func (o *GetWalletParams) WithHTTPClient(client *http.Client) *GetWalletParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get wallet params
func (o *GetWalletParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get wallet params
func (o *GetWalletParams) WithNamespace(namespace string) *GetWalletParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get wallet params
func (o *GetWalletParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithWalletID adds the walletID to the get wallet params
func (o *GetWalletParams) WithWalletID(walletID string) *GetWalletParams {
	o.SetWalletID(walletID)
	return o
}

// SetWalletID adds the walletId to the get wallet params
func (o *GetWalletParams) SetWalletID(walletID string) {
	o.WalletID = walletID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWalletParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param walletId
	if err := r.SetPathParam("walletId", o.WalletID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
