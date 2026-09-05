package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func sendJSON(w http.ResponseWriter, resp apiResponse, status int) {
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("error ao fazer marshall de json", "error", err, "response", resp)
		sendJSON(w, apiResponse{Error: "Something Went Wrong"}, http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("erro ao enviar a resposta", "error", err)
		return
	}
}
