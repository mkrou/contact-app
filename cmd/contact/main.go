package main

import (
	"context"
	"contact/internal/app/contacts"
	"contact/internal/app/transport"
	"contact/internal/pkg/config"
	"contact/internal/pkg/postgres"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"net/http"
	"os"
)

func main() {
	//initialize logger
	logger := log.NewJSONLogger(os.Stderr)

	//load configuration from env variables or .env file
	config, err := config.NewConfig(context.Background(), logger)
	if err != nil {
		_ = level.Error(logger).Log("exit", err)
		return
	}
	_ = level.Info(logger).Log("msg", "config has parsed")

	//initialize db connection
	db := postgres.NewPostgresDB(config, logger)
	defer db.Close()

	//initialize repositories
	contactPGRepository := contacts.NewContactPostgresRepository(db)

	//initialize services
	contactService := contacts.NewContactService(contactPGRepository)

	//initialize endpoints
	contactEndpoints := contacts.NewContactEndpoints(contactService)

	//initialize transport
	transport := transport.NewTransport(contactEndpoints)

	//start http server
	_ = level.Info(logger).Log("msg", "server has started")
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), transport); err != nil {
		_ = level.Error(logger).Log("exit", err)
	}
}
