// Code generated by go-swagger; DO NOT EDIT.

package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ListResourceGroupsHandlerFunc turns a function with the right signature into a list resource groups handler
type ListResourceGroupsHandlerFunc func(ListResourceGroupsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListResourceGroupsHandlerFunc) Handle(params ListResourceGroupsParams) middleware.Responder {
	return fn(params)
}

// ListResourceGroupsHandler interface for that can handle valid list resource groups params
type ListResourceGroupsHandler interface {
	Handle(ListResourceGroupsParams) middleware.Responder
}

// NewListResourceGroups creates a new http.Handler for the list resource groups operation
func NewListResourceGroups(ctx *middleware.Context, handler ListResourceGroupsHandler) *ListResourceGroups {
	return &ListResourceGroups{Context: ctx, Handler: handler}
}

/*ListResourceGroups swagger:route GET / resources listResourceGroups

List the root resource groups

*/
type ListResourceGroups struct {
	Context *middleware.Context
	Handler ListResourceGroupsHandler
}

func (o *ListResourceGroups) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListResourceGroupsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
