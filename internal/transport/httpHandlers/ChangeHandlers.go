package httpHandlers

import (
	"GoCloudPlaylist/internal/models"
	"GoCloudPlaylist/pkg/timeConverting"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"net/http"
)

// @Summary Add a new song to the playlist
// @Description Adds a new song to the playlist with the given name and duration (duration format 00:01:30)
// @Tags Playlist
// @Accept json
// @Produce json
// @Param song body Song true "Song object to add to the playlist"
// @Success 200 {object} []models.Song "List of all songs in the playlist"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /add_song [post]
func (h *HttpHandlers) AddSong(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("body reading error")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var song Song
	err = json.Unmarshal(body, &song)
	if err != nil {
		h.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("body unmarshalling error")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if song.Name == "" {
		h.Pl.Logger.WithLevel(zerolog.WarnLevel).Msg("empty name")
		http.Error(w, errors.New("empty name").Error(), http.StatusBadRequest)
		return
	}
	dur, err := timeConverting.ParseTimeToSeconds(song.Duration)
	if err != nil {
		h.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("time parsing error")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ok := h.Pl.AddNewSong(models.Song{Name: song.Name, Duration: dur})
	if !ok {
		h.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(errors.New("new song adding error")).Msg("song already exist")
		http.Error(w, errors.New("song adding error, song already exist or incorrect input").Error(), http.StatusBadRequest)
		return
	}

	list, err := h.Pl.GetList()
	if err != nil {
		h.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("playlist getting error in AddSong")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(list)
	if err != nil {
		h.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("response marshalling error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.Pl.Logger.Info().Msg(fmt.Sprintf("[%v] added into playlist", models.Song{Name: song.Name, Duration: dur}))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// @Summary Delete song from playlist
// @Description Deletes the song with the given name
// @Tags Playlist
// @Accept json
// @Produce json
// @Param name query string true "Song's name to delete"
// @Success 200 {object} []models.Song "List of all songs in the playlist"
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal server error"
// @Router /delete_song [get]
func (h *HttpHandlers) DeleteSong(w http.ResponseWriter, r *http.Request) {
	songName := r.URL.Query().Get("name")
	if songName == "" {
		h.Pl.Logger.WithLevel(zerolog.WarnLevel).Msg("method GET query is empty")
		http.Error(w, errors.New("empty link").Error(), http.StatusBadRequest)
		return
	}

	err := h.Pl.DeleteSong(songName)
	if err != nil {
		h.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("song deleting error")
		if errors.Is(err, errors.New("can't delete song while playing")) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		return
	}

	list, err := h.Pl.GetList()
	if err != nil {
		h.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("playlist getting error in DeleteSong")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(list)
	if err != nil {
		h.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("result playlist marshalling error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.Pl.Logger.Info().Msg(fmt.Sprintf("[%s] deleted from playlist", songName))

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
