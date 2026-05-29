package cache

import (
	"os"
	"path/filepath"
	"strings"
)

func Get(track, artist string) (string, bool) {
	name := sanitize(artist + "-" + track)

	data, err := os.ReadFile(
		filepath.Join("song_meta", name+".lrc"),
	)

	if err != nil {
		return "", false
	}

	return string(data), true
}

func Save(track, artist, lyrics string) error {
	_ = os.MkdirAll("song_meta", 0755)

	name := sanitize(artist + "-" + track)

	return os.WriteFile(
		filepath.Join("song_meta", name+".lrc"),
		[]byte(lyrics),
		0644,
	)
}

func sanitize(s string) string {
	s = strings.ReplaceAll(s, "/", "_")
	s = strings.ReplaceAll(s, "\\", "_")
	return s
}
