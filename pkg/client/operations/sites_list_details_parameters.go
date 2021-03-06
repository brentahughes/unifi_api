// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewSitesListDetailsParams creates a new SitesListDetailsParams object
// with the default values initialized.
func NewSitesListDetailsParams() *SitesListDetailsParams {

	return &SitesListDetailsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSitesListDetailsParamsWithTimeout creates a new SitesListDetailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSitesListDetailsParamsWithTimeout(timeout time.Duration) *SitesListDetailsParams {

	return &SitesListDetailsParams{

		timeout: timeout,
	}
}

// NewSitesListDetailsParamsWithContext creates a new SitesListDetailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSitesListDetailsParamsWithContext(ctx context.Context) *SitesListDetailsParams {

	return &SitesListDetailsParams{

		Context: ctx,
	}
}

// NewSitesListDetailsParamsWithHTTPClient creates a new SitesListDetailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSitesListDetailsParamsWithHTTPClient(client *http.Client) *SitesListDetailsParams {

	return &SitesListDetailsParams{
		HTTPClient: client,
	}
}

/*SitesListDetailsParams contains all the parameters to send to the API endpoint
for the sites list details operation typically these are written to a http.Request
*/
type SitesListDetailsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the sites list details params
func (o *SitesListDetailsParams) WithTimeout(timeout time.Duration) *SitesListDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sites list details params
func (o *SitesListDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sites list details params
func (o *SitesListDetailsParams) WithContext(ctx context.Context) *SitesListDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sites list details params
func (o *SitesListDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sites list details params
func (o *SitesListDetailsParams) WithHTTPClient(client *http.Client) *SitesListDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sites list details params
func (o *SitesListDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SitesListDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
