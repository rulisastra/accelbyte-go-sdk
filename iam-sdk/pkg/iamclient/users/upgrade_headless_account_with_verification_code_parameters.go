// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewUpgradeHeadlessAccountWithVerificationCodeParams creates a new UpgradeHeadlessAccountWithVerificationCodeParams object
// with the default values initialized.
func NewUpgradeHeadlessAccountWithVerificationCodeParams() *UpgradeHeadlessAccountWithVerificationCodeParams {
	var ()
	return &UpgradeHeadlessAccountWithVerificationCodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpgradeHeadlessAccountWithVerificationCodeParamsWithTimeout creates a new UpgradeHeadlessAccountWithVerificationCodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpgradeHeadlessAccountWithVerificationCodeParamsWithTimeout(timeout time.Duration) *UpgradeHeadlessAccountWithVerificationCodeParams {
	var ()
	return &UpgradeHeadlessAccountWithVerificationCodeParams{

		timeout: timeout,
	}
}

// NewUpgradeHeadlessAccountWithVerificationCodeParamsWithContext creates a new UpgradeHeadlessAccountWithVerificationCodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpgradeHeadlessAccountWithVerificationCodeParamsWithContext(ctx context.Context) *UpgradeHeadlessAccountWithVerificationCodeParams {
	var ()
	return &UpgradeHeadlessAccountWithVerificationCodeParams{

		Context: ctx,
	}
}

// NewUpgradeHeadlessAccountWithVerificationCodeParamsWithHTTPClient creates a new UpgradeHeadlessAccountWithVerificationCodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpgradeHeadlessAccountWithVerificationCodeParamsWithHTTPClient(client *http.Client) *UpgradeHeadlessAccountWithVerificationCodeParams {
	var ()
	return &UpgradeHeadlessAccountWithVerificationCodeParams{
		HTTPClient: client,
	}
}

/*UpgradeHeadlessAccountWithVerificationCodeParams contains all the parameters to send to the API endpoint
for the upgrade headless account with verification code operation typically these are written to a http.Request
*/
type UpgradeHeadlessAccountWithVerificationCodeParams struct {

	/*Body*/
	Body *iamclientmodels.ModelUpgradeHeadlessAccountWithVerificationCodeRequest
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

// WithTimeout adds the timeout to the upgrade headless account with verification code params
func (o *UpgradeHeadlessAccountWithVerificationCodeParams) WithTimeout(timeout time.Duration) *UpgradeHeadlessAccountWithVerificationCodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upgrade headless account with verification code params
func (o *UpgradeHeadlessAccountWithVerificationCodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upgrade headless account with verification code params
func (o *UpgradeHeadlessAccountWithVerificationCodeParams) WithContext(ctx context.Context) *UpgradeHeadlessAccountWithVerificationCodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upgrade headless account with verification code params
func (o *UpgradeHeadlessAccountWithVerificationCodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upgrade headless account with verification code params
func (o *UpgradeHeadlessAccountWithVerificationCodeParams) WithHTTPClient(client *http.Client) *UpgradeHeadlessAccountWithVerificationCodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upgrade headless account with verification code params
func (o *UpgradeHeadlessAccountWithVerificationCodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the upgrade headless account with verification code params
func (o *UpgradeHeadlessAccountWithVerificationCodeParams) WithBody(body *iamclientmodels.ModelUpgradeHeadlessAccountWithVerificationCodeRequest) *UpgradeHeadlessAccountWithVerificationCodeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upgrade headless account with verification code params
func (o *UpgradeHeadlessAccountWithVerificationCodeParams) SetBody(body *iamclientmodels.ModelUpgradeHeadlessAccountWithVerificationCodeRequest) {
	o.Body = body
}

// WithNamespace adds the namespace to the upgrade headless account with verification code params
func (o *UpgradeHeadlessAccountWithVerificationCodeParams) WithNamespace(namespace string) *UpgradeHeadlessAccountWithVerificationCodeParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the upgrade headless account with verification code params
func (o *UpgradeHeadlessAccountWithVerificationCodeParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithUserID adds the userID to the upgrade headless account with verification code params
func (o *UpgradeHeadlessAccountWithVerificationCodeParams) WithUserID(userID string) *UpgradeHeadlessAccountWithVerificationCodeParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the upgrade headless account with verification code params
func (o *UpgradeHeadlessAccountWithVerificationCodeParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UpgradeHeadlessAccountWithVerificationCodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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