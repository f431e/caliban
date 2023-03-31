package caliban

import "net/http"

type charityOpt func(*helixParams)
type charityOpts struct {
	// TODO
}

var ChrtyOpts *charityOpts

func init() {
	ChrtyOpts = &charityOpts{
		// TODO
	}
}

// GetCharityCampaign
func (S *Session) GetCharityCampaign(bcastId string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	url := newHelixURL("charity/campaigns", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// GetCharityCampaignDonations
func (S *Session) GetCharityCampaignDonations(bcastId string, opts ...charityOpt) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("charity/donations", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}
