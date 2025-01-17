// Code generated by go-swagger; DO NOT EDIT.

package users_v4

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

// NewAdminRemoveUserRoleV4Params creates a new AdminRemoveUserRoleV4Params object
// with the default values initialized.
func NewAdminRemoveUserRoleV4Params() *AdminRemoveUserRoleV4Params {
	var ()
	return &AdminRemoveUserRoleV4Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminRemoveUserRoleV4ParamsWithTimeout creates a new AdminRemoveUserRoleV4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminRemoveUserRoleV4ParamsWithTimeout(timeout time.Duration) *AdminRemoveUserRoleV4Params {
	var ()
	return &AdminRemoveUserRoleV4Params{

		timeout: timeout,
	}
}

// NewAdminRemoveUserRoleV4ParamsWithContext creates a new AdminRemoveUserRoleV4Params object
// with the default values initialized, and the ability to set a context for a request
func NewAdminRemoveUserRoleV4ParamsWithContext(ctx context.Context) *AdminRemoveUserRoleV4Params {
	var ()
	return &AdminRemoveUserRoleV4Params{

		Context: ctx,
	}
}

// NewAdminRemoveUserRoleV4ParamsWithHTTPClient creates a new AdminRemoveUserRoleV4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminRemoveUserRoleV4ParamsWithHTTPClient(client *http.Client) *AdminRemoveUserRoleV4Params {
	var ()
	return &AdminRemoveUserRoleV4Params{
		HTTPClient: client,
	}
}

/*AdminRemoveUserRoleV4Params contains all the parameters to send to the API endpoint
for the admin remove user role v4 operation typically these are written to a http.Request
*/
type AdminRemoveUserRoleV4Params struct {

	/*Body
	  Object of Role ID with allowed namespaces to be removed

	*/
	Body *iamclientmodels.ModelRemoveUserRoleV4Request
	/*Namespace
	  Namespace, only accept alphabet and numeric

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

// WithTimeout adds the timeout to the admin remove user role v4 params
func (o *AdminRemoveUserRoleV4Params) WithTimeout(timeout time.Duration) *AdminRemoveUserRoleV4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin remove user role v4 params
func (o *AdminRemoveUserRoleV4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin remove user role v4 params
func (o *AdminRemoveUserRoleV4Params) WithContext(ctx context.Context) *AdminRemoveUserRoleV4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin remove user role v4 params
func (o *AdminRemoveUserRoleV4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin remove user role v4 params
func (o *AdminRemoveUserRoleV4Params) WithHTTPClient(client *http.Client) *AdminRemoveUserRoleV4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin remove user role v4 params
func (o *AdminRemoveUserRoleV4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin remove user role v4 params
func (o *AdminRemoveUserRoleV4Params) WithBody(body *iamclientmodels.ModelRemoveUserRoleV4Request) *AdminRemoveUserRoleV4Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin remove user role v4 params
func (o *AdminRemoveUserRoleV4Params) SetBody(body *iamclientmodels.ModelRemoveUserRoleV4Request) {
	o.Body = body
}

// WithNamespace adds the namespace to the admin remove user role v4 params
func (o *AdminRemoveUserRoleV4Params) WithNamespace(namespace string) *AdminRemoveUserRoleV4Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the admin remove user role v4 params
func (o *AdminRemoveUserRoleV4Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the admin remove user role v4 params
func (o *AdminRemoveUserRoleV4Params) WithUserID(userID string) *AdminRemoveUserRoleV4Params {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the admin remove user role v4 params
func (o *AdminRemoveUserRoleV4Params) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AdminRemoveUserRoleV4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
