package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"GoCloudPlaylist/internal/config"
	"GoCloudPlaylist/internal/playlist"
	PlaylistServer "GoCloudPlaylist/internal/server"
	PlaylistStorage "GoCloudPlaylist/internal/storage"
)

// @title GoCloudPlaylist API
// @version 1.0
// @description API Server for GoCloudPlaylist Application

// @host localhost:8080
// @BasePath /
func main() {
	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)

	cfgPath := flag.String("config", "./config.yaml", "Path to yaml configuration file")
	flag.Parse()

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	logger.Info().Msg("service start")

	conf := config.NewConfigStruct()
	err := conf.LoadConfig(*cfgPath)
	if err != nil {
		logger.Fatal().Err(err).Msg("config loading error")
	}

	storage, err := PlaylistStorage.InitStorage(conf.DbUrl)
	if err != nil {
		logger.Fatal().Err(err).Msg("storage init error")
	}
	defer storage.Conn.Close(context.Background())

	pl := playlist.Init()
	pl.Logger = &logger
	err = storage.LoadToStorageIfNotExistBaseSongsSet()
	if err != nil {
		logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("base songs set loading error")
	}
	storageList, err := storage.Load()
	if err != nil {
		logger.Fatal().Err(err).Msg("playlist init error")
	}
	pl.LoadListToPlaylistFromStorage(storageList)
	logger.Info().Msg("playlist initialized")

	serv := PlaylistServer.New(pl, storage)
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

	<-signalCh
	fmt.Printf("\nGracefully stopping...\n")
	pl.ExitChan <- struct{}{}
	list, err := pl.GetList()
	if err != nil {
		logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("error getting list to upload to storage")
	} else {
		err = storage.Upload(list)
		if err != nil {
			logger.WithLevel(zerolog.WarnLevel).Err(err).Msg("storage upload error")
		}
		logger.Info().Msg("state uploaded")
	}

	wg.Wait()
	logger.Info().Msg("service exit")
}
