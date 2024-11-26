// Code generated by go-swagger; DO NOT EDIT.

package order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"EDA_GO/ordering/orderingclient/models"
)

// GetOrderReader is a Reader for the GetOrder structure.
type GetOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetOrderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOrderOK creates a GetOrderOK with default headers values
func NewGetOrderOK() *GetOrderOK {
	return &GetOrderOK{}
}

/* GetOrderOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetOrderOK struct {
	Payload *models.OrderingpbGetOrderResponse
}

// IsSuccess returns true when this get order o k response has a 2xx status code
func (o *GetOrderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get order o k response has a 3xx status code
func (o *GetOrderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get order o k response has a 4xx status code
func (o *GetOrderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get order o k response has a 5xx status code
func (o *GetOrderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get order o k response a status code equal to that given
func (o *GetOrderOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOrderOK) Error() string {
	return fmt.Sprintf("[GET /api/ordering/{id}][%d] getOrderOK  %+v", 200, o.Payload)
}

func (o *GetOrderOK) String() string {
	return fmt.Sprintf("[GET /api/ordering/{id}][%d] getOrderOK  %+v", 200, o.Payload)
}

func (o *GetOrderOK) GetPayload() *models.OrderingpbGetOrderResponse {
	return o.Payload
}

func (o *GetOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrderingpbGetOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderDefault creates a GetOrderDefault with default headers values
func NewGetOrderDefault(code int) *GetOrderDefault {
	return &GetOrderDefault{
		_statusCode: code,
	}
}

/* GetOrderDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetOrderDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the get order default response
func (o *GetOrderDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get order default response has a 2xx status code
func (o *GetOrderDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get order default response has a 3xx status code
func (o *GetOrderDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get order default response has a 4xx status code
func (o *GetOrderDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get order default response has a 5xx status code
func (o *GetOrderDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get order default response a status code equal to that given
func (o *GetOrderDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetOrderDefault) Error() string {
	return fmt.Sprintf("[GET /api/ordering/{id}][%d] getOrder default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrderDefault) String() string {
	return fmt.Sprintf("[GET /api/ordering/{id}][%d] getOrder default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrderDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *GetOrderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
