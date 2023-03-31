package caliban

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type scheduleOpt func(*helixParams)
type scheduleOpts struct {
	// TODO
}

var SchdOpts *scheduleOpts

func init() {
	SchdOpts = &scheduleOpts{
		// TODO
	}
}

// GetChannelStreamSchedule
func (S *Session) GetChannelStreamSchedule(bcastId string, opts ...scheduleOpt) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("schedule", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// GetChannelICalendar
func (S *Session) GetChannelICalendar(bcastId string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	url := newHelixURL("icalendar", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// UpdateChannelStreamSchedule
func (S *Session) UpdateChannelStreamSchedule(bcastId string, opts ...scheduleOpt) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("schedule/settings", p.Encode())
	return S.Do(http.MethodPatch, url.String(), nil)
}

// CreateChannelStreamScheduleSegment
func (S *Session) CreateChannelStreamScheduleSegment(bcastId string, opts ...scheduleOpt) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	for _, opt := range opts {
		opt(p)
	}
	rawBody, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("schedule/segment", p.Encode())
	return S.Do(http.MethodPost, url.String(), bytes.NewBuffer(rawBody))
}

// UpdateChannelStreamScheduleSegment
func (S *Session) UpdateChannelStreamScheduleSegment(bcastId, id string, opts ...scheduleOpt) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	for _, opt := range opts {
		opt(p)
	}
	rawBody, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("schedule/segment", p.Encode())
	return S.Do(http.MethodPatch, url.String(), bytes.NewBuffer(rawBody))
}

// DeleteChannelStreamScheduleSegment
func (S *Session) DeleteChannelStreamScheduleSegment(bcastId, segId string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	p.Add("id", segId)
	url := newHelixURL("schedule/segment", p.Encode())
	return S.Do(http.MethodDelete, url.String(), nil)
}
