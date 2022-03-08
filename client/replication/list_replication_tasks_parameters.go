// Code generated by go-swagger; DO NOT EDIT.

package replication

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

// NewListReplicationTasksParams creates a new ListReplicationTasksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListReplicationTasksParams() *ListReplicationTasksParams {
	return &ListReplicationTasksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListReplicationTasksParamsWithTimeout creates a new ListReplicationTasksParams object
// with the ability to set a timeout on a request.
func NewListReplicationTasksParamsWithTimeout(timeout time.Duration) *ListReplicationTasksParams {
	return &ListReplicationTasksParams{
		timeout: timeout,
	}
}

// NewListReplicationTasksParamsWithContext creates a new ListReplicationTasksParams object
// with the ability to set a context for a request.
func NewListReplicationTasksParamsWithContext(ctx context.Context) *ListReplicationTasksParams {
	return &ListReplicationTasksParams{
		Context: ctx,
	}
}

// NewListReplicationTasksParamsWithHTTPClient creates a new ListReplicationTasksParams object
// with the ability to set a custom HTTPClient for a request.
func NewListReplicationTasksParamsWithHTTPClient(client *http.Client) *ListReplicationTasksParams {
	return &ListReplicationTasksParams{
		HTTPClient: client,
	}
}

/* ListReplicationTasksParams contains all the parameters to send to the API endpoint
   for the list replication tasks operation.

   Typically these are written to a http.Request.
*/
type ListReplicationTasksParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* ID.

	   The ID of the execution that the tasks belongs to.

	   Format: int64
	*/
	ID int64

	/* Page.

	   The page number

	   Format: int64
	   Default: 1
	*/
	Page *int64

	/* PageSize.

	   The size of per page

	   Format: int64
	   Default: 10
	*/
	PageSize *int64

	/* ResourceType.

	   The resource type.
	*/
	ResourceType *string

	/* Sort.

	   Sort the resource list in ascending or descending order. e.g. sort by field1 in ascending orderr and field2 in descending order with 'sort=field1,-field2'
	*/
	Sort *string

	/* Status.

	   The task status.
	*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list replication tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListReplicationTasksParams) WithDefaults() *ListReplicationTasksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list replication tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListReplicationTasksParams) SetDefaults() {
	var (
		pageDefault = int64(1)

		pageSizeDefault = int64(10)
	)

	val := ListReplicationTasksParams{
		Page:     &pageDefault,
		PageSize: &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list replication tasks params
func (o *ListReplicationTasksParams) WithTimeout(timeout time.Duration) *ListReplicationTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list replication tasks params
func (o *ListReplicationTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list replication tasks params
func (o *ListReplicationTasksParams) WithContext(ctx context.Context) *ListReplicationTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list replication tasks params
func (o *ListReplicationTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list replication tasks params
func (o *ListReplicationTasksParams) WithHTTPClient(client *http.Client) *ListReplicationTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list replication tasks params
func (o *ListReplicationTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the list replication tasks params
func (o *ListReplicationTasksParams) WithXRequestID(xRequestID *string) *ListReplicationTasksParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the list replication tasks params
func (o *ListReplicationTasksParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the list replication tasks params
func (o *ListReplicationTasksParams) WithID(id int64) *ListReplicationTasksParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the list replication tasks params
func (o *ListReplicationTasksParams) SetID(id int64) {
	o.ID = id
}

// WithPage adds the page to the list replication tasks params
func (o *ListReplicationTasksParams) WithPage(page *int64) *ListReplicationTasksParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list replication tasks params
func (o *ListReplicationTasksParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the list replication tasks params
func (o *ListReplicationTasksParams) WithPageSize(pageSize *int64) *ListReplicationTasksParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list replication tasks params
func (o *ListReplicationTasksParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithResourceType adds the resourceType to the list replication tasks params
func (o *ListReplicationTasksParams) WithResourceType(resourceType *string) *ListReplicationTasksParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the list replication tasks params
func (o *ListReplicationTasksParams) SetResourceType(resourceType *string) {
	o.ResourceType = resourceType
}

// WithSort adds the sort to the list replication tasks params
func (o *ListReplicationTasksParams) WithSort(sort *string) *ListReplicationTasksParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the list replication tasks params
func (o *ListReplicationTasksParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithStatus adds the status to the list replication tasks params
func (o *ListReplicationTasksParams) WithStatus(status *string) *ListReplicationTasksParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the list replication tasks params
func (o *ListReplicationTasksParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *ListReplicationTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.ResourceType != nil {

		// query param resource_type
		var qrResourceType string

		if o.ResourceType != nil {
			qrResourceType = *o.ResourceType
		}
		qResourceType := qrResourceType
		if qResourceType != "" {

			if err := r.SetQueryParam("resource_type", qResourceType); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string

		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {

			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}
	}

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
