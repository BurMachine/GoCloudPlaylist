package gprcEndpoints

import (
	api "GoCloudPlaylist/pkg/api"
	"context"
	"log"
)

func (s *GrpcEndpoints) VersionGet(ctx context.Context, req *api.VersionRequest) (*api.PlaylistResponse, error) {
	log.Println("Привет из Version Get")
	return &api.PlaylistResponse{Playlist: []*api.Song{
		{
			Name:     "Demolisher",
			Duration: "3:66",
		},
	}}, nil
}

func (s *GrpcEndpoints) UploadVersionToPlaylist(ctx context.Context, req *api.VersionRequest) (*api.PlaylistResponse, error) {

	return nil, nil
}
