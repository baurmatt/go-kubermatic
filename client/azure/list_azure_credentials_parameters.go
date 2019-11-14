// Code generated by go-swagger; DO NOT EDIT.

package azure

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
)

// NewListAzureCredentialsParams creates a new ListAzureCredentialsParams object
// with the default values initialized.
func NewListAzureCredentialsParams() *ListAzureCredentialsParams {

	return &ListAzureCredentialsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAzureCredentialsParamsWithTimeout creates a new ListAzureCredentialsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAzureCredentialsParamsWithTimeout(timeout time.Duration) *ListAzureCredentialsParams {

	return &ListAzureCredentialsParams{

		timeout: timeout,
	}
}

// NewListAzureCredentialsParamsWithContext creates a new ListAzureCredentialsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAzureCredentialsParamsWithContext(ctx context.Context) *ListAzureCredentialsParams {

	return &ListAzureCredentialsParams{

		Context: ctx,
	}
}

// NewListAzureCredentialsParamsWithHTTPClient creates a new ListAzureCredentialsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAzureCredentialsParamsWithHTTPClient(client *http.Client) *ListAzureCredentialsParams {

	return &ListAzureCredentialsParams{
		HTTPClient: client,
	}
}

/*ListAzureCredentialsParams contains all the parameters to send to the API endpoint
for the list azure credentials operation typically these are written to a http.Request
*/
type ListAzureCredentialsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list azure credentials params
func (o *ListAzureCredentialsParams) WithTimeout(timeout time.Duration) *ListAzureCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list azure credentials params
func (o *ListAzureCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list azure credentials params
func (o *ListAzureCredentialsParams) WithContext(ctx context.Context) *ListAzureCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list azure credentials params
func (o *ListAzureCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list azure credentials params
func (o *ListAzureCredentialsParams) WithHTTPClient(client *http.Client) *ListAzureCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list azure credentials params
func (o *ListAzureCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListAzureCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
