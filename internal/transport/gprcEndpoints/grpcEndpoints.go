package gprcEndpoints

import (
	"GoCloudPlaylist/internal/playlist"
	"GoCloudPlaylist/internal/storage"
	api "GoCloudPlaylist/pkg/api"
)

type GrpcEndpoints struct {
	api.GoCloudPlaylistServer
	Pl *playlist.Playlist
	Db *storage.PlaylistStorage
}
