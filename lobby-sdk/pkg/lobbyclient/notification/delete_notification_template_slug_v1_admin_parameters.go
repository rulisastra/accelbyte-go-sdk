// Code generated by go-swagger; DO NOT EDIT.

package notification

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

// NewDeleteNotificationTemplateSlugV1AdminParams creates a new DeleteNotificationTemplateSlugV1AdminParams object
// with the default values initialized.
func NewDeleteNotificationTemplateSlugV1AdminParams() *DeleteNotificationTemplateSlugV1AdminParams {
	var ()
	return &DeleteNotificationTemplateSlugV1AdminParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNotificationTemplateSlugV1AdminParamsWithTimeout creates a new DeleteNotificationTemplateSlugV1AdminParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteNotificationTemplateSlugV1AdminParamsWithTimeout(timeout time.Duration) *DeleteNotificationTemplateSlugV1AdminParams {
	var ()
	return &DeleteNotificationTemplateSlugV1AdminParams{

		timeout: timeout,
	}
}

// NewDeleteNotificationTemplateSlugV1AdminParamsWithContext creates a new DeleteNotificationTemplateSlugV1AdminParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteNotificationTemplateSlugV1AdminParamsWithContext(ctx context.Context) *DeleteNotificationTemplateSlugV1AdminParams {
	var ()
	return &DeleteNotificationTemplateSlugV1AdminParams{

		Context: ctx,
	}
}

// NewDeleteNotificationTemplateSlugV1AdminParamsWithHTTPClient creates a new DeleteNotificationTemplateSlugV1AdminParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteNotificationTemplateSlugV1AdminParamsWithHTTPClient(client *http.Client) *DeleteNotificationTemplateSlugV1AdminParams {
	var ()
	return &DeleteNotificationTemplateSlugV1AdminParams{
		HTTPClient: client,
	}
}

/*DeleteNotificationTemplateSlugV1AdminParams contains all the parameters to send to the API endpoint
for the delete notification template slug v1 admin operation typically these are written to a http.Request
*/
type DeleteNotificationTemplateSlugV1AdminParams struct {

	/*Namespace
	  namespace

	*/
	Namespace string
	/*TemplateSlug
	  template slug

	*/
	TemplateSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete notification template slug v1 admin params
func (o *DeleteNotificationTemplateSlugV1AdminParams) WithTimeout(timeout time.Duration) *DeleteNotificationTemplateSlugV1AdminParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete notification template slug v1 admin params
func (o *DeleteNotificationTemplateSlugV1AdminParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete notification template slug v1 admin params
func (o *DeleteNotificationTemplateSlugV1AdminParams) WithContext(ctx context.Context) *DeleteNotificationTemplateSlugV1AdminParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete notification template slug v1 admin params
func (o *DeleteNotificationTemplateSlugV1AdminParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete notification template slug v1 admin params
func (o *DeleteNotificationTemplateSlugV1AdminParams) WithHTTPClient(client *http.Client) *DeleteNotificationTemplateSlugV1AdminParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete notification template slug v1 admin params
func (o *DeleteNotificationTemplateSlugV1AdminParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the delete notification template slug v1 admin params
func (o *DeleteNotificationTemplateSlugV1AdminParams) WithNamespace(namespace string) *DeleteNotificationTemplateSlugV1AdminParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete notification template slug v1 admin params
func (o *DeleteNotificationTemplateSlugV1AdminParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithTemplateSlug adds the templateSlug to the delete notification template slug v1 admin params
func (o *DeleteNotificationTemplateSlugV1AdminParams) WithTemplateSlug(templateSlug string) *DeleteNotificationTemplateSlugV1AdminParams {
	o.SetTemplateSlug(templateSlug)
	return o
}

// SetTemplateSlug adds the templateSlug to the delete notification template slug v1 admin params
func (o *DeleteNotificationTemplateSlugV1AdminParams) SetTemplateSlug(templateSlug string) {
	o.TemplateSlug = templateSlug
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNotificationTemplateSlugV1AdminParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param templateSlug
	if err := r.SetPathParam("templateSlug", o.TemplateSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
