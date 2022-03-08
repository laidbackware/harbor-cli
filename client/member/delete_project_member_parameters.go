// Code generated by go-swagger; DO NOT EDIT.

package member

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

// NewDeleteProjectMemberParams creates a new DeleteProjectMemberParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteProjectMemberParams() *DeleteProjectMemberParams {
	return &DeleteProjectMemberParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProjectMemberParamsWithTimeout creates a new DeleteProjectMemberParams object
// with the ability to set a timeout on a request.
func NewDeleteProjectMemberParamsWithTimeout(timeout time.Duration) *DeleteProjectMemberParams {
	return &DeleteProjectMemberParams{
		timeout: timeout,
	}
}

// NewDeleteProjectMemberParamsWithContext creates a new DeleteProjectMemberParams object
// with the ability to set a context for a request.
func NewDeleteProjectMemberParamsWithContext(ctx context.Context) *DeleteProjectMemberParams {
	return &DeleteProjectMemberParams{
		Context: ctx,
	}
}

// NewDeleteProjectMemberParamsWithHTTPClient creates a new DeleteProjectMemberParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteProjectMemberParamsWithHTTPClient(client *http.Client) *DeleteProjectMemberParams {
	return &DeleteProjectMemberParams{
		HTTPClient: client,
	}
}

/* DeleteProjectMemberParams contains all the parameters to send to the API endpoint
   for the delete project member operation.

   Typically these are written to a http.Request.
*/
type DeleteProjectMemberParams struct {

	/* XIsResourceName.

	   The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.
	*/
	XIsResourceName *bool

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* Mid.

	   Member ID.

	   Format: int64
	*/
	Mid int64

	/* ProjectNameOrID.

	   The name or id of the project
	*/
	ProjectNameOrID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete project member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProjectMemberParams) WithDefaults() *DeleteProjectMemberParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete project member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProjectMemberParams) SetDefaults() {
	var (
		xIsResourceNameDefault = bool(false)
	)

	val := DeleteProjectMemberParams{
		XIsResourceName: &xIsResourceNameDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete project member params
func (o *DeleteProjectMemberParams) WithTimeout(timeout time.Duration) *DeleteProjectMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete project member params
func (o *DeleteProjectMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete project member params
func (o *DeleteProjectMemberParams) WithContext(ctx context.Context) *DeleteProjectMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete project member params
func (o *DeleteProjectMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete project member params
func (o *DeleteProjectMemberParams) WithHTTPClient(client *http.Client) *DeleteProjectMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete project member params
func (o *DeleteProjectMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXIsResourceName adds the xIsResourceName to the delete project member params
func (o *DeleteProjectMemberParams) WithXIsResourceName(xIsResourceName *bool) *DeleteProjectMemberParams {
	o.SetXIsResourceName(xIsResourceName)
	return o
}

// SetXIsResourceName adds the xIsResourceName to the delete project member params
func (o *DeleteProjectMemberParams) SetXIsResourceName(xIsResourceName *bool) {
	o.XIsResourceName = xIsResourceName
}

// WithXRequestID adds the xRequestID to the delete project member params
func (o *DeleteProjectMemberParams) WithXRequestID(xRequestID *string) *DeleteProjectMemberParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete project member params
func (o *DeleteProjectMemberParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithMid adds the mid to the delete project member params
func (o *DeleteProjectMemberParams) WithMid(mid int64) *DeleteProjectMemberParams {
	o.SetMid(mid)
	return o
}

// SetMid adds the mid to the delete project member params
func (o *DeleteProjectMemberParams) SetMid(mid int64) {
	o.Mid = mid
}

// WithProjectNameOrID adds the projectNameOrID to the delete project member params
func (o *DeleteProjectMemberParams) WithProjectNameOrID(projectNameOrID string) *DeleteProjectMemberParams {
	o.SetProjectNameOrID(projectNameOrID)
	return o
}

// SetProjectNameOrID adds the projectNameOrId to the delete project member params
func (o *DeleteProjectMemberParams) SetProjectNameOrID(projectNameOrID string) {
	o.ProjectNameOrID = projectNameOrID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProjectMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XIsResourceName != nil {

		// header param X-Is-Resource-Name
		if err := r.SetHeaderParam("X-Is-Resource-Name", swag.FormatBool(*o.XIsResourceName)); err != nil {
			return err
		}
	}

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	// path param mid
	if err := r.SetPathParam("mid", swag.FormatInt64(o.Mid)); err != nil {
		return err
	}

	// path param project_name_or_id
	if err := r.SetPathParam("project_name_or_id", o.ProjectNameOrID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
