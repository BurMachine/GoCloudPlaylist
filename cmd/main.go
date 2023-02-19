package main

import (
	"GoCloudPlaylist/internal/config"
	"GoCloudPlaylist/internal/playlist"
	PlaylistServer "GoCloudPlaylist/internal/server"
	"flag"
	"github.com/rs/zerolog"
	"log"
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
	a := pl.Pause()
	time.Sleep(3 * time.Second)
	pl.Next()

	time.Sleep(5 * time.Second)
	b := pl.Play()
	log.Println(a, b)

	wg.Wait()
	logger.Info().Msg("service exit")
}
