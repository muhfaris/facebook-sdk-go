package facebook

import (
	"encoding/json"
	"log"
	"net/http"
)

// Graph is graph facebook request
func (app *App) Graph(url, method string, params map[string]string) FBResponse {
	switch method {
	case http.MethodPost:

	default:
		// method GET
		return app.responseAnomaliFB(app.get(url, params))
	}
	return FBResponse{}
}

func (app *App) responseAnomaliFB(resp FBResponse) FBResponse {
	if resp.Response.StatusCode != http.StatusOK {
		var fbError FBGraphError
		err := json.Unmarshal(resp.Result.([]byte), &fbError)
		if err != nil {
			return FBResponse{Error: CreateError("Unmarshal response body facebook", err)}
		}
		return FBResponse{Error: fbError}
	}

	log.Println(string(resp.Result.([]byte)))
	return resp
}
