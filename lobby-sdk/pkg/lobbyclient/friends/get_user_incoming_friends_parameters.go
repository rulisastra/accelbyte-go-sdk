// Code generated by go-swagger; DO NOT EDIT.

package friends

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

// NewGetUserIncomingFriendsParams creates a new GetUserIncomingFriendsParams object
// with the default values initialized.
func NewGetUserIncomingFriendsParams() *GetUserIncomingFriendsParams {
	var ()
	return &GetUserIncomingFriendsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserIncomingFriendsParamsWithTimeout creates a new GetUserIncomingFriendsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserIncomingFriendsParamsWithTimeout(timeout time.Duration) *GetUserIncomingFriendsParams {
	var ()
	return &GetUserIncomingFriendsParams{

		timeout: timeout,
	}
}

// NewGetUserIncomingFriendsParamsWithContext creates a new GetUserIncomingFriendsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserIncomingFriendsParamsWithContext(ctx context.Context) *GetUserIncomingFriendsParams {
	var ()
	return &GetUserIncomingFriendsParams{

		Context: ctx,
	}
}

// NewGetUserIncomingFriendsParamsWithHTTPClient creates a new GetUserIncomingFriendsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserIncomingFriendsParamsWithHTTPClient(client *http.Client) *GetUserIncomingFriendsParams {
	var ()
	return &GetUserIncomingFriendsParams{
		HTTPClient: client,
	}
}

/*GetUserIncomingFriendsParams contains all the parameters to send to the API endpoint
for the get user incoming friends operation typically these are written to a http.Request
*/
type GetUserIncomingFriendsParams struct {

	/*Namespace
	  namespace

	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user incoming friends params
func (o *GetUserIncomingFriendsParams) WithTimeout(timeout time.Duration) *GetUserIncomingFriendsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user incoming friends params
func (o *GetUserIncomingFriendsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user incoming friends params
func (o *GetUserIncomingFriendsParams) WithContext(ctx context.Context) *GetUserIncomingFriendsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user incoming friends params
func (o *GetUserIncomingFriendsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user incoming friends params
func (o *GetUserIncomingFriendsParams) WithHTTPClient(client *http.Client) *GetUserIncomingFriendsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user incoming friends params
func (o *GetUserIncomingFriendsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get user incoming friends params
func (o *GetUserIncomingFriendsParams) WithNamespace(namespace string) *GetUserIncomingFriendsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get user incoming friends params
func (o *GetUserIncomingFriendsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserIncomingFriendsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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