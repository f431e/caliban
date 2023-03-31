package caliban

import (
	"net/http"
	"strconv"
)

type searchOption func(*helixParams)
type searchOptions struct {
	Count func(int) searchOption
	Next  func(string) searchOption
	Live  func(bool) searchOption
}

var SrchOpts *searchOptions

func init() {
	SrchOpts = &searchOptions{
		Count: func(n int) searchOption {
			return func(p *helixParams) {
				p.Add("first", strconv.Itoa(n))
			}
		},
		Next: func(cur string) searchOption {
			return func(p *helixParams) {
				p.Add("after", cur)
			}
		},
		Live: func(live bool) searchOption {
			return func(p *helixParams) {
				p.Add("live_only", strconv.FormatBool(live))
			}
		},
	}
}

// SearchCategories
func (S *Session) SearchCategories(query string, opts ...searchOption) *helixResp {
	p := newHelixParams()
	p.Add("query", query)
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("search/categories", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// SearchChannels
func (S *Session) SearchChannels(query string, opts ...searchOption) *helixResp {
	p := newHelixParams()
	p.Add("query", query)
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("search/channels", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}
