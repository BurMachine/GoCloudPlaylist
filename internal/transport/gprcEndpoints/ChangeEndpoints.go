package gprcEndpoints

import (
	api "GoCloudPlaylist/pkg/api"
	"context"
)

func (s *GrpcEndpoints) AddSong(ctx context.Context, req *api.AddRequest) (*api.PlaylistResponse, error) {

	return nil, nil
}

func (s *GrpcEndpoints) DeleteSong(ctx context.Context, req *api.SongNameForDelete) (*api.PlaylistResponse, error) {

	return nil, nil
}
