package caliban

import (
	"bufio"
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/oauth2"
)

var cacheDir string

// establish user cache directory
func init() {
	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		log.Fatal(err)
	}
	cacheDir = userCacheDir + "/caliban/"
	err = os.Mkdir(cacheDir, 0600)
	if err != nil && !errors.Is(err, os.ErrExist) {
		log.Fatal(err)
	}
}

const (
	localPort = "11000"
	redirUrl  = "http://localhost:" + localPort + "/caliban/authed"
)

// Session
type Session struct {
	AuthConfig oauth2.Config
	AuthTime   time.Time
	AuthToken  oauth2.Token
	httpClient *http.Client
	Tag        string
}

// NewSession
func NewSession(ctx context.Context, clientId, label string) (*Session, error) {
	// label helps prevent name overlap of cached sessions
	if len(label) == 0 {
		return nil, fmt.Errorf("label cannot be empty")
	}
	S := &Session{Tag: label + "-" + clientId}

	// load cached session if available
	if S.Load(ctx) {
		return S, nil
	}

	// else start from scratch & get user input
	S.AuthConfig.ClientID = clientId
	S.AuthConfig.Endpoint = twitchAuthEPs
	S.AuthConfig.RedirectURL = redirUrl
	S.AuthConfig.ClientSecret = promptInput("Enter your Twitch client-secret: ")

	// TODO: handle scopes properly
	scopes := promptInput("Enter a comma delimited string of requested Twitch scopes: ")
	S.AuthConfig.Scopes = strings.Split(scopes, ",")

	// state
	stateStr := genStateStr(16)

	// gen auth url and print
	authCodeUrl := S.AuthConfig.AuthCodeURL(stateStr)
	fmt.Printf("Follow the Twitch link below authorize caliban:\n\n%v\n\n", authCodeUrl)

	// launch server and wait for results
	stateCh := make(chan string, 1)
	codeCh := make(chan string, 1)
	go serveRedir(stateCh, codeCh)
	retStateStr := <-stateCh
	retCode := <-codeCh

	// confirm correct state
	if retStateStr != stateStr {
		// TODO: handle bad state properly
		log.Fatal("bad sate")
	}

	// exchange code for token
	token, err := S.AuthConfig.Exchange(ctx, retCode)
	if err != nil {
		return S, err
	}

	// register auth time with session
	S.AuthTime = time.Now()

	// TODO: Saving the authToken to Session as we won't have easy access later
	// in order to cache it. If a token expires during use and exchanges are
	// made this could (depending on how Twitch handles tokens?) render this
	// token useless... or I could be mistaken. Either way, this *could* be a
	// problem that requires manual re-auth at some point.
	// See: https://github.com/golang/oauth2/issues/84
	S.AuthToken = *token

	// instantiate client with token and context
	S.httpClient = S.AuthConfig.Client(ctx, token)
	return S, err
}

// validateAuth
func (S *Session) validateAuth() {
	// TODO
	if time.Since(S.AuthTime) > time.Hour {
		req, err := http.NewRequest(http.MethodGet, twitchAuthValidateUrl, nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Add("Client-Id", S.AuthConfig.ClientID)

		resp, err := S.httpClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		rawBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		authResp := &validationResp{}
		err = json.Unmarshal(rawBody, &authResp)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode == 401 {
			log.Fatal("auth invalid, clear user cache and re-auth")
		}
		S.AuthConfig.Scopes = authResp.Scopes
		S.AuthTime = time.Now()
	}
}

// Load checks for a cached Session file, loading the data if present.
func (S *Session) Load(ctx context.Context) bool {
	file, err := os.Open(cacheDir + S.Tag)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		log.Fatal(err)
	}
	defer file.Close()
	dec := gob.NewDecoder(file)
	if err := dec.Decode(S); err != nil {
		log.Fatal(err)
	}
	// instantiate client with recovered token
	S.httpClient = S.AuthConfig.Client(ctx, &S.AuthToken)
	return true
}

// Save writes the current session to a gob file in the user's cache directory.
// The authed user's client id and a label of the user's chosing will be used
// for the file name. Ex. 'TheLabel-z9b0<TheClientId>uuurgq4fffn2o'
func (S *Session) Save() error {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(S)
	sf := cacheDir + S.Tag
	if err = os.WriteFile(sf, buf.Bytes(), os.FileMode(0600)); err != nil {
		return fmt.Errorf("Error writing to file: %w", err)
	}
	return nil
}

// Do forms a request from the provided inputs, adds any needed Headers, then
// calls the client's Do() with the assembled request, returning a new
// helixResp.
func (S *Session) Do(method string, url string, body io.Reader) *helixResp {
	S.validateAuth()
	req, err := http.NewRequest(method, url, body)
	req.Header.Add("Client-Id", S.AuthConfig.ClientID)
	req.Header.Add("Content-Type", "application/json")
	resp, err := S.httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	r := &helixResp{}
	err = json.Unmarshal(rawBody, r)
	if err != nil {
		log.Fatal(err)
	}
	r.StatusCode = resp.StatusCode
	r.Status = resp.Status
	return r
}

// HasScope checks if the current Session (authorized id/token) has been granted
// the given scope.
func (S *Session) HasScope(scope string) bool {
	for _, v := range S.AuthConfig.Scopes {
		if v == scope {
			return true
		}
	}
	return false
}

// Identity returns any relevant information we have about the authenticated
// user. Currently a place-holder waiting to be implemented.
func (S *Session) Identity() string {
	s := "Twitch ClientId: %s\n"
	return fmt.Sprintf(s, S.AuthConfig.ClientID)
}

// PrettyString returns a 'pretty' string representation of result R.
func (R *helixResp) PrettyString() string {
	pj := &bytes.Buffer{}
	if R.Data != nil {
		if err := json.Indent(pj, R.Data, "", "    "); err != nil {
			log.Fatal(err)
		}
	}
	fmtStr := "\nStatus: %s\n\n  Data:\n%s\n\nCursor: %s\n\n"
	return fmt.Sprintf(fmtStr, R.Status, pj, R.Pagination["cursor"])
}

// serveRedir serves a local redirect page for auth token retreival.
func serveRedir(s chan string, c chan string) {
	http.HandleFunc("/caliban/authed",
		func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			s <- r.Form.Get("state")
			c <- r.Form.Get("code")
		})
	err := http.ListenAndServe(":"+localPort, nil)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

// genStateStr returns a sudo-random string of length n.
func genStateStr(n int) string {
	const pool = `0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
	bs := make([]byte, n)
	for i := range bs {
		bs[i] = pool[rand.Intn(len(pool))]
	}
	return string(bs)
}

// promptInput
func promptInput(prompt string) string {
	var out string
	var err error
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, prompt)
		out, err = r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if out != "" {
			break
		}
	}
	return strings.TrimSpace(out)
}
