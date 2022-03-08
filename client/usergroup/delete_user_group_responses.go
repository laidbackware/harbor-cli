// Code generated by go-swagger; DO NOT EDIT.

package usergroup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/laidbackware/harbor-cli/models"
)

// DeleteUserGroupReader is a Reader for the DeleteUserGroup structure.
type DeleteUserGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUserGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUserGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteUserGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteUserGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteUserGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserGroupOK creates a DeleteUserGroupOK with default headers values
func NewDeleteUserGroupOK() *DeleteUserGroupOK {
	return &DeleteUserGroupOK{}
}

/* DeleteUserGroupOK describes a response with status code 200, with default header values.

Success
*/
type DeleteUserGroupOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *DeleteUserGroupOK) Error() string {
	return fmt.Sprintf("[DELETE /usergroups/{group_id}][%d] deleteUserGroupOK ", 200)
}

func (o *DeleteUserGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewDeleteUserGroupBadRequest creates a DeleteUserGroupBadRequest with default headers values
func NewDeleteUserGroupBadRequest() *DeleteUserGroupBadRequest {
	return &DeleteUserGroupBadRequest{}
}

/* DeleteUserGroupBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteUserGroupBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *DeleteUserGroupBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /usergroups/{group_id}][%d] deleteUserGroupBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteUserGroupBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteUserGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteUserGroupUnauthorized creates a DeleteUserGroupUnauthorized with default headers values
func NewDeleteUserGroupUnauthorized() *DeleteUserGroupUnauthorized {
	return &DeleteUserGroupUnauthorized{}
}

/* DeleteUserGroupUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteUserGroupUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *DeleteUserGroupUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /usergroups/{group_id}][%d] deleteUserGroupUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteUserGroupUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteUserGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteUserGroupForbidden creates a DeleteUserGroupForbidden with default headers values
func NewDeleteUserGroupForbidden() *DeleteUserGroupForbidden {
	return &DeleteUserGroupForbidden{}
}

/* DeleteUserGroupForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteUserGroupForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *DeleteUserGroupForbidden) Error() string {
	return fmt.Sprintf("[DELETE /usergroups/{group_id}][%d] deleteUserGroupForbidden  %+v", 403, o.Payload)
}
func (o *DeleteUserGroupForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteUserGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteUserGroupInternalServerError creates a DeleteUserGroupInternalServerError with default headers values
func NewDeleteUserGroupInternalServerError() *DeleteUserGroupInternalServerError {
	return &DeleteUserGroupInternalServerError{}
}

/* DeleteUserGroupInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteUserGroupInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *DeleteUserGroupInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /usergroups/{group_id}][%d] deleteUserGroupInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteUserGroupInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteUserGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
