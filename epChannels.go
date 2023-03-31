package caliban

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type channelsOpt func(*helixParams)
type channelsOpts struct {
	// TODO
}

var ChansOpts *channelsOpts

func init() {
	ChansOpts = &channelsOpts{
		// TODO
	}
}

// GetChannelInformation
func (S *Session) GetChannelInformation(bcastIds ...string) *helixResp {
	p := newHelixParams()
	for _, id := range bcastIds {
		p.Add("broadcaster_id", id)
	}
	url := newHelixURL("channels", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// ModifyChannelInformation
func (S *Session) ModifyChannelInformation(bcastrId string, opts ...channelsOpt) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastrId)
	for _, opt := range opts {
		opt(p)
	}
	body, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("channels", p.Encode())
	return S.Do(http.MethodPatch, url.String(), bytes.NewBuffer(body))
}

// GetChannelEditors
func (S *Session) GetChannelEditors(bcastId string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	url := newHelixURL("channels/editors", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// GetFollowedChannels
func (S *Session) GetFollowedChannels(userId string, opts ...channelsOpt) *helixResp {
	p := newHelixParams()
	p.Add("user_id", userId)
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("channels/followed", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// GetChannelFollowers
func (S *Session) GetChannelFollowers(bcastId string, opts ...channelsOpt) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("channels/followers", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}
