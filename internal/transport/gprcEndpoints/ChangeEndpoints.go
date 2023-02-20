package gprcEndpoints

import (
	"GoCloudPlaylist/internal/playlist"
	api "GoCloudPlaylist/pkg/api"
	"GoCloudPlaylist/pkg/timeConverting"
	"context"
	"errors"
)

func (s *GrpcEndpoints) AddSong(ctx context.Context, req *api.AddRequest) (*api.PlaylistResponse, error) {
	time, err := timeConverting.ParseTimeToSeconds(req.Time)
	if err != nil {
		return nil, err
	}
	println(time)
	ok := s.Pl.AddNewSong(playlist.Song{Name: req.Name, Duration: time})
	if !ok {
		return nil, errors.New("new song adding error")
	}
	list, err := s.Pl.GetList()
	if err != nil {
		return nil, err
	}
	var res api.PlaylistResponse
	for _, song := range list {
		dur := timeConverting.ConvertFromSecondsToString(song.Duration)
		songRes := api.Song{
			Name:     song.Name,
			Duration: dur,
		}
		res.Playlist = append(res.Playlist, &songRes)
	}

	return &res, nil
}

func (s *GrpcEndpoints) DeleteSong(ctx context.Context, req *api.SongNameForDelete) (*api.PlaylistResponse, error) {

	return nil, nil
}
