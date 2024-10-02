package tem2024

import (
	"database/sql"
	"errors"
)

type MusicItemSql struct {
	Id          int    		`json:"id" db:"id"`
	SongName    string 		`json:"song" db:"song_name" binding:"required"`
	GroupName   string 		`json:"group" db:"group_name" binding:"required"`
	ReleaseDate sql.NullString		`json:"releaseDate" db:"release_date"`
	SongText 	sql.NullString 		`json:"text" db:"song_text"`
	LinkUrl 	sql.NullString 		`json:"link" db:"link_url"`
}

type MusicItem struct {
	Id          int    		`json:"id" db:"id"`
	SongName    string 		`json:"song" db:"song_name" binding:"required"`
	GroupName   string 		`json:"group" db:"group_name" binding:"required"`
	ReleaseDate string		`json:"releaseDate" db:"release_date"`
	SongText 	string 		`json:"text" db:"song_text"`
	LinkUrl 	string		`json:"link" db:"link_url"`
}

type QueryParams struct {
	SortBy 		string 		`form:"sort_by" binding:"required"`
	Desc 		bool 		`form:"desc"`
	Limit 		int 		`form:"limit" binding:"required"`
}

type CoupletMusicText struct {
	Couplet 	string 		`json:"couplet"`
}

type FullMusicText struct {
	MusicText 	string 		`db:"song_text"`	
}

type CreateMusicInput struct {
	SongName    string 		`json:"song" binding:"required" default:"Конь"`
	GroupName   string 		`json:"group" binding:"required" default:"Любэ"`
}

type PageMusicItemsResponse struct {
	Page 		[]MusicItem `json:"page"`
}

type UpdateMusicInput struct {
	SongName    *string 	`json:"song" default:"Конь"`
	GroupName   *string 	`json:"group" default:"Любэ"`
	ReleaseDate *string		`json:"releaseDate" default:"2023.12.31"`
	SongText 	*string 	`json:"text"`
	LinkUrl 	*string 	`json:"link"`
}

func (i UpdateMusicInput) Validate() error {
	if i.SongName == nil && i.GroupName == nil && i.ReleaseDate == nil && i.SongText == nil && i.LinkUrl == nil {
		return errors.New("update music structure has no values")
	}

	return nil
}