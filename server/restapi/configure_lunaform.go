package restapi

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/drewsonne/lunaform/server/restapi/operations"

	"github.com/drewsonne/lunaform/backend/database"
	"github.com/drewsonne/lunaform/backend/identity"
	"log"
	"net/http/httputil"
	"net/http/httptest"
	"os"
	"github.com/go-openapi/swag"
	"github.com/drewsonne/lunaform/server/models"
	"github.com/pborman/uuid"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target ../server --name TerraformServer --spec ../swagger.yml --principal models.Principal

var version string

func configureAPI(api *operations.LunaformAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError
	//api.BearerAuthenticator =

	var idp identity.Provider
	var dbDriver database.Driver
	var err error

	cfg, err := parseCliConfiguration()
	if err != nil {
		panic(err)
	}

	switch cfg.Backend.DatabaseType {
	case "memory":
		dbDriver, err = database.NewMemoryDBDriver()
	default:
		err = fmt.Errorf("unexpected Database type: '%s'", cfg.Backend.DatabaseType)
	}

	if err != nil {
		panic(err)
	}

	db := database.NewDatabaseWithDriver(dbDriver)

	switch cfg.Backend.IdentityType {
	case "memory":
		idp = identity.NewMemoryIdentityProvider()
	case "database":
		idp, err = identity.NewDatabaseIdentityProvider(db)
	default:
		panic(fmt.Sprintf("Unexpected Identity Provider type: '%s'", cfg.Backend.IdentityType))
	}

	if err != nil {
		panic(err)
	}

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	oh := NewContextHelper(api.Context())

	api.APIKeyAuth = func(s string) (p *models.Principal, err error) {
		user := models.ResourceAuthUser{}
		if err = db.Read("lf-auth-apikeys", s, &user); err != nil {
			return
		}

		return &models.Principal{user}, nil
	}

	configureRootUser(&db)

	// Controllers for /
	api.ResourcesListResourceGroupsHandler = ListResourceGroupsController(idp, oh)

	// Controllers for /{group}
	api.ResourcesListResourcesHandler = ListResourcesController(idp, oh)

	// Controllers for /tf/modules
	api.ModulesListModulesHandler = ListTfModulesController(idp, oh, db)
	api.ModulesCreateModuleHandler = CreateTfModuleController(idp, oh, db)
	api.ModulesGetModuleHandler = GetTfModuleController(idp, oh, db)

	// Controllers for /tf/stacks
	api.StacksDeployStackHandler = CreateTfStackController(idp, oh, db)
	api.StacksListStacksHandler = ListTfStacksController(idp, oh, db)
	api.StacksGetStackHandler = GetTfStackController(idp, oh, db)

	// Controllers for /tf/workspaces
	api.WorkspacesListWorkspacesHandler = ListTfWorkspacesController(idp, oh, db)
	api.WorkspacesCreateWorkspaceHandler = CreateTfWorkspaceController(idp, oh, db)

	api.ServerShutdown = func() {
		dbDriver.Close()
	}

	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "Version",
			Options: map[string]string{
				"one": "two",
			},
		},
	}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {

}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {

	handler = logRequest(handler)
	debug := os.Getenv("DEBUG")
	if debug == "1" {
		return logResponse("lunaform", handler)
	} else {
		return handler
	}
}

func logRequest(handler http.Handler) http.Handler {
	//return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	//	handler.ServeHTTP(w, r)
	//})
	return handler
}

func logResponse(prefix string, h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Save a copy of this request for debugging.
		requestDump, err := httputil.DumpRequest(r, false)
		if err != nil {
			log.Println(err)
		}
		log.Println(prefix, string(requestDump))

		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, r)

		dump, err := httputil.DumpResponse(rec.Result(), false)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(prefix, string(dump))

		// we copy the captured response headers to our new response
		for k, v := range rec.Header() {
			w.Header()[k] = v
		}

		// grab the captured response body
		data := rec.Body.Bytes()
		w.WriteHeader(rec.Code)

		w.Write(data)
	}
}

func configureRootUser(db *database.Database) (err error) {
	userRecords := []*models.ResourceAuthUser{}
	if err = db.List("lf-auth-user", &userRecords); err != nil {
		return
	}

	var foundAdmin bool
	for _, user := range userRecords {
		if user.Name == "admin" {
			foundAdmin = true
		}
	}

	if !foundAdmin || cliconfig.AdminApiKey != "" {

		if cliconfig.AdminApiKey == "" {
			cliconfig.AdminApiKey = uuid.New()
		}

		adminUser := &models.ResourceAuthUser{
			Name:      "Adminstrator",
			Shortname: "admin",
			Groups:    []string{"admin"},
			ID:        uuid.New(),
			APIKeys:   []string{cliconfig.AdminApiKey},
		}
		if err = db.Create("lf-auth-user", adminUser.ID, adminUser); err != nil {
			return
		}
		for _, keys := range adminUser.APIKeys {
			if err = db.Create("lf-auth-apikeys", keys, adminUser); err != nil {
				return
			}
			fmt.Printf("Generated api-key for admin user '%s'", keys)
		}
	}

	return
}
