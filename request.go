package facebook

import (
	"net/http"

	"github.com/muhfaris/request"
)

//validate is validate data
func (config *Config) validate(method, path string) error {
	if ok := HTTPMethodsValidate(method); !ok {
		return AppError{Code: AppErrorInvalidMethod}
	}

	if config.Version == "" {
		return AppError{Code: AppErrorInvalidVersion}
	}

	if config.Token == "" {
		return AppError{Code: AppErrorInvalidToken}
	}

	// setup
	config.generateAPI(path)
	config.setMethod(method)
	config.setToken()

	return nil
}

// MarketingAPI is request marketing facebook api
func (config *Config) MarketingAPI(path string, method string, params map[string]string) FBResponse {
	err := config.validate(method, path)
	if err != nil {
		return FBResponse{Error: err}
	}

	return config.send(params)
}

func (config *Config) send(params map[string]string) FBResponse {
	switch config.method {
	case http.MethodGet:
		req := request.ReqApp{
			URL:           config.api,
			ContentType:   request.MimeTypeJSON,
			Authorization: config.Token,
			QueryString:   params,
		}

		response, err := req.GET()
		if err != nil {
			return FBResponse{Error: CreateError("GET request to facebook", err)}
		}
		return FBResponse{Response: response.Response, Result: response.Body}

	}

	return FBResponse{}
}
