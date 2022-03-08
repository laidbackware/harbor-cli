// Code generated by go-swagger; DO NOT EDIT.

package project_metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/laidbackware/harbor-cli/models"
)

// DeleteProjectMetadataReader is a Reader for the DeleteProjectMetadata structure.
type DeleteProjectMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteProjectMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteProjectMetadataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteProjectMetadataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteProjectMetadataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteProjectMetadataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteProjectMetadataConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteProjectMetadataInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteProjectMetadataOK creates a DeleteProjectMetadataOK with default headers values
func NewDeleteProjectMetadataOK() *DeleteProjectMetadataOK {
	return &DeleteProjectMetadataOK{}
}

/* DeleteProjectMetadataOK describes a response with status code 200, with default header values.

Success
*/
type DeleteProjectMetadataOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *DeleteProjectMetadataOK) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/metadatas/{meta_name}][%d] deleteProjectMetadataOK ", 200)
}

func (o *DeleteProjectMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewDeleteProjectMetadataBadRequest creates a DeleteProjectMetadataBadRequest with default headers values
func NewDeleteProjectMetadataBadRequest() *DeleteProjectMetadataBadRequest {
	return &DeleteProjectMetadataBadRequest{}
}

/* DeleteProjectMetadataBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteProjectMetadataBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *DeleteProjectMetadataBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/metadatas/{meta_name}][%d] deleteProjectMetadataBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteProjectMetadataBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteProjectMetadataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteProjectMetadataUnauthorized creates a DeleteProjectMetadataUnauthorized with default headers values
func NewDeleteProjectMetadataUnauthorized() *DeleteProjectMetadataUnauthorized {
	return &DeleteProjectMetadataUnauthorized{}
}

/* DeleteProjectMetadataUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteProjectMetadataUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *DeleteProjectMetadataUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/metadatas/{meta_name}][%d] deleteProjectMetadataUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteProjectMetadataUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteProjectMetadataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteProjectMetadataForbidden creates a DeleteProjectMetadataForbidden with default headers values
func NewDeleteProjectMetadataForbidden() *DeleteProjectMetadataForbidden {
	return &DeleteProjectMetadataForbidden{}
}

/* DeleteProjectMetadataForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteProjectMetadataForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *DeleteProjectMetadataForbidden) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/metadatas/{meta_name}][%d] deleteProjectMetadataForbidden  %+v", 403, o.Payload)
}
func (o *DeleteProjectMetadataForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteProjectMetadataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteProjectMetadataNotFound creates a DeleteProjectMetadataNotFound with default headers values
func NewDeleteProjectMetadataNotFound() *DeleteProjectMetadataNotFound {
	return &DeleteProjectMetadataNotFound{}
}

/* DeleteProjectMetadataNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteProjectMetadataNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *DeleteProjectMetadataNotFound) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/metadatas/{meta_name}][%d] deleteProjectMetadataNotFound  %+v", 404, o.Payload)
}
func (o *DeleteProjectMetadataNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteProjectMetadataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteProjectMetadataConflict creates a DeleteProjectMetadataConflict with default headers values
func NewDeleteProjectMetadataConflict() *DeleteProjectMetadataConflict {
	return &DeleteProjectMetadataConflict{}
}

/* DeleteProjectMetadataConflict describes a response with status code 409, with default header values.

Conflict
*/
type DeleteProjectMetadataConflict struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *DeleteProjectMetadataConflict) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/metadatas/{meta_name}][%d] deleteProjectMetadataConflict  %+v", 409, o.Payload)
}
func (o *DeleteProjectMetadataConflict) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteProjectMetadataConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteProjectMetadataInternalServerError creates a DeleteProjectMetadataInternalServerError with default headers values
func NewDeleteProjectMetadataInternalServerError() *DeleteProjectMetadataInternalServerError {
	return &DeleteProjectMetadataInternalServerError{}
}

/* DeleteProjectMetadataInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteProjectMetadataInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *DeleteProjectMetadataInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/metadatas/{meta_name}][%d] deleteProjectMetadataInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteProjectMetadataInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteProjectMetadataInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
