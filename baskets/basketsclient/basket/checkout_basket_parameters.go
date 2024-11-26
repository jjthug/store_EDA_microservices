// Code generated by go-swagger; DO NOT EDIT.

package basket

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

// NewCheckoutBasketParams creates a new CheckoutBasketParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCheckoutBasketParams() *CheckoutBasketParams {
	return &CheckoutBasketParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCheckoutBasketParamsWithTimeout creates a new CheckoutBasketParams object
// with the ability to set a timeout on a request.
func NewCheckoutBasketParamsWithTimeout(timeout time.Duration) *CheckoutBasketParams {
	return &CheckoutBasketParams{
		timeout: timeout,
	}
}

// NewCheckoutBasketParamsWithContext creates a new CheckoutBasketParams object
// with the ability to set a context for a request.
func NewCheckoutBasketParamsWithContext(ctx context.Context) *CheckoutBasketParams {
	return &CheckoutBasketParams{
		Context: ctx,
	}
}

// NewCheckoutBasketParamsWithHTTPClient creates a new CheckoutBasketParams object
// with the ability to set a custom HTTPClient for a request.
func NewCheckoutBasketParamsWithHTTPClient(client *http.Client) *CheckoutBasketParams {
	return &CheckoutBasketParams{
		HTTPClient: client,
	}
}

/* CheckoutBasketParams contains all the parameters to send to the API endpoint
   for the checkout basket operation.

   Typically these are written to a http.Request.
*/
type CheckoutBasketParams struct {

	// Body.
	Body *models.CheckoutBasketParamsBody

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the checkout basket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CheckoutBasketParams) WithDefaults() *CheckoutBasketParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the checkout basket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CheckoutBasketParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the checkout basket params
func (o *CheckoutBasketParams) WithTimeout(timeout time.Duration) *CheckoutBasketParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the checkout basket params
func (o *CheckoutBasketParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the checkout basket params
func (o *CheckoutBasketParams) WithContext(ctx context.Context) *CheckoutBasketParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the checkout basket params
func (o *CheckoutBasketParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the checkout basket params
func (o *CheckoutBasketParams) WithHTTPClient(client *http.Client) *CheckoutBasketParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the checkout basket params
func (o *CheckoutBasketParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the checkout basket params
func (o *CheckoutBasketParams) WithBody(body *models.CheckoutBasketParamsBody) *CheckoutBasketParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the checkout basket params
func (o *CheckoutBasketParams) SetBody(body *models.CheckoutBasketParamsBody) {
	o.Body = body
}

// WithID adds the id to the checkout basket params
func (o *CheckoutBasketParams) WithID(id string) *CheckoutBasketParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the checkout basket params
func (o *CheckoutBasketParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CheckoutBasketParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
