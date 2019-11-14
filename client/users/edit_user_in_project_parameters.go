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

// NewEditUserInProjectParams creates a new EditUserInProjectParams object
// with the default values initialized.
func NewEditUserInProjectParams() *EditUserInProjectParams {
	var ()
	return &EditUserInProjectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEditUserInProjectParamsWithTimeout creates a new EditUserInProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEditUserInProjectParamsWithTimeout(timeout time.Duration) *EditUserInProjectParams {
	var ()
	return &EditUserInProjectParams{

		timeout: timeout,
	}
}

// NewEditUserInProjectParamsWithContext creates a new EditUserInProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewEditUserInProjectParamsWithContext(ctx context.Context) *EditUserInProjectParams {
	var ()
	return &EditUserInProjectParams{

		Context: ctx,
	}
}

// NewEditUserInProjectParamsWithHTTPClient creates a new EditUserInProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEditUserInProjectParamsWithHTTPClient(client *http.Client) *EditUserInProjectParams {
	var ()
	return &EditUserInProjectParams{
		HTTPClient: client,
	}
}

/*EditUserInProjectParams contains all the parameters to send to the API endpoint
for the edit user in project operation typically these are written to a http.Request
*/
type EditUserInProjectParams struct {

	/*Body*/
	Body *models.User
	/*ProjectID*/
	ProjectID string
	/*UserID*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the edit user in project params
func (o *EditUserInProjectParams) WithTimeout(timeout time.Duration) *EditUserInProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edit user in project params
func (o *EditUserInProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edit user in project params
func (o *EditUserInProjectParams) WithContext(ctx context.Context) *EditUserInProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edit user in project params
func (o *EditUserInProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edit user in project params
func (o *EditUserInProjectParams) WithHTTPClient(client *http.Client) *EditUserInProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edit user in project params
func (o *EditUserInProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the edit user in project params
func (o *EditUserInProjectParams) WithBody(body *models.User) *EditUserInProjectParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the edit user in project params
func (o *EditUserInProjectParams) SetBody(body *models.User) {
	o.Body = body
}

// WithProjectID adds the projectID to the edit user in project params
func (o *EditUserInProjectParams) WithProjectID(projectID string) *EditUserInProjectParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the edit user in project params
func (o *EditUserInProjectParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithUserID adds the userID to the edit user in project params
func (o *EditUserInProjectParams) WithUserID(userID string) *EditUserInProjectParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the edit user in project params
func (o *EditUserInProjectParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *EditUserInProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param user_id
	if err := r.SetPathParam("user_id", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
