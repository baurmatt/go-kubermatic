// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// ListVSphereNetworksNoCredentialsReader is a Reader for the ListVSphereNetworksNoCredentials structure.
type ListVSphereNetworksNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListVSphereNetworksNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListVSphereNetworksNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListVSphereNetworksNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListVSphereNetworksNoCredentialsOK creates a ListVSphereNetworksNoCredentialsOK with default headers values
func NewListVSphereNetworksNoCredentialsOK() *ListVSphereNetworksNoCredentialsOK {
	return &ListVSphereNetworksNoCredentialsOK{}
}

/*ListVSphereNetworksNoCredentialsOK handles this case with default header values.

VSphereNetwork
*/
type ListVSphereNetworksNoCredentialsOK struct {
	Payload []*models.VSphereNetwork
}

func (o *ListVSphereNetworksNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/vsphere/networks][%d] listVSphereNetworksNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListVSphereNetworksNoCredentialsOK) GetPayload() []*models.VSphereNetwork {
	return o.Payload
}

func (o *ListVSphereNetworksNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVSphereNetworksNoCredentialsDefault creates a ListVSphereNetworksNoCredentialsDefault with default headers values
func NewListVSphereNetworksNoCredentialsDefault(code int) *ListVSphereNetworksNoCredentialsDefault {
	return &ListVSphereNetworksNoCredentialsDefault{
		_statusCode: code,
	}
}

/*ListVSphereNetworksNoCredentialsDefault handles this case with default header values.

errorResponse
*/
type ListVSphereNetworksNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list v sphere networks no credentials default response
func (o *ListVSphereNetworksNoCredentialsDefault) Code() int {
	return o._statusCode
}

func (o *ListVSphereNetworksNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/vsphere/networks][%d] listVSphereNetworksNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListVSphereNetworksNoCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListVSphereNetworksNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
