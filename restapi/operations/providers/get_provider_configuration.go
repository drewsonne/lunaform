// Code generated by go-swagger; DO NOT EDIT.

package providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/getlunaform/lunaform/models"
)

// GetProviderConfigurationHandlerFunc turns a function with the right signature into a get provider configuration handler
type GetProviderConfigurationHandlerFunc func(GetProviderConfigurationParams, *models.ResourceAuthUser) middleware.Responder

// Handle executing the request and returning a response
func (fn GetProviderConfigurationHandlerFunc) Handle(params GetProviderConfigurationParams, principal *models.ResourceAuthUser) middleware.Responder {
	return fn(params, principal)
}

// GetProviderConfigurationHandler interface for that can handle valid get provider configuration params
type GetProviderConfigurationHandler interface {
	Handle(GetProviderConfigurationParams, *models.ResourceAuthUser) middleware.Responder
}

// NewGetProviderConfiguration creates a new http.Handler for the get provider configuration operation
func NewGetProviderConfiguration(ctx *middleware.Context, handler GetProviderConfigurationHandler) *GetProviderConfiguration {
	return &GetProviderConfiguration{Context: ctx, Handler: handler}
}

/*GetProviderConfiguration swagger:route GET /tf/provider/{provider-name}/configuration/{id} providers getProviderConfiguration

Get Configuration for Provider

*/
type GetProviderConfiguration struct {
	Context *middleware.Context
	Handler GetProviderConfigurationHandler
}

func (o *GetProviderConfiguration) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetProviderConfigurationParams()

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
