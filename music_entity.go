package tem2024

import (
	"errors"
	"time"
)

type MusicItem struct {
	Id          int    		`json:"id" db:"id"`
	SongName    string 		`json:"song" db:"song_name" binding:"required"`
	GroupName   string 		`json:"group" db:"group_name" binding:"required"`
	ReleaseDate time.Time 	`json:"releaseDate" db:"release_date"`
	SongText 	string 		`json:"text" db:"song_text"`
	LinkUrl 	string 		`json:"link" db:"link_url"`
}

type QueryParams struct {
	SortBy 	string 	`form:"sort_by" binding:"required"`
	Desc 	bool 	`form:"desc" binding:"required"`
	Limit 	int 	`form:"limit" binding:"required"`
}

type CoupletMusicText struct {
	Couplet 	string 		`json:"couplet"`
}

type CreateMusicInput struct {
	SongName    string 		`json:"song" binding:"required"`
	GroupName   string 		`json:"group" binding:"required"`
}

type GetPageMusicItemsResponse struct {
	Page 		[]MusicItem `json:"page"`
}

type UpdateMusicInput struct {
	SongName    *string 	`json:"song"`
	GroupName   *string 	`json:"group"`
	ReleaseDate *time.Time 	`json:"releaseDate"`
	SongText 	*string 	`json:"text"`
	LinkUrl 	*string 	`json:"link"`
}

func (i UpdateMusicInput) Validate() error {
	if i.SongName == nil && i.GroupName == nil && i.ReleaseDate == nil && i.SongText == nil && i.LinkUrl == nil {
		return errors.New("update music structure has no values")
	}

	return nil
}