// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetOryAccessControlPolicyRole GetOryAccessControlPolicyRole get ory access control policy role
// swagger:model GetOryAccessControlPolicyRole
type GetOryAccessControlPolicyRole struct {

	// The ORY Access Control Policy flavor. Can be "regex", "glob", and "exact".
	//
	// in: path
	// Required: true
	Flavor *string `json:"flavor"`

	// The ID of the ORY Access Control Policy Role.
	//
	// in: path
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this get ory access control policy role
func (m *GetOryAccessControlPolicyRole) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlavor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetOryAccessControlPolicyRole) validateFlavor(formats strfmt.Registry) error {

	if err := validate.Required("flavor", "body", m.Flavor); err != nil {
		return err
	}

	return nil
}

func (m *GetOryAccessControlPolicyRole) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetOryAccessControlPolicyRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetOryAccessControlPolicyRole) UnmarshalBinary(b []byte) error {
	var res GetOryAccessControlPolicyRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
