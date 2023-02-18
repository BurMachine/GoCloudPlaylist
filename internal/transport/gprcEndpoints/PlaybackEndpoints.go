package gprcEndpoints

import (
	api "GoCloudPlaylist/pkg/api"
	"context"
)

func (s *GrpcEndpoints) PlaySong(ctx context.Context, req *api.Empty) (*api.SongProc, error) {

	return nil, nil
}

func (s *GrpcEndpoints) PauseSong(ctx context.Context, req *api.Empty) (*api.SongProc, error) {

	return nil, nil
}

func (s *GrpcEndpoints) Next(ctx context.Context, req *api.Empty) (*api.SongProc, error) {

	return nil, nil
}

func (s *GrpcEndpoints) Prev(ctx context.Context, req *api.Empty) (*api.SongProc, error) {

	return nil, nil
}

func (s *GrpcEndpoints) Status(ctx context.Context, req *api.Empty) (*api.SongProc, error) {

	return nil, nil
}
