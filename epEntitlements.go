package caliban

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type entitlementsOpt func(*helixParams)
type entitlementsOpts struct {
	// TODO
}

var EntOpts *entitlementsOpts

func init() {
	EntOpts = &entitlementsOpts{
		// TODO
	}
}

// GetDropsEntitlements
func (S *Session) GetDropsEntitlements(opts ...entitlementsOpt) *helixResp {
	p := newHelixParams()
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("entitlements/drops", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// UpdateDropsEntitlements
func (S *Session) UpdateDropsEntitlements(opts ...entitlementsOpt) *helixResp {
	p := newHelixParams()
	for _, opt := range opts {
		opt(p)
	}
	rawBody, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("entitlements/drops", "")
	return S.Do(http.MethodPatch, url.String(), bytes.NewBuffer(rawBody))
}
