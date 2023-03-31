package caliban

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type chatOpt func(*helixParams)
type chatOpts struct {
	// TODO
}

var ChtOpts *chatOpts

func init() {
	ChtOpts = &chatOpts{
		// TODO
	}
}

// GetChatters
func (S *Session) GetChatters(bcastId, modId string, opts ...chatOpt) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	p.Add("moderator_id", modId)
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("chat/chatters", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// GetChannelEmotes
func (S *Session) GetChannelEmotes(bcastId string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	url := newHelixURL("chat/emotes", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// GetGlobalEmotes
func (S *Session) GetGlobalEmotes() *helixResp {
	url := newHelixURL("chat/emotes/global", "")
	return S.Do(http.MethodGet, url.String(), nil)
}

// GetEmoteSets
func (S *Session) GetEmoteSets(emoteSetIds ...string) *helixResp {
	p := newHelixParams()
	for _, id := range emoteSetIds {
		p.Add("emote_set_id", id)
	}
	url := newHelixURL("chat/emotes/set", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// GetChannelChatBadges
func (S *Session) GetChannelChatBadges(bcastId string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	url := newHelixURL("chat/badges", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// GetGlobalChatBadges
func (S *Session) GetGlobalChatBadges() *helixResp {
	url := newHelixURL("chat/badges/global", "")
	return S.Do(http.MethodGet, url.String(), nil)
}

// GetChatSettings
func (S *Session) GetChatSettings(bcastId string, opts ...chatOpt) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("chat/settings", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// UpdateChatSettings
func (S *Session) UpdateChatSettings(bcastId, modId string, opts ...chatOpt) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	for _, opt := range opts {
		opt(p)
	}
	rawBody, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("chat/settings", p.Encode())
	return S.Do(http.MethodPatch, url.String(), bytes.NewBuffer(rawBody))
}

// SendChatAnnouncement
func (S *Session) SendChatAnnouncement(bcastId, modId string, opts ...chatOpt) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	p.Add("moderator_id", modId)
	for _, opt := range opts {
		opt(p)
	}
	rawBody, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("chat/announcements", p.Encode())
	return S.Do(http.MethodPost, url.String(), bytes.NewBuffer(rawBody))
}

// SendShoutout
func (S *Session) SendShoutout(frmBcastId, toBcastId, modId string) *helixResp {
	p := newHelixParams()
	p.Add("from_broadcaster_id", frmBcastId)
	p.Add("to_broadcaster_id", toBcastId)
	p.Add("moderator_id", modId)
	url := newHelixURL("chat/shoutouts", p.Encode())
	return S.Do(http.MethodPost, url.String(), nil)
}

// GetUserChatColor
func (S *Session) GetUserChatColor(userIds ...string) *helixResp {
	p := newHelixParams()
	for _, id := range userIds {
		p.Add("user_id", id)
	}
	url := newHelixURL("chat/color", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// UpdateUserChatColor
func (S *Session) UpdateUserChatColor(userId, color string) *helixResp {
	p := newHelixParams()
	p.Add("user_id", userId)
	p.Add("color", color)
	url := newHelixURL("chat/color", p.Encode())
	return S.Do(http.MethodPut, url.String(), nil)
}
