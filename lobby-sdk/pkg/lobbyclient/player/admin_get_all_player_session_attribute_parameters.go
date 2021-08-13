// Code generated by go-swagger; DO NOT EDIT.

package player

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

// NewAdminGetAllPlayerSessionAttributeParams creates a new AdminGetAllPlayerSessionAttributeParams object
// with the default values initialized.
func NewAdminGetAllPlayerSessionAttributeParams() *AdminGetAllPlayerSessionAttributeParams {
	var ()
	return &AdminGetAllPlayerSessionAttributeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminGetAllPlayerSessionAttributeParamsWithTimeout creates a new AdminGetAllPlayerSessionAttributeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminGetAllPlayerSessionAttributeParamsWithTimeout(timeout time.Duration) *AdminGetAllPlayerSessionAttributeParams {
	var ()
	return &AdminGetAllPlayerSessionAttributeParams{

		timeout: timeout,
	}
}

// NewAdminGetAllPlayerSessionAttributeParamsWithContext creates a new AdminGetAllPlayerSessionAttributeParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdminGetAllPlayerSessionAttributeParamsWithContext(ctx context.Context) *AdminGetAllPlayerSessionAttributeParams {
	var ()
	return &AdminGetAllPlayerSessionAttributeParams{

		Context: ctx,
	}
}

// NewAdminGetAllPlayerSessionAttributeParamsWithHTTPClient creates a new AdminGetAllPlayerSessionAttributeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminGetAllPlayerSessionAttributeParamsWithHTTPClient(client *http.Client) *AdminGetAllPlayerSessionAttributeParams {
	var ()
	return &AdminGetAllPlayerSessionAttributeParams{
		HTTPClient: client,
	}
}

/*AdminGetAllPlayerSessionAttributeParams contains all the parameters to send to the API endpoint
for the admin get all player session attribute operation typically these are written to a http.Request
*/
type AdminGetAllPlayerSessionAttributeParams struct {

	/*Namespace
	  namespace

	*/
	Namespace string
	/*UserID
	  user ID

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin get all player session attribute params
func (o *AdminGetAllPlayerSessionAttributeParams) WithTimeout(timeout time.Duration) *AdminGetAllPlayerSessionAttributeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin get all player session attribute params
func (o *AdminGetAllPlayerSessionAttributeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin get all player session attribute params
func (o *AdminGetAllPlayerSessionAttributeParams) WithContext(ctx context.Context) *AdminGetAllPlayerSessionAttributeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin get all player session attribute params
func (o *AdminGetAllPlayerSessionAttributeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin get all player session attribute params
func (o *AdminGetAllPlayerSessionAttributeParams) WithHTTPClient(client *http.Client) *AdminGetAllPlayerSessionAttributeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin get all player session attribute params
func (o *AdminGetAllPlayerSessionAttributeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the admin get all player session attribute params
func (o *AdminGetAllPlayerSessionAttributeParams) WithNamespace(namespace string) *AdminGetAllPlayerSessionAttributeParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin get all player session attribute params
func (o *AdminGetAllPlayerSessionAttributeParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the admin get all player session attribute params
func (o *AdminGetAllPlayerSessionAttributeParams) WithUserID(userID string) *AdminGetAllPlayerSessionAttributeParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the admin get all player session attribute params
func (o *AdminGetAllPlayerSessionAttributeParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AdminGetAllPlayerSessionAttributeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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