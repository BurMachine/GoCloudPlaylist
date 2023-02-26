package httpHandlers

import (
	"GoCloudPlaylist/pkg/timeConverting"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"net/http"
)

func (h *HttpHandlers) PlaySong(w http.ResponseWriter, r *http.Request) {
	songProc := h.Pl.Play()
	timeString := timeConverting.ConvertFromSongProcToString(songProc)
	dur := timeConverting.ConvertFromSecondsToString(songProc.Duration)
	resp := processingResponse{
		Name:     songProc.Name,
		Duration: dur,
		Status:   fmt.Sprintf("%s plays at %s", songProc.Name, timeString),
	}
	res, err := json.Marshal(resp)
	if err != nil {
		h.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("response marshalling error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.Pl.Logger.Info().Msg(fmt.Sprintf("playing [%v] at %s", songProc, timeString))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(res)
	return
}

func (h *HttpHandlers) PauseSong(w http.ResponseWriter, r *http.Request) {
	songProc := h.Pl.Pause()
	timeString := timeConverting.ConvertFromSongProcToString(songProc)
	dur := timeConverting.ConvertFromSecondsToString(songProc.Duration)
	resp := processingResponse{
		Name:     songProc.Name,
		Duration: dur,
		Status:   fmt.Sprintf("%s paused at %s", songProc.Name, timeString),
	}
	res, err := json.Marshal(resp)
	if err != nil {
		h.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("response marshalling error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.Pl.Logger.Info().Msg(fmt.Sprintf("paused [%v] at %s", songProc, timeString))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(res)
}

func (h *HttpHandlers) NextSong(w http.ResponseWriter, r *http.Request) {
	songProc := h.Pl.Next()
	timeString := timeConverting.ConvertFromSongProcToString(songProc)

	var status string
	if songProc.Exist {
		status = fmt.Sprintf("Switched to next song: %s", songProc.Name)
	} else {
		status = "The next song does not exist, so you are at the end of the playlist."
	}
	dur := timeConverting.ConvertFromSecondsToString(songProc.Duration)
	resp := processingResponse{
		Name:     songProc.Name,
		Duration: dur,
		Status:   status,
	}

	res, err := json.Marshal(resp)
	if err != nil {
		h.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("response marshalling error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.Pl.Logger.Info().Msg(fmt.Sprintf("next song: [%v] at %s (exist: %v)", songProc, timeString, songProc.Exist))

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *HttpHandlers) PrevSong(w http.ResponseWriter, r *http.Request) {
	songProc := h.Pl.Prev()
	timeString := timeConverting.ConvertFromSongProcToString(songProc)

	var status string
	if songProc.Exist {
		status = fmt.Sprintf("Switched to previous song: %s", songProc.Name)
	} else {
		status = "The previous song does not exist, so you are at the beginning of the playlist."
	}

	dur := timeConverting.ConvertFromSecondsToString(songProc.Duration)
	resp := processingResponse{
		Name:     songProc.Name,
		Duration: dur,
		Status:   status,
	}

	res, err := json.Marshal(resp)
	if err != nil {
		h.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("response marshalling error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.Pl.Logger.Info().Msg(fmt.Sprintf("prev song: [%v] at %s (exist: %v)", songProc, timeString, songProc.Exist))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *HttpHandlers) Status(w http.ResponseWriter, r *http.Request) {
	songProc := h.Pl.Status()
	timeString := timeConverting.ConvertFromSongProcToString(songProc)

	var status string
	if songProc.Playing {
		status = fmt.Sprintf("Playback status: %s playing on %s", songProc.Name, timeString)
	} else {
		status = fmt.Sprintf("Playback status: %s paused on %s", songProc.Name, timeString)
	}

	dur := timeConverting.ConvertFromSecondsToString(songProc.Duration)
	resp := processingResponse{
		Name:     songProc.Name,
		Duration: dur,
		Status:   status,
	}

	res, err := json.Marshal(resp)
	if err != nil {
		h.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("response marshalling error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.Pl.Logger.Info().Msg(fmt.Sprintf("status song: [%v] at %s(playing: %v)", songProc, timeString, songProc.Playing))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
