package handlers

import (
	"encoding/json"
	"http-proxy-task1/models"
	"io"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

var reqStore sync.Map

// PrHandler godoc
// @Summary Proxy request
// @Description Proxy an HTTP request to the specified URL.
// @Tags proxy
// @Accept  json
// @Produce  json
// @Param   request body models.Request true "Proxy request"
// @Success 200 {object} models.Response
// @Router /proxy [post]
func PrHandler(w http.ResponseWriter, r *http.Request) {
	var req models.Request
	//body decode
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Fail", http.StatusBadRequest)
		return
	}

	//newclient !!!
	client := &http.Client{}

	//new req
	prReq, err := http.NewRequest(req.Method, req.URL, nil)
	if err != nil {
		http.Error(w, "Failed to create new request", http.StatusInternalServerError)
	}
	//set headers
	for k, v := range req.Headers {
		prReq.Header.Set(k, v)
	}

	resp, err := client.Do(prReq)
	if err != nil {
		http.Error(w, "Failed to send request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	//body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	headers := make(map[string]string)
	for k, v := range resp.Header {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}

	reqID := generateReqID()

	reqStore.Store(reqID, map[string]interface{}{
		"request":  req,
		"response": string(body),
	})

	response := models.Response{
		ID:      reqID,
		Status:  resp.StatusCode,
		Headers: headers,
		Length:  len(body),
	}

	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode", http.StatusInternalServerError)
		return
	}

}

func generateReqID() string {
	return uuid.New().String()
}
