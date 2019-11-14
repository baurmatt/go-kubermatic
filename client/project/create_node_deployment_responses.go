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

// CreateNodeDeploymentReader is a Reader for the CreateNodeDeployment structure.
type CreateNodeDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNodeDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateNodeDeploymentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCreateNodeDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateNodeDeploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateNodeDeploymentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateNodeDeploymentCreated creates a CreateNodeDeploymentCreated with default headers values
func NewCreateNodeDeploymentCreated() *CreateNodeDeploymentCreated {
	return &CreateNodeDeploymentCreated{}
}

/*CreateNodeDeploymentCreated handles this case with default header values.

NodeDeployment
*/
type CreateNodeDeploymentCreated struct {
	Payload *models.NodeDeployment
}

func (o *CreateNodeDeploymentCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments][%d] createNodeDeploymentCreated  %+v", 201, o.Payload)
}

func (o *CreateNodeDeploymentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeDeployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNodeDeploymentUnauthorized creates a CreateNodeDeploymentUnauthorized with default headers values
func NewCreateNodeDeploymentUnauthorized() *CreateNodeDeploymentUnauthorized {
	return &CreateNodeDeploymentUnauthorized{}
}

/*CreateNodeDeploymentUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type CreateNodeDeploymentUnauthorized struct {
}

func (o *CreateNodeDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments][%d] createNodeDeploymentUnauthorized ", 401)
}

func (o *CreateNodeDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateNodeDeploymentForbidden creates a CreateNodeDeploymentForbidden with default headers values
func NewCreateNodeDeploymentForbidden() *CreateNodeDeploymentForbidden {
	return &CreateNodeDeploymentForbidden{}
}

/*CreateNodeDeploymentForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type CreateNodeDeploymentForbidden struct {
}

func (o *CreateNodeDeploymentForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments][%d] createNodeDeploymentForbidden ", 403)
}

func (o *CreateNodeDeploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateNodeDeploymentDefault creates a CreateNodeDeploymentDefault with default headers values
func NewCreateNodeDeploymentDefault(code int) *CreateNodeDeploymentDefault {
	return &CreateNodeDeploymentDefault{
		_statusCode: code,
	}
}

/*CreateNodeDeploymentDefault handles this case with default header values.

errorResponse
*/
type CreateNodeDeploymentDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create node deployment default response
func (o *CreateNodeDeploymentDefault) Code() int {
	return o._statusCode
}

func (o *CreateNodeDeploymentDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments][%d] createNodeDeployment default  %+v", o._statusCode, o.Payload)
}

func (o *CreateNodeDeploymentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
