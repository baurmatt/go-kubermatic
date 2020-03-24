// Code generated by go-swagger; DO NOT EDIT.

package hetzner

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

// NewListHetznerSizesParams creates a new ListHetznerSizesParams object
// with the default values initialized.
func NewListHetznerSizesParams() *ListHetznerSizesParams {

	return &ListHetznerSizesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListHetznerSizesParamsWithTimeout creates a new ListHetznerSizesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListHetznerSizesParamsWithTimeout(timeout time.Duration) *ListHetznerSizesParams {

	return &ListHetznerSizesParams{

		timeout: timeout,
	}
}

// NewListHetznerSizesParamsWithContext creates a new ListHetznerSizesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListHetznerSizesParamsWithContext(ctx context.Context) *ListHetznerSizesParams {

	return &ListHetznerSizesParams{

		Context: ctx,
	}
}

// NewListHetznerSizesParamsWithHTTPClient creates a new ListHetznerSizesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListHetznerSizesParamsWithHTTPClient(client *http.Client) *ListHetznerSizesParams {

	return &ListHetznerSizesParams{
		HTTPClient: client,
	}
}

/*ListHetznerSizesParams contains all the parameters to send to the API endpoint
for the list hetzner sizes operation typically these are written to a http.Request
*/
type ListHetznerSizesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list hetzner sizes params
func (o *ListHetznerSizesParams) WithTimeout(timeout time.Duration) *ListHetznerSizesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list hetzner sizes params
func (o *ListHetznerSizesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list hetzner sizes params
func (o *ListHetznerSizesParams) WithContext(ctx context.Context) *ListHetznerSizesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list hetzner sizes params
func (o *ListHetznerSizesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list hetzner sizes params
func (o *ListHetznerSizesParams) WithHTTPClient(client *http.Client) *ListHetznerSizesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list hetzner sizes params
func (o *ListHetznerSizesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListHetznerSizesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
