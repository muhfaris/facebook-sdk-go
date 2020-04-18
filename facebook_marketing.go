package facebook

import (
	"encoding/json"
	"net/http"
)

// Marketing is graph facebook request
func (app *App) Marketing(url, method string, params map[string]string) FBResponse {
	switch method {
	case http.MethodPost:

	default:
		// method GET
		return app.responseAnomaliFBM(app.get(url, params))
	}
	return FBResponse{}
}

func (app *App) responseAnomaliFBM(resp FBResponse) FBResponse {
	if resp.Response.StatusCode != http.StatusOK {
		var fbError FBMarketingError
		err := json.Unmarshal(resp.Result.([]byte), &fbError)
		if err != nil {
			return FBResponse{Error: CreateError("Unmarshal response body facebook", err)}
		}
		return FBResponse{Error: fbError}
	}

	return resp
}
