// Filename: cmd/api/routes.go
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// This new function is a middleware that adds the CORS header.
func (app *application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// IMPORTANT: This line allows requests specifically from your Flutter app's origin.
		// If your Flutter app's port changes, you'll need to update this value.
			w.Header().Set("Access-Control-Allow-Origin", "*")
		// This passes the request along to the next handler (our router).
		next.ServeHTTP(w, r)
	})
}

func (app *application) routes() http.Handler {
	// Initialize the router from the julienschmidt package.
	router := httprouter.New()

	// This is the key part: it maps the GET request for the URL
	// /v1/healthcheck to our healthcheckHandler function.
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	// We wrap our router with the CORS middleware before returning it.
	// This ensures every request gets the CORS header added.
	return app.enableCORS(router)
}
