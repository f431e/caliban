package caliban

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type eventSubOpt func(*helixParams)
type eventSubOpts struct {
}

var EventSubOpts *eventSubOpts

func init() {
	EventSubOpts = &eventSubOpts{}
}

// CreateEventSubSubscription
func (S *Session) CreateEventSubSubscription(opts ...eventSubOpt) *helixResp {
	p := newHelixParams()
	for _, opt := range opts {
		opt(p)
	}
	rawBody, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("eventsub/subscriptions", p.Encode())
	return S.Do(http.MethodPost, url.String(), bytes.NewBuffer(rawBody))
}

// DeleteEventSubSubscription
func (S *Session) DeleteEventSubSubscription(id string) *helixResp {
	p := newHelixParams()
	p.Add("id", id)
	url := newHelixURL("eventsub/subscription", p.Encode())
	return S.Do(http.MethodDelete, url.String(), nil)
}

// GetEventSubSubscriptions
func (S *Session) GetEventSubSubscriptions(opts ...eventSubOpt) *helixResp {
	p := newHelixParams()
	for _, opt := range opts {
		opt(p)
	}
	rawBody, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("eventsub/subscriptions", p.Encode())
	return S.Do(http.MethodGet, url.String(), bytes.NewBuffer(rawBody))
}
