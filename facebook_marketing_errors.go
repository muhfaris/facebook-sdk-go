package facebook

import "fmt"

// FBMarketingError is wrap error response
type FBMarketingError struct {
	Data FBMarketingErrorWrap `json:"data"`
}

// FBMarketingErrorWrap is error response
type FBMarketingErrorWrap struct {
	Error FBMarketingErrorData `json:"error"`
}

// FBMarketingErrorData is body error data
type FBMarketingErrorData struct {
	Message      string `json:"message"`
	Type         string `json:"type"`
	Code         int    `json:"code"`
	ErrorSubcode int    `json:"error_subcode"` // subcode for authentication related errors.
	UserTitle    string `json:"error_user_title"`
	UserMessage  string `json:"error_user_msg"`
	IsTransient  bool   `json:"is_transient"`
	TraceID      string `json:"fbtrace_id"`
}

// CreateErrorMarketing is create error to facebook format
func CreateErrorMarketing(msg string, err error) FBMarketingError {
	return FBMarketingError{
		Data: FBMarketingErrorWrap{
			Error: FBMarketingErrorData{
				Message: fmt.Sprintf("%s - %v", msg, err),
				Type:    TypeFBSDKError,
			},
		},
	}
}
