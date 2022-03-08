// Code generated by go-swagger; DO NOT EDIT.

package robotv1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/laidbackware/harbor-cli/models"
)

// ListRobotV1Reader is a Reader for the ListRobotV1 structure.
type ListRobotV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRobotV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRobotV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListRobotV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListRobotV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListRobotV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListRobotV1OK creates a ListRobotV1OK with default headers values
func NewListRobotV1OK() *ListRobotV1OK {
	return &ListRobotV1OK{}
}

/* ListRobotV1OK describes a response with status code 200, with default header values.

Success
*/
type ListRobotV1OK struct {

	/* Link refers to the previous page and next page
	 */
	Link string

	/* The total count of robot accounts
	 */
	XTotalCount int64

	Payload []*models.Robot
}

func (o *ListRobotV1OK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/robots][%d] listRobotV1OK  %+v", 200, o.Payload)
}
func (o *ListRobotV1OK) GetPayload() []*models.Robot {
	return o.Payload
}

func (o *ListRobotV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// hydrates response header X-Total-Count
	hdrXTotalCount := response.GetHeader("X-Total-Count")

	if hdrXTotalCount != "" {
		valxTotalCount, err := swag.ConvertInt64(hdrXTotalCount)
		if err != nil {
			return errors.InvalidType("X-Total-Count", "header", "int64", hdrXTotalCount)
		}
		o.XTotalCount = valxTotalCount
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRobotV1BadRequest creates a ListRobotV1BadRequest with default headers values
func NewListRobotV1BadRequest() *ListRobotV1BadRequest {
	return &ListRobotV1BadRequest{}
}

/* ListRobotV1BadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ListRobotV1BadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListRobotV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/robots][%d] listRobotV1BadRequest  %+v", 400, o.Payload)
}
func (o *ListRobotV1BadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListRobotV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRobotV1NotFound creates a ListRobotV1NotFound with default headers values
func NewListRobotV1NotFound() *ListRobotV1NotFound {
	return &ListRobotV1NotFound{}
}

/* ListRobotV1NotFound describes a response with status code 404, with default header values.

Not found
*/
type ListRobotV1NotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListRobotV1NotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/robots][%d] listRobotV1NotFound  %+v", 404, o.Payload)
}
func (o *ListRobotV1NotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListRobotV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRobotV1InternalServerError creates a ListRobotV1InternalServerError with default headers values
func NewListRobotV1InternalServerError() *ListRobotV1InternalServerError {
	return &ListRobotV1InternalServerError{}
}

/* ListRobotV1InternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ListRobotV1InternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ListRobotV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/robots][%d] listRobotV1InternalServerError  %+v", 500, o.Payload)
}
func (o *ListRobotV1InternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListRobotV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
