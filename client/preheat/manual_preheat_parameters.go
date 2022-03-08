// Code generated by go-swagger; DO NOT EDIT.

package preheat

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

	"github.com/laidbackware/harbor-cli/models"
)

// NewManualPreheatParams creates a new ManualPreheatParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewManualPreheatParams() *ManualPreheatParams {
	return &ManualPreheatParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewManualPreheatParamsWithTimeout creates a new ManualPreheatParams object
// with the ability to set a timeout on a request.
func NewManualPreheatParamsWithTimeout(timeout time.Duration) *ManualPreheatParams {
	return &ManualPreheatParams{
		timeout: timeout,
	}
}

// NewManualPreheatParamsWithContext creates a new ManualPreheatParams object
// with the ability to set a context for a request.
func NewManualPreheatParamsWithContext(ctx context.Context) *ManualPreheatParams {
	return &ManualPreheatParams{
		Context: ctx,
	}
}

// NewManualPreheatParamsWithHTTPClient creates a new ManualPreheatParams object
// with the ability to set a custom HTTPClient for a request.
func NewManualPreheatParamsWithHTTPClient(client *http.Client) *ManualPreheatParams {
	return &ManualPreheatParams{
		HTTPClient: client,
	}
}

/* ManualPreheatParams contains all the parameters to send to the API endpoint
   for the manual preheat operation.

   Typically these are written to a http.Request.
*/
type ManualPreheatParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* Policy.

	   The policy schema info
	*/
	Policy *models.PreheatPolicy

	/* PreheatPolicyName.

	   Preheat Policy Name
	*/
	PreheatPolicyName string

	/* ProjectName.

	   The name of the project
	*/
	ProjectName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the manual preheat params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ManualPreheatParams) WithDefaults() *ManualPreheatParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the manual preheat params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ManualPreheatParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the manual preheat params
func (o *ManualPreheatParams) WithTimeout(timeout time.Duration) *ManualPreheatParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the manual preheat params
func (o *ManualPreheatParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the manual preheat params
func (o *ManualPreheatParams) WithContext(ctx context.Context) *ManualPreheatParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the manual preheat params
func (o *ManualPreheatParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the manual preheat params
func (o *ManualPreheatParams) WithHTTPClient(client *http.Client) *ManualPreheatParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the manual preheat params
func (o *ManualPreheatParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the manual preheat params
func (o *ManualPreheatParams) WithXRequestID(xRequestID *string) *ManualPreheatParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the manual preheat params
func (o *ManualPreheatParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithPolicy adds the policy to the manual preheat params
func (o *ManualPreheatParams) WithPolicy(policy *models.PreheatPolicy) *ManualPreheatParams {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the manual preheat params
func (o *ManualPreheatParams) SetPolicy(policy *models.PreheatPolicy) {
	o.Policy = policy
}

// WithPreheatPolicyName adds the preheatPolicyName to the manual preheat params
func (o *ManualPreheatParams) WithPreheatPolicyName(preheatPolicyName string) *ManualPreheatParams {
	o.SetPreheatPolicyName(preheatPolicyName)
	return o
}

// SetPreheatPolicyName adds the preheatPolicyName to the manual preheat params
func (o *ManualPreheatParams) SetPreheatPolicyName(preheatPolicyName string) {
	o.PreheatPolicyName = preheatPolicyName
}

// WithProjectName adds the projectName to the manual preheat params
func (o *ManualPreheatParams) WithProjectName(projectName string) *ManualPreheatParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the manual preheat params
func (o *ManualPreheatParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WriteToRequest writes these params to a swagger request
func (o *ManualPreheatParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}
	if o.Policy != nil {
		if err := r.SetBodyParam(o.Policy); err != nil {
			return err
		}
	}

	// path param preheat_policy_name
	if err := r.SetPathParam("preheat_policy_name", o.PreheatPolicyName); err != nil {
		return err
	}

	// path param project_name
	if err := r.SetPathParam("project_name", o.ProjectName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}