package playlist

import (
	"container/list"
	_ "container/list"
)

type Song struct {
	Name     string
	Duration int
}

func Init() *list.List {
	return list.New()
}

func AddSong(name string, duration int, playlist *list.List) *list.List {
	song := Song{Name: name, Duration: duration}
	playlist.PushBack(song)
	return playlist
}
