package httpHandlers

import (
	gorilla "github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type HttpHandlers struct {
	Logger *zerolog.Logger
	Mux    *gorilla.Router
}

func (h *HttpHandlers) Register() {
	h.Mux.HandleFunc("/add_song", h.AddSong)
}
