package caliban

import (
	"net/http"
	"strconv"
)

type hypeTrnOpt func(*helixParams)
type hypeTrnOpts struct {
	After func(string) hypeTrnOpt
	First func(int) hypeTrnOpt
}

var HypeTrnOpts *hypeTrnOpts

func init() {
	HypeTrnOpts = &hypeTrnOpts{
		After: func(cursor string) hypeTrnOpt {
			return func(p *helixParams) {
				p.Add("after", cursor)
			}
		},
		First: func(n int) hypeTrnOpt {
			return func(p *helixParams) {
				p.Add("first", strconv.Itoa(n))
			}
		},
	}
}

// GetHypeTrainEvents
func (S *Session) GetHypeTrainEvents(bcastId string, opts ...hypeTrnOpt) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("hypetrain/events", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}
