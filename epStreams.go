package caliban

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type streamsOpt func(*helixParams)
type streamsOpts struct {
	Count        func(n int) streamsOpt
	ForGameId    func(gameIds ...string) streamsOpt
	ForUserId    func(ids ...string) streamsOpt
	ForUserLogin func(logins ...string) streamsOpt
	LiveOnly     func() streamsOpt
	Lang         func(langs ...string) streamsOpt
	Next         func(n int, cursor string) streamsOpt
	Prev         func(n int, cursor string) streamsOpt
}

var StrmOpts *streamsOpts

func init() {
	StrmOpts = &streamsOpts{
		Count: func(n int) streamsOpt {
			return func(p *helixParams) {
				p.Add("first", strconv.Itoa(n))
			}
		},
		ForGameId: func(gameIds ...string) streamsOpt {
			return func(p *helixParams) {
				for _, gameId := range gameIds {
					p.Add("game_id", gameId)
				}
			}
		},
		ForUserId: func(ids ...string) streamsOpt {
			return func(p *helixParams) {
				for _, id := range ids {
					p.Add("user_id", id)
				}
			}
		},
		ForUserLogin: func(logins ...string) streamsOpt {
			return func(p *helixParams) {
				for _, login := range logins {
					p.Add("user_login", login)
				}
			}
		},
		Lang: func(langs ...string) streamsOpt {
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

// GetStreamKey
func (S *Session) GetStreamKey(bcasterId string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcasterId)
	url := newHelixURL("streams/key", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// GetStreams
func (S *Session) GetStreams(opts ...streamsOpt) *helixResp {
	p := newHelixParams()
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("streams", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// GetFollowedStreams
func (S *Session) GetFollowedStreams(userId string, opts ...streamsOpt) *helixResp {
	p := newHelixParams()
	p.Add("user_id", userId)
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("streams/followed", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// CreateStreamMarker
func (S *Session) CreateStreamMarker(userId, desc string) *helixResp {
	p := newHelixParams()
	p.body["user_id"] = userId
	p.body["description"] = desc
	rawBody, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("streams/markers", "")
	return S.Do(http.MethodPost, url.String(), bytes.NewBuffer(rawBody))
}

// GetStreamMarkers
func (S *Session) GetStreamMarkers(opts ...streamsOpt) *helixResp {
	p := newHelixParams()
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("streams/markers", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}
