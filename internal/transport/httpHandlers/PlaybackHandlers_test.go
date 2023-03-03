package httpHandlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"GoCloudPlaylist/internal/models"
	"GoCloudPlaylist/internal/playlist"
	"github.com/rs/zerolog"
)

func TestHttpHandlers_PlayPauseStatus(t *testing.T) {
	logger := zerolog.New(io.Discard)
	pl := playlist.Init()
	pl.Logger = &logger
	go pl.Run()
	h := HttpHandlers{
		Pl: pl,
	}
	pl.AddNewSong(models.Song{
		Name:     "string",
		Duration: 10,
	})

	t.Run("play test", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/play", nil)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.PlaySong)
		handler.ServeHTTP(w, req)

		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	})
	t.Run("pause test", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/pause", nil)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.PauseSong)
		handler.ServeHTTP(w, req)

		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	})
	t.Run("status test", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/status", nil)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.Status)
		handler.ServeHTTP(w, req)

		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	})
}

func TestHttpHandlers_NextPrevSong(t *testing.T) {
	logger := zerolog.New(io.Discard)
	pl := playlist.Init()
	pl.Logger = &logger
	go pl.Run()
	h := HttpHandlers{
		Pl: pl,
	}

	t.Run("next test #1", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/next_song", nil)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.NextSong)
		handler.ServeHTTP(w, req)

		if status := w.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	})
	pl.AddNewSong(models.Song{
		Name:     "string",
		Duration: 10,
	})
	pl.AddNewSong(models.Song{
		Name:     "qwerty",
		Duration: 10,
	})
	t.Run("status test", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/status", nil)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.Status)
		handler.ServeHTTP(w, req)

		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	})

	t.Run("next test #2", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/next_song", nil)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.NextSong)
		handler.ServeHTTP(w, req)

		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

	})

	t.Run("next test #3", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/next_song", nil)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.NextSong)
		handler.ServeHTTP(w, req)

		if status := w.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	})

	t.Run("prev test #1", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/prev_song", nil)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.PrevSong)
		handler.ServeHTTP(w, req)

		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	})

	t.Run("prev test #2", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/prev_song", nil)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.PrevSong)
		handler.ServeHTTP(w, req)

		if status := w.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	})
}
