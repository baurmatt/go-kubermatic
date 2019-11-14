// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kubermatic/go-kubermatic/models"
)

// NewAddUserToProjectParams creates a new AddUserToProjectParams object
// with the default values initialized.
func NewAddUserToProjectParams() *AddUserToProjectParams {
	var ()
	return &AddUserToProjectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddUserToProjectParamsWithTimeout creates a new AddUserToProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddUserToProjectParamsWithTimeout(timeout time.Duration) *AddUserToProjectParams {
	var ()
	return &AddUserToProjectParams{

		timeout: timeout,
	}
}

// NewAddUserToProjectParamsWithContext creates a new AddUserToProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddUserToProjectParamsWithContext(ctx context.Context) *AddUserToProjectParams {
	var ()
	return &AddUserToProjectParams{

		Context: ctx,
	}
}

// NewAddUserToProjectParamsWithHTTPClient creates a new AddUserToProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddUserToProjectParamsWithHTTPClient(client *http.Client) *AddUserToProjectParams {
	var ()
	return &AddUserToProjectParams{
		HTTPClient: client,
	}
}

/*AddUserToProjectParams contains all the parameters to send to the API endpoint
for the add user to project operation typically these are written to a http.Request
*/
type AddUserToProjectParams struct {

	/*Body*/
	Body *models.User
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add user to project params
func (o *AddUserToProjectParams) WithTimeout(timeout time.Duration) *AddUserToProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add user to project params
func (o *AddUserToProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add user to project params
func (o *AddUserToProjectParams) WithContext(ctx context.Context) *AddUserToProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add user to project params
func (o *AddUserToProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add user to project params
func (o *AddUserToProjectParams) WithHTTPClient(client *http.Client) *AddUserToProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add user to project params
func (o *AddUserToProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add user to project params
func (o *AddUserToProjectParams) WithBody(body *models.User) *AddUserToProjectParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add user to project params
func (o *AddUserToProjectParams) SetBody(body *models.User) {
	o.Body = body
}

// WithProjectID adds the projectID to the add user to project params
func (o *AddUserToProjectParams) WithProjectID(projectID string) *AddUserToProjectParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the add user to project params
func (o *AddUserToProjectParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *AddUserToProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
