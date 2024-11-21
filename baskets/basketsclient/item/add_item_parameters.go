// Code generated by go-swagger; DO NOT EDIT.

package item

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

	"EDA_GO/baskets/basketsclient/models"
)

// NewAddItemParams creates a new AddItemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddItemParams() *AddItemParams {
	return &AddItemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddItemParamsWithTimeout creates a new AddItemParams object
// with the ability to set a timeout on a request.
func NewAddItemParamsWithTimeout(timeout time.Duration) *AddItemParams {
	return &AddItemParams{
		timeout: timeout,
	}
}

// NewAddItemParamsWithContext creates a new AddItemParams object
// with the ability to set a context for a request.
func NewAddItemParamsWithContext(ctx context.Context) *AddItemParams {
	return &AddItemParams{
		Context: ctx,
	}
}

// NewAddItemParamsWithHTTPClient creates a new AddItemParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddItemParamsWithHTTPClient(client *http.Client) *AddItemParams {
	return &AddItemParams{
		HTTPClient: client,
	}
}

/*
AddItemParams contains all the parameters to send to the API endpoint

	for the add item operation.

	Typically these are written to a http.Request.
*/
type AddItemParams struct {

	// Body.
	Body *models.BasketServiceAddItemBody

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddItemParams) WithDefaults() *AddItemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddItemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add item params
func (o *AddItemParams) WithTimeout(timeout time.Duration) *AddItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add item params
func (o *AddItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add item params
func (o *AddItemParams) WithContext(ctx context.Context) *AddItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add item params
func (o *AddItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add item params
func (o *AddItemParams) WithHTTPClient(client *http.Client) *AddItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add item params
func (o *AddItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add item params
func (o *AddItemParams) WithBody(body *models.BasketServiceAddItemBody) *AddItemParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add item params
func (o *AddItemParams) SetBody(body *models.BasketServiceAddItemBody) {
	o.Body = body
}

// WithID adds the id to the add item params
func (o *AddItemParams) WithID(id string) *AddItemParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the add item params
func (o *AddItemParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AddItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
