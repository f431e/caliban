package caliban

import "net/http"

// StartRaid
func (S *Session) StartRaid(frmBcastId, toBcastId string) *helixResp {
	p := newHelixParams()
	p.Add("from_broadcaster_id", frmBcastId)
	p.Add("to_broadcaster_id", toBcastId)
	url := newHelixURL("raids", p.Encode())
	return S.Do(http.MethodPost, url.String(), nil)
}

// CancelRaid
func (S *Session) CancelRaid(bcastId string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	url := newHelixURL("raids", p.Encode())
	return S.Do(http.MethodDelete, url.String(), nil)
}
