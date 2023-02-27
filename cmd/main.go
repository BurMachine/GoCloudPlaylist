package main

import (
	"GoCloudPlaylist/internal/config"
	"GoCloudPlaylist/internal/playlist"
	PlaylistServer "GoCloudPlaylist/internal/server"
	PlaylistStorage "GoCloudPlaylist/internal/storage"
	"context"
	"flag"
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)

	cfgPath := flag.String("config", "./config.yaml", "Path to yaml configuration file")
	flag.Parse()

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()

	conf := config.NewConfigStruct()
	err := conf.LoadConfig(*cfgPath)
	if err != nil {
		logger.Fatal().Err(err).Msg("config loading error")
	}

	// Postgres
	storage, err := PlaylistStorage.InitStorage(conf.DbUrl)
	if err != nil {
		logger.Fatal().Err(err).Msg("storage init error")
	}
	defer storage.Conn.Close(context.Background())

	// Инициализация плейлиста
	pl := playlist.Init()
	pl.Logger = &logger
	storageList, err := storage.Load()
	if err != nil {
		logger.Fatal().Err(err).Msg("playlist init error")
	}
	pl.LoadListToPlaylistFromStorage(storageList)

	// Инициализация сервера
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
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	wg.Wait()
	logger.Info().Msg("service exit")
}
