package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/api/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/api/v1/contributions", app.createContributionHandler)
	router.HandlerFunc(http.MethodGet, "/api/v1/contributions/:id", app.showContributionHandler)

	return router
}
