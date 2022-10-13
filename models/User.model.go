package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `gorm:"not null;uniqueIndex:idx_email;size:255" json:"email"`
	Password  string `gorm:"not null;size:255" json:"password"`
}
