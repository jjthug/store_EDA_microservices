// Code generated by go-swagger; DO NOT EDIT.

package item

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

// AddItemReader is a Reader for the AddItem structure.
type AddItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddItemDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddItemOK creates a AddItemOK with default headers values
func NewAddItemOK() *AddItemOK {
	return &AddItemOK{}
}

/*
AddItemOK describes a response with status code 200, with default header values.

A successful response.
*/
type AddItemOK struct {
	Payload models.BasketspbAddItemResponse
}

// IsSuccess returns true when this add item o k response has a 2xx status code
func (o *AddItemOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add item o k response has a 3xx status code
func (o *AddItemOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add item o k response has a 4xx status code
func (o *AddItemOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add item o k response has a 5xx status code
func (o *AddItemOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add item o k response a status code equal to that given
func (o *AddItemOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the add item o k response
func (o *AddItemOK) Code() int {
	return 200
}

func (o *AddItemOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/baskets/{id}/addItem][%d] addItemOK %s", 200, payload)
}

func (o *AddItemOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/baskets/{id}/addItem][%d] addItemOK %s", 200, payload)
}

func (o *AddItemOK) GetPayload() models.BasketspbAddItemResponse {
	return o.Payload
}

func (o *AddItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddItemDefault creates a AddItemDefault with default headers values
func NewAddItemDefault(code int) *AddItemDefault {
	return &AddItemDefault{
		_statusCode: code,
	}
}

/*
AddItemDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AddItemDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this add item default response has a 2xx status code
func (o *AddItemDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this add item default response has a 3xx status code
func (o *AddItemDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this add item default response has a 4xx status code
func (o *AddItemDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this add item default response has a 5xx status code
func (o *AddItemDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this add item default response a status code equal to that given
func (o *AddItemDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the add item default response
func (o *AddItemDefault) Code() int {
	return o._statusCode
}

func (o *AddItemDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/baskets/{id}/addItem][%d] addItem default %s", o._statusCode, payload)
}

func (o *AddItemDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/baskets/{id}/addItem][%d] addItem default %s", o._statusCode, payload)
}

func (o *AddItemDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *AddItemDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
