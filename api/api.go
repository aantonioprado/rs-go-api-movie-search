package api

import (
	"aantonioprado/rs-go-api-movie-search/omdb"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Response struct {
	Error  string `json:"error,omitempty"`
	Data   any    `json:"data,omitempty"`
	Status string `json:"status,omitempty"`
}

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("Failed to marshal response", "error", err)
		sendJSON(
			w,
			Response{Error: "something went wrong"},
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("Failed to write response to client", "error", err)
		return
	}
}

func NewHandler(apiKey string) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/", handleGetSearchMovie(apiKey))

	return r
}

func handleGetSearchMovie(apiKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("s")
		res, err := omdb.Search(apiKey, search)

		if err != nil {
			sendJSON(
				w,
				Response{Error: "something wrong with omdb"},
				http.StatusBadGateway,
			)
			return
		}

		sendJSON(
			w,
			Response{
				Data:   res,
				Status: "OK",
			},
			http.StatusOK,
		)
	}
}
