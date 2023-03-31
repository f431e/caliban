package caliban

import (
	"net/http"
)

// GetCreatorGoals
func (S *Session) GetCreatorGoals(bcastId string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	url := newHelixURL("goals", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}
