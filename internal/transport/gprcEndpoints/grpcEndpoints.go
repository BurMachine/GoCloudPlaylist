package gprcEndpoints

import (
	"GoCloudPlaylist/internal/playlist"
	api "GoCloudPlaylist/pkg/api"
)

type GrpcEndpoints struct {
	api.GoCloudPlaylistServer
	Pl *playlist.Playlist
}
