// Code generated by go-swagger; DO NOT EDIT.

package workspaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/getlunaform/lunaform/models"
)

// CreateWorkspaceHandlerFunc turns a function with the right signature into a create workspace handler
type CreateWorkspaceHandlerFunc func(CreateWorkspaceParams, *models.ResourceAuthUser) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateWorkspaceHandlerFunc) Handle(params CreateWorkspaceParams, principal *models.ResourceAuthUser) middleware.Responder {
	return fn(params, principal)
}

// CreateWorkspaceHandler interface for that can handle valid create workspace params
type CreateWorkspaceHandler interface {
	Handle(CreateWorkspaceParams, *models.ResourceAuthUser) middleware.Responder
}

// NewCreateWorkspace creates a new http.Handler for the create workspace operation
func NewCreateWorkspace(ctx *middleware.Context, handler CreateWorkspaceHandler) *CreateWorkspace {
	return &CreateWorkspace{Context: ctx, Handler: handler}
}

/*CreateWorkspace swagger:route PUT /tf/workspace/{name} workspaces createWorkspace

Create a Terraform workspace

*/
type CreateWorkspace struct {
	Context *middleware.Context
	Handler CreateWorkspaceHandler
}

func (o *CreateWorkspace) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateWorkspaceParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.ResourceAuthUser
	if uprinc != nil {
		principal = uprinc.(*models.ResourceAuthUser) // this is really a models.ResourceAuthUser, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
