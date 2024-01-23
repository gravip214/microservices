// musixmatch_service.go
package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"microservice/models"
	"net/http"
	"net/url"
	"os"
	"sync"
)

func GetLyrics(artist, track string) (models.MusixmatchLyrics, error) {
	apiKey := os.Getenv("MUSIXMATCH_API_KEY")
	url := fmt.Sprintf("https://api.musixmatch.com/ws/1.1/matcher.lyrics.get?format=json&callback=callback&q_artist=%s&q_track=%s&apikey=%s", url.QueryEscape(artist), url.QueryEscape(track), url.QueryEscape(apiKey))

	resp, err := http.Get(url)
	if err != nil {
		return models.MusixmatchLyrics{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.MusixmatchLyrics{}, err
	}

	var response struct {
		Message struct {
			Body models.MusixmatchLyrics `json:"body"`
		} `json:"message"`
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("GetLyrics:", string(body))
		fmt.Println("errors Lyrics:", err.Error())
		return models.MusixmatchLyrics{}, err
	}

	return response.Message.Body, nil
}
func GetLyricsInfo(wg *sync.WaitGroup, artist, track string) models.MusixmatchLyrics {
	defer wg.Done()
	if val, err := GetLyrics(artist, track); err == nil {
		return val
	}
	return models.MusixmatchLyrics{}
}
