package caliban

import (
	"net/http"
)

type subsOpt func(*helixParams)
type subsOpts struct {
	// TODO
}

var SubsOpts *subsOpts

func init() {
	SubsOpts = &subsOpts{
		// TODO
	}
}

// GetBroadcasterSubscriptions
func (S *Session) GetBroadcasterSubscriptions(bcastId string, opts ...subsOpt) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("subscriptions", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// CheckUserSubscription
func (S *Session) CheckUserSubscription(bcastId, userId string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	p.Add("user_id", userId)
	url := newHelixURL("subscriptions/user", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}
