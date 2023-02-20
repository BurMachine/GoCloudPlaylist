package gprcEndpoints

import (
	api "GoCloudPlaylist/pkg/api"
	"GoCloudPlaylist/pkg/timeConverting"
	"context"
	"fmt"
)

func (s *GrpcEndpoints) PlaySong(ctx context.Context, req *api.Empty) (*api.SongProc, error) {
	songProc := s.Pl.Play()
	timeString := timeConverting.ConvertFromSongProcToString(songProc)
	s.Pl.Logger.Info().Msg(fmt.Sprintf("playing %v at %s", songProc, timeString))
	return &api.SongProc{
		Name:   songProc.Name,
		Time:   timeConverting.ConvertFromSecondsToString(songProc.Duration),
		Status: fmt.Sprintf("%s plays at %s", songProc.Name, timeString),
	}, nil
}

func (s *GrpcEndpoints) PauseSong(ctx context.Context, req *api.Empty) (*api.SongProc, error) {
	songProc := s.Pl.Pause()
	timeString := timeConverting.ConvertFromSongProcToString(songProc)
	s.Pl.Logger.Info().Msg(fmt.Sprintf("paused [%v] at %s", songProc, timeString))
	return &api.SongProc{
		Name:   songProc.Name,
		Time:   timeConverting.ConvertFromSecondsToString(songProc.Duration),
		Status: fmt.Sprintf("%s is paused at %s", songProc.Name, timeString),
	}, nil
}

func (s *GrpcEndpoints) Next(ctx context.Context, req *api.Empty) (*api.SongProc, error) {
	songProc := s.Pl.Next()
	timeString := timeConverting.ConvertFromSongProcToString(songProc)
	s.Pl.Logger.Info().Msg(fmt.Sprintf("next song: [%v] at %s (exist: %v)", songProc, timeString, songProc.Exist))

	var status string
	if songProc.Exist {
		status = fmt.Sprintf("Switched to next song: %s", songProc.Name)
	} else {
		status = "The next song does not exist, so you are at the end of the playlist."
	}
	return &api.SongProc{
		Name:   songProc.Name,
		Time:   timeConverting.ConvertFromSecondsToString(songProc.Duration),
		Status: status,
	}, nil
}

func (s *GrpcEndpoints) Prev(ctx context.Context, req *api.Empty) (*api.SongProc, error) {
	songProc := s.Pl.Prev()
	timeString := timeConverting.ConvertFromSongProcToString(songProc)
	s.Pl.Logger.Info().Msg(fmt.Sprintf("prev song: [%v] at %s (exist: %v)", songProc, timeString, songProc.Exist))
	var status string
	if songProc.Exist {
		status = fmt.Sprintf("Switched to previous song: %s", songProc.Name)
	} else {
		status = "The previous song does not exist, so you are at the beginning of the playlist."
	}
	return &api.SongProc{
		Name:   songProc.Name,
		Time:   timeConverting.ConvertFromSecondsToString(songProc.Duration),
		Status: status,
	}, nil
}

func (s *GrpcEndpoints) Status(ctx context.Context, req *api.Empty) (*api.SongProc, error) {
	songProc := s.Pl.Status()
	timeString := timeConverting.ConvertFromSongProcToString(songProc)
	s.Pl.Logger.Info().Msg(fmt.Sprintf("status song: [%v] at %s(playing: %v)", songProc, timeString, songProc.Playing))
	var status string
	if songProc.Playing {
		status = fmt.Sprintf("Playback status: %s playing on %s", songProc.Name, timeString)
	} else {
		status = fmt.Sprintf("Playback status: %s paused on %s", songProc.Name, timeString)
	}
	return &api.SongProc{
		Name:   songProc.Name,
		Time:   timeConverting.ConvertFromSecondsToString(songProc.Duration),
		Status: status,
	}, nil
}
