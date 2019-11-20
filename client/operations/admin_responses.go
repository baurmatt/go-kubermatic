// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kubermatic/go-kubermatic/models"
)

// AdminReader is a Reader for the Admin structure.
type AdminReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAdminOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewAdminUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewAdminForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAdminDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAdminOK creates a AdminOK with default headers values
func NewAdminOK() *AdminOK {
	return &AdminOK{}
}

/*AdminOK handles this case with default header values.

Admin
*/
type AdminOK struct {
	Payload []*models.Admin
}

func (o *AdminOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin][%d] adminOK  %+v", 200, o.Payload)
}

func (o *AdminOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUnauthorized creates a AdminUnauthorized with default headers values
func NewAdminUnauthorized() *AdminUnauthorized {
	return &AdminUnauthorized{}
}

/*AdminUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type AdminUnauthorized struct {
}

func (o *AdminUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin][%d] adminUnauthorized ", 401)
}

func (o *AdminUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminForbidden creates a AdminForbidden with default headers values
func NewAdminForbidden() *AdminForbidden {
	return &AdminForbidden{}
}

/*AdminForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type AdminForbidden struct {
}

func (o *AdminForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin][%d] adminForbidden ", 403)
}

func (o *AdminForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdminDefault creates a AdminDefault with default headers values
func NewAdminDefault(code int) *AdminDefault {
	return &AdminDefault{
		_statusCode: code,
	}
}

/*AdminDefault handles this case with default header values.

errorResponse
*/
type AdminDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the admin default response
func (o *AdminDefault) Code() int {
	return o._statusCode
}

func (o *AdminDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin][%d] admin default  %+v", o._statusCode, o.Payload)
}

func (o *AdminDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}