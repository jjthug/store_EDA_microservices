// Code generated by go-swagger; DO NOT EDIT.

package customer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"EDA_GO/customers/customersclient/models"
)

// DisableCustomerReader is a Reader for the DisableCustomer structure.
type DisableCustomerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DisableCustomerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDisableCustomerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDisableCustomerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDisableCustomerOK creates a DisableCustomerOK with default headers values
func NewDisableCustomerOK() *DisableCustomerOK {
	return &DisableCustomerOK{}
}

/* DisableCustomerOK describes a response with status code 200, with default header values.

A successful response.
*/
type DisableCustomerOK struct {
	Payload models.CustomerspbDisableCustomerResponse
}

// IsSuccess returns true when this disable customer o k response has a 2xx status code
func (o *DisableCustomerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this disable customer o k response has a 3xx status code
func (o *DisableCustomerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable customer o k response has a 4xx status code
func (o *DisableCustomerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this disable customer o k response has a 5xx status code
func (o *DisableCustomerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this disable customer o k response a status code equal to that given
func (o *DisableCustomerOK) IsCode(code int) bool {
	return code == 200
}

func (o *DisableCustomerOK) Error() string {
	return fmt.Sprintf("[PUT /api/customers/{id}/disable][%d] disableCustomerOK  %+v", 200, o.Payload)
}

func (o *DisableCustomerOK) String() string {
	return fmt.Sprintf("[PUT /api/customers/{id}/disable][%d] disableCustomerOK  %+v", 200, o.Payload)
}

func (o *DisableCustomerOK) GetPayload() models.CustomerspbDisableCustomerResponse {
	return o.Payload
}

func (o *DisableCustomerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDisableCustomerDefault creates a DisableCustomerDefault with default headers values
func NewDisableCustomerDefault(code int) *DisableCustomerDefault {
	return &DisableCustomerDefault{
		_statusCode: code,
	}
}

/* DisableCustomerDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DisableCustomerDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the disable customer default response
func (o *DisableCustomerDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this disable customer default response has a 2xx status code
func (o *DisableCustomerDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this disable customer default response has a 3xx status code
func (o *DisableCustomerDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this disable customer default response has a 4xx status code
func (o *DisableCustomerDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this disable customer default response has a 5xx status code
func (o *DisableCustomerDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this disable customer default response a status code equal to that given
func (o *DisableCustomerDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DisableCustomerDefault) Error() string {
	return fmt.Sprintf("[PUT /api/customers/{id}/disable][%d] DisableCustomer default  %+v", o._statusCode, o.Payload)
}

func (o *DisableCustomerDefault) String() string {
	return fmt.Sprintf("[PUT /api/customers/{id}/disable][%d] DisableCustomer default  %+v", o._statusCode, o.Payload)
}

func (o *DisableCustomerDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *DisableCustomerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}