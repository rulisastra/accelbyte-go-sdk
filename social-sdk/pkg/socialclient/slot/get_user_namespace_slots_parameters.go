// Code generated by go-swagger; DO NOT EDIT.

package slot

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

// NewGetUserNamespaceSlotsParams creates a new GetUserNamespaceSlotsParams object
// with the default values initialized.
func NewGetUserNamespaceSlotsParams() *GetUserNamespaceSlotsParams {
	var ()
	return &GetUserNamespaceSlotsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserNamespaceSlotsParamsWithTimeout creates a new GetUserNamespaceSlotsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserNamespaceSlotsParamsWithTimeout(timeout time.Duration) *GetUserNamespaceSlotsParams {
	var ()
	return &GetUserNamespaceSlotsParams{

		timeout: timeout,
	}
}

// NewGetUserNamespaceSlotsParamsWithContext creates a new GetUserNamespaceSlotsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserNamespaceSlotsParamsWithContext(ctx context.Context) *GetUserNamespaceSlotsParams {
	var ()
	return &GetUserNamespaceSlotsParams{

		Context: ctx,
	}
}

// NewGetUserNamespaceSlotsParamsWithHTTPClient creates a new GetUserNamespaceSlotsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserNamespaceSlotsParamsWithHTTPClient(client *http.Client) *GetUserNamespaceSlotsParams {
	var ()
	return &GetUserNamespaceSlotsParams{
		HTTPClient: client,
	}
}

/*GetUserNamespaceSlotsParams contains all the parameters to send to the API endpoint
for the get user namespace slots operation typically these are written to a http.Request
*/
type GetUserNamespaceSlotsParams struct {

	/*Namespace
	  Namespace ID

	*/
	Namespace string
	/*UserID
	  User ID

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user namespace slots params
func (o *GetUserNamespaceSlotsParams) WithTimeout(timeout time.Duration) *GetUserNamespaceSlotsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user namespace slots params
func (o *GetUserNamespaceSlotsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user namespace slots params
func (o *GetUserNamespaceSlotsParams) WithContext(ctx context.Context) *GetUserNamespaceSlotsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user namespace slots params
func (o *GetUserNamespaceSlotsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user namespace slots params
func (o *GetUserNamespaceSlotsParams) WithHTTPClient(client *http.Client) *GetUserNamespaceSlotsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user namespace slots params
func (o *GetUserNamespaceSlotsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get user namespace slots params
func (o *GetUserNamespaceSlotsParams) WithNamespace(namespace string) *GetUserNamespaceSlotsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get user namespace slots params
func (o *GetUserNamespaceSlotsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the get user namespace slots params
func (o *GetUserNamespaceSlotsParams) WithUserID(userID string) *GetUserNamespaceSlotsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get user namespace slots params
func (o *GetUserNamespaceSlotsParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserNamespaceSlotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
