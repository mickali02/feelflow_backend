// Filename: cmd/api/server.go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func (app *application) serve() error {
	// This is the server setup code from your slides (Web Server (1)).
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(), // This points to our router
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.logger.Info("starting server", "addr", srv.Addr, "env", app.config.env)

	return srv.ListenAndServe()
}