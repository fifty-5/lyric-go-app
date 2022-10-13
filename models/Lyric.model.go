package models

import "gorm.io/gorm"

type Lyric struct {
	gorm.Model

	ID       int     `gorm:"not null;uniqueIndex:idx_id" json:"id"`
	Name     string  `gorm:"index:idx_name" json:"name"`
	Artist   string  `gorm:"index:idx_artist" json:"artist"`
	Album    string  `gorm:"index:idx_album" json:"album"`
	Artwork  string  `json:"artwork"`
	Duration float32 `json:"duration"`
	Price    float32 `json:"price"`
	Origin   string  `json:"origin"`
}

type LyricFormatted struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Artist   string  `son:"artist"`
	Album    string  `json:"album"`
	Artwork  string  `json:"artwork"`
	Duration float32 `json:"duration"`
	Price    float32 `json:"price"`
	Origin   string  `json:"origin"`
}

func (u Lyric) ToFormatted() LyricFormatted {
	return LyricFormatted{
		Id:       u.ID,
		Name:     u.Name,
		Artist:   u.Artist,
		Album:    u.Album,
		Artwork:  u.Artwork,
		Duration: u.Duration,
		Price:    u.Price,
		Origin:   u.Origin,
	}
}
