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

func insertAndGetLyrics(ch chan models.ResultApple, ch2 chan models.ResultChartLyric) []models.LyricFormatted {
	// contiene info de apple
	std := <-ch
	// contiene info de chartlyrics
	std2 := <-ch2
	data := []models.LyricFormatted{}

	for i := 0; i < len(std2.Results); i++ {
		input := std2.Results[i]

		if input.TrackId == 0 {
			continue
		}

		lyric := searchLyric(input.TrackId)

		if lyric.ID == 0 {

			lyric = models.Lyric{
				ID:      input.TrackId,
				Name:    input.Song,
				Artist:  input.Artist,
				Artwork: input.SongUrl,
				Origin:  "chartlyrics",
			}

			utils.DB.Create(&lyric)
		}

		data = append(data, lyric.ToFormatted())
	}

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
	chApple := make(chan models.ResultApple)
	chChart := make(chan models.ResultChartLyric)

	wg.Add(2)

	client := utils.HttpClient()

	go utils.SendRequestApple(client, *params, wg, chApple)
	go utils.SendRequestChart(client, *params, wg, chChart)

	response := insertAndGetLyrics(chApple, chChart)

	wg.Wait()

	return c.Status(200).JSON(response)
}
