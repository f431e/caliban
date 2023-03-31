package caliban

import (
	"net/http"
	"strconv"
)

type clipsOpt func(*helixParams)
type clipsOpts struct {
	ByClipId   func(...string) clipsOpt
	Count      func(int) clipsOpt
	End        func(string) clipsOpt
	ForBcastId func(string) clipsOpt
	ForGameId  func(string) clipsOpt
	Next       func(int, string) clipsOpt
	Prev       func(int, string) clipsOpt
	Start      func(string) clipsOpt
}

var ClpsOpts *clipsOpts

func init() {
	ClpsOpts = &clipsOpts{
		ByClipId: func(ids ...string) clipsOpt {
			return func(p *helixParams) {
				for _, id := range ids {
					p.Add("id", id)
				}
			}
		},
		Count: func(n int) clipsOpt {
			return func(p *helixParams) {
				p.Add("first", strconv.Itoa(n))
			}
		},
		End: func(end string) clipsOpt {
			return func(p *helixParams) {
				p.Add("ended_at", end)
			}
		},
		ForBcastId: func(id string) clipsOpt {
			return func(p *helixParams) {
				p.Add("broadcaster_id", id)
			}
		},
		ForGameId: func(id string) clipsOpt {
			return func(p *helixParams) {
				p.Add("game_id", id)
			}
		},
		Next: func(n int, cursor string) clipsOpt {
			return func(p *helixParams) {
				p.Add("first", strconv.Itoa(n))
				p.Add("after", cursor)
			}
		},
		Prev: func(n int, cursor string) clipsOpt {
			return func(p *helixParams) {
				p.Add("first", strconv.Itoa(n))
				p.Add("before", cursor)
			}
		},
		Start: func(start string) clipsOpt {
			return func(p *helixParams) {
				p.Add("started_at", start)
			}
		},
	}
}

func (S *Session) CreateClip(bcastId string, hasDelay bool) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	p.Add("has_delay", strconv.FormatBool(hasDelay))
	url := newHelixURL("clips", p.Encode())
	return S.Do(http.MethodPost, url.String(), nil)
}

func (S *Session) GetClips(opts ...clipsOpt) *helixResp {
	p := newHelixParams()
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("clips", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}
