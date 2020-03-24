// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// AssignSSHKeyToClusterReader is a Reader for the AssignSSHKeyToCluster structure.
type AssignSSHKeyToClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignSSHKeyToClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAssignSSHKeyToClusterCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAssignSSHKeyToClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAssignSSHKeyToClusterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAssignSSHKeyToClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAssignSSHKeyToClusterCreated creates a AssignSSHKeyToClusterCreated with default headers values
func NewAssignSSHKeyToClusterCreated() *AssignSSHKeyToClusterCreated {
	return &AssignSSHKeyToClusterCreated{}
}

/*AssignSSHKeyToClusterCreated handles this case with default header values.

SSHKey
*/
type AssignSSHKeyToClusterCreated struct {
	Payload *models.SSHKey
}

func (o *AssignSSHKeyToClusterCreated) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/sshkeys/{key_id}][%d] assignSshKeyToClusterCreated  %+v", 201, o.Payload)
}

func (o *AssignSSHKeyToClusterCreated) GetPayload() *models.SSHKey {
	return o.Payload
}

func (o *AssignSSHKeyToClusterCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SSHKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignSSHKeyToClusterUnauthorized creates a AssignSSHKeyToClusterUnauthorized with default headers values
func NewAssignSSHKeyToClusterUnauthorized() *AssignSSHKeyToClusterUnauthorized {
	return &AssignSSHKeyToClusterUnauthorized{}
}

/*AssignSSHKeyToClusterUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type AssignSSHKeyToClusterUnauthorized struct {
}

func (o *AssignSSHKeyToClusterUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/sshkeys/{key_id}][%d] assignSshKeyToClusterUnauthorized ", 401)
}

func (o *AssignSSHKeyToClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssignSSHKeyToClusterForbidden creates a AssignSSHKeyToClusterForbidden with default headers values
func NewAssignSSHKeyToClusterForbidden() *AssignSSHKeyToClusterForbidden {
	return &AssignSSHKeyToClusterForbidden{}
}

/*AssignSSHKeyToClusterForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type AssignSSHKeyToClusterForbidden struct {
}

func (o *AssignSSHKeyToClusterForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/sshkeys/{key_id}][%d] assignSshKeyToClusterForbidden ", 403)
}

func (o *AssignSSHKeyToClusterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssignSSHKeyToClusterDefault creates a AssignSSHKeyToClusterDefault with default headers values
func NewAssignSSHKeyToClusterDefault(code int) *AssignSSHKeyToClusterDefault {
	return &AssignSSHKeyToClusterDefault{
		_statusCode: code,
	}
}

/*AssignSSHKeyToClusterDefault handles this case with default header values.

errorResponse
*/
type AssignSSHKeyToClusterDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the assign SSH key to cluster default response
func (o *AssignSSHKeyToClusterDefault) Code() int {
	return o._statusCode
}

func (o *AssignSSHKeyToClusterDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/sshkeys/{key_id}][%d] assignSSHKeyToCluster default  %+v", o._statusCode, o.Payload)
}

func (o *AssignSSHKeyToClusterDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AssignSSHKeyToClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
