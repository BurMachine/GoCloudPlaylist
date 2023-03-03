package gprcEndpoints

import (
	"context"
	"io"
	"testing"

	"GoCloudPlaylist/internal/playlist"
	__ "GoCloudPlaylist/pkg/api"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestAddSong(t *testing.T) {
	logger := zerolog.New(io.Discard)
	pl := playlist.Init()
	pl.Logger = &logger
	go pl.Run()
	g := GrpcEndpoints{
		Pl: pl,
	}

	req := &__.AddRequest{
		Name: "string",
		Time: "00:00:10",
	}
	t.Run("add test #1", func(t *testing.T) {
		resp, err := g.AddSong(context.Background(), req)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		if resp.Playlist[0].Name != "string" || resp.Playlist[0].Duration != "00:00:10" {
			t.Error(err)
			t.Fail()
		}
	})

	t.Run("add test #2", func(t *testing.T) {
		req = &__.AddRequest{
			Name: "qwe",
			Time: "14",
		}
		_, err := g.AddSong(context.Background(), req)
		if err.Error() != status.Error(codes.FailedPrecondition, "incorrect duration format").Error() {
			t.Error(err)
			t.Fail()
		}
	})

	t.Run("add test #3", func(t *testing.T) {
		req = &__.AddRequest{
			Name: "",
			Time: "14:00:00",
		}
		_, err := g.AddSong(context.Background(), req)
		if err.Error() != status.Error(codes.FailedPrecondition, "empty name").Error() {
			t.Error(err)
			t.Fail()
		}
	})
}

func TestDeleteSong(t *testing.T) {
	logger := zerolog.New(io.Discard)
	pl := playlist.Init()
	pl.Logger = &logger
	go pl.Run()
	g := GrpcEndpoints{
		Pl: pl,
	}

	t.Run("add test #1", func(t *testing.T) {
		req := &__.AddRequest{
			Name: "string",
			Time: "00:00:10",
		}
		resp, err := g.AddSong(context.Background(), req)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		if resp.Playlist[0].Name != "string" || resp.Playlist[0].Duration != "00:00:10" {
			t.Error(err)
			t.Fail()
		}
	})

	t.Run("delete test #1", func(t *testing.T) {
		reqdel := &__.SongNameForDelete{Name: "string"}
		_, err := g.DeleteSong(context.Background(), reqdel)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
	})

	t.Run("delete test #2", func(t *testing.T) {
		reqdel := &__.SongNameForDelete{Name: "not_exist"}
		_, err := g.DeleteSong(context.Background(), reqdel)
		if err.Error() != status.Error(codes.FailedPrecondition, "song does not exist in playlist").Error() {
			t.Error(err)
			t.Fail()
		}
	})
}
