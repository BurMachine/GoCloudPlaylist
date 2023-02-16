package PlaylistServer

import (
	"GoCloudPlaylist/internal/config"
	"GoCloudPlaylist/internal/transport/httpHandlers"
	gorilla "github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"net/http"
	"sync"
)

type PlaylistServer struct {
	Mux    *gorilla.Router
	Logger *zerolog.Logger
	Conf   *config.Conf
}

func New() *PlaylistServer {
	return &PlaylistServer{Mux: gorilla.NewRouter()}
}

func (s PlaylistServer) Run() {
	var errHttp error
	wg := &sync.WaitGroup{}
	httpH := &httpHandlers.HttpHandlers{
		Logger: s.Logger,
		Mux:    s.Mux,
	}
	httpH.Register()
	go func() {
		wg.Add(1)
		errHttp = http.ListenAndServe(s.Conf.AddrHttp, s.Mux)
		wg.Done()
	}()
	s.Logger.Info().Msg("http server start")
	wg.Wait()
	if errHttp != nil {
		s.Logger.Fatal().Err(errHttp).Msg("http server starting error")
	}
}
