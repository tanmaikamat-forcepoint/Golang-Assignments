package helper

import (
	"encoding/json"
	"net/http"
)

type ErrorMessage struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func SendInvalidAuthError(w http.ResponseWriter) {
	SendErrorWithCustomMessage(w, "Authentication Erro")

}

func SendInvalidBodyError(w *http.ResponseWriter) {

}

func SendErrorWithCustomMessage(w http.ResponseWriter, message string) {

	sendError(w, ErrorMessage{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	})

}

func sendError(w http.ResponseWriter, message ErrorMessage) {
	w.WriteHeader(message.StatusCode)
	json.NewEncoder(w).Encode(message)

}
