package main

import (
	"GoCloudPlaylist/internal/config"
	"GoCloudPlaylist/internal/playlist"
	PlaylistServer "GoCloudPlaylist/internal/server"
	"flag"
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"sync"
	"time"
)

func main() {
	cfgPath := flag.String("config", "./config.yaml", "Path to yaml configuration file")
	flag.Parse()

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()

	conf := config.NewConfigStruct()
	err := conf.LoadConfig(*cfgPath)
	if err != nil {
		logger.Fatal().Err(err).Msg("config loading error")
	}
	serv := PlaylistServer.New()
	serv.Logger = &logger

	wg := &sync.WaitGroup{}
	serv.Conf = conf
	go func() {
		wg.Add(1)
		serv.Run()
		wg.Done()
	}()

	pl := playlist.Init()
	pl.Logger = &logger
	go func() {
		wg.Add(1)
		pl.Run()
		wg.Done()
	}()
	//time.Sleep(15 * time.Second)

	time.Sleep(3 * time.Second)
	pl.Play()
	pl.AddNewSong(playlist.Song{Name: "Kingslayer", Duration: 12})
	l, err := pl.GetList()
	if err != nil {
		println("((((")
	}
	fmt.Println(l)
	time.Sleep(5 * time.Second)
	err = pl.DeleteSong("Run Free")
	if err != nil {
		println("((((1")
	}
	time.Sleep(2 * time.Second)
	l, err = pl.GetList()
	if err != nil {
		println("((((")
	}
	fmt.Println(l)
	pl.Status()

	wg.Wait()
	logger.Info().Msg("service exit")
}
