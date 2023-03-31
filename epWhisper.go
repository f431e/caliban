package caliban

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// SendWhisper
func (S *Session) SendWhisper(frmUsrId, toUsrId, mess string) *helixResp {
	p := newHelixParams()
	p.Add("from_user_id", frmUsrId)
	p.Add("to_user_id", toUsrId)
	p.body["message"] = mess
	rawBody, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("whispers", p.Encode())
	return S.Do(http.MethodPost, url.String(), bytes.NewBuffer(rawBody))
}
