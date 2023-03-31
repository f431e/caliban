package caliban

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type predictionsOption func(*helixParams)
type predictionsOptions struct {
	// TODO
}

var PredOpts *predictionsOptions

func init() {
	PredOpts = &predictionsOptions{
		// TODO
	}
}

// GetPredictions
func (S *Session) GetPredictions(bcastId string, opts ...predictionsOption) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("predictions", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// CreatePrediction
func (S *Session) CreatePrediction(bcastId string, opts ...predictionsOption) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	for _, opt := range opts {
		opt(p)
	}
	rawBody, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("predictions", p.Encode())
	return S.Do(http.MethodPost, url.String(), bytes.NewBuffer(rawBody))
}

// EndPrediction
func (S *Session) EndPrediction(bcastId string, opts ...predictionsOption) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	for _, opt := range opts {
		opt(p)
	}
	rawBody, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("predictions", p.Encode())
	return S.Do(http.MethodPatch, url.String(), bytes.NewBuffer(rawBody))
}
