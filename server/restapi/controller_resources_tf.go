package restapi

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/drewsonne/terraform-server/backend/identity"
	"github.com/drewsonne/terraform-server/server/models"
	"github.com/drewsonne/terraform-server/server/restapi/operations/tf"
	"github.com/drewsonne/terraform-server/backend/database"
	"github.com/pborman/uuid"
)

// ListResourcesController provides a list of resources under the identity tag. This is an exploratory read-only endpoint.
var ListTfModulesController = func(idp identity.Provider, ch ContextHelper, db database.Database) tf.ListModulesHandlerFunc {
	return tf.ListModulesHandlerFunc(func(params tf.ListModulesParams) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		records, err := db.List("tf-modules")
		if err != nil {
			var statuscode int64 = 500
			return tf.NewListModulesInternalServerError().WithPayload(&models.ServerError{
				StatusCode: &statuscode,
				Status:     String("Internal Server Error"),
				Message:    String(err.Error()),
			})
		}

		modules := make([]*models.ResponseTfModule, len(records))
		for i, record := range records {
			modules[i] = &models.ResponseTfModule{
				Name:  &record.Value,
				Links: halSelfLink(ch.FQEndpoint + "/" + record.Value),
			}
		}

		return tf.NewListModulesOK().WithPayload(&models.ResponseListTfModules{
			Links:    halRootRscLinks(ch),
			Embedded: modules,
		})
	})
}

var CreateTfModuleController = func(idp identity.Provider, ch ContextHelper, db database.Database) tf.CreateModuleHandlerFunc {
	return tf.CreateModuleHandlerFunc(func(params tf.CreateModuleParams) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		newId := uuid.New()
		db.Create("tf-module", newId, params)

		response := &models.ResponseTfModule{
			Links: halRootRscLinks(ch),
			VcsID: String(newId),
		}
		if params.TerraformModule == nil {
			return tf.NewCreateModuleBadRequest()
		} else {
			response.Name = params.TerraformModule.Name
			response.Type = *params.TerraformModule.Type
			return tf.NewCreateModuleCreated().WithPayload(response)
		}

	})
}
