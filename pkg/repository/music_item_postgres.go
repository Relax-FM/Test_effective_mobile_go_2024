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
	logrus.Debug(query)
	err := r.db.Get(&id, query, item.SongName, item.GroupName)

	return id, err
}


func (r *MusicItemPostgres) GetAllMusic(input tem2024.QueryParams) ([]tem2024.PageMusicItemsResponse, error) {

	validOrderParams := map[string]struct{}{
		"id":           {},
		"song_name":    {},
		"group_name":   {},
		"release_date": {},
		"song_text":    {},
	}

	var allMusicElements []tem2024.MusicItemSql

	_, exists := validOrderParams[input.SortBy]
	if !exists {
		input.SortBy = "id"
	}

	var descAsc string
	if input.Desc{
		descAsc = "DESC"
	} else {
		descAsc = "ASC"
	}

	query := fmt.Sprintf("SELECT * FROM %s ORDER BY %s %s", mainTable, input.SortBy, descAsc)
	err := r.db.Select(&allMusicElements, query)

	logrus.Debug(allMusicElements)

	var pages []tem2024.PageMusicItemsResponse
	page := tem2024.PageMusicItemsResponse{
		Page: []tem2024.MusicItem{},
	}
	count := 0

	for _, element := range allMusicElements {
		if count == input.Limit {
			logrus.Debug(page)
			pages = append(pages, page)
			page.Page = []tem2024.MusicItem{}
			count = 0
		}
		page.Page = append(page.Page, tem2024.MusicItem{
			Id          : element.Id,
			SongName    : element.SongName,
			GroupName   : element.GroupName,
			ReleaseDate : element.ReleaseDate.String,
			SongText 	: element.SongText.String,
			LinkUrl 	: element.LinkUrl.String,
		})
		count++
	}
	pages = append(pages, page)
	page.Page = []tem2024.MusicItem{}
	count = 0

	return pages, err
}

func (r *MusicItemPostgres) GetById(itemId int) ([]tem2024.CoupletMusicText, error) {
	var fullText tem2024.FullMusicText

	logrus.Debug(itemId)

	query := fmt.Sprintf("SELECT song_text FROM %s WHERE id=$1", mainTable)
	err := r.db.Get(&fullText, query, itemId)

	logrus.Debug(fullText)

	// Пагинация по куплетам (\n\n)
	var list []tem2024.CoupletMusicText

	splitedTexts := strings.Split(fullText.MusicText, "\n\n")

	for _, splitedText := range splitedTexts{
		list = append(list, tem2024.CoupletMusicText{
			Couplet: splitedText,
		})
	}

	logrus.Debug(list)

	return list, err
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
		logrus.Debug(onlyDate)
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