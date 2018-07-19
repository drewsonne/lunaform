package restapi

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/drewsonne/lunaform/backend/identity"
	models "github.com/getlunaform/lunaform-models-go"
	"github.com/drewsonne/lunaform/backend/database"

	"strings"
	operations "github.com/drewsonne/lunaform/server/restapi/operations/modules"
	"github.com/drewsonne/lunaform/server/helpers"
	"github.com/go-openapi/swag"
)

// ListResourcesController provides a list of resources under the identity tag. This is an exploratory read-only endpoint.
var ListTfModulesController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.ListModulesHandlerFunc {
	return operations.ListModulesHandlerFunc(func(params operations.ListModulesParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		modules := make([]*models.ResourceTfModule, 0)
		if err := db.List(DB_TABLE_TF_MODULE, &modules); err != nil {
			return operations.NewListModulesInternalServerError().WithPayload(&models.ServerError{
				StatusCode: HTTP_INTERNAL_SERVER_ERROR,
				Status:     HTTP_INTERNAL_SERVER_ERROR_STATUS,
				Message:    swag.String(err.Error()),
			})
		}

		for _, module := range modules {
			module.GenerateLinks(strings.TrimSuffix(ch.Endpoint, "s"))
			module.Embedded = nil
		}

		return operations.NewListModulesOK().WithPayload(&models.ResponseListTfModules{
			Links: helpers.HalRootRscLinks(ch),
			Embedded: &models.ResourceListTfModule{
				Modules: modules,
			},
		})
	})
}

var CreateTfModuleController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.CreateModuleHandlerFunc {
	return operations.CreateModuleHandlerFunc(func(params operations.CreateModuleParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		tfm := params.TerraformModule
		tfm.ID = idGenerator.MustGenerate()

		if err := db.Create(DB_TABLE_TF_MODULE, tfm.ID, tfm); err != nil {
			return operations.NewCreateModuleBadRequest().WithPayload(&models.ServerError{
				StatusCode: HTTP_INTERNAL_SERVER_ERROR,
				Status:     HTTP_INTERNAL_SERVER_ERROR_STATUS,
				Message:    swag.String(err.Error()),
			})
		}

		tfm.Links = helpers.HalRootRscLinks(ch)
		tfm.Embedded = &models.ResourceListTfStack{
			Stacks: make([]*models.ResourceTfStack, 0),
		}
		return operations.NewCreateModuleCreated().WithPayload(tfm)
	})
}

var GetTfModuleController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.GetModuleHandlerFunc {
	return operations.GetModuleHandlerFunc(func(params operations.GetModuleParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		module := &models.ResourceTfModule{}
		if err := db.Read(DB_TABLE_TF_MODULE, params.ID, module); err != nil {
			return operations.NewGetModuleInternalServerError().WithPayload(&models.ServerError{
				StatusCode: HTTP_INTERNAL_SERVER_ERROR,
				Status:     HTTP_INTERNAL_SERVER_ERROR_STATUS,
				Message:    swag.String(err.Error()),
			})
		} else if module == nil {
			return operations.NewGetModuleNotFound().WithPayload(&models.ServerError{
				StatusCode: HTTP_NOT_FOUND,
				Status:     HTTP_NOT_FOUND_STATUS,
				Message:    swag.String("Could not find module with id '" + params.ID + "'"),
			})
		} else {

			module.Embedded = &models.ResourceListTfStack{
				Stacks: make([]*models.ResourceTfStack, 0),
			}

			module.Links = helpers.HalSelfLink(
				helpers.HalDocLink(nil, ch.OperationID),
				ch.Endpoint,
			)

			return operations.NewGetModuleOK().WithPayload(module)
		}
	})
}
