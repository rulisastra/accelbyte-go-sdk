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

	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclientmodels"
)

// NewAdminAddRoleMembersV3Params creates a new AdminAddRoleMembersV3Params object
// with the default values initialized.
func NewAdminAddRoleMembersV3Params() *AdminAddRoleMembersV3Params {
	var ()
	return &AdminAddRoleMembersV3Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminAddRoleMembersV3ParamsWithTimeout creates a new AdminAddRoleMembersV3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminAddRoleMembersV3ParamsWithTimeout(timeout time.Duration) *AdminAddRoleMembersV3Params {
	var ()
	return &AdminAddRoleMembersV3Params{

		timeout: timeout,
	}
}

// NewAdminAddRoleMembersV3ParamsWithContext creates a new AdminAddRoleMembersV3Params object
// with the default values initialized, and the ability to set a context for a request
func NewAdminAddRoleMembersV3ParamsWithContext(ctx context.Context) *AdminAddRoleMembersV3Params {
	var ()
	return &AdminAddRoleMembersV3Params{

		Context: ctx,
	}
}

// NewAdminAddRoleMembersV3ParamsWithHTTPClient creates a new AdminAddRoleMembersV3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminAddRoleMembersV3ParamsWithHTTPClient(client *http.Client) *AdminAddRoleMembersV3Params {
	var ()
	return &AdminAddRoleMembersV3Params{
		HTTPClient: client,
	}
}

/*AdminAddRoleMembersV3Params contains all the parameters to send to the API endpoint
for the admin add role members v3 operation typically these are written to a http.Request
*/
type AdminAddRoleMembersV3Params struct {

	/*Body*/
	Body *iamclientmodels.ModelRoleMembersRequestV3
	/*RoleID
	  Role id

	*/
	RoleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin add role members v3 params
func (o *AdminAddRoleMembersV3Params) WithTimeout(timeout time.Duration) *AdminAddRoleMembersV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin add role members v3 params
func (o *AdminAddRoleMembersV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin add role members v3 params
func (o *AdminAddRoleMembersV3Params) WithContext(ctx context.Context) *AdminAddRoleMembersV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin add role members v3 params
func (o *AdminAddRoleMembersV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin add role members v3 params
func (o *AdminAddRoleMembersV3Params) WithHTTPClient(client *http.Client) *AdminAddRoleMembersV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin add role members v3 params
func (o *AdminAddRoleMembersV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin add role members v3 params
func (o *AdminAddRoleMembersV3Params) WithBody(body *iamclientmodels.ModelRoleMembersRequestV3) *AdminAddRoleMembersV3Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin add role members v3 params
func (o *AdminAddRoleMembersV3Params) SetBody(body *iamclientmodels.ModelRoleMembersRequestV3) {
	o.Body = body
}

// WithRoleID adds the roleID to the admin add role members v3 params
func (o *AdminAddRoleMembersV3Params) WithRoleID(roleID string) *AdminAddRoleMembersV3Params {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the admin add role members v3 params
func (o *AdminAddRoleMembersV3Params) SetRoleID(roleID string) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *AdminAddRoleMembersV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param roleId
	if err := r.SetPathParam("roleId", o.RoleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}