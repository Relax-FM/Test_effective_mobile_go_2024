package service

import (
	tem2024 "github.com/Relax-FM/Test_effective_mobile_go_2024"
	"github.com/Relax-FM/Test_effective_mobile_go_2024/pkg/repository"
)

type MusicItem interface {
	Create(item tem2024.CreateMusicInput) (int, error)
	GetAllMusic(input tem2024.QueryParams) ([]tem2024.PageMusicItemsResponse, error)
	GetById(itemId int) ([]tem2024.CoupletMusicText, error)
	Delete(itemId int) error
	Update(itemId int, input tem2024.UpdateMusicInput) error
}

type Service struct {
	MusicItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		MusicItem:      NewMusicItemService(repos.MusicItem),
	}
}