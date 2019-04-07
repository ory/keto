// Code generated by go-swagger; DO NOT EDIT.

package engines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ory/keto/sdk/go/keto/models"
)

// AddOryAccessControlPolicyRoleMembersReader is a Reader for the AddOryAccessControlPolicyRoleMembers structure.
type AddOryAccessControlPolicyRoleMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddOryAccessControlPolicyRoleMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddOryAccessControlPolicyRoleMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewAddOryAccessControlPolicyRoleMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddOryAccessControlPolicyRoleMembersOK creates a AddOryAccessControlPolicyRoleMembersOK with default headers values
func NewAddOryAccessControlPolicyRoleMembersOK() *AddOryAccessControlPolicyRoleMembersOK {
	return &AddOryAccessControlPolicyRoleMembersOK{}
}

/*AddOryAccessControlPolicyRoleMembersOK handles this case with default header values.

oryAccessControlPolicyRole
*/
type AddOryAccessControlPolicyRoleMembersOK struct {
	Payload *models.Role
}

func (o *AddOryAccessControlPolicyRoleMembersOK) Error() string {
	return fmt.Sprintf("[PUT /engines/acp/ory/{flavor}/roles/{id}/members][%d] addOryAccessControlPolicyRoleMembersOK  %+v", 200, o.Payload)
}

func (o *AddOryAccessControlPolicyRoleMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddOryAccessControlPolicyRoleMembersInternalServerError creates a AddOryAccessControlPolicyRoleMembersInternalServerError with default headers values
func NewAddOryAccessControlPolicyRoleMembersInternalServerError() *AddOryAccessControlPolicyRoleMembersInternalServerError {
	return &AddOryAccessControlPolicyRoleMembersInternalServerError{}
}

/*AddOryAccessControlPolicyRoleMembersInternalServerError handles this case with default header values.

The standard error format
*/
type AddOryAccessControlPolicyRoleMembersInternalServerError struct {
	Payload *AddOryAccessControlPolicyRoleMembersInternalServerErrorBody
}

func (o *AddOryAccessControlPolicyRoleMembersInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /engines/acp/ory/{flavor}/roles/{id}/members][%d] addOryAccessControlPolicyRoleMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *AddOryAccessControlPolicyRoleMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddOryAccessControlPolicyRoleMembersInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddOryAccessControlPolicyRoleMembersInternalServerErrorBody add ory access control policy role members internal server error body
swagger:model AddOryAccessControlPolicyRoleMembersInternalServerErrorBody
*/
type AddOryAccessControlPolicyRoleMembersInternalServerErrorBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// details
	Details []map[string]interface{} `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// request
	Request string `json:"request,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this add ory access control policy role members internal server error body
func (o *AddOryAccessControlPolicyRoleMembersInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddOryAccessControlPolicyRoleMembersInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddOryAccessControlPolicyRoleMembersInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res AddOryAccessControlPolicyRoleMembersInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
