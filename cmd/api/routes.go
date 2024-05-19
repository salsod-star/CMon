package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/api/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/api/v1/contributions", app.createContributionHandler)
	router.HandlerFunc(http.MethodGet, "/api/v1/contributions/:id", app.showContributionHandler)
	router.HandlerFunc(http.MethodPut, "/api/v1/contributions/:id", app.updateContributionHandler)
	router.HandlerFunc(http.MethodDelete, "/api/v1/contributions/:id", app.deleteContributionHandler)

	return router
}
