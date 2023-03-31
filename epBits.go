package caliban

import (
	"net/http"
	"strconv"
)

type bitsOption func(*helixParams)
type bitsOptions struct {
	Count      func(n int) streamsOpt
	GameIds    func(gameIds ...string) streamsOpt
	UserIds    func(ids ...string) streamsOpt
	UserLogins func(logins ...string) streamsOpt
	LiveOnly   func() streamsOpt
	Langs      func(langs ...string) streamsOpt
	Next       func(n int, cursor string) streamsOpt
	Prev       func(n int, cursor string) streamsOpt
}

var BitsOpts *bitsOptions

func init() {
	BitsOpts = &bitsOptions{
		Count: func(n int) streamsOpt {
			return func(p *helixParams) {
				p.Add("first", strconv.Itoa(n))
			}
		},
		GameIds: func(gameIds ...string) streamsOpt {
			return func(p *helixParams) {
				for _, gameId := range gameIds {
					p.Add("game_id", gameId)
				}
			}
		},
		UserIds: func(ids ...string) streamsOpt {
			return func(p *helixParams) {
				for _, id := range ids {
					p.Add("user_id", id)
				}
			}
		},
		UserLogins: func(logins ...string) streamsOpt {
			return func(p *helixParams) {
				for _, login := range logins {
					p.Add("user_login", login)
				}
			}
		},
		Langs: func(langs ...string) streamsOpt {
			return func(p *helixParams) {
				for _, lang := range langs {
					p.Add("language", lang)
				}
			}
		},
		LiveOnly: func() streamsOpt {
			return func(p *helixParams) {
				p.Add("type", "live")
			}
		},
		Next: func(n int, cursor string) streamsOpt {
			return func(p *helixParams) {
				p.Add("first", strconv.Itoa(n))
				p.Add("after", cursor)
			}
		},
		Prev: func(n int, cursor string) streamsOpt {
			return func(p *helixParams) {
				p.Add("first", strconv.Itoa(n))
				p.Add("before", cursor)
			}
		},
	}
}

// GetBitsLeaderboard
func (S *Session) GetBitsLeaderboard(opts ...bitsOption) *helixResp {
	p := newHelixParams()
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("bits/leaderboard", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// GetCheermotes
func (S *Session) GetCheermotes(bcastId string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	url := newHelixURL("bits/cheermotes", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// GetExtensionTransactions
func (S *Session) GetExtensionTransactions(opts ...bitsOption) *helixResp {
	p := newHelixParams()
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("extensions/transactions", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}
