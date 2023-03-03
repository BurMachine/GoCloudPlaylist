package storage

import (
	"context"

	"GoCloudPlaylist/internal/models"
	"GoCloudPlaylist/pkg/timeConverting"
)

func (s *PlaylistStorage) Load() ([]models.Song, error) {
	var songs []models.Song

	rows, err := s.Conn.Query(context.Background(), "SELECT * FROM playlist")
	if err != nil {
		return songs, err
	}

	for rows.Next() {
		var name string
		var duration string
		if err := rows.Scan(&name, &duration); err != nil {
			return songs, err
		}
		durationValue, err := timeConverting.ParseTimeToSeconds(duration)
		if err != nil {
			return songs, err
		}
		song := models.Song{
			Name:     name,
			Duration: durationValue,
		}
		songs = append(songs, song)
	}

	if err := rows.Err(); err != nil {
		return songs, err
	}
	return songs, err
}
