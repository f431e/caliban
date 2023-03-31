package caliban

import (
	"net/http"
)

type musicOpt func(*helixParams)
type musicOpts struct {
	// TODO
}

var MusicOpts *musicOpts

func init() {
	MusicOpts = &musicOpts{
		// TODO
	}
}

// GetSoundtrackCurrentTrack
func (S *Session) GetSoundtrackCurrentTrack(bcastId string) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	url := newHelixURL("soundtrack/current_track", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// GetSoundtrackPlaylist
func (S *Session) GetSoundtrackPlaylist(id string, opts ...musicOpt) *helixResp {
	p := newHelixParams()
	p.Add("id", id)
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("soundtrack/playlist", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// GetSoundtrackPlaylists
func (S *Session) GetSoundtrackPlaylists(opts ...musicOpt) *helixResp {
	p := newHelixParams()
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("soundtrack/playlists", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}
