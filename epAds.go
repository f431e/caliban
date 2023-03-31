package caliban

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// StartCommercial
func (S *Session) StartCommercial(bcastId string, length int) *helixResp {
	p := newHelixParams()
	p.body["broadcaster_id"] = bcastId
	p.body["length"] = length
	rawBody, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("channels/commercial", "")
	return S.Do(http.MethodPost, url.String(), bytes.NewBuffer(rawBody))
}
