// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
	"github.com/goharbor/harbor/src/pkg/quota/types"
)

// Quota The quota object
//
// swagger:model Quota
type Quota struct {

	// the creation time of the quota
	// Format: date-time
	CreationTime strfmt.DateTime `json:"creation_time,omitempty"`

	// The hard limits of the quota
	Hard types.ResourceList `json:"hard,omitempty"`

	// ID of the quota
	ID int64 `json:"id,omitempty"`

	// The reference object of the quota
	Ref QuotaRefObject `json:"ref,omitempty"`

	// the update time of the quota
	// Format: date-time
	UpdateTime strfmt.DateTime `json:"update_time,omitempty"`

	// The used status of the quota
	Used types.ResourceList `json:"used,omitempty"`
}

// Validate validates this quota
func (m *Quota) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Quota) validateCreationTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationTime) { // not required
		return nil
	}

	if err := validate.FormatOf("creation_time", "body", "date-time", m.CreationTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Quota) validateHard(formats strfmt.Registry) error {
	if swag.IsZero(m.Hard) { // not required
		return nil
	}

	if err := m.Hard.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("hard")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("hard")
		}
		return err
	}

	return nil
}

func (m *Quota) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Quota) validateUsed(formats strfmt.Registry) error {
	if swag.IsZero(m.Used) { // not required
		return nil
	}

	if err := m.Used.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("used")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("used")
		}
		return err
	}

	return nil
}

// ContextValidate validate this quota based on the context it is used
func (m *Quota) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Quota) contextValidateHard(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Hard.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("hard")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("hard")
		}
		return err
	}

	return nil
}

func (m *Quota) contextValidateUsed(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Used.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("used")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("used")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Quota) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Quota) UnmarshalBinary(b []byte) error {
	var res Quota
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}