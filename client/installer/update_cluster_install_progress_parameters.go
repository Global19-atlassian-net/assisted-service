// Code generated by go-swagger; DO NOT EDIT.

package installer

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

// NewUpdateClusterInstallProgressParams creates a new UpdateClusterInstallProgressParams object
// with the default values initialized.
func NewUpdateClusterInstallProgressParams() *UpdateClusterInstallProgressParams {
	var ()
	return &UpdateClusterInstallProgressParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateClusterInstallProgressParamsWithTimeout creates a new UpdateClusterInstallProgressParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateClusterInstallProgressParamsWithTimeout(timeout time.Duration) *UpdateClusterInstallProgressParams {
	var ()
	return &UpdateClusterInstallProgressParams{

		timeout: timeout,
	}
}

// NewUpdateClusterInstallProgressParamsWithContext creates a new UpdateClusterInstallProgressParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateClusterInstallProgressParamsWithContext(ctx context.Context) *UpdateClusterInstallProgressParams {
	var ()
	return &UpdateClusterInstallProgressParams{

		Context: ctx,
	}
}

// NewUpdateClusterInstallProgressParamsWithHTTPClient creates a new UpdateClusterInstallProgressParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateClusterInstallProgressParamsWithHTTPClient(client *http.Client) *UpdateClusterInstallProgressParams {
	var ()
	return &UpdateClusterInstallProgressParams{
		HTTPClient: client,
	}
}

/*UpdateClusterInstallProgressParams contains all the parameters to send to the API endpoint
for the update cluster install progress operation typically these are written to a http.Request
*/
type UpdateClusterInstallProgressParams struct {

	/*ClusterProgress
	  Cluster install progress value.

	*/
	ClusterProgress string
	/*ClusterID
	  The ID of the cluster to retrieve.

	*/
	ClusterID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update cluster install progress params
func (o *UpdateClusterInstallProgressParams) WithTimeout(timeout time.Duration) *UpdateClusterInstallProgressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update cluster install progress params
func (o *UpdateClusterInstallProgressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update cluster install progress params
func (o *UpdateClusterInstallProgressParams) WithContext(ctx context.Context) *UpdateClusterInstallProgressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update cluster install progress params
func (o *UpdateClusterInstallProgressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update cluster install progress params
func (o *UpdateClusterInstallProgressParams) WithHTTPClient(client *http.Client) *UpdateClusterInstallProgressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update cluster install progress params
func (o *UpdateClusterInstallProgressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterProgress adds the clusterProgress to the update cluster install progress params
func (o *UpdateClusterInstallProgressParams) WithClusterProgress(clusterProgress string) *UpdateClusterInstallProgressParams {
	o.SetClusterProgress(clusterProgress)
	return o
}

// SetClusterProgress adds the clusterProgress to the update cluster install progress params
func (o *UpdateClusterInstallProgressParams) SetClusterProgress(clusterProgress string) {
	o.ClusterProgress = clusterProgress
}

// WithClusterID adds the clusterID to the update cluster install progress params
func (o *UpdateClusterInstallProgressParams) WithClusterID(clusterID strfmt.UUID) *UpdateClusterInstallProgressParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the update cluster install progress params
func (o *UpdateClusterInstallProgressParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateClusterInstallProgressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.ClusterProgress); err != nil {
		return err
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
