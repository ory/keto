// Code generated by go-swagger; DO NOT EDIT.

package write

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/ory/keto/internal/httpclient/models"
)

// CreateRelationTupleReader is a Reader for the CreateRelationTuple structure.
type CreateRelationTupleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRelationTupleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRelationTupleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRelationTupleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateRelationTupleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRelationTupleCreated creates a CreateRelationTupleCreated with default headers values
func NewCreateRelationTupleCreated() *CreateRelationTupleCreated {
	return &CreateRelationTupleCreated{}
}

/* CreateRelationTupleCreated describes a response with status code 201, with default header values.

InternalRelationTuple
*/
type CreateRelationTupleCreated struct {
	Payload *models.InternalRelationTuple
}

func (o *CreateRelationTupleCreated) Error() string {
	return fmt.Sprintf("[PUT /relationtuple][%d] createRelationTupleCreated  %+v", 201, o.Payload)
}
func (o *CreateRelationTupleCreated) GetPayload() *models.InternalRelationTuple {
	return o.Payload
}

func (o *CreateRelationTupleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalRelationTuple)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRelationTupleBadRequest creates a CreateRelationTupleBadRequest with default headers values
func NewCreateRelationTupleBadRequest() *CreateRelationTupleBadRequest {
	return &CreateRelationTupleBadRequest{}
}

/* CreateRelationTupleBadRequest describes a response with status code 400, with default header values.

The standard error format
*/
type CreateRelationTupleBadRequest struct {
	Payload *CreateRelationTupleBadRequestBody
}

func (o *CreateRelationTupleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /relationtuple][%d] createRelationTupleBadRequest  %+v", 400, o.Payload)
}
func (o *CreateRelationTupleBadRequest) GetPayload() *CreateRelationTupleBadRequestBody {
	return o.Payload
}

func (o *CreateRelationTupleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateRelationTupleBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRelationTupleInternalServerError creates a CreateRelationTupleInternalServerError with default headers values
func NewCreateRelationTupleInternalServerError() *CreateRelationTupleInternalServerError {
	return &CreateRelationTupleInternalServerError{}
}

/* CreateRelationTupleInternalServerError describes a response with status code 500, with default header values.

The standard error format
*/
type CreateRelationTupleInternalServerError struct {
	Payload *CreateRelationTupleInternalServerErrorBody
}

func (o *CreateRelationTupleInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /relationtuple][%d] createRelationTupleInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateRelationTupleInternalServerError) GetPayload() *CreateRelationTupleInternalServerErrorBody {
	return o.Payload
}

func (o *CreateRelationTupleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateRelationTupleInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateRelationTupleBadRequestBody create relation tuple bad request body
swagger:model CreateRelationTupleBadRequestBody
*/
type CreateRelationTupleBadRequestBody struct {

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

// Validate validates this create relation tuple bad request body
func (o *CreateRelationTupleBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create relation tuple bad request body based on context it is used
func (o *CreateRelationTupleBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateRelationTupleBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateRelationTupleBadRequestBody) UnmarshalBinary(b []byte) error {
	var res CreateRelationTupleBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateRelationTupleInternalServerErrorBody create relation tuple internal server error body
swagger:model CreateRelationTupleInternalServerErrorBody
*/
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
func (o *CreateRelationTupleInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create relation tuple internal server error body based on context it is used
func (o *CreateRelationTupleInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateRelationTupleInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateRelationTupleInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res CreateRelationTupleInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
