package facebook

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FBResponse is wrap response facebook
type FBResponse struct {
	Response *http.Response
	Result   interface{}
	Error    interface{}
}

// Byte response to byte
func (fr FBResponse) Byte() ([]byte, error) {
	return json.Marshal(fr.Result)
}

// FBResponseGraph is response graph api
type FBResponseGraph struct {
}

// Graph is extract data for graph response
func (fr FBResponse) Graph() FBGraphError {
	return fr.Error.(FBGraphError)
}

// FBResponseMarketing is response graph api
type FBResponseMarketing struct {
}

// Marketing is extract data for graph response
func (fr FBResponse) Marketing() FBResponseMarketing {
	return FBResponseMarketing{}
}

// Request is request API
func (app *App) Request(ty, path, method string, params map[string]string) FBResponse {
	if ok := validateAPI(ty); !ok {
		return FBResponse{Error: "Type of api not valid"}
	}

	if ok := HTTPMethodsValidate(method); !ok {
		return FBResponse{Error: "Http method invalid"}
	}

	url := app.buildURL(path)

	switch ty {
	case MarketingAPI:
		return app.Marketing(url, method, params)
	default:
		// GraphAPI:
		return app.Graph(url, method, params)
	}
}

// ValidateAPI is validation type of api
func validateAPI(ty string) bool {
	for _, api := range APITypes {
		if api == ty {
			return true
		}
	}
	return false
}

func (app *App) buildURL(path string) string {
	base := app.BaseURL()
	return fmt.Sprintf("%s/%s", base, path)
}

// BaseURL is create base url
func (app *App) BaseURL() string {
	return fmt.Sprintf("%s/%s", FacebookAPI, app.version)
}

// BearerToken is format token facebook
func (app *App) BearerToken() string {
	return fmt.Sprintf("Bearer %s", app.token)
}
