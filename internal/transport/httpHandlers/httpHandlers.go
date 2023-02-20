package httpHandlers

import (
	"GoCloudPlaylist/internal/playlist"
	gorilla "github.com/gorilla/mux"
	"net/http"
)

type HttpHandlers struct {
	Mux *gorilla.Router
	Pl  *playlist.Playlist
}

func (h *HttpHandlers) Register() {
	h.Mux.HandleFunc("/play", h.PlaySong).Methods(http.MethodGet)
	h.Mux.HandleFunc("/pause", h.PauseSong).Methods(http.MethodGet)
	h.Mux.HandleFunc("/next_song", h.NextSong).Methods(http.MethodGet)
	h.Mux.HandleFunc("/prev_song", h.PrevSong).Methods(http.MethodGet)
	h.Mux.HandleFunc("/status", h.Status).Methods(http.MethodGet)
	h.Mux.HandleFunc("/add_song", h.AddSong).Methods(http.MethodPost)
}
