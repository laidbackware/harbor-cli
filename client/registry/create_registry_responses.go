// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/laidbackware/harbor-cli/models"
)

// CreateRegistryReader is a Reader for the CreateRegistry structure.
type CreateRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRegistryCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRegistryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateRegistryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRegistryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateRegistryConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateRegistryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRegistryCreated creates a CreateRegistryCreated with default headers values
func NewCreateRegistryCreated() *CreateRegistryCreated {
	return &CreateRegistryCreated{}
}

/* CreateRegistryCreated describes a response with status code 201, with default header values.

Created
*/
type CreateRegistryCreated struct {

	/* The location of the resource
	 */
	Location string

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *CreateRegistryCreated) Error() string {
	return fmt.Sprintf("[POST /registries][%d] createRegistryCreated ", 201)
}

func (o *CreateRegistryCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewCreateRegistryBadRequest creates a CreateRegistryBadRequest with default headers values
func NewCreateRegistryBadRequest() *CreateRegistryBadRequest {
	return &CreateRegistryBadRequest{}
}

/* CreateRegistryBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateRegistryBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *CreateRegistryBadRequest) Error() string {
	return fmt.Sprintf("[POST /registries][%d] createRegistryBadRequest  %+v", 400, o.Payload)
}
func (o *CreateRegistryBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateRegistryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateRegistryUnauthorized creates a CreateRegistryUnauthorized with default headers values
func NewCreateRegistryUnauthorized() *CreateRegistryUnauthorized {
	return &CreateRegistryUnauthorized{}
}

/* CreateRegistryUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateRegistryUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *CreateRegistryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /registries][%d] createRegistryUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateRegistryUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateRegistryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateRegistryForbidden creates a CreateRegistryForbidden with default headers values
func NewCreateRegistryForbidden() *CreateRegistryForbidden {
	return &CreateRegistryForbidden{}
}

/* CreateRegistryForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateRegistryForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *CreateRegistryForbidden) Error() string {
	return fmt.Sprintf("[POST /registries][%d] createRegistryForbidden  %+v", 403, o.Payload)
}
func (o *CreateRegistryForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateRegistryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateRegistryConflict creates a CreateRegistryConflict with default headers values
func NewCreateRegistryConflict() *CreateRegistryConflict {
	return &CreateRegistryConflict{}
}

/* CreateRegistryConflict describes a response with status code 409, with default header values.

Conflict
*/
type CreateRegistryConflict struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *CreateRegistryConflict) Error() string {
	return fmt.Sprintf("[POST /registries][%d] createRegistryConflict  %+v", 409, o.Payload)
}
func (o *CreateRegistryConflict) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateRegistryConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateRegistryInternalServerError creates a CreateRegistryInternalServerError with default headers values
func NewCreateRegistryInternalServerError() *CreateRegistryInternalServerError {
	return &CreateRegistryInternalServerError{}
}

/* CreateRegistryInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type CreateRegistryInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *CreateRegistryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /registries][%d] createRegistryInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateRegistryInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateRegistryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
