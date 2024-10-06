package repository

import (
	tem2024 "github.com/Relax-FM/Test_effective_mobile_go_2024"
	"github.com/jmoiron/sqlx"
)

type MusicItem interface {
	Create(item tem2024.CreateMusicInput) (int, error)
	GetAllMusic(input tem2024.QueryParams) ([]tem2024.PageMusicItemsResponse, error)
	GetById(itemId int) ([]tem2024.CoupletMusicText, error)
	Delete(itemId int) error
	Update(itemId int, input tem2024.UpdateMusicInput) error
}

type Repository struct {
	MusicItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		MusicItem:      NewMusicItemPostgres(db),
	}
}