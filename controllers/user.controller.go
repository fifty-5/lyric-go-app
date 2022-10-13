package controllers

import (
	"github/chino/go-music-api/models"
	"github/chino/go-music-api/utils"

	"github.com/gofiber/fiber/v2"
)

func Singin(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	user.Password = utils.GeneratePasswordHash(user.Password)

	if user.Email == "" {
		response := utils.SetError("Email es requerido.")

		return c.Status(404).JSON(response)
	}

	if user.Password == "" {
		response := utils.SetError("Contraseña es requerida.")

		return c.Status(404).JSON(response)
	}

	// validate email repetido
	utils.DB.Where("email = ?", user.Email).First(&user)

	if user.ID != 0 {
		response := utils.SetError("El email ya existe en la base de datos.")

		return c.Status(404).JSON(response)
	}

	result := utils.DB.Create(&user)

	if result.Error != nil {
		response := utils.SetError(result.Error.Error())

		return c.Status(404).JSON(response)
	}

	response := models.Success{
		Response: "Registro realizado correctamente.",
	}

	return c.Status(200).JSON(response)
}
