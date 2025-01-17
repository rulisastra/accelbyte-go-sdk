// Code generated by go-swagger; DO NOT EDIT.

package o_auth2_0

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

// NewRetrieveUserThirdPartyPlatformTokenV3Params creates a new RetrieveUserThirdPartyPlatformTokenV3Params object
// with the default values initialized.
func NewRetrieveUserThirdPartyPlatformTokenV3Params() *RetrieveUserThirdPartyPlatformTokenV3Params {
	var ()
	return &RetrieveUserThirdPartyPlatformTokenV3Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewRetrieveUserThirdPartyPlatformTokenV3ParamsWithTimeout creates a new RetrieveUserThirdPartyPlatformTokenV3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewRetrieveUserThirdPartyPlatformTokenV3ParamsWithTimeout(timeout time.Duration) *RetrieveUserThirdPartyPlatformTokenV3Params {
	var ()
	return &RetrieveUserThirdPartyPlatformTokenV3Params{

		timeout: timeout,
	}
}

// NewRetrieveUserThirdPartyPlatformTokenV3ParamsWithContext creates a new RetrieveUserThirdPartyPlatformTokenV3Params object
// with the default values initialized, and the ability to set a context for a request
func NewRetrieveUserThirdPartyPlatformTokenV3ParamsWithContext(ctx context.Context) *RetrieveUserThirdPartyPlatformTokenV3Params {
	var ()
	return &RetrieveUserThirdPartyPlatformTokenV3Params{

		Context: ctx,
	}
}

// NewRetrieveUserThirdPartyPlatformTokenV3ParamsWithHTTPClient creates a new RetrieveUserThirdPartyPlatformTokenV3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRetrieveUserThirdPartyPlatformTokenV3ParamsWithHTTPClient(client *http.Client) *RetrieveUserThirdPartyPlatformTokenV3Params {
	var ()
	return &RetrieveUserThirdPartyPlatformTokenV3Params{
		HTTPClient: client,
	}
}

/*RetrieveUserThirdPartyPlatformTokenV3Params contains all the parameters to send to the API endpoint
for the retrieve user third party platform token v3 operation typically these are written to a http.Request
*/
type RetrieveUserThirdPartyPlatformTokenV3Params struct {

	/*Namespace
	  Namespace, only accept alphabet and numeric

	*/
	Namespace string
	/*PlatformID
	  Platform ID

	*/
	PlatformID string
	/*UserID
	  UserId

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the retrieve user third party platform token v3 params
func (o *RetrieveUserThirdPartyPlatformTokenV3Params) WithTimeout(timeout time.Duration) *RetrieveUserThirdPartyPlatformTokenV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve user third party platform token v3 params
func (o *RetrieveUserThirdPartyPlatformTokenV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve user third party platform token v3 params
func (o *RetrieveUserThirdPartyPlatformTokenV3Params) WithContext(ctx context.Context) *RetrieveUserThirdPartyPlatformTokenV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve user third party platform token v3 params
func (o *RetrieveUserThirdPartyPlatformTokenV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve user third party platform token v3 params
func (o *RetrieveUserThirdPartyPlatformTokenV3Params) WithHTTPClient(client *http.Client) *RetrieveUserThirdPartyPlatformTokenV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve user third party platform token v3 params
func (o *RetrieveUserThirdPartyPlatformTokenV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the retrieve user third party platform token v3 params
func (o *RetrieveUserThirdPartyPlatformTokenV3Params) WithNamespace(namespace string) *RetrieveUserThirdPartyPlatformTokenV3Params {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the retrieve user third party platform token v3 params
func (o *RetrieveUserThirdPartyPlatformTokenV3Params) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithPlatformID adds the platformID to the retrieve user third party platform token v3 params
func (o *RetrieveUserThirdPartyPlatformTokenV3Params) WithPlatformID(platformID string) *RetrieveUserThirdPartyPlatformTokenV3Params {
	o.SetPlatformID(platformID)
	return o
}

// SetPlatformID adds the platformId to the retrieve user third party platform token v3 params
func (o *RetrieveUserThirdPartyPlatformTokenV3Params) SetPlatformID(platformID string) {
	o.PlatformID = platformID
}

// WithUserID adds the userID to the retrieve user third party platform token v3 params
func (o *RetrieveUserThirdPartyPlatformTokenV3Params) WithUserID(userID string) *RetrieveUserThirdPartyPlatformTokenV3Params {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the retrieve user third party platform token v3 params
func (o *RetrieveUserThirdPartyPlatformTokenV3Params) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *RetrieveUserThirdPartyPlatformTokenV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param platformId
	if err := r.SetPathParam("platformId", o.PlatformID); err != nil {
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
