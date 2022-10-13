package utils

import "github/chino/go-music-api/models"

// set error message
func SetError(message string) models.Error {
	var err models.Error

	err.IsError = true
	err.Message = message
	return err
}
