// Code generated by go-swagger; DO NOT EDIT.

package scan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/laidbackware/harbor-cli/models"
)

// ScanArtifactReader is a Reader for the ScanArtifact structure.
type ScanArtifactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScanArtifactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewScanArtifactAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewScanArtifactBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewScanArtifactUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewScanArtifactForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewScanArtifactNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewScanArtifactInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewScanArtifactAccepted creates a ScanArtifactAccepted with default headers values
func NewScanArtifactAccepted() *ScanArtifactAccepted {
	return &ScanArtifactAccepted{}
}

/* ScanArtifactAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ScanArtifactAccepted struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *ScanArtifactAccepted) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan][%d] scanArtifactAccepted ", 202)
}

func (o *ScanArtifactAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewScanArtifactBadRequest creates a ScanArtifactBadRequest with default headers values
func NewScanArtifactBadRequest() *ScanArtifactBadRequest {
	return &ScanArtifactBadRequest{}
}

/* ScanArtifactBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ScanArtifactBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ScanArtifactBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan][%d] scanArtifactBadRequest  %+v", 400, o.Payload)
}
func (o *ScanArtifactBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ScanArtifactBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewScanArtifactUnauthorized creates a ScanArtifactUnauthorized with default headers values
func NewScanArtifactUnauthorized() *ScanArtifactUnauthorized {
	return &ScanArtifactUnauthorized{}
}

/* ScanArtifactUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ScanArtifactUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ScanArtifactUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan][%d] scanArtifactUnauthorized  %+v", 401, o.Payload)
}
func (o *ScanArtifactUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ScanArtifactUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewScanArtifactForbidden creates a ScanArtifactForbidden with default headers values
func NewScanArtifactForbidden() *ScanArtifactForbidden {
	return &ScanArtifactForbidden{}
}

/* ScanArtifactForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ScanArtifactForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ScanArtifactForbidden) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan][%d] scanArtifactForbidden  %+v", 403, o.Payload)
}
func (o *ScanArtifactForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ScanArtifactForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewScanArtifactNotFound creates a ScanArtifactNotFound with default headers values
func NewScanArtifactNotFound() *ScanArtifactNotFound {
	return &ScanArtifactNotFound{}
}

/* ScanArtifactNotFound describes a response with status code 404, with default header values.

Not found
*/
type ScanArtifactNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ScanArtifactNotFound) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan][%d] scanArtifactNotFound  %+v", 404, o.Payload)
}
func (o *ScanArtifactNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ScanArtifactNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewScanArtifactInternalServerError creates a ScanArtifactInternalServerError with default headers values
func NewScanArtifactInternalServerError() *ScanArtifactInternalServerError {
	return &ScanArtifactInternalServerError{}
}

/* ScanArtifactInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ScanArtifactInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *ScanArtifactInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/scan][%d] scanArtifactInternalServerError  %+v", 500, o.Payload)
}
func (o *ScanArtifactInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ScanArtifactInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
