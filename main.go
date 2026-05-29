package main

import (
	"fmt"
	"time"

	"github.com/helloirfanaditya/go-lyric/config"
	"github.com/helloirfanaditya/go-lyric/internal/cache"
	"github.com/helloirfanaditya/go-lyric/internal/lyrics"
	"github.com/helloirfanaditya/go-lyric/internal/spotify"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	spotify.Init(
		cfg.SpotifyClientID,
		cfg.SpotifyClientSecret,
		cfg.SpotifyRedirectURL,
	)

	if !spotify.LoginWithSavedToken() {
		if err := spotify.Login(); err != nil {
			panic(err)
		}
	}

	var (
		lastSong string
		lines    []lyrics.LyricLine
	)

	for {
		track, err := spotify.CurrentTrack()
		if err != nil {
			time.Sleep(time.Second)
			continue
		}

		if track == nil {
			time.Sleep(time.Second)
			continue
		}

		songKey := track.Artist + "-" + track.Name

		if songKey != lastSong {
			lastSong = songKey

			lines = nil

			if cachedLyrics, found := cache.Get(
				track.Name,
				track.Artist,
			); found {

				lines = lyrics.ParseLRC(cachedLyrics)

			} else {

				syncedLyrics, err := lyrics.Get(
					track.Name,
					track.Artist,
				)
				if err != nil {
					continue
				}

				_ = cache.Save(
					track.Name,
					track.Artist,
					syncedLyrics,
				)

				lines = lyrics.ParseLRC(
					syncedLyrics,
				)
			}
		}

		if len(lines) == 0 {
			fmt.Print("\033[H\033[2J")
			fmt.Printf("Now Playing : %s - %s\n\n", track.Artist, track.Name)
			fmt.Println("Loading lyrics...")

			time.Sleep(200 * time.Millisecond)
			continue
		}

		current := lyrics.CurrentLine(
			track.ProgressMs,
			lines,
		)

		fmt.Print("\033[H\033[2J")

		fmt.Printf("Now Playing : %s - %s\n\n",
			track.Artist,
			track.Name,
		)

		// Auto-scroll window
		const viewportSize = 15

		start := current - viewportSize/2
		if start < 0 {
			start = 0
		}

		end := start + viewportSize
		if end > len(lines) {
			end = len(lines)
		}

		if end-start < viewportSize {
			start = end - viewportSize
			if start < 0 {
				start = 0
			}
		}

		for i := start; i < end; i++ {
			if i == current {
				fmt.Printf("\033[1;32m▶ %s\033[0m\n", lines[i].Text)
			} else {
				fmt.Printf("  %s\n", lines[i].Text)
			}
		}

	}
}
