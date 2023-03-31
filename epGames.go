package caliban

import (
	"net/http"
)

type gamesOption func(*helixParams)
type gamesOptions struct {
	// TODO
}

var GamesOpt *gamesOptions

func init() {
	GamesOpt = &gamesOptions{
		// TODO
	}
}

// GetTopGames
func (S *Session) GetTopGames(opts ...gamesOption) *helixResp {
	p := newHelixParams()
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("games/top", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// GetGames
func (S *Session) GetGames(opts ...gamesOption) *helixResp {
	p := newHelixParams()
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("games", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}
