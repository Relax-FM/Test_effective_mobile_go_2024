package repository

import (
	"fmt"
	"strings"

	tem2024 "github.com/Relax-FM/Test_effective_mobile_go_2024"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type MusicItemPostgres struct {
	db *sqlx.DB
}

func NewMusicItemPostgres(db *sqlx.DB) *MusicItemPostgres {
	return &MusicItemPostgres{db: db}
}

func (r *MusicItemPostgres) Create(item tem2024.CreateMusicInput) (int, error) {

	var id int
	query := fmt.Sprintf("INSERT INTO %s (song_name, group_name) VALUES ($1, $2) RETURNING id", mainTable)
	logrus.Print(query)
	err := r.db.Get(&id, query, item.SongName, item.GroupName)

	return id, err
}

func (r *MusicItemPostgres) GetAllMusic(listId tem2024.QueryParams) ([]tem2024.GetPageMusicItemsResponse, error) {
	return nil, nil
}

func (r *MusicItemPostgres) GetById(itemId int) ([]tem2024.CoupletMusicText, error) {
	return nil, nil
}

func (r *MusicItemPostgres) Delete(itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1",mainTable)
	_, err := r.db.Exec(query, itemId)

	return err
}

func (r *MusicItemPostgres) Update(itemId int, input tem2024.UpdateMusicInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.SongName != nil {
		setValues = append(setValues, fmt.Sprintf("song_name=$%d", argId))
		args = append(args, *input.SongName)
		argId++
	}

	if input.GroupName != nil {
		setValues = append(setValues, fmt.Sprintf("group_name=$%d", argId))
		args = append(args, *input.GroupName)
		argId++
	}

	if input.ReleaseDate != nil {
		setValues = append(setValues, fmt.Sprintf("release_date=$%d", argId))
		onlyDate := *input.ReleaseDate
		logrus.Print(onlyDate)
		args = append(args, onlyDate)
		argId++
	}

	if input.SongText != nil {
		setValues = append(setValues, fmt.Sprintf("song_text=$%d", argId))
		args = append(args, *input.SongText)
		argId++
	}

	if input.LinkUrl != nil {
		setValues = append(setValues, fmt.Sprintf("link_url=$%d", argId))
		args = append(args, *input.LinkUrl)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", mainTable, setQuery, argId)
	args = append(args, itemId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}