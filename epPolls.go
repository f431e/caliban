package caliban

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type pollsOption func(*helixParams)
type pollsOptions struct {
	// TODO
}

var PollOpts *pollsOptions

func init() {
	PollOpts = &pollsOptions{
		// TODO
	}
}

// GetPolls
func (S *Session) GetPolls(bcastId string, opts ...pollsOption) *helixResp {
	p := newHelixParams()
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("polls", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// CreatePoll
func (S *Session) CreatePoll(bcastId string, opts ...pollsOption) *helixResp {
	p := newHelixParams()
	p.body["broadcaster_id"] = bcastId
	for _, opt := range opts {
		opt(p)
	}
	rawBody, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("polls", "")
	return S.Do(http.MethodPost, url.String(), bytes.NewBuffer(rawBody))
}

// EndPoll
func (S *Session) EndPoll(bcastId, id, status string) *helixResp {
	p := newHelixParams()
	p.body["broadcaster_id"] = bcastId
	p.body["poll_id"] = id
	p.body["status"] = status
	rawBody, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("polls", "")
	return S.Do(http.MethodPatch, url.String(), bytes.NewBuffer(rawBody))
}
