package main

import (
	"context"
	"net/http"
	"os"

	dphttp "github.com/ONSdigital/dp-net/http"
	"github.com/ONSdigital/dp-zebedee-api-stub/config"
	"github.com/ONSdigital/dp-zebedee-api-stub/handlers"
	"github.com/ONSdigital/log.go/log"
	"github.com/gorilla/mux"
)

func main() {
	if err := run(); err != nil {
		log.Event(context.Background(), "unexpected application error", log.Error(err), log.ERROR)
		os.Exit(1)
	}
}

func run() error {
	log.Namespace = "dp-zebedee-api-stub"

	cfg, err := config.Get()
	if err != nil {
		return err
	}

	log.Event(context.Background(), "config", log.INFO, log.Data{
		"config": cfg,
	})

	r := mux.NewRouter()
	r.HandleFunc("/health", handlers.HealthCheck).Methods(http.MethodGet)
	r.HandleFunc("/identity", handlers.GetIdentity(cfg.Identities)).Methods(http.MethodGet)
	r.HandleFunc("/serviceInstancePermissions", handlers.GetPermissions(cfg.Identities)).Methods(http.MethodGet)
	r.HandleFunc("/serviceDatasetPermissions", handlers.GetPermissions(cfg.Identities)).Methods(http.MethodGet)
	r.HandleFunc("/userInstancePermissions", handlers.GetPermissions(cfg.Identities)).Methods(http.MethodGet)
	r.HandleFunc("/userDatasetPermissions", handlers.GetPermissions(cfg.Identities)).Methods(http.MethodGet)

	log.Event(context.Background(), "starting stub", log.INFO)
	if err := dphttp.NewServer(cfg.BindAddr, r).ListenAndServe(); err != nil {
		return err
	}

	return nil
}
