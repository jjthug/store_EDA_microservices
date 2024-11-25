// Code generated by go-swagger; DO NOT EDIT.

package basket

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"EDA_GO/baskets/basketsclient/models"
)

// CheckoutBasketReader is a Reader for the CheckoutBasket structure.
type CheckoutBasketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckoutBasketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckoutBasketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCheckoutBasketDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCheckoutBasketOK creates a CheckoutBasketOK with default headers values
func NewCheckoutBasketOK() *CheckoutBasketOK {
	return &CheckoutBasketOK{}
}

/*
CheckoutBasketOK describes a response with status code 200, with default header values.

A successful response.
*/
type CheckoutBasketOK struct {
	Payload models.BasketspbCheckoutBasketResponse
}

// IsSuccess returns true when this checkout basket o k response has a 2xx status code
func (o *CheckoutBasketOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this checkout basket o k response has a 3xx status code
func (o *CheckoutBasketOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this checkout basket o k response has a 4xx status code
func (o *CheckoutBasketOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this checkout basket o k response has a 5xx status code
func (o *CheckoutBasketOK) IsServerError() bool {
	return false
}

// IsCode returns true when this checkout basket o k response a status code equal to that given
func (o *CheckoutBasketOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the checkout basket o k response
func (o *CheckoutBasketOK) Code() int {
	return 200
}

func (o *CheckoutBasketOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/baskets/{id}/checkout][%d] checkoutBasketOK %s", 200, payload)
}

func (o *CheckoutBasketOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/baskets/{id}/checkout][%d] checkoutBasketOK %s", 200, payload)
}

func (o *CheckoutBasketOK) GetPayload() models.BasketspbCheckoutBasketResponse {
	return o.Payload
}

func (o *CheckoutBasketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckoutBasketDefault creates a CheckoutBasketDefault with default headers values
func NewCheckoutBasketDefault(code int) *CheckoutBasketDefault {
	return &CheckoutBasketDefault{
		_statusCode: code,
	}
}

/*
CheckoutBasketDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CheckoutBasketDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this checkout basket default response has a 2xx status code
func (o *CheckoutBasketDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this checkout basket default response has a 3xx status code
func (o *CheckoutBasketDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this checkout basket default response has a 4xx status code
func (o *CheckoutBasketDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this checkout basket default response has a 5xx status code
func (o *CheckoutBasketDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this checkout basket default response a status code equal to that given
func (o *CheckoutBasketDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the checkout basket default response
func (o *CheckoutBasketDefault) Code() int {
	return o._statusCode
}

func (o *CheckoutBasketDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/baskets/{id}/checkout][%d] checkoutBasket default %s", o._statusCode, payload)
}

func (o *CheckoutBasketDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/baskets/{id}/checkout][%d] checkoutBasket default %s", o._statusCode, payload)
}

func (o *CheckoutBasketDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *CheckoutBasketDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
