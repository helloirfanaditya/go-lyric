package spotify

import (
	"encoding/json"
	"os"

	"golang.org/x/oauth2"
)

const tokenFile = ".spotify_token.json"

func SaveToken(token *oauth2.Token) error {
	f, err := os.Create(tokenFile)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(token)
}

func LoadToken() (*oauth2.Token, error) {
	f, err := os.Open(tokenFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var token oauth2.Token

	if err := json.NewDecoder(f).Decode(&token); err != nil {
		return nil, err
	}

	return &token, nil
}
