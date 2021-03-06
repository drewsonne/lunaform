// Code generated by go-swagger; DO NOT EDIT.

package workspaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/getlunaform/lunaform/models"
)

// CreateWorkspaceOKCode is the HTTP code returned for type CreateWorkspaceOK
const CreateWorkspaceOKCode int = 200

/*CreateWorkspaceOK OK

swagger:response createWorkspaceOK
*/
type CreateWorkspaceOK struct {

	/*
	  In: Body
	*/
	Payload *models.ResourceTfWorkspace `json:"body,omitempty"`
}

// NewCreateWorkspaceOK creates CreateWorkspaceOK with default headers values
func NewCreateWorkspaceOK() *CreateWorkspaceOK {

	return &CreateWorkspaceOK{}
}

// WithPayload adds the payload to the create workspace o k response
func (o *CreateWorkspaceOK) WithPayload(payload *models.ResourceTfWorkspace) *CreateWorkspaceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create workspace o k response
func (o *CreateWorkspaceOK) SetPayload(payload *models.ResourceTfWorkspace) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateWorkspaceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateWorkspaceCreatedCode is the HTTP code returned for type CreateWorkspaceCreated
const CreateWorkspaceCreatedCode int = 201

/*CreateWorkspaceCreated Created

swagger:response createWorkspaceCreated
*/
type CreateWorkspaceCreated struct {

	/*
	  In: Body
	*/
	Payload *models.ResourceTfWorkspace `json:"body,omitempty"`
}

// NewCreateWorkspaceCreated creates CreateWorkspaceCreated with default headers values
func NewCreateWorkspaceCreated() *CreateWorkspaceCreated {

	return &CreateWorkspaceCreated{}
}

// WithPayload adds the payload to the create workspace created response
func (o *CreateWorkspaceCreated) WithPayload(payload *models.ResourceTfWorkspace) *CreateWorkspaceCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create workspace created response
func (o *CreateWorkspaceCreated) SetPayload(payload *models.ResourceTfWorkspace) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateWorkspaceCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateWorkspaceBadRequestCode is the HTTP code returned for type CreateWorkspaceBadRequest
const CreateWorkspaceBadRequestCode int = 400

/*CreateWorkspaceBadRequest Bad Request

swagger:response createWorkspaceBadRequest
*/
type CreateWorkspaceBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ServerError `json:"body,omitempty"`
}

// NewCreateWorkspaceBadRequest creates CreateWorkspaceBadRequest with default headers values
func NewCreateWorkspaceBadRequest() *CreateWorkspaceBadRequest {

	return &CreateWorkspaceBadRequest{}
}

// WithPayload adds the payload to the create workspace bad request response
func (o *CreateWorkspaceBadRequest) WithPayload(payload *models.ServerError) *CreateWorkspaceBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create workspace bad request response
func (o *CreateWorkspaceBadRequest) SetPayload(payload *models.ServerError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateWorkspaceBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateWorkspaceNotFoundCode is the HTTP code returned for type CreateWorkspaceNotFound
const CreateWorkspaceNotFoundCode int = 404

/*CreateWorkspaceNotFound Not Found

swagger:response createWorkspaceNotFound
*/
type CreateWorkspaceNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ServerError `json:"body,omitempty"`
}

// NewCreateWorkspaceNotFound creates CreateWorkspaceNotFound with default headers values
func NewCreateWorkspaceNotFound() *CreateWorkspaceNotFound {

	return &CreateWorkspaceNotFound{}
}

// WithPayload adds the payload to the create workspace not found response
func (o *CreateWorkspaceNotFound) WithPayload(payload *models.ServerError) *CreateWorkspaceNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create workspace not found response
func (o *CreateWorkspaceNotFound) SetPayload(payload *models.ServerError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateWorkspaceNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateWorkspaceInternalServerErrorCode is the HTTP code returned for type CreateWorkspaceInternalServerError
const CreateWorkspaceInternalServerErrorCode int = 500

/*CreateWorkspaceInternalServerError Internal Server Error

swagger:response createWorkspaceInternalServerError
*/
type CreateWorkspaceInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ServerError `json:"body,omitempty"`
}

// NewCreateWorkspaceInternalServerError creates CreateWorkspaceInternalServerError with default headers values
func NewCreateWorkspaceInternalServerError() *CreateWorkspaceInternalServerError {

	return &CreateWorkspaceInternalServerError{}
}

// WithPayload adds the payload to the create workspace internal server error response
func (o *CreateWorkspaceInternalServerError) WithPayload(payload *models.ServerError) *CreateWorkspaceInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create workspace internal server error response
func (o *CreateWorkspaceInternalServerError) SetPayload(payload *models.ServerError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateWorkspaceInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
