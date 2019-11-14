// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kubermatic/go-kubermatic/models"
)

// GetRoleBindingReader is a Reader for the GetRoleBinding structure.
type GetRoleBindingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoleBindingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRoleBindingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetRoleBindingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetRoleBindingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetRoleBindingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRoleBindingOK creates a GetRoleBindingOK with default headers values
func NewGetRoleBindingOK() *GetRoleBindingOK {
	return &GetRoleBindingOK{}
}

/*GetRoleBindingOK handles this case with default header values.

RoleBinding
*/
type GetRoleBindingOK struct {
	Payload *models.RoleBinding
}

func (o *GetRoleBindingOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}/bindings/{binding_id}][%d] getRoleBindingOK  %+v", 200, o.Payload)
}

func (o *GetRoleBindingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoleBinding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoleBindingUnauthorized creates a GetRoleBindingUnauthorized with default headers values
func NewGetRoleBindingUnauthorized() *GetRoleBindingUnauthorized {
	return &GetRoleBindingUnauthorized{}
}

/*GetRoleBindingUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type GetRoleBindingUnauthorized struct {
}

func (o *GetRoleBindingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}/bindings/{binding_id}][%d] getRoleBindingUnauthorized ", 401)
}

func (o *GetRoleBindingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRoleBindingForbidden creates a GetRoleBindingForbidden with default headers values
func NewGetRoleBindingForbidden() *GetRoleBindingForbidden {
	return &GetRoleBindingForbidden{}
}

/*GetRoleBindingForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type GetRoleBindingForbidden struct {
}

func (o *GetRoleBindingForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}/bindings/{binding_id}][%d] getRoleBindingForbidden ", 403)
}

func (o *GetRoleBindingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRoleBindingDefault creates a GetRoleBindingDefault with default headers values
func NewGetRoleBindingDefault(code int) *GetRoleBindingDefault {
	return &GetRoleBindingDefault{
		_statusCode: code,
	}
}

/*GetRoleBindingDefault handles this case with default header values.

errorResponse
*/
type GetRoleBindingDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get role binding default response
func (o *GetRoleBindingDefault) Code() int {
	return o._statusCode
}

func (o *GetRoleBindingDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/roles/{namespace}/{role_id}/bindings/{binding_id}][%d] getRoleBinding default  %+v", o._statusCode, o.Payload)
}

func (o *GetRoleBindingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
