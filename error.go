package facebook

import (
	"encoding/json"
	"errors"
)

type (
	Error struct {
		Message      string
		Type         string
		Code         int
		ErrorSubcode int    // subcode for authentication related errors.
		UserTitle    string `json:"error_user_title"`
		UserMessage  string `json:"error_user_msg"`
		IsTransient  bool   `json:"is_transient"`
		TraceID      string `json:"fbtrace_id"`
	}
)

func (e *Error) Msg() string {
	return e.Message
}

func (e *Error) Error() error {
	return errors.New(e.Message)
}

func (e *Error) isError() bool {
	return (e.TraceID != "")
}

func newError(response []byte) *Error {
	var error Error
	if err := json.Unmarshal(response, &error); err != nil {
		return &error
	}

	return &error
}
