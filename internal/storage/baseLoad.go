package storage

import (
	"GoCloudPlaylist/internal/models"
	"GoCloudPlaylist/pkg/timeConverting"
	"context"
)

var schemaIns = `
INSERT INTO playlist (song, duration)
SELECT $1, $2
WHERE
    NOT EXISTS (
        SELECT song FROM playlist WHERE song = $3
    );
`

func (s *PlaylistStorage) LoadToStorageIfNotExistBaseSongsSet() error {
	songs := []models.Song{
		{
			Name:     "Bohemian Rhapsody - Queen",
			Duration: 30,
		},
		{
			Name:     "Waka Waka - Shakira",
			Duration: 25,
		},
		{
			Name:     "Paper Planes - M.I.A.",
			Duration: 27,
		},
		{
			Name:     "Nothing Else Matters - Metallica",
			Duration: 32,
		},
		{
			Name:     "Lose Yourself",
			Duration: 31,
		},
		{
			Name:     "Figure.09 - Linkin Park",
			Duration: 30,
		},
		{
			Name:     "Chop Suey! - System of a down",
			Duration: 28,
		},
	}
	for _, song := range songs {
		_, err := s.Conn.Exec(context.Background(), schemaIns, song.Name, timeConverting.ConvertFromSecondsToString(song.Duration),
			song.Name)
		if err != nil {
			return err
		}
	}
	return nil
}
