// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/baurmatt/go-kubermatic/models"
)

// PatchClusterRoleReader is a Reader for the PatchClusterRole structure.
type PatchClusterRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchClusterRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchClusterRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchClusterRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchClusterRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchClusterRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchClusterRoleOK creates a PatchClusterRoleOK with default headers values
func NewPatchClusterRoleOK() *PatchClusterRoleOK {
	return &PatchClusterRoleOK{}
}

/*PatchClusterRoleOK handles this case with default header values.

ClusterRole
*/
type PatchClusterRoleOK struct {
	Payload *models.ClusterRole
}

func (o *PatchClusterRoleOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}][%d] patchClusterRoleOK  %+v", 200, o.Payload)
}

func (o *PatchClusterRoleOK) GetPayload() *models.ClusterRole {
	return o.Payload
}

func (o *PatchClusterRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchClusterRoleUnauthorized creates a PatchClusterRoleUnauthorized with default headers values
func NewPatchClusterRoleUnauthorized() *PatchClusterRoleUnauthorized {
	return &PatchClusterRoleUnauthorized{}
}

/*PatchClusterRoleUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type PatchClusterRoleUnauthorized struct {
}

func (o *PatchClusterRoleUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}][%d] patchClusterRoleUnauthorized ", 401)
}

func (o *PatchClusterRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchClusterRoleForbidden creates a PatchClusterRoleForbidden with default headers values
func NewPatchClusterRoleForbidden() *PatchClusterRoleForbidden {
	return &PatchClusterRoleForbidden{}
}

/*PatchClusterRoleForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type PatchClusterRoleForbidden struct {
}

func (o *PatchClusterRoleForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}][%d] patchClusterRoleForbidden ", 403)
}

func (o *PatchClusterRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchClusterRoleDefault creates a PatchClusterRoleDefault with default headers values
func NewPatchClusterRoleDefault(code int) *PatchClusterRoleDefault {
	return &PatchClusterRoleDefault{
		_statusCode: code,
	}
}

/*PatchClusterRoleDefault handles this case with default header values.

errorResponse
*/
type PatchClusterRoleDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch cluster role default response
func (o *PatchClusterRoleDefault) Code() int {
	return o._statusCode
}

func (o *PatchClusterRoleDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/clusterroles/{role_id}][%d] patchClusterRole default  %+v", o._statusCode, o.Payload)
}

func (o *PatchClusterRoleDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PatchClusterRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
