package middlewares

import (
	"github/chino/go-music-api/models"
	"github/chino/go-music-api/utils"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

var ConfigDefault = basicauth.Config{
	Next: func(c *fiber.Ctx) bool {
		// solo esta ruta para registro no cuenta con auth basic
		if c.OriginalURL() == "/signin" && c.Method() == "POST" {
			return true
		}

		return false
	},
	Realm: "Forbidden",
	Authorizer: func(email, pass string) bool {
		var user models.User

		utils.DB.Where("email = ?", email).First(&user)

		if user.ID != 0 && utils.ValidatePasswordHash(pass, user.Password) {
			return true
		}

		return false
	},
	Unauthorized: func(c *fiber.Ctx) error {
		response := utils.SetError("No autorizado para solicitar este recurso")

		return c.Status(401).JSON(response)
	},
}
