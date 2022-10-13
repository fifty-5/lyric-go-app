package controllers

import (
	"github/chino/go-music-api/models"
	"github/chino/go-music-api/utils"
	"sync"

	"github.com/gofiber/fiber/v2"
)

func searchLyric(id int) models.Lyric {
	var lyric models.Lyric

	utils.DB.First(&lyric, id)

	return lyric
}

func insertAndGetLyrics(ch chan models.ResultApple) []models.LyricFormatted {
	std := <-ch
	data := []models.LyricFormatted{}

	for i := 0; i < len(std.Results); i++ {
		input := std.Results[i]

		lyric := searchLyric(input.Id)

		if lyric.ID == 0 {
			input.Map()

			lyric = models.Lyric{
				ID:       input.Id,
				Name:     input.Name,
				Artist:   input.Artist,
				Album:    input.Album,
				Artwork:  input.Artwork,
				Duration: input.Duration,
				Price:    input.Price,
				Origin:   input.Origin,
			}

			utils.DB.Create(&lyric)
		}

		data = append(data, lyric.ToFormatted())
	}

	return data
}

func GetLyrics(c *fiber.Ctx) error {
	params := new(models.Params)

	if err := c.QueryParser(params); err != nil {
		response := utils.SetError(err.Error())

		return c.Status(404).JSON(response)
	}

	wg := &sync.WaitGroup{}
	ch := make(chan models.ResultApple)

	wg.Add(1)

	client := utils.HttpClient()

	go utils.SendRequest(client, *params, wg, ch)

	response := insertAndGetLyrics(ch)

	wg.Wait()

	return c.Status(200).JSON(response)
}
