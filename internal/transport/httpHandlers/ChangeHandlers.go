package httpHandlers

import (
	"fmt"
	"net/http"
)

func (h *HttpHandlers) AddSong(w http.ResponseWriter, r *http.Request) {

	h.Logger.Info().Msg("Privet")

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "Privet add)")
}

func (h *HttpHandlers) DeleteSong(w http.ResponseWriter, r *http.Request) {

	h.Logger.Info().Msg("Privet")

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "Privet delete)")
}
