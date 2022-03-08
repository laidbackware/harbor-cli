// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/laidbackware/harbor-cli/models"
)

// LastTriggerReader is a Reader for the LastTrigger structure.
type LastTriggerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LastTriggerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLastTriggerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLastTriggerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewLastTriggerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewLastTriggerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLastTriggerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLastTriggerOK creates a LastTriggerOK with default headers values
func NewLastTriggerOK() *LastTriggerOK {
	return &LastTriggerOK{}
}

/* LastTriggerOK describes a response with status code 200, with default header values.

Test webhook connection successfully.
*/
type LastTriggerOK struct {
	Payload []*models.WebhookLastTrigger
}

func (o *LastTriggerOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/lasttrigger][%d] lastTriggerOK  %+v", 200, o.Payload)
}
func (o *LastTriggerOK) GetPayload() []*models.WebhookLastTrigger {
	return o.Payload
}

func (o *LastTriggerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLastTriggerBadRequest creates a LastTriggerBadRequest with default headers values
func NewLastTriggerBadRequest() *LastTriggerBadRequest {
	return &LastTriggerBadRequest{}
}

/* LastTriggerBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type LastTriggerBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *LastTriggerBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/lasttrigger][%d] lastTriggerBadRequest  %+v", 400, o.Payload)
}
func (o *LastTriggerBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *LastTriggerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewLastTriggerUnauthorized creates a LastTriggerUnauthorized with default headers values
func NewLastTriggerUnauthorized() *LastTriggerUnauthorized {
	return &LastTriggerUnauthorized{}
}

/* LastTriggerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type LastTriggerUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *LastTriggerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/lasttrigger][%d] lastTriggerUnauthorized  %+v", 401, o.Payload)
}
func (o *LastTriggerUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *LastTriggerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewLastTriggerForbidden creates a LastTriggerForbidden with default headers values
func NewLastTriggerForbidden() *LastTriggerForbidden {
	return &LastTriggerForbidden{}
}

/* LastTriggerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type LastTriggerForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *LastTriggerForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/lasttrigger][%d] lastTriggerForbidden  %+v", 403, o.Payload)
}
func (o *LastTriggerForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *LastTriggerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewLastTriggerInternalServerError creates a LastTriggerInternalServerError with default headers values
func NewLastTriggerInternalServerError() *LastTriggerInternalServerError {
	return &LastTriggerInternalServerError{}
}

/* LastTriggerInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type LastTriggerInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *LastTriggerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/lasttrigger][%d] lastTriggerInternalServerError  %+v", 500, o.Payload)
}
func (o *LastTriggerInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *LastTriggerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
