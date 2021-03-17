package sdk

import (
	"encoding/json"
	"fmt"
	"net/url"
)

const endpoint = "https://graph.facebook.com"

// FacebookAPI interface
type FacebookAPI interface {
	ChangeToken(string) error
	ChangeGraph(int)

	Authenticate(state string, scopes ...string) *Authenticate
	Get(string, ...requestOptions) Response
	Post(string, ...requestOptions) Response
	Delete(string, ...requestOptions) Response

	Batch(ArrayOfBatchBodyRequest) BatchResponse
}

// Facebook is wrap facebook data
type Facebook struct {
	AppID       string // AppID
	AppKey      string
	Debug       string
	Version     string
	Token       string
	RedirectURL string
	Graph       int
}

// NewFacebook is create new object facebook
func NewFacebook(fb Facebook) (FacebookAPI, error) {
	return NewGraph(
		GWithURL(fb.url()),
		GWithAppID(fb.AppID),
		GWithAppKey(fb.AppKey),
		GWithToken(fb.Token),
		GWithRedirectURL(fb.RedirectURL),
		GWithGraph(fb.Graph),
		gSecretProof(fb.AppKey, fb.Token),
	), nil
}

func (f Facebook) url() string {
	return fmt.Sprintf("%s/%s", endpoint, f.Version)
}

func (f Facebook) body() ([]byte, error) {
	return json.Marshal(f.body)
}

// ParamQuery is for url query
type ParamQuery map[string]interface{}

// DebugType is type of debug
func (pq ParamQuery) DebugType() []string {
	return []string{
		"all", "info", "warning",
	}
}

// FindDebugType is type of debug
func (pq ParamQuery) FindDebugType(code string) (string, error) {
	for _, dt := range pq.DebugType() {
		if dt == code {
			return dt, nil
		}
	}

	return "", fmt.Errorf("error debug type invalid")
}

// EncodeURL to url
func (pq ParamQuery) EncodeURL(debug string) string {
	var URLQuery = url.Values{}
	for keyQuery, valueQuery := range pq {
		URLQuery.Add(keyQuery, valueQuery.(string))
	}

	dbg, _ := pq.FindDebugType(debug)
	if dbg != "" {
		URLQuery.Add("debug", dbg)
	}

	return URLQuery.Encode()
}
