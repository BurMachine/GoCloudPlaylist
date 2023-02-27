package main

import (
	"GoCloudPlaylist/internal/config"
	"GoCloudPlaylist/internal/playlist"
	PlaylistServer "GoCloudPlaylist/internal/server"
	"flag"
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"os/signal"
	"sync"
)

func main() {
	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, os.Interrupt, os.Kill)

	cfgPath := flag.String("config", "./config.yaml", "Path to yaml configuration file")
	flag.Parse()

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()

	conf := config.NewConfigStruct()
	err := conf.LoadConfig(*cfgPath)
	if err != nil {
		logger.Fatal().Err(err).Msg("config loading error")
	}

	// Инициализация плейлиста
	pl := playlist.Init()
	pl.Logger = &logger

	serv := PlaylistServer.New(pl)
	serv.Logger = &logger

	wg := &sync.WaitGroup{}
	serv.Conf = conf
	go func() {
		serv.Run()
	}()

	go func() {
		wg.Add(1)
		pl.Run()
		wg.Done()
	}()
	//time.Sleep(15 * time.Second)
	//
	//time.Sleep(3 * time.Second)
	////pl.Play()
	//pl.AddNewSong(playlist.Song{Name: "Kingslayer", Duration: 12})

	<-signalCh
	fmt.Printf("\nGracefully stopping...\n")
	pl.ExitChan <- struct{}{}
	wg.Wait()
	logger.Info().Msg("service exit")
}
