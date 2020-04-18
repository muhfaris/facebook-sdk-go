package facebook

import (
	"github.com/muhfaris/request"
)

func (app *App) get(url string, params map[string]string) FBResponse {
	auth := app.BearerToken()
	req, err := request.New(url, "application/json", auth, "", params)
	if err != nil {
		return FBResponse{Error: CreateError("Pre-request to facebook", err)}
	}

	resp, err := req.GET()
	if err != nil {
		return FBResponse{Error: CreateError("GET request to facebook", err)}
	}

	return FBResponse{Response: resp.Response, Result: resp.Body}
}
