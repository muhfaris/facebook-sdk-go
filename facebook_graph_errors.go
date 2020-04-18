package facebook

import "fmt"

// FBGraphError is wrap error response
type FBGraphError struct {
	Error FBGraphErrorData `json:"error"`
}

// FBGraphErrorData is body error data
type FBGraphErrorData struct {
	Message      string `json:"message"`
	Type         string `json:"type"`
	Code         int    `json:"code"`
	ErrorSubcode int    `json:"error_subcode"` // subcode for authentication related errors.
	UserTitle    string `json:"error_user_title"`
	UserMessage  string `json:"error_user_msg"`
	IsTransient  bool   `json:"is_transient"`
	TraceID      string `json:"fbtrace_id"`
}

const TypeFBSDKError = "FBSDKError"

// CreateError is create error to facebook format
func CreateError(msg string, err error) FBGraphError {
	return FBGraphError{
		Error: FBGraphErrorData{
			Message: fmt.Sprintf("%s - %v", msg, err),
			Type:    TypeFBSDKError,
		},
	}
}
