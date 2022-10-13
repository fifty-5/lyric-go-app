package main

import (
	"github/chino/go-music-api/controllers"
	"github/chino/go-music-api/middlewares"
	"github/chino/go-music-api/models"
	"github/chino/go-music-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// load env
	utils.LoadEnv()

	// db connection
	utils.DBConnection()

	// run migration
	utils.DB.AutoMigrate(models.Lyric{})
	utils.DB.AutoMigrate(models.User{})

	app := fiber.New()

	// middlewares
	app.Use(cors.New())
	app.Use(basicauth.New(middlewares.ConfigDefault))

	// routes
	app.Get("/lyric", controllers.GetLyrics)
	app.Post("/signin", controllers.Singin)

	app.Listen(":3000")
}
