package repository

import "github.com/jmoiron/sqlx"

type MusicItem interface {
	// Create(listId int, item todo.TodoItem) (int, error)
	// GetAll(userId, listId int) ([]todo.TodoItem, error)
	// GetById(userId, itemId int) (todo.TodoItem, error)
	// Delete(userId, itemId int) error
	// Update(userId, itemId int, input todo.UpdateItemInput) error
}

type Repository struct {
	MusicItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		MusicItem:      NewMusicItemPostgres(db),
	}
}