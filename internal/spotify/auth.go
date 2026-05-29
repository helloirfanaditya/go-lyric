package spotify

import (
	"context"
	"fmt"
	"net/http"

	gospotify "github.com/zmb3/spotify/v2"
	auth "github.com/zmb3/spotify/v2/auth"
)

var (
	Authenticator *auth.Authenticator
	Client        *gospotify.Client
)

func Init(clientID, clientSecret, redirectURL string) {
	Authenticator = auth.New(
		auth.WithClientID(clientID),
		auth.WithClientSecret(clientSecret),
		auth.WithRedirectURL(redirectURL),
		auth.WithScopes(
			auth.ScopeUserReadCurrentlyPlaying,
			auth.ScopeUserReadPlaybackState,
		),
	)
}

func Login() error {
	ch := make(chan *gospotify.Client)

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		token, err := Authenticator.Token(r.Context(), "state", r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := SaveToken(token); err != nil {
			fmt.Println("failed save token:", err)
		}

		client := gospotify.New(Authenticator.Client(r.Context(), token))

		fmt.Fprintln(w, "Spotify login success. You can close this tab.")

		ch <- client
	})

	go func() {
		_ = http.ListenAndServe(":8888", nil)
	}()

	url := Authenticator.AuthURL("state")

	fmt.Println("Open this URL:")
	fmt.Println(url)

	Client = <-ch

	return nil
}

func Context() context.Context {
	return context.Background()
}

func LoginWithSavedToken() bool {
	token, err := LoadToken()
	if err != nil {
		return false
	}

	Client = gospotify.New(
		Authenticator.Client(context.Background(), token),
	)

	return true
}
