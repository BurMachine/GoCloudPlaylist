package httpHandlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"GoCloudPlaylist/internal/models"
	"GoCloudPlaylist/internal/playlist"
	gorilla "github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

func TestHttpHandlers_AddSong(t *testing.T) {
	logger := zerolog.New(io.Discard)
	pl := playlist.Init()
	pl.Logger = &logger
	go pl.Run()
	h := HttpHandlers{
		Pl: pl,
	}

	t.Run("add test #1", func(t *testing.T) {
		song, err := json.Marshal(Song{
			Name:     "Happy New Year",
			Duration: "00:00:10",
		})
		if err != nil {
			t.Fatal(err)
		}
		body := bytes.NewReader(song)
		req, err := http.NewRequest("POST", "/add_song", body)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.AddSong)
		handler.ServeHTTP(w, req)

		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	})

	t.Run("add test #2", func(t *testing.T) {
		song, err := json.Marshal(Song{
			Name:     "Happy New Year",
			Duration: "10",
		})
		if err != nil {
			t.Fatal(err)
		}
		body := bytes.NewReader(song)
		req, err := http.NewRequest("POST", "/add_song", body)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.AddSong)
		handler.ServeHTTP(w, req)

		if status := w.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}
	})
	t.Run("add test #3", func(t *testing.T) {
		song, err := json.Marshal(Song{
			Name:     "",
			Duration: "00:00:10",
		})
		if err != nil {
			t.Fatal(err)
		}
		body := bytes.NewReader(song)
		req, err := http.NewRequest("POST", "/add_song", body)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.AddSong)
		handler.ServeHTTP(w, req)

		if status := w.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}
	})
}

func TestHttpHandlers_DeleteSong(t *testing.T) {
	logger := zerolog.New(io.Discard)
	pl := playlist.Init()
	pl.Logger = &logger
	go pl.Run()
	h := HttpHandlers{
		Pl: pl,
	}

	t.Run("delete test #1", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/delete_song?name=", nil)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.DeleteSong)
		handler.ServeHTTP(w, req)

		if status := w.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}
	})

	t.Run("delete test #2", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/delete_song?name=string", nil)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.DeleteSong)
		handler.ServeHTTP(w, req)

		if status := w.Code; status != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusNotFound)
		}
	})

	pl.AddNewSong(models.Song{
		Name:     "string",
		Duration: 10,
	})

	t.Run("delete test #3", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/delete_song?name=not_found", nil)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.DeleteSong)
		handler.ServeHTTP(w, req)

		if status := w.Code; status != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusNotFound)
		}
	})

	t.Run("register test #1", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/delete_song?name=string", nil)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		handler := http.HandlerFunc(h.DeleteSong)
		handler.ServeHTTP(w, req)

		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	})
	t.Run("delete test #4", func(t *testing.T) {
		h.Mux = gorilla.NewRouter()
		h.Register()
	})
}
