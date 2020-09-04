// Code generated by go-swagger; DO NOT EDIT.

package serviceaccounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/baurmatt/go-kubermatic/models"
)

// NewAddServiceAccountToProjectParams creates a new AddServiceAccountToProjectParams object
// with the default values initialized.
func NewAddServiceAccountToProjectParams() *AddServiceAccountToProjectParams {
	var ()
	return &AddServiceAccountToProjectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddServiceAccountToProjectParamsWithTimeout creates a new AddServiceAccountToProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddServiceAccountToProjectParamsWithTimeout(timeout time.Duration) *AddServiceAccountToProjectParams {
	var ()
	return &AddServiceAccountToProjectParams{

		timeout: timeout,
	}
}

// NewAddServiceAccountToProjectParamsWithContext creates a new AddServiceAccountToProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddServiceAccountToProjectParamsWithContext(ctx context.Context) *AddServiceAccountToProjectParams {
	var ()
	return &AddServiceAccountToProjectParams{

		Context: ctx,
	}
}

// NewAddServiceAccountToProjectParamsWithHTTPClient creates a new AddServiceAccountToProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddServiceAccountToProjectParamsWithHTTPClient(client *http.Client) *AddServiceAccountToProjectParams {
	var ()
	return &AddServiceAccountToProjectParams{
		HTTPClient: client,
	}
}

/*AddServiceAccountToProjectParams contains all the parameters to send to the API endpoint
for the add service account to project operation typically these are written to a http.Request
*/
type AddServiceAccountToProjectParams struct {

	/*Body*/
	Body *models.ServiceAccount
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add service account to project params
func (o *AddServiceAccountToProjectParams) WithTimeout(timeout time.Duration) *AddServiceAccountToProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add service account to project params
func (o *AddServiceAccountToProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add service account to project params
func (o *AddServiceAccountToProjectParams) WithContext(ctx context.Context) *AddServiceAccountToProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add service account to project params
func (o *AddServiceAccountToProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add service account to project params
func (o *AddServiceAccountToProjectParams) WithHTTPClient(client *http.Client) *AddServiceAccountToProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add service account to project params
func (o *AddServiceAccountToProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add service account to project params
func (o *AddServiceAccountToProjectParams) WithBody(body *models.ServiceAccount) *AddServiceAccountToProjectParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add service account to project params
func (o *AddServiceAccountToProjectParams) SetBody(body *models.ServiceAccount) {
	o.Body = body
}

// WithProjectID adds the projectID to the add service account to project params
func (o *AddServiceAccountToProjectParams) WithProjectID(projectID string) *AddServiceAccountToProjectParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the add service account to project params
func (o *AddServiceAccountToProjectParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *AddServiceAccountToProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
