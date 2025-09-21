// Filename: cmd/api/main.go
package main

import (
	"flag"
	"log/slog"
	"os"
)

// This version number is used in the healthcheck response.
const version = "1.0.0"

// This holds all configuration settings for our application.
// We're keeping the structure your teacher showed.
type config struct {
	port int
	env  string
}

// This holds dependencies that we'll share across the application.
// For now, it's just the config and the logger. This is the
// Dependency Injection (DI) pattern from your slides.
type application struct {
	config config
	logger *slog.Logger
}

func main() {
	// Declare an instance of the config struct.
	var cfg config

	// Read the port and env from command-line flags.
	// Default to port 4000 and "development" environment.
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// Set up the logger.
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Create an instance of our application struct, containing the config and logger.
	app := &application{
		config: cfg,
		logger: logger,
	}

	// Call the serve() method to start the server.
	// We pass the error handling to another file to keep main() clean.
	err := app.serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}