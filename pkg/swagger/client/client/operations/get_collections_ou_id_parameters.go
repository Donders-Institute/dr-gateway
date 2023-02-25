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
	"github.com/go-openapi/strfmt"
)

// NewGetCollectionsOuIDParams creates a new GetCollectionsOuIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCollectionsOuIDParams() *GetCollectionsOuIDParams {
	return &GetCollectionsOuIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCollectionsOuIDParamsWithTimeout creates a new GetCollectionsOuIDParams object
// with the ability to set a timeout on a request.
func NewGetCollectionsOuIDParamsWithTimeout(timeout time.Duration) *GetCollectionsOuIDParams {
	return &GetCollectionsOuIDParams{
		timeout: timeout,
	}
}

// NewGetCollectionsOuIDParamsWithContext creates a new GetCollectionsOuIDParams object
// with the ability to set a context for a request.
func NewGetCollectionsOuIDParamsWithContext(ctx context.Context) *GetCollectionsOuIDParams {
	return &GetCollectionsOuIDParams{
		Context: ctx,
	}
}

// NewGetCollectionsOuIDParamsWithHTTPClient creates a new GetCollectionsOuIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCollectionsOuIDParamsWithHTTPClient(client *http.Client) *GetCollectionsOuIDParams {
	return &GetCollectionsOuIDParams{
		HTTPClient: client,
	}
}

/*
GetCollectionsOuIDParams contains all the parameters to send to the API endpoint

	for the get collections ou ID operation.

	Typically these are written to a http.Request.
*/
type GetCollectionsOuIDParams struct {

	/* ID.

	   short name of the organizational unit
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get collections ou ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCollectionsOuIDParams) WithDefaults() *GetCollectionsOuIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get collections ou ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCollectionsOuIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get collections ou ID params
func (o *GetCollectionsOuIDParams) WithTimeout(timeout time.Duration) *GetCollectionsOuIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get collections ou ID params
func (o *GetCollectionsOuIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get collections ou ID params
func (o *GetCollectionsOuIDParams) WithContext(ctx context.Context) *GetCollectionsOuIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get collections ou ID params
func (o *GetCollectionsOuIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get collections ou ID params
func (o *GetCollectionsOuIDParams) WithHTTPClient(client *http.Client) *GetCollectionsOuIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get collections ou ID params
func (o *GetCollectionsOuIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get collections ou ID params
func (o *GetCollectionsOuIDParams) WithID(id string) *GetCollectionsOuIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get collections ou ID params
func (o *GetCollectionsOuIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCollectionsOuIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
