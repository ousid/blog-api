package helpers

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Data    any    `json:"data.omitempty"`
	Message string `json:"message.omitempty"`
	Error   string `json:"error,omitempty"`
}

func ValidationFailed(w http.ResponseWriter, errs []ValidationError) {
	JSON(w, http.StatusUnprocessableEntity, map[string]any{
		"message": "Validation failed",
		"errors":  errs,
	})
}

func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func Success(w http.ResponseWriter, data any) {
	JSON(w, http.StatusOK, APIResponse{Data: data})
}

func Created(w http.ResponseWriter, data any) {
	JSON(w, http.StatusCreated, APIResponse{Data: data})
}

func NotFound(w http.ResponseWriter, message string) {
	JSON(w, http.StatusNotFound, APIResponse{Error: message})
}

func UnprocessableEntity(w http.ResponseWriter, message string) {
	JSON(w, http.StatusUnprocessableEntity, APIResponse{Error: message})
}

func BadRequest(w http.ResponseWriter, message string) {
	JSON(w, http.StatusBadRequest, APIResponse{Error: message})
}

func ServerError(w http.ResponseWriter, message string) {
	JSON(w, http.StatusInternalServerError, APIResponse{Error: message})
}

func NotContent(w http.ResponseWriter) {
	JSON(w, http.StatusNoContent, nil)
}
