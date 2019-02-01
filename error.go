package facebook

type (
	ErrorResponse struct {
		Data struct {
			Code      int    `json:"code"`
			FbtraceID string `json:"fbtrace_id"`
			Message   string `json:"message"`
			Type      string `json:"type"`
			Error     error
		} `json:"error"`
	}
)

func (e *ErrorResponse) Message() string {
	return e.Data.Message
}

func (e *ErrorResponse) Errors() error {
	return e.Data.Error
}

func NewError(err error, msg string, status int) *ErrorResponse {
	return &ErrorResponse{}
}

func NewErrorWrapf(err error, prefix, suffix, msg string, status int) *ErrorResponse {
	return &ErrorResponse{}
}
