// Code generated by go-swagger; DO NOT EDIT.

package providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/getlunaform/lunaform/models"
)

// GetProviderAcceptedCode is the HTTP code returned for type GetProviderAccepted
const GetProviderAcceptedCode int = 202

/*GetProviderAccepted OK

swagger:response getProviderAccepted
*/
type GetProviderAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.ResourceTfProvider `json:"body,omitempty"`
}

// NewGetProviderAccepted creates GetProviderAccepted with default headers values
func NewGetProviderAccepted() *GetProviderAccepted {

	return &GetProviderAccepted{}
}

// WithPayload adds the payload to the get provider accepted response
func (o *GetProviderAccepted) WithPayload(payload *models.ResourceTfProvider) *GetProviderAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get provider accepted response
func (o *GetProviderAccepted) SetPayload(payload *models.ResourceTfProvider) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProviderAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProviderNotFoundCode is the HTTP code returned for type GetProviderNotFound
const GetProviderNotFoundCode int = 404

/*GetProviderNotFound Not Found

swagger:response getProviderNotFound
*/
type GetProviderNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ServerError `json:"body,omitempty"`
}

// NewGetProviderNotFound creates GetProviderNotFound with default headers values
func NewGetProviderNotFound() *GetProviderNotFound {

	return &GetProviderNotFound{}
}

// WithPayload adds the payload to the get provider not found response
func (o *GetProviderNotFound) WithPayload(payload *models.ServerError) *GetProviderNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get provider not found response
func (o *GetProviderNotFound) SetPayload(payload *models.ServerError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProviderNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProviderInternalServerErrorCode is the HTTP code returned for type GetProviderInternalServerError
const GetProviderInternalServerErrorCode int = 500

/*GetProviderInternalServerError Internal Server Error

swagger:response getProviderInternalServerError
*/
type GetProviderInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ServerError `json:"body,omitempty"`
}

// NewGetProviderInternalServerError creates GetProviderInternalServerError with default headers values
func NewGetProviderInternalServerError() *GetProviderInternalServerError {

	return &GetProviderInternalServerError{}
}

// WithPayload adds the payload to the get provider internal server error response
func (o *GetProviderInternalServerError) WithPayload(payload *models.ServerError) *GetProviderInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get provider internal server error response
func (o *GetProviderInternalServerError) SetPayload(payload *models.ServerError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProviderInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
