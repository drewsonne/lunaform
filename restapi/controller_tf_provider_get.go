package restapi

import (
	"github.com/getlunaform/lunaform/backend/identity"
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/backend/database"
	operation "github.com/getlunaform/lunaform/restapi/operations/providers"
	"github.com/getlunaform/lunaform/models"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
)

func GetTfProviderController(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operation.GetProviderHandlerFunc {
	return func(params operation.GetProviderParams, user *models.ResourceAuthUser) middleware.Responder {
		ch.SetRequest(params.HTTPRequest)

		prov, code, err := buildGetTfProviderResponse(db, params.Name)
		if err != nil {
			return NewServerError(code, err.Error())
		}

		prov.Links = helpers.HalAddCuries(ch, helpers.HalSelfLink(
			helpers.HalDocLink(nil, ch.OperationID),
			ch.Endpoint,
		))

		return operation.NewGetProviderAccepted().WithPayload(prov)
	}
}

func buildGetTfProviderResponse(db database.Database, providerName string) (provider *models.ResourceTfProvider, errCode int, err error) {
	provider = &models.ResourceTfProvider{}
	if err := db.Read(DB_TABLE_TF_PROVIDER, providerName, provider); err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return
}
