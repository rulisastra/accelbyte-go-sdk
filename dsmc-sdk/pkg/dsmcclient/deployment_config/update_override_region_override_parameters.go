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

// NewUpdateOverrideRegionOverrideParams creates a new UpdateOverrideRegionOverrideParams object
// with the default values initialized.
func NewUpdateOverrideRegionOverrideParams() *UpdateOverrideRegionOverrideParams {
	var ()
	return &UpdateOverrideRegionOverrideParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOverrideRegionOverrideParamsWithTimeout creates a new UpdateOverrideRegionOverrideParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateOverrideRegionOverrideParamsWithTimeout(timeout time.Duration) *UpdateOverrideRegionOverrideParams {
	var ()
	return &UpdateOverrideRegionOverrideParams{

		timeout: timeout,
	}
}

// NewUpdateOverrideRegionOverrideParamsWithContext creates a new UpdateOverrideRegionOverrideParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateOverrideRegionOverrideParamsWithContext(ctx context.Context) *UpdateOverrideRegionOverrideParams {
	var ()
	return &UpdateOverrideRegionOverrideParams{

		Context: ctx,
	}
}

// NewUpdateOverrideRegionOverrideParamsWithHTTPClient creates a new UpdateOverrideRegionOverrideParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateOverrideRegionOverrideParamsWithHTTPClient(client *http.Client) *UpdateOverrideRegionOverrideParams {
	var ()
	return &UpdateOverrideRegionOverrideParams{
		HTTPClient: client,
	}
}

/*UpdateOverrideRegionOverrideParams contains all the parameters to send to the API endpoint
for the update override region override operation typically these are written to a http.Request
*/
type UpdateOverrideRegionOverrideParams struct {

	/*Body*/
	Body *dsmcclientmodels.ModelsUpdateRegionOverrideRequest
	/*Deployment
	  deployment name

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
	/*Version
	  version

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update override region override params
func (o *UpdateOverrideRegionOverrideParams) WithTimeout(timeout time.Duration) *UpdateOverrideRegionOverrideParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update override region override params
func (o *UpdateOverrideRegionOverrideParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update override region override params
func (o *UpdateOverrideRegionOverrideParams) WithContext(ctx context.Context) *UpdateOverrideRegionOverrideParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update override region override params
func (o *UpdateOverrideRegionOverrideParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update override region override params
func (o *UpdateOverrideRegionOverrideParams) WithHTTPClient(client *http.Client) *UpdateOverrideRegionOverrideParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update override region override params
func (o *UpdateOverrideRegionOverrideParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update override region override params
func (o *UpdateOverrideRegionOverrideParams) WithBody(body *dsmcclientmodels.ModelsUpdateRegionOverrideRequest) *UpdateOverrideRegionOverrideParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update override region override params
func (o *UpdateOverrideRegionOverrideParams) SetBody(body *dsmcclientmodels.ModelsUpdateRegionOverrideRequest) {
	o.Body = body
}

// WithDeployment adds the deployment to the update override region override params
func (o *UpdateOverrideRegionOverrideParams) WithDeployment(deployment string) *UpdateOverrideRegionOverrideParams {
	o.SetDeployment(deployment)
	return o
}

// SetDeployment adds the deployment to the update override region override params
func (o *UpdateOverrideRegionOverrideParams) SetDeployment(deployment string) {
	o.Deployment = deployment
}

// WithNamespace adds the namespace to the update override region override params
func (o *UpdateOverrideRegionOverrideParams) WithNamespace(namespace string) *UpdateOverrideRegionOverrideParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the update override region override params
func (o *UpdateOverrideRegionOverrideParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithRegion adds the region to the update override region override params
func (o *UpdateOverrideRegionOverrideParams) WithRegion(region string) *UpdateOverrideRegionOverrideParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the update override region override params
func (o *UpdateOverrideRegionOverrideParams) SetRegion(region string) {
	o.Region = region
}

// WithVersion adds the version to the update override region override params
func (o *UpdateOverrideRegionOverrideParams) WithVersion(version string) *UpdateOverrideRegionOverrideParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the update override region override params
func (o *UpdateOverrideRegionOverrideParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOverrideRegionOverrideParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
