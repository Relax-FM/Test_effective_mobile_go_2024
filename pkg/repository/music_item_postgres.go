package repository

import (
	tem2024 "github.com/Relax-FM/Test_effective_mobile_go_2024"
	"github.com/jmoiron/sqlx"
)

type MusicItemPostgres struct {
	db *sqlx.DB
}

func NewMusicItemPostgres(db *sqlx.DB) *MusicItemPostgres {
	return &MusicItemPostgres{db: db}
}

func (r *MusicItemPostgres) Create(item tem2024.CreateMusicInput) (int, error) {
	return
}

func (r *MusicItemPostgres) getAllMusic(listId tem2024.QueryParams) ([]tem2024.GetPageMusicItemsResponse, error) {
	return
}

func (r *MusicItemPostgres) GetById(itemId int) ([]tem2024.CoupletMusicText, error) {
	return
}

func (r *MusicItemPostgres) Delete(itemId int) error {
	return
}

func (r *MusicItemPostgres) Update(itemId int, input tem2024.UpdateMusicInput) error {
	return
}