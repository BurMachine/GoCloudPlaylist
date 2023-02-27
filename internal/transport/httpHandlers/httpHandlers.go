package httpHandlers

import (
	"GoCloudPlaylist/internal/playlist"
	"context"
	"fmt"
	gorilla "github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

type HttpHandlers struct {
	Mux *gorilla.Router
	Pl  *playlist.Playlist
}

func (h *HttpHandlers) Register() {
	h.Mux.HandleFunc("/play", h.Middleware(h.PlaySong)).Methods(http.MethodGet)
	h.Mux.HandleFunc("/pause", h.Middleware(h.PauseSong)).Methods(http.MethodGet)
	h.Mux.HandleFunc("/next_song", h.Middleware(h.NextSong)).Methods(http.MethodGet)
	h.Mux.HandleFunc("/prev_song", h.Middleware(h.PrevSong)).Methods(http.MethodGet)
	h.Mux.HandleFunc("/status", h.Middleware(h.Status)).Methods(http.MethodGet)
	h.Mux.HandleFunc("/add_song", h.Middleware(h.AddSong)).Methods(http.MethodPost)
	h.Mux.HandleFunc("/delete_song", h.Middleware(h.DeleteSong)).Methods(http.MethodGet)

	swagHandler := httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/swagger.json"), // The url pointing to API definition"
	)

	h.Mux.Handle("/swagger/doc.json", swagHandler)
}

func (h *HttpHandlers) Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.Pl.Logger.Info().Msg(fmt.Sprintf("method: [%s] with path: [%s]", r.Method, r.URL.String()))
		next.ServeHTTP(w, r.WithContext(context.Background()))
	}
}
