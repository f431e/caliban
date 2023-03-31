package caliban

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type usersOpt func(*helixParams)
type usersOpts struct {
	ByLogin func(...string) usersOpt
	ById    func(...string) usersOpt
	Count   func(int) usersOpt
	Next    func(int, string) usersOpt
	Source  func(string) usersOpt
	Reason  func(string) usersOpt
}

var UsrOpts *usersOpts

func init() {
	UsrOpts = &usersOpts{
		ByLogin: func(logins ...string) usersOpt {
			return func(p *helixParams) {
				for _, login := range logins {
					p.Add("login", login)
				}
			}
		},
		ById: func(ids ...string) usersOpt {
			return func(p *helixParams) {
				for _, id := range ids {
					p.Add("id", id)
				}
			}
		},
		Count: func(n int) usersOpt {
			return func(p *helixParams) {
				p.Add("first", strconv.Itoa(n))
			}
		},
		Next: func(n int, cursor string) usersOpt {
			return func(p *helixParams) {
				p.Add("first", strconv.Itoa(n))
				p.Add("after", cursor)
			}
		},
		Source: func(source string) usersOpt {
			return func(p *helixParams) {
				p.Add("source_context", source)
			}
		},
		Reason: func(reason string) usersOpt {
			return func(p *helixParams) {
				p.Add("reason", reason)
			}
		},
	}
}

// GetUsers
func (S *Session) GetUsers(opts ...usersOpt) *helixResp {
	p := newHelixParams()
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("users", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// UpdateUser
func (S *Session) UpdateUser(description string) *helixResp {
	p := newHelixParams()
	p.Add("description", description)
	url := newHelixURL("users", p.Encode())
	return S.Do(http.MethodPut, url.String(), nil)
}

// GetUsersFollows
// Deprecation: This endpoint is deprecated and will be decommissioned
// on August 3, 2023. Access to this endpoint is limited to client IDs
// that have called the endpoint on or before February 17, 2023.
func (S *Session) GetUsersFollows(opts ...usersOpt) *helixResp {
	p := newHelixParams()
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("users/follows", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// GetUserBlockList
func (S *Session) GetUserBlockList(bcastId string, opts ...usersOpt) *helixResp {
	p := newHelixParams()
	p.Add("broadcaster_id", bcastId)
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("users/blocks", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// BlockUser
func (S *Session) BlockUser(targetUserId string, opts ...usersOpt) *helixResp {
	p := newHelixParams()
	p.Add("target_user_id", targetUserId)
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("users/blocks", p.Encode())
	return S.Do(http.MethodPut, url.String(), nil)
}

// UnblockUser
func (S *Session) UnblockUser(targetUserId string) *helixResp {
	p := newHelixParams()
	p.Add("target_user_id", targetUserId)
	url := newHelixURL("users/blocks", p.Encode())
	return S.Do(http.MethodDelete, url.String(), nil)
}

// GetUserExtensions
func (S *Session) GetUserExtensions() *helixResp {
	url := newHelixURL("users/extensions/list", "")
	return S.Do(http.MethodGet, url.String(), nil)
}

// GetUserActiveExtensions
func (S *Session) GetUserActiveExtensions(opts ...usersOpt) *helixResp {
	p := newHelixParams()
	for _, opt := range opts {
		opt(p)
	}
	url := newHelixURL("users/extensions", p.Encode())
	return S.Do(http.MethodGet, url.String(), nil)
}

// UpdateUserExtensions
func (S *Session) UpdateUserExtensions(opts ...usersOpt) *helixResp {
	p := newHelixParams()
	for _, opt := range opts {
		opt(p)
	}
	rawBody, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	url := newHelixURL("users/extensions", "")
	return S.Do(http.MethodPut, url.String(), bytes.NewBuffer(rawBody))
}
