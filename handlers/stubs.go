package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ONSdigital/log.go/log"
)

const (
	florenceToken = "X-Florence-Token"
	serviceToken  = "Authorization"
)

type Identity struct {
	ID          string   `json:"id"`
	Identifier  string   `json:"identifier"`
	Permissions []string `json:"permissions"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Event(r.Context(), "health check invoked", log.INFO)
	return
}

func GetIdentity(identities map[string]*Identity) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Event(r.Context(), "handling get identity request", log.INFO)

		identity, exists := getIdentity(identities, r)
		if !exists {
			writeErrorResponse(w, http.StatusUnauthorized, "unknown identity")
			return
		}

		entity := struct {
			ID         string `json:"id"`
			Identifier string `json:"identifier"`
		}{
			ID:         identity.ID,
			Identifier: identity.Identifier,
		}

		b, err := json.Marshal(entity)
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, "internal server error")
			return
		}

		w.Header().Add("content-type", "application/json")
		w.Write(b)
	}
}

func GetPermissions(identities map[string]*Identity) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Event(r.Context(), "handling get identity request", log.INFO)

		identity, exists := getIdentity(identities, r)
		if !exists {
			writeErrorResponse(w, http.StatusUnauthorized, "unknown identity")
			return
		}

		entity := struct {
			Permissions []string `json:"permissions"`
		}{
			Permissions: identity.Permissions,
		}

		b, err := json.Marshal(entity)
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, "internal server error")
			return
		}

		w.Header().Add("content-type", "application/json")
		w.Write(b)
	}
}

func getIdentity(identities map[string]*Identity, r *http.Request) (*Identity, bool) {
	token := r.Header.Get(florenceToken)
	if token == "" {
		token = strings.TrimPrefix(r.Header.Get(serviceToken), "Bearer ")
	}

	i, exists := identities[token]
	return i, exists
}

func writeErrorResponse(w http.ResponseWriter, status int, message string) {
	body := struct {
		Message string
	}{
		Message: message,
	}

	b, _ := json.Marshal(body)

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(status)
	w.Write(b)
}
