// Code generated by go-swagger; DO NOT EDIT.

package deployment_config

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

	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclientmodels"
)

// NewCreateRootRegionOverrideParams creates a new CreateRootRegionOverrideParams object
// with the default values initialized.
func NewCreateRootRegionOverrideParams() *CreateRootRegionOverrideParams {
	var ()
	return &CreateRootRegionOverrideParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRootRegionOverrideParamsWithTimeout creates a new CreateRootRegionOverrideParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRootRegionOverrideParamsWithTimeout(timeout time.Duration) *CreateRootRegionOverrideParams {
	var ()
	return &CreateRootRegionOverrideParams{

		timeout: timeout,
	}
}

// NewCreateRootRegionOverrideParamsWithContext creates a new CreateRootRegionOverrideParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRootRegionOverrideParamsWithContext(ctx context.Context) *CreateRootRegionOverrideParams {
	var ()
	return &CreateRootRegionOverrideParams{

		Context: ctx,
	}
}

// NewCreateRootRegionOverrideParamsWithHTTPClient creates a new CreateRootRegionOverrideParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRootRegionOverrideParamsWithHTTPClient(client *http.Client) *CreateRootRegionOverrideParams {
	var ()
	return &CreateRootRegionOverrideParams{
		HTTPClient: client,
	}
}

/*CreateRootRegionOverrideParams contains all the parameters to send to the API endpoint
for the create root region override operation typically these are written to a http.Request
*/
type CreateRootRegionOverrideParams struct {

	/*Body*/
	Body *dsmcclientmodels.ModelsCreateRegionOverrideRequest
	/*Deployment
	  deployment of the game

	*/
	Deployment string
	/*Namespace
	  namespace of the game

	*/
	Namespace string
	/*Region
	  region

	*/
	Region string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create root region override params
func (o *CreateRootRegionOverrideParams) WithTimeout(timeout time.Duration) *CreateRootRegionOverrideParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create root region override params
func (o *CreateRootRegionOverrideParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create root region override params
func (o *CreateRootRegionOverrideParams) WithContext(ctx context.Context) *CreateRootRegionOverrideParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create root region override params
func (o *CreateRootRegionOverrideParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create root region override params
func (o *CreateRootRegionOverrideParams) WithHTTPClient(client *http.Client) *CreateRootRegionOverrideParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create root region override params
func (o *CreateRootRegionOverrideParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create root region override params
func (o *CreateRootRegionOverrideParams) WithBody(body *dsmcclientmodels.ModelsCreateRegionOverrideRequest) *CreateRootRegionOverrideParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create root region override params
func (o *CreateRootRegionOverrideParams) SetBody(body *dsmcclientmodels.ModelsCreateRegionOverrideRequest) {
	o.Body = body
}

// WithDeployment adds the deployment to the create root region override params
func (o *CreateRootRegionOverrideParams) WithDeployment(deployment string) *CreateRootRegionOverrideParams {
	o.SetDeployment(deployment)
	return o
}

// SetDeployment adds the deployment to the create root region override params
func (o *CreateRootRegionOverrideParams) SetDeployment(deployment string) {
	o.Deployment = deployment
}

// WithNamespace adds the namespace to the create root region override params
func (o *CreateRootRegionOverrideParams) WithNamespace(namespace string) *CreateRootRegionOverrideParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the create root region override params
func (o *CreateRootRegionOverrideParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithRegion adds the region to the create root region override params
func (o *CreateRootRegionOverrideParams) WithRegion(region string) *CreateRootRegionOverrideParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the create root region override params
func (o *CreateRootRegionOverrideParams) SetRegion(region string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRootRegionOverrideParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param deployment
	if err := r.SetPathParam("deployment", o.Deployment); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param region
	if err := r.SetPathParam("region", o.Region); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
