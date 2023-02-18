package httpHandlers

import (
	"fmt"
	"net/http"
)

func (h *HttpHandlers) PlaySong(w http.ResponseWriter, r *http.Request) {

	h.Logger.Info().Msg("Privet play")

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "Privet play)")
}
