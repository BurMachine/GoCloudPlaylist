package gprcEndpoints

import (
	"context"
	"io"
	"testing"

	"GoCloudPlaylist/internal/models"
	"GoCloudPlaylist/internal/playlist"
	__ "GoCloudPlaylist/pkg/api"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGrpcEndpoints_PlaySong_PauseSong_StatusSong(t *testing.T) {
	logger := zerolog.New(io.Discard)
	pl := playlist.Init()
	pl.Logger = &logger
	go pl.Run()
	g := GrpcEndpoints{
		Pl: pl,
	}
	pl.AddNewSong(models.Song{
		Name:     "string",
		Duration: 10,
	})
	t.Run("play test", func(t *testing.T) {
		resp, err := g.PlaySong(context.Background(), &__.Empty{})
		if err != status.Errorf(codes.OK, "OK") {
			t.Error(err)
			t.Fail()
		}
		if resp.Name != "string" || resp.Time != "00:00:10" {
			t.Error(err)
			t.Fail()
		}
	})

	t.Run("pause test", func(t *testing.T) {
		resp, err := g.PauseSong(context.Background(), &__.Empty{})
		if err != status.Errorf(codes.OK, "OK") {
			t.Error(err)
			t.Fail()
		}
		if resp.Name != "string" || resp.Time != "00:00:10" {
			t.Error(err)
			t.Fail()
		}
	})

	t.Run("status test", func(t *testing.T) {
		resp, err := g.Status(context.Background(), &__.Empty{})
		if err != status.Errorf(codes.OK, "OK") {
			t.Error(err)
			t.Fail()
		}
		if resp.Name != "string" || resp.Time != "00:00:10" {
			t.Error(err)
			t.Fail()
		}
	})
}

func TestGrpcEndpoints_Next_Prev(t *testing.T) {
	logger := zerolog.New(io.Discard)
	pl := playlist.Init()
	pl.Logger = &logger
	go pl.Run()
	g := GrpcEndpoints{
		Pl: pl,
	}
	pl.AddNewSong(models.Song{
		Name:     "string",
		Duration: 10,
	})
	pl.AddNewSong(models.Song{
		Name:     "qwerty",
		Duration: 10,
	})

	t.Run("play test", func(t *testing.T) {
		resp, err := g.PlaySong(context.Background(), &__.Empty{})
		if err != status.Errorf(codes.OK, "OK") {
			t.Error(err)
			t.Fail()
		}

		resp, err = g.PauseSong(context.Background(), &__.Empty{})
		if err != status.Errorf(codes.OK, "OK") {
			t.Error(err)
			t.Fail()
		}
		if resp.Name != "string" || resp.Time != "00:00:10" {
			t.Error(err)
			t.Fail()
		}
	})
	t.Run("next test #1", func(t *testing.T) {
		resp, err := g.Next(context.Background(), &__.Empty{})
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		if resp.Name != "qwerty" || resp.Time != "00:00:10" {
			t.Error(err)
			t.Fail()
		}
	})
	t.Run("next test #2", func(t *testing.T) {
		_, err := g.Next(context.Background(), &__.Empty{})
		statusString := "The next song does not exist, so you are at the end of the playlist."
		if err.Error() != status.Errorf(codes.NotFound, statusString).Error() {
			t.Error(err)
			t.Fail()
		}
	})

	t.Run("prev test #1", func(t *testing.T) {
		resp, err := g.Prev(context.Background(), &__.Empty{})
		if err != status.Errorf(codes.OK, "OK") {
			t.Error(err)
			t.Fail()
		}
		if resp.Name != "string" || resp.Time != "00:00:10" {
			t.Error(err)
			t.Fail()
		}
	})

	t.Run("next test #2", func(t *testing.T) {
		_, err := g.Prev(context.Background(), &__.Empty{})
		statusString := "The previous song does not exist, so you are at the beginning of the playlist."
		if err.Error() != status.Errorf(codes.NotFound, statusString).Error() {
			t.Error(err)
			t.Fail()
		}
	})
}
