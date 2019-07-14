package facebook

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type HttpClient interface {
	Do(req *http.Request) (resp *http.Response, err error)
	Get(url string) (resp *http.Response, err error)
	Post(url string, bodyType string, body io.Reader) (resp *http.Response, err error)
}

func (app *App) GET(path string, params Params) (Response, error) {
	URIquery := app.prepareQuery(path, params)
	log.Println(URIquery)
	request, err := http.NewRequest(http.MethodGet, URIquery, nil)
	if err != nil {
		return nil, err
	}

	response, err := app.request(request)
	log.Println(string(response), err)
	if err != nil {
		return nil, err
	}

	checkErr := newError(response)
	if checkErr.isError() {
		return nil, checkErr.Error()
	}

	var resp Response
	json.Unmarshal(response, &resp)

	return resp, nil
}

func (app *App) POST(path string, params Params) {
	//return graph(path, http.MethodPost, params)
}

func (app *App) DELETE(path string, params Params) {
	//	return graph(path, http.MethodDelete, params)
}

func graph(path, method string, params Params) {

}

func (app *App) request(r *http.Request) ([]byte, error) {
	r.Header.Set("Content-Type", "application/json")
	r.Header.Add("Authorization", app.token)

	resp, err := app.httpClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return data, errors.New(ErrorCantRequestFacebook)
	}

	return data, nil
}

func setGetQuery(params Params) string {
	tempQuery := url.Values{}
	for i, j := range params {
		tempQuery.Add(i, j.(string))
	}

	return tempQuery.Encode()
}

func (app *App) prepareQuery(path string, params Params) string {
	return fmt.Sprintf(
		"%s/%s/%s?%s",
		APIFacebook,
		DefaultVersion,
		path,
		setGetQuery(params),
	)
}
