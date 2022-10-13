package utils

import (
	"encoding/json"
	"fmt"
	"github/chino/go-music-api/models"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func HttpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

func validateLimit(limit string) string {
	if limit != "" {
		return limit
	} else {
		return "2"
	}
}

func SendRequest(client *http.Client, params models.Params, wg *sync.WaitGroup, ch chan models.ResultApple) {
	req, err := http.NewRequest(http.MethodGet, "https://itunes.apple.com/search", nil)
	if err != nil {
		log.Fatal(err)
	}

	// añadiedo query params
	q := req.URL.Query()
	q.Add("term", fmt.Sprintf("%s %s %s", params.Name, params.Artist, params.Album))
	q.Add("media", "music")
	q.Add("limit", validateLimit(params.Limit))

	// encode query params
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}

	if resp.StatusCode != 200 {
		log.Fatalf("not results.")
	}

	jsonBody := string(responseBody)

	// response map data
	var response models.ResultApple

	json.Unmarshal([]byte(jsonBody), &response)

	ch <- response

	wg.Done()
}
