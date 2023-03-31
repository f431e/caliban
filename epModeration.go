package caliban

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type moderationOpt func(*helixParams)
type moderationOpts struct {
	// TODO
}

var ModOpts *moderationOpts

func init() {
	ModOpts = &moderationOpts{
		// TODO
	}
}

// CheckAutoModStatus
func (S *Session) CheckAutoModStatus(bcastId string, opts ...moderationOpt) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	rawBody, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("moderation/enforcements/status", p.Encode())
	return S.Do(http.MethodPost, url.String(), bytes.NewBuffer(rawBody))
}

// ManageHeldAutoModMessages
func (S *Session) ManageHeldAutoModMessages(userId, msgId, action string) *helixResp {
	p := newHelixParams()
	p.body["user_id"] = userId
	p.body["msg_id"] = msgId
	p.body["action"] = action
	rawBody, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("moderation/automod/message", "")
	return S.Do(http.MethodPost, url.String(), bytes.NewBuffer(rawBody))
}

// GetAutoModSettings
func (S *Session) GetAutoModSettings(bcastId, modId string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	p.Add("moderator_id", modId)
	url := newHelixURL("moderation/automod/settings", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// UpdateAutoModSettings - incomplete
func (S *Session) UpdateAutoModSettings(opts ...moderationOpt) *helixResp {
	p := newHelixParams()
	// p.Add("broadcaster_id", bcastId)
	// p.Add("moderator_id", modId)
	rawBody, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("moderation/automod/settings", p.Encode())
	return S.Do(http.MethodPut, url.String(), bytes.NewBuffer(rawBody))
}

// GetBannedUsers
func (S *Session) GetBannedUsers(bcastId string, opts ...moderationOpt) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	for _, opt := range opts {
		opt(p)
	}
	rawBody, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("moderation/banned", p.Encode())
	return S.Do(http.MethodGet, url.String(), bytes.NewBuffer(rawBody))
}

// BanUser
func (S *Session) BanUser(bcastId, modId string, opts ...moderationOpt) *helixResp {
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
	url := newHelixURL("moderation/bans", p.Encode())
	return S.Do(http.MethodPost, url.String(), bytes.NewBuffer(rawBody))
}

// UnbanUser
func (S *Session) UnbanUser(bcastId, modId, userId string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	p.Add("moderator_id", modId)
	p.Add("user_id", userId)
	url := newHelixURL("moderation/bans", p.Encode())
	return S.Do(http.MethodDelete, url.String(), nil)
}

// GetBlockedTerms
func (S *Session) GetBlockedTerms(bcastId, modId string, opts ...moderationOpt) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	p.Add("moderator_id", modId)
	url := newHelixURL("moderation/blocked_terms", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// AddBlockedTerm
func (S *Session) AddBlockedTerm(bcastId, modId, text string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	p.Add("moderator_id", modId)
	p.body["text"] = text
	rawBody, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("moderation/blocked_terms", p.Encode())
	return S.Do(http.MethodPost, url.String(), bytes.NewBuffer(rawBody))
}

// RemoveBlockedTerm
func (S *Session) RemoveBlockedTerm(bcastId, modId, id string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	p.Add("moderator_id", modId)
	p.Add("id", id)
	url := newHelixURL("moderation/blocked_terms", p.Encode())
	return S.Do(http.MethodDelete, url.String(), nil)
}

// DeleteChatMessages
func (S *Session) DeleteChatMessages(bcastId, modId, messId string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	p.Add("moderator_id", modId)
	p.Add("message_id", messId)
	url := newHelixURL("moderation/chat", p.Encode())
	return S.Do(http.MethodDelete, url.String(), nil)
}

// GetModerators
func (S *Session) GetModerators(bcastId string, opts ...moderationOpt) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("moderation/moderators", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// AddChannelModerator
func (S *Session) AddChannelModerator(bcastId, userId string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	p.Add("user_id", userId)
	url := newHelixURL("moderation/moderators", p.Encode())
	return S.Do(http.MethodPost, url.String(), nil)
}

// RemoveChannelModerator
func (S *Session) RemoveChannelModerator(bcastId, userId string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	p.Add("user_id", userId)
	url := newHelixURL("moderation/moderators", p.Encode())
	return S.Do(http.MethodDelete, url.String(), nil)
}

// GetVIPs
func (S *Session) GetVIPs(bcastId string, opts ...moderationOpt) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("channels/vips", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// AddChannelVIP
func (S *Session) AddChannelVIP(bcastId, userId string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	p.Add("user_id", userId)
	url := newHelixURL("channels/vips", p.Encode())
	return S.Do(http.MethodPost, url.String(), nil)
}

// RemoveChannelVIP
func (S *Session) RemoveChannelVIP(bcastId, userId string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	p.Add("user_id", userId)
	url := newHelixURL("channels/vips", p.Encode())
	return S.Do(http.MethodDelete, url.String(), nil)
}

// UpdateShieldModeStatus
func (S *Session) UpdateShieldModeStatus(bcastId, modId string, active bool) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	p.Add("moderator_id", modId)
	p.body["is_active"] = active
	rawBody, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("moderation/shield_mode", p.Encode())
	return S.Do(http.MethodPut, url.String(), bytes.NewBuffer(rawBody))
}

// GetShieldModeStatus
func (S *Session) GetShieldModeStatus(bcastId, modId string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	p.Add("moderator_id", modId)
	url := newHelixURL("moderation/shield_mode", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}
