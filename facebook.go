package facebook

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type (
	App struct {
		Url               string
		Token             string
		Version           string
		AppSecretKey      string
		IsAppSecretProof  bool
		AppSecretProofKey string
		Debug             string
		context           context.Context
		HttpClient        http.Client
	}
)

type HttpClient interface {
	Do(req *http.Request) (resp *http.Response, err error)
	Get(url string) (resp *http.Response, err error)
	Post(url string, bodyType string, body io.Reader) (resp *http.Response, err error)
}

func New(token, appSecret string) *App {
	return &App{
		Url:          DefaultURL,
		Token:        token,
		AppSecretKey: appSecret,
		Version:      DefaultVersion,
	}
}

// If facebook returns error in response, returns error details in res and set err.
func (session *App) Api(path string, method Method, params Params) (Result, *ErrorResponse) {
	return session.grap(path, method, params)
}

// Get is a short hand of Api(path, GET, params).
func (session *App) Get(path string, params Params) (Result, *ErrorResponse) {
	return session.Api(path, GET, params)
}

// Post is a short hand of Api(path, POST, params).
func (session *App) Post(path string, params Params) (Result, *ErrorResponse) {
	return session.Api(path, POST, params)
}

// Delete is a short hand of Api(path, DELETE, params).
func (session *App) Delete(path string, params Params) (Result, *ErrorResponse) {
	return session.Api(path, DELETE, params)
}

func (session *App) prepareRequest(path, query string) (token string, url string) {
	url = session.setApi(path, query)
	token = session.setToken()
	return
}

func (session *App) prepareParams(params Params) string {
	tempQuery := url.Values{}
	for i, j := range params {
		tempQuery.Add(i, j.(string))
	}

	if isEmptyString(session.AppSecretProofKey) {
		tempQuery.Add(TextSecretProof, session.AppSecretProofKey)
	}

	tempQuery.Add(TextDebug, session.Debug)

	return tempQuery.Encode()
}

func (session *App) setApi(path, query string) string {
	return fmt.Sprintf("%s/%s%s?%s", session.Url, session.Version, path, query)
}

func (session *App) setToken() string {
	return fmt.Sprintf("%s %s", TextBearer, session.Token)
}

func (session *App) setSecretProof() {
	if session.IsAppSecretProof {
		session.AppSecretProofKey = GenerateSecretProof(session.Token, session.AppSecretKey)
		return
	}

	session.AppSecretProofKey = ""
}

func (session *App) grap(path string, method Method, params Params) (Result, *ErrorResponse) {
	session.setSecretProof()
	query := session.prepareParams(params)
	token, url := session.prepareRequest(path, query)

	switch {
	case method == "GET":
		data, err := session.sendGetRequest(url, token, "GET")
		return data, err
	}
	return nil, nil
}

func (session *App) sendGetRequest(url, token string, method Method) (Result, *ErrorResponse) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, NewError(err, ErrorCantRequestFacebook, http.StatusInternalServerError)
	}

	_, data, errRes := session.sendRequest(token, request)

	if errRes != nil {
		return nil, errRes
	}

	return data, nil
}

func (session *App) sendRequest(token string, req *http.Request) (*http.Response, Result, *ErrorResponse) {
	var result map[string]interface{}

	if session.context != nil {
		req = req.WithContext(session.context)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	resp, err := session.HttpClient.Do(req)

	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, resp.Body)
	data := buf.Bytes()

	if err != nil {
		return resp, nil, NewError(err, ErrorCantRequestFacebook, http.StatusInternalServerError)
	}

	errResp, isError := isErrorResponse(data, resp)
	if isError {
		return resp, nil, &errResp
	}

	json.Unmarshal(data, &result)
	return resp, result, nil
}
