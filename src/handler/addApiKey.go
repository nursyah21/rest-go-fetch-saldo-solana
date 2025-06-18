package handler

import (
	"encoding/json"
	"fetch-saldo/src/helper"
	"fetch-saldo/src/models"
	"io"
	"net/http"
)

func AddApiKey(w http.ResponseWriter, r *http.Request) {
	xSecret := r.Header.Get("X-Secret")
	if xSecret != helper.SECRET_KEY {
		http.Error(w, `{"error": "Invalid or Missing X-Secret"}`, http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, `{"error": "Failed to read request body"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	type request struct {
		ApiKey string `json:"api_key"`
	}
	var req request
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	if req.ApiKey == "" {
		http.Error(w, `{"error": "Missing required fields: api_key"}`, http.StatusBadRequest)
		return
	}

	if exists := helper.GetAPIKeyCache(req.ApiKey); exists {
		http.Error(w, `{"error": "Api key already exists"}`, http.StatusBadRequest)
		return
	}

	if exists := models.ApiExist(req.ApiKey); exists {
		helper.SetAPIKeyCache(req.ApiKey, exists)

		http.Error(w, `{"error": "Api key already exists"}`, http.StatusBadRequest)
		return
	}

	if err := models.CreateAPI(req.ApiKey); err != nil {
		http.Error(w, `{"error": "Failed to save API key"}`, http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "API key created successfully",
	}

	json.NewEncoder(w).Encode(response)
}
