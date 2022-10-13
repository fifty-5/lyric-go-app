package models

type Success struct {
	Response string `json:"response"`
}

type Error struct {
	IsError bool   `json:"error"`
	Message string `json:"message"`
}
