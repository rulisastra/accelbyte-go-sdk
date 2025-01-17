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

// NewPublishTemplateLocalizationV1AdminParams creates a new PublishTemplateLocalizationV1AdminParams object
// with the default values initialized.
func NewPublishTemplateLocalizationV1AdminParams() *PublishTemplateLocalizationV1AdminParams {
	var ()
	return &PublishTemplateLocalizationV1AdminParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPublishTemplateLocalizationV1AdminParamsWithTimeout creates a new PublishTemplateLocalizationV1AdminParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPublishTemplateLocalizationV1AdminParamsWithTimeout(timeout time.Duration) *PublishTemplateLocalizationV1AdminParams {
	var ()
	return &PublishTemplateLocalizationV1AdminParams{

		timeout: timeout,
	}
}

// NewPublishTemplateLocalizationV1AdminParamsWithContext creates a new PublishTemplateLocalizationV1AdminParams object
// with the default values initialized, and the ability to set a context for a request
func NewPublishTemplateLocalizationV1AdminParamsWithContext(ctx context.Context) *PublishTemplateLocalizationV1AdminParams {
	var ()
	return &PublishTemplateLocalizationV1AdminParams{

		Context: ctx,
	}
}

// NewPublishTemplateLocalizationV1AdminParamsWithHTTPClient creates a new PublishTemplateLocalizationV1AdminParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPublishTemplateLocalizationV1AdminParamsWithHTTPClient(client *http.Client) *PublishTemplateLocalizationV1AdminParams {
	var ()
	return &PublishTemplateLocalizationV1AdminParams{
		HTTPClient: client,
	}
}

/*PublishTemplateLocalizationV1AdminParams contains all the parameters to send to the API endpoint
for the publish template localization v1 admin operation typically these are written to a http.Request
*/
type PublishTemplateLocalizationV1AdminParams struct {

	/*Namespace
	  namespace

	*/
	Namespace string
	/*TemplateLanguage
	  template language, follows IETF BCP 47 standard

	*/
	TemplateLanguage string
	/*TemplateSlug
	  Template Identifier, only alphabet characters and hyphens are permitted

	*/
	TemplateSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the publish template localization v1 admin params
func (o *PublishTemplateLocalizationV1AdminParams) WithTimeout(timeout time.Duration) *PublishTemplateLocalizationV1AdminParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the publish template localization v1 admin params
func (o *PublishTemplateLocalizationV1AdminParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the publish template localization v1 admin params
func (o *PublishTemplateLocalizationV1AdminParams) WithContext(ctx context.Context) *PublishTemplateLocalizationV1AdminParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the publish template localization v1 admin params
func (o *PublishTemplateLocalizationV1AdminParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the publish template localization v1 admin params
func (o *PublishTemplateLocalizationV1AdminParams) WithHTTPClient(client *http.Client) *PublishTemplateLocalizationV1AdminParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the publish template localization v1 admin params
func (o *PublishTemplateLocalizationV1AdminParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the publish template localization v1 admin params
func (o *PublishTemplateLocalizationV1AdminParams) WithNamespace(namespace string) *PublishTemplateLocalizationV1AdminParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the publish template localization v1 admin params
func (o *PublishTemplateLocalizationV1AdminParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithTemplateLanguage adds the templateLanguage to the publish template localization v1 admin params
func (o *PublishTemplateLocalizationV1AdminParams) WithTemplateLanguage(templateLanguage string) *PublishTemplateLocalizationV1AdminParams {
	o.SetTemplateLanguage(templateLanguage)
	return o
}

// SetTemplateLanguage adds the templateLanguage to the publish template localization v1 admin params
func (o *PublishTemplateLocalizationV1AdminParams) SetTemplateLanguage(templateLanguage string) {
	o.TemplateLanguage = templateLanguage
}

// WithTemplateSlug adds the templateSlug to the publish template localization v1 admin params
func (o *PublishTemplateLocalizationV1AdminParams) WithTemplateSlug(templateSlug string) *PublishTemplateLocalizationV1AdminParams {
	o.SetTemplateSlug(templateSlug)
	return o
}

// SetTemplateSlug adds the templateSlug to the publish template localization v1 admin params
func (o *PublishTemplateLocalizationV1AdminParams) SetTemplateSlug(templateSlug string) {
	o.TemplateSlug = templateSlug
}

// WriteToRequest writes these params to a swagger request
func (o *PublishTemplateLocalizationV1AdminParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param templateLanguage
	if err := r.SetPathParam("templateLanguage", o.TemplateLanguage); err != nil {
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
