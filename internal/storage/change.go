package storage

import (
	"context"
)

var addShema = `INSERT INTO playlist (song, duration) VALUES ($1, $2)`
var delSchema = `DELETE FROM playlist WHERE song=$1;`

func (s *PlaylistStorage) Add(name, dur string) error {
	_, err := s.Conn.Exec(context.Background(), addShema, name, dur)
	if err != nil {
		return err
	}
	return nil
}

func (s PlaylistStorage) Delete(name string) error {
	_, err := s.Conn.Exec(context.Background(), delSchema, name)
	if err != nil {
		return err
	}
	return nil
}
