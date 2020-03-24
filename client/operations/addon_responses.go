// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/go-kubermatic/models"
)

// AddonReader is a Reader for the Addon structure.
type AddonReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddonReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddonOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddonUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddonForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddonDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddonOK creates a AddonOK with default headers values
func NewAddonOK() *AddonOK {
	return &AddonOK{}
}

/*AddonOK handles this case with default header values.

AccessibleAddons
*/
type AddonOK struct {
	Payload models.AccessibleAddons
}

func (o *AddonOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/addons][%d] addonOK  %+v", 200, o.Payload)
}

func (o *AddonOK) GetPayload() models.AccessibleAddons {
	return o.Payload
}

func (o *AddonOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddonUnauthorized creates a AddonUnauthorized with default headers values
func NewAddonUnauthorized() *AddonUnauthorized {
	return &AddonUnauthorized{}
}

/*AddonUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type AddonUnauthorized struct {
}

func (o *AddonUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/addons][%d] addonUnauthorized ", 401)
}

func (o *AddonUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddonForbidden creates a AddonForbidden with default headers values
func NewAddonForbidden() *AddonForbidden {
	return &AddonForbidden{}
}

/*AddonForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type AddonForbidden struct {
}

func (o *AddonForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/addons][%d] addonForbidden ", 403)
}

func (o *AddonForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddonDefault creates a AddonDefault with default headers values
func NewAddonDefault(code int) *AddonDefault {
	return &AddonDefault{
		_statusCode: code,
	}
}

/*AddonDefault handles this case with default header values.

errorResponse
*/
type AddonDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the addon default response
func (o *AddonDefault) Code() int {
	return o._statusCode
}

func (o *AddonDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/addons][%d] addon default  %+v", o._statusCode, o.Payload)
}

func (o *AddonDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddonDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
