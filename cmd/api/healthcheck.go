// Filename: cmd/api/healthcheck.go
package main

import (
	"encoding/json"
	"net/http"
)

// This is our helper for writing JSON responses. It's a simplified
// version of the helpers.go file from your slides.
func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

// This is the handler function that will be executed when our
// /v1/healthcheck endpoint is requested.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// The data we want to send in the response.
	// Using map[string]any for the envelope pattern from your slides.
	env := map[string]any{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	// Use our helper to send the response.
	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}