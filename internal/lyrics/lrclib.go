package lyrics

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type SearchResult struct {
	PlainLyrics  string `json:"plainLyrics"`
	SyncedLyrics string `json:"syncedLyrics"`
}

func Get(track, artist string) (string, error) {
	endpoint := fmt.Sprintf(
		"https://lrclib.net/api/search?track_name=%s&artist_name=%s",
		url.QueryEscape(track),
		url.QueryEscape(artist),
	)

	resp, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result []SearchResult

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result) == 0 {
		return "Lyrics not found", nil
	}

	if result[0].SyncedLyrics != "" {
		return result[0].SyncedLyrics, nil
	}

	return result[0].PlainLyrics, nil
}
