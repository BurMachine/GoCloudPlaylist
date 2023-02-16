package main

import (
	"GoCloudPlaylist/internal/config"
	PlaylistServer "GoCloudPlaylist/internal/server"
	"flag"
	"github.com/rs/zerolog"
	"os"
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
	serv.Conf = conf
	serv.Run()
	//pl := playlist.Init()
	//playlist.AddSong("Demolisher", 120, pl)
	//for e := pl.Front(); e != nil; e = e.Next() {
	//	fmt.Println(e.Value)
	//}
}
