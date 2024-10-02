package service

import "github.com/Relax-FM/Test_effective_mobile_go_2024/pkg/repository"

type MusicItemService struct {
	repo repository.MusicItem
}

func NewMusicItemService(repo repository.MusicItem) *MusicItemService {
	return &MusicItemService{repo: repo}
}