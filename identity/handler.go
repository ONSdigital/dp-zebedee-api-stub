package identity

import (
	"encoding/json"
	"net/http"
	"strings"
)

const (
	florenceToken = "X-Florence-Token"
	serviceToken  = "Authorization"
)

var (
	identities = map[string]string{
		"7e0d1238-cf25-4239-adfb-7f1a460a0580": "Weyland-Yutani Corporation",
	}
)

type Identity struct {
	ID         string `json:"id"`
	Identifier string `json:"identifier"`
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
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
