package PlaylistServer

import (
	"GoCloudPlaylist/internal/config"
	"GoCloudPlaylist/internal/playlist"
	"GoCloudPlaylist/internal/transport/gprcEndpoints"
	"GoCloudPlaylist/internal/transport/httpHandlers"
	api "GoCloudPlaylist/pkg/api"
	"flag"
	gorilla "github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"sync"
)

type PlaylistServer struct {
	Mux           *gorilla.Router
	Logger        *zerolog.Logger
	Conf          *config.Conf
	HttpHandlers  *httpHandlers.HttpHandlers
	GrpcEndpoints *gprcEndpoints.GrpcEndpoints
}

func New(pl *playlist.Playlist) *PlaylistServer {
	return &PlaylistServer{Mux: gorilla.NewRouter(), GrpcEndpoints: &gprcEndpoints.GrpcEndpoints{Pl: pl},
		HttpHandlers: &httpHandlers.HttpHandlers{Pl: pl}}
}

func (s PlaylistServer) Run() {
	// HTTP
	var errHttp error
	wg := &sync.WaitGroup{}
	s.HttpHandlers.Mux = s.Mux
	s.HttpHandlers.Register()
	go func() {
		wg.Add(1)
		errHttp = http.ListenAndServe(s.Conf.AddrHttp, s.Mux)
		wg.Done()
	}()
	s.Logger.Info().Msg("http server start")

	// gRPC
	grpcServerEndpoint := flag.String("grpc-server-endpoint", s.Conf.AddrGrpc, "gRPC server endpoint")

	lis, err := net.Listen("tcp", *grpcServerEndpoint)
	if err != nil {
		s.Logger.Fatal().Err(err).Msg("failed to listen")
	}
	grpcServer := grpc.NewServer()
	api.RegisterGoCloudPlaylistServer(grpcServer, s.GrpcEndpoints)

	var grpcErr error
	go func() {
		wg.Add(1)
		grpcErr = grpcServer.Serve(lis)
		wg.Done()
	}()
	s.Logger.Info().Msg("gRPC server start")

	wg.Wait()
	if errHttp != nil {
		s.Logger.Fatal().Err(errHttp).Msg("http server starting error")
	}
	if grpcErr != nil {
		s.Logger.Fatal().Err(grpcErr).Msg("grpc server starting error")
	}
}
