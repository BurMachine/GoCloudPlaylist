package main

import (
	"GoCloudPlaylist/internal/playlist"
	"fmt"
)

func main() {
	pl := playlist.Init()
	playlist.AddSong("Demolisher", 120, pl)
	for e := pl.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
