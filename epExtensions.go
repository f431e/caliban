package caliban

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type extensionOpt func(*helixParams)
type extensionOpts struct {
	// TODO
}

var ExtOpts *extensionOpts

func init() {
	ExtOpts = &extensionOpts{
		// TODO
	}
}

// GetExtensionConfigurationSegment
func (S *Session) GetExtensionConfigurationSegment(opts ...extensionOpt) *helixResp {
	p := newHelixParams()
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("extensions/configurations", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// SetExtensionConfigurationSegment
func (S *Session) SetExtensionConfigurationSegment(opts ...extensionOpt) *helixResp {
	p := newHelixParams()
	for _, opt := range opts {
		opt(p)
	}
	rawBody, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("extensions/required_configuration", "")
	return S.Do(http.MethodPut, url.String(), bytes.NewBuffer(rawBody))
}

// SetExtensionRequiredConfiguration
func (S *Session) SetExtensionRequiredConfiguration(bcastId string, opts ...extensionOpt) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("extensions/required_configuration", p.Encode())
	rawBody, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	return S.Do(http.MethodPut, url.String(), bytes.NewBuffer(rawBody))
}

// SendExtensionPubSubMessage
func (S *Session) SendExtensionPubSubMessage(opts ...extensionOpt) *helixResp {
	p := newHelixParams()
	for _, opt := range opts {
		opt(p)
	}
	rawBody, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("extensions/pubsub", p.Encode())
	return S.Do(http.MethodPost, url.String(), bytes.NewBuffer(rawBody))
}

// GetExtensionLiveChannels
func (S *Session) GetExtensionLiveChannels(extId, after string, first int) *helixResp {
	p := newHelixParams()
	p.Add("extension_id", extId)
	p.Add("first", strconv.Itoa(first))
	p.Add("after", after)
	url := newHelixURL("extensions/live", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// GetExtensionSecrets
func (S *Session) GetExtensionSecrets() *helixResp {
	url := newHelixURL("extensions/jwt/secrets", "")
	return S.Do(http.MethodGet, url.String(), nil)
}

// CreateExtensionSecret
func (S *Session) CreateExtensionSecret(extId string, delay int) *helixResp {
	p := newHelixParams()
	p.Add("extension_id", extId)
	p.Add("delay", strconv.Itoa(delay))
	url := newHelixURL("extensions/jwt/secrets", p.Encode())
	return S.Do(http.MethodPost, url.String(), nil)
}

// SendExtensionChatMessage
func (S *Session) SendExtensionChatMessage(bcastId string, opts ...extensionOpt) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	for _, opt := range opts {
		opt(p)
	}
	rawBody, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("extensions/chat", p.Encode())
	return S.Do(http.MethodPost, url.String(), bytes.NewBuffer(rawBody))
}

// GetExtensions
func (S *Session) GetExtensions(extId, extVer string) *helixResp {
	p := newHelixParams()
	p.Add("extension_id", extId)
	p.Add("extension_version", extVer)
	url := newHelixURL("extensions", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// GetReleasedExtensions
func (S *Session) GetReleasedExtensions(extId, extVer string) *helixResp {
	p := newHelixParams()
	p.Add("extension_id", extId)
	p.Add("extension_version", extVer)
	url := newHelixURL("extensions/released", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// GetExtensionBitsProducts
func (S *Session) GetExtensionBitsProducts(includeAll bool) *helixResp {
	p := newHelixParams()
	p.Add("should_include_all", strconv.FormatBool(includeAll))
	url := newHelixURL("bits/extensions", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// UpdateExtensionBitsProduct
func (S *Session) UpdateExtensionBitsProduct(opts ...extensionOpt) *helixResp {
	p := newHelixParams()
	for _, opt := range opts {
		opt(p)
	}
	rawBody, err := json.Marshal(p.body)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("bits/extensions", "")
	return S.Do(http.MethodPut, url.String(), bytes.NewBuffer(rawBody))
}
