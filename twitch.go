package caliban

import (
	"encoding/json"
	"net/url"

	"golang.org/x/oauth2"
)

// twitch urls/host
const (
	twitchApiHost         = "api.twitch.tv"
	twitchAuthValidateUrl = "https://id.twitch.tv/oauth2/validate"
)

// videoTypes
const (
	archiveVideos videoType = iota + 1
	highlightVideos
	uploadVideos
)

// timePeriods
const (
	lastDay timePeriod = iota + 1
	lastMonth
	lastWeek
)

// twitch authentication endpoints
var twitchAuthEPs = oauth2.Endpoint{
	AuthURL:  "https://id.twitch.tv/oauth2/authorize",
	TokenURL: "https://id.twitch.tv/oauth2/token",
}

type (
	// body params for moderation endpoint
	// AutoModData struct {
	// 	MsgId   string `json:"msg_id,omitempty"`
	// 	MsgText string `json:"msg_text,omitempty"`
	// }
	// AutoModSettings struct {
	// 	BroadcasterId           string `json:"broadcaster_id,omitempty"`
	// 	ModeratorId             string `json:"moderator_id,omitempty"`
	// 	OverallLevel            int    `json:"overall_level,omitempty"`
	// 	Disability              int    `json:"disability,omitempty"`
	// 	Aggression              int    `json:"aggression,omitempty"`
	// 	SexualitySexOrGender    int    `json:"sexuality_sex_or_gender,omitempty"`
	// 	Misogyny                int    `json:"misogyny,omitempty"`
	// 	Bullying                int    `json:"bullying,omitempty"`
	// 	Swearing                int    `json:"swearing,omitempty"`
	// 	RaceEthnicityOrReligion int    `json:"race_ethnicity_or_religion,omitempty"`
	// 	SexBasedTerms           int    `json:"sex_based_terms,omitempty"`
	// }
	// parameters for helix api
	helixParams struct {
		body map[string]interface{}
		url.Values
	}
	// general tiwtch response
	helixResp struct {
		Status     string            `json:"status,omitempty"`
		StatusCode int               `json:"status_code,omitempty"`
		Data       json.RawMessage   `json:"data,omitempty"`
		Pagination map[string]string `json:"pagination,omitempty"`
	}
	// url for helix api
	helixURL struct{ url.URL }
	// timePeriod
	timePeriod uint8
	// auth validation response
	validationResp struct {
		ClientId  string   `json:"client_id,omitempty"`
		Login     string   `json:"login,omitempty"`
		Scopes    []string `json:"scopes,omitempty"`
		UserId    string   `json:"user_id,omitempty"`
		ExpiresIn int      `json:"expires_in,omitempty"`
		Status    int      `json:"status,omitempty"`
		Message   string   `json:"message,omitempty"`
	}
	// videoType
	videoType uint8
)

// newHelixParams
func newHelixParams() *helixParams {
	return &helixParams{
		map[string]interface{}{},
		url.Values{},
	}
}

// newHelixURL
func newHelixURL(endPoint, rawQuery string) *helixURL {
	return &helixURL{
		url.URL{
			Scheme:   "https",
			Host:     twitchApiHost,
			Path:     "helix/" + endPoint,
			RawQuery: rawQuery,
		},
	}
}

// timePeriod stringer
func (t timePeriod) String() string {
	return [...]string{"day", "month", "week"}[t-1]
}

// videoType stringer
func (v videoType) String() string {
	return [...]string{"archive", "highlight", "upload"}[v-1]
}
