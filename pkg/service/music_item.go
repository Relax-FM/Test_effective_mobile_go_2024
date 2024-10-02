package service

import (
	tem2024 "github.com/Relax-FM/Test_effective_mobile_go_2024"
	"github.com/Relax-FM/Test_effective_mobile_go_2024/pkg/repository"
)

type MusicItemService struct {
	repo repository.MusicItem
}

func NewMusicItemService(repo repository.MusicItem) *MusicItemService {
	return &MusicItemService{repo: repo}
}

func (s *MusicItemService) Create(item tem2024.CreateMusicInput) (int, error) {
	return s.repo.Create(item)
}

func (s *MusicItemService) getAllMusic(listId tem2024.QueryParams) ([]tem2024.GetPageMusicItemsResponse, error) {
	return s.repo.getAllMusic(listId)
}

func (s *MusicItemService) GetById(itemId int) ([]tem2024.CoupletMusicText, error) {
	return s.repo.GetById(itemId)
}

func (s *MusicItemService) Delete(itemId int) error {
	return s.repo.Delete(itemId)
}

func (s *MusicItemService) Update(itemId int, input tem2024.UpdateMusicInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(itemId, input)
}