package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePasswordHash(password string) string {
	if password == "" {
		return ""
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err == nil {
		return string(bytes)
	}

	log.Fatal(err)

	return ""
}

func ValidatePasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
