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
)

// RemoveOryAccessControlPolicyRoleMembersReader is a Reader for the RemoveOryAccessControlPolicyRoleMembers structure.
type RemoveOryAccessControlPolicyRoleMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveOryAccessControlPolicyRoleMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRemoveOryAccessControlPolicyRoleMembersCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewRemoveOryAccessControlPolicyRoleMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveOryAccessControlPolicyRoleMembersCreated creates a RemoveOryAccessControlPolicyRoleMembersCreated with default headers values
func NewRemoveOryAccessControlPolicyRoleMembersCreated() *RemoveOryAccessControlPolicyRoleMembersCreated {
	return &RemoveOryAccessControlPolicyRoleMembersCreated{}
}

/*RemoveOryAccessControlPolicyRoleMembersCreated handles this case with default header values.

An empty response
*/
type RemoveOryAccessControlPolicyRoleMembersCreated struct {
}

func (o *RemoveOryAccessControlPolicyRoleMembersCreated) Error() string {
	return fmt.Sprintf("[DELETE /engines/acp/ory/{flavor}/roles/{id}/members/{member}][%d] removeOryAccessControlPolicyRoleMembersCreated ", 201)
}

func (o *RemoveOryAccessControlPolicyRoleMembersCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveOryAccessControlPolicyRoleMembersInternalServerError creates a RemoveOryAccessControlPolicyRoleMembersInternalServerError with default headers values
func NewRemoveOryAccessControlPolicyRoleMembersInternalServerError() *RemoveOryAccessControlPolicyRoleMembersInternalServerError {
	return &RemoveOryAccessControlPolicyRoleMembersInternalServerError{}
}

/*RemoveOryAccessControlPolicyRoleMembersInternalServerError handles this case with default header values.

The standard error format
*/
type RemoveOryAccessControlPolicyRoleMembersInternalServerError struct {
	Payload *RemoveOryAccessControlPolicyRoleMembersInternalServerErrorBody
}

func (o *RemoveOryAccessControlPolicyRoleMembersInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /engines/acp/ory/{flavor}/roles/{id}/members/{member}][%d] removeOryAccessControlPolicyRoleMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveOryAccessControlPolicyRoleMembersInternalServerError) GetPayload() *RemoveOryAccessControlPolicyRoleMembersInternalServerErrorBody {
	return o.Payload
}

func (o *RemoveOryAccessControlPolicyRoleMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RemoveOryAccessControlPolicyRoleMembersInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RemoveOryAccessControlPolicyRoleMembersInternalServerErrorBody remove ory access control policy role members internal server error body
swagger:model RemoveOryAccessControlPolicyRoleMembersInternalServerErrorBody
*/
type RemoveOryAccessControlPolicyRoleMembersInternalServerErrorBody struct {

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

// Validate validates this remove ory access control policy role members internal server error body
func (o *RemoveOryAccessControlPolicyRoleMembersInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoveOryAccessControlPolicyRoleMembersInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveOryAccessControlPolicyRoleMembersInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res RemoveOryAccessControlPolicyRoleMembersInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
