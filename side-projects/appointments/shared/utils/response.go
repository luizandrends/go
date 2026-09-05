package utils

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type ApiResponse[T any] struct {
	Error string `json:"error,omitempty"`
	Data  T      `json:"data,omitempty"`
}

func SendJSON[T any](w http.ResponseWriter, resp ApiResponse[T], status int) {
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("error ao fazer marshall de json", "error", err, "response", resp)
		SendJSON(w, ApiResponse[any]{Error: "Something Went Wrong"}, http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("erro ao enviar a resposta", "error", err)
		return
	}
}
