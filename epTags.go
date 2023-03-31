package caliban

import (
	"net/http"
)

type tagsOpt func(*helixParams)
type tagsOpts struct {
	// TODO
}

var TagsOpts *tagsOpts

func init() {
	TagsOpts = &tagsOpts{
		// TODO
	}
}

// GetAllStreamTags
func (S *Session) GetAllStreamTags(opts ...tagsOpt) *helixResp {
	p := newHelixParams()
	url := newHelixURL("tags/streams", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// GetStreamTags
func (S *Session) GetStreamTags(bcastId string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	url := newHelixURL("streams/tags", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}
