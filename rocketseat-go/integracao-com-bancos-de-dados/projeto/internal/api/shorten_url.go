package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/luizandrends/projeto/internal/store"
)

type shoretnURLRequest struct {
	URL string `json:"url"`
}

type shortenURLResponse struct {
	Code string `json:"code"`
}

func handleShortenURL(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body shoretnURLRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJSON(
				w,
				apiResponse{Error: "invalid body"},
				http.StatusUnprocessableEntity,
			)
		}

		if _, err := url.Parse(body.URL); err != nil {
			sendJSON(
				w,
				apiResponse{Error: "invalid url passed"},
				http.StatusBadRequest,
			)
		}
		code, err := store.SaveShortenedURL(r.Context(), body.URL)

		if err != nil {
			slog.Error("Failed to create code")
			sendJSON(
				w,
				apiResponse{Error: "invalid body"},
				http.StatusInternalServerError,
			)
			return
		}
		sendJSON(w, apiResponse{Data: shortenURLResponse{Code: code}}, http.StatusCreated)
	}
}
