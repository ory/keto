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

// InternalRelationTuple internal relation tuple
//
// swagger:model InternalRelationTuple
type InternalRelationTuple struct {

	// Namespace of the Relation Tuple
	// Required: true
	Namespace *string `json:"namespace"`

	// Object of the Relation Tuple
	// Required: true
	Object *string `json:"object"`

	// Relation of the Relation Tuple
	// Required: true
	Relation *string `json:"relation"`

	// SubjectID of the Relation Tuple
	//
	// Either SubjectSet or SubjectID are required.
	SubjectID string `json:"subject_id,omitempty"`

	// subject set
	SubjectSet *SubjectSet `json:"subject_set,omitempty"`
}

// Validate validates this internal relation tuple
func (m *InternalRelationTuple) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubjectSet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InternalRelationTuple) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *InternalRelationTuple) validateObject(formats strfmt.Registry) error {

	if err := validate.Required("object", "body", m.Object); err != nil {
		return err
	}

	return nil
}

func (m *InternalRelationTuple) validateRelation(formats strfmt.Registry) error {

	if err := validate.Required("relation", "body", m.Relation); err != nil {
		return err
	}

	return nil
}

func (m *InternalRelationTuple) validateSubjectSet(formats strfmt.Registry) error {
	if swag.IsZero(m.SubjectSet) { // not required
		return nil
	}

	if m.SubjectSet != nil {
		if err := m.SubjectSet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subject_set")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subject_set")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this internal relation tuple based on the context it is used
func (m *InternalRelationTuple) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSubjectSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InternalRelationTuple) contextValidateSubjectSet(ctx context.Context, formats strfmt.Registry) error {

	if m.SubjectSet != nil {
		if err := m.SubjectSet.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subject_set")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subject_set")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InternalRelationTuple) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InternalRelationTuple) UnmarshalBinary(b []byte) error {
	var res InternalRelationTuple
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
