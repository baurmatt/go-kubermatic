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

// DeleteNodeForClusterLegacyReader is a Reader for the DeleteNodeForClusterLegacy structure.
type DeleteNodeForClusterLegacyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNodeForClusterLegacyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNodeForClusterLegacyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteNodeForClusterLegacyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteNodeForClusterLegacyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteNodeForClusterLegacyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNodeForClusterLegacyOK creates a DeleteNodeForClusterLegacyOK with default headers values
func NewDeleteNodeForClusterLegacyOK() *DeleteNodeForClusterLegacyOK {
	return &DeleteNodeForClusterLegacyOK{}
}

/*DeleteNodeForClusterLegacyOK handles this case with default header values.

EmptyResponse is a empty response
*/
type DeleteNodeForClusterLegacyOK struct {
}

func (o *DeleteNodeForClusterLegacyOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes/{node_id}][%d] deleteNodeForClusterLegacyOK ", 200)
}

func (o *DeleteNodeForClusterLegacyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNodeForClusterLegacyUnauthorized creates a DeleteNodeForClusterLegacyUnauthorized with default headers values
func NewDeleteNodeForClusterLegacyUnauthorized() *DeleteNodeForClusterLegacyUnauthorized {
	return &DeleteNodeForClusterLegacyUnauthorized{}
}

/*DeleteNodeForClusterLegacyUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type DeleteNodeForClusterLegacyUnauthorized struct {
}

func (o *DeleteNodeForClusterLegacyUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes/{node_id}][%d] deleteNodeForClusterLegacyUnauthorized ", 401)
}

func (o *DeleteNodeForClusterLegacyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNodeForClusterLegacyForbidden creates a DeleteNodeForClusterLegacyForbidden with default headers values
func NewDeleteNodeForClusterLegacyForbidden() *DeleteNodeForClusterLegacyForbidden {
	return &DeleteNodeForClusterLegacyForbidden{}
}

/*DeleteNodeForClusterLegacyForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type DeleteNodeForClusterLegacyForbidden struct {
}

func (o *DeleteNodeForClusterLegacyForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes/{node_id}][%d] deleteNodeForClusterLegacyForbidden ", 403)
}

func (o *DeleteNodeForClusterLegacyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNodeForClusterLegacyDefault creates a DeleteNodeForClusterLegacyDefault with default headers values
func NewDeleteNodeForClusterLegacyDefault(code int) *DeleteNodeForClusterLegacyDefault {
	return &DeleteNodeForClusterLegacyDefault{
		_statusCode: code,
	}
}

/*DeleteNodeForClusterLegacyDefault handles this case with default header values.

errorResponse
*/
type DeleteNodeForClusterLegacyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete node for cluster legacy default response
func (o *DeleteNodeForClusterLegacyDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNodeForClusterLegacyDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes/{node_id}][%d] deleteNodeForClusterLegacy default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteNodeForClusterLegacyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteNodeForClusterLegacyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
