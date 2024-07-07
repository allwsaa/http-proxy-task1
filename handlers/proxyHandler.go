package handlers

import (
	"encoding/json"
	"http-proxy-task1/models"
	"io"
	"net/http"
	"sync"
)

var reqStore sync.Map

func prHandler(w http.ResponseWriter, r *http.Request) {
	var req models.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Fail", http.StatusBadRequest)
		return
	}

	//chck url
	resp, err := http.Get(req.URL)
	if err != nil {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	//body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

}
