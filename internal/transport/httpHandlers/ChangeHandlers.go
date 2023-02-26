package httpHandlers

import (
	"GoCloudPlaylist/internal/playlist"
	"GoCloudPlaylist/pkg/timeConverting"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"net/http"
)

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

	dur, err := timeConverting.ParseTimeToSeconds(song.Duration)
	if err != nil {
		h.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("time parsing error")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ok := h.Pl.AddNewSong(playlist.Song{Name: song.Name, Duration: dur})
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

	h.Pl.Logger.Info().Msg(fmt.Sprintf("[%v] added into playlist", playlist.Song{Name: song.Name, Duration: dur}))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

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
