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

// GetOidcClusterKubeconfigReader is a Reader for the GetOidcClusterKubeconfig structure.
type GetOidcClusterKubeconfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOidcClusterKubeconfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOidcClusterKubeconfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetOidcClusterKubeconfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOidcClusterKubeconfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetOidcClusterKubeconfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOidcClusterKubeconfigOK creates a GetOidcClusterKubeconfigOK with default headers values
func NewGetOidcClusterKubeconfigOK() *GetOidcClusterKubeconfigOK {
	return &GetOidcClusterKubeconfigOK{}
}

/*GetOidcClusterKubeconfigOK handles this case with default header values.

Kubeconfig is a clusters kubeconfig
*/
type GetOidcClusterKubeconfigOK struct {
	Payload *models.Config
}

func (o *GetOidcClusterKubeconfigOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/oidckubeconfig][%d] getOidcClusterKubeconfigOK  %+v", 200, o.Payload)
}

func (o *GetOidcClusterKubeconfigOK) GetPayload() *models.Config {
	return o.Payload
}

func (o *GetOidcClusterKubeconfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Config)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOidcClusterKubeconfigUnauthorized creates a GetOidcClusterKubeconfigUnauthorized with default headers values
func NewGetOidcClusterKubeconfigUnauthorized() *GetOidcClusterKubeconfigUnauthorized {
	return &GetOidcClusterKubeconfigUnauthorized{}
}

/*GetOidcClusterKubeconfigUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type GetOidcClusterKubeconfigUnauthorized struct {
}

func (o *GetOidcClusterKubeconfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/oidckubeconfig][%d] getOidcClusterKubeconfigUnauthorized ", 401)
}

func (o *GetOidcClusterKubeconfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOidcClusterKubeconfigForbidden creates a GetOidcClusterKubeconfigForbidden with default headers values
func NewGetOidcClusterKubeconfigForbidden() *GetOidcClusterKubeconfigForbidden {
	return &GetOidcClusterKubeconfigForbidden{}
}

/*GetOidcClusterKubeconfigForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type GetOidcClusterKubeconfigForbidden struct {
}

func (o *GetOidcClusterKubeconfigForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/oidckubeconfig][%d] getOidcClusterKubeconfigForbidden ", 403)
}

func (o *GetOidcClusterKubeconfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOidcClusterKubeconfigDefault creates a GetOidcClusterKubeconfigDefault with default headers values
func NewGetOidcClusterKubeconfigDefault(code int) *GetOidcClusterKubeconfigDefault {
	return &GetOidcClusterKubeconfigDefault{
		_statusCode: code,
	}
}

/*GetOidcClusterKubeconfigDefault handles this case with default header values.

errorResponse
*/
type GetOidcClusterKubeconfigDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get oidc cluster kubeconfig default response
func (o *GetOidcClusterKubeconfigDefault) Code() int {
	return o._statusCode
}

func (o *GetOidcClusterKubeconfigDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/oidckubeconfig][%d] getOidcClusterKubeconfig default  %+v", o._statusCode, o.Payload)
}

func (o *GetOidcClusterKubeconfigDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetOidcClusterKubeconfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
