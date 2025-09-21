// Filename: cmd/api/routes.go
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	// Initialize the router from the julienschmidt package.
	router := httprouter.New()

	// This is the key part: it maps the GET request for the URL
	// /v1/healthcheck to our healthcheckHandler function.
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	// Return the router.
	return router
}