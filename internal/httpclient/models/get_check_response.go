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
)

// GetCheckResponse GetCheckResponse Represents the response for a check request.
//
// The content of the allowed field is mirrored in the HTTP status code.
//
// swagger:model getCheckResponse
type GetCheckResponse struct {

	// whether the relation tuple is allowed
	// Required: true
	Allowed *bool `json:"allowed"`
}

// Validate validates this get check response
func (m *GetCheckResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCheckResponse) validateAllowed(formats strfmt.Registry) error {

	if err := validate.Required("allowed", "body", m.Allowed); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get check response based on context it is used
func (m *GetCheckResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetCheckResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCheckResponse) UnmarshalBinary(b []byte) error {
	var res GetCheckResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
