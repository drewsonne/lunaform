package restapi

import (
	"github.com/drewsonne/lunaform/backend/identity"
	"github.com/drewsonne/lunaform/backend/database"
	"github.com/go-openapi/runtime/middleware"
	operations "github.com/drewsonne/lunaform/server/restapi/operations/workspaces"
	"github.com/drewsonne/lunaform/server/models"
	"strings"
	"github.com/drewsonne/lunaform/server/helpers"
)

var ListTfWorkspacesController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.ListWorkspacesHandlerFunc {
	return operations.ListWorkspacesHandlerFunc(func(params operations.ListWorkspacesParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		workspaces := []*models.ResourceTfWorkspace{}
		err := db.List(DB_TABLE_TF_WORKSPACE, &workspaces)
		if err != nil {
			return operations.NewListWorkspacesInternalServerError().WithPayload(&models.ServerError{
				StatusCode: HTTP_INTERNAL_SERVER_ERROR,
				Status:     HTTP_INTERNAL_SERVER_ERROR_STATUS,
				Message:    helpers.String(err.Error()),
			})
		}

		for _, workspace := range workspaces {
			workspace.Links = helpers.HalSelfLink(strings.TrimSuffix(ch.FQEndpoint, "s") + "/" + *workspace.Name)
		}

		return operations.NewListWorkspacesOK().WithPayload(&models.ResponseListTfWorkspaces{
			Links: helpers.HalRootRscLinks(ch),
			Embedded: &models.ResourceListTfWorkspace{
				Workspaces: workspaces,
			},
		})
	})
}

var CreateTfWorkspaceController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.CreateWorkspaceHandlerFunc {
	return operations.CreateWorkspaceHandlerFunc(func(params operations.CreateWorkspaceParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		tfw := params.TerraformWorkspace
		tfw.Name = helpers.String(params.Name)
		tfw.Modules = []*models.ResourceTfModule{}

		existingWorkspace := models.ResourceTfWorkspace{}

		halRscLinks := helpers.HalSelfLink(strings.TrimSuffix(ch.FQEndpoint, "s") + "/" + params.Name)
		halRscLinks.Doc = helpers.HalDocLink(ch).Doc

		if err := db.Read(DB_TABLE_TF_WORKSPACE, params.Name, &existingWorkspace); err != nil {
			if _, isNotFound := err.(database.RecordDoesNotExistError); isNotFound {
				if err := db.Create(DB_TABLE_TF_WORKSPACE, params.Name, tfw); err == nil {
					r = operations.NewCreateWorkspaceCreated().WithPayload(tfw)
				} else {
					r = operations.NewCreateWorkspaceBadRequest().WithPayload(&models.ServerError{
						StatusCode: HTTP_BAD_REQUEST,
						Status:     HTTP_BAD_REQUEST_STATUS,
						Message:    helpers.String(err.Error()),
					})
				}
			} else {
				r = operations.NewCreateWorkspaceInternalServerError().WithPayload(&models.ServerError{
					StatusCode: HTTP_INTERNAL_SERVER_ERROR,
					Status:     HTTP_INTERNAL_SERVER_ERROR_STATUS,
					Message:    helpers.String(err.Error()),
				})
			}
		} else {
			if err := db.Update(DB_TABLE_TF_WORKSPACE, params.Name, existingWorkspace); err != nil {
				r = operations.NewCreateWorkspaceInternalServerError().WithPayload(&models.ServerError{
					StatusCode: HTTP_INTERNAL_SERVER_ERROR,
					Status:     HTTP_INTERNAL_SERVER_ERROR_STATUS,
					Message:    helpers.String(err.Error()),
				})
			} else {
				r = operations.NewCreateWorkspaceOK().WithPayload(tfw)
			}
		}

		return
	})
}

var GetTfWorkspaceController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.DescribeWorkspaceHandlerFunc {
	return operations.DescribeWorkspaceHandlerFunc(func(params operations.DescribeWorkspaceParams, p *models.Principal) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		workspace := &models.ResourceTfWorkspace{}
		if err := db.Read(DB_TABLE_TF_WORKSPACE, params.Name, workspace); err != nil {
			return operations.NewDescribeWorkspaceInternalServerError().WithPayload(&models.ServerError{
				StatusCode: HTTP_INTERNAL_SERVER_ERROR,
				Status:     HTTP_INTERNAL_SERVER_ERROR_STATUS,
				Message:    helpers.String(err.Error()),
			})
		} else if workspace == nil {
			return operations.NewDescribeWorkspaceNotFound().WithPayload(&models.ServerError{
				StatusCode: HTTP_NOT_FOUND,
				Status:     HTTP_NOT_FOUND_STATUS,
				Message:    helpers.String("Could not find workspace with name '" + params.Name + "'"),
			})
		} else {
			workspace.Links = helpers.HalSelfLink(ch.FQEndpoint)
			workspace.Links.Doc = helpers.HalDocLink(ch).Doc
			return operations.NewDescribeWorkspaceOK().WithPayload(workspace)
		}
	})
}
