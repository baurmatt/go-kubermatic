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

// GetClusterReader is a Reader for the GetCluster structure.
type GetClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetClusterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterOK creates a GetClusterOK with default headers values
func NewGetClusterOK() *GetClusterOK {
	return &GetClusterOK{}
}

/*GetClusterOK handles this case with default header values.

Cluster
*/
type GetClusterOK struct {
	Payload *models.Cluster
}

func (o *GetClusterOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}][%d] getClusterOK  %+v", 200, o.Payload)
}

func (o *GetClusterOK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *GetClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterUnauthorized creates a GetClusterUnauthorized with default headers values
func NewGetClusterUnauthorized() *GetClusterUnauthorized {
	return &GetClusterUnauthorized{}
}

/*GetClusterUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type GetClusterUnauthorized struct {
}

func (o *GetClusterUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}][%d] getClusterUnauthorized ", 401)
}

func (o *GetClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterForbidden creates a GetClusterForbidden with default headers values
func NewGetClusterForbidden() *GetClusterForbidden {
	return &GetClusterForbidden{}
}

/*GetClusterForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type GetClusterForbidden struct {
}

func (o *GetClusterForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}][%d] getClusterForbidden ", 403)
}

func (o *GetClusterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetClusterDefault creates a GetClusterDefault with default headers values
func NewGetClusterDefault(code int) *GetClusterDefault {
	return &GetClusterDefault{
		_statusCode: code,
	}
}

/*GetClusterDefault handles this case with default header values.

errorResponse
*/
type GetClusterDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get cluster default response
func (o *GetClusterDefault) Code() int {
	return o._statusCode
}

func (o *GetClusterDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}][%d] getCluster default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
