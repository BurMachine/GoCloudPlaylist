package storage

import (
	"context"

	"GoCloudPlaylist/internal/models"
	"GoCloudPlaylist/pkg/timeConverting"
)

var uploadScheme = `INSERT INTO playlist (song, duration) VALUES ($1, $2);`

var cleanScheme = `DELETE FROM playlist;`

// Upload Выгрузка данных в хранилище
func (s *PlaylistStorage) Upload(list []models.Song) error {
	_, err := s.Conn.Exec(context.Background(), cleanScheme)
	if err != nil {
		return err
	}
	for _, song := range list {
		time := timeConverting.ConvertFromSecondsToString(song.Duration)
		_, err := s.Conn.Exec(context.Background(), uploadScheme, song.Name, time)
		if err != nil {
			return err
		}
	}
	return nil
}
