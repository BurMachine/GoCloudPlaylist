package gprcEndpoints

import (
	"GoCloudPlaylist/internal/models"
	api "GoCloudPlaylist/pkg/api"
	"GoCloudPlaylist/pkg/timeConverting"
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
)

func (s *GrpcEndpoints) AddSong(ctx context.Context, req *api.AddRequest) (*api.PlaylistResponse, error) {
	time, err := timeConverting.ParseTimeToSeconds(req.Time)
	if err != nil {
		s.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("parse time to seconds error")
		return nil, err
	}
	ok := s.Pl.AddNewSong(models.Song{Name: req.Name, Duration: time})
	if !ok {
		s.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(errors.New("new song adding error")).Msg("song already exist")
		return nil, errors.New("new song adding error, song already exist or incorrect input")
	}
	list, err := s.Pl.GetList()
	if err != nil {
		s.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("playlist getting error in AddSong")
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
	s.Pl.Logger.Info().Msg(fmt.Sprintf("[%v] added into playlist", models.Song{Name: req.Name, Duration: time}))
	return &res, nil
}

func (s *GrpcEndpoints) DeleteSong(ctx context.Context, req *api.SongNameForDelete) (*api.PlaylistResponse, error) {
	var res api.PlaylistResponse
	err := s.Pl.DeleteSong(req.Name)
	if err != nil {
		s.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("song deleting error")
		return &res, err
	}

	list, err := s.Pl.GetList()
	if err != nil {
		s.Pl.Logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("playlist getting error in DeleteSong")
		return nil, err
	}

	for _, song := range list {
		dur := timeConverting.ConvertFromSecondsToString(song.Duration)
		songRes := api.Song{
			Name:     song.Name,
			Duration: dur,
		}
		res.Playlist = append(res.Playlist, &songRes)
	}
	s.Pl.Logger.Info().Msg(fmt.Sprintf("[%s] deleted from playlist", req.Name))
	return &res, nil
}
