package models

import "math"

type Params struct {
	Name   string `query:"name"`
	Album  string `query:"album"`
	Artist string `query:"artist"`
	Limit  string `query:"limit"`
}

type LyricApple struct {
	Id       int     `json:"trackId"`
	Artist   string  `json:"artistName"`
	Album    string  `json:"collectionName"`
	Name     string  `json:"trackName"`
	Artwork  string  `json:"artworkUrl100"`
	Duration float32 `json:"trackTimeMillis"`
	Price    float32 `json:"trackPrice"`
	Origin   string  `json:"origin"`
}

type ResultApple struct {
	ResultCount int
	Results     []LyricApple
}

func (std *LyricApple) Map() {
	if std.Origin == "" {
		std.Origin = "apple"
	}

	if std.Duration != 0 {
		duration := std.Duration / 60 / 1000
		std.Duration = float32(math.Floor(float64(duration)*100) / 100)
	}
}
