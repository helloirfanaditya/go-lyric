# Go Lyric 🎵

A lightweight cross-platform terminal lyrics viewer for Spotify, built with Go.

Go Lyric automatically detects the song currently playing on Spotify, fetches synchronized lyrics from LRCLIB, caches them locally, and displays them in a clean karaoke-style terminal interface.

## Features

* Spotify integration
* Real-time synced lyrics
* Karaoke-style lyric highlighting
* Auto-scrolling lyrics
* Local lyric caching
* Spotify token persistence
* Cross-platform support

  * Windows
  * macOS
  * Linux
* Lightweight and fast
* Built entirely in Go

## Preview

```text
🎵 YOASOBI - Idol

  無敵の笑顔で荒らすメディア
  知りたいその秘密ミステリアス
▶ 天才的なアイドル様
  今日何食べた？
  好きな本は？
```

## Requirements

* Go 1.24+
* Spotify Premium account (recommended)
* Spotify Desktop App running

## Setup

### 1. Create a Spotify Application

Visit:

https://developer.spotify.com/dashboard

Create a new application and configure:

```text
Redirect URI:
http://127.0.0.1:8888/callback
```

Copy:

* Client ID
* Client Secret

### 2. Create Environment File

Create a `.env` file in the project root.

```env
SPOTIFY_CLIENT_ID=YOUR_CLIENT_ID
SPOTIFY_CLIENT_SECRET=YOUR_CLIENT_SECRET
SPOTIFY_REDIRECT_URL=http://127.0.0.1:8888/callback
```

### 3. Install Dependencies

```bash
go mod tidy
```

### 4. Run

```bash
go run . // or make run
```

On first launch, a browser window will open for Spotify authentication.

After successful login, a token will be stored locally and reused for future launches.

## Build

### Windows

```bash
GOOS=windows GOARCH=amd64 go build -o go-lyric.exe
```

### macOS (Apple Silicon)

```bash
GOOS=darwin GOARCH=arm64 go build -o go-lyric
```

### Linux

```bash
GOOS=linux GOARCH=amd64 go build -o go-lyric
```

## Cache

Lyrics are automatically cached locally.

```text
song_meta/
├── YOASOBI-Idol.lrc
├── Aimer-Zankyou_Sanka.lrc
└── ...
```

When a song is played again, Go Lyric loads the lyrics directly from cache instead of requesting LRCLIB.

## Technologies

* Go
* Spotify Web API
* LRCLIB
* Viper

## Roadmap

* [ ] Bubble Tea UI
* [ ] Album artwork support
* [ ] Theme customization
* [ ] Keyboard shortcuts
* [ ] Search lyrics manually
* [ ] Offline mode
* [ ] Desktop notifications
* [ ] Multi-language lyrics support

## License

MIT License

## Disclaimer

This project is not affiliated with Spotify.

Lyrics are provided by LRCLIB and remain the property of their respective copyright holders.
