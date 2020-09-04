// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/baurmatt/go-kubermatic/models"
)

// LogoutCurrentUserReader is a Reader for the LogoutCurrentUser structure.
type LogoutCurrentUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogoutCurrentUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLogoutCurrentUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewLogoutCurrentUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewLogoutCurrentUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLogoutCurrentUserOK creates a LogoutCurrentUserOK with default headers values
func NewLogoutCurrentUserOK() *LogoutCurrentUserOK {
	return &LogoutCurrentUserOK{}
}

/*LogoutCurrentUserOK handles this case with default header values.

EmptyResponse is a empty response
*/
type LogoutCurrentUserOK struct {
}

func (o *LogoutCurrentUserOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/me/logout][%d] logoutCurrentUserOK ", 200)
}

func (o *LogoutCurrentUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLogoutCurrentUserUnauthorized creates a LogoutCurrentUserUnauthorized with default headers values
func NewLogoutCurrentUserUnauthorized() *LogoutCurrentUserUnauthorized {
	return &LogoutCurrentUserUnauthorized{}
}

/*LogoutCurrentUserUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type LogoutCurrentUserUnauthorized struct {
}

func (o *LogoutCurrentUserUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/me/logout][%d] logoutCurrentUserUnauthorized ", 401)
}

func (o *LogoutCurrentUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLogoutCurrentUserDefault creates a LogoutCurrentUserDefault with default headers values
func NewLogoutCurrentUserDefault(code int) *LogoutCurrentUserDefault {
	return &LogoutCurrentUserDefault{
		_statusCode: code,
	}
}

/*LogoutCurrentUserDefault handles this case with default header values.

errorResponse
*/
type LogoutCurrentUserDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the logout current user default response
func (o *LogoutCurrentUserDefault) Code() int {
	return o._statusCode
}

func (o *LogoutCurrentUserDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/me/logout][%d] logoutCurrentUser default  %+v", o._statusCode, o.Payload)
}

func (o *LogoutCurrentUserDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LogoutCurrentUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
