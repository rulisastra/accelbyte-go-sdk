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
	"github.com/go-openapi/swag"
)

// NewAdminGetRolesV4Params creates a new AdminGetRolesV4Params object
// with the default values initialized.
func NewAdminGetRolesV4Params() *AdminGetRolesV4Params {
	var ()
	return &AdminGetRolesV4Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminGetRolesV4ParamsWithTimeout creates a new AdminGetRolesV4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminGetRolesV4ParamsWithTimeout(timeout time.Duration) *AdminGetRolesV4Params {
	var ()
	return &AdminGetRolesV4Params{

		timeout: timeout,
	}
}

// NewAdminGetRolesV4ParamsWithContext creates a new AdminGetRolesV4Params object
// with the default values initialized, and the ability to set a context for a request
func NewAdminGetRolesV4ParamsWithContext(ctx context.Context) *AdminGetRolesV4Params {
	var ()
	return &AdminGetRolesV4Params{

		Context: ctx,
	}
}

// NewAdminGetRolesV4ParamsWithHTTPClient creates a new AdminGetRolesV4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminGetRolesV4ParamsWithHTTPClient(client *http.Client) *AdminGetRolesV4Params {
	var ()
	return &AdminGetRolesV4Params{
		HTTPClient: client,
	}
}

/*AdminGetRolesV4Params contains all the parameters to send to the API endpoint
for the admin get roles v4 operation typically these are written to a http.Request
*/
type AdminGetRolesV4Params struct {

	/*AdminRole
	  - true if the expected result should only returns records with adminRole = true
	          - false if the expected result should only returns records with adminRole = false
	          - empty (omitted) if the expected result should returns records with no wildcard filter at all

	*/
	AdminRole *bool
	/*After
	  The cursor that points to query data for the next page

	*/
	After *string
	/*Before
	  The cursor that points to query data for the previous page

	*/
	Before *string
	/*IsWildcard
	  - true if the expected result should only returns records with wildcard = true
	          - false if the expected result should only returns records with wildcard = false
	          - empty (omitted) if the expected result should returns records with no wildcard filter at all

	*/
	IsWildcard *bool
	/*Limit
	  the maximum number of data that may be returned (1...100)

	*/
	Limit *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin get roles v4 params
func (o *AdminGetRolesV4Params) WithTimeout(timeout time.Duration) *AdminGetRolesV4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin get roles v4 params
func (o *AdminGetRolesV4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin get roles v4 params
func (o *AdminGetRolesV4Params) WithContext(ctx context.Context) *AdminGetRolesV4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin get roles v4 params
func (o *AdminGetRolesV4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin get roles v4 params
func (o *AdminGetRolesV4Params) WithHTTPClient(client *http.Client) *AdminGetRolesV4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin get roles v4 params
func (o *AdminGetRolesV4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdminRole adds the adminRole to the admin get roles v4 params
func (o *AdminGetRolesV4Params) WithAdminRole(adminRole *bool) *AdminGetRolesV4Params {
	o.SetAdminRole(adminRole)
	return o
}

// SetAdminRole adds the adminRole to the admin get roles v4 params
func (o *AdminGetRolesV4Params) SetAdminRole(adminRole *bool) {
	o.AdminRole = adminRole
}

// WithAfter adds the after to the admin get roles v4 params
func (o *AdminGetRolesV4Params) WithAfter(after *string) *AdminGetRolesV4Params {
	o.SetAfter(after)
	return o
}

// SetAfter adds the after to the admin get roles v4 params
func (o *AdminGetRolesV4Params) SetAfter(after *string) {
	o.After = after
}

// WithBefore adds the before to the admin get roles v4 params
func (o *AdminGetRolesV4Params) WithBefore(before *string) *AdminGetRolesV4Params {
	o.SetBefore(before)
	return o
}

// SetBefore adds the before to the admin get roles v4 params
func (o *AdminGetRolesV4Params) SetBefore(before *string) {
	o.Before = before
}

// WithIsWildcard adds the isWildcard to the admin get roles v4 params
func (o *AdminGetRolesV4Params) WithIsWildcard(isWildcard *bool) *AdminGetRolesV4Params {
	o.SetIsWildcard(isWildcard)
	return o
}

// SetIsWildcard adds the isWildcard to the admin get roles v4 params
func (o *AdminGetRolesV4Params) SetIsWildcard(isWildcard *bool) {
	o.IsWildcard = isWildcard
}

// WithLimit adds the limit to the admin get roles v4 params
func (o *AdminGetRolesV4Params) WithLimit(limit *int64) *AdminGetRolesV4Params {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the admin get roles v4 params
func (o *AdminGetRolesV4Params) SetLimit(limit *int64) {
	o.Limit = limit
}

// WriteToRequest writes these params to a swagger request
func (o *AdminGetRolesV4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AdminRole != nil {

		// query param adminRole
		var qrAdminRole bool
		if o.AdminRole != nil {
			qrAdminRole = *o.AdminRole
		}
		qAdminRole := swag.FormatBool(qrAdminRole)
		if qAdminRole != "" {
			if err := r.SetQueryParam("adminRole", qAdminRole); err != nil {
				return err
			}
		}

	}

	if o.After != nil {

		// query param after
		var qrAfter string
		if o.After != nil {
			qrAfter = *o.After
		}
		qAfter := qrAfter
		if qAfter != "" {
			if err := r.SetQueryParam("after", qAfter); err != nil {
				return err
			}
		}

	}

	if o.Before != nil {

		// query param before
		var qrBefore string
		if o.Before != nil {
			qrBefore = *o.Before
		}
		qBefore := qrBefore
		if qBefore != "" {
			if err := r.SetQueryParam("before", qBefore); err != nil {
				return err
			}
		}

	}

	if o.IsWildcard != nil {

		// query param isWildcard
		var qrIsWildcard bool
		if o.IsWildcard != nil {
			qrIsWildcard = *o.IsWildcard
		}
		qIsWildcard := swag.FormatBool(qrIsWildcard)
		if qIsWildcard != "" {
			if err := r.SetQueryParam("isWildcard", qIsWildcard); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
