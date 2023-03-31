package caliban

import (
	"net/http"
	"strconv"
)

type analyticsOpt func(*helixParams)
type analyticsOpts struct {
	Count          func(int) analyticsOpt
	End            func(string) analyticsOpt
	ForExtensionId func(string) analyticsOpt
	ForGameId      func(string) analyticsOpt
	Next           func(int, string) analyticsOpt
	ReportType     func(string) analyticsOpt
	Start          func(string) analyticsOpt
}

var AnaOpts *analyticsOpts

func init() {
	AnaOpts = &analyticsOpts{
		Count: func(n int) analyticsOpt {
			return func(p *helixParams) {
				p.Add("first", strconv.Itoa(n))
			}
		},
		End: func(end string) analyticsOpt {
			return func(p *helixParams) {
				p.Add("ended_at", end)
			}
		},
		ForExtensionId: func(id string) analyticsOpt {
			return func(p *helixParams) {
				p.Set("extension_id", id)
			}
		},
		ForGameId: func(id string) analyticsOpt {
			return func(p *helixParams) {
				p.Add("game_id", id)
			}
		},
		Next: func(n int, cursor string) analyticsOpt {
			return func(p *helixParams) {
				p.Add("first", strconv.Itoa(n))
				p.Add("after", cursor)
			}
		},
		ReportType: func(typ string) analyticsOpt {
			return func(p *helixParams) {
				p.Set("type", typ)
			}
		},
		Start: func(start string) analyticsOpt {
			return func(p *helixParams) {
				p.Add("started_at", start)
			}
		},
	}
}

// GetExtensionAnalytics
func (S *Session) GetExtensionAnalytics(opts ...analyticsOpt) *helixResp {
	p := newHelixParams()
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("analytics/extensions", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// GetGameAnalytics
func (S *Session) GetGameAnalytics(opts ...analyticsOpt) *helixResp {
	p := newHelixParams()
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("analytics/games", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}
