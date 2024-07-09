package handlers

import (
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
