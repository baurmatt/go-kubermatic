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

// BindUserToClusterRoleReader is a Reader for the BindUserToClusterRole structure.
type BindUserToClusterRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BindUserToClusterRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBindUserToClusterRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewBindUserToClusterRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewBindUserToClusterRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewBindUserToClusterRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBindUserToClusterRoleOK creates a BindUserToClusterRoleOK with default headers values
func NewBindUserToClusterRoleOK() *BindUserToClusterRoleOK {
	return &BindUserToClusterRoleOK{}
}

/*BindUserToClusterRoleOK handles this case with default header values.

ClusterRoleBinding
*/
type BindUserToClusterRoleOK struct {
	Payload *models.ClusterRoleBinding
}

func (o *BindUserToClusterRoleOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings][%d] bindUserToClusterRoleOK  %+v", 200, o.Payload)
}

func (o *BindUserToClusterRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterRoleBinding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBindUserToClusterRoleUnauthorized creates a BindUserToClusterRoleUnauthorized with default headers values
func NewBindUserToClusterRoleUnauthorized() *BindUserToClusterRoleUnauthorized {
	return &BindUserToClusterRoleUnauthorized{}
}

/*BindUserToClusterRoleUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type BindUserToClusterRoleUnauthorized struct {
}

func (o *BindUserToClusterRoleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings][%d] bindUserToClusterRoleUnauthorized ", 401)
}

func (o *BindUserToClusterRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBindUserToClusterRoleForbidden creates a BindUserToClusterRoleForbidden with default headers values
func NewBindUserToClusterRoleForbidden() *BindUserToClusterRoleForbidden {
	return &BindUserToClusterRoleForbidden{}
}

/*BindUserToClusterRoleForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type BindUserToClusterRoleForbidden struct {
}

func (o *BindUserToClusterRoleForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings][%d] bindUserToClusterRoleForbidden ", 403)
}

func (o *BindUserToClusterRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBindUserToClusterRoleDefault creates a BindUserToClusterRoleDefault with default headers values
func NewBindUserToClusterRoleDefault(code int) *BindUserToClusterRoleDefault {
	return &BindUserToClusterRoleDefault{
		_statusCode: code,
	}
}

/*BindUserToClusterRoleDefault handles this case with default header values.

errorResponse
*/
type BindUserToClusterRoleDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the bind user to cluster role default response
func (o *BindUserToClusterRoleDefault) Code() int {
	return o._statusCode
}

func (o *BindUserToClusterRoleDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}/clusterbindings][%d] bindUserToClusterRole default  %+v", o._statusCode, o.Payload)
}

func (o *BindUserToClusterRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
