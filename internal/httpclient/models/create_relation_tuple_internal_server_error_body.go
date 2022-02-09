// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateRelationTupleInternalServerErrorBody CreateRelationTupleInternalServerErrorBody create relation tuple internal server error body
//
// swagger:model CreateRelationTupleInternalServerErrorBody
type CreateRelationTupleInternalServerErrorBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// details
	Details []interface{} `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// request
	Request string `json:"request,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this create relation tuple internal server error body
func (m *CreateRelationTupleInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create relation tuple internal server error body based on context it is used
func (m *CreateRelationTupleInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateRelationTupleInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateRelationTupleInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res CreateRelationTupleInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}