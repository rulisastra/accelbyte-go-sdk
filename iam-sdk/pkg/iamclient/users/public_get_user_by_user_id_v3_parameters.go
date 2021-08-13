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

// NewPublicGetUserByUserIDV3Params creates a new PublicGetUserByUserIDV3Params object
// with the default values initialized.
func NewPublicGetUserByUserIDV3Params() *PublicGetUserByUserIDV3Params {
	var ()
	return &PublicGetUserByUserIDV3Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublicGetUserByUserIDV3ParamsWithTimeout creates a new PublicGetUserByUserIDV3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublicGetUserByUserIDV3ParamsWithTimeout(timeout time.Duration) *PublicGetUserByUserIDV3Params {
	var ()
	return &PublicGetUserByUserIDV3Params{

		timeout: timeout,
	}
}

// NewPublicGetUserByUserIDV3ParamsWithContext creates a new PublicGetUserByUserIDV3Params object
// with the default values initialized, and the ability to set a context for a request
func NewPublicGetUserByUserIDV3ParamsWithContext(ctx context.Context) *PublicGetUserByUserIDV3Params {
	var ()
	return &PublicGetUserByUserIDV3Params{

		Context: ctx,
	}
}

// NewPublicGetUserByUserIDV3ParamsWithHTTPClient creates a new PublicGetUserByUserIDV3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublicGetUserByUserIDV3ParamsWithHTTPClient(client *http.Client) *PublicGetUserByUserIDV3Params {
	var ()
	return &PublicGetUserByUserIDV3Params{
		HTTPClient: client,
	}
}

/*PublicGetUserByUserIDV3Params contains all the parameters to send to the API endpoint
for the public get user by user Id v3 operation typically these are written to a http.Request
*/
type PublicGetUserByUserIDV3Params struct {

	/*Namespace
	  Namespace, only accept alphabet and numeric

	*/
	Namespace string
	/*UserID
	  User ID, should follow UUID version 4 without hyphen

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the public get user by user Id v3 params
func (o *PublicGetUserByUserIDV3Params) WithTimeout(timeout time.Duration) *PublicGetUserByUserIDV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the public get user by user Id v3 params
func (o *PublicGetUserByUserIDV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the public get user by user Id v3 params
func (o *PublicGetUserByUserIDV3Params) WithContext(ctx context.Context) *PublicGetUserByUserIDV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the public get user by user Id v3 params
func (o *PublicGetUserByUserIDV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the public get user by user Id v3 params
func (o *PublicGetUserByUserIDV3Params) WithHTTPClient(client *http.Client) *PublicGetUserByUserIDV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the public get user by user Id v3 params
func (o *PublicGetUserByUserIDV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the public get user by user Id v3 params
func (o *PublicGetUserByUserIDV3Params) WithNamespace(namespace string) *PublicGetUserByUserIDV3Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the public get user by user Id v3 params
func (o *PublicGetUserByUserIDV3Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the public get user by user Id v3 params
func (o *PublicGetUserByUserIDV3Params) WithUserID(userID string) *PublicGetUserByUserIDV3Params {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the public get user by user Id v3 params
func (o *PublicGetUserByUserIDV3Params) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PublicGetUserByUserIDV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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