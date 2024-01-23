package models

type LastFMTrack struct {
	Artist struct {
		Name string `json:"name"`
		Info string `json:"info"`
	} `json:"artist"`
	Name string `json:"name"`
	MusixmatchLyrics
}

type MusixmatchLyrics struct {
	Lyrics struct {
		Body string `json:"lyrics_body"`
	} `json:"lyrics"`
}
