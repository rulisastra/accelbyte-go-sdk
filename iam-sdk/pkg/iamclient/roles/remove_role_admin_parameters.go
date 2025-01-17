// Code generated by go-swagger; DO NOT EDIT.

package roles

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

// NewRemoveRoleAdminParams creates a new RemoveRoleAdminParams object
// with the default values initialized.
func NewRemoveRoleAdminParams() *RemoveRoleAdminParams {
	var ()
	return &RemoveRoleAdminParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveRoleAdminParamsWithTimeout creates a new RemoveRoleAdminParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveRoleAdminParamsWithTimeout(timeout time.Duration) *RemoveRoleAdminParams {
	var ()
	return &RemoveRoleAdminParams{

		timeout: timeout,
	}
}

// NewRemoveRoleAdminParamsWithContext creates a new RemoveRoleAdminParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveRoleAdminParamsWithContext(ctx context.Context) *RemoveRoleAdminParams {
	var ()
	return &RemoveRoleAdminParams{

		Context: ctx,
	}
}

// NewRemoveRoleAdminParamsWithHTTPClient creates a new RemoveRoleAdminParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveRoleAdminParamsWithHTTPClient(client *http.Client) *RemoveRoleAdminParams {
	var ()
	return &RemoveRoleAdminParams{
		HTTPClient: client,
	}
}

/*RemoveRoleAdminParams contains all the parameters to send to the API endpoint
for the remove role admin operation typically these are written to a http.Request
*/
type RemoveRoleAdminParams struct {

	/*RoleID
	  Role id

	*/
	RoleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove role admin params
func (o *RemoveRoleAdminParams) WithTimeout(timeout time.Duration) *RemoveRoleAdminParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove role admin params
func (o *RemoveRoleAdminParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove role admin params
func (o *RemoveRoleAdminParams) WithContext(ctx context.Context) *RemoveRoleAdminParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove role admin params
func (o *RemoveRoleAdminParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove role admin params
func (o *RemoveRoleAdminParams) WithHTTPClient(client *http.Client) *RemoveRoleAdminParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove role admin params
func (o *RemoveRoleAdminParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoleID adds the roleID to the remove role admin params
func (o *RemoveRoleAdminParams) WithRoleID(roleID string) *RemoveRoleAdminParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the remove role admin params
func (o *RemoveRoleAdminParams) SetRoleID(roleID string) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveRoleAdminParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param roleId
	if err := r.SetPathParam("roleId", o.RoleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
