package health

import (
	"net/http"

	"github.com/ONSdigital/log.go/log"
)

func CheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Event(r.Context(), "health check invoked", log.INFO)
	return
}
