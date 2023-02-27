package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
	"os"
)

type PlaylistStorage struct {
	Conn   *pgx.Conn
	Logger *zerolog.Logger
}

var schema = `CREATE TABLE IF NOT EXISTS playlist (
	song VARCHAR(1000) NOT NULL,
	duration VARCHAR(50) NOT NULL
);`

func InitStorage(url string) (*PlaylistStorage, error) {
	var res PlaylistStorage

	dsn := os.Getenv("POSTGRES_URI")
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil || dsn == "" {
		conn, err = pgx.Connect(context.Background(), url)
		if err != nil {
			return &res, err
		}
	}
	_, err = conn.Exec(context.Background(), schema)
	if err != nil {
		return nil, fmt.Errorf("table creating error: %v", err)
	}

	res.Conn = conn
	return &res, nil
}
