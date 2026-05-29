package spotify

type Track struct {
	Name       string
	Artist     string
	ProgressMs int64
}

func CurrentTrack() (*Track, error) {
	state, err := Client.PlayerState(Context())
	if err != nil {
		return nil, err
	}

	if state == nil || state.Item == nil {
		return nil, nil
	}

	var artist string

	if len(state.Item.Artists) > 0 {
		artist = state.Item.Artists[0].Name
	}

	return &Track{
		Name:       state.Item.Name,
		Artist:     artist,
		ProgressMs: int64(state.Progress),
	}, nil
}
