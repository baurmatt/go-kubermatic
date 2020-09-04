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

// GetCurrentUserReader is a Reader for the GetCurrentUser structure.
type GetCurrentUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrentUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCurrentUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCurrentUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCurrentUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCurrentUserOK creates a GetCurrentUserOK with default headers values
func NewGetCurrentUserOK() *GetCurrentUserOK {
	return &GetCurrentUserOK{}
}

/*GetCurrentUserOK handles this case with default header values.

User
*/
type GetCurrentUserOK struct {
	Payload *models.User
}

func (o *GetCurrentUserOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/me][%d] getCurrentUserOK  %+v", 200, o.Payload)
}

func (o *GetCurrentUserOK) GetPayload() *models.User {
	return o.Payload
}

func (o *GetCurrentUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentUserUnauthorized creates a GetCurrentUserUnauthorized with default headers values
func NewGetCurrentUserUnauthorized() *GetCurrentUserUnauthorized {
	return &GetCurrentUserUnauthorized{}
}

/*GetCurrentUserUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type GetCurrentUserUnauthorized struct {
}

func (o *GetCurrentUserUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/me][%d] getCurrentUserUnauthorized ", 401)
}

func (o *GetCurrentUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCurrentUserDefault creates a GetCurrentUserDefault with default headers values
func NewGetCurrentUserDefault(code int) *GetCurrentUserDefault {
	return &GetCurrentUserDefault{
		_statusCode: code,
	}
}

/*GetCurrentUserDefault handles this case with default header values.

errorResponse
*/
type GetCurrentUserDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get current user default response
func (o *GetCurrentUserDefault) Code() int {
	return o._statusCode
}

func (o *GetCurrentUserDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/me][%d] getCurrentUser default  %+v", o._statusCode, o.Payload)
}

func (o *GetCurrentUserDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetCurrentUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
