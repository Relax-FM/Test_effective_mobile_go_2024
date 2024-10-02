package tem2024

import (
	"time"
)

type UpdateMusicInput struct {
	SongName    string `json:"song"`
	GroupName   string `json:"group"`
	ReleaseDate time.Time `json:"releaseDate"`
	SongText 	string `json:"text"`
	LinkUrl  	string `json:"link"`
}

type MusicItem struct {
	Id          int    `json:"id" db:"id"`
	SongName    string `json:"song" db:"song_name" binding:"required"`
	GroupName   string `json:"group" db:"group_name" binding:"required"`
	ReleaseDate time.Time `json:"releaseDate" db:"release_date"`
	SongText 	string `json:"text" db:"song_text"`
	LinkUrl 	string `json:"link" db:"link_url"`
}

type MusicText struct {
	Couplet string `json:"couplet"`
}

type CreateMusicInput struct {
	SongName    string `json:"song" binding:"required"`
	GroupName   string `json:"group" binding:"required"`
}

func (i UpdateMusicInput) Validate() error {
	
}