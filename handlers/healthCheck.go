package handlers

import (
	"net/http"
)

// HealthCheckHandler godoc
// @Summary Health check
// @Description Check if the service is running.
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Router /health [get]
func HealthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
