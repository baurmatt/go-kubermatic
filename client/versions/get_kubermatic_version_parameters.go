// Code generated by go-swagger; DO NOT EDIT.

package versions

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
)

// NewGetKubermaticVersionParams creates a new GetKubermaticVersionParams object
// with the default values initialized.
func NewGetKubermaticVersionParams() *GetKubermaticVersionParams {

	return &GetKubermaticVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKubermaticVersionParamsWithTimeout creates a new GetKubermaticVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKubermaticVersionParamsWithTimeout(timeout time.Duration) *GetKubermaticVersionParams {

	return &GetKubermaticVersionParams{

		timeout: timeout,
	}
}

// NewGetKubermaticVersionParamsWithContext creates a new GetKubermaticVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKubermaticVersionParamsWithContext(ctx context.Context) *GetKubermaticVersionParams {

	return &GetKubermaticVersionParams{

		Context: ctx,
	}
}

// NewGetKubermaticVersionParamsWithHTTPClient creates a new GetKubermaticVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKubermaticVersionParamsWithHTTPClient(client *http.Client) *GetKubermaticVersionParams {

	return &GetKubermaticVersionParams{
		HTTPClient: client,
	}
}

/*GetKubermaticVersionParams contains all the parameters to send to the API endpoint
for the get kubermatic version operation typically these are written to a http.Request
*/
type GetKubermaticVersionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get kubermatic version params
func (o *GetKubermaticVersionParams) WithTimeout(timeout time.Duration) *GetKubermaticVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get kubermatic version params
func (o *GetKubermaticVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get kubermatic version params
func (o *GetKubermaticVersionParams) WithContext(ctx context.Context) *GetKubermaticVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get kubermatic version params
func (o *GetKubermaticVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get kubermatic version params
func (o *GetKubermaticVersionParams) WithHTTPClient(client *http.Client) *GetKubermaticVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get kubermatic version params
func (o *GetKubermaticVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetKubermaticVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
