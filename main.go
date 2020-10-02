package main

import (
	"context"
	"net/http"
	"os"

	dphttp "github.com/ONSdigital/dp-net/http"
	"github.com/ONSdigital/dp-zebedee-api-stub/health"
	"github.com/ONSdigital/dp-zebedee-api-stub/identity"
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

	r := mux.NewRouter()
	r.HandleFunc("/identity", identity.GetHandler).Methods(http.MethodGet)
	r.HandleFunc("/health", health.CheckHandler).Methods(http.MethodGet)

	if err := dphttp.NewServer(":8082", r).ListenAndServe(); err != nil {
		return err
	}

	return nil
}
