package identity

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

type Model struct {
	ID         string `json:"id"`
	Identifier string `json:"identifier"`
}

func GetIdentityHandler(identities map[string]Model) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Event(r.Context(), "handling get identity request", log.INFO)

		id := r.Header.Get(florenceToken)
		if id == "" {
			id = strings.TrimPrefix(r.Header.Get(serviceToken), "Bearer ")
		}

		i, exists := identities[id]
		if !exists {
			writeErrorResponse(w, http.StatusUnauthorized, "unknown identity")
			return
		}

		b, err := json.Marshal(i)
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, "internal server error")
			return
		}

		w.Header().Add("content-type", "application/json")
		w.Write(b)
	}
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
