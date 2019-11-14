// Code generated by go-swagger; DO NOT EDIT.

package addon

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

// NewGetAddonParams creates a new GetAddonParams object
// with the default values initialized.
func NewGetAddonParams() *GetAddonParams {
	var ()
	return &GetAddonParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAddonParamsWithTimeout creates a new GetAddonParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAddonParamsWithTimeout(timeout time.Duration) *GetAddonParams {
	var ()
	return &GetAddonParams{

		timeout: timeout,
	}
}

// NewGetAddonParamsWithContext creates a new GetAddonParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAddonParamsWithContext(ctx context.Context) *GetAddonParams {
	var ()
	return &GetAddonParams{

		Context: ctx,
	}
}

// NewGetAddonParamsWithHTTPClient creates a new GetAddonParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAddonParamsWithHTTPClient(client *http.Client) *GetAddonParams {
	var ()
	return &GetAddonParams{
		HTTPClient: client,
	}
}

/*GetAddonParams contains all the parameters to send to the API endpoint
for the get addon operation typically these are written to a http.Request
*/
type GetAddonParams struct {

	/*AddonID*/
	AddonID string
	/*ClusterID*/
	ClusterID string
	/*Dc*/
	Dc string
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get addon params
func (o *GetAddonParams) WithTimeout(timeout time.Duration) *GetAddonParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get addon params
func (o *GetAddonParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get addon params
func (o *GetAddonParams) WithContext(ctx context.Context) *GetAddonParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get addon params
func (o *GetAddonParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get addon params
func (o *GetAddonParams) WithHTTPClient(client *http.Client) *GetAddonParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get addon params
func (o *GetAddonParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddonID adds the addonID to the get addon params
func (o *GetAddonParams) WithAddonID(addonID string) *GetAddonParams {
	o.SetAddonID(addonID)
	return o
}

// SetAddonID adds the addonId to the get addon params
func (o *GetAddonParams) SetAddonID(addonID string) {
	o.AddonID = addonID
}

// WithClusterID adds the clusterID to the get addon params
func (o *GetAddonParams) WithClusterID(clusterID string) *GetAddonParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get addon params
func (o *GetAddonParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDc adds the dc to the get addon params
func (o *GetAddonParams) WithDc(dc string) *GetAddonParams {
	o.SetDc(dc)
	return o
}

// SetDc adds the dc to the get addon params
func (o *GetAddonParams) SetDc(dc string) {
	o.Dc = dc
}

// WithProjectID adds the projectID to the get addon params
func (o *GetAddonParams) WithProjectID(projectID string) *GetAddonParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get addon params
func (o *GetAddonParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAddonParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param addon_id
	if err := r.SetPathParam("addon_id", o.AddonID); err != nil {
		return err
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param dc
	if err := r.SetPathParam("dc", o.Dc); err != nil {
		return err
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
