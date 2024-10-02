package repository

import "github.com/jmoiron/sqlx"

type MusicItemPostgres struct {
	db *sqlx.DB
}

func NewMusicItemPostgres(db *sqlx.DB) *MusicItemPostgres {
	return &MusicItemPostgres{db: db}
}