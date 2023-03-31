package caliban

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type chanPointsOpt func(*helixParams)
type chanPointsOpts struct {
	// TODO
}

var ChanPointsOpts *chanPointsOpts

func init() {
	ChanPointsOpts = &chanPointsOpts{
		// TODO
	}
}

// CreateCustomRewards
func (S *Session) CreateCustomRewards(bcastId string, opts ...chanPointsOpt) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	for _, opt := range opts {
		opt(p)
	}
	body, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("channel_points/custom_rewards", p.Encode())
	return S.Do(http.MethodPost, url.String(), bytes.NewBuffer(body))
}

// DeleteCustomReward
func (S *Session) DeleteCustomReward(bCastId, id string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bCastId)
	p.Add("id", id)
	url := newHelixURL("channel_points/custom_rewards", p.Encode())
	return S.Do(http.MethodDelete, url.String(), nil)
}

// GetCustomReward
func (S *Session) GetCustomReward(bCastId, onlyManRewards string, ids ...string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bCastId)
	p.Add("only_manageable_rewards", onlyManRewards)
	for _, id := range ids {
		p.Add("id", id)
	}
	url := newHelixURL("channel_points/custom_rewards", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// GetCustomRewardRedemption
func (S *Session) GetCustomRewardRedemption(bcastId string, opts ...chanPointsOpt) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("channel_points/custom_rewards/redemptions", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// UpdateCustomReward
func (S *Session) UpdateCustomReward(bcastId, id string, opts ...chanPointsOpt) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	p.Add("id", id)
	for _, opt := range opts {
		opt(p)
	}
	body, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("channel_points/custom_rewards", p.Encode())
	return S.Do(http.MethodPatch, url.String(), bytes.NewBuffer(body))
}

// UpdateRedemptionStatus
func (S *Session) UpdateRedemptionStatus(bcastId, rewardId, status string, ids ...string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	p.Add("reward_id", rewardId)
	for _, id := range ids {
		p.Add("id", id)
	}
	p.body["status"] = status
	body, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("channel_points/custom_rewards/redemptions", p.Encode())
	return S.Do(http.MethodPatch, url.String(), bytes.NewBuffer(body))
}
