// image_search_service.go
package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sync"
)

var artistinfo = make(map[string]string)
var mx = sync.Mutex{}

func GetArtistImage(artist string) (string, error) {
	customSearchAPIKey := os.Getenv("GOOGLE_CUSTOM_SEARCH_API_KEY")
	customSearchEngineID := os.Getenv("GOOGLE_CUSTOM_SEARCH_ENGINE_ID")

	url := fmt.Sprintf("https://www.googleapis.com/customsearch/v1?q=%s&key=%s&cx=%s&searchType=image&num=1", url.QueryEscape(artist), url.QueryEscape(customSearchAPIKey), url.QueryEscape(customSearchEngineID))

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response struct {
		Items []struct {
			Link string `json:"link"`
		} `json:"items"`
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("GetArtistImage: ", string(body))
		return "", err
	}

	if len(response.Items) == 0 {
		return "", fmt.Errorf("no images found for the artist")
	}

	return response.Items[0].Link, nil
}

func GetArtistInfo(wg *sync.WaitGroup, artist string) string {
	defer wg.Done()
	if data, ok := artistinfo[artist]; ok {
		return data
	}
	if val, err := GetArtistImage(artist); err == nil {
		mx.Lock()
		artistinfo[artist] = val
		mx.Unlock()
		return val
	} else {
		fmt.Println(err)
		return val
	}
	return ""
}
