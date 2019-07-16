// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UpsertOryAccessControlPolicyInternalServerError UpsertOryAccessControlPolicyInternalServerError UpsertOryAccessControlPolicyInternalServerError handles this case with default header values.
//
// The standard error format
// swagger:model UpsertOryAccessControlPolicyInternalServerError
type UpsertOryAccessControlPolicyInternalServerError struct {

	// payload
	Payload *UpsertOryAccessControlPolicyInternalServerErrorBody `json:"Payload,omitempty"`
}

// Validate validates this upsert ory access control policy internal server error
func (m *UpsertOryAccessControlPolicyInternalServerError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayload(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpsertOryAccessControlPolicyInternalServerError) validatePayload(formats strfmt.Registry) error {

	if swag.IsZero(m.Payload) { // not required
		return nil
	}

	if m.Payload != nil {
		if err := m.Payload.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Payload")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpsertOryAccessControlPolicyInternalServerError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpsertOryAccessControlPolicyInternalServerError) UnmarshalBinary(b []byte) error {
	var res UpsertOryAccessControlPolicyInternalServerError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
