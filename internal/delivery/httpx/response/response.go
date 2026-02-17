package response

import (
	"encoding/json"
	"net/http"
)

type successResponse struct {
	Success bool           `json:"success" binding:"required"`
	Data    any            `json:"data" binding:"required"`
	Meta    map[string]any `json:"meta,omitempty"`
}

type errorResponse struct {
	Success bool   `json:"success" binding:"required"`
	Error   string `json:"error"`
}

func toJson(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func SUCCESS(w http.ResponseWriter, status int, data any, meta map[string]any) {
	toJson(w, status, successResponse{
		Success: true,
		Data:    data,
		Meta:    meta,
	})
}

func FAILED(w http.ResponseWriter, status int, err error) {
	toJson(w, status, errorResponse{
		Success: false,
		Error:   err.Error(),
	})
}
