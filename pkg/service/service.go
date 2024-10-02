package service

import "github.com/Relax-FM/Test_effective_mobile_go_2024/pkg/repository"

type MusicItem interface {
	// Create(userId, listId int, item todo.TodoItem) (int, error)
	// GetAll(userId, listId int) ([]todo.TodoItem, error)
	// GetById(userId, itemId int) (todo.TodoItem, error)
	// Delete(userId, itemId int) error
	// Update(userId, itemId int, input todo.UpdateItemInput) error
}

type Service struct {
	MusicItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		MusicItem:      NewMusicItemService(repos.MusicItem),
	}
}