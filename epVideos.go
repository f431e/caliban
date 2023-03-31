package caliban

import (
	"net/http"
	"strconv"
)

type videosOpt func(*helixParams)
type videosOpts struct {
	ByVideoId func(...string) videosOpt
	Count     func(int) videosOpt
	ForGameId func(string) videosOpt
	ForUserId func(string) videosOpt
	Lang      func(string) videosOpt
	Next      func(int, string) videosOpt
	Only      func(videoType) videosOpt
	Period    func(timePeriod) videosOpt
	Prev      func(int, string) videosOpt
	TopTrend  func(timePeriod) videosOpt
	TopViews  func(timePeriod) videosOpt
}

var VidOpts *videosOpts

func init() {
	VidOpts = &videosOpts{
		ByVideoId: func(ids ...string) videosOpt {
			return func(p *helixParams) {
				for _, id := range ids {
					p.Add("id", id)
				}
			}
		},
		Count: func(n int) videosOpt {
			return func(p *helixParams) {
				p.Add("first", strconv.Itoa(n))
			}
		},
		ForGameId: func(id string) videosOpt {
			return func(p *helixParams) {
				p.Add("game_id", id)
			}
		},
		ForUserId: func(id string) videosOpt {
			return func(p *helixParams) {
				p.Add("user_id", id)
			}
		},
		Lang: func(l string) videosOpt {
			return func(p *helixParams) {
				p.Add("language", l)
			}
		},
		Next: func(n int, cursor string) videosOpt {
			return func(p *helixParams) {
				p.Add("first", strconv.Itoa(n))
				p.Add("after", cursor)
			}
		},
		Only: func(typ videoType) videosOpt {
			return func(p *helixParams) {
				p.Add("type", typ.String())
			}
		},
		Period: func(per timePeriod) videosOpt {
			return func(p *helixParams) {
				p.Add("period", per.String())
			}
		},
		Prev: func(n int, cursor string) videosOpt {
			return func(p *helixParams) {
				p.Add("first", strconv.Itoa(n))
				p.Add("before", cursor)
			}
		},
		TopTrend: func(per timePeriod) videosOpt {
			return func(p *helixParams) {
				p.Add("period", per.String())
				p.Add("sort", "trending")
			}
		},
		TopViews: func(per timePeriod) videosOpt {
			return func(p *helixParams) {
				p.Add("period", per.String())
				p.Add("sort", "views")
			}
		},
	}
}

// GetVideos
func (S *Session) GetVideos(opts ...videosOpt) *helixResp {
	p := newHelixParams()
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("videos", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// DeleteVideos
func (S *Session) DeleteVideos(opts ...videosOpt) *helixResp {
	p := newHelixParams()
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("videos", p.Encode())
	return S.Do(http.MethodDelete, url.String(), nil)
}
