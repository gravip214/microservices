// lastfm_service.go
package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"microservice/models"
	"net/http"
	"os"
	"sync"
)

func GetTopTrack(region string) (models.LastFMTrack, error) {
	var wg sync.WaitGroup
	apiKey := os.Getenv("LASTFM_API_KEY")
	url := fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=geo.gettoptracks&country=%s&api_key=%s&format=json", region, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return models.LastFMTrack{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.LastFMTrack{}, err
	}

	var response struct {
		Tracks struct {
			Track []models.LastFMTrack `json:"track"`
		} `json:"tracks"`
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return models.LastFMTrack{}, err
	}

	if len(response.Tracks.Track) == 0 {
		return models.LastFMTrack{}, fmt.Errorf("no top tracks found for the region")
	}
	if len(response.Tracks.Track) > 0 {
		value := response.Tracks.Track[0]
		wg.Add(1)
		go func() {
			response.Tracks.Track[0].Artist.Info = GetArtistInfo(&wg, value.Artist.Name)
		}()
		wg.Add(1)
		go func() {
			response.Tracks.Track[0].Lyrics = GetLyricsInfo(&wg, value.Artist.Name, value.Name).Lyrics
		}()

	}
	wg.Wait()
	return response.Tracks.Track[0], nil
}
