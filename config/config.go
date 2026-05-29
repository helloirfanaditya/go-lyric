package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	SpotifyClientID     string
	SpotifyClientSecret string
	SpotifyRedirectURL  string
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		SpotifyClientID:     viper.GetString("SPOTIFY_CLIENT_ID"),
		SpotifyClientSecret: viper.GetString("SPOTIFY_CLIENT_SECRET"),
		SpotifyRedirectURL:  viper.GetString("SPOTIFY_REDIRECT_URL"),
	}, nil
}
