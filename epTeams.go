package caliban

import (
	"net/http"
)

type teamsOpt func(*helixParams)
type teamsOpts struct {
	ByTeamName func(string) teamsOpt
	ByTeamId   func(string) teamsOpt
}

var TeamsOpts *teamsOpts

func init() {
	TeamsOpts = &teamsOpts{
		ByTeamName: func(name string) teamsOpt {
			return func(p *helixParams) {
				p.Add("name", name)
			}
		},
		ByTeamId: func(id string) teamsOpt {
			return func(p *helixParams) {
				p.Add("id", id)
			}
		},
	}
}

// GetChannelTeams
func (S *Session) GetChannelTeams(bcastId string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	url := newHelixURL("teams/channel", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// GetTeams
func (S *Session) GetTeams(opts ...teamsOpt) *helixResp {
	p := newHelixParams()
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("teams", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}
